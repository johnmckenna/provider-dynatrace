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

type ProcessVisibilityInitParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Number) The maximum amount of processes that host may report is 100
	// The maximum amount of processes that host may report is **100**
	MaxProcesses *float64 `json:"maxProcesses,omitempty" tf:"max_processes,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type ProcessVisibilityObservation struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number) The maximum amount of processes that host may report is 100
	// The maximum amount of processes that host may report is **100**
	MaxProcesses *float64 `json:"maxProcesses,omitempty" tf:"max_processes,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type ProcessVisibilityParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Number) The maximum amount of processes that host may report is 100
	// The maximum amount of processes that host may report is **100**
	// +kubebuilder:validation:Optional
	MaxProcesses *float64 `json:"maxProcesses,omitempty" tf:"max_processes,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// ProcessVisibilitySpec defines the desired state of ProcessVisibility
type ProcessVisibilitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProcessVisibilityParameters `json:"forProvider"`
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
	InitProvider ProcessVisibilityInitParameters `json:"initProvider,omitempty"`
}

// ProcessVisibilityStatus defines the observed state of ProcessVisibility.
type ProcessVisibilityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProcessVisibilityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProcessVisibility is the Schema for the ProcessVisibilitys API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type ProcessVisibility struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.maxProcesses) || (has(self.initProvider) && has(self.initProvider.maxProcesses))",message="spec.forProvider.maxProcesses is a required parameter"
	Spec   ProcessVisibilitySpec   `json:"spec"`
	Status ProcessVisibilityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProcessVisibilityList contains a list of ProcessVisibilitys
type ProcessVisibilityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProcessVisibility `json:"items"`
}

// Repository type metadata.
var (
	ProcessVisibility_Kind             = "ProcessVisibility"
	ProcessVisibility_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProcessVisibility_Kind}.String()
	ProcessVisibility_KindAPIVersion   = ProcessVisibility_Kind + "." + CRDGroupVersion.String()
	ProcessVisibility_GroupVersionKind = CRDGroupVersion.WithKind(ProcessVisibility_Kind)
)

func init() {
	SchemeBuilder.Register(&ProcessVisibility{}, &ProcessVisibilityList{})
}