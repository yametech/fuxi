package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TektonStoreSpec defines the desired state of TektonStore
type TektonStoreSpec struct {
	TektonResourceType string `json:"tektonResourceType, omitempty"`
	Data               string `json:"data, omitempty"`
	Author             string `json:"author, omitempty"`
	Forks              uint32 `json:"forks, omitempty"`
	ParamsDescription  string `json:"paramsDescription, omitempty"`
}

// TektonStoreStatus defines the observed state of TektonStore
type TektonStoreStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TektonStore is the Schema for the tektonstores API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=tektonstores,scope=Namespaced
type TektonStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TektonStoreSpec   `json:"spec,omitempty"`
	Status TektonStoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TektonStoreList contains a list of TektonStore
type TektonStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TektonStore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TektonStore{}, &TektonStoreList{})
}
