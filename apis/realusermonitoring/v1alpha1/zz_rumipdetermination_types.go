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

type RumIPDeterminationInitParameters struct {

	// (String) Client IP header name
	// Client IP header name
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`
}

type RumIPDeterminationObservation struct {

	// (String) Client IP header name
	// Client IP header name
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`
}

type RumIPDeterminationParameters struct {

	// (String) Client IP header name
	// Client IP header name
	// +kubebuilder:validation:Optional
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// +kubebuilder:validation:Optional
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`
}

// RumIPDeterminationSpec defines the desired state of RumIPDetermination
type RumIPDeterminationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RumIPDeterminationParameters `json:"forProvider"`
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
	InitProvider RumIPDeterminationInitParameters `json:"initProvider,omitempty"`
}

// RumIPDeterminationStatus defines the observed state of RumIPDetermination.
type RumIPDeterminationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RumIPDeterminationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RumIPDetermination is the Schema for the RumIPDeterminations API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type RumIPDetermination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.headerName) || (has(self.initProvider) && has(self.initProvider.headerName))",message="spec.forProvider.headerName is a required parameter"
	Spec   RumIPDeterminationSpec   `json:"spec"`
	Status RumIPDeterminationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RumIPDeterminationList contains a list of RumIPDeterminations
type RumIPDeterminationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RumIPDetermination `json:"items"`
}

// Repository type metadata.
var (
	RumIPDetermination_Kind             = "RumIPDetermination"
	RumIPDetermination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RumIPDetermination_Kind}.String()
	RumIPDetermination_KindAPIVersion   = RumIPDetermination_Kind + "." + CRDGroupVersion.String()
	RumIPDetermination_GroupVersionKind = CRDGroupVersion.WithKind(RumIPDetermination_Kind)
)

func init() {
	SchemeBuilder.Register(&RumIPDetermination{}, &RumIPDeterminationList{})
}
