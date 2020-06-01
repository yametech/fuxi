package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BaseRoleSpec defines the desired state of BaseRole
type BaseRoleSpec struct {
	RoleName string `json:"role_name, omitempty"`
}

// BaseRoleStatus defines the observed state of BaseRole
type BaseRoleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BaseRole is the Schema for the baseroles API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=baseroles,scope=Namespaced
type BaseRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BaseRoleSpec   `json:"spec,omitempty"`
	Status BaseRoleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BaseRoleList contains a list of BaseRole
type BaseRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BaseRole `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BaseRole{}, &BaseRoleList{})
}
