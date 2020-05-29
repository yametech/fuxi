package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WorkloadsSpec defines the desired state of Workloads
type WorkloadsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	// Workloads ResourceType [Deployment,Statefulset...]
	AppName           *string `json:"appName"`
	ResourceType      *string `json:"resourceType"`
	GenerateTimestamp *int64  `json:"generateTimestamp"`
	Metadata          *string `json:"metadata"`
}

// WorkloadsStatus defines the observed state of Workloads
type WorkloadsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Workloads is the Schema for the workloads API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=workloads,scope=Namespaced
type Workloads struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkloadsSpec   `json:"spec,omitempty"`
	Status WorkloadsStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkloadsList contains a list of Workloads
type WorkloadsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workloads `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Workloads{}, &WorkloadsList{})
}
