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

type CookieInitParameters struct {

	// (String) Enclose placeholder values in brackets, for example {email}
	// Enclose placeholder values in brackets, for example \{email\}
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (String) Enclose placeholder values in brackets, for example {email}
	// Enclose placeholder values in brackets, for example \{email\}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Enclose placeholder values in brackets, for example {email}
	// Enclose placeholder values in brackets, for example \{email\}
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Enclose placeholder values in brackets, for example {email}
	// Enclose placeholder values in brackets, for example \{email\}
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CookieObservation struct {

	// (String) Enclose placeholder values in brackets, for example {email}
	// Enclose placeholder values in brackets, for example \{email\}
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (String) Enclose placeholder values in brackets, for example {email}
	// Enclose placeholder values in brackets, for example \{email\}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Enclose placeholder values in brackets, for example {email}
	// Enclose placeholder values in brackets, for example \{email\}
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Enclose placeholder values in brackets, for example {email}
	// Enclose placeholder values in brackets, for example \{email\}
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CookieParameters struct {

	// (String) Enclose placeholder values in brackets, for example {email}
	// Enclose placeholder values in brackets, for example \{email\}
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// (String) Enclose placeholder values in brackets, for example {email}
	// Enclose placeholder values in brackets, for example \{email\}
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) Enclose placeholder values in brackets, for example {email}
	// Enclose placeholder values in brackets, for example \{email\}
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Enclose placeholder values in brackets, for example {email}
	// Enclose placeholder values in brackets, for example \{email\}
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type CookiesInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Cookie []CookieInitParameters `json:"cookie,omitempty" tf:"cookie,omitempty"`
}

type CookiesObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Cookie []CookieObservation `json:"cookie,omitempty" tf:"cookie,omitempty"`
}

type CookiesParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Cookie []CookieParameters `json:"cookie" tf:"cookie,omitempty"`
}

type HTTPMonitorCookiesInitParameters struct {

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	Cookies []CookiesInitParameters `json:"cookies,omitempty" tf:"cookies,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The scope of this setting (HTTP_CHECK)
	// The scope of this setting (HTTP_CHECK)
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type HTTPMonitorCookiesObservation struct {

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	Cookies []CookiesObservation `json:"cookies,omitempty" tf:"cookies,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The scope of this setting (HTTP_CHECK)
	// The scope of this setting (HTTP_CHECK)
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type HTTPMonitorCookiesParameters struct {

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	// +kubebuilder:validation:Optional
	Cookies []CookiesParameters `json:"cookies,omitempty" tf:"cookies,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The scope of this setting (HTTP_CHECK)
	// The scope of this setting (HTTP_CHECK)
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// HTTPMonitorCookiesSpec defines the desired state of HTTPMonitorCookies
type HTTPMonitorCookiesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HTTPMonitorCookiesParameters `json:"forProvider"`
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
	InitProvider HTTPMonitorCookiesInitParameters `json:"initProvider,omitempty"`
}

// HTTPMonitorCookiesStatus defines the observed state of HTTPMonitorCookies.
type HTTPMonitorCookiesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HTTPMonitorCookiesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HTTPMonitorCookies is the Schema for the HTTPMonitorCookiess API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type HTTPMonitorCookies struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scope) || (has(self.initProvider) && has(self.initProvider.scope))",message="spec.forProvider.scope is a required parameter"
	Spec   HTTPMonitorCookiesSpec   `json:"spec"`
	Status HTTPMonitorCookiesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HTTPMonitorCookiesList contains a list of HTTPMonitorCookiess
type HTTPMonitorCookiesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTPMonitorCookies `json:"items"`
}

// Repository type metadata.
var (
	HTTPMonitorCookies_Kind             = "HTTPMonitorCookies"
	HTTPMonitorCookies_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HTTPMonitorCookies_Kind}.String()
	HTTPMonitorCookies_KindAPIVersion   = HTTPMonitorCookies_Kind + "." + CRDGroupVersion.String()
	HTTPMonitorCookies_GroupVersionKind = CRDGroupVersion.WithKind(HTTPMonitorCookies_Kind)
)

func init() {
	SchemeBuilder.Register(&HTTPMonitorCookies{}, &HTTPMonitorCookiesList{})
}
