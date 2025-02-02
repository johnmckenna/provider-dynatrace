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

type HubSubscriptionsInitParameters struct {

	// (Block List, Max: 1) Subscriptions (see below for nested schema)
	// Subscriptions
	TokenSubscriptions []TokenSubscriptionsInitParameters `json:"tokenSubscriptions,omitempty" tf:"token_subscriptions,omitempty"`
}

type HubSubscriptionsObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Max: 1) Subscriptions (see below for nested schema)
	// Subscriptions
	TokenSubscriptions []TokenSubscriptionsObservation `json:"tokenSubscriptions,omitempty" tf:"token_subscriptions,omitempty"`
}

type HubSubscriptionsParameters struct {

	// (Block List, Max: 1) Subscriptions (see below for nested schema)
	// Subscriptions
	// +kubebuilder:validation:Optional
	TokenSubscriptions []TokenSubscriptionsParameters `json:"tokenSubscriptions,omitempty" tf:"token_subscriptions,omitempty"`
}

type TokenSubscriptionInitParameters struct {

	// (String) no documentation available
	// no documentation available
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Name of subscription
	// Name of subscription
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Subscription token
	// Subscription token
	Token *string `json:"token,omitempty" tf:"token,omitempty"`
}

type TokenSubscriptionObservation struct {

	// (String) no documentation available
	// no documentation available
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Name of subscription
	// Name of subscription
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Subscription token
	// Subscription token
	Token *string `json:"token,omitempty" tf:"token,omitempty"`
}

type TokenSubscriptionParameters struct {

	// (String) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// (String) Name of subscription
	// Name of subscription
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) Subscription token
	// Subscription token
	// +kubebuilder:validation:Optional
	Token *string `json:"token" tf:"token,omitempty"`
}

type TokenSubscriptionsInitParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	TokenSubscription []TokenSubscriptionInitParameters `json:"tokenSubscription,omitempty" tf:"token_subscription,omitempty"`
}

type TokenSubscriptionsObservation struct {

	// (Block List, Min: 1) (see below for nested schema)
	TokenSubscription []TokenSubscriptionObservation `json:"tokenSubscription,omitempty" tf:"token_subscription,omitempty"`
}

type TokenSubscriptionsParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	TokenSubscription []TokenSubscriptionParameters `json:"tokenSubscription" tf:"token_subscription,omitempty"`
}

// HubSubscriptionsSpec defines the desired state of HubSubscriptions
type HubSubscriptionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HubSubscriptionsParameters `json:"forProvider"`
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
	InitProvider HubSubscriptionsInitParameters `json:"initProvider,omitempty"`
}

// HubSubscriptionsStatus defines the observed state of HubSubscriptions.
type HubSubscriptionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HubSubscriptionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HubSubscriptions is the Schema for the HubSubscriptionss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type HubSubscriptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HubSubscriptionsSpec   `json:"spec"`
	Status            HubSubscriptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HubSubscriptionsList contains a list of HubSubscriptionss
type HubSubscriptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HubSubscriptions `json:"items"`
}

// Repository type metadata.
var (
	HubSubscriptions_Kind             = "HubSubscriptions"
	HubSubscriptions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HubSubscriptions_Kind}.String()
	HubSubscriptions_KindAPIVersion   = HubSubscriptions_Kind + "." + CRDGroupVersion.String()
	HubSubscriptions_GroupVersionKind = CRDGroupVersion.WithKind(HubSubscriptions_Kind)
)

func init() {
	SchemeBuilder.Register(&HubSubscriptions{}, &HubSubscriptionsList{})
}
