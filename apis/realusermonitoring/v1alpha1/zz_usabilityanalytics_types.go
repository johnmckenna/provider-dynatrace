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

type UsabilityAnalyticsInitParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// loading or failed page resources. Rage click counts are compiled for each session and considered in the User Experience Score .
	// With this setting enabled, a rage click count is compiled for each monitored user session.
	// Three or more rapid clicks within the same area of a web page are considered to be rage clicks. Rage clicks commonly reflect slow-loading or failed page resources. Rage click counts are compiled for each session and considered in the [User Experience Score](https://dt-url.net/39034wt) .
	// With this setting enabled, a rage click count is compiled for each monitored user session.
	DetectRageClicks *bool `json:"detectRageClicks,omitempty" tf:"detect_rage_clicks,omitempty"`
}

type UsabilityAnalyticsObservation struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// loading or failed page resources. Rage click counts are compiled for each session and considered in the User Experience Score .
	// With this setting enabled, a rage click count is compiled for each monitored user session.
	// Three or more rapid clicks within the same area of a web page are considered to be rage clicks. Rage clicks commonly reflect slow-loading or failed page resources. Rage click counts are compiled for each session and considered in the [User Experience Score](https://dt-url.net/39034wt) .
	// With this setting enabled, a rage click count is compiled for each monitored user session.
	DetectRageClicks *bool `json:"detectRageClicks,omitempty" tf:"detect_rage_clicks,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type UsabilityAnalyticsParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// loading or failed page resources. Rage click counts are compiled for each session and considered in the User Experience Score .
	// With this setting enabled, a rage click count is compiled for each monitored user session.
	// Three or more rapid clicks within the same area of a web page are considered to be rage clicks. Rage clicks commonly reflect slow-loading or failed page resources. Rage click counts are compiled for each session and considered in the [User Experience Score](https://dt-url.net/39034wt) .
	// With this setting enabled, a rage click count is compiled for each monitored user session.
	// +kubebuilder:validation:Optional
	DetectRageClicks *bool `json:"detectRageClicks,omitempty" tf:"detect_rage_clicks,omitempty"`
}

// UsabilityAnalyticsSpec defines the desired state of UsabilityAnalytics
type UsabilityAnalyticsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UsabilityAnalyticsParameters `json:"forProvider"`
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
	InitProvider UsabilityAnalyticsInitParameters `json:"initProvider,omitempty"`
}

// UsabilityAnalyticsStatus defines the observed state of UsabilityAnalytics.
type UsabilityAnalyticsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UsabilityAnalyticsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UsabilityAnalytics is the Schema for the UsabilityAnalyticss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type UsabilityAnalytics struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.detectRageClicks) || (has(self.initProvider) && has(self.initProvider.detectRageClicks))",message="spec.forProvider.detectRageClicks is a required parameter"
	Spec   UsabilityAnalyticsSpec   `json:"spec"`
	Status UsabilityAnalyticsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UsabilityAnalyticsList contains a list of UsabilityAnalyticss
type UsabilityAnalyticsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UsabilityAnalytics `json:"items"`
}

// Repository type metadata.
var (
	UsabilityAnalytics_Kind             = "UsabilityAnalytics"
	UsabilityAnalytics_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UsabilityAnalytics_Kind}.String()
	UsabilityAnalytics_KindAPIVersion   = UsabilityAnalytics_Kind + "." + CRDGroupVersion.String()
	UsabilityAnalytics_GroupVersionKind = CRDGroupVersion.WithKind(UsabilityAnalytics_Kind)
)

func init() {
	SchemeBuilder.Register(&UsabilityAnalytics{}, &UsabilityAnalyticsList{})
}
