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

type PgAlertingInitParameters struct {

	// (String) Possible Values: ON_INSTANCE_COUNT_VIOLATION, ON_PGI_UNAVAILABILITY
	// Possible Values: `ON_INSTANCE_COUNT_VIOLATION`, `ON_PGI_UNAVAILABILITY`
	AlertingMode *string `json:"alertingMode,omitempty" tf:"alerting_mode,omitempty"`

	// (Boolean) Enable (true) or disable (false) process group availability monitoring
	// Enable (`true`) or disable (`false`) process group availability monitoring
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Number) Open a new problem if the number of active process instances in the group is fewer than X
	// Open a new problem if the number of active process instances in the group is fewer than X
	MinimumInstanceThreshold *float64 `json:"minimumInstanceThreshold,omitempty" tf:"minimum_instance_threshold,omitempty"`

	// (String) The process group ID for availability monitoring
	// The process group ID for availability monitoring
	ProcessGroup *string `json:"processGroup,omitempty" tf:"process_group,omitempty"`
}

type PgAlertingObservation struct {

	// (String) Possible Values: ON_INSTANCE_COUNT_VIOLATION, ON_PGI_UNAVAILABILITY
	// Possible Values: `ON_INSTANCE_COUNT_VIOLATION`, `ON_PGI_UNAVAILABILITY`
	AlertingMode *string `json:"alertingMode,omitempty" tf:"alerting_mode,omitempty"`

	// (Boolean) Enable (true) or disable (false) process group availability monitoring
	// Enable (`true`) or disable (`false`) process group availability monitoring
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number) Open a new problem if the number of active process instances in the group is fewer than X
	// Open a new problem if the number of active process instances in the group is fewer than X
	MinimumInstanceThreshold *float64 `json:"minimumInstanceThreshold,omitempty" tf:"minimum_instance_threshold,omitempty"`

	// (String) The process group ID for availability monitoring
	// The process group ID for availability monitoring
	ProcessGroup *string `json:"processGroup,omitempty" tf:"process_group,omitempty"`
}

type PgAlertingParameters struct {

	// (String) Possible Values: ON_INSTANCE_COUNT_VIOLATION, ON_PGI_UNAVAILABILITY
	// Possible Values: `ON_INSTANCE_COUNT_VIOLATION`, `ON_PGI_UNAVAILABILITY`
	// +kubebuilder:validation:Optional
	AlertingMode *string `json:"alertingMode,omitempty" tf:"alerting_mode,omitempty"`

	// (Boolean) Enable (true) or disable (false) process group availability monitoring
	// Enable (`true`) or disable (`false`) process group availability monitoring
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Number) Open a new problem if the number of active process instances in the group is fewer than X
	// Open a new problem if the number of active process instances in the group is fewer than X
	// +kubebuilder:validation:Optional
	MinimumInstanceThreshold *float64 `json:"minimumInstanceThreshold,omitempty" tf:"minimum_instance_threshold,omitempty"`

	// (String) The process group ID for availability monitoring
	// The process group ID for availability monitoring
	// +kubebuilder:validation:Optional
	ProcessGroup *string `json:"processGroup,omitempty" tf:"process_group,omitempty"`
}

// PgAlertingSpec defines the desired state of PgAlerting
type PgAlertingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PgAlertingParameters `json:"forProvider"`
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
	InitProvider PgAlertingInitParameters `json:"initProvider,omitempty"`
}

// PgAlertingStatus defines the observed state of PgAlerting.
type PgAlertingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PgAlertingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PgAlerting is the Schema for the PgAlertings API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type PgAlerting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.processGroup) || (has(self.initProvider) && has(self.initProvider.processGroup))",message="spec.forProvider.processGroup is a required parameter"
	Spec   PgAlertingSpec   `json:"spec"`
	Status PgAlertingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PgAlertingList contains a list of PgAlertings
type PgAlertingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PgAlerting `json:"items"`
}

// Repository type metadata.
var (
	PgAlerting_Kind             = "PgAlerting"
	PgAlerting_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PgAlerting_Kind}.String()
	PgAlerting_KindAPIVersion   = PgAlerting_Kind + "." + CRDGroupVersion.String()
	PgAlerting_GroupVersionKind = CRDGroupVersion.WithKind(PgAlerting_Kind)
)

func init() {
	SchemeBuilder.Register(&PgAlerting{}, &PgAlertingList{})
}
