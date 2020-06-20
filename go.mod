module github.com/yametech/fuxi

go 1.14

require (
	code.gitea.io/sdk/gitea v0.12.0
	contrib.go.opencensus.io/exporter/prometheus v0.1.0 // indirect
	github.com/alauda/kube-ovn v0.10.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-openapi/spec v0.19.6
	github.com/go-openapi/swag v0.19.7 // indirect
	github.com/go-resty/resty/v2 v2.1.0
	github.com/gorilla/websocket v1.4.1
	github.com/igm/sockjs-go v2.0.1+incompatible
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.16.0
	github.com/micro/micro v1.16.0
	github.com/nats-io/nats.go v1.9.1 // indirect
	github.com/operator-framework/operator-sdk v0.16.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/pflag v1.0.5
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	github.com/tektoncd/pipeline v0.11.1
	github.com/yametech/nuwa v0.0.0
	go.opencensus.io v0.22.2 // indirect
	go.uber.org/zap v1.13.0
	golang.org/x/tools v0.0.0-20200305224536-de023d59a5d1 // indirect
	google.golang.org/genproto v0.0.0-20191108220845-16a3f7862a1a // indirect
	k8s.io/api v0.18.4
	k8s.io/apiextensions-apiserver v0.18.2
	k8s.io/apimachinery v0.18.4
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6
	k8s.io/metrics v0.0.0
	k8s.io/sample-controller v0.0.0-20190326030654-b8f621986e45
	knative.dev/pkg v0.0.0-20200207155214-fef852970f43
	sigs.k8s.io/controller-runtime v0.6.0
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	github.com/docker/go-connections => github.com/docker/go-connections v0.3.0
	github.com/docker/go-units => github.com/docker/go-units v0.3.3
	github.com/docker/libnetwork => github.com/docker/libnetwork v0.0.0-20180830151422-a9cd636e3789
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.1
	github.com/micro/go-micro => github.com/micro/go-micro v1.16.0
	github.com/yametech/nuwa => github.com/yametech/nuwa v1.0.1-0.20200602142225-f13837e897ae
	k8s.io/api => k8s.io/api v0.0.0-20191114100237-2cd11237263f // 1.15.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190918201827-3de75813f604
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191004115701-31ade1b30762 // 1.15.6
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20191016112112-5190913f932d
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20191016114015-74ad18325ed5
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191114101336-8cba805ad12d // 1.15.6
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20191016115326-20453efc2458
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20191016115129-c07a134afb42
	k8s.io/code-generator => k8s.io/code-generator v0.17.0
	k8s.io/component-base => k8s.io/component-base v0.0.0-20191016111319-039242c015a9
	k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190828162817-608eb1dad4ac
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.18.3
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20191016112429-9587704a8ad4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20191016114939-2b2b218dc1df
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20191016114407-2e83b6f20229
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20191016114748-65049c67a58b
	k8s.io/kubectl => k8s.io/kubectl v0.0.0-20191016120415-2ed914427d51
	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20191016114556-7841ed97f1b2
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20191016115753-cf0698c3a16b
	k8s.io/metrics => k8s.io/metrics v0.18.4
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20191016112829-06bb3c9d77c9
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.3.0
)
