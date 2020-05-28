package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BasePermissionSpec defines the desired state of BasePermission
type BasePermissionSpec struct {
	// [超级管理员],[运维管理,业务运维],[开发主管,业务管理,业务开发]
	Name string `json:"name"`
	// [777],[770,077],[077,037,007]
	Value    uint32 `json:"value"`
	IsDelete bool   `json:"is_delete"`
	Comment  string `json:"comment"`
}

// BasePermissionStatus defines the observed state of BasePermission
type BasePermissionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BasePermission is the Schema for the basepermissions API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=basepermissions,scope=Namespaced
type BasePermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BasePermissionSpec   `json:"spec,omitempty"`
	Status BasePermissionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BasePermissionList contains a list of BasePermission
type BasePermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BasePermission `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BasePermission{}, &BasePermissionList{})
}
