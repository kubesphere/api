// Package v1alpha1 contains API Schema definitions for the tower v1alpha1 API group
// +k8s:openapi-gen=true
// +kubebuilder:object:generate=true
// +k8s:defaulter-gen=TypeMeta
// +groupName=appstore.kubesphere.io

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

const GroupName = "appstore.kubesphere.io"

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	AddToScheme = SchemeBuilder.AddToScheme
)

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func init() {
	SchemeBuilder.Register(
		&Application{},
		&ApplicationList{},
		&ApplicationClass{},
		&ApplicationClassList{},
		&ApplicationResource{},
		&ApplicationResourceList{},
	)
}
