package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BaseRolePermSpec defines the desired state of BaseRolePerm
type BaseRolePermSpec struct {
	PermissionId uint32 `json:"permission_id,omitempty"`
	RoleId       uint32 `json:"role_id,omitempty"`
}

// BaseRolePermStatus defines the observed state of BaseRolePerm
type BaseRolePermStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BaseRolePerm is the Schema for the baseroleperms API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=baseroleperms,scope=Namespaced
type BaseRolePerm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BaseRolePermSpec   `json:"spec,omitempty"`
	Status BaseRolePermStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BaseRolePermList contains a list of BaseRolePerm
type BaseRolePermList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BaseRolePerm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BaseRolePerm{}, &BaseRolePermList{})
}
