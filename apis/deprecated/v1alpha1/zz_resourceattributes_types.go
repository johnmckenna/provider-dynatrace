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

type KeysInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type KeysObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type KeysParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule" tf:"rule,omitempty"`
}

type ResourceAttributesInitParameters struct {

	// list (see below for nested schema)
	// Attribute key allow-list
	Keys []KeysInitParameters `json:"keys,omitempty" tf:"keys,omitempty"`
}

type ResourceAttributesObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// list (see below for nested schema)
	// Attribute key allow-list
	Keys []KeysObservation `json:"keys,omitempty" tf:"keys,omitempty"`
}

type ResourceAttributesParameters struct {

	// list (see below for nested schema)
	// Attribute key allow-list
	// +kubebuilder:validation:Optional
	Keys []KeysParameters `json:"keys,omitempty" tf:"keys,omitempty"`
}

type RuleInitParameters struct {

	// (String) Attribute key service.name is automatically captured by default
	// Attribute key **service.name** is automatically captured by default
	AttributeKey *string `json:"attributeKey,omitempty" tf:"attribute_key,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Possible Values: MASK_ENTIRE_VALUE, MASK_ONLY_CONFIDENTIAL_DATA, NOT_MASKED
	// Possible Values: `MASK_ENTIRE_VALUE`, `MASK_ONLY_CONFIDENTIAL_DATA`, `NOT_MASKED`
	Masking *string `json:"masking,omitempty" tf:"masking,omitempty"`
}

type RuleObservation struct {

	// (String) Attribute key service.name is automatically captured by default
	// Attribute key **service.name** is automatically captured by default
	AttributeKey *string `json:"attributeKey,omitempty" tf:"attribute_key,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Possible Values: MASK_ENTIRE_VALUE, MASK_ONLY_CONFIDENTIAL_DATA, NOT_MASKED
	// Possible Values: `MASK_ENTIRE_VALUE`, `MASK_ONLY_CONFIDENTIAL_DATA`, `NOT_MASKED`
	Masking *string `json:"masking,omitempty" tf:"masking,omitempty"`
}

type RuleParameters struct {

	// (String) Attribute key service.name is automatically captured by default
	// Attribute key **service.name** is automatically captured by default
	// +kubebuilder:validation:Optional
	AttributeKey *string `json:"attributeKey" tf:"attribute_key,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// (String) Possible Values: MASK_ENTIRE_VALUE, MASK_ONLY_CONFIDENTIAL_DATA, NOT_MASKED
	// Possible Values: `MASK_ENTIRE_VALUE`, `MASK_ONLY_CONFIDENTIAL_DATA`, `NOT_MASKED`
	// +kubebuilder:validation:Optional
	Masking *string `json:"masking" tf:"masking,omitempty"`
}

// ResourceAttributesSpec defines the desired state of ResourceAttributes
type ResourceAttributesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceAttributesParameters `json:"forProvider"`
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
	InitProvider ResourceAttributesInitParameters `json:"initProvider,omitempty"`
}

// ResourceAttributesStatus defines the observed state of ResourceAttributes.
type ResourceAttributesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceAttributesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ResourceAttributes is the Schema for the ResourceAttributess API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type ResourceAttributes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceAttributesSpec   `json:"spec"`
	Status            ResourceAttributesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceAttributesList contains a list of ResourceAttributess
type ResourceAttributesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceAttributes `json:"items"`
}

// Repository type metadata.
var (
	ResourceAttributes_Kind             = "ResourceAttributes"
	ResourceAttributes_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceAttributes_Kind}.String()
	ResourceAttributes_KindAPIVersion   = ResourceAttributes_Kind + "." + CRDGroupVersion.String()
	ResourceAttributes_GroupVersionKind = CRDGroupVersion.WithKind(ResourceAttributes_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceAttributes{}, &ResourceAttributesList{})
}
