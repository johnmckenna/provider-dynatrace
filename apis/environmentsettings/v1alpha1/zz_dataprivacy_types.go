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

type DataCollectionInitParameters struct {

	// collection and opt-in mode enabled, Real User Monitoring data isn't captured until dtrum.enable() is called for specific user sessions.
	// With [Data-collection and opt-in mode](https://dt-url.net/7l3p0p3h) enabled, Real User Monitoring data isn't captured until dtrum.enable() is called for specific user sessions.
	OptInModeEnabled *bool `json:"optInModeEnabled,omitempty" tf:"opt_in_mode_enabled,omitempty"`
}

type DataCollectionObservation struct {

	// collection and opt-in mode enabled, Real User Monitoring data isn't captured until dtrum.enable() is called for specific user sessions.
	// With [Data-collection and opt-in mode](https://dt-url.net/7l3p0p3h) enabled, Real User Monitoring data isn't captured until dtrum.enable() is called for specific user sessions.
	OptInModeEnabled *bool `json:"optInModeEnabled,omitempty" tf:"opt_in_mode_enabled,omitempty"`
}

type DataCollectionParameters struct {

	// collection and opt-in mode enabled, Real User Monitoring data isn't captured until dtrum.enable() is called for specific user sessions.
	// With [Data-collection and opt-in mode](https://dt-url.net/7l3p0p3h) enabled, Real User Monitoring data isn't captured until dtrum.enable() is called for specific user sessions.
	// +kubebuilder:validation:Optional
	OptInModeEnabled *bool `json:"optInModeEnabled" tf:"opt_in_mode_enabled,omitempty"`
}

type DataPrivacyInitParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// in mode. (see below for nested schema)
	// To provide your end users with the ability to decide for themselves if their activities should be tracked to measure application performance and usage, enable opt-in mode.
	DataCollection []DataCollectionInitParameters `json:"dataCollection,omitempty" tf:"data_collection,omitempty"`

	// (Block List, Min: 1, Max: 1) Most modern web browsers have a privacy feature called "Do Not Track" that individual users may have enabled on their devices. Customize how Dynatrace should behave when it encounters this setting. (see below for nested schema)
	// Most modern web browsers have a privacy feature called ["Do Not Track"](https://dt-url.net/sb3n0pnl) that individual users may have enabled on their devices. Customize how Dynatrace should behave when it encounters this setting.
	DoNotTrack []DoNotTrackInitParameters `json:"doNotTrack,omitempty" tf:"do_not_track,omitempty"`

	// (Block List, Min: 1, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	Masking []MaskingInitParameters `json:"masking,omitempty" tf:"masking,omitempty"`

	// (Block List, Min: 1, Max: 1) User tracking (see below for nested schema)
	// User tracking
	UserTracking []UserTrackingInitParameters `json:"userTracking,omitempty" tf:"user_tracking,omitempty"`
}

