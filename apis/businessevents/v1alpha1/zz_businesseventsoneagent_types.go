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

type BusinessEventsOneagentInitParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Block List, Min: 1, Max: 1) Event meta data (see below for nested schema)
	// Event meta data
	Event []EventInitParameters `json:"event,omitempty" tf:"event,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) Rule name
	// Rule name
	RuleName *string `json:"ruleName,omitempty" tf:"rule_name,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// (Block List, Min: 1, Max: 1) Define conditions to trigger business events from incoming web requests. Triggers are connected by AND logic per capture rule. If you set multiple trigger rules, all of them need to be fulfilled to capture a business event. (see below for nested schema)
	// Define conditions to trigger business events from incoming web requests. Triggers are connected by AND logic per capture rule. If you set multiple trigger rules, all of them need to be fulfilled to capture a business event.
	Triggers []TriggersInitParameters `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type BusinessEventsOneagentObservation struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Block List, Min: 1, Max: 1) Event meta data (see below for nested schema)
	// Event meta data
	Event []EventObservation `json:"event,omitempty" tf:"event,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) Rule name
	// Rule name
	RuleName *string `json:"ruleName,omitempty" tf:"rule_name,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// (Block List, Min: 1, Max: 1) Define conditions to trigger business events from incoming web requests. Triggers are connected by AND logic per capture rule. If you set multiple trigger rules, all of them need to be fulfilled to capture a business event. (see below for nested schema)
	// Define conditions to trigger business events from incoming web requests. Triggers are connected by AND logic per capture rule. If you set multiple trigger rules, all of them need to be fulfilled to capture a business event.
	Triggers []TriggersObservation `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type BusinessEventsOneagentParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Block List, Min: 1, Max: 1) Event meta data (see below for nested schema)
	// Event meta data
	// +kubebuilder:validation:Optional
	Event []EventParameters `json:"event,omitempty" tf:"event,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// +kubebuilder:validation:Optional
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) Rule name
	// Rule name
	// +kubebuilder:validation:Optional
	RuleName *string `json:"ruleName,omitempty" tf:"rule_name,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// (Block List, Min: 1, Max: 1) Define conditions to trigger business events from incoming web requests. Triggers are connected by AND logic per capture rule. If you set multiple trigger rules, all of them need to be fulfilled to capture a business event. (see below for nested schema)
	// Define conditions to trigger business events from incoming web requests. Triggers are connected by AND logic per capture rule. If you set multiple trigger rules, all of them need to be fulfilled to capture a business event.
	// +kubebuilder:validation:Optional
	Triggers []TriggersParameters `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type CategoryInitParameters struct {

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Fixed value
	// Fixed value
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Constant_string, Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Constant_string`, `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

type CategoryObservation struct {

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Fixed value
	// Fixed value
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Constant_string, Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Constant_string`, `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

type CategoryParameters struct {

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Fixed value
	// Fixed value
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Constant_string, Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Constant_string`, `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType" tf:"source_type,omitempty"`
}

type DataInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	EventDataFieldComplex []EventDataFieldComplexInitParameters `json:"eventDataFieldComplex,omitempty" tf:"event_data_field_complex,omitempty"`
}

type DataObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	EventDataFieldComplex []EventDataFieldComplexObservation `json:"eventDataFieldComplex,omitempty" tf:"event_data_field_complex,omitempty"`
}

type DataParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	EventDataFieldComplex []EventDataFieldComplexParameters `json:"eventDataFieldComplex" tf:"event_data_field_complex,omitempty"`
}

type EventDataFieldComplexInitParameters struct {

	// (String) Field name to be added to data.
	// Field name to be added to data.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Fixed value
	// no documentation available
	Source []SourceInitParameters `json:"source,omitempty" tf:"source,omitempty"`
}

type EventDataFieldComplexObservation struct {

	// (String) Field name to be added to data.
	// Field name to be added to data.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Fixed value
	// no documentation available
	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`
}

type EventDataFieldComplexParameters struct {

	// (String) Field name to be added to data.
	// Field name to be added to data.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) Fixed value
	// no documentation available
	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source" tf:"source,omitempty"`
}

