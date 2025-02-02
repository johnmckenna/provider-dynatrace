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

type AutomationWorkflowSlackInitParameters struct {

	// (String, Deprecated) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) The name of the Slack connection
	// The name of the Slack connection
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive) The bot token obtained from the Slack App Management UI
	// The bot token obtained from the Slack App Management UI
	TokenSecretRef v1.SecretKeySelector `json:"tokenSecretRef" tf:"-"`
}

type AutomationWorkflowSlackObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String, Deprecated) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) The name of the Slack connection
	// The name of the Slack connection
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AutomationWorkflowSlackParameters struct {

	// (String, Deprecated) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// +kubebuilder:validation:Optional
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) The name of the Slack connection
	// The name of the Slack connection
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive) The bot token obtained from the Slack App Management UI
	// The bot token obtained from the Slack App Management UI
	// +kubebuilder:validation:Optional
	TokenSecretRef v1.SecretKeySelector `json:"tokenSecretRef" tf:"-"`
}

// AutomationWorkflowSlackSpec defines the desired state of AutomationWorkflowSlack
type AutomationWorkflowSlackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutomationWorkflowSlackParameters `json:"forProvider"`
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
	InitProvider AutomationWorkflowSlackInitParameters `json:"initProvider,omitempty"`
}

// AutomationWorkflowSlackStatus defines the observed state of AutomationWorkflowSlack.
type AutomationWorkflowSlackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutomationWorkflowSlackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AutomationWorkflowSlack is the Schema for the AutomationWorkflowSlacks API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type AutomationWorkflowSlack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tokenSecretRef)",message="spec.forProvider.tokenSecretRef is a required parameter"
	Spec   AutomationWorkflowSlackSpec   `json:"spec"`
	Status AutomationWorkflowSlackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutomationWorkflowSlackList contains a list of AutomationWorkflowSlacks
type AutomationWorkflowSlackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutomationWorkflowSlack `json:"items"`
}

// Repository type metadata.
var (
	AutomationWorkflowSlack_Kind             = "AutomationWorkflowSlack"
	AutomationWorkflowSlack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AutomationWorkflowSlack_Kind}.String()
	AutomationWorkflowSlack_KindAPIVersion   = AutomationWorkflowSlack_Kind + "." + CRDGroupVersion.String()
	AutomationWorkflowSlack_GroupVersionKind = CRDGroupVersion.WithKind(AutomationWorkflowSlack_Kind)
)

func init() {
	SchemeBuilder.Register(&AutomationWorkflowSlack{}, &AutomationWorkflowSlackList{})
}
