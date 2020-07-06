package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TektonGraphSpec defines the desired state of TektonGraph
type TektonGraphSpec struct {
	// +optional
	Data string `json:"data, omitempty"`
}

// TektonGraphStatus defines the observed state of TektonGraph
type TektonGraphStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TektonGraph is the Schema for the tektongraphs API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=tektongraphs,scope=Namespaced
type TektonGraph struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TektonGraphSpec   `json:"spec,omitempty"`
	Status TektonGraphStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TektonGraphList contains a list of TektonGraph
type TektonGraphList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TektonGraph `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TektonGraph{}, &TektonGraphList{})
}
