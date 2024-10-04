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

type OneagentSideMaskingInitParameters struct {

	// (Boolean) Exclude email addresses from URLs
	// Exclude email addresses from URLs
	IsEmailMaskingEnabled *bool `json:"isEmailMaskingEnabled,omitempty" tf:"is_email_masking_enabled,omitempty"`

	// (Boolean) Exclude IBANs and payment card numbers from URLs
	// Exclude IBANs and payment card numbers from URLs
	IsFinancialMaskingEnabled *bool `json:"isFinancialMaskingEnabled,omitempty" tf:"is_financial_masking_enabled,omitempty"`

	// (Boolean) Exclude hexadecimal IDs and consecutive numbers above 5 digits from URLs
	// Exclude hexadecimal IDs and consecutive numbers above 5 digits from URLs
	IsNumbersMaskingEnabled *bool `json:"isNumbersMaskingEnabled,omitempty" tf:"is_numbers_masking_enabled,omitempty"`

	// (Boolean) Exclude query parameters from URLs and web requests
	// Exclude query parameters from URLs and web requests
	IsQueryMaskingEnabled *bool `json:"isQueryMaskingEnabled,omitempty" tf:"is_query_masking_enabled,omitempty"`

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ProcessGroupID *string `json:"processGroupId,omitempty" tf:"process_group_id,omitempty"`
}

type OneagentSideMaskingObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Exclude email addresses from URLs
	// Exclude email addresses from URLs
	IsEmailMaskingEnabled *bool `json:"isEmailMaskingEnabled,omitempty" tf:"is_email_masking_enabled,omitempty"`

	// (Boolean) Exclude IBANs and payment card numbers from URLs
	// Exclude IBANs and payment card numbers from URLs
	IsFinancialMaskingEnabled *bool `json:"isFinancialMaskingEnabled,omitempty" tf:"is_financial_masking_enabled,omitempty"`

	// (Boolean) Exclude hexadecimal IDs and consecutive numbers above 5 digits from URLs
	// Exclude hexadecimal IDs and consecutive numbers above 5 digits from URLs
	IsNumbersMaskingEnabled *bool `json:"isNumbersMaskingEnabled,omitempty" tf:"is_numbers_masking_enabled,omitempty"`

	// (Boolean) Exclude query parameters from URLs and web requests
	// Exclude query parameters from URLs and web requests
	IsQueryMaskingEnabled *bool `json:"isQueryMaskingEnabled,omitempty" tf:"is_query_masking_enabled,omitempty"`

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ProcessGroupID *string `json:"processGroupId,omitempty" tf:"process_group_id,omitempty"`
}

type OneagentSideMaskingParameters struct {

	// (Boolean) Exclude email addresses from URLs
	// Exclude email addresses from URLs
	// +kubebuilder:validation:Optional
	IsEmailMaskingEnabled *bool `json:"isEmailMaskingEnabled,omitempty" tf:"is_email_masking_enabled,omitempty"`

	// (Boolean) Exclude IBANs and payment card numbers from URLs
	// Exclude IBANs and payment card numbers from URLs
	// +kubebuilder:validation:Optional
	IsFinancialMaskingEnabled *bool `json:"isFinancialMaskingEnabled,omitempty" tf:"is_financial_masking_enabled,omitempty"`

	// (Boolean) Exclude hexadecimal IDs and consecutive numbers above 5 digits from URLs
	// Exclude hexadecimal IDs and consecutive numbers above 5 digits from URLs
	// +kubebuilder:validation:Optional
	IsNumbersMaskingEnabled *bool `json:"isNumbersMaskingEnabled,omitempty" tf:"is_numbers_masking_enabled,omitempty"`

	// (Boolean) Exclude query parameters from URLs and web requests
	// Exclude query parameters from URLs and web requests
	// +kubebuilder:validation:Optional
	IsQueryMaskingEnabled *bool `json:"isQueryMaskingEnabled,omitempty" tf:"is_query_masking_enabled,omitempty"`

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// +kubebuilder:validation:Optional
	ProcessGroupID *string `json:"processGroupId,omitempty" tf:"process_group_id,omitempty"`
}

// OneagentSideMaskingSpec defines the desired state of OneagentSideMasking
type OneagentSideMaskingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OneagentSideMaskingParameters `json:"forProvider"`
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
	InitProvider OneagentSideMaskingInitParameters `json:"initProvider,omitempty"`
}

// OneagentSideMaskingStatus defines the observed state of OneagentSideMasking.
type OneagentSideMaskingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OneagentSideMaskingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OneagentSideMasking is the Schema for the OneagentSideMaskings API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type OneagentSideMasking struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isEmailMaskingEnabled) || (has(self.initProvider) && has(self.initProvider.isEmailMaskingEnabled))",message="spec.forProvider.isEmailMaskingEnabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isFinancialMaskingEnabled) || (has(self.initProvider) && has(self.initProvider.isFinancialMaskingEnabled))",message="spec.forProvider.isFinancialMaskingEnabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isNumbersMaskingEnabled) || (has(self.initProvider) && has(self.initProvider.isNumbersMaskingEnabled))",message="spec.forProvider.isNumbersMaskingEnabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isQueryMaskingEnabled) || (has(self.initProvider) && has(self.initProvider.isQueryMaskingEnabled))",message="spec.forProvider.isQueryMaskingEnabled is a required parameter"
	Spec   OneagentSideMaskingSpec   `json:"spec"`
	Status OneagentSideMaskingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OneagentSideMaskingList contains a list of OneagentSideMaskings
type OneagentSideMaskingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OneagentSideMasking `json:"items"`
}

// Repository type metadata.
var (
	OneagentSideMasking_Kind             = "OneagentSideMasking"
	OneagentSideMasking_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OneagentSideMasking_Kind}.String()
	OneagentSideMasking_KindAPIVersion   = OneagentSideMasking_Kind + "." + CRDGroupVersion.String()
	OneagentSideMasking_GroupVersionKind = CRDGroupVersion.WithKind(OneagentSideMasking_Kind)
)

func init() {
	SchemeBuilder.Register(&OneagentSideMasking{}, &OneagentSideMaskingList{})
}
