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

type DevobsAgentOptinInitParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The scope of this setting (PROCESS_GROUP, CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	// The scope of this setting (PROCESS_GROUP, CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type DevobsAgentOptinObservation struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The scope of this setting (PROCESS_GROUP, CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	// The scope of this setting (PROCESS_GROUP, CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type DevobsAgentOptinParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The scope of this setting (PROCESS_GROUP, CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	// The scope of this setting (PROCESS_GROUP, CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// DevobsAgentOptinSpec defines the desired state of DevobsAgentOptin
type DevobsAgentOptinSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DevobsAgentOptinParameters `json:"forProvider"`
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
	InitProvider DevobsAgentOptinInitParameters `json:"initProvider,omitempty"`
}

// DevobsAgentOptinStatus defines the observed state of DevobsAgentOptin.
type DevobsAgentOptinStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DevobsAgentOptinObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DevobsAgentOptin is the Schema for the DevobsAgentOptins API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type DevobsAgentOptin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	Spec   DevobsAgentOptinSpec   `json:"spec"`
	Status DevobsAgentOptinStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DevobsAgentOptinList contains a list of DevobsAgentOptins
type DevobsAgentOptinList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevobsAgentOptin `json:"items"`
}

// Repository type metadata.
var (
	DevobsAgentOptin_Kind             = "DevobsAgentOptin"
	DevobsAgentOptin_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DevobsAgentOptin_Kind}.String()
	DevobsAgentOptin_KindAPIVersion   = DevobsAgentOptin_Kind + "." + CRDGroupVersion.String()
	DevobsAgentOptin_GroupVersionKind = CRDGroupVersion.WithKind(DevobsAgentOptin_Kind)
)

func init() {
	SchemeBuilder.Register(&DevobsAgentOptin{}, &DevobsAgentOptinList{})
}
