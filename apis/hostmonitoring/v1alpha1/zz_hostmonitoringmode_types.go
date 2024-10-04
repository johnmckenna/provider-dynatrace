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

type HostMonitoringModeInitParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// (String) Possible Values: DISCOVERY, FULL_STACK, INFRA_ONLY
	// Possible Values: `DISCOVERY`, `FULL_STACK`, `INFRA_ONLY`
	MonitoringMode *string `json:"monitoringMode,omitempty" tf:"monitoring_mode,omitempty"`
}

type HostMonitoringModeObservation struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Possible Values: DISCOVERY, FULL_STACK, INFRA_ONLY
	// Possible Values: `DISCOVERY`, `FULL_STACK`, `INFRA_ONLY`
	MonitoringMode *string `json:"monitoringMode,omitempty" tf:"monitoring_mode,omitempty"`
}

type HostMonitoringModeParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// +kubebuilder:validation:Optional
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// (String) Possible Values: DISCOVERY, FULL_STACK, INFRA_ONLY
	// Possible Values: `DISCOVERY`, `FULL_STACK`, `INFRA_ONLY`
	// +kubebuilder:validation:Optional
	MonitoringMode *string `json:"monitoringMode,omitempty" tf:"monitoring_mode,omitempty"`
}

// HostMonitoringModeSpec defines the desired state of HostMonitoringMode
type HostMonitoringModeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostMonitoringModeParameters `json:"forProvider"`
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
	InitProvider HostMonitoringModeInitParameters `json:"initProvider,omitempty"`
}

// HostMonitoringModeStatus defines the observed state of HostMonitoringMode.
type HostMonitoringModeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostMonitoringModeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HostMonitoringMode is the Schema for the HostMonitoringModes API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type HostMonitoringMode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostId) || (has(self.initProvider) && has(self.initProvider.hostId))",message="spec.forProvider.hostId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.monitoringMode) || (has(self.initProvider) && has(self.initProvider.monitoringMode))",message="spec.forProvider.monitoringMode is a required parameter"
	Spec   HostMonitoringModeSpec   `json:"spec"`
	Status HostMonitoringModeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostMonitoringModeList contains a list of HostMonitoringModes
type HostMonitoringModeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostMonitoringMode `json:"items"`
}

// Repository type metadata.
var (
	HostMonitoringMode_Kind             = "HostMonitoringMode"
	HostMonitoringMode_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostMonitoringMode_Kind}.String()
	HostMonitoringMode_KindAPIVersion   = HostMonitoringMode_Kind + "." + CRDGroupVersion.String()
	HostMonitoringMode_GroupVersionKind = CRDGroupVersion.WithKind(HostMonitoringMode_Kind)
)

func init() {
	SchemeBuilder.Register(&HostMonitoringMode{}, &HostMonitoringModeList{})
}