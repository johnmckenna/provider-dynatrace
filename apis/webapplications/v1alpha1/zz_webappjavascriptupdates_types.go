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

type WebAppJavascriptUpdatesInitParameters struct {

	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Possible Values: `CUSTOM`, `LATEST_IE7_10_SUPPORTED`, `LATEST_STABLE`, `PREVIOUS_STABLE`
	JavascriptVersion *string `json:"javascriptVersion,omitempty" tf:"javascript_version,omitempty"`
}

type WebAppJavascriptUpdatesObservation struct {

	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Possible Values: `CUSTOM`, `LATEST_IE7_10_SUPPORTED`, `LATEST_STABLE`, `PREVIOUS_STABLE`
	JavascriptVersion *string `json:"javascriptVersion,omitempty" tf:"javascript_version,omitempty"`
}

type WebAppJavascriptUpdatesParameters struct {

	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Possible Values: `CUSTOM`, `LATEST_IE7_10_SUPPORTED`, `LATEST_STABLE`, `PREVIOUS_STABLE`
	// +kubebuilder:validation:Optional
	JavascriptVersion *string `json:"javascriptVersion,omitempty" tf:"javascript_version,omitempty"`
}

// WebAppJavascriptUpdatesSpec defines the desired state of WebAppJavascriptUpdates
type WebAppJavascriptUpdatesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebAppJavascriptUpdatesParameters `json:"forProvider"`
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
	InitProvider WebAppJavascriptUpdatesInitParameters `json:"initProvider,omitempty"`
}

// WebAppJavascriptUpdatesStatus defines the observed state of WebAppJavascriptUpdates.
type WebAppJavascriptUpdatesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebAppJavascriptUpdatesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// WebAppJavascriptUpdates is the Schema for the WebAppJavascriptUpdatess API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type WebAppJavascriptUpdates struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.javascriptVersion) || (has(self.initProvider) && has(self.initProvider.javascriptVersion))",message="spec.forProvider.javascriptVersion is a required parameter"
	Spec   WebAppJavascriptUpdatesSpec   `json:"spec"`
	Status WebAppJavascriptUpdatesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebAppJavascriptUpdatesList contains a list of WebAppJavascriptUpdatess
type WebAppJavascriptUpdatesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebAppJavascriptUpdates `json:"items"`
}

// Repository type metadata.
var (
	WebAppJavascriptUpdates_Kind             = "WebAppJavascriptUpdates"
	WebAppJavascriptUpdates_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebAppJavascriptUpdates_Kind}.String()
	WebAppJavascriptUpdates_KindAPIVersion   = WebAppJavascriptUpdates_Kind + "." + CRDGroupVersion.String()
	WebAppJavascriptUpdates_GroupVersionKind = CRDGroupVersion.WithKind(WebAppJavascriptUpdates_Kind)
)

func init() {
	SchemeBuilder.Register(&WebAppJavascriptUpdates{}, &WebAppJavascriptUpdatesList{})
}
