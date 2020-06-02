package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BaseRoleUserSpec defines the desired state of BaseRoleUser
type BaseRoleUserSpec struct {
	RoleId string `json:"role_id,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

// BaseRoleUserStatus defines the observed state of BaseRoleUser
type BaseRoleUserStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BaseRoleUser is the Schema for the baseroleusers API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=baseroleusers,scope=Namespaced
type BaseRoleUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BaseRoleUserSpec   `json:"spec,omitempty"`
	Status BaseRoleUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BaseRoleUserList contains a list of BaseRoleUser
type BaseRoleUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BaseRoleUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BaseRoleUser{}, &BaseRoleUserList{})
}
