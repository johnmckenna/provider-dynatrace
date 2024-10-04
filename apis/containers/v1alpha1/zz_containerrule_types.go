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

type ContainerRuleInitParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) Possible Values: MONITORING_OFF, MONITORING_ON
	// Possible Values: `MONITORING_OFF`, `MONITORING_ON`
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// (String) Possible Values: CONTAINS, ENDS, EQUALS, EXISTS, NOT_CONTAINS, NOT_ENDS, NOT_EQUALS, NOT_EXISTS, NOT_STARTS, STARTS
	// Possible Values: `CONTAINS`, `ENDS`, `EQUALS`, `EXISTS`, `NOT_CONTAINS`, `NOT_ENDS`, `NOT_EQUALS`, `NOT_EXISTS`, `NOT_STARTS`, `STARTS`
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// (String) Possible Values: CONTAINER_NAME, IMAGE_NAME, KUBERNETES_BASEPODNAME, KUBERNETES_CONTAINERNAME, KUBERNETES_FULLPODNAME, KUBERNETES_NAMESPACE, KUBERNETES_PODUID
	// Possible Values: `CONTAINER_NAME`, `IMAGE_NAME`, `KUBERNETES_BASEPODNAME`, `KUBERNETES_CONTAINERNAME`, `KUBERNETES_FULLPODNAME`, `KUBERNETES_NAMESPACE`, `KUBERNETES_PODUID`
	Property *string `json:"property,omitempty" tf:"property,omitempty"`

	// (String) Condition value
	// Condition value
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ContainerRuleObservation struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) Possible Values: MONITORING_OFF, MONITORING_ON
	// Possible Values: `MONITORING_OFF`, `MONITORING_ON`
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// (String) Possible Values: CONTAINS, ENDS, EQUALS, EXISTS, NOT_CONTAINS, NOT_ENDS, NOT_EQUALS, NOT_EXISTS, NOT_STARTS, STARTS
	// Possible Values: `CONTAINS`, `ENDS`, `EQUALS`, `EXISTS`, `NOT_CONTAINS`, `NOT_ENDS`, `NOT_EQUALS`, `NOT_EXISTS`, `NOT_STARTS`, `STARTS`
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// (String) Possible Values: CONTAINER_NAME, IMAGE_NAME, KUBERNETES_BASEPODNAME, KUBERNETES_CONTAINERNAME, KUBERNETES_FULLPODNAME, KUBERNETES_NAMESPACE, KUBERNETES_PODUID
	// Possible Values: `CONTAINER_NAME`, `IMAGE_NAME`, `KUBERNETES_BASEPODNAME`, `KUBERNETES_CONTAINERNAME`, `KUBERNETES_FULLPODNAME`, `KUBERNETES_NAMESPACE`, `KUBERNETES_PODUID`
	Property *string `json:"property,omitempty" tf:"property,omitempty"`

	// (String) Condition value
	// Condition value
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ContainerRuleParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// +kubebuilder:validation:Optional
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) Possible Values: MONITORING_OFF, MONITORING_ON
	// Possible Values: `MONITORING_OFF`, `MONITORING_ON`
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// (String) Possible Values: CONTAINS, ENDS, EQUALS, EXISTS, NOT_CONTAINS, NOT_ENDS, NOT_EQUALS, NOT_EXISTS, NOT_STARTS, STARTS
	// Possible Values: `CONTAINS`, `ENDS`, `EQUALS`, `EXISTS`, `NOT_CONTAINS`, `NOT_ENDS`, `NOT_EQUALS`, `NOT_EXISTS`, `NOT_STARTS`, `STARTS`
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// (String) Possible Values: CONTAINER_NAME, IMAGE_NAME, KUBERNETES_BASEPODNAME, KUBERNETES_CONTAINERNAME, KUBERNETES_FULLPODNAME, KUBERNETES_NAMESPACE, KUBERNETES_PODUID
	// Possible Values: `CONTAINER_NAME`, `IMAGE_NAME`, `KUBERNETES_BASEPODNAME`, `KUBERNETES_CONTAINERNAME`, `KUBERNETES_FULLPODNAME`, `KUBERNETES_NAMESPACE`, `KUBERNETES_PODUID`
	// +kubebuilder:validation:Optional
	Property *string `json:"property,omitempty" tf:"property,omitempty"`

	// (String) Condition value
	// Condition value
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// ContainerRuleSpec defines the desired state of ContainerRule
type ContainerRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerRuleParameters `json:"forProvider"`
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
	InitProvider ContainerRuleInitParameters `json:"initProvider,omitempty"`
}

// ContainerRuleStatus defines the observed state of ContainerRule.
type ContainerRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ContainerRule is the Schema for the ContainerRules API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type ContainerRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mode) || (has(self.initProvider) && has(self.initProvider.mode))",message="spec.forProvider.mode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.operator) || (has(self.initProvider) && has(self.initProvider.operator))",message="spec.forProvider.operator is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.property) || (has(self.initProvider) && has(self.initProvider.property))",message="spec.forProvider.property is a required parameter"
	Spec   ContainerRuleSpec   `json:"spec"`
	Status ContainerRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerRuleList contains a list of ContainerRules
type ContainerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerRule `json:"items"`
}

// Repository type metadata.
var (
	ContainerRule_Kind             = "ContainerRule"
	ContainerRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContainerRule_Kind}.String()
	ContainerRule_KindAPIVersion   = ContainerRule_Kind + "." + CRDGroupVersion.String()
	ContainerRule_GroupVersionKind = CRDGroupVersion.WithKind(ContainerRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ContainerRule{}, &ContainerRuleList{})
}