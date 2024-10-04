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

type TransactionStartFiltersInitParameters struct {

	// (Set of String) You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// +listType=set
	CicsTerminalTransactionIds []*string `json:"cicsTerminalTransactionIds,omitempty" tf:"cics_terminal_transaction_ids,omitempty"`

	// (Set of String) You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// +listType=set
	CicsTransactionIds []*string `json:"cicsTransactionIds,omitempty" tf:"cics_transaction_ids,omitempty"`

	// (Set of String) You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// +listType=set
	ImsTerminalTransactionIds []*string `json:"imsTerminalTransactionIds,omitempty" tf:"ims_terminal_transaction_ids,omitempty"`

	// (Set of String) You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// +listType=set
	ImsTransactionIds []*string `json:"imsTransactionIds,omitempty" tf:"ims_transaction_ids,omitempty"`
}

type TransactionStartFiltersObservation struct {

	// (Set of String) You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// +listType=set
	CicsTerminalTransactionIds []*string `json:"cicsTerminalTransactionIds,omitempty" tf:"cics_terminal_transaction_ids,omitempty"`

	// (Set of String) You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// +listType=set
	CicsTransactionIds []*string `json:"cicsTransactionIds,omitempty" tf:"cics_transaction_ids,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of String) You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// +listType=set
	ImsTerminalTransactionIds []*string `json:"imsTerminalTransactionIds,omitempty" tf:"ims_terminal_transaction_ids,omitempty"`

	// (Set of String) You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// +listType=set
	ImsTransactionIds []*string `json:"imsTransactionIds,omitempty" tf:"ims_transaction_ids,omitempty"`
}

type TransactionStartFiltersParameters struct {

	// (Set of String) You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// +kubebuilder:validation:Optional
	// +listType=set
	CicsTerminalTransactionIds []*string `json:"cicsTerminalTransactionIds,omitempty" tf:"cics_terminal_transaction_ids,omitempty"`

	// (Set of String) You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// +kubebuilder:validation:Optional
	// +listType=set
	CicsTransactionIds []*string `json:"cicsTransactionIds,omitempty" tf:"cics_transaction_ids,omitempty"`

	// (Set of String) You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// +kubebuilder:validation:Optional
	// +listType=set
	ImsTerminalTransactionIds []*string `json:"imsTerminalTransactionIds,omitempty" tf:"ims_terminal_transaction_ids,omitempty"`

	// (Set of String) You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	// +kubebuilder:validation:Optional
	// +listType=set
	ImsTransactionIds []*string `json:"imsTransactionIds,omitempty" tf:"ims_transaction_ids,omitempty"`
}

// TransactionStartFiltersSpec defines the desired state of TransactionStartFilters
type TransactionStartFiltersSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransactionStartFiltersParameters `json:"forProvider"`
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
	InitProvider TransactionStartFiltersInitParameters `json:"initProvider,omitempty"`
}

// TransactionStartFiltersStatus defines the observed state of TransactionStartFilters.
type TransactionStartFiltersStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransactionStartFiltersObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TransactionStartFilters is the Schema for the TransactionStartFilterss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type TransactionStartFilters struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransactionStartFiltersSpec   `json:"spec"`
	Status            TransactionStartFiltersStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransactionStartFiltersList contains a list of TransactionStartFilterss
type TransactionStartFiltersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransactionStartFilters `json:"items"`
}

// Repository type metadata.
var (
	TransactionStartFilters_Kind             = "TransactionStartFilters"
	TransactionStartFilters_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransactionStartFilters_Kind}.String()
	TransactionStartFilters_KindAPIVersion   = TransactionStartFilters_Kind + "." + CRDGroupVersion.String()
	TransactionStartFilters_GroupVersionKind = CRDGroupVersion.WithKind(TransactionStartFilters_Kind)
)

func init() {
	SchemeBuilder.Register(&TransactionStartFilters{}, &TransactionStartFiltersList{})
}
