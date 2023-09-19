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

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope="Cluster"
// +kubebuilder:printcolumn:name="Provisioner",type="string",JSONPath=".spec.pluginInfo.name"

// ClusterInfo is the Schema for the clusterinfos API. the API is use to store telemetry data.
type ClusterInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterInfoSpec `json:"spec,omitempty"`

	// +optional
	Status *ClusterInfoStatus `json:"status,omitempty"`
}

// ClusterInfoSpec nothing in Spec. only use collect cluster telemetry data
type ClusterInfoSpec struct {
}

// ClusterInfoStatus store cluster telemetry data
type ClusterInfoStatus struct {
	// when to sync data to ksCloud
	SyncTime *metav1.Time `json:"syncTime,omitempty"`

	// extension which cluster has installed. refer to subscriptions.kubesphere.io
	Extension []Extension `json:"extension,omitempty"`

	// cluster info which kubesphere use. refer to clusters.cluster.kubesphere.io
	Clusters []Cluster `json:"clusters,omitempty"`

	// the platform resources total.
	Platform Platform `json:"platform,omitempty"`

	// kubesphere cloud id
	CloudId string `json:"cloudId,omitempty"`

	// collection time
	TotalTime *metav1.Time `json:"ts,omitempty"`
}

type Extension struct {
	// extension name
	Name string `json:"name,omitempty"`
	// extension version
	Version string `json:"version,omitempty"`
	// extension create time
	CreateTime string `json:"cTime,omitempty"`
}

type Cluster struct {
	// cluster role
	Role string `json:"role,omitempty"`
	// cluster name
	Name string `json:"name,omitempty"`
	//cluster uid
	Uid string `json:"uid,omitempty"`
	// cluster namespace id
	Nid string `json:"nid,omitempty"`
	// kubesphere version
	KsVersion string `json:"ksVersion,omitempty"`
	// kubernetes cluster version
	ClusterVersion string `json:"clusterVersion,omitempty"`
	// Namepace number of cluster
	Namespace int `json:"namespace,omitempty"`
	// nodes of cluster
	Nodes []Node `json:"nodes,omitempty"`
}

type Node struct {
	// node uid
	Uid string `json:"uid,omitempty"`
	// node name
	Name string `json:"name,omitempty"`
	// node roles
	Role []string `json:"role,omitempty"`
	// node arch
	Arch string `json:"arch,omitempty"`
	// node containerRuntime
	ContainerRuntime string `json:"containerRuntime,omitempty"`
	// node kernel
	Kernel string `json:"kernel,omitempty"`
	// node kubeProxy
	KubeProxy string `json:"kubeProxy,omitempty"`
	// node kubelet
	Kubelet string `json:"kubelet,omitempty"`
	// node operator system
	Os string `json:"os,omitempty"`
	// os operator system image
	OsImage string `json:"osImage,omitempty"`
}

type Platform struct {
	// workspace number of cluster
	Workspace int `json:"workspace,omitempty"`
	// user number of cluster
	User int `json:"user,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterInfoList contains a list of ClusterInfo
type ClusterInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterInfo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterInfo{}, &ClusterInfoList{})
}
