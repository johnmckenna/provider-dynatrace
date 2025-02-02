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

type InfraopsAppSettingsInitParameters struct {

	// (Boolean) When set to true, the app will display monitoring candidates in the Hosts table
	// When set to true, the app will display monitoring candidates in the Hosts table
	ShowMonitoringCandidates *bool `json:"showMonitoringCandidates,omitempty" tf:"show_monitoring_candidates,omitempty"`

	// (Boolean) When set to true, the app will display app only hosts in the Hosts table
	// When set to true, the app will display app only hosts in the Hosts table
	ShowStandaloneHosts *bool `json:"showStandaloneHosts,omitempty" tf:"show_standalone_hosts,omitempty"`
}

type InfraopsAppSettingsObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) When set to true, the app will display monitoring candidates in the Hosts table
	// When set to true, the app will display monitoring candidates in the Hosts table
	ShowMonitoringCandidates *bool `json:"showMonitoringCandidates,omitempty" tf:"show_monitoring_candidates,omitempty"`

	// (Boolean) When set to true, the app will display app only hosts in the Hosts table
	// When set to true, the app will display app only hosts in the Hosts table
	ShowStandaloneHosts *bool `json:"showStandaloneHosts,omitempty" tf:"show_standalone_hosts,omitempty"`
}

type InfraopsAppSettingsParameters struct {

	// (Boolean) When set to true, the app will display monitoring candidates in the Hosts table
	// When set to true, the app will display monitoring candidates in the Hosts table
	// +kubebuilder:validation:Optional
	ShowMonitoringCandidates *bool `json:"showMonitoringCandidates,omitempty" tf:"show_monitoring_candidates,omitempty"`

	// (Boolean) When set to true, the app will display app only hosts in the Hosts table
	// When set to true, the app will display app only hosts in the Hosts table
	// +kubebuilder:validation:Optional
	ShowStandaloneHosts *bool `json:"showStandaloneHosts,omitempty" tf:"show_standalone_hosts,omitempty"`
}

// InfraopsAppSettingsSpec defines the desired state of InfraopsAppSettings
type InfraopsAppSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InfraopsAppSettingsParameters `json:"forProvider"`
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
	InitProvider InfraopsAppSettingsInitParameters `json:"initProvider,omitempty"`
}

// InfraopsAppSettingsStatus defines the observed state of InfraopsAppSettings.
type InfraopsAppSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InfraopsAppSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InfraopsAppSettings is the Schema for the InfraopsAppSettingss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type InfraopsAppSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.showMonitoringCandidates) || (has(self.initProvider) && has(self.initProvider.showMonitoringCandidates))",message="spec.forProvider.showMonitoringCandidates is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.showStandaloneHosts) || (has(self.initProvider) && has(self.initProvider.showStandaloneHosts))",message="spec.forProvider.showStandaloneHosts is a required parameter"
	Spec   InfraopsAppSettingsSpec   `json:"spec"`
	Status InfraopsAppSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InfraopsAppSettingsList contains a list of InfraopsAppSettingss
type InfraopsAppSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InfraopsAppSettings `json:"items"`
}

// Repository type metadata.
var (
	InfraopsAppSettings_Kind             = "InfraopsAppSettings"
	InfraopsAppSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InfraopsAppSettings_Kind}.String()
	InfraopsAppSettings_KindAPIVersion   = InfraopsAppSettings_Kind + "." + CRDGroupVersion.String()
	InfraopsAppSettings_GroupVersionKind = CRDGroupVersion.WithKind(InfraopsAppSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&InfraopsAppSettings{}, &InfraopsAppSettingsList{})
}
