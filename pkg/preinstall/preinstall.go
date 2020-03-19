package preinstall

import (
	"encoding/json"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/micro/cors"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/micro/micro/plugin"
	"github.com/opentracing/opentracing-go"
	"github.com/yametech/fuxi/pkg/db"
	"github.com/yametech/fuxi/pkg/k8s/client"
	dynclient "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/mysql"
	"github.com/yametech/fuxi/thirdparty/lib/token"
	"github.com/yametech/fuxi/thirdparty/lib/tracer"
	"github.com/yametech/fuxi/thirdparty/lib/wrapper/tracer/opentracing/gin2micro"
	"github.com/yametech/fuxi/util/common"
	k8sjson "k8s.io/apimachinery/pkg/util/json"
	clientcmdapiV1 "k8s.io/client-go/tools/clientcmd/api/v1"
	"time"
)

func InitDB(opt *mysql.Option) (err error) {
	if db.DB != nil {
		return nil
	}
	db.DB, err = mysql.WithOptionsCreateClient(opt)
	if err != nil {
		return err
	}

	return nil
}

// CheckAndSync table check if not exits table create it or sync table structure
//func CheckAndSync(val interface{}, created bool, forced bool) error {
//	if created {
//		return DB.CreateTableIfNotExist(val)
//	}
//	// TODO: sync columns , forced is true then sync(add/delete) column
//
//	return nil
//}

//InitGateWay init a gateway
func InitGateWay() {
	// token2 := &token.Token{}
	// whitelist2 := &whitelist.Whitelist{}

	if err := plugin.Register(cors.NewPlugin()); err != nil {
		log.Fatal(err)
	}
	// err := plugin.Register(
	// 	plugin.NewPlugin(
	// 		plugin.WithName("auth"),
	// 		plugin.WithHandler(auth.JWTAuthWrapper(token2, whitelist2)),
	// 		plugin.WithFlag(EtcdStringFlag()),
	// 		plugin.WithInit(func(ctx *cli.Context) error {
	// 			conf := common.NewConfigServer(
	// 				ctx.String("etcd_address"),
	// 				common.ConfigPrefix,
	// 			)
	// 			if err := whitelist2.InitConfig(conf, "go", "micro", "urls", "list"); err != nil {
	// 				return err
	// 			}
	// 			return token2.InitConfig(conf, "go", "micro", "jwt", "key")
	// 		}),
	// 	))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = plugin.Register(
	// 	plugin.NewPlugin(
	// 		plugin.WithName("tracer"),
	// 		plugin.WithHandler(stdhttp.TracerWrapper),
	// 	))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = plugin.Register(
	// 	plugin.NewPlugin(
	// 		plugin.WithName("breaker"),
	// 		plugin.WithHandler(hystrix.BreakerWrapper),
	// 	))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = plugin.Register(
	// 	plugin.NewPlugin(
	// 		plugin.WithName("metrics"),
	// 		plugin.WithHandler(prometheus.MetricsWrapper),
	// 	))
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

// EtcdStringFlag a Etcd String Flag
func EtcdStringFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:   "etcd_address",
		Usage:  "etcd address for config K/V",
		EnvVar: "ETCD_ADDRESS",
		Value:  "fuxi.io",
	}
}

//todo:考虑把参数封装成ServiceOption struct
//todo: RegisterTTL RegisterInterval 也基于参数的方式配置进来
//InitApi init a api
func InitApi(sampling int, name, version, tracingAddr string) (web.Service, *token.Token, error) {
	// New Service
	token := &token.Token{}

	gin2micro.SetSamplingFrequency(sampling)
	t, io, err := tracer.NewTracer(name, tracingAddr)
	if err != nil {
		return web.NewService(), token, err
	}
	defer func() {
		if err := io.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	opentracing.SetGlobalTracer(t)

	service := web.NewService(
		web.Name(name),
		web.Version(common.Version(version)),
		web.RegisterTTL(time.Second*15),
		web.RegisterInterval(time.Second*10),
		web.Flags(EtcdStringFlag()),
		web.Action(func(ctx *cli.Context) {
			configureDb(ctx.String("etcd_address"))
			conf := common.NewConfigServer(
				ctx.String("etcd_address"),
				common.ConfigPrefix,
			)
			token.InitConfig(conf, "go", "micro", "jwt", "key")

			// 20200310 add by laik, Need to use in development, temporarily join
			// Initialise kubeclietn on web api support workload service.
			if err := InitKubeClient(ctx.String("etcd_address")); err != nil {
				panic(err)
			}
			if err := InitKubeClientFactoryResourceHandler(); err != nil {
				panic(err)
			}

			if err := InitKubeDynamicClient(ctx.String("etcd_address")); err != nil {
				panic(err)
			}
		}),
	)

	// Initialise service
	if err := service.Init(); err != nil {
		return service, nil, err
	}

	return service, token, nil
}

//InitSrv  init a Service
func InitSrv(name, version string) micro.Service {
	service := micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.Flags(EtcdStringFlag()),
		micro.Action(func(ctx *cli.Context) {
			configureDb(ctx.String("etcd_address"))
			InitKubeClient(ctx.String("etcd_address"))
		}),
	)

	return service

}

//configDb configure db and init db
func configureDb(addr string) {
	configServer := NewConfig(addr)
	mysqlOption := mysql.NewOption()
	err := json.Unmarshal(configServer.Get("go", "micro", "database", "db").Bytes(), mysqlOption)
	if err != nil {
		log.Fatal("unmarshal bytes error: " + err.Error())
	}
	if err = InitDB(mysqlOption); err != nil {
		log.Fatalf("init db error use option %#v: %s", mysqlOption, err)
	}

	log.Info("initialize database")
}

type k8sConfig struct {
	Name   string                `json:"name"`
	Config clientcmdapiV1.Config `json:"config"`
}

//InitKubeClient init k8 clinet
func InitKubeClient(addr string) error {
	var configs k8sConfig
	configServer := NewConfig(addr)
	conf := configServer.Get("go", "micro", "kubernetes", "k8s").Bytes()
	err := k8sjson.Unmarshal(conf, &configs)
	if err != nil {
		log.Fatal("unmarshal bytes error: " + err.Error())
	}
	res, restConf, err := client.BuildClient(configs.Name, configs.Config)
	if err != nil {
		return err
	}
	client.RestConf = restConf
	client.K8sClient = res

	return err
}

func InitKubeDynamicClient(addr string) error {
	var configs k8sConfig
	configServer := NewConfig(addr)
	conf := configServer.Get("go", "micro", "kubernetes", "k8s").Bytes()
	err := k8sjson.Unmarshal(conf, &configs)
	if err != nil {
		log.Fatal("unmarshal bytes error: " + err.Error())
	}
	_, err = dynclient.NewCacheInformerFactory(configs.Name, configs.Config)
	if err != nil {
		return err
	}

	return nil
}

func InitKubeClientFactoryResourceHandler() error {
	if client.K8sClient == nil {
		return fmt.Errorf("%s", "kubernetes client needs to start first")
	}
	cacheFactor, err := client.BuildCacheController(client.K8sClient, client.RestConf)
	if err != nil {
		return err
	}
	client.K8sResourceHandler = client.NewResourceHandler(client.K8sClient, cacheFactor)

	return nil
}

func NewConfig(addr string) config.Config {
	log.Info("start load config from etcd adddress:" + addr)
	configure := common.NewConfigServer(addr, common.ConfigPrefix)
	configServer := config.NewConfig()
	if err := configServer.Load(configure); err != nil {
		log.Fatal("load config from source error: " + err.Error())
	}
	return configServer
}
