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

type HubPermissionsInitParameters struct {

	// (String) Name
	// Name
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Contact Email
	// Contact Email
	Email *string `json:"email,omitempty" tf:"email,omitempty"`
}

type HubPermissionsObservation struct {

	// (String) Name
	// Name
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Contact Email
	// Contact Email
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HubPermissionsParameters struct {

	// (String) Name
	// Name
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Contact Email
	// Contact Email
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`
}

// HubPermissionsSpec defines the desired state of HubPermissions
type HubPermissionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HubPermissionsParameters `json:"forProvider"`
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
	InitProvider HubPermissionsInitParameters `json:"initProvider,omitempty"`
}

// HubPermissionsStatus defines the observed state of HubPermissions.
type HubPermissionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HubPermissionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HubPermissions is the Schema for the HubPermissionss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type HubPermissions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || (has(self.initProvider) && has(self.initProvider.email))",message="spec.forProvider.email is a required parameter"
	Spec   HubPermissionsSpec   `json:"spec"`
	Status HubPermissionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HubPermissionsList contains a list of HubPermissionss
type HubPermissionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HubPermissions `json:"items"`
}

// Repository type metadata.
var (
	HubPermissions_Kind             = "HubPermissions"
	HubPermissions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HubPermissions_Kind}.String()
	HubPermissions_KindAPIVersion   = HubPermissions_Kind + "." + CRDGroupVersion.String()
	HubPermissions_GroupVersionKind = CRDGroupVersion.WithKind(HubPermissions_Kind)
)

func init() {
	SchemeBuilder.Register(&HubPermissions{}, &HubPermissionsList{})
}