type EventInitParameters struct {

	// (Block List, Min: 1, Max: 1) Event category (see below for nested schema)
	// Event category
	Category []CategoryInitParameters `json:"category,omitempty" tf:"category,omitempty"`

	// (Block List, Max: 1) Additional attributes for the business event. (see below for nested schema)
	// Additional attributes for the business event.
	Data []DataInitParameters `json:"data,omitempty" tf:"data,omitempty"`

	// (Block List, Min: 1, Max: 1) Event provider (see below for nested schema)
	// Event provider
	Provider []ProviderInitParameters `json:"provider,omitempty" tf:"provider,omitempty"`

	// (Block List, Min: 1, Max: 1) Event type (see below for nested schema)
	// Event type
	Type []TypeInitParameters `json:"type,omitempty" tf:"type,omitempty"`
}

type EventObservation struct {

	// (Block List, Min: 1, Max: 1) Event category (see below for nested schema)
	// Event category
	Category []CategoryObservation `json:"category,omitempty" tf:"category,omitempty"`

	// (Block List, Max: 1) Additional attributes for the business event. (see below for nested schema)
	// Additional attributes for the business event.
	Data []DataObservation `json:"data,omitempty" tf:"data,omitempty"`

	// (Block List, Min: 1, Max: 1) Event provider (see below for nested schema)
	// Event provider
	Provider []ProviderObservation `json:"provider,omitempty" tf:"provider,omitempty"`

	// (Block List, Min: 1, Max: 1) Event type (see below for nested schema)
	// Event type
	Type []TypeObservation `json:"type,omitempty" tf:"type,omitempty"`
}

type EventParameters struct {

	// (Block List, Min: 1, Max: 1) Event category (see below for nested schema)
	// Event category
	// +kubebuilder:validation:Optional
	Category []CategoryParameters `json:"category" tf:"category,omitempty"`

	// (Block List, Max: 1) Additional attributes for the business event. (see below for nested schema)
	// Additional attributes for the business event.
	// +kubebuilder:validation:Optional
	Data []DataParameters `json:"data,omitempty" tf:"data,omitempty"`

	// (Block List, Min: 1, Max: 1) Event provider (see below for nested schema)
	// Event provider
	// +kubebuilder:validation:Optional
	Provider []ProviderParameters `json:"provider" tf:"provider,omitempty"`

	// (Block List, Min: 1, Max: 1) Event type (see below for nested schema)
	// Event type
	// +kubebuilder:validation:Optional
	Type []TypeParameters `json:"type" tf:"type,omitempty"`
}

