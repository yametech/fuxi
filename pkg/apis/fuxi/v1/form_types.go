package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type DataNode struct {
	Title    string     `json:"title,omitempty"`
	Key      string     `json:"key,omitempty"`
	NodeType string     `json:"node_type,omitempty"`
	Children []DataNode `json:"children, omitempty"`
}

// FormSpec defines the desired state of Form
type FormSpec struct {
	Tree        []DataNode `json:"tree, omitempty"`
	PropsSchema string     `json:"props_schema, omitempty"`
}

// FormStatus defines the observed state of Form
type FormStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Form is the Schema for the forms API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=forms,scope=Namespaced
type Form struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FormSpec   `json:"spec,omitempty"`
	Status FormStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FormList contains a list of Form
type FormList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Form `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Form{}, &FormList{})
}
