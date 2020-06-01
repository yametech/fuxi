package preinstall

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/web"
	"github.com/micro/micro/plugin"
	"github.com/yametech/fuxi/pkg/kubernetes/clientv1"
	"github.com/yametech/fuxi/pkg/kubernetes/clientv2"
	"github.com/yametech/fuxi/thirdparty/lib/token"
	"github.com/yametech/fuxi/thirdparty/lib/whitelist"
	"github.com/yametech/fuxi/util/common"
	k8sjson "k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	clientcmdapiV1 "k8s.io/client-go/tools/clientcmd/api/v1"
	"time"
)

func NewGoMicroPlugin(name string, handlers []plugin.Handler, etcdFlag cli.StringFlag, init func(*cli.Context) error) plugin.Plugin {
	return plugin.NewPlugin(
		plugin.WithName(name),           // "auth"
		plugin.WithHandler(handlers...), //auth.JWTAuthWrapper(gateWayInstall.token, gateWayInstall.whitelist)
		plugin.WithFlag(etcdFlag),
		plugin.WithInit(init),
		//func(ctx *cli.Context) error {
		//	conf := common.NewConfigServer(
		//		ctx.String("etcd_address"),
		//		common.ConfigPrefix,
		//	)
		//	if err := gateWayInstall.whitelist.InitConfig(conf, "go", "micro", "urls", "list"); err != nil {
		//		return err
		//	}
		//	return gateWayInstall.token.InitConfig(conf, "go", "micro", "jwt", "key")
		//}
	)
}

type GateWayInstall struct {
	token     *token.Token
	whitelist *whitelist.Whitelist
}

//InitGateWay init a gateway
func InitGateWayInstall(microPlugins ...plugin.Plugin) (*GateWayInstall, error) {
	gateWayInstall := &GateWayInstall{
		&token.Token{},
		&whitelist.Whitelist{},
	}

	for _, microPlugin := range microPlugins {
		if err := plugin.Register(microPlugin); err != nil {
			return nil, err
		}
	}
	return gateWayInstall, nil
}

// EtcdStringFlag a Etcd String Flag
func EtcdStringFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:   "etcd_address",
		Usage:  "etcd address for config K/V",
		EnvVar: "ETCD_ADDRESS",
		//Value:  "gz.nuwa.xyz:32428",
		Value: "fuxi.io:12379",
	}
}

//--registry_address 10.200.100.200:2379 fuxi.io

//todo:考虑把参数封装成ServiceOption struct
//todo: RegisterTTL RegisterInterval 也基于参数的方式配置进来
//InitApi init a api
func InitApi(sampling int, name, version, tracingAddr string) (web.Service, *token.Token, *ApiInstallConfigure, error) {
	// New Service
	token := &token.Token{}
	//gin2micro.SetSamplingFrequency(sampling)
	//t, io, err := tracer.NewTracer(name, tracingAddr)
	//if err != nil {
	//	return web.NewService(), token, err
	//}
	//defer func() {
	//	if err := io.Close(); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//opentracing.SetGlobalTracer(t)

	var apiInstallConfigure *ApiInstallConfigure
	service := web.NewService(
		web.Name(name),
		web.Version(common.Version(version)),
		web.RegisterTTL(time.Second*15),
		web.RegisterInterval(time.Second*10),
		web.Flags(EtcdStringFlag()),
		web.Action(func(ctx *cli.Context) {
			defaultInstallConfigure, err := NewDefaultInstallConfigure(ctx.String("etcd_address"))
			if err != nil {
				panic(err)
			}
			apiInstallConfigure = &ApiInstallConfigure{
				DefaultInstallConfigure: *defaultInstallConfigure,
			}
			token.InitConfig(apiInstallConfigure.SystemConfigServer, "go", "micro", "jwt", "key")
		}),
	)

	if err := service.Init(); err != nil {
		return nil, nil, nil, err
	}

	return service, token, apiInstallConfigure, nil
}

// InitService init a Service
func InitService(name, version string) (micro.Service, *ApiInstallConfigure) {
	var apiInstallConfigure *ApiInstallConfigure
	service := micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		//micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.Flags(EtcdStringFlag()),
		micro.Action(func(ctx *cli.Context) {
			defaultInstallConfigure, err := NewDefaultInstallConfigure(ctx.String("etcd_address"))
			if err != nil {
				panic(err)
			}
			apiInstallConfigure = &ApiInstallConfigure{
				DefaultInstallConfigure: *defaultInstallConfigure,
			}
		}),
	)
	return service, apiInstallConfigure
}

type GateWayInstallConfigure struct {
	DefaultInstallConfigure
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
	K8sConfig          *K8sConfig
	// kubernetes
	RestConfig *rest.Config
	ClientV1   *kubernetes.Clientset
	ClientV2   *clientv2.CacheInformerFactory
}

func NewDefaultInstallConfigure(addr string) (*DefaultInstallConfigure, error) {
	systemConfig, systemConfigServer, err := createSystemConfig(addr)
	if err != nil {
		return nil, err
	}
	defaultInstallConfigure := &DefaultInstallConfigure{
		Addr:               addr,
		SystemConfig:       systemConfig,
		SystemConfigServer: systemConfigServer,
	}

	k8sBytes := systemConfig.Get("go", "micro", "kubernetes", "k8s").Bytes()
	k8sConfig, err := createKubernetesJsonConfig(k8sBytes)
	if err != nil {
		return nil, err
	}
	defaultInstallConfigure.K8sConfig = k8sConfig

	ClientV1, restConf, err := clientv1.BuildClient(k8sConfig.Name, k8sConfig.Config)
	if err != nil {
		return nil, err
	}
	defaultInstallConfigure.ClientV1 = ClientV1
	defaultInstallConfigure.RestConfig = restConf

	clientV2, err := clientv2.NewCacheInformerFactory(k8sConfig.Name, k8sConfig.Config)
	if err != nil {
		return nil, err
	}
	defaultInstallConfigure.ClientV2 = clientV2

	return defaultInstallConfigure, nil
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
