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

type ParameterInitParameters struct {

	// (String) Query parameter name
	// Query parameter name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Query parameter value
	// Query parameter value
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// (Boolean) Query parameter value is undefined
	// Query parameter value is undefined
	ValueIsUndefined *bool `json:"valueIsUndefined,omitempty" tf:"value_is_undefined,omitempty"`
}

type ParameterObservation struct {

	// (String) Query parameter name
	// Query parameter name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Query parameter value
	// Query parameter value
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// (Boolean) Query parameter value is undefined
	// Query parameter value is undefined
	ValueIsUndefined *bool `json:"valueIsUndefined,omitempty" tf:"value_is_undefined,omitempty"`
}

type ParameterParameters struct {

	// (String) Query parameter name
	// Query parameter name
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) Query parameter value
	// Query parameter value
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// (Boolean) Query parameter value is undefined
	// Query parameter value is undefined
	// +kubebuilder:validation:Optional
	ValueIsUndefined *bool `json:"valueIsUndefined,omitempty" tf:"value_is_undefined,omitempty"`
}

type QueryParametersInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Parameter []ParameterInitParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type QueryParametersObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Parameter []ParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type QueryParametersParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter" tf:"parameter,omitempty"`
}

type URLBasedSamplingInitParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Select the scaling factor for the current sampling rate of the system. Possible values: IncreaseCapturing128Times, IncreaseCapturing64Times, IncreaseCapturing32Times, IncreaseCapturing16Times, IncreaseCapturing8Times, IncreaseCapturing4Times, IncreaseCapturing2Times, ReduceCapturingByFactor2, ReduceCapturingByFactor4, ReduceCapturingByFactor8, ReduceCapturingByFactor16, ReduceCapturingByFactor32, ReduceCapturingByFactor64, ReduceCapturingByFactor128
	// Select the scaling factor for the current sampling rate of the system. Possible values: `IncreaseCapturing128Times`, `IncreaseCapturing64Times`, `IncreaseCapturing32Times`, `IncreaseCapturing16Times`, `IncreaseCapturing8Times`, `IncreaseCapturing4Times`, `IncreaseCapturing2Times`, `ReduceCapturingByFactor2`, `ReduceCapturingByFactor4`, `ReduceCapturingByFactor8`, `ReduceCapturingByFactor16`, `ReduceCapturingByFactor32`, `ReduceCapturingByFactor64`, `ReduceCapturingByFactor128`
	Factor *string `json:"factor,omitempty" tf:"factor,omitempty"`

	// (Set of String) Possible values: GET, POST, PUT, DELETE, HEAD, CONNECT, OPTIONS, TRACE, PATCH
	// Possible values: `GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`
	// +listType=set
	HTTPMethod []*string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// (Boolean) The scaling factor for the defined URL will be applied to any HTTP method.
	// The scaling factor for the defined URL will be applied to any HTTP method.
	HTTPMethodAny *bool `json:"httpMethodAny,omitempty" tf:"http_method_any,omitempty"`

	// (Boolean) The matching URLs will always be ignored, also if Adaptive Traffic Management is not active.
	// The matching URLs will always be ignored, also if Adaptive Traffic Management is not active.
	Ignore *bool `json:"ignore,omitempty" tf:"ignore,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) Path of the URL.
	// Path of the URL.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Path comparison condition. Possible values: EQUALS, DOES_NOT_EQUAL, CONTAINS, DOES_NOT_CONTAIN, STARTS_WITH, DOES_NOT_START_WITH, ENDS_WITH, DOES_NOT_END_WITH
	// Path comparison condition. Possible values: `EQUALS`, `DOES_NOT_EQUAL`, `CONTAINS`, `DOES_NOT_CONTAIN`, `STARTS_WITH`, `DOES_NOT_START_WITH`, `ENDS_WITH`, `DOES_NOT_END_WITH`
	PathComparisonType *string `json:"pathComparisonType,omitempty" tf:"path_comparison_type,omitempty"`

	// (Block List, Max: 1) Add URL parameters in any order. All specified parameters must be present in the query of an URL to get a match. (see below for nested schema)
	// Add URL parameters in any order. **All** specified parameters must be present in the query of an URL to get a match.
	QueryParameters []QueryParametersInitParameters `json:"queryParameters,omitempty" tf:"query_parameters,omitempty"`

	// (String) The scope of this setting (PROCESS_GROUP_INSTANCE, PROCESS_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (PROCESS_GROUP_INSTANCE, PROCESS_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type URLBasedSamplingObservation struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Select the scaling factor for the current sampling rate of the system. Possible values: IncreaseCapturing128Times, IncreaseCapturing64Times, IncreaseCapturing32Times, IncreaseCapturing16Times, IncreaseCapturing8Times, IncreaseCapturing4Times, IncreaseCapturing2Times, ReduceCapturingByFactor2, ReduceCapturingByFactor4, ReduceCapturingByFactor8, ReduceCapturingByFactor16, ReduceCapturingByFactor32, ReduceCapturingByFactor64, ReduceCapturingByFactor128
	// Select the scaling factor for the current sampling rate of the system. Possible values: `IncreaseCapturing128Times`, `IncreaseCapturing64Times`, `IncreaseCapturing32Times`, `IncreaseCapturing16Times`, `IncreaseCapturing8Times`, `IncreaseCapturing4Times`, `IncreaseCapturing2Times`, `ReduceCapturingByFactor2`, `ReduceCapturingByFactor4`, `ReduceCapturingByFactor8`, `ReduceCapturingByFactor16`, `ReduceCapturingByFactor32`, `ReduceCapturingByFactor64`, `ReduceCapturingByFactor128`
	Factor *string `json:"factor,omitempty" tf:"factor,omitempty"`

	// (Set of String) Possible values: GET, POST, PUT, DELETE, HEAD, CONNECT, OPTIONS, TRACE, PATCH
	// Possible values: `GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`
	// +listType=set
	HTTPMethod []*string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// (Boolean) The scaling factor for the defined URL will be applied to any HTTP method.
	// The scaling factor for the defined URL will be applied to any HTTP method.
	HTTPMethodAny *bool `json:"httpMethodAny,omitempty" tf:"http_method_any,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) The matching URLs will always be ignored, also if Adaptive Traffic Management is not active.
	// The matching URLs will always be ignored, also if Adaptive Traffic Management is not active.
	Ignore *bool `json:"ignore,omitempty" tf:"ignore,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) Path of the URL.
	// Path of the URL.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Path comparison condition. Possible values: EQUALS, DOES_NOT_EQUAL, CONTAINS, DOES_NOT_CONTAIN, STARTS_WITH, DOES_NOT_START_WITH, ENDS_WITH, DOES_NOT_END_WITH
	// Path comparison condition. Possible values: `EQUALS`, `DOES_NOT_EQUAL`, `CONTAINS`, `DOES_NOT_CONTAIN`, `STARTS_WITH`, `DOES_NOT_START_WITH`, `ENDS_WITH`, `DOES_NOT_END_WITH`
	PathComparisonType *string `json:"pathComparisonType,omitempty" tf:"path_comparison_type,omitempty"`

	// (Block List, Max: 1) Add URL parameters in any order. All specified parameters must be present in the query of an URL to get a match. (see below for nested schema)
	// Add URL parameters in any order. **All** specified parameters must be present in the query of an URL to get a match.
	QueryParameters []QueryParametersObservation `json:"queryParameters,omitempty" tf:"query_parameters,omitempty"`

	// (String) The scope of this setting (PROCESS_GROUP_INSTANCE, PROCESS_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (PROCESS_GROUP_INSTANCE, PROCESS_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type URLBasedSamplingParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Select the scaling factor for the current sampling rate of the system. Possible values: IncreaseCapturing128Times, IncreaseCapturing64Times, IncreaseCapturing32Times, IncreaseCapturing16Times, IncreaseCapturing8Times, IncreaseCapturing4Times, IncreaseCapturing2Times, ReduceCapturingByFactor2, ReduceCapturingByFactor4, ReduceCapturingByFactor8, ReduceCapturingByFactor16, ReduceCapturingByFactor32, ReduceCapturingByFactor64, ReduceCapturingByFactor128
	// Select the scaling factor for the current sampling rate of the system. Possible values: `IncreaseCapturing128Times`, `IncreaseCapturing64Times`, `IncreaseCapturing32Times`, `IncreaseCapturing16Times`, `IncreaseCapturing8Times`, `IncreaseCapturing4Times`, `IncreaseCapturing2Times`, `ReduceCapturingByFactor2`, `ReduceCapturingByFactor4`, `ReduceCapturingByFactor8`, `ReduceCapturingByFactor16`, `ReduceCapturingByFactor32`, `ReduceCapturingByFactor64`, `ReduceCapturingByFactor128`
	// +kubebuilder:validation:Optional
	Factor *string `json:"factor,omitempty" tf:"factor,omitempty"`

	// (Set of String) Possible values: GET, POST, PUT, DELETE, HEAD, CONNECT, OPTIONS, TRACE, PATCH
	// Possible values: `GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`
	// +kubebuilder:validation:Optional
	// +listType=set
	HTTPMethod []*string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// (Boolean) The scaling factor for the defined URL will be applied to any HTTP method.
	// The scaling factor for the defined URL will be applied to any HTTP method.
	// +kubebuilder:validation:Optional
	HTTPMethodAny *bool `json:"httpMethodAny,omitempty" tf:"http_method_any,omitempty"`

	// (Boolean) The matching URLs will always be ignored, also if Adaptive Traffic Management is not active.
	// The matching URLs will always be ignored, also if Adaptive Traffic Management is not active.
	// +kubebuilder:validation:Optional
	Ignore *bool `json:"ignore,omitempty" tf:"ignore,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// +kubebuilder:validation:Optional
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (String) Path of the URL.
	// Path of the URL.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Path comparison condition. Possible values: EQUALS, DOES_NOT_EQUAL, CONTAINS, DOES_NOT_CONTAIN, STARTS_WITH, DOES_NOT_START_WITH, ENDS_WITH, DOES_NOT_END_WITH
	// Path comparison condition. Possible values: `EQUALS`, `DOES_NOT_EQUAL`, `CONTAINS`, `DOES_NOT_CONTAIN`, `STARTS_WITH`, `DOES_NOT_START_WITH`, `ENDS_WITH`, `DOES_NOT_END_WITH`
	// +kubebuilder:validation:Optional
	PathComparisonType *string `json:"pathComparisonType,omitempty" tf:"path_comparison_type,omitempty"`

	// (Block List, Max: 1) Add URL parameters in any order. All specified parameters must be present in the query of an URL to get a match. (see below for nested schema)
	// Add URL parameters in any order. **All** specified parameters must be present in the query of an URL to get a match.
	// +kubebuilder:validation:Optional
	QueryParameters []QueryParametersParameters `json:"queryParameters,omitempty" tf:"query_parameters,omitempty"`

	// (String) The scope of this setting (PROCESS_GROUP_INSTANCE, PROCESS_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (PROCESS_GROUP_INSTANCE, PROCESS_GROUP). Omit this property if you want to cover the whole environment.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// URLBasedSamplingSpec defines the desired state of URLBasedSampling
type URLBasedSamplingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     URLBasedSamplingParameters `json:"forProvider"`
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
	InitProvider URLBasedSamplingInitParameters `json:"initProvider,omitempty"`
}

// URLBasedSamplingStatus defines the observed state of URLBasedSampling.
type URLBasedSamplingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        URLBasedSamplingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// URLBasedSampling is the Schema for the URLBasedSamplings API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type URLBasedSampling struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.httpMethodAny) || (has(self.initProvider) && has(self.initProvider.httpMethodAny))",message="spec.forProvider.httpMethodAny is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ignore) || (has(self.initProvider) && has(self.initProvider.ignore))",message="spec.forProvider.ignore is a required parameter"
	Spec   URLBasedSamplingSpec   `json:"spec"`
	Status URLBasedSamplingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// URLBasedSamplingList contains a list of URLBasedSamplings
type URLBasedSamplingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []URLBasedSampling `json:"items"`
}

// Repository type metadata.
var (
	URLBasedSampling_Kind             = "URLBasedSampling"
	URLBasedSampling_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: URLBasedSampling_Kind}.String()
	URLBasedSampling_KindAPIVersion   = URLBasedSampling_Kind + "." + CRDGroupVersion.String()
	URLBasedSampling_GroupVersionKind = CRDGroupVersion.WithKind(URLBasedSampling_Kind)
)

func init() {
	SchemeBuilder.Register(&URLBasedSampling{}, &URLBasedSamplingList{})
}