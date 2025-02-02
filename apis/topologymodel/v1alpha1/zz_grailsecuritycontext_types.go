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

type GrailSecurityContextInitParameters struct {

	// sensitive name of a property of the destination type.
	// The case-sensitive name of a property of the destination type.
	DestinationProperty *string `json:"destinationProperty,omitempty" tf:"destination_property,omitempty"`

	// (String) Type of the entity whose security context to override.
	// Type of the entity whose security context to override.
	EntityType *string `json:"entityType,omitempty" tf:"entity_type,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`
}

type GrailSecurityContextObservation struct {

	// sensitive name of a property of the destination type.
	// The case-sensitive name of a property of the destination type.
	DestinationProperty *string `json:"destinationProperty,omitempty" tf:"destination_property,omitempty"`

	// (String) Type of the entity whose security context to override.
	// Type of the entity whose security context to override.
	EntityType *string `json:"entityType,omitempty" tf:"entity_type,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`
}

type GrailSecurityContextParameters struct {

	// sensitive name of a property of the destination type.
	// The case-sensitive name of a property of the destination type.
	// +kubebuilder:validation:Optional
	DestinationProperty *string `json:"destinationProperty,omitempty" tf:"destination_property,omitempty"`

	// (String) Type of the entity whose security context to override.
	// Type of the entity whose security context to override.
	// +kubebuilder:validation:Optional
	EntityType *string `json:"entityType,omitempty" tf:"entity_type,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// +kubebuilder:validation:Optional
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`
}

// GrailSecurityContextSpec defines the desired state of GrailSecurityContext
type GrailSecurityContextSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GrailSecurityContextParameters `json:"forProvider"`
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
	InitProvider GrailSecurityContextInitParameters `json:"initProvider,omitempty"`
}

// GrailSecurityContextStatus defines the observed state of GrailSecurityContext.
type GrailSecurityContextStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GrailSecurityContextObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GrailSecurityContext is the Schema for the GrailSecurityContexts API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type GrailSecurityContext struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destinationProperty) || (has(self.initProvider) && has(self.initProvider.destinationProperty))",message="spec.forProvider.destinationProperty is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.entityType) || (has(self.initProvider) && has(self.initProvider.entityType))",message="spec.forProvider.entityType is a required parameter"
	Spec   GrailSecurityContextSpec   `json:"spec"`
	Status GrailSecurityContextStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GrailSecurityContextList contains a list of GrailSecurityContexts
type GrailSecurityContextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GrailSecurityContext `json:"items"`
}

// Repository type metadata.
var (
	GrailSecurityContext_Kind             = "GrailSecurityContext"
	GrailSecurityContext_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GrailSecurityContext_Kind}.String()
	GrailSecurityContext_KindAPIVersion   = GrailSecurityContext_Kind + "." + CRDGroupVersion.String()
	GrailSecurityContext_GroupVersionKind = CRDGroupVersion.WithKind(GrailSecurityContext_Kind)
)

func init() {
	SchemeBuilder.Register(&GrailSecurityContext{}, &GrailSecurityContextList{})
}