type DataPrivacyObservation struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// in mode. (see below for nested schema)
	// To provide your end users with the ability to decide for themselves if their activities should be tracked to measure application performance and usage, enable opt-in mode.
	DataCollection []DataCollectionObservation `json:"dataCollection,omitempty" tf:"data_collection,omitempty"`

	// (Block List, Min: 1, Max: 1) Most modern web browsers have a privacy feature called "Do Not Track" that individual users may have enabled on their devices. Customize how Dynatrace should behave when it encounters this setting. (see below for nested schema)
	// Most modern web browsers have a privacy feature called ["Do Not Track"](https://dt-url.net/sb3n0pnl) that individual users may have enabled on their devices. Customize how Dynatrace should behave when it encounters this setting.
	DoNotTrack []DoNotTrackObservation `json:"doNotTrack,omitempty" tf:"do_not_track,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Min: 1, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	Masking []MaskingObservation `json:"masking,omitempty" tf:"masking,omitempty"`

	// (Block List, Min: 1, Max: 1) User tracking (see below for nested schema)
	// User tracking
	UserTracking []UserTrackingObservation `json:"userTracking,omitempty" tf:"user_tracking,omitempty"`
}

type DataPrivacyParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// in mode. (see below for nested schema)
	// To provide your end users with the ability to decide for themselves if their activities should be tracked to measure application performance and usage, enable opt-in mode.
	// +kubebuilder:validation:Optional
	DataCollection []DataCollectionParameters `json:"dataCollection,omitempty" tf:"data_collection,omitempty"`

	// (Block List, Min: 1, Max: 1) Most modern web browsers have a privacy feature called "Do Not Track" that individual users may have enabled on their devices. Customize how Dynatrace should behave when it encounters this setting. (see below for nested schema)
	// Most modern web browsers have a privacy feature called ["Do Not Track"](https://dt-url.net/sb3n0pnl) that individual users may have enabled on their devices. Customize how Dynatrace should behave when it encounters this setting.
	// +kubebuilder:validation:Optional
	DoNotTrack []DoNotTrackParameters `json:"doNotTrack,omitempty" tf:"do_not_track,omitempty"`

	// (Block List, Min: 1, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	// +kubebuilder:validation:Optional
	Masking []MaskingParameters `json:"masking,omitempty" tf:"masking,omitempty"`

	// (Block List, Min: 1, Max: 1) User tracking (see below for nested schema)
	// User tracking
	// +kubebuilder:validation:Optional
	UserTracking []UserTrackingParameters `json:"userTracking,omitempty" tf:"user_tracking,omitempty"`
}

type DoNotTrackInitParameters struct {

	// (Boolean) Comply with "Do Not Track" browser settings
	// Comply with "Do Not Track" browser settings
	ComplyWithDoNotTrack *bool `json:"complyWithDoNotTrack,omitempty" tf:"comply_with_do_not_track,omitempty"`

	// (Block List, Min: 1, Max: 1) Most modern web browsers have a privacy feature called "Do Not Track" that individual users may have enabled on their devices. Customize how Dynatrace should behave when it encounters this setting. (see below for nested schema)
	// Possible Values: `Anonymous`, `Disable_rum`
	DoNotTrack *string `json:"doNotTrack,omitempty" tf:"do_not_track,omitempty"`
}

type DoNotTrackObservation struct {

	// (Boolean) Comply with "Do Not Track" browser settings
	// Comply with "Do Not Track" browser settings
	ComplyWithDoNotTrack *bool `json:"complyWithDoNotTrack,omitempty" tf:"comply_with_do_not_track,omitempty"`

	// (Block List, Min: 1, Max: 1) Most modern web browsers have a privacy feature called "Do Not Track" that individual users may have enabled on their devices. Customize how Dynatrace should behave when it encounters this setting. (see below for nested schema)
	// Possible Values: `Anonymous`, `Disable_rum`
	DoNotTrack *string `json:"doNotTrack,omitempty" tf:"do_not_track,omitempty"`
}

type DoNotTrackParameters struct {

	// (Boolean) Comply with "Do Not Track" browser settings
	// Comply with "Do Not Track" browser settings
	// +kubebuilder:validation:Optional
	ComplyWithDoNotTrack *bool `json:"complyWithDoNotTrack" tf:"comply_with_do_not_track,omitempty"`

	// (Block List, Min: 1, Max: 1) Most modern web browsers have a privacy feature called "Do Not Track" that individual users may have enabled on their devices. Customize how Dynatrace should behave when it encounters this setting. (see below for nested schema)
	// Possible Values: `Anonymous`, `Disable_rum`
	// +kubebuilder:validation:Optional
	DoNotTrack *string `json:"doNotTrack,omitempty" tf:"do_not_track,omitempty"`
}

type MaskingInitParameters struct {

	// (String, Deprecated) Possible Values: All, Public
	// Possible Values: `All`, `Public`
	IPAddressMasking *string `json:"ipAddressMasking,omitempty" tf:"ip_address_masking,omitempty"`

	// users to determine the regions from which they access your application. To learn more, visit Mask IPs and GPS coordinates.. Dynatrace also captures GPS data from mobile apps that provide their users with the option of sharing geolocation data. On the server side, Dynatrace captures IP addresses to enable detailed troubleshooting for Dynatrace service calls.
	// Dynatrace captures the IP addresses of your end-users to determine the regions from which they access your application. To learn more, visit [Mask IPs and GPS coordinates](https://dt-url.net/mask-end-users-ip-addresses).. Dynatrace also captures GPS data from mobile apps that provide their users with the option of sharing geolocation data. On the server side, Dynatrace captures IP addresses to enable detailed troubleshooting for Dynatrace service calls.
	//
	// Once enabled, IP address masking sets the last octet of monitored IPv4 addresses and the last 80 bits of IPv6 addresses to zeroes. GPS coordinates are rounded up to 1 decimal place (~10 km). This masking occurs in memory. Full IP addresses are never written to disk. Location lookups are made using anonymized IP addresses and GPS coordinates.
	IPAddressMaskingEnabled *bool `json:"ipAddressMaskingEnabled,omitempty" tf:"ip_address_masking_enabled,omitempty"`

	// side to enable detailed performance analysis of your applications. For complete details, visit Mask personal data in URIs.. URIs and request headers contain personal data. When this setting is enabled, Dynatrace automatically detects UUIDs, credit card numbers, email addresses, IP addresses, and other IDs and replaces those values with placeholders. The personal data is then masked in PurePath analysis, error analysis, user action naming for RUM, and elsewhere in Dynatrace.
	// Dynatrace captures the URIs and request headers sent from desktop and mobile browsers. Dynatrace also captures full URIs on the server-side to enable detailed performance analysis of your applications. For complete details, visit [Mask personal data in URIs](https://dt-url.net/mask-personal-data-in-URIs).. URIs and request headers contain personal data. When this setting is enabled, Dynatrace automatically detects UUIDs, credit card numbers, email addresses, IP addresses, and other IDs and replaces those values with placeholders. The personal data is then masked in PurePath analysis, error analysis, user action naming for RUM, and elsewhere in Dynatrace.
	PersonalDataURIMaskingEnabled *bool `json:"personalDataUriMaskingEnabled,omitempty" tf:"personal_data_uri_masking_enabled,omitempty"`

	// (Boolean) When Dynatrace detects a user action that triggers a page load or an AJAX/XHR action. To learn more about masking user actions, visit Mask user actions.. When Dynatrace detects a user action that triggers a page load or an AJAX/XHR action, it constructs a name for the user action based on:
	// When Dynatrace detects a user action that triggers a page load or an AJAX/XHR action. To learn more about masking user actions, visit [Mask user actions](https://dt-url.net/mask-user-action).. When Dynatrace detects a user action that triggers a page load or an AJAX/XHR action, it constructs a name for the user action based on:
	//
	// - User event type (click on..., loading of page..., or keypress on...)
	// - Title, caption, label, value, ID, className, or other available property of the related HTML element (for example, an image, button, checkbox, or text input field).
	//
	// In most instances, the default approach to user-action naming works well, resulting in user-action names such as:
	//
	// - click on "Search" on page /search.html
	// - keypress on "Feedback" on page /contact.html
	// - touch on "Homescreen" of page /list.jsf
	//
	// In rare circumstances, confidential data (for example, email addresses, usernames, or account numbers) can be unintentionally included in user action names because the confidential data itself is included in an HTML element label, attribute, or other value (for example, click on "my Account Number: 1231231"...). If such confidential data appears in your application's user action names, enable the Mask user action names setting. This setting replaces specific HTML element names and values with generic HTML element names. With user-action name masking enabled, the user action names listed above appear as:
	//
	// - click on INPUT on page /search.html
	// - keypress on TEXTAREA on page /contact.html
	// - touch on DIV of page /list.jsf
	UserActionMaskingEnabled *bool `json:"userActionMaskingEnabled,omitempty" tf:"user_action_masking_enabled,omitempty"`
}

type MaskingObservation struct {

	// (String, Deprecated) Possible Values: All, Public
	// Possible Values: `All`, `Public`
	IPAddressMasking *string `json:"ipAddressMasking,omitempty" tf:"ip_address_masking,omitempty"`

	// users to determine the regions from which they access your application. To learn more, visit Mask IPs and GPS coordinates.. Dynatrace also captures GPS data from mobile apps that provide their users with the option of sharing geolocation data. On the server side, Dynatrace captures IP addresses to enable detailed troubleshooting for Dynatrace service calls.
	// Dynatrace captures the IP addresses of your end-users to determine the regions from which they access your application. To learn more, visit [Mask IPs and GPS coordinates](https://dt-url.net/mask-end-users-ip-addresses).. Dynatrace also captures GPS data from mobile apps that provide their users with the option of sharing geolocation data. On the server side, Dynatrace captures IP addresses to enable detailed troubleshooting for Dynatrace service calls.
	//
	// Once enabled, IP address masking sets the last octet of monitored IPv4 addresses and the last 80 bits of IPv6 addresses to zeroes. GPS coordinates are rounded up to 1 decimal place (~10 km). This masking occurs in memory. Full IP addresses are never written to disk. Location lookups are made using anonymized IP addresses and GPS coordinates.
	IPAddressMaskingEnabled *bool `json:"ipAddressMaskingEnabled,omitempty" tf:"ip_address_masking_enabled,omitempty"`

	// side to enable detailed performance analysis of your applications. For complete details, visit Mask personal data in URIs.. URIs and request headers contain personal data. When this setting is enabled, Dynatrace automatically detects UUIDs, credit card numbers, email addresses, IP addresses, and other IDs and replaces those values with placeholders. The personal data is then masked in PurePath analysis, error analysis, user action naming for RUM, and elsewhere in Dynatrace.
	// Dynatrace captures the URIs and request headers sent from desktop and mobile browsers. Dynatrace also captures full URIs on the server-side to enable detailed performance analysis of your applications. For complete details, visit [Mask personal data in URIs](https://dt-url.net/mask-personal-data-in-URIs).. URIs and request headers contain personal data. When this setting is enabled, Dynatrace automatically detects UUIDs, credit card numbers, email addresses, IP addresses, and other IDs and replaces those values with placeholders. The personal data is then masked in PurePath analysis, error analysis, user action naming for RUM, and elsewhere in Dynatrace.
	PersonalDataURIMaskingEnabled *bool `json:"personalDataUriMaskingEnabled,omitempty" tf:"personal_data_uri_masking_enabled,omitempty"`

	// (Boolean) When Dynatrace detects a user action that triggers a page load or an AJAX/XHR action. To learn more about masking user actions, visit Mask user actions.. When Dynatrace detects a user action that triggers a page load or an AJAX/XHR action, it constructs a name for the user action based on:
	// When Dynatrace detects a user action that triggers a page load or an AJAX/XHR action. To learn more about masking user actions, visit [Mask user actions](https://dt-url.net/mask-user-action).. When Dynatrace detects a user action that triggers a page load or an AJAX/XHR action, it constructs a name for the user action based on:
	//
	// - User event type (click on..., loading of page..., or keypress on...)
	// - Title, caption, label, value, ID, className, or other available property of the related HTML element (for example, an image, button, checkbox, or text input field).
	//
	// In most instances, the default approach to user-action naming works well, resulting in user-action names such as:
	//
	// - click on "Search" on page /search.html
	// - keypress on "Feedback" on page /contact.html
	// - touch on "Homescreen" of page /list.jsf
	//
	// In rare circumstances, confidential data (for example, email addresses, usernames, or account numbers) can be unintentionally included in user action names because the confidential data itself is included in an HTML element label, attribute, or other value (for example, click on "my Account Number: 1231231"...). If such confidential data appears in your application's user action names, enable the Mask user action names setting. This setting replaces specific HTML element names and values with generic HTML element names. With user-action name masking enabled, the user action names listed above appear as:
	//
	// - click on INPUT on page /search.html
	// - keypress on TEXTAREA on page /contact.html
	// - touch on DIV of page /list.jsf
	UserActionMaskingEnabled *bool `json:"userActionMaskingEnabled,omitempty" tf:"user_action_masking_enabled,omitempty"`
}

type MaskingParameters struct {

	// (String, Deprecated) Possible Values: All, Public
	// Possible Values: `All`, `Public`
	// +kubebuilder:validation:Optional
	IPAddressMasking *string `json:"ipAddressMasking,omitempty" tf:"ip_address_masking,omitempty"`

	// users to determine the regions from which they access your application. To learn more, visit Mask IPs and GPS coordinates.. Dynatrace also captures GPS data from mobile apps that provide their users with the option of sharing geolocation data. On the server side, Dynatrace captures IP addresses to enable detailed troubleshooting for Dynatrace service calls.
	// Dynatrace captures the IP addresses of your end-users to determine the regions from which they access your application. To learn more, visit [Mask IPs and GPS coordinates](https://dt-url.net/mask-end-users-ip-addresses).. Dynatrace also captures GPS data from mobile apps that provide their users with the option of sharing geolocation data. On the server side, Dynatrace captures IP addresses to enable detailed troubleshooting for Dynatrace service calls.
	//
	// Once enabled, IP address masking sets the last octet of monitored IPv4 addresses and the last 80 bits of IPv6 addresses to zeroes. GPS coordinates are rounded up to 1 decimal place (~10 km). This masking occurs in memory. Full IP addresses are never written to disk. Location lookups are made using anonymized IP addresses and GPS coordinates.
	// +kubebuilder:validation:Optional
	IPAddressMaskingEnabled *bool `json:"ipAddressMaskingEnabled,omitempty" tf:"ip_address_masking_enabled,omitempty"`

	// side to enable detailed performance analysis of your applications. For complete details, visit Mask personal data in URIs.. URIs and request headers contain personal data. When this setting is enabled, Dynatrace automatically detects UUIDs, credit card numbers, email addresses, IP addresses, and other IDs and replaces those values with placeholders. The personal data is then masked in PurePath analysis, error analysis, user action naming for RUM, and elsewhere in Dynatrace.
	// Dynatrace captures the URIs and request headers sent from desktop and mobile browsers. Dynatrace also captures full URIs on the server-side to enable detailed performance analysis of your applications. For complete details, visit [Mask personal data in URIs](https://dt-url.net/mask-personal-data-in-URIs).. URIs and request headers contain personal data. When this setting is enabled, Dynatrace automatically detects UUIDs, credit card numbers, email addresses, IP addresses, and other IDs and replaces those values with placeholders. The personal data is then masked in PurePath analysis, error analysis, user action naming for RUM, and elsewhere in Dynatrace.
	// +kubebuilder:validation:Optional
	PersonalDataURIMaskingEnabled *bool `json:"personalDataUriMaskingEnabled" tf:"personal_data_uri_masking_enabled,omitempty"`

	// (Boolean) When Dynatrace detects a user action that triggers a page load or an AJAX/XHR action. To learn more about masking user actions, visit Mask user actions.. When Dynatrace detects a user action that triggers a page load or an AJAX/XHR action, it constructs a name for the user action based on:
	// When Dynatrace detects a user action that triggers a page load or an AJAX/XHR action. To learn more about masking user actions, visit [Mask user actions](https://dt-url.net/mask-user-action).. When Dynatrace detects a user action that triggers a page load or an AJAX/XHR action, it constructs a name for the user action based on:
	//
	// - User event type (click on..., loading of page..., or keypress on...)
	// - Title, caption, label, value, ID, className, or other available property of the related HTML element (for example, an image, button, checkbox, or text input field).
	//
	// In most instances, the default approach to user-action naming works well, resulting in user-action names such as:
	//
	// - click on "Search" on page /search.html
	// - keypress on "Feedback" on page /contact.html
	// - touch on "Homescreen" of page /list.jsf
	//
	// In rare circumstances, confidential data (for example, email addresses, usernames, or account numbers) can be unintentionally included in user action names because the confidential data itself is included in an HTML element label, attribute, or other value (for example, click on "my Account Number: 1231231"...). If such confidential data appears in your application's user action names, enable the Mask user action names setting. This setting replaces specific HTML element names and values with generic HTML element names. With user-action name masking enabled, the user action names listed above appear as:
	//
	// - click on INPUT on page /search.html
	// - keypress on TEXTAREA on page /contact.html
	// - touch on DIV of page /list.jsf
	// +kubebuilder:validation:Optional
	UserActionMaskingEnabled *bool `json:"userActionMaskingEnabled" tf:"user_action_masking_enabled,omitempty"`
}

type UserTrackingInitParameters struct {

	// user devices to identify returning users.
	// When enabled, Dynatrace places a [persistent cookie](https://dt-url.net/313o0p4n) on all end-user devices to identify returning users.
	PersistentCookieEnabled *bool `json:"persistentCookieEnabled,omitempty" tf:"persistent_cookie_enabled,omitempty"`
}

type UserTrackingObservation struct {

	// user devices to identify returning users.
	// When enabled, Dynatrace places a [persistent cookie](https://dt-url.net/313o0p4n) on all end-user devices to identify returning users.
	PersistentCookieEnabled *bool `json:"persistentCookieEnabled,omitempty" tf:"persistent_cookie_enabled,omitempty"`
}

type UserTrackingParameters struct {

	// user devices to identify returning users.
	// When enabled, Dynatrace places a [persistent cookie](https://dt-url.net/313o0p4n) on all end-user devices to identify returning users.
	// +kubebuilder:validation:Optional
	PersistentCookieEnabled *bool `json:"persistentCookieEnabled" tf:"persistent_cookie_enabled,omitempty"`
}

// DataPrivacySpec defines the desired state of DataPrivacy
type DataPrivacySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataPrivacyParameters `json:"forProvider"`
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
	InitProvider DataPrivacyInitParameters `json:"initProvider,omitempty"`
}

// DataPrivacyStatus defines the observed state of DataPrivacy.
type DataPrivacyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataPrivacyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DataPrivacy is the Schema for the DataPrivacys API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type DataPrivacy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataCollection) || (has(self.initProvider) && has(self.initProvider.dataCollection))",message="spec.forProvider.dataCollection is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.doNotTrack) || (has(self.initProvider) && has(self.initProvider.doNotTrack))",message="spec.forProvider.doNotTrack is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.masking) || (has(self.initProvider) && has(self.initProvider.masking))",message="spec.forProvider.masking is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userTracking) || (has(self.initProvider) && has(self.initProvider.userTracking))",message="spec.forProvider.userTracking is a required parameter"
	Spec   DataPrivacySpec   `json:"spec"`
	Status DataPrivacyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataPrivacyList contains a list of DataPrivacys
type DataPrivacyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataPrivacy `json:"items"`
}

// Repository type metadata.
var (
	DataPrivacy_Kind             = "DataPrivacy"
	DataPrivacy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataPrivacy_Kind}.String()
	DataPrivacy_KindAPIVersion   = DataPrivacy_Kind + "." + CRDGroupVersion.String()
	DataPrivacy_GroupVersionKind = CRDGroupVersion.WithKind(DataPrivacy_Kind)
)

func init() {
	SchemeBuilder.Register(&DataPrivacy{}, &DataPrivacyList{})
}
