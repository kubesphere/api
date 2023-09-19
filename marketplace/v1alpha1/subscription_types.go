package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories="marketplace",scope="Cluster"

type Subscription struct {
	metav1.TypeMeta    `json:",inline"`
	metav1.ObjectMeta  `json:"metadata,omitempty"`
	SubscriptionSpec   SubscriptionSpec   `json:"spec"`
	SubscriptionStatus SubscriptionStatus `json:"status,omitempty"`
}

type SubscriptionSpec struct {
	ExtensionName string `json:"extensionName"`
}

type SubscriptionStatus struct {
	ExpiredAt          *metav1.Time `json:"expiredAt"`
	StartedAt          *metav1.Time `json:"startedAt"`
	CreatedAt          *metav1.Time `json:"createdAt"`
	ExtensionID        string       `json:"extensionID"`
	ExtraInfo          string       `json:"extraInfo"`
	OrderID            string       `json:"orderID"`
	SubscriptionID     string       `json:"subscriptionID"`
	UpdatedAt          *metav1.Time `json:"updatedAt"`
	UserID             string       `json:"userID"`
	UserSubscriptionID string       `json:"userSubscriptionID"`
}

// +kubebuilder:object:root=true

type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}
