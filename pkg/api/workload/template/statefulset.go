package template

import nuwav1 "github.com/yametech/nuwa/api/v1"

// StatefulSet1Request nuwa statefulset request
type StatefulSet1Request struct {
	Model CommonTemplate         `json:"model" from:"model"`
	Spec  nuwav1.StatefulSetSpec `json:"spec" form:"spec"`
}

// StatefulSet1Response nuwa statefulset response
type StatefulSet1Response struct {
	StatefulSetList nuwav1.StatefulSetList `json:"statefulset_list"`
}
