package common

import nuwav1 "github.com/yametech/nuwa/api/v1"

// fuxi label
const (
	// BaseServiceStoreageNamespace User/Role/Department data store namespace
	BaseServiceStoreageNamespace = "kube-system"
	// NamespaceLabelForDepartment  eg: fuxi.kubernetes.io/department: dxp
	NamespaceLabelForDepartment = "fuxi.kubernetes.io/department"

	// NamespaceLabelForDepartment  eg: fuxi.kubernetes.io/namespce: dxp
	// the dxp is user allow access namespaces
	DeployResourceLabelForNamespace = "fuxi.kubernetes.io/namespce"
)

// nuwa annotations
const (
	// NamespaceAnnotationForNodeResource
	// eg: nuwa.kubernetes.io/default_resource_limit: '[{"zone":"A","rack":"W-01","host":"node1"},{"zone":"A","rack":"W-02","host":"node2"}]'
	// the identified by the operation and maintenance, and the release resources are limited to these areas
	NamespaceAnnotationForNodeResource = nuwav1.NuwaLimitFlag

	// NamespaceAnnotationForNodeResource
	// eg: nuwa.kubernetes.io/default_storage_limit: '["a","b","c"]'
	NamespaceAnnotationForStorageClass = "fuxi.kubernetes.io/default_storage_limit"
)
