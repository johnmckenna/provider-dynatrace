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

type VmwareInitParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) This string should have one of the following formats:
	// This string should have one of the following formats:
	// - $prefix(parameter) - property value starting with 'parameter'
	// - $eq(parameter) - property value exactly matching 'parameter'
	// - $suffix(parameter) - property value ends with 'parameter'
	// - $contains(parameter) - property value contains 'parameter'
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) Specify the IP address or name of the vCenter or standalone ESXi host:
	// Specify the IP address or name of the vCenter or standalone ESXi host:
	Ipaddress *string `json:"ipaddress,omitempty" tf:"ipaddress,omitempty"`

	// (String) Name this connection
	// Name this connection
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// (String, Sensitive) no documentation available
	// no documentation available
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// (String) Provide user credentials for the vCenter or standalone ESXi host:
	// Provide user credentials for the vCenter or standalone ESXi host:
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type VmwareObservation struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) This string should have one of the following formats:
	// This string should have one of the following formats:
	// - $prefix(parameter) - property value starting with 'parameter'
	// - $eq(parameter) - property value exactly matching 'parameter'
	// - $suffix(parameter) - property value ends with 'parameter'
	// - $contains(parameter) - property value contains 'parameter'
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Specify the IP address or name of the vCenter or standalone ESXi host:
	// Specify the IP address or name of the vCenter or standalone ESXi host:
	Ipaddress *string `json:"ipaddress,omitempty" tf:"ipaddress,omitempty"`

	// (String) Name this connection
	// Name this connection
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// (String) Provide user credentials for the vCenter or standalone ESXi host:
	// Provide user credentials for the vCenter or standalone ESXi host:
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type VmwareParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) This string should have one of the following formats:
	// This string should have one of the following formats:
	// - $prefix(parameter) - property value starting with 'parameter'
	// - $eq(parameter) - property value exactly matching 'parameter'
	// - $suffix(parameter) - property value ends with 'parameter'
	// - $contains(parameter) - property value contains 'parameter'
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) Specify the IP address or name of the vCenter or standalone ESXi host:
	// Specify the IP address or name of the vCenter or standalone ESXi host:
	// +kubebuilder:validation:Optional
	Ipaddress *string `json:"ipaddress,omitempty" tf:"ipaddress,omitempty"`

	// (String) Name this connection
	// Name this connection
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// (String, Sensitive) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// (String) Provide user credentials for the vCenter or standalone ESXi host:
	// Provide user credentials for the vCenter or standalone ESXi host:
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// VmwareSpec defines the desired state of Vmware
type VmwareSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VmwareParameters `json:"forProvider"`
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
	InitProvider VmwareInitParameters `json:"initProvider,omitempty"`
}

// VmwareStatus defines the observed state of Vmware.
type VmwareStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VmwareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Vmware is the Schema for the Vmwares API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type Vmware struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipaddress) || (has(self.initProvider) && has(self.initProvider.ipaddress))",message="spec.forProvider.ipaddress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.label) || (has(self.initProvider) && has(self.initProvider.label))",message="spec.forProvider.label is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   VmwareSpec   `json:"spec"`
	Status VmwareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VmwareList contains a list of Vmwares
type VmwareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vmware `json:"items"`
}

// Repository type metadata.
var (
	Vmware_Kind             = "Vmware"
	Vmware_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vmware_Kind}.String()
	Vmware_KindAPIVersion   = Vmware_Kind + "." + CRDGroupVersion.String()
	Vmware_GroupVersionKind = CRDGroupVersion.WithKind(Vmware_Kind)
)

func init() {
	SchemeBuilder.Register(&Vmware{}, &VmwareList{})
}