package preinstall

import (
	"net/http"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/web"
	"github.com/micro/micro/plugin"
	"github.com/yametech/fuxi/pkg/kubernetes/clientv1"
	"github.com/yametech/fuxi/pkg/kubernetes/clientv2"
	"github.com/yametech/fuxi/thirdparty/lib/token"
	"github.com/yametech/fuxi/thirdparty/lib/whitelist"
	"github.com/yametech/fuxi/thirdparty/lib/wrapper/auth"
	"github.com/yametech/fuxi/util/common"
	k8sjson "k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	clientcmdapiV1 "k8s.io/client-go/tools/clientcmd/api/v1"
)

// defaultETCDFlag a etcd String Flag
// dev Value:  "gz.nuwa.xyz:32428",
// value = "sdmssd.io:2379"
// value = "gz.nuwa.xyz:32428"
// value = "linux:30755"
func defaultETCDFlag() cli.StringFlag {
	flag := cli.StringFlag{
		Name:   "etcd_address",
		Usage:  "etcd address for config K/V",
		EnvVar: "ETCD_ADDRESS",
		Value:  "etcd.kube-system.svc.cluster.local:2379",
	}
	return flag
}

func inClusterFlag() cli.StringFlag {
	flag := cli.StringFlag{
		Name:   "in_cluster",
		Usage:  "in_cluster=true",
		EnvVar: "IN_CLUSTER",
		Value:  "",
	}
	return flag
}

//InitGateWay init a gateway
func InitGatewayInstallConfigure(name string, loginHandle http.Handler, microPlugins ...plugin.Plugin) (*GateWayInstallConfigure, error) {
	gwic := &GateWayInstallConfigure{
		Token:     &token.Token{},
		Whitelist: &whitelist.Whitelist{},
	}
	err := plugin.Register(
		plugin.NewPlugin(
			plugin.WithName("auth"),
			plugin.WithHandler(
				auth.JWTAuthWrapper(gwic.Token, gwic.Whitelist, loginHandle),
			),
			plugin.WithFlag(
				defaultETCDFlag(),
				inClusterFlag(),
			),
			plugin.WithInit(func(ctx *cli.Context) error {
				defaultInstallConfigure, err :=
					NewDefaultInstallConfigure(
						ctx.String("etcd_address"),
						ctx.String("in_cluster"),
					)
				if err != nil {
					return err
				}
				gwic.DefaultInstallConfigure = *defaultInstallConfigure

				_whileList, err := whitelist.InitConfig(gwic.SystemConfigServer, "go", "micro", "urls", "list")
				if err != nil {
					return err
				}
				gwic.Whitelist = _whileList

				_token, err := token.InitConfig(gwic.SystemConfigServer, "go", "micro", "urls", "list")
				if err != nil {
					return err
				}
				gwic.Token = _token

				return nil
			}),
		))
	if err != nil {
		return nil, err
	}

	for _, microPlugin := range microPlugins {
		if err := plugin.Register(microPlugin); err != nil {
			return nil, err
		}
	}
	return gwic, nil
}

//InitApi init a api
func InitApi(sampling int, name, version, tracingAddr string) (web.Service, *ApiInstallConfigure, error) {
	var apiInstallConfigure *ApiInstallConfigure
	service := web.NewService(
		web.Name(name),
		web.Version(common.Version(version)),
		web.RegisterTTL(time.Second*15),
		web.RegisterInterval(time.Second*10),
		web.Flags(
			defaultETCDFlag(),
			inClusterFlag(),
		),
		web.Action(func(ctx *cli.Context) {
			defaultInstallConfigure, err :=
				NewDefaultInstallConfigure(
					ctx.String("etcd_address"),
					ctx.String("in_cluster"),
				)
			if err != nil {
				panic(err)
			}
			apiInstallConfigure = &ApiInstallConfigure{
				DefaultInstallConfigure: *defaultInstallConfigure,
			}
		}),
	)
	if err := service.Init(); err != nil {
		return nil, nil, err
	}
	return service, apiInstallConfigure, nil
}

