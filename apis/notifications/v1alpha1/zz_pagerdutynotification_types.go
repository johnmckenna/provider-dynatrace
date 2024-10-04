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

type PagerDutyNotificationInitParameters struct {

	// (String, Sensitive) The API key to access PagerDuty
	// The API key to access PagerDuty
	APIKeySecretRef *v1.SecretKeySelector `json:"apiKeySecretRef,omitempty" tf:"-"`

	// (String) The name of the PagerDuty account
	// The name of the PagerDuty account
	Account *string `json:"account,omitempty" tf:"account,omitempty"`

	// (Boolean) The configuration is enabled (true) or disabled (false)
	// The configuration is enabled (`true`) or disabled (`false`)
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// (String) The ID of these settings when referred to from resources requiring the REST API V1 keys
	// The ID of these settings when referred to from resources requiring the REST API V1 keys
	LegacyID *string `json:"legacyId,omitempty" tf:"legacy_id,omitempty"`

	// (String) The name of the notification configuration
	// The name of the notification configuration
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the associated alerting profile
	// The ID of the associated alerting profile
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`

	// (String) The name of the PagerDuty Service
	// The name of the PagerDuty Service
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type PagerDutyNotificationObservation struct {

	// (String) The name of the PagerDuty account
	// The name of the PagerDuty account
	Account *string `json:"account,omitempty" tf:"account,omitempty"`

	// (Boolean) The configuration is enabled (true) or disabled (false)
	// The configuration is enabled (`true`) or disabled (`false`)
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The ID of these settings when referred to from resources requiring the REST API V1 keys
	// The ID of these settings when referred to from resources requiring the REST API V1 keys
	LegacyID *string `json:"legacyId,omitempty" tf:"legacy_id,omitempty"`

	// (String) The name of the notification configuration
	// The name of the notification configuration
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the associated alerting profile
	// The ID of the associated alerting profile
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`

	// (String) The name of the PagerDuty Service
	// The name of the PagerDuty Service
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type PagerDutyNotificationParameters struct {

	// (String, Sensitive) The API key to access PagerDuty
	// The API key to access PagerDuty
	// +kubebuilder:validation:Optional
	APIKeySecretRef *v1.SecretKeySelector `json:"apiKeySecretRef,omitempty" tf:"-"`

	// (String) The name of the PagerDuty account
	// The name of the PagerDuty account
	// +kubebuilder:validation:Optional
	Account *string `json:"account,omitempty" tf:"account,omitempty"`

	// (Boolean) The configuration is enabled (true) or disabled (false)
	// The configuration is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// (String) The ID of these settings when referred to from resources requiring the REST API V1 keys
	// The ID of these settings when referred to from resources requiring the REST API V1 keys
	// +kubebuilder:validation:Optional
	LegacyID *string `json:"legacyId,omitempty" tf:"legacy_id,omitempty"`

	// (String) The name of the notification configuration
	// The name of the notification configuration
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the associated alerting profile
	// The ID of the associated alerting profile
	// +kubebuilder:validation:Optional
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`

	// (String) The name of the PagerDuty Service
	// The name of the PagerDuty Service
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

// PagerDutyNotificationSpec defines the desired state of PagerDutyNotification
type PagerDutyNotificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PagerDutyNotificationParameters `json:"forProvider"`
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
	InitProvider PagerDutyNotificationInitParameters `json:"initProvider,omitempty"`
}

// PagerDutyNotificationStatus defines the observed state of PagerDutyNotification.
type PagerDutyNotificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PagerDutyNotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PagerDutyNotification is the Schema for the PagerDutyNotifications API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type PagerDutyNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.account) || (has(self.initProvider) && has(self.initProvider.account))",message="spec.forProvider.account is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.active) || (has(self.initProvider) && has(self.initProvider.active))",message="spec.forProvider.active is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.profile) || (has(self.initProvider) && has(self.initProvider.profile))",message="spec.forProvider.profile is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.service) || (has(self.initProvider) && has(self.initProvider.service))",message="spec.forProvider.service is a required parameter"
	Spec   PagerDutyNotificationSpec   `json:"spec"`
	Status PagerDutyNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PagerDutyNotificationList contains a list of PagerDutyNotifications
type PagerDutyNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PagerDutyNotification `json:"items"`
}

// Repository type metadata.
var (
	PagerDutyNotification_Kind             = "PagerDutyNotification"
	PagerDutyNotification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PagerDutyNotification_Kind}.String()
	PagerDutyNotification_KindAPIVersion   = PagerDutyNotification_Kind + "." + CRDGroupVersion.String()
	PagerDutyNotification_GroupVersionKind = CRDGroupVersion.WithKind(PagerDutyNotification_Kind)
)

func init() {
	SchemeBuilder.Register(&PagerDutyNotification{}, &PagerDutyNotificationList{})
}