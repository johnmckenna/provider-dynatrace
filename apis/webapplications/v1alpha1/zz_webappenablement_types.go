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

type RumInitParameters struct {

	// (Number) (Field has overlap with dynatrace_web_application) Percentage of user sessions captured and analyzed
	// (Field has overlap with `dynatrace_web_application`) Percentage of user sessions captured and analyzed
	CostAndTrafficControl *float64 `json:"costAndTrafficControl,omitempty" tf:"cost_and_traffic_control,omitempty"`

	// (Boolean) (Field has overlap with dynatrace_web_application) This setting is enabled (true) or disabled (false)
	// (Field has overlap with `dynatrace_web_application`) This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type RumObservation struct {

	// (Number) (Field has overlap with dynatrace_web_application) Percentage of user sessions captured and analyzed
	// (Field has overlap with `dynatrace_web_application`) Percentage of user sessions captured and analyzed
	CostAndTrafficControl *float64 `json:"costAndTrafficControl,omitempty" tf:"cost_and_traffic_control,omitempty"`

	// (Boolean) (Field has overlap with dynatrace_web_application) This setting is enabled (true) or disabled (false)
	// (Field has overlap with `dynatrace_web_application`) This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type RumParameters struct {

	// (Number) (Field has overlap with dynatrace_web_application) Percentage of user sessions captured and analyzed
	// (Field has overlap with `dynatrace_web_application`) Percentage of user sessions captured and analyzed
	// +kubebuilder:validation:Optional
	CostAndTrafficControl *float64 `json:"costAndTrafficControl" tf:"cost_and_traffic_control,omitempty"`

	// (Boolean) (Field has overlap with dynatrace_web_application) This setting is enabled (true) or disabled (false)
	// (Field has overlap with `dynatrace_web_application`) This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type SessionReplayInitParameters struct {

	// (Number) (Field has overlap with dynatrace_web_application) Percentage of user sessions captured and analyzed
	// (Field has overlap with `dynatrace_web_application`) [Percentage of user sessions recorded with Session Replay](https://dt-url.net/sr-cost-traffic-control). For example, if you have 50% for RUM and 50% for Session Replay, it results in 25% of sessions recorded with Session Replay.
	CostAndTrafficControl *float64 `json:"costAndTrafficControl,omitempty" tf:"cost_and_traffic_control,omitempty"`

	// (Boolean) (Field has overlap with dynatrace_web_application) This setting is enabled (true) or disabled (false)
	// (Field has overlap with `dynatrace_web_application`) This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SessionReplayObservation struct {

	// (Number) (Field has overlap with dynatrace_web_application) Percentage of user sessions captured and analyzed
	// (Field has overlap with `dynatrace_web_application`) [Percentage of user sessions recorded with Session Replay](https://dt-url.net/sr-cost-traffic-control). For example, if you have 50% for RUM and 50% for Session Replay, it results in 25% of sessions recorded with Session Replay.
	CostAndTrafficControl *float64 `json:"costAndTrafficControl,omitempty" tf:"cost_and_traffic_control,omitempty"`

	// (Boolean) (Field has overlap with dynatrace_web_application) This setting is enabled (true) or disabled (false)
	// (Field has overlap with `dynatrace_web_application`) This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SessionReplayParameters struct {

	// (Number) (Field has overlap with dynatrace_web_application) Percentage of user sessions captured and analyzed
	// (Field has overlap with `dynatrace_web_application`) [Percentage of user sessions recorded with Session Replay](https://dt-url.net/sr-cost-traffic-control). For example, if you have 50% for RUM and 50% for Session Replay, it results in 25% of sessions recorded with Session Replay.
	// +kubebuilder:validation:Optional
	CostAndTrafficControl *float64 `json:"costAndTrafficControl" tf:"cost_and_traffic_control,omitempty"`

	// (Boolean) (Field has overlap with dynatrace_web_application) This setting is enabled (true) or disabled (false)
	// (Field has overlap with `dynatrace_web_application`) This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type WebAppEnablementInitParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (Block List, Min: 1, Max: 1) Capture and analyze all user actions within your application. Enable Real User Monitoring (RUM) to monitor and improve your application's performance, identify errors, and gain insight into your user's behavior and experience. (see below for nested schema)
	// Capture and analyze all user actions within your application. Enable [Real User Monitoring (RUM)](https://dt-url.net/1n2b0prq) to monitor and improve your application's performance, identify errors, and gain insight into your user's behavior and experience.
	Rum []RumInitParameters `json:"rum,omitempty" tf:"rum,omitempty"`

	// like experience while providing best-in-class security and data protection. (see below for nested schema)
	// [Session Replay](https://dt-url.net/session-replay) captures all user interactions within your application and replays them in a movie-like experience while providing [best-in-class security and data protection](https://dt-url.net/b303zxj).
	SessionReplay []SessionReplayInitParameters `json:"sessionReplay,omitempty" tf:"session_replay,omitempty"`
}

type WebAppEnablementObservation struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Min: 1, Max: 1) Capture and analyze all user actions within your application. Enable Real User Monitoring (RUM) to monitor and improve your application's performance, identify errors, and gain insight into your user's behavior and experience. (see below for nested schema)
	// Capture and analyze all user actions within your application. Enable [Real User Monitoring (RUM)](https://dt-url.net/1n2b0prq) to monitor and improve your application's performance, identify errors, and gain insight into your user's behavior and experience.
	Rum []RumObservation `json:"rum,omitempty" tf:"rum,omitempty"`

	// like experience while providing best-in-class security and data protection. (see below for nested schema)
	// [Session Replay](https://dt-url.net/session-replay) captures all user interactions within your application and replays them in a movie-like experience while providing [best-in-class security and data protection](https://dt-url.net/b303zxj).
	SessionReplay []SessionReplayObservation `json:"sessionReplay,omitempty" tf:"session_replay,omitempty"`
}

type WebAppEnablementParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (Block List, Min: 1, Max: 1) Capture and analyze all user actions within your application. Enable Real User Monitoring (RUM) to monitor and improve your application's performance, identify errors, and gain insight into your user's behavior and experience. (see below for nested schema)
	// Capture and analyze all user actions within your application. Enable [Real User Monitoring (RUM)](https://dt-url.net/1n2b0prq) to monitor and improve your application's performance, identify errors, and gain insight into your user's behavior and experience.
	// +kubebuilder:validation:Optional
	Rum []RumParameters `json:"rum,omitempty" tf:"rum,omitempty"`

	// like experience while providing best-in-class security and data protection. (see below for nested schema)
	// [Session Replay](https://dt-url.net/session-replay) captures all user interactions within your application and replays them in a movie-like experience while providing [best-in-class security and data protection](https://dt-url.net/b303zxj).
	// +kubebuilder:validation:Optional
	SessionReplay []SessionReplayParameters `json:"sessionReplay,omitempty" tf:"session_replay,omitempty"`
}

// WebAppEnablementSpec defines the desired state of WebAppEnablement
type WebAppEnablementSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebAppEnablementParameters `json:"forProvider"`
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
	InitProvider WebAppEnablementInitParameters `json:"initProvider,omitempty"`
}

// WebAppEnablementStatus defines the observed state of WebAppEnablement.
type WebAppEnablementStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebAppEnablementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// WebAppEnablement is the Schema for the WebAppEnablements API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type WebAppEnablement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rum) || (has(self.initProvider) && has(self.initProvider.rum))",message="spec.forProvider.rum is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sessionReplay) || (has(self.initProvider) && has(self.initProvider.sessionReplay))",message="spec.forProvider.sessionReplay is a required parameter"
	Spec   WebAppEnablementSpec   `json:"spec"`
	Status WebAppEnablementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebAppEnablementList contains a list of WebAppEnablements
type WebAppEnablementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebAppEnablement `json:"items"`
}

// Repository type metadata.
var (
	WebAppEnablement_Kind             = "WebAppEnablement"
	WebAppEnablement_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebAppEnablement_Kind}.String()
	WebAppEnablement_KindAPIVersion   = WebAppEnablement_Kind + "." + CRDGroupVersion.String()
	WebAppEnablement_GroupVersionKind = CRDGroupVersion.WithKind(WebAppEnablement_Kind)
)

func init() {
	SchemeBuilder.Register(&WebAppEnablement{}, &WebAppEnablementList{})
}
