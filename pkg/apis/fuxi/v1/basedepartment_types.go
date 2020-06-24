package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Stack struct {
	Address string `json:"address,omitempty"`
	// +optional
	Token string `json:"token,omitempty"`
	// +optional
	User string `json:"user,omitempty"`
	// +optional
	Password string `json:"password,omitempty"`
}

// BaseDepartmentSpec defines the desired state of BaseDepartment
type BaseDepartmentSpec struct {
	Namespace []string `json:"namespace,omitempty"`
	// +optional
	DefaultNamespace string `json:"defaultNamespace,omitempty"`
	// +optional
	Gits []Stack `json:"gits,omitempty"`
	// +optional
	Registers []Stack `json:"registers,omitempty"`
}

// BaseDepartmentStatus defines the observed state of BaseDepartment
type BaseDepartmentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BaseDepartment is the Schema for the basedepartments API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=basedepartments,scope=Namespaced
type BaseDepartment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BaseDepartmentSpec   `json:"spec,omitempty"`
	Status BaseDepartmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BaseDepartmentList contains a list of BaseDepartment
type BaseDepartmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BaseDepartment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BaseDepartment{}, &BaseDepartmentList{})
}
