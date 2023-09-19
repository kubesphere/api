package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +enum

type ApplicationResourcePhase string

const (
	ResourcePending   ApplicationResourcePhase = "Pending"
	ResourceAvailable ApplicationResourcePhase = "Available"
	ResourceFailed    ApplicationResourcePhase = "Failed"
)

type ApplicationResourceSpec struct {
	ApplicationRef *v1.ObjectReference `json:"claimRef,omitempty" protobuf:"bytes,4,opt,name=claimRef"`

	ApplicationClassName string `json:"className,omitempty"`
}

type ApplicationResourceStatus struct {
	Phase ApplicationResourcePhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=PersistentVolumePhase"`

	Message string `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`

	Reason string `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope="Cluster"

type ApplicationResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationResourceSpec   `json:"spec,omitempty"`
	Status ApplicationResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type ApplicationResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationResource `json:"items"`
}
