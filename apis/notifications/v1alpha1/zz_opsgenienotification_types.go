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

type OpsGenieNotificationInitParameters struct {

	// (String, Sensitive) The API key to access OpsGenie
	// The API key to access OpsGenie
	APIKeySecretRef *v1.SecretKeySelector `json:"apiKeySecretRef,omitempty" tf:"-"`

	// (Boolean) The configuration is enabled (true) or disabled (false)
	// The configuration is enabled (`true`) or disabled (`false`)
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// (String) The region domain of the OpsGenie
	// The region domain of the OpsGenie
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (String) The ID of these settings when referred to from resources requiring the REST API V1 keys
	// The ID of these settings when referred to from resources requiring the REST API V1 keys
	LegacyID *string `json:"legacyId,omitempty" tf:"legacy_id,omitempty"`

	// (String) The content of the message.  You can use the following placeholders:  * {ProblemID}: The display number of the reported problem.  * {ProblemImpact}: The impact level of the problem. Possible values are APPLICATION, SERVICE, and INFRASTRUCTURE.  * {ProblemSeverity}: The severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, and CUSTOM_ALERT.  * {ProblemTitle}: A short description of the problem
	// The content of the message.  You can use the following placeholders:  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) The name of the notification configuration
	// The name of the notification configuration
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the associated alerting profile
	// The ID of the associated alerting profile
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`
}

type OpsGenieNotificationObservation struct {

	// (Boolean) The configuration is enabled (true) or disabled (false)
	// The configuration is enabled (`true`) or disabled (`false`)
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// (String) The region domain of the OpsGenie
	// The region domain of the OpsGenie
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The ID of these settings when referred to from resources requiring the REST API V1 keys
	// The ID of these settings when referred to from resources requiring the REST API V1 keys
	LegacyID *string `json:"legacyId,omitempty" tf:"legacy_id,omitempty"`

	// (String) The content of the message.  You can use the following placeholders:  * {ProblemID}: The display number of the reported problem.  * {ProblemImpact}: The impact level of the problem. Possible values are APPLICATION, SERVICE, and INFRASTRUCTURE.  * {ProblemSeverity}: The severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, and CUSTOM_ALERT.  * {ProblemTitle}: A short description of the problem
	// The content of the message.  You can use the following placeholders:  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) The name of the notification configuration
	// The name of the notification configuration
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the associated alerting profile
	// The ID of the associated alerting profile
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`
}

type OpsGenieNotificationParameters struct {

	// (String, Sensitive) The API key to access OpsGenie
	// The API key to access OpsGenie
	// +kubebuilder:validation:Optional
	APIKeySecretRef *v1.SecretKeySelector `json:"apiKeySecretRef,omitempty" tf:"-"`

	// (Boolean) The configuration is enabled (true) or disabled (false)
	// The configuration is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// (String) The region domain of the OpsGenie
	// The region domain of the OpsGenie
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (String) The ID of these settings when referred to from resources requiring the REST API V1 keys
	// The ID of these settings when referred to from resources requiring the REST API V1 keys
	// +kubebuilder:validation:Optional
	LegacyID *string `json:"legacyId,omitempty" tf:"legacy_id,omitempty"`

	// (String) The content of the message.  You can use the following placeholders:  * {ProblemID}: The display number of the reported problem.  * {ProblemImpact}: The impact level of the problem. Possible values are APPLICATION, SERVICE, and INFRASTRUCTURE.  * {ProblemSeverity}: The severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, and CUSTOM_ALERT.  * {ProblemTitle}: A short description of the problem
	// The content of the message.  You can use the following placeholders:  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) The name of the notification configuration
	// The name of the notification configuration
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the associated alerting profile
	// The ID of the associated alerting profile
	// +kubebuilder:validation:Optional
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`
}

// OpsGenieNotificationSpec defines the desired state of OpsGenieNotification
type OpsGenieNotificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OpsGenieNotificationParameters `json:"forProvider"`
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
	InitProvider OpsGenieNotificationInitParameters `json:"initProvider,omitempty"`
}

// OpsGenieNotificationStatus defines the observed state of OpsGenieNotification.
type OpsGenieNotificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OpsGenieNotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OpsGenieNotification is the Schema for the OpsGenieNotifications API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type OpsGenieNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.active) || (has(self.initProvider) && has(self.initProvider.active))",message="spec.forProvider.active is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domain) || (has(self.initProvider) && has(self.initProvider.domain))",message="spec.forProvider.domain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.message) || (has(self.initProvider) && has(self.initProvider.message))",message="spec.forProvider.message is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.profile) || (has(self.initProvider) && has(self.initProvider.profile))",message="spec.forProvider.profile is a required parameter"
	Spec   OpsGenieNotificationSpec   `json:"spec"`
	Status OpsGenieNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsGenieNotificationList contains a list of OpsGenieNotifications
type OpsGenieNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsGenieNotification `json:"items"`
}

// Repository type metadata.
var (
	OpsGenieNotification_Kind             = "OpsGenieNotification"
	OpsGenieNotification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OpsGenieNotification_Kind}.String()
	OpsGenieNotification_KindAPIVersion   = OpsGenieNotification_Kind + "." + CRDGroupVersion.String()
	OpsGenieNotification_GroupVersionKind = CRDGroupVersion.WithKind(OpsGenieNotification_Kind)
)

func init() {
	SchemeBuilder.Register(&OpsGenieNotification{}, &OpsGenieNotificationList{})
}