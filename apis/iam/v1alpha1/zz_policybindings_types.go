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

type PolicyBindingsInitParameters struct {

	// The UUID of the cluster. The attribute `policies` must contain ONLY policies defined for that cluster.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The ID of the environment (https://<environmentid>.live.dynatrace.com). The attribute `policies` must contain ONLY policies defined for that environment.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The name of the policy
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// A list of IDs referring to policies bound to that group. It's not possible to mix policies here that are defined for different scopes (different clusters or environments) than specified via attributes `cluster` or `environment`.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type PolicyBindingsObservation struct {

	// The UUID of the cluster. The attribute `policies` must contain ONLY policies defined for that cluster.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The ID of the environment (https://<environmentid>.live.dynatrace.com). The attribute `policies` must contain ONLY policies defined for that environment.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The name of the policy
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of IDs referring to policies bound to that group. It's not possible to mix policies here that are defined for different scopes (different clusters or environments) than specified via attributes `cluster` or `environment`.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type PolicyBindingsParameters struct {

	// The UUID of the cluster. The attribute `policies` must contain ONLY policies defined for that cluster.
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The ID of the environment (https://<environmentid>.live.dynatrace.com). The attribute `policies` must contain ONLY policies defined for that environment.
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The name of the policy
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// A list of IDs referring to policies bound to that group. It's not possible to mix policies here that are defined for different scopes (different clusters or environments) than specified via attributes `cluster` or `environment`.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

// PolicyBindingsSpec defines the desired state of PolicyBindings
type PolicyBindingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyBindingsParameters `json:"forProvider"`
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
	InitProvider PolicyBindingsInitParameters `json:"initProvider,omitempty"`
}

// PolicyBindingsStatus defines the observed state of PolicyBindings.
type PolicyBindingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyBindingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PolicyBindings is the Schema for the PolicyBindingss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type PolicyBindings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.group) || (has(self.initProvider) && has(self.initProvider.group))",message="spec.forProvider.group is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policies) || (has(self.initProvider) && has(self.initProvider.policies))",message="spec.forProvider.policies is a required parameter"
	Spec   PolicyBindingsSpec   `json:"spec"`
	Status PolicyBindingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyBindingsList contains a list of PolicyBindingss
type PolicyBindingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyBindings `json:"items"`
}

// Repository type metadata.
var (
	PolicyBindings_Kind             = "PolicyBindings"
	PolicyBindings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyBindings_Kind}.String()
	PolicyBindings_KindAPIVersion   = PolicyBindings_Kind + "." + CRDGroupVersion.String()
	PolicyBindings_GroupVersionKind = CRDGroupVersion.WithKind(PolicyBindings_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyBindings{}, &PolicyBindingsList{})
}