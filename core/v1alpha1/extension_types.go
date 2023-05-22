/*
Copyright 2022 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InstallationMode string

const (
	InstallationModeHostOnly = "HostOnly"
	InstallationMulticluster = "Multicluster"
)

// Vendor describes an extension vendor.
type Vendor struct {
	// Name is a username or organization name
	Name string `json:"name,omitempty"`
	// URL is an optional URL to an address for the named vendor
	URL string `json:"url,omitempty"`
	// Email is an optional email address to contact the named vendor
	Email string `json:"email,omitempty"`
}

// ExtensionInfo describes an extension's basic information.
type ExtensionInfo struct {
	DisplayName Locales `json:"displayName,omitempty"`
	Description Locales `json:"description,omitempty"`
	Icon        string  `json:"icon,omitempty"`
	Vendor      Vendor  `json:"vendor,omitempty"`
}

// ExtensionSpec only contains basic extension information copied from the latest ExtensionVersion.
type ExtensionSpec struct {
	*ExtensionInfo `json:",inline"`
}

// ExtensionVersionSpec contains the details of a specific version extension.
type ExtensionVersionSpec struct {
	*ExtensionInfo `json:",inline"`
	Version        string   `json:"version,omitempty"`
	Keywords       []string `json:"keywords,omitempty"`
	Sources        []string `json:"sources,omitempty"`
	Repository     string   `json:"repository,omitempty"`
	// KubeVersion is a SemVer constraint specifying the version of Kubernetes required.
	// eg: >= 1.2.0, see https://github.com/Masterminds/semver for more info.
	KubeVersion string `json:"kubeVersion,omitempty"`
	// KSVersion is a SemVer constraint specifying the version of KubeSphere required.
	// eg: >= 1.2.0, see https://github.com/Masterminds/semver for more info.
	KSVersion string `json:"ksVersion,omitempty"`
	Home      string `json:"home,omitempty"`
	// ChartDataRef refers to a configMap which contains raw chart data.
	ChartDataRef *ConfigMapKeyRef `json:"chartDataRef,omitempty"`
	ChartURL     string           `json:"chartURL,omitempty"`
	// +kubebuilder:default:=HostOnly
	InstallationMode     InstallationMode     `json:"installationMode,omitempty"`
	ExternalDependencies []ExternalDependency `json:"externalDependencies,omitempty"`
}

type ExternalDependency struct {
	Name string `json:"name"`
	// Type of dependency, default to extension
	// +optional
	Type     string `json:"type,omitempty"`
	Version  string `json:"version"`
	Required bool   `json:"required"`
}

type ConfigMapKeyRef struct {
	corev1.ConfigMapKeySelector `json:",inline"`
	Namespace                   string `json:"namespace"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories="extensions",scope="Cluster"

type ExtensionVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExtensionVersionSpec `json:"spec,omitempty"`
}

type CategorySpec struct {
	DisplayName Locales `json:"displayName,omitempty"`
	Description Locales `json:"description,omitempty"`
	Icon        string  `json:"icon,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories="extensions",scope="Cluster"

// Category can help us group the extensions.
type Category struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CategorySpec `json:"spec,omitempty"`
}

type ExtensionVersionInfo struct {
	Version           string      `json:"version"`
	CreationTimestamp metav1.Time `json:"creationTimestamp,omitempty"`
}

type ExtensionStatus struct {
	State              string                 `json:"state,omitempty"`
	SubscribedVersion  string                 `json:"subscribedVersion,omitempty"`
	RecommendedVersion string                 `json:"recommendedVersion,omitempty"`
	Versions           []ExtensionVersionInfo `json:"versions,omitempty"`
	Conditions         []metav1.Condition     `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories="extensions",scope="Cluster"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"

// Extension is synchronized from the Repository.
// An extension can contain multiple versions.
type Extension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExtensionSpec   `json:"spec,omitempty"`
	Status            ExtensionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type ExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Extension `json:"items"`
}

// +kubebuilder:object:root=true

type ExtensionVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExtensionVersion `json:"items"`
}

// +kubebuilder:object:root=true

type CategoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Category `json:"items"`
}
