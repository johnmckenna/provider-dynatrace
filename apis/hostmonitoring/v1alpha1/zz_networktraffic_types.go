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

type ExcludeIPInitParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	IPAddressForm []IPAddressFormInitParameters `json:"ipAddressForm,omitempty" tf:"ip_address_form,omitempty"`
}

type ExcludeIPObservation struct {

	// (Block List, Min: 1) (see below for nested schema)
	IPAddressForm []IPAddressFormObservation `json:"ipAddressForm,omitempty" tf:"ip_address_form,omitempty"`
}

type ExcludeIPParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	IPAddressForm []IPAddressFormParameters `json:"ipAddressForm" tf:"ip_address_form,omitempty"`
}

type ExcludeNicInitParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	NicForm []NicFormInitParameters `json:"nicForm,omitempty" tf:"nic_form,omitempty"`
}

type ExcludeNicObservation struct {

	// (Block List, Min: 1) (see below for nested schema)
	NicForm []NicFormObservation `json:"nicForm,omitempty" tf:"nic_form,omitempty"`
}

type ExcludeNicParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	NicForm []NicFormParameters `json:"nicForm" tf:"nic_form,omitempty"`
}

type IPAddressFormInitParameters struct {

	// (String) IP address
	// IP address
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
}

type IPAddressFormObservation struct {

	// (String) IP address
	// IP address
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
}

type IPAddressFormParameters struct {

	// (String) IP address
	// IP address
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`
}

type NetworkTrafficInitParameters struct {

	// (Block List, Max: 1) Providing a host IP address, you will exclude network traffic only in calculating connectivity (other metrics will still be calculated). (see below for nested schema)
	// Providing a host IP address, you will exclude network traffic only in calculating connectivity (other metrics will still be calculated).
	ExcludeIP []ExcludeIPInitParameters `json:"excludeIp,omitempty" tf:"exclude_ip,omitempty"`

	// (Block List, Max: 1) Selecting a network interface, you will exclude all network traffic on that interface from being monitored. You can select from the list below what to not monitor, or input it manually using the "other one" option. (see below for nested schema)
	// Selecting a network interface, you will exclude all network traffic on that interface from being monitored. You can select from the list below what to not monitor, or input it manually using the "other one" option.
	ExcludeNic []ExcludeNicInitParameters `json:"excludeNic,omitempty" tf:"exclude_nic,omitempty"`

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`
}

