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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	Automatic UpgradeStrategy = "Automatic"
	Manual    UpgradeStrategy = "Manual"
)

type Placement struct {
	// +listType=set
	// +optional
	Clusters        []string              `json:"clusters,omitempty"`
	ClusterSelector *metav1.LabelSelector `json:"clusterSelector,omitempty"`
}

type ClusterScheduling struct {
	Placement *Placement        `json:"placement,omitempty"`
	Overrides map[string]string `json:"overrides,omitempty"`
}

type SubscriptionState struct {
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	State              string      `json:"state"`
}

type InstallationStatus struct {
	State           string              `json:"state,omitempty"`
	ReleaseName     string              `json:"releaseName,omitempty"`
	TargetNamespace string              `json:"targetNamespace,omitempty"`
	JobName         string              `json:"jobName,omitempty"`
	Conditions      []metav1.Condition  `json:"conditions,omitempty"`
	StateHistory    []SubscriptionState `json:"stateHistory,omitempty"`
}

type ExtensionRef struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type UpgradeStrategy string

type SubscriptionSpec struct {
	Extension ExtensionRef `json:"extension"`
	Enabled   bool         `json:"enabled"`
	// +kubebuilder:default:=Manual
	UpgradeStrategy   UpgradeStrategy    `json:"upgradeStrategy,omitempty"`
	Config            string             `json:"config,omitempty"`
	ClusterScheduling *ClusterScheduling `json:"clusterScheduling,omitempty"`
}

type SubscriptionStatus struct {
	InstallationStatus `json:",inline"`
	// ClusterSchedulingStatuses describes the subchart installation status of the extension
	ClusterSchedulingStatuses map[string]InstallationStatus `json:"clusterSchedulingStatuses,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories="extensions",scope="Cluster"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"

// Subscription describes the configuration and the extension version to be subscribed.
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionSpec   `json:"spec,omitempty"`
	Status            SubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}
