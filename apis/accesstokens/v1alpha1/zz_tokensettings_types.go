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

type TokenSettingsInitParameters struct {

	// (Boolean) Check out this blog post to find out more about the new Dynatrace API token format.
	// Check out this [blog post](http://www.dynatrace.com/blog/further-increased-security-of-your-api-tokens-by-automating-token-protection/) to find out more about the new Dynatrace API token format.
	NewTokenFormat *bool `json:"newTokenFormat,omitempty" tf:"new_token_format,omitempty"`

	// (Boolean) Allow users of this environment to generate personal access tokens based on user permissions.
	// Note that existing personal access tokens will become unusable while this setting is disabled.
	// Allow users of this environment to generate personal access tokens based on user permissions.
	// Note that existing personal access tokens will become unusable while this setting is disabled.
	PersonalTokens *bool `json:"personalTokens,omitempty" tf:"personal_tokens,omitempty"`
}

type TokenSettingsObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Check out this blog post to find out more about the new Dynatrace API token format.
	// Check out this [blog post](http://www.dynatrace.com/blog/further-increased-security-of-your-api-tokens-by-automating-token-protection/) to find out more about the new Dynatrace API token format.
	NewTokenFormat *bool `json:"newTokenFormat,omitempty" tf:"new_token_format,omitempty"`

	// (Boolean) Allow users of this environment to generate personal access tokens based on user permissions.
	// Note that existing personal access tokens will become unusable while this setting is disabled.
	// Allow users of this environment to generate personal access tokens based on user permissions.
	// Note that existing personal access tokens will become unusable while this setting is disabled.
	PersonalTokens *bool `json:"personalTokens,omitempty" tf:"personal_tokens,omitempty"`
}

type TokenSettingsParameters struct {

	// (Boolean) Check out this blog post to find out more about the new Dynatrace API token format.
	// Check out this [blog post](http://www.dynatrace.com/blog/further-increased-security-of-your-api-tokens-by-automating-token-protection/) to find out more about the new Dynatrace API token format.
	// +kubebuilder:validation:Optional
	NewTokenFormat *bool `json:"newTokenFormat,omitempty" tf:"new_token_format,omitempty"`

	// (Boolean) Allow users of this environment to generate personal access tokens based on user permissions.
	// Note that existing personal access tokens will become unusable while this setting is disabled.
	// Allow users of this environment to generate personal access tokens based on user permissions.
	// Note that existing personal access tokens will become unusable while this setting is disabled.
	// +kubebuilder:validation:Optional
	PersonalTokens *bool `json:"personalTokens,omitempty" tf:"personal_tokens,omitempty"`
}

// TokenSettingsSpec defines the desired state of TokenSettings
type TokenSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TokenSettingsParameters `json:"forProvider"`
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
	InitProvider TokenSettingsInitParameters `json:"initProvider,omitempty"`
}

// TokenSettingsStatus defines the observed state of TokenSettings.
type TokenSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TokenSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TokenSettings is the Schema for the TokenSettingss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type TokenSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.newTokenFormat) || (has(self.initProvider) && has(self.initProvider.newTokenFormat))",message="spec.forProvider.newTokenFormat is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.personalTokens) || (has(self.initProvider) && has(self.initProvider.personalTokens))",message="spec.forProvider.personalTokens is a required parameter"
	Spec   TokenSettingsSpec   `json:"spec"`
	Status TokenSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TokenSettingsList contains a list of TokenSettingss
type TokenSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TokenSettings `json:"items"`
}

// Repository type metadata.
var (
	TokenSettings_Kind             = "TokenSettings"
	TokenSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TokenSettings_Kind}.String()
	TokenSettings_KindAPIVersion   = TokenSettings_Kind + "." + CRDGroupVersion.String()
	TokenSettings_GroupVersionKind = CRDGroupVersion.WithKind(TokenSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&TokenSettings{}, &TokenSettingsList{})
}
