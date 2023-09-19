package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Maintainers struct {
	// Name is a username or organization name
	Name string `json:"name,omitempty"`
	// URL is an optional URL to an address for the named maintainer
	URL string `json:"url,omitempty"`
	// Email is an optional email address to contact the named maintainer
	Email string `json:"email,omitempty"`
}

type ApplicationClassSpec struct {
	AppVersion     string        `json:"appVersion,omitempty"`
	PackageVersion string        `json:"packageVersion,omitempty"`
	Description    string        `json:"description,omitempty"`
	Icon           string        `json:"icon,omitempty"`
	Maintainers    []Maintainers `json:"maintainers,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope="Cluster"

type ApplicationClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Provisioner string `json:"provisioner" protobuf:"bytes,2,opt,name=provisioner"`

	Parameters map[string]string `json:"parameters,omitempty" protobuf:"bytes,3,rep,name=parameters"`

	Spec ApplicationClassSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

type ApplicationClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationClass `json:"items"`
}
