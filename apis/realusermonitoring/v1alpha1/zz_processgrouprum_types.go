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

type ProcessGroupRumInitParameters struct {

	// (Boolean) Allows OneAgent to:
	// Allows OneAgent to:
	// * automatically inject the RUM JavaScript tag into each page delivered by this process group
	// * provide the necessary info to correlate RUM data with server-side PurePaths
	// * forward beacons to the cluster
	// * deliver the monitoring code
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// PROCESS_GROUP-XXXXXXXXXXXXXXXX
	// The scope of this setting - PROCESS_GROUP-XXXXXXXXXXXXXXXX
	ProcessGroupID *string `json:"processGroupId,omitempty" tf:"process_group_id,omitempty"`
}

type ProcessGroupRumObservation struct {

	// (Boolean) Allows OneAgent to:
	// Allows OneAgent to:
	// * automatically inject the RUM JavaScript tag into each page delivered by this process group
	// * provide the necessary info to correlate RUM data with server-side PurePaths
	// * forward beacons to the cluster
	// * deliver the monitoring code
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// PROCESS_GROUP-XXXXXXXXXXXXXXXX
	// The scope of this setting - PROCESS_GROUP-XXXXXXXXXXXXXXXX
	ProcessGroupID *string `json:"processGroupId,omitempty" tf:"process_group_id,omitempty"`
}

type ProcessGroupRumParameters struct {

	// (Boolean) Allows OneAgent to:
	// Allows OneAgent to:
	// * automatically inject the RUM JavaScript tag into each page delivered by this process group
	// * provide the necessary info to correlate RUM data with server-side PurePaths
	// * forward beacons to the cluster
	// * deliver the monitoring code
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// PROCESS_GROUP-XXXXXXXXXXXXXXXX
	// The scope of this setting - PROCESS_GROUP-XXXXXXXXXXXXXXXX
	// +kubebuilder:validation:Optional
	ProcessGroupID *string `json:"processGroupId,omitempty" tf:"process_group_id,omitempty"`
}

// ProcessGroupRumSpec defines the desired state of ProcessGroupRum
type ProcessGroupRumSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProcessGroupRumParameters `json:"forProvider"`
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
	InitProvider ProcessGroupRumInitParameters `json:"initProvider,omitempty"`
}

// ProcessGroupRumStatus defines the observed state of ProcessGroupRum.
type ProcessGroupRumStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProcessGroupRumObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProcessGroupRum is the Schema for the ProcessGroupRums API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type ProcessGroupRum struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enable) || (has(self.initProvider) && has(self.initProvider.enable))",message="spec.forProvider.enable is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.processGroupId) || (has(self.initProvider) && has(self.initProvider.processGroupId))",message="spec.forProvider.processGroupId is a required parameter"
	Spec   ProcessGroupRumSpec   `json:"spec"`
	Status ProcessGroupRumStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProcessGroupRumList contains a list of ProcessGroupRums
type ProcessGroupRumList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProcessGroupRum `json:"items"`
}

// Repository type metadata.
var (
	ProcessGroupRum_Kind             = "ProcessGroupRum"
	ProcessGroupRum_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProcessGroupRum_Kind}.String()
	ProcessGroupRum_KindAPIVersion   = ProcessGroupRum_Kind + "." + CRDGroupVersion.String()
	ProcessGroupRum_GroupVersionKind = CRDGroupVersion.WithKind(ProcessGroupRum_Kind)
)

func init() {
	SchemeBuilder.Register(&ProcessGroupRum{}, &ProcessGroupRumList{})
}
