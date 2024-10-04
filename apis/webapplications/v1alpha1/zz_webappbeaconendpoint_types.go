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

type WebAppBeaconEndpointInitParameters struct {

	// The scope of this setting
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Possible Values: `ACTIVEGATE`, `DEFAULT_CONFIG`, `ONEAGENT`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// You can specify either path segments or an absolute URL.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Learn more about [sending beacon data via CORS](https://dt-url.net/r7038sa)
	UseCors *bool `json:"useCors,omitempty" tf:"use_cors,omitempty"`
}

type WebAppBeaconEndpointObservation struct {

	// The scope of this setting
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Possible Values: `ACTIVEGATE`, `DEFAULT_CONFIG`, `ONEAGENT`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// You can specify either path segments or an absolute URL.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Learn more about [sending beacon data via CORS](https://dt-url.net/r7038sa)
	UseCors *bool `json:"useCors,omitempty" tf:"use_cors,omitempty"`
}

type WebAppBeaconEndpointParameters struct {

	// The scope of this setting
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Possible Values: `ACTIVEGATE`, `DEFAULT_CONFIG`, `ONEAGENT`
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// You can specify either path segments or an absolute URL.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Learn more about [sending beacon data via CORS](https://dt-url.net/r7038sa)
	// +kubebuilder:validation:Optional
	UseCors *bool `json:"useCors,omitempty" tf:"use_cors,omitempty"`
}

// WebAppBeaconEndpointSpec defines the desired state of WebAppBeaconEndpoint
type WebAppBeaconEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebAppBeaconEndpointParameters `json:"forProvider"`
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
	InitProvider WebAppBeaconEndpointInitParameters `json:"initProvider,omitempty"`
}

// WebAppBeaconEndpointStatus defines the observed state of WebAppBeaconEndpoint.
type WebAppBeaconEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebAppBeaconEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// WebAppBeaconEndpoint is the Schema for the WebAppBeaconEndpoints API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type WebAppBeaconEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.applicationId) || (has(self.initProvider) && has(self.initProvider.applicationId))",message="spec.forProvider.applicationId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   WebAppBeaconEndpointSpec   `json:"spec"`
	Status WebAppBeaconEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebAppBeaconEndpointList contains a list of WebAppBeaconEndpoints
type WebAppBeaconEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebAppBeaconEndpoint `json:"items"`
}

// Repository type metadata.
var (
	WebAppBeaconEndpoint_Kind             = "WebAppBeaconEndpoint"
	WebAppBeaconEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebAppBeaconEndpoint_Kind}.String()
	WebAppBeaconEndpoint_KindAPIVersion   = WebAppBeaconEndpoint_Kind + "." + CRDGroupVersion.String()
	WebAppBeaconEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(WebAppBeaconEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&WebAppBeaconEndpoint{}, &WebAppBeaconEndpointList{})
}
