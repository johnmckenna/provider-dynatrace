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

type DocumentInitParameters struct {

	// The user context the executions of the document will happen with
	Actor *string `json:"actor,omitempty" tf:"actor,omitempty"`

	// Document content as JSON
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The name/name of the document
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the owner of this document
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Specifies whether the document is private or readable by everybody
	Private *bool `json:"private,omitempty" tf:"private,omitempty"`

	// Type of the document. Possible Values are `dashboard` and `notebook`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DocumentObservation struct {

	// The user context the executions of the document will happen with
	Actor *string `json:"actor,omitempty" tf:"actor,omitempty"`

	// Document content as JSON
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name/name of the document
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the owner of this document
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Specifies whether the document is private or readable by everybody
	Private *bool `json:"private,omitempty" tf:"private,omitempty"`

	// Type of the document. Possible Values are `dashboard` and `notebook`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The version of the document
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type DocumentParameters struct {

	// The user context the executions of the document will happen with
	// +kubebuilder:validation:Optional
	Actor *string `json:"actor,omitempty" tf:"actor,omitempty"`

	// Document content as JSON
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The name/name of the document
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the owner of this document
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Specifies whether the document is private or readable by everybody
	// +kubebuilder:validation:Optional
	Private *bool `json:"private,omitempty" tf:"private,omitempty"`

	// Type of the document. Possible Values are `dashboard` and `notebook`
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DocumentSpec defines the desired state of Document
type DocumentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DocumentParameters `json:"forProvider"`
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
	InitProvider DocumentInitParameters `json:"initProvider,omitempty"`
}

// DocumentStatus defines the observed state of Document.
type DocumentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DocumentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Document is the Schema for the Documents API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type Document struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.content) || (has(self.initProvider) && has(self.initProvider.content))",message="spec.forProvider.content is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   DocumentSpec   `json:"spec"`
	Status DocumentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentList contains a list of Documents
type DocumentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Document `json:"items"`
}

// Repository type metadata.
var (
	Document_Kind             = "Document"
	Document_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Document_Kind}.String()
	Document_KindAPIVersion   = Document_Kind + "." + CRDGroupVersion.String()
	Document_GroupVersionKind = CRDGroupVersion.WithKind(Document_Kind)
)

func init() {
	SchemeBuilder.Register(&Document{}, &DocumentList{})
}
