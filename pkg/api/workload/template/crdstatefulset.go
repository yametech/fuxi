package template

import nuwav1 "github.com/yametech/nuwa/api/v1"

type CRDStatefulSetRequest struct {
	Model CommonTemplate         `json:"model" from:"model"`
	Spec  nuwav1.StatefulSetSpec `json:"spec" form:"spec"`
}

type CRDStatefulSetResponse struct {
	StatfulSetList nuwav1.StatefulSetList `json:"statefulset_list"`
}
