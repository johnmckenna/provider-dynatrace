// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IssueTrackingInitParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// critical
	// Set a label to identify these issues, for example, `release_blocker` or `non-critical`
	Issuelabel *string `json:"issuelabel,omitempty" tf:"issuelabel,omitempty"`

	// (String) You can use the following placeholders to automatically insert values from the Release monitoring page in your query: {NAME}, {VERSION}, {STAGE}, {PRODUCT}.
	// You can use the following placeholders to automatically insert values from the **Release monitoring** page in your query: `{NAME}`, `{VERSION}`, `{STAGE}`, `{PRODUCT}`.
	Issuequery *string `json:"issuequery,omitempty" tf:"issuequery,omitempty"`

	// (String) Possible Values: ERROR, INFO, RESOLVED
	// Possible Values: `ERROR`, `INFO`, `RESOLVED`
	Issuetheme *string `json:"issuetheme,omitempty" tf:"issuetheme,omitempty"`

	// (String) Possible Values: GITHUB, GITLAB, JIRA, JIRA_CLOUD, JIRA_ON_PREMISE, SERVICENOW
	// Possible Values: `GITHUB`, `GITLAB`, `JIRA`, `JIRA_CLOUD`, `JIRA_ON_PREMISE`, `SERVICENOW`
	Issuetrackersystem *string `json:"issuetrackersystem,omitempty" tf:"issuetrackersystem,omitempty"`

	// (String) Password
	// Password
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// (String, Sensitive) Token
	// Token
	TokenSecretRef *v1.SecretKeySelector `json:"tokenSecretRef,omitempty" tf:"-"`

	// now.com/)
	// For Jira, use the base URL (for example, https://jira.yourcompany.com); for GitHub, use the repository URL (for example, https://github.com/org/repo); for GitLab, use the specific project API for a single project (for example, https://gitlab.com/api/v4/projects/:projectId), and the specific group API for a multiple projects (for example, https://gitlab.com/api/v4/groups/:groupId); for ServiceNow, use your company instance URL (for example, https://yourinstance.service-now.com/)
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// (String) Username
	// Username
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type IssueTrackingObservation struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// critical
	// Set a label to identify these issues, for example, `release_blocker` or `non-critical`
	Issuelabel *string `json:"issuelabel,omitempty" tf:"issuelabel,omitempty"`

	// (String) You can use the following placeholders to automatically insert values from the Release monitoring page in your query: {NAME}, {VERSION}, {STAGE}, {PRODUCT}.
	// You can use the following placeholders to automatically insert values from the **Release monitoring** page in your query: `{NAME}`, `{VERSION}`, `{STAGE}`, `{PRODUCT}`.
	Issuequery *string `json:"issuequery,omitempty" tf:"issuequery,omitempty"`

	// (String) Possible Values: ERROR, INFO, RESOLVED
	// Possible Values: `ERROR`, `INFO`, `RESOLVED`
	Issuetheme *string `json:"issuetheme,omitempty" tf:"issuetheme,omitempty"`

	// (String) Possible Values: GITHUB, GITLAB, JIRA, JIRA_CLOUD, JIRA_ON_PREMISE, SERVICENOW
	// Possible Values: `GITHUB`, `GITLAB`, `JIRA`, `JIRA_CLOUD`, `JIRA_ON_PREMISE`, `SERVICENOW`
	Issuetrackersystem *string `json:"issuetrackersystem,omitempty" tf:"issuetrackersystem,omitempty"`

	// (String) Password
	// Password
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// now.com/)
	// For Jira, use the base URL (for example, https://jira.yourcompany.com); for GitHub, use the repository URL (for example, https://github.com/org/repo); for GitLab, use the specific project API for a single project (for example, https://gitlab.com/api/v4/projects/:projectId), and the specific group API for a multiple projects (for example, https://gitlab.com/api/v4/groups/:groupId); for ServiceNow, use your company instance URL (for example, https://yourinstance.service-now.com/)
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// (String) Username
	// Username
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type IssueTrackingParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// +kubebuilder:validation:Optional
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// critical
	// Set a label to identify these issues, for example, `release_blocker` or `non-critical`
	// +kubebuilder:validation:Optional
	Issuelabel *string `json:"issuelabel,omitempty" tf:"issuelabel,omitempty"`

	// (String) You can use the following placeholders to automatically insert values from the Release monitoring page in your query: {NAME}, {VERSION}, {STAGE}, {PRODUCT}.
	// You can use the following placeholders to automatically insert values from the **Release monitoring** page in your query: `{NAME}`, `{VERSION}`, `{STAGE}`, `{PRODUCT}`.
	// +kubebuilder:validation:Optional
	Issuequery *string `json:"issuequery,omitempty" tf:"issuequery,omitempty"`

	// (String) Possible Values: ERROR, INFO, RESOLVED
	// Possible Values: `ERROR`, `INFO`, `RESOLVED`
	// +kubebuilder:validation:Optional
	Issuetheme *string `json:"issuetheme,omitempty" tf:"issuetheme,omitempty"`

	// (String) Possible Values: GITHUB, GITLAB, JIRA, JIRA_CLOUD, JIRA_ON_PREMISE, SERVICENOW
	// Possible Values: `GITHUB`, `GITLAB`, `JIRA`, `JIRA_CLOUD`, `JIRA_ON_PREMISE`, `SERVICENOW`
	// +kubebuilder:validation:Optional
	Issuetrackersystem *string `json:"issuetrackersystem,omitempty" tf:"issuetrackersystem,omitempty"`

	// (String) Password
	// Password
	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// (String, Sensitive) Token
	// Token
	// +kubebuilder:validation:Optional
	TokenSecretRef *v1.SecretKeySelector `json:"tokenSecretRef,omitempty" tf:"-"`

	// now.com/)
	// For Jira, use the base URL (for example, https://jira.yourcompany.com); for GitHub, use the repository URL (for example, https://github.com/org/repo); for GitLab, use the specific project API for a single project (for example, https://gitlab.com/api/v4/projects/:projectId), and the specific group API for a multiple projects (for example, https://gitlab.com/api/v4/groups/:groupId); for ServiceNow, use your company instance URL (for example, https://yourinstance.service-now.com/)
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// (String) Username
	// Username
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// IssueTrackingSpec defines the desired state of IssueTracking
type IssueTrackingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IssueTrackingParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider IssueTrackingInitParameters `json:"initProvider,omitempty"`
}

