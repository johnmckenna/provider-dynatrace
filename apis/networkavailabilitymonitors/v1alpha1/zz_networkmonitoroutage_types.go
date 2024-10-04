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

type NetworkMonitorOutageInitParameters struct {

	// (Number) Alert if all locations are unable to access my target address
	// Alert if all locations are unable to access my target address
	GlobalConsecutiveOutageCountThreshold *float64 `json:"globalConsecutiveOutageCountThreshold,omitempty" tf:"global_consecutive_outage_count_threshold,omitempty"`

	// (Boolean) Generate a problem and send an alert when the monitor is unavailable at all configured locations.
	// Generate a problem and send an alert when the monitor is unavailable at all configured locations.
	GlobalOutages *bool `json:"globalOutages,omitempty" tf:"global_outages,omitempty"`

	// (Number) are unable to access my target address
	// are unable to access my target address
	LocalConsecutiveOutageCountThreshold *float64 `json:"localConsecutiveOutageCountThreshold,omitempty" tf:"local_consecutive_outage_count_threshold,omitempty"`

	// (Number) Alert if at least
	// Alert if at least
	LocalLocationOutageCountThreshold *float64 `json:"localLocationOutageCountThreshold,omitempty" tf:"local_location_outage_count_threshold,omitempty"`

	// (Boolean) Generate a problem and send an alert when the monitor is unavailable for one or more consecutive runs at any location.
	// Generate a problem and send an alert when the monitor is unavailable for one or more consecutive runs at any location.
	LocalOutages *bool `json:"localOutages,omitempty" tf:"local_outages,omitempty"`

	// (String) The scope of this setting (MULTIPROTOCOL_MONITOR). Omit this property if you want to cover the whole environment.
	// The scope of this setting (MULTIPROTOCOL_MONITOR). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type NetworkMonitorOutageObservation struct {

	// (Number) Alert if all locations are unable to access my target address
	// Alert if all locations are unable to access my target address
	GlobalConsecutiveOutageCountThreshold *float64 `json:"globalConsecutiveOutageCountThreshold,omitempty" tf:"global_consecutive_outage_count_threshold,omitempty"`

	// (Boolean) Generate a problem and send an alert when the monitor is unavailable at all configured locations.
	// Generate a problem and send an alert when the monitor is unavailable at all configured locations.
	GlobalOutages *bool `json:"globalOutages,omitempty" tf:"global_outages,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number) are unable to access my target address
	// are unable to access my target address
	LocalConsecutiveOutageCountThreshold *float64 `json:"localConsecutiveOutageCountThreshold,omitempty" tf:"local_consecutive_outage_count_threshold,omitempty"`

	// (Number) Alert if at least
	// Alert if at least
	LocalLocationOutageCountThreshold *float64 `json:"localLocationOutageCountThreshold,omitempty" tf:"local_location_outage_count_threshold,omitempty"`

	// (Boolean) Generate a problem and send an alert when the monitor is unavailable for one or more consecutive runs at any location.
	// Generate a problem and send an alert when the monitor is unavailable for one or more consecutive runs at any location.
	LocalOutages *bool `json:"localOutages,omitempty" tf:"local_outages,omitempty"`

	// (String) The scope of this setting (MULTIPROTOCOL_MONITOR). Omit this property if you want to cover the whole environment.
	// The scope of this setting (MULTIPROTOCOL_MONITOR). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type NetworkMonitorOutageParameters struct {

	// (Number) Alert if all locations are unable to access my target address
	// Alert if all locations are unable to access my target address
	// +kubebuilder:validation:Optional
	GlobalConsecutiveOutageCountThreshold *float64 `json:"globalConsecutiveOutageCountThreshold,omitempty" tf:"global_consecutive_outage_count_threshold,omitempty"`

	// (Boolean) Generate a problem and send an alert when the monitor is unavailable at all configured locations.
	// Generate a problem and send an alert when the monitor is unavailable at all configured locations.
	// +kubebuilder:validation:Optional
	GlobalOutages *bool `json:"globalOutages,omitempty" tf:"global_outages,omitempty"`

	// (Number) are unable to access my target address
	// are unable to access my target address
	// +kubebuilder:validation:Optional
	LocalConsecutiveOutageCountThreshold *float64 `json:"localConsecutiveOutageCountThreshold,omitempty" tf:"local_consecutive_outage_count_threshold,omitempty"`

	// (Number) Alert if at least
	// Alert if at least
	// +kubebuilder:validation:Optional
	LocalLocationOutageCountThreshold *float64 `json:"localLocationOutageCountThreshold,omitempty" tf:"local_location_outage_count_threshold,omitempty"`

	// (Boolean) Generate a problem and send an alert when the monitor is unavailable for one or more consecutive runs at any location.
	// Generate a problem and send an alert when the monitor is unavailable for one or more consecutive runs at any location.
	// +kubebuilder:validation:Optional
	LocalOutages *bool `json:"localOutages,omitempty" tf:"local_outages,omitempty"`

	// (String) The scope of this setting (MULTIPROTOCOL_MONITOR). Omit this property if you want to cover the whole environment.
	// The scope of this setting (MULTIPROTOCOL_MONITOR). Omit this property if you want to cover the whole environment.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// NetworkMonitorOutageSpec defines the desired state of NetworkMonitorOutage
type NetworkMonitorOutageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkMonitorOutageParameters `json:"forProvider"`
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
	InitProvider NetworkMonitorOutageInitParameters `json:"initProvider,omitempty"`
}

// NetworkMonitorOutageStatus defines the observed state of NetworkMonitorOutage.
type NetworkMonitorOutageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkMonitorOutageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NetworkMonitorOutage is the Schema for the NetworkMonitorOutages API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type NetworkMonitorOutage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.globalOutages) || (has(self.initProvider) && has(self.initProvider.globalOutages))",message="spec.forProvider.globalOutages is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.localOutages) || (has(self.initProvider) && has(self.initProvider.localOutages))",message="spec.forProvider.localOutages is a required parameter"
	Spec   NetworkMonitorOutageSpec   `json:"spec"`
	Status NetworkMonitorOutageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkMonitorOutageList contains a list of NetworkMonitorOutages
type NetworkMonitorOutageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkMonitorOutage `json:"items"`
}

// Repository type metadata.
var (
	NetworkMonitorOutage_Kind             = "NetworkMonitorOutage"
	NetworkMonitorOutage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkMonitorOutage_Kind}.String()
	NetworkMonitorOutage_KindAPIVersion   = NetworkMonitorOutage_Kind + "." + CRDGroupVersion.String()
	NetworkMonitorOutage_GroupVersionKind = CRDGroupVersion.WithKind(NetworkMonitorOutage_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkMonitorOutage{}, &NetworkMonitorOutageList{})
}