// +build tools

// Place any runtime dependencies as imports in this file.
// Go modules will be forced to download and install them.
package hack

import (
	//_ "k8s.io/code-generator/cmd/clientv2-gen"
	//_ "k8s.io/code-generator/cmd/deepcopy-gen"
	////_ "k8s.io/code-generator/cmd/defaulter-gen"
	//_ "k8s.io/code-generator/cmd/informer-gen"
	//_ "k8s.io/code-generator/cmd/lister-gen"
	//_ "k8s.io/kube-openapi/cmd/openapi-gen"
	//_ "k8s.io/code-generator"
	_ "github.com/beorn7/perks/quantile"
)
