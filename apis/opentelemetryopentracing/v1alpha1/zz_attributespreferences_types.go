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

type AttributesPreferencesInitParameters struct {

	// (String) Possible Values: ALLOW_ALL_ATTRIBUTES, BLOCK_ALL_ATTRIBUTES
	// Possible Values: `ALLOW_ALL_ATTRIBUTES`, `BLOCK_ALL_ATTRIBUTES`
	PersistenceMode *string `json:"persistenceMode,omitempty" tf:"persistence_mode,omitempty"`
}

type AttributesPreferencesObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Possible Values: ALLOW_ALL_ATTRIBUTES, BLOCK_ALL_ATTRIBUTES
	// Possible Values: `ALLOW_ALL_ATTRIBUTES`, `BLOCK_ALL_ATTRIBUTES`
	PersistenceMode *string `json:"persistenceMode,omitempty" tf:"persistence_mode,omitempty"`
}

type AttributesPreferencesParameters struct {

	// (String) Possible Values: ALLOW_ALL_ATTRIBUTES, BLOCK_ALL_ATTRIBUTES
	// Possible Values: `ALLOW_ALL_ATTRIBUTES`, `BLOCK_ALL_ATTRIBUTES`
	// +kubebuilder:validation:Optional
	PersistenceMode *string `json:"persistenceMode,omitempty" tf:"persistence_mode,omitempty"`
}

// AttributesPreferencesSpec defines the desired state of AttributesPreferences
type AttributesPreferencesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AttributesPreferencesParameters `json:"forProvider"`
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
	InitProvider AttributesPreferencesInitParameters `json:"initProvider,omitempty"`
}

// AttributesPreferencesStatus defines the observed state of AttributesPreferences.
type AttributesPreferencesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AttributesPreferencesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AttributesPreferences is the Schema for the AttributesPreferencess API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type AttributesPreferences struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.persistenceMode) || (has(self.initProvider) && has(self.initProvider.persistenceMode))",message="spec.forProvider.persistenceMode is a required parameter"
	Spec   AttributesPreferencesSpec   `json:"spec"`
	Status AttributesPreferencesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttributesPreferencesList contains a list of AttributesPreferencess
type AttributesPreferencesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AttributesPreferences `json:"items"`
}

// Repository type metadata.
var (
	AttributesPreferences_Kind             = "AttributesPreferences"
	AttributesPreferences_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AttributesPreferences_Kind}.String()
	AttributesPreferences_KindAPIVersion   = AttributesPreferences_Kind + "." + CRDGroupVersion.String()
	AttributesPreferences_GroupVersionKind = CRDGroupVersion.WithKind(AttributesPreferences_Kind)
)

func init() {
	SchemeBuilder.Register(&AttributesPreferences{}, &AttributesPreferencesList{})
}