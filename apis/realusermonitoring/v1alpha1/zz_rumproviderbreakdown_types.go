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

type DomainNamePatternInitParameters struct {

	// with pattern for this content provider's domain
	// Use a ends-with pattern for this content provider's domain
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type DomainNamePatternListInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	DomainNamePattern []DomainNamePatternInitParameters `json:"domainNamePattern,omitempty" tf:"domain_name_pattern,omitempty"`
}

type DomainNamePatternListObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	DomainNamePattern []DomainNamePatternObservation `json:"domainNamePattern,omitempty" tf:"domain_name_pattern,omitempty"`
}

type DomainNamePatternListParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	DomainNamePattern []DomainNamePatternParameters `json:"domainNamePattern" tf:"domain_name_pattern,omitempty"`
}

type DomainNamePatternObservation struct {

	// with pattern for this content provider's domain
	// Use a ends-with pattern for this content provider's domain
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type DomainNamePatternParameters struct {

	// with pattern for this content provider's domain
	// Use a ends-with pattern for this content provider's domain
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern" tf:"pattern,omitempty"`
}

type RumProviderBreakdownInitParameters struct {

	// (Block List, Min: 1, Max: 1) Domain name pattern (see below for nested schema)
	// Domain name pattern
	DomainNamePatternList []DomainNamePatternListInitParameters `json:"domainNamePatternList,omitempty" tf:"domain_name_pattern_list,omitempty"`

	// (String) Specify an URL for the provider's brand icon
	// Specify an URL for the provider's brand icon
	IconURL *string `json:"iconUrl,omitempty" tf:"icon_url,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// party detection.
	// Send the patterns of this provider to Dynatrace to help us improve 3rd-party detection.
	ReportPublicImprovement *bool `json:"reportPublicImprovement,omitempty" tf:"report_public_improvement,omitempty"`

	// (String) Resource name
	// Resource name
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	// (String) Possible Values: FirstParty, ThirdParty, Cdn
	// Possible Values: `FirstParty`, `ThirdParty`, `Cdn`
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type RumProviderBreakdownObservation struct {

	// (Block List, Min: 1, Max: 1) Domain name pattern (see below for nested schema)
	// Domain name pattern
	DomainNamePatternList []DomainNamePatternListObservation `json:"domainNamePatternList,omitempty" tf:"domain_name_pattern_list,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Specify an URL for the provider's brand icon
	// Specify an URL for the provider's brand icon
	IconURL *string `json:"iconUrl,omitempty" tf:"icon_url,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// party detection.
	// Send the patterns of this provider to Dynatrace to help us improve 3rd-party detection.
	ReportPublicImprovement *bool `json:"reportPublicImprovement,omitempty" tf:"report_public_improvement,omitempty"`

	// (String) Resource name
	// Resource name
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	// (String) Possible Values: FirstParty, ThirdParty, Cdn
	// Possible Values: `FirstParty`, `ThirdParty`, `Cdn`
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type RumProviderBreakdownParameters struct {

	// (Block List, Min: 1, Max: 1) Domain name pattern (see below for nested schema)
	// Domain name pattern
	// +kubebuilder:validation:Optional
	DomainNamePatternList []DomainNamePatternListParameters `json:"domainNamePatternList,omitempty" tf:"domain_name_pattern_list,omitempty"`

	// (String) Specify an URL for the provider's brand icon
	// Specify an URL for the provider's brand icon
	// +kubebuilder:validation:Optional
	IconURL *string `json:"iconUrl,omitempty" tf:"icon_url,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// +kubebuilder:validation:Optional
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// party detection.
	// Send the patterns of this provider to Dynatrace to help us improve 3rd-party detection.
	// +kubebuilder:validation:Optional
	ReportPublicImprovement *bool `json:"reportPublicImprovement,omitempty" tf:"report_public_improvement,omitempty"`

	// (String) Resource name
	// Resource name
	// +kubebuilder:validation:Optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	// (String) Possible Values: FirstParty, ThirdParty, Cdn
	// Possible Values: `FirstParty`, `ThirdParty`, `Cdn`
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

// RumProviderBreakdownSpec defines the desired state of RumProviderBreakdown
type RumProviderBreakdownSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RumProviderBreakdownParameters `json:"forProvider"`
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
	InitProvider RumProviderBreakdownInitParameters `json:"initProvider,omitempty"`
}

// RumProviderBreakdownStatus defines the observed state of RumProviderBreakdown.
type RumProviderBreakdownStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RumProviderBreakdownObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RumProviderBreakdown is the Schema for the RumProviderBreakdowns API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type RumProviderBreakdown struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domainNamePatternList) || (has(self.initProvider) && has(self.initProvider.domainNamePatternList))",message="spec.forProvider.domainNamePatternList is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.reportPublicImprovement) || (has(self.initProvider) && has(self.initProvider.reportPublicImprovement))",message="spec.forProvider.reportPublicImprovement is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceName) || (has(self.initProvider) && has(self.initProvider.resourceName))",message="spec.forProvider.resourceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceType) || (has(self.initProvider) && has(self.initProvider.resourceType))",message="spec.forProvider.resourceType is a required parameter"
	Spec   RumProviderBreakdownSpec   `json:"spec"`
	Status RumProviderBreakdownStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RumProviderBreakdownList contains a list of RumProviderBreakdowns
type RumProviderBreakdownList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RumProviderBreakdown `json:"items"`
}

// Repository type metadata.
var (
	RumProviderBreakdown_Kind             = "RumProviderBreakdown"
	RumProviderBreakdown_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RumProviderBreakdown_Kind}.String()
	RumProviderBreakdown_KindAPIVersion   = RumProviderBreakdown_Kind + "." + CRDGroupVersion.String()
	RumProviderBreakdown_GroupVersionKind = CRDGroupVersion.WithKind(RumProviderBreakdown_Kind)
)

func init() {
	SchemeBuilder.Register(&RumProviderBreakdown{}, &RumProviderBreakdownList{})
}