type ProviderInitParameters struct {

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Fixed value
	// Fixed value
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Constant_string, Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Constant_string`, `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

type ProviderObservation struct {

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Fixed value
	// Fixed value
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Constant_string, Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Constant_string`, `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

type ProviderParameters struct {

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Fixed value
	// Fixed value
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Constant_string, Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Constant_string`, `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType" tf:"source_type,omitempty"`
}

type SourceInitParameters struct {

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Fixed value
	// Fixed value
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Constant_string, Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Constant_string`, `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

type SourceObservation struct {

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Fixed value
	// Fixed value
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Constant_string, Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Constant_string`, `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

type SourceParameters struct {

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Fixed value
	// Fixed value
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Constant_string, Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Constant_string`, `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType" tf:"source_type,omitempty"`
}

type TriggerInitParameters struct {

	// (Boolean) Case sensitive
	// Case sensitive
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// (String) Fixed value
	// no documentation available
	Source []TriggerSourceInitParameters `json:"source,omitempty" tf:"source,omitempty"`

	// (Block List, Min: 1, Max: 1) Event type (see below for nested schema)
	// Possible Values: `CONTAINS`, `ENDS_WITH`, `EQUALS`, `EXISTS`, `N_CONTAINS`, `N_ENDS_WITH`, `N_EQUALS`, `N_EXISTS`, `N_STARTS_WITH`, `STARTS_WITH`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) no documentation available
	// no documentation available
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TriggerObservation struct {

	// (Boolean) Case sensitive
	// Case sensitive
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// (String) Fixed value
	// no documentation available
	Source []TriggerSourceObservation `json:"source,omitempty" tf:"source,omitempty"`

	// (Block List, Min: 1, Max: 1) Event type (see below for nested schema)
	// Possible Values: `CONTAINS`, `ENDS_WITH`, `EQUALS`, `EXISTS`, `N_CONTAINS`, `N_ENDS_WITH`, `N_EQUALS`, `N_EXISTS`, `N_STARTS_WITH`, `STARTS_WITH`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) no documentation available
	// no documentation available
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TriggerParameters struct {

	// (Boolean) Case sensitive
	// Case sensitive
	// +kubebuilder:validation:Optional
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// (String) Fixed value
	// no documentation available
	// +kubebuilder:validation:Optional
	Source []TriggerSourceParameters `json:"source" tf:"source,omitempty"`

	// (Block List, Min: 1, Max: 1) Event type (see below for nested schema)
	// Possible Values: `CONTAINS`, `ENDS_WITH`, `EQUALS`, `EXISTS`, `N_CONTAINS`, `N_ENDS_WITH`, `N_EQUALS`, `N_EXISTS`, `N_STARTS_WITH`, `STARTS_WITH`
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// (String) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TriggerSourceInitParameters struct {

	// (String) Possible Values: Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type TriggerSourceObservation struct {

	// (String) Possible Values: Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type TriggerSourceParameters struct {

	// (String) Possible Values: Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	// +kubebuilder:validation:Optional
	DataSource *string `json:"dataSource" tf:"data_source,omitempty"`

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type TriggersInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Trigger []TriggerInitParameters `json:"trigger,omitempty" tf:"trigger,omitempty"`
}

type TriggersObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Trigger []TriggerObservation `json:"trigger,omitempty" tf:"trigger,omitempty"`
}

type TriggersParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Trigger []TriggerParameters `json:"trigger" tf:"trigger,omitempty"`
}

type TypeInitParameters struct {

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Fixed value
	// Fixed value
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Constant_string, Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Constant_string`, `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

type TypeObservation struct {

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Fixed value
	// Fixed value
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Constant_string, Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Constant_string`, `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

type TypeParameters struct {

	// (String) See our documentation
	// [See our documentation](https://dt-url.net/ei034bx)
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Fixed value
	// Fixed value
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Constant_string, Request_body, Request_headers, Request_method, Request_parameters, Request_path, Request_url, Response_body, Response_headers, Response_statusCode
	// Possible Values: `Constant_string`, `Request_body`, `Request_headers`, `Request_method`, `Request_parameters`, `Request_path`, `Request_url`, `Response_body`, `Response_headers`, `Response_statusCode`
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType" tf:"source_type,omitempty"`
}

// BusinessEventsOneagentSpec defines the desired state of BusinessEventsOneagent
type BusinessEventsOneagentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BusinessEventsOneagentParameters `json:"forProvider"`
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
	InitProvider BusinessEventsOneagentInitParameters `json:"initProvider,omitempty"`
}

// BusinessEventsOneagentStatus defines the observed state of BusinessEventsOneagent.
type BusinessEventsOneagentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BusinessEventsOneagentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BusinessEventsOneagent is the Schema for the BusinessEventsOneagents API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type BusinessEventsOneagent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.event) || (has(self.initProvider) && has(self.initProvider.event))",message="spec.forProvider.event is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ruleName) || (has(self.initProvider) && has(self.initProvider.ruleName))",message="spec.forProvider.ruleName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.triggers) || (has(self.initProvider) && has(self.initProvider.triggers))",message="spec.forProvider.triggers is a required parameter"
	Spec   BusinessEventsOneagentSpec   `json:"spec"`
	Status BusinessEventsOneagentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BusinessEventsOneagentList contains a list of BusinessEventsOneagents
type BusinessEventsOneagentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BusinessEventsOneagent `json:"items"`
}

// Repository type metadata.
var (
	BusinessEventsOneagent_Kind             = "BusinessEventsOneagent"
	BusinessEventsOneagent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BusinessEventsOneagent_Kind}.String()
	BusinessEventsOneagent_KindAPIVersion   = BusinessEventsOneagent_Kind + "." + CRDGroupVersion.String()
	BusinessEventsOneagent_GroupVersionKind = CRDGroupVersion.WithKind(BusinessEventsOneagent_Kind)
)

func init() {
	SchemeBuilder.Register(&BusinessEventsOneagent{}, &BusinessEventsOneagentList{})
}
