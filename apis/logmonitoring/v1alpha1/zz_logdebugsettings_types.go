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

type LogDebugSettingsInitParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type LogDebugSettingsObservation struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LogDebugSettingsParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

// LogDebugSettingsSpec defines the desired state of LogDebugSettings
type LogDebugSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogDebugSettingsParameters `json:"forProvider"`
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
	InitProvider LogDebugSettingsInitParameters `json:"initProvider,omitempty"`
}

// LogDebugSettingsStatus defines the observed state of LogDebugSettings.
type LogDebugSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogDebugSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LogDebugSettings is the Schema for the LogDebugSettingss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type LogDebugSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	Spec   LogDebugSettingsSpec   `json:"spec"`
	Status LogDebugSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogDebugSettingsList contains a list of LogDebugSettingss
type LogDebugSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogDebugSettings `json:"items"`
}

// Repository type metadata.
var (
	LogDebugSettings_Kind             = "LogDebugSettings"
	LogDebugSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogDebugSettings_Kind}.String()
	LogDebugSettings_KindAPIVersion   = LogDebugSettings_Kind + "." + CRDGroupVersion.String()
	LogDebugSettings_GroupVersionKind = CRDGroupVersion.WithKind(LogDebugSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&LogDebugSettings{}, &LogDebugSettingsList{})
}