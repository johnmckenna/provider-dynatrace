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

type AutomationWorkflowAwsConnectionsInitParameters struct {

	// (String) Name
	// Name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Possible Values: WebIdentity
	// Possible Values: `WebIdentity`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	WebIdentity []WebIdentityInitParameters `json:"webIdentity,omitempty" tf:"web_identity,omitempty"`
}

type AutomationWorkflowAwsConnectionsObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Name
	// Name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Possible Values: WebIdentity
	// Possible Values: `WebIdentity`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	WebIdentity []WebIdentityParameters `json:"webIdentity,omitempty" tf:"web_identity,omitempty"`
}

type AutomationWorkflowAwsConnectionsParameters struct {

	// (String) Name
	// Name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Possible Values: WebIdentity
	// Possible Values: `WebIdentity`
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	// +kubebuilder:validation:Optional
	WebIdentity []WebIdentityParameters `json:"webIdentity,omitempty" tf:"web_identity,omitempty"`
}

type WebIdentityInitParameters struct {
	PolicyArns []*string `json:"policyArnsSecretRef,omitempty" tf:"-"`

	// (String, Sensitive) The ARN of the AWS role that should be assumed
	// The ARN of the AWS role that should be assumed
	RoleArnSecretRef v1.SecretKeySelector `json:"roleArnSecretRef" tf:"-"`
}

type WebIdentityObservation struct {
}

type WebIdentityParameters struct {

	// (List of String, Sensitive) An optional list of policies that can be used to restrict the AWS role
	// An optional list of policies that can be used to restrict the AWS role
	// +kubebuilder:validation:Optional
	PolicyArnsSecretRef *[]v1.SecretKeySelector `json:"policyArnsSecretRef,omitempty" tf:"-"`

	// (String, Sensitive) The ARN of the AWS role that should be assumed
	// The ARN of the AWS role that should be assumed
	// +kubebuilder:validation:Optional
	RoleArnSecretRef v1.SecretKeySelector `json:"roleArnSecretRef" tf:"-"`
}

// AutomationWorkflowAwsConnectionsSpec defines the desired state of AutomationWorkflowAwsConnections
type AutomationWorkflowAwsConnectionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutomationWorkflowAwsConnectionsParameters `json:"forProvider"`
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
	InitProvider AutomationWorkflowAwsConnectionsInitParameters `json:"initProvider,omitempty"`
}

// AutomationWorkflowAwsConnectionsStatus defines the observed state of AutomationWorkflowAwsConnections.
type AutomationWorkflowAwsConnectionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutomationWorkflowAwsConnectionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AutomationWorkflowAwsConnections is the Schema for the AutomationWorkflowAwsConnectionss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type AutomationWorkflowAwsConnections struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   AutomationWorkflowAwsConnectionsSpec   `json:"spec"`
	Status AutomationWorkflowAwsConnectionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutomationWorkflowAwsConnectionsList contains a list of AutomationWorkflowAwsConnectionss
type AutomationWorkflowAwsConnectionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutomationWorkflowAwsConnections `json:"items"`
}

// Repository type metadata.
var (
	AutomationWorkflowAwsConnections_Kind             = "AutomationWorkflowAwsConnections"
	AutomationWorkflowAwsConnections_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AutomationWorkflowAwsConnections_Kind}.String()
	AutomationWorkflowAwsConnections_KindAPIVersion   = AutomationWorkflowAwsConnections_Kind + "." + CRDGroupVersion.String()
	AutomationWorkflowAwsConnections_GroupVersionKind = CRDGroupVersion.WithKind(AutomationWorkflowAwsConnections_Kind)
)

func init() {
	SchemeBuilder.Register(&AutomationWorkflowAwsConnections{}, &AutomationWorkflowAwsConnectionsList{})
}
