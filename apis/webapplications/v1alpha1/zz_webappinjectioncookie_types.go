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

type WebAppInjectionCookieInitParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// level domains). Before typing a domain name here, confirm that the domain will accept cookies from your browser. For details, see the list of forbidden top-level domains.
	// Specify an alternative domain for cookies set by Dynatrace. Keep in mind that your browser may not allow placement of cookies on certain domains (for example, top-level domains). Before typing a domain name here, confirm that the domain will accept cookies from your browser. For details, see the list of [forbidden top-level domains](https://dt-url.net/9n6b0pfz).
	CookiePlacementDomain *string `json:"cookiePlacementDomain,omitempty" tf:"cookie_placement_domain,omitempty"`

	// (String) Possible Values: LAX, NONE, NOTSET, STRICT
	// Possible Values: `LAX`, `NONE`, `NOTSET`, `STRICT`
	SameSiteCookieAttribute *string `json:"sameSiteCookieAttribute,omitempty" tf:"same_site_cookie_attribute,omitempty"`

	// compliance security scanners. Be aware that with this setting enabled Dynatrace correlation of user actions with server-side web requests is only possible over SSL connections.
	// If your application is only accessible via SSL, you can add the Secure attribute to all cookies set by Dynatrace. This setting prevents the display of warnings from PCI-compliance security scanners. Be aware that with this setting enabled Dynatrace correlation of user actions with server-side web requests is only possible over SSL connections.
	UseSecureCookieAttribute *bool `json:"useSecureCookieAttribute,omitempty" tf:"use_secure_cookie_attribute,omitempty"`
}

type WebAppInjectionCookieObservation struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// level domains). Before typing a domain name here, confirm that the domain will accept cookies from your browser. For details, see the list of forbidden top-level domains.
	// Specify an alternative domain for cookies set by Dynatrace. Keep in mind that your browser may not allow placement of cookies on certain domains (for example, top-level domains). Before typing a domain name here, confirm that the domain will accept cookies from your browser. For details, see the list of [forbidden top-level domains](https://dt-url.net/9n6b0pfz).
	CookiePlacementDomain *string `json:"cookiePlacementDomain,omitempty" tf:"cookie_placement_domain,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Possible Values: LAX, NONE, NOTSET, STRICT
	// Possible Values: `LAX`, `NONE`, `NOTSET`, `STRICT`
	SameSiteCookieAttribute *string `json:"sameSiteCookieAttribute,omitempty" tf:"same_site_cookie_attribute,omitempty"`

	// compliance security scanners. Be aware that with this setting enabled Dynatrace correlation of user actions with server-side web requests is only possible over SSL connections.
	// If your application is only accessible via SSL, you can add the Secure attribute to all cookies set by Dynatrace. This setting prevents the display of warnings from PCI-compliance security scanners. Be aware that with this setting enabled Dynatrace correlation of user actions with server-side web requests is only possible over SSL connections.
	UseSecureCookieAttribute *bool `json:"useSecureCookieAttribute,omitempty" tf:"use_secure_cookie_attribute,omitempty"`
}

type WebAppInjectionCookieParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// level domains). Before typing a domain name here, confirm that the domain will accept cookies from your browser. For details, see the list of forbidden top-level domains.
	// Specify an alternative domain for cookies set by Dynatrace. Keep in mind that your browser may not allow placement of cookies on certain domains (for example, top-level domains). Before typing a domain name here, confirm that the domain will accept cookies from your browser. For details, see the list of [forbidden top-level domains](https://dt-url.net/9n6b0pfz).
	// +kubebuilder:validation:Optional
	CookiePlacementDomain *string `json:"cookiePlacementDomain,omitempty" tf:"cookie_placement_domain,omitempty"`

	// (String) Possible Values: LAX, NONE, NOTSET, STRICT
	// Possible Values: `LAX`, `NONE`, `NOTSET`, `STRICT`
	// +kubebuilder:validation:Optional
	SameSiteCookieAttribute *string `json:"sameSiteCookieAttribute,omitempty" tf:"same_site_cookie_attribute,omitempty"`

	// compliance security scanners. Be aware that with this setting enabled Dynatrace correlation of user actions with server-side web requests is only possible over SSL connections.
	// If your application is only accessible via SSL, you can add the Secure attribute to all cookies set by Dynatrace. This setting prevents the display of warnings from PCI-compliance security scanners. Be aware that with this setting enabled Dynatrace correlation of user actions with server-side web requests is only possible over SSL connections.
	// +kubebuilder:validation:Optional
	UseSecureCookieAttribute *bool `json:"useSecureCookieAttribute,omitempty" tf:"use_secure_cookie_attribute,omitempty"`
}

// WebAppInjectionCookieSpec defines the desired state of WebAppInjectionCookie
type WebAppInjectionCookieSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebAppInjectionCookieParameters `json:"forProvider"`
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
	InitProvider WebAppInjectionCookieInitParameters `json:"initProvider,omitempty"`
}

// WebAppInjectionCookieStatus defines the observed state of WebAppInjectionCookie.
type WebAppInjectionCookieStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebAppInjectionCookieObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// WebAppInjectionCookie is the Schema for the WebAppInjectionCookies API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type WebAppInjectionCookie struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.applicationId) || (has(self.initProvider) && has(self.initProvider.applicationId))",message="spec.forProvider.applicationId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sameSiteCookieAttribute) || (has(self.initProvider) && has(self.initProvider.sameSiteCookieAttribute))",message="spec.forProvider.sameSiteCookieAttribute is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.useSecureCookieAttribute) || (has(self.initProvider) && has(self.initProvider.useSecureCookieAttribute))",message="spec.forProvider.useSecureCookieAttribute is a required parameter"
	Spec   WebAppInjectionCookieSpec   `json:"spec"`
	Status WebAppInjectionCookieStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebAppInjectionCookieList contains a list of WebAppInjectionCookies
type WebAppInjectionCookieList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebAppInjectionCookie `json:"items"`
}

// Repository type metadata.
var (
	WebAppInjectionCookie_Kind             = "WebAppInjectionCookie"
	WebAppInjectionCookie_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebAppInjectionCookie_Kind}.String()
	WebAppInjectionCookie_KindAPIVersion   = WebAppInjectionCookie_Kind + "." + CRDGroupVersion.String()
	WebAppInjectionCookie_GroupVersionKind = CRDGroupVersion.WithKind(WebAppInjectionCookie_Kind)
)

func init() {
	SchemeBuilder.Register(&WebAppInjectionCookie{}, &WebAppInjectionCookieList{})
}