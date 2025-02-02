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

type PolicyInitParameters_2 struct {

	// (String) The UUID of the cluster in case the policy should be applied to all environments of this cluster.
	// The UUID of the cluster in case the policy should be applied to all environments of this cluster.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// (String) An optional description text for the policy
	// An optional description text for the policy
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of the environment if the policy should be applied to a specific environment
	// The ID of the environment if the policy should be applied to a specific environment
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// (String) The name of the policy
	// The name of the policy
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Statement Query of the policy
	// The Statement Query of the policy
	StatementQuery *string `json:"statementQuery,omitempty" tf:"statement_query,omitempty"`

	// (Set of String) Tags for this policy
	// Tags for this policy
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PolicyObservation_2 struct {

	// (String) The UUID of the cluster in case the policy should be applied to all environments of this cluster.
	// The UUID of the cluster in case the policy should be applied to all environments of this cluster.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// (String) An optional description text for the policy
	// An optional description text for the policy
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of the environment if the policy should be applied to a specific environment
	// The ID of the environment if the policy should be applied to a specific environment
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the policy
	// The name of the policy
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Statement Query of the policy
	// The Statement Query of the policy
	StatementQuery *string `json:"statementQuery,omitempty" tf:"statement_query,omitempty"`

	// (Set of String) Tags for this policy
	// Tags for this policy
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PolicyParameters_2 struct {

	// (String) The UUID of the cluster in case the policy should be applied to all environments of this cluster.
	// The UUID of the cluster in case the policy should be applied to all environments of this cluster.
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// (String) An optional description text for the policy
	// An optional description text for the policy
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of the environment if the policy should be applied to a specific environment
	// The ID of the environment if the policy should be applied to a specific environment
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// (String) The name of the policy
	// The name of the policy
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Statement Query of the policy
	// The Statement Query of the policy
	// +kubebuilder:validation:Optional
	StatementQuery *string `json:"statementQuery,omitempty" tf:"statement_query,omitempty"`

	// (Set of String) Tags for this policy
	// Tags for this policy
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters_2 `json:"forProvider"`
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
	InitProvider PolicyInitParameters_2 `json:"initProvider,omitempty"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Policy is the Schema for the Policys API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.statementQuery) || (has(self.initProvider) && has(self.initProvider.statementQuery))",message="spec.forProvider.statementQuery is a required parameter"
	Spec   PolicySpec   `json:"spec"`
	Status PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
