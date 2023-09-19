package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +enum

type ApplicationPhase string

const (
	ApplicationPending   ApplicationPhase = "Pending"
	ApplicationAvailable ApplicationPhase = "Available"
	ApplicationFailed    ApplicationPhase = "Failed"

	StateInstalling      = "Installing"
	StateInstallFailed   = "InstallFailed"
	StateUpgrading       = "Upgrading"
	StateUpgradeFailed   = "UpgradeFailed"
	StateInstalled       = "Installed"
	StateUninstalling    = "Uninstalling"
	StateUninstalled     = "Uninstalled"
	StateUninstallFailed = "UninstallFailed"
)

type ApplicationSpec struct {
	ApplicationClassName string `json:"className"`
	Parameters           string `json:"parameters"`
}

type ApplicationStatus struct {
	ResourcePhase  map[string]string  `json:"parameters,omitempty" protobuf:"bytes,3,rep,name=phase"`
	Conditions     []metav1.Condition `json:"conditions,omitempty"`
	ResourceStatus `json:",inline"`
	Message        string `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
}

type ResourceStatus struct {
	State        string `json:"state,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
}

// +kubebuilder:object:root=true

type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}