type NetworkTrafficObservation struct {

	// (Block List, Max: 1) Providing a host IP address, you will exclude network traffic only in calculating connectivity (other metrics will still be calculated). (see below for nested schema)
	// Providing a host IP address, you will exclude network traffic only in calculating connectivity (other metrics will still be calculated).
	ExcludeIP []ExcludeIPObservation `json:"excludeIp,omitempty" tf:"exclude_ip,omitempty"`

	// (Block List, Max: 1) Selecting a network interface, you will exclude all network traffic on that interface from being monitored. You can select from the list below what to not monitor, or input it manually using the "other one" option. (see below for nested schema)
	// Selecting a network interface, you will exclude all network traffic on that interface from being monitored. You can select from the list below what to not monitor, or input it manually using the "other one" option.
	ExcludeNic []ExcludeNicObservation `json:"excludeNic,omitempty" tf:"exclude_nic,omitempty"`

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NetworkTrafficParameters struct {

	// (Block List, Max: 1) Providing a host IP address, you will exclude network traffic only in calculating connectivity (other metrics will still be calculated). (see below for nested schema)
	// Providing a host IP address, you will exclude network traffic only in calculating connectivity (other metrics will still be calculated).
	// +kubebuilder:validation:Optional
	ExcludeIP []ExcludeIPParameters `json:"excludeIp,omitempty" tf:"exclude_ip,omitempty"`

	// (Block List, Max: 1) Selecting a network interface, you will exclude all network traffic on that interface from being monitored. You can select from the list below what to not monitor, or input it manually using the "other one" option. (see below for nested schema)
	// Selecting a network interface, you will exclude all network traffic on that interface from being monitored. You can select from the list below what to not monitor, or input it manually using the "other one" option.
	// +kubebuilder:validation:Optional
	ExcludeNic []ExcludeNicParameters `json:"excludeNic,omitempty" tf:"exclude_nic,omitempty"`

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// +kubebuilder:validation:Optional
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`
}

type NicFormInitParameters struct {

	// (String) Network interface
	// Network interface
	Interface *string `json:"interface,omitempty" tf:"interface,omitempty"`

	// (String) Possible Values: OS_TYPE_AIX, OS_TYPE_DARWIN, OS_TYPE_HPUX, OS_TYPE_LINUX, OS_TYPE_SOLARIS, OS_TYPE_UNKNOWN, OS_TYPE_WINDOWS, OS_TYPE_ZOS
	// Possible Values: `OS_TYPE_AIX`, `OS_TYPE_DARWIN`, `OS_TYPE_HPUX`, `OS_TYPE_LINUX`, `OS_TYPE_SOLARIS`, `OS_TYPE_UNKNOWN`, `OS_TYPE_WINDOWS`, `OS_TYPE_ZOS`
	Os *string `json:"os,omitempty" tf:"os,omitempty"`
}

type NicFormObservation struct {

	// (String) Network interface
	// Network interface
	Interface *string `json:"interface,omitempty" tf:"interface,omitempty"`

	// (String) Possible Values: OS_TYPE_AIX, OS_TYPE_DARWIN, OS_TYPE_HPUX, OS_TYPE_LINUX, OS_TYPE_SOLARIS, OS_TYPE_UNKNOWN, OS_TYPE_WINDOWS, OS_TYPE_ZOS
	// Possible Values: `OS_TYPE_AIX`, `OS_TYPE_DARWIN`, `OS_TYPE_HPUX`, `OS_TYPE_LINUX`, `OS_TYPE_SOLARIS`, `OS_TYPE_UNKNOWN`, `OS_TYPE_WINDOWS`, `OS_TYPE_ZOS`
	Os *string `json:"os,omitempty" tf:"os,omitempty"`
}

type NicFormParameters struct {

	// (String) Network interface
	// Network interface
	// +kubebuilder:validation:Optional
	Interface *string `json:"interface" tf:"interface,omitempty"`

	// (String) Possible Values: OS_TYPE_AIX, OS_TYPE_DARWIN, OS_TYPE_HPUX, OS_TYPE_LINUX, OS_TYPE_SOLARIS, OS_TYPE_UNKNOWN, OS_TYPE_WINDOWS, OS_TYPE_ZOS
	// Possible Values: `OS_TYPE_AIX`, `OS_TYPE_DARWIN`, `OS_TYPE_HPUX`, `OS_TYPE_LINUX`, `OS_TYPE_SOLARIS`, `OS_TYPE_UNKNOWN`, `OS_TYPE_WINDOWS`, `OS_TYPE_ZOS`
	// +kubebuilder:validation:Optional
	Os *string `json:"os" tf:"os,omitempty"`
}

// NetworkTrafficSpec defines the desired state of NetworkTraffic
type NetworkTrafficSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkTrafficParameters `json:"forProvider"`
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
	InitProvider NetworkTrafficInitParameters `json:"initProvider,omitempty"`
}

// NetworkTrafficStatus defines the observed state of NetworkTraffic.
type NetworkTrafficStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkTrafficObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NetworkTraffic is the Schema for the NetworkTraffics API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type NetworkTraffic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostId) || (has(self.initProvider) && has(self.initProvider.hostId))",message="spec.forProvider.hostId is a required parameter"
	Spec   NetworkTrafficSpec   `json:"spec"`
	Status NetworkTrafficStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkTrafficList contains a list of NetworkTraffics
type NetworkTrafficList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkTraffic `json:"items"`
}

// Repository type metadata.
var (
	NetworkTraffic_Kind             = "NetworkTraffic"
	NetworkTraffic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkTraffic_Kind}.String()
	NetworkTraffic_KindAPIVersion   = NetworkTraffic_Kind + "." + CRDGroupVersion.String()
	NetworkTraffic_GroupVersionKind = CRDGroupVersion.WithKind(NetworkTraffic_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkTraffic{}, &NetworkTrafficList{})
}