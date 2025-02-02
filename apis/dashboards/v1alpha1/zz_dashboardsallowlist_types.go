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

type AllowlistInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Urlpattern []UrlpatternInitParameters `json:"urlpattern,omitempty" tf:"urlpattern,omitempty"`
}

type AllowlistObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Urlpattern []UrlpatternObservation `json:"urlpattern,omitempty" tf:"urlpattern,omitempty"`
}

type AllowlistParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Urlpattern []UrlpatternParameters `json:"urlpattern" tf:"urlpattern,omitempty"`
}

type DashboardsAllowlistInitParameters struct {

	// (Block List, Max: 1) List of URL pattern matchers (see below for nested schema)
	// List of URL pattern matchers
	Allowlist []AllowlistInitParameters `json:"allowlist,omitempty" tf:"allowlist,omitempty"`
}

type DashboardsAllowlistObservation struct {

	// (Block List, Max: 1) List of URL pattern matchers (see below for nested schema)
	// List of URL pattern matchers
	Allowlist []AllowlistObservation `json:"allowlist,omitempty" tf:"allowlist,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DashboardsAllowlistParameters struct {

	// (Block List, Max: 1) List of URL pattern matchers (see below for nested schema)
	// List of URL pattern matchers
	// +kubebuilder:validation:Optional
	Allowlist []AllowlistParameters `json:"allowlist,omitempty" tf:"allowlist,omitempty"`
}

type UrlpatternInitParameters struct {

	// (String) Possible Values: Equals, StartsWith
	// Possible Values: `Equals`, `StartsWith`
	Rule *string `json:"rule,omitempty" tf:"rule,omitempty"`

	// (String) Pattern
	// Pattern
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type UrlpatternObservation struct {

	// (String) Possible Values: Equals, StartsWith
	// Possible Values: `Equals`, `StartsWith`
	Rule *string `json:"rule,omitempty" tf:"rule,omitempty"`

	// (String) Pattern
	// Pattern
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type UrlpatternParameters struct {

	// (String) Possible Values: Equals, StartsWith
	// Possible Values: `Equals`, `StartsWith`
	// +kubebuilder:validation:Optional
	Rule *string `json:"rule" tf:"rule,omitempty"`

	// (String) Pattern
	// Pattern
	// +kubebuilder:validation:Optional
	Template *string `json:"template" tf:"template,omitempty"`
}

// DashboardsAllowlistSpec defines the desired state of DashboardsAllowlist
type DashboardsAllowlistSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DashboardsAllowlistParameters `json:"forProvider"`
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
	InitProvider DashboardsAllowlistInitParameters `json:"initProvider,omitempty"`
}

// DashboardsAllowlistStatus defines the observed state of DashboardsAllowlist.
type DashboardsAllowlistStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DashboardsAllowlistObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DashboardsAllowlist is the Schema for the DashboardsAllowlists API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type DashboardsAllowlist struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DashboardsAllowlistSpec   `json:"spec"`
	Status            DashboardsAllowlistStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DashboardsAllowlistList contains a list of DashboardsAllowlists
type DashboardsAllowlistList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DashboardsAllowlist `json:"items"`
}

// Repository type metadata.
var (
	DashboardsAllowlist_Kind             = "DashboardsAllowlist"
	DashboardsAllowlist_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DashboardsAllowlist_Kind}.String()
	DashboardsAllowlist_KindAPIVersion   = DashboardsAllowlist_Kind + "." + CRDGroupVersion.String()
	DashboardsAllowlist_GroupVersionKind = CRDGroupVersion.WithKind(DashboardsAllowlist_Kind)
)

func init() {
	SchemeBuilder.Register(&DashboardsAllowlist{}, &DashboardsAllowlistList{})
}
