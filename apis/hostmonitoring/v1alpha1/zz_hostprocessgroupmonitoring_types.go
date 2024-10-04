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

type HostProcessGroupMonitoringInitParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// (String) Possible Values: DEFAULT, MONITORING_OFF, MONITORING_ON
	// Possible Values: `DEFAULT`, `MONITORING_OFF`, `MONITORING_ON`
	MonitoringState *string `json:"monitoringState,omitempty" tf:"monitoring_state,omitempty"`

	// (String) Process group
	// Process group
	ProcessGroup *string `json:"processGroup,omitempty" tf:"process_group,omitempty"`
}

type HostProcessGroupMonitoringObservation struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Possible Values: DEFAULT, MONITORING_OFF, MONITORING_ON
	// Possible Values: `DEFAULT`, `MONITORING_OFF`, `MONITORING_ON`
	MonitoringState *string `json:"monitoringState,omitempty" tf:"monitoring_state,omitempty"`

	// (String) Process group
	// Process group
	ProcessGroup *string `json:"processGroup,omitempty" tf:"process_group,omitempty"`
}

type HostProcessGroupMonitoringParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// +kubebuilder:validation:Optional
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// (String) Possible Values: DEFAULT, MONITORING_OFF, MONITORING_ON
	// Possible Values: `DEFAULT`, `MONITORING_OFF`, `MONITORING_ON`
	// +kubebuilder:validation:Optional
	MonitoringState *string `json:"monitoringState,omitempty" tf:"monitoring_state,omitempty"`

	// (String) Process group
	// Process group
	// +kubebuilder:validation:Optional
	ProcessGroup *string `json:"processGroup,omitempty" tf:"process_group,omitempty"`
}

// HostProcessGroupMonitoringSpec defines the desired state of HostProcessGroupMonitoring
type HostProcessGroupMonitoringSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostProcessGroupMonitoringParameters `json:"forProvider"`
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
	InitProvider HostProcessGroupMonitoringInitParameters `json:"initProvider,omitempty"`
}

// HostProcessGroupMonitoringStatus defines the observed state of HostProcessGroupMonitoring.
type HostProcessGroupMonitoringStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostProcessGroupMonitoringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HostProcessGroupMonitoring is the Schema for the HostProcessGroupMonitorings API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type HostProcessGroupMonitoring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostId) || (has(self.initProvider) && has(self.initProvider.hostId))",message="spec.forProvider.hostId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.monitoringState) || (has(self.initProvider) && has(self.initProvider.monitoringState))",message="spec.forProvider.monitoringState is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.processGroup) || (has(self.initProvider) && has(self.initProvider.processGroup))",message="spec.forProvider.processGroup is a required parameter"
	Spec   HostProcessGroupMonitoringSpec   `json:"spec"`
	Status HostProcessGroupMonitoringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostProcessGroupMonitoringList contains a list of HostProcessGroupMonitorings
type HostProcessGroupMonitoringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostProcessGroupMonitoring `json:"items"`
}

// Repository type metadata.
var (
	HostProcessGroupMonitoring_Kind             = "HostProcessGroupMonitoring"
	HostProcessGroupMonitoring_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostProcessGroupMonitoring_Kind}.String()
	HostProcessGroupMonitoring_KindAPIVersion   = HostProcessGroupMonitoring_Kind + "." + CRDGroupVersion.String()
	HostProcessGroupMonitoring_GroupVersionKind = CRDGroupVersion.WithKind(HostProcessGroupMonitoring_Kind)
)

func init() {
	SchemeBuilder.Register(&HostProcessGroupMonitoring{}, &HostProcessGroupMonitoringList{})
}
