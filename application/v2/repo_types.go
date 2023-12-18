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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kubesphere.io/api/constants"
)

type RepoCredential struct {
	// chart repository username
	Username string `json:"username,omitempty"`
	// chart repository password
	Password string `json:"password,omitempty"`
	// identify HTTPS client using this SSL certificate file
	CertFile string `json:"certFile,omitempty"`
	// identify HTTPS client using this SSL key file
	KeyFile string `json:"keyFile,omitempty"`
	// verify certificates of HTTPS-enabled servers using this CA bundle
	CAFile string `json:"caFile,omitempty"`
	// skip tls certificate checks for the repository, default is ture
	InsecureSkipTLSVerify *bool `json:"insecureSkipTLSVerify,omitempty"`
}

// RepoSpec defines the desired state of Repo
type RepoSpec struct {
	//  repo url
	Url string `json:"url"`
	//  repo credential
	Credential RepoCredential `json:"credential,omitempty"`
	// chart repo description from frontend
	Description string `json:"description,omitempty"`
	// sync period in seconds, no sync when SyncPeriod=0, the minimum SyncPeriod is 180s
	SyncPeriod int `json:"syncPeriod,omitempty"`
}

// RepoStatus defines the observed state of Repo
type RepoStatus struct {
	LastUpdateTime *metav1.Time `json:"lastUpdateTime,omitempty"`
	State          string       `json:"state,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,path=repos,shortName=repo
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Workspace",type="string",JSONPath=".metadata.labels.kubesphere\\.io/workspace"
// +kubebuilder:printcolumn:name="url",type=string,JSONPath=`.spec.url`
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// Repo is the Schema for the repoes API
type Repo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RepoSpec   `json:"spec,omitempty"`
	Status RepoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepoList contains a list of Repo
type RepoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repo `json:"items"`
}

func (in *Repo) GetWorkspace() string {
	return getValue(in.Labels, constants.WorkspaceLabelKey)
}

func (in *Repo) GetCreator() string {
	return getValue(in.Annotations, constants.CreatorAnnotationKey)
}
