/*
Copyright 2020 The KubeSphere Authors.

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

package v2

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kubesphere.io/api/constants"
)

// ApplicationReleaseSpec defines the desired state of ApplicationRelease
type ApplicationReleaseSpec struct {
	AppID        string `json:"appID"`
	AppVersionID string `json:"appVersionID"`
	Values       []byte `json:"values,omitempty"`
	AppType      string `json:"appType,omitempty"`
}

// ApplicationReleaseStatus defines the observed state of ApplicationRelease
type ApplicationReleaseStatus struct {
	State   string `json:"state"`
	Message string `json:"message,omitempty"`
	// This is used to compare whether the spec has been modified to determine if an upgrade is needed.
	SpecHash          string            `json:"specHash,omitempty"`
	JobName           string            `json:"jobName,omitempty"`
	LastUpdate        metav1.Time       `json:"lastUpdate,omitempty"`
	LastDeployed      *metav1.Time      `json:"lastDeployed,omitempty"`
	RealTimeResources []json.RawMessage `json:"realTimeResources,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName=apprls
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="App Name",type=string,JSONPath=".spec.AppID"
// +kubebuilder:printcolumn:name="Workspace",type="string",JSONPath=".metadata.labels.kubesphere\\.io/workspace"
// +kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=".metadata.labels.kubesphere\\.io/cluster"
// +kubebuilder:printcolumn:name="Namespace",type="string",JSONPath=".metadata.labels.kubesphere\\.io/namespace"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ApplicationRelease is the Schema for the applicationreleases API
type ApplicationRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationReleaseSpec   `json:"spec,omitempty"`
	Status ApplicationReleaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationReleaseList contains a list of ApplicationRelease
type ApplicationReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationRelease `json:"items"`
}

func (in *ApplicationRelease) GetCreator() string {
	return getValue(in.Annotations, constants.CreatorAnnotationKey)
}

func (in *ApplicationRelease) GetRlsCluster() string {
	name := getValue(in.Labels, constants.ClusterNameLabelKey)
	if name != "" {
		return name
	}
	//todo remove hardcode
	return "host"
}

func (in *ApplicationRelease) GetRlsNamespace() string {
	ns := getValue(in.Labels, constants.NamespaceLabelKey)
	if ns == "" {
		return "default"
	}
	return ns
}

func (in *ApplicationRelease) HashSpec() string {
	specJSON, _ := json.Marshal(in.Spec)
	return fmt.Sprintf("%x", md5.Sum(specJSON))
}