// InitService init a Service
func InitService(name, version string) (micro.Service, *ApiInstallConfigure) {
	var apiInstallConfigure *ApiInstallConfigure
	service := micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Flags(
			defaultETCDFlag(),
			inClusterFlag(),
		),
		micro.Action(func(ctx *cli.Context) {
			defaultInstallConfigure, err :=
				NewDefaultInstallConfigure(
					ctx.String("etcd_address"),
					ctx.String("in_cluster"),
				)
			if err != nil {
				panic(err)
			}
			apiInstallConfigure = &ApiInstallConfigure{
				DefaultInstallConfigure: *defaultInstallConfigure,
			}
		}),
	)
	// Initialise service
	service.Init()
	return service, apiInstallConfigure
}

type GateWayInstallConfigure struct {
	DefaultInstallConfigure
	Token     *token.Token
	Whitelist *whitelist.Whitelist
}

type ApiInstallConfigure struct {
	DefaultInstallConfigure
}

type ServiceInstallConfigure struct {
	DefaultInstallConfigure
}

type K8sConfig struct {
	Name   string                `json:"name"`
	Config clientcmdapiV1.Config `json:"config"`
}

type DefaultInstallConfigure struct {
	// configure server address
	Addr string
	// system config
	SystemConfig config.Config
	// custom system config server for etcd
	SystemConfigServer *common.ConfigServer
	//K8sConfig          *K8sConfig
	// kubernetes
	RestConfig *rest.Config
	ClientV1   *kubernetes.Clientset
	ClientV2   *clientv2.CacheInformerFactory
}

func NewDefaultInstallConfigure(addr string, mode string) (*DefaultInstallConfigure, error) {
	systemConfig, systemConfigServer, err := createSystemConfig(addr)
	if err != nil {
		return nil, err
	}
	defaultInstallConfigure := &DefaultInstallConfigure{
		Addr:               addr,
		SystemConfig:       systemConfig,
		SystemConfigServer: systemConfigServer,
	}
	if mode != "" {
		var err error
		ClientV1, restConf, err := createInClusterConfig()
		if err != nil {
			return nil, err
		}
		defaultInstallConfigure.ClientV1 = ClientV1
		defaultInstallConfigure.RestConfig = restConf

		clientV2, err := clientv2.NewCacheInformerFactory(restConf.ServerName, restConf, nil)
		if err != nil {
			return nil, err
		}
		defaultInstallConfigure.ClientV2 = clientV2

		return defaultInstallConfigure, nil
	}

	k8sBytes := systemConfig.Get("go", "micro", "kubernetes", "k8s").Bytes()
	k8sConfig, err := createKubernetesJsonConfig(k8sBytes)
	if err != nil {
		return nil, err
	}
	//defaultInstallConfigure.K8sConfig = k8sConfig
	ClientV1, restConf, err := clientv1.BuildClient(k8sConfig.Name, k8sConfig.Config)
	if err != nil {
		return nil, err
	}
	defaultInstallConfigure.ClientV1 = ClientV1
	defaultInstallConfigure.RestConfig = restConf

	clientV2, err := clientv2.NewCacheInformerFactory(k8sConfig.Name, nil, &k8sConfig.Config)
	if err != nil {
		return nil, err
	}
	defaultInstallConfigure.ClientV2 = clientV2

	return defaultInstallConfigure, nil
}

func createInClusterConfig() (*kubernetes.Clientset, *rest.Config, error) {
	restConfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, nil, err
	}
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, nil, err
	}
	return clientSet, restConfig, nil
}

func createSystemConfig(addr string) (config.Config, *common.ConfigServer, error) {
	configureServer := common.NewConfigServer(addr, common.ConfigPrefix)
	config := config.NewConfig()
	if err := config.Load(configureServer); err != nil {
		return nil, nil, err
	}
	return config, configureServer, nil
}

func createKubernetesJsonConfig(data []byte) (*K8sConfig, error) {
	config := K8sConfig{}
	err := k8sjson.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
