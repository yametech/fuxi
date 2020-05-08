package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FormRenderSpec defines the desired state of FormRender
type FormRenderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	ID          string `json:"id"`
	PropsSchema string `json:"props_schema,omitempty"`
	UiSchema    string `json:"ui_schema,omitempty"`
}

// FormRenderStatus defines the observed state of FormRender
type FormRenderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FormRender is the Schema for the formrenders API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=formrenders,scope=Namespaced
type FormRender struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FormRenderSpec   `json:"spec,omitempty"`
	Status FormRenderStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FormRenderList contains a list of FormRender
type FormRenderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FormRender `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FormRender{}, &FormRenderList{})
}