// IssueTrackingStatus defines the observed state of IssueTracking.
type IssueTrackingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IssueTrackingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IssueTracking is the Schema for the IssueTrackings API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type IssueTracking struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.issuelabel) || (has(self.initProvider) && has(self.initProvider.issuelabel))",message="spec.forProvider.issuelabel is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.issuequery) || (has(self.initProvider) && has(self.initProvider.issuequery))",message="spec.forProvider.issuequery is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.issuetheme) || (has(self.initProvider) && has(self.initProvider.issuetheme))",message="spec.forProvider.issuetheme is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.issuetrackersystem) || (has(self.initProvider) && has(self.initProvider.issuetrackersystem))",message="spec.forProvider.issuetrackersystem is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.url) || (has(self.initProvider) && has(self.initProvider.url))",message="spec.forProvider.url is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   IssueTrackingSpec   `json:"spec"`
	Status IssueTrackingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IssueTrackingList contains a list of IssueTrackings
type IssueTrackingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IssueTracking `json:"items"`
}

// Repository type metadata.
var (
	IssueTracking_Kind             = "IssueTracking"
	IssueTracking_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IssueTracking_Kind}.String()
	IssueTracking_KindAPIVersion   = IssueTracking_Kind + "." + CRDGroupVersion.String()
	IssueTracking_GroupVersionKind = CRDGroupVersion.WithKind(IssueTracking_Kind)
)

func init() {
	SchemeBuilder.Register(&IssueTracking{}, &IssueTrackingList{})
}