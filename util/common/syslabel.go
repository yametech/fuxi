package common

// fuxi label
const (
	// NamespaceLabelForDepartment  eg: fuxi.kubernetes.io/department: dxp
	NamespaceLabelForDepartment = "fuxi.kubernetes.io/department"
)

// nuwa annotations
const (
	// NamespaceAnnotationForNodeResource
	// eg: nuwa.kubernetes.io/default_resource_limit: '[{"zone":"A","rack":"W-01","host":"node1"},{"zone":"A","rack":"W-02","host":"node2"}]'
	// the identified by the operation and maintenance, and the release resources are limited to these areas
	NamespaceAnnotationForNodeResource = "nuwa.kubernetes.io/default_resource_limit"
)
