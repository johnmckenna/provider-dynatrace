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

type EulaSettingsInitParameters struct {

	// (Boolean) Display end user terms to new users logging in to the environment
	// Display end user terms to new users logging in to the environment
	EnableEula *bool `json:"enableEula,omitempty" tf:"enable_eula,omitempty"`

	// (String) The scope of this setting (environment)
	// The scope of this setting (environment)
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type EulaSettingsObservation struct {

	// (Boolean) Display end user terms to new users logging in to the environment
	// Display end user terms to new users logging in to the environment
	EnableEula *bool `json:"enableEula,omitempty" tf:"enable_eula,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The scope of this setting (environment)
	// The scope of this setting (environment)
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type EulaSettingsParameters struct {

	// (Boolean) Display end user terms to new users logging in to the environment
	// Display end user terms to new users logging in to the environment
	// +kubebuilder:validation:Optional
	EnableEula *bool `json:"enableEula,omitempty" tf:"enable_eula,omitempty"`

	// (String) The scope of this setting (environment)
	// The scope of this setting (environment)
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// EulaSettingsSpec defines the desired state of EulaSettings
type EulaSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EulaSettingsParameters `json:"forProvider"`
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
	InitProvider EulaSettingsInitParameters `json:"initProvider,omitempty"`
}

// EulaSettingsStatus defines the observed state of EulaSettings.
type EulaSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EulaSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EulaSettings is the Schema for the EulaSettingss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type EulaSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enableEula) || (has(self.initProvider) && has(self.initProvider.enableEula))",message="spec.forProvider.enableEula is a required parameter"
	Spec   EulaSettingsSpec   `json:"spec"`
	Status EulaSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EulaSettingsList contains a list of EulaSettingss
type EulaSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EulaSettings `json:"items"`
}

// Repository type metadata.
var (
	EulaSettings_Kind             = "EulaSettings"
	EulaSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EulaSettings_Kind}.String()
	EulaSettings_KindAPIVersion   = EulaSettings_Kind + "." + CRDGroupVersion.String()
	EulaSettings_GroupVersionKind = CRDGroupVersion.WithKind(EulaSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&EulaSettings{}, &EulaSettingsList{})
}