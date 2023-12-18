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

const (
	RepoIDLabelKey              = "application.kubesphere.io/repo-name"
	ReqUserAnnotationKey        = "application.kubesphere.io/req-user"
	AppIDLabelKey               = "application.kubesphere.io/app-id"
	AppTypeLabelKey             = "application.kubesphere.io/app-type"
	AppStoreLabelKey            = "application.kubesphere.io/app-store"
	AppCategoryNameKey          = "application.kubesphere.io/app-category-name"
	LatestAppVersionKey         = "application.kubesphere.io/latest-app-version"
	AppMaintainersKey           = "application.kubesphere.io/app-maintainers"
	AppReleaseReferenceLabelKey = "application.kubesphere.io/app-release-name"
	UncategorizedCategoryID     = "kubesphere-app-uncategorized"
	StatusActive                = "active"
	StatusSuccessful            = "successful"
	StatusCreating              = "creating"
	StatusDeleting              = "deleting"
	StatusUpgrading             = "upgrading"
	StatusFailed                = "failed"
	StatusCreated               = "created"
	StatusUpgraded              = "upgraded"
	StatusNosync                = "nosync"
	AppTypeHelm                 = "helm"
	AppTypeYaml                 = "yaml"
	AppTypeEdge                 = "edge"
	BinaryKey                   = "BinaryKey"
	MaxNumOfVersions            = 10
	SystemWorkspace             = "system-workspace"
	// App review status: draft, submitted, passed, rejected, suspended, active
	ReviewStatusDraft     = "draft"
	ReviewStatusSubmitted = "submitted"
	ReviewStatusPassed    = "passed"
	ReviewStatusRejected  = "rejected"
	ReviewStatusSuspended = "suspended"
	ReviewStatusActive    = "active"
)
