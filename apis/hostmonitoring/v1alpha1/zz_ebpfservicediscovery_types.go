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

type EbpfServiceDiscoveryInitParameters struct {

	// (Boolean) When disabled, Dynatrace can only detect services in Full stack mode.
	// When disabled, Dynatrace can only detect services in Full stack mode.
	Ebpf *bool `json:"ebpf,omitempty" tf:"ebpf,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type EbpfServiceDiscoveryObservation struct {

	// (Boolean) When disabled, Dynatrace can only detect services in Full stack mode.
	// When disabled, Dynatrace can only detect services in Full stack mode.
	Ebpf *bool `json:"ebpf,omitempty" tf:"ebpf,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type EbpfServiceDiscoveryParameters struct {

	// (Boolean) When disabled, Dynatrace can only detect services in Full stack mode.
	// When disabled, Dynatrace can only detect services in Full stack mode.
	// +kubebuilder:validation:Optional
	Ebpf *bool `json:"ebpf,omitempty" tf:"ebpf,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// EbpfServiceDiscoverySpec defines the desired state of EbpfServiceDiscovery
type EbpfServiceDiscoverySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EbpfServiceDiscoveryParameters `json:"forProvider"`
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
	InitProvider EbpfServiceDiscoveryInitParameters `json:"initProvider,omitempty"`
}

// EbpfServiceDiscoveryStatus defines the observed state of EbpfServiceDiscovery.
type EbpfServiceDiscoveryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EbpfServiceDiscoveryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EbpfServiceDiscovery is the Schema for the EbpfServiceDiscoverys API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type EbpfServiceDiscovery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ebpf) || (has(self.initProvider) && has(self.initProvider.ebpf))",message="spec.forProvider.ebpf is a required parameter"
	Spec   EbpfServiceDiscoverySpec   `json:"spec"`
	Status EbpfServiceDiscoveryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EbpfServiceDiscoveryList contains a list of EbpfServiceDiscoverys
type EbpfServiceDiscoveryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EbpfServiceDiscovery `json:"items"`
}

// Repository type metadata.
var (
	EbpfServiceDiscovery_Kind             = "EbpfServiceDiscovery"
	EbpfServiceDiscovery_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EbpfServiceDiscovery_Kind}.String()
	EbpfServiceDiscovery_KindAPIVersion   = EbpfServiceDiscovery_Kind + "." + CRDGroupVersion.String()
	EbpfServiceDiscovery_GroupVersionKind = CRDGroupVersion.WithKind(EbpfServiceDiscovery_Kind)
)

func init() {
	SchemeBuilder.Register(&EbpfServiceDiscovery{}, &EbpfServiceDiscoveryList{})
}