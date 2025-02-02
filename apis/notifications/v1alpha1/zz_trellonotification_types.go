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

type TrelloNotificationInitParameters struct {

	// (Boolean) The configuration is enabled (true) or disabled (false)
	// The configuration is enabled (`true`) or disabled (`false`)
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// (String) The application key for the Trello account
	// The application key for the Trello account
	ApplicationKey *string `json:"applicationKey,omitempty" tf:"application_key,omitempty"`

	// (String, Sensitive) The application token for the Trello account
	// The application token for the Trello account
	AuthorizationTokenSecretRef *v1.SecretKeySelector `json:"authorizationTokenSecretRef,omitempty" tf:"-"`

	// (String) The Trello board to which the card should be assigned
	// The Trello board to which the card should be assigned
	BoardID *string `json:"boardId,omitempty" tf:"board_id,omitempty"`

	// (String) The description of the Trello card.   You can use same placeholders as in card text
	// The description of the Trello card.   You can use same placeholders as in card text
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of these settings when referred to from resources requiring the REST API V1 keys
	// The ID of these settings when referred to from resources requiring the REST API V1 keys
	LegacyID *string `json:"legacyId,omitempty" tf:"legacy_id,omitempty"`

	// (String) The Trello list to which the card should be assigned
	// The Trello list to which the card should be assigned
	ListID *string `json:"listId,omitempty" tf:"list_id,omitempty"`

	// (String) The name of the notification configuration
	// The name of the notification configuration
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the associated alerting profile
	// The ID of the associated alerting profile
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`

	// (String) The Trello list to which the card of the resolved problem should be assigned
	// The Trello list to which the card of the resolved problem should be assigned
	ResolvedListID *string `json:"resolvedListId,omitempty" tf:"resolved_list_id,omitempty"`

	// formatted string.  * {ProblemID}: The display number of the reported problem.  * {ProblemImpact}: The impact level of the problem. Possible values are APPLICATION, SERVICE, and INFRASTRUCTURE.  * {ProblemSeverity}: The severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, and CUSTOM_ALERT.  * {ProblemTitle}: A short description of the problem.  * {ProblemURL}: The URL of the problem within Dynatrace.  * {State}: The state of the problem. Possible values are OPEN and RESOLVED.  * {Tags}: The list of tags that are defined for all impacted entities, separated by commas
	// The text of the generated Trello card.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas
	Text *string `json:"text,omitempty" tf:"text,omitempty"`
}

type TrelloNotificationObservation struct {

	// (Boolean) The configuration is enabled (true) or disabled (false)
	// The configuration is enabled (`true`) or disabled (`false`)
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// (String) The application key for the Trello account
	// The application key for the Trello account
	ApplicationKey *string `json:"applicationKey,omitempty" tf:"application_key,omitempty"`

	// (String) The Trello board to which the card should be assigned
	// The Trello board to which the card should be assigned
	BoardID *string `json:"boardId,omitempty" tf:"board_id,omitempty"`

	// (String) The description of the Trello card.   You can use same placeholders as in card text
	// The description of the Trello card.   You can use same placeholders as in card text
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The ID of these settings when referred to from resources requiring the REST API V1 keys
	// The ID of these settings when referred to from resources requiring the REST API V1 keys
	LegacyID *string `json:"legacyId,omitempty" tf:"legacy_id,omitempty"`

	// (String) The Trello list to which the card should be assigned
	// The Trello list to which the card should be assigned
	ListID *string `json:"listId,omitempty" tf:"list_id,omitempty"`

	// (String) The name of the notification configuration
	// The name of the notification configuration
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the associated alerting profile
	// The ID of the associated alerting profile
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`

	// (String) The Trello list to which the card of the resolved problem should be assigned
	// The Trello list to which the card of the resolved problem should be assigned
	ResolvedListID *string `json:"resolvedListId,omitempty" tf:"resolved_list_id,omitempty"`

	// formatted string.  * {ProblemID}: The display number of the reported problem.  * {ProblemImpact}: The impact level of the problem. Possible values are APPLICATION, SERVICE, and INFRASTRUCTURE.  * {ProblemSeverity}: The severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, and CUSTOM_ALERT.  * {ProblemTitle}: A short description of the problem.  * {ProblemURL}: The URL of the problem within Dynatrace.  * {State}: The state of the problem. Possible values are OPEN and RESOLVED.  * {Tags}: The list of tags that are defined for all impacted entities, separated by commas
	// The text of the generated Trello card.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas
	Text *string `json:"text,omitempty" tf:"text,omitempty"`
}

