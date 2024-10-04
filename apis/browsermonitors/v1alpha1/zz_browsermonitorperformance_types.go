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

type BrowserMonitorPerformanceInitParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The scope of this setting (SYNTHETIC_TEST)
	// The scope of this setting (SYNTHETIC_TEST)
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// (Block List, Max: 1) Performance thresholds (see below for nested schema)
	// Performance thresholds
	Thresholds []BrowserMonitorPerformanceThresholdsInitParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type BrowserMonitorPerformanceObservation struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The scope of this setting (SYNTHETIC_TEST)
	// The scope of this setting (SYNTHETIC_TEST)
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// (Block List, Max: 1) Performance thresholds (see below for nested schema)
	// Performance thresholds
	Thresholds []BrowserMonitorPerformanceThresholdsObservation `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type BrowserMonitorPerformanceParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The scope of this setting (SYNTHETIC_TEST)
	// The scope of this setting (SYNTHETIC_TEST)
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// (Block List, Max: 1) Performance thresholds (see below for nested schema)
	// Performance thresholds
	// +kubebuilder:validation:Optional
	Thresholds []BrowserMonitorPerformanceThresholdsParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type BrowserMonitorPerformanceThresholdsInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Threshold []ThresholdsThresholdInitParameters `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type BrowserMonitorPerformanceThresholdsObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Threshold []ThresholdsThresholdObservation `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type BrowserMonitorPerformanceThresholdsParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Threshold []ThresholdsThresholdParameters `json:"threshold" tf:"threshold,omitempty"`
}

type ThresholdsThresholdInitParameters struct {

	// (String) Synthetic event
	// Synthetic event
	Event *string `json:"event,omitempty" tf:"event,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	// Threshold (in seconds)
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type ThresholdsThresholdObservation struct {

	// (String) Synthetic event
	// Synthetic event
	Event *string `json:"event,omitempty" tf:"event,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	// Threshold (in seconds)
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type ThresholdsThresholdParameters struct {

	// (String) Synthetic event
	// Synthetic event
	// +kubebuilder:validation:Optional
	Event *string `json:"event" tf:"event,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	// Threshold (in seconds)
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`
}

// BrowserMonitorPerformanceSpec defines the desired state of BrowserMonitorPerformance
type BrowserMonitorPerformanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BrowserMonitorPerformanceParameters `json:"forProvider"`
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
	InitProvider BrowserMonitorPerformanceInitParameters `json:"initProvider,omitempty"`
}

// BrowserMonitorPerformanceStatus defines the observed state of BrowserMonitorPerformance.
type BrowserMonitorPerformanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BrowserMonitorPerformanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BrowserMonitorPerformance is the Schema for the BrowserMonitorPerformances API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type BrowserMonitorPerformance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scope) || (has(self.initProvider) && has(self.initProvider.scope))",message="spec.forProvider.scope is a required parameter"
	Spec   BrowserMonitorPerformanceSpec   `json:"spec"`
	Status BrowserMonitorPerformanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BrowserMonitorPerformanceList contains a list of BrowserMonitorPerformances
type BrowserMonitorPerformanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BrowserMonitorPerformance `json:"items"`
}

// Repository type metadata.
var (
	BrowserMonitorPerformance_Kind             = "BrowserMonitorPerformance"
	BrowserMonitorPerformance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BrowserMonitorPerformance_Kind}.String()
	BrowserMonitorPerformance_KindAPIVersion   = BrowserMonitorPerformance_Kind + "." + CRDGroupVersion.String()
	BrowserMonitorPerformance_GroupVersionKind = CRDGroupVersion.WithKind(BrowserMonitorPerformance_Kind)
)

func init() {
	SchemeBuilder.Register(&BrowserMonitorPerformance{}, &BrowserMonitorPerformanceList{})
}
