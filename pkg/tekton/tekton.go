package tekton

import (
	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
)

var TektonClient tektonclient.Interface
