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

type ApplicationDetectionRuleV2InitParameters struct {

	// (String) Select an existing application or create a new one.
	// Select an existing application or create a new one.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (String) (v1.274) Add a description for your rule
	// (v1.274) Add a description for your rule
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) Possible Values: DOMAIN_CONTAINS, DOMAIN_ENDS_WITH, DOMAIN_EQUALS, DOMAIN_MATCHES, DOMAIN_STARTS_WITH, URL_CONTAINS, URL_ENDS_WITH, URL_EQUALS, URL_STARTS_WITH
	// Possible Values: `DOMAIN_CONTAINS`, `DOMAIN_ENDS_WITH`, `DOMAIN_EQUALS`, `DOMAIN_MATCHES`, `DOMAIN_STARTS_WITH`, `URL_CONTAINS`, `URL_ENDS_WITH`, `URL_EQUALS`, `URL_STARTS_WITH`
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// (String) Pattern
	// Pattern
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type ApplicationDetectionRuleV2Observation struct {

	// (String) Select an existing application or create a new one.
	// Select an existing application or create a new one.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (String) (v1.274) Add a description for your rule
	// (v1.274) Add a description for your rule
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) Possible Values: DOMAIN_CONTAINS, DOMAIN_ENDS_WITH, DOMAIN_EQUALS, DOMAIN_MATCHES, DOMAIN_STARTS_WITH, URL_CONTAINS, URL_ENDS_WITH, URL_EQUALS, URL_STARTS_WITH
	// Possible Values: `DOMAIN_CONTAINS`, `DOMAIN_ENDS_WITH`, `DOMAIN_EQUALS`, `DOMAIN_MATCHES`, `DOMAIN_STARTS_WITH`, `URL_CONTAINS`, `URL_ENDS_WITH`, `URL_EQUALS`, `URL_STARTS_WITH`
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// (String) Pattern
	// Pattern
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type ApplicationDetectionRuleV2Parameters struct {

	// (String) Select an existing application or create a new one.
	// Select an existing application or create a new one.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (String) (v1.274) Add a description for your rule
	// (v1.274) Add a description for your rule
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// +kubebuilder:validation:Optional
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) Possible Values: DOMAIN_CONTAINS, DOMAIN_ENDS_WITH, DOMAIN_EQUALS, DOMAIN_MATCHES, DOMAIN_STARTS_WITH, URL_CONTAINS, URL_ENDS_WITH, URL_EQUALS, URL_STARTS_WITH
	// Possible Values: `DOMAIN_CONTAINS`, `DOMAIN_ENDS_WITH`, `DOMAIN_EQUALS`, `DOMAIN_MATCHES`, `DOMAIN_STARTS_WITH`, `URL_CONTAINS`, `URL_ENDS_WITH`, `URL_EQUALS`, `URL_STARTS_WITH`
	// +kubebuilder:validation:Optional
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// (String) Pattern
	// Pattern
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

// ApplicationDetectionRuleV2Spec defines the desired state of ApplicationDetectionRuleV2
type ApplicationDetectionRuleV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationDetectionRuleV2Parameters `json:"forProvider"`
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
	InitProvider ApplicationDetectionRuleV2InitParameters `json:"initProvider,omitempty"`
}

// ApplicationDetectionRuleV2Status defines the observed state of ApplicationDetectionRuleV2.
type ApplicationDetectionRuleV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationDetectionRuleV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ApplicationDetectionRuleV2 is the Schema for the ApplicationDetectionRuleV2s API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type ApplicationDetectionRuleV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.applicationId) || (has(self.initProvider) && has(self.initProvider.applicationId))",message="spec.forProvider.applicationId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.matcher) || (has(self.initProvider) && has(self.initProvider.matcher))",message="spec.forProvider.matcher is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.pattern) || (has(self.initProvider) && has(self.initProvider.pattern))",message="spec.forProvider.pattern is a required parameter"
	Spec   ApplicationDetectionRuleV2Spec   `json:"spec"`
	Status ApplicationDetectionRuleV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationDetectionRuleV2List contains a list of ApplicationDetectionRuleV2s
type ApplicationDetectionRuleV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationDetectionRuleV2 `json:"items"`
}

// Repository type metadata.
var (
	ApplicationDetectionRuleV2_Kind             = "ApplicationDetectionRuleV2"
	ApplicationDetectionRuleV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationDetectionRuleV2_Kind}.String()
	ApplicationDetectionRuleV2_KindAPIVersion   = ApplicationDetectionRuleV2_Kind + "." + CRDGroupVersion.String()
	ApplicationDetectionRuleV2_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationDetectionRuleV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationDetectionRuleV2{}, &ApplicationDetectionRuleV2List{})
}
