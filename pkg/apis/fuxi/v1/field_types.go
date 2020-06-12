package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type SelectStore struct {
	Key   string `json:"key, omitempty" bson:",omitempty"`
	Value string `json:"value, omitempty" bson:",omitempty"`
}

// FieldSpec defines the desired state of Field
type FieldSpec struct {
	// FormRender Type
	FieldType string `json:"field_type"`
	// +optional
	FormDataConfig string `json:"form_data_config, omitempty"`
	// +optional
	PropsSchema string `json:"props_schema, omitempty"`
}

// FieldStatus defines the observed state of Field
type FieldStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Field is the Schema for the fields API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=fields,scope=Namespaced
type Field struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FieldSpec   `json:"spec,omitempty"`
	Status FieldStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FieldList contains a list of Field
type FieldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Field `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Field{}, &FieldList{})
}