type TrelloNotificationParameters struct {

	// (Boolean) The configuration is enabled (true) or disabled (false)
	// The configuration is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// (String) The application key for the Trello account
	// The application key for the Trello account
	// +kubebuilder:validation:Optional
	ApplicationKey *string `json:"applicationKey,omitempty" tf:"application_key,omitempty"`

	// (String, Sensitive) The application token for the Trello account
	// The application token for the Trello account
	// +kubebuilder:validation:Optional
	AuthorizationTokenSecretRef *v1.SecretKeySelector `json:"authorizationTokenSecretRef,omitempty" tf:"-"`

	// (String) The Trello board to which the card should be assigned
	// The Trello board to which the card should be assigned
	// +kubebuilder:validation:Optional
	BoardID *string `json:"boardId,omitempty" tf:"board_id,omitempty"`

	// (String) The description of the Trello card.   You can use same placeholders as in card text
	// The description of the Trello card.   You can use same placeholders as in card text
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of these settings when referred to from resources requiring the REST API V1 keys
	// The ID of these settings when referred to from resources requiring the REST API V1 keys
	// +kubebuilder:validation:Optional
	LegacyID *string `json:"legacyId,omitempty" tf:"legacy_id,omitempty"`

	// (String) The Trello list to which the card should be assigned
	// The Trello list to which the card should be assigned
	// +kubebuilder:validation:Optional
	ListID *string `json:"listId,omitempty" tf:"list_id,omitempty"`

	// (String) The name of the notification configuration
	// The name of the notification configuration
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the associated alerting profile
	// The ID of the associated alerting profile
	// +kubebuilder:validation:Optional
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`

	// (String) The Trello list to which the card of the resolved problem should be assigned
	// The Trello list to which the card of the resolved problem should be assigned
	// +kubebuilder:validation:Optional
	ResolvedListID *string `json:"resolvedListId,omitempty" tf:"resolved_list_id,omitempty"`

	// formatted string.  * {ProblemID}: The display number of the reported problem.  * {ProblemImpact}: The impact level of the problem. Possible values are APPLICATION, SERVICE, and INFRASTRUCTURE.  * {ProblemSeverity}: The severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, and CUSTOM_ALERT.  * {ProblemTitle}: A short description of the problem.  * {ProblemURL}: The URL of the problem within Dynatrace.  * {State}: The state of the problem. Possible values are OPEN and RESOLVED.  * {Tags}: The list of tags that are defined for all impacted entities, separated by commas
	// The text of the generated Trello card.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas
	// +kubebuilder:validation:Optional
	Text *string `json:"text,omitempty" tf:"text,omitempty"`
}

// TrelloNotificationSpec defines the desired state of TrelloNotification
type TrelloNotificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrelloNotificationParameters `json:"forProvider"`
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
	InitProvider TrelloNotificationInitParameters `json:"initProvider,omitempty"`
}

// TrelloNotificationStatus defines the observed state of TrelloNotification.
type TrelloNotificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrelloNotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TrelloNotification is the Schema for the TrelloNotifications API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type TrelloNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.active) || (has(self.initProvider) && has(self.initProvider.active))",message="spec.forProvider.active is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.applicationKey) || (has(self.initProvider) && has(self.initProvider.applicationKey))",message="spec.forProvider.applicationKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.boardId) || (has(self.initProvider) && has(self.initProvider.boardId))",message="spec.forProvider.boardId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.listId) || (has(self.initProvider) && has(self.initProvider.listId))",message="spec.forProvider.listId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.profile) || (has(self.initProvider) && has(self.initProvider.profile))",message="spec.forProvider.profile is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resolvedListId) || (has(self.initProvider) && has(self.initProvider.resolvedListId))",message="spec.forProvider.resolvedListId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.text) || (has(self.initProvider) && has(self.initProvider.text))",message="spec.forProvider.text is a required parameter"
	Spec   TrelloNotificationSpec   `json:"spec"`
	Status TrelloNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrelloNotificationList contains a list of TrelloNotifications
type TrelloNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrelloNotification `json:"items"`
}

// Repository type metadata.
var (
	TrelloNotification_Kind             = "TrelloNotification"
	TrelloNotification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrelloNotification_Kind}.String()
	TrelloNotification_KindAPIVersion   = TrelloNotification_Kind + "." + CRDGroupVersion.String()
	TrelloNotification_GroupVersionKind = CRDGroupVersion.WithKind(TrelloNotification_Kind)
)

func init() {
	SchemeBuilder.Register(&TrelloNotification{}, &TrelloNotificationList{})
}
