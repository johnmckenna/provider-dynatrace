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

type DimensionInitParameters struct {

	// (String) Display name
	// Display name
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) Dimension key
	// Dimension key
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type DimensionObservation struct {

	// (String) Display name
	// Display name
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) Dimension key
	// Dimension key
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type DimensionParameters struct {

	// (String) Display name
	// Display name
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) Dimension key
	// Dimension key
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`
}

type DimensionsInitParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	Dimension []DimensionInitParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`
}

type DimensionsObservation struct {

	// (Block List, Min: 1) (see below for nested schema)
	Dimension []DimensionObservation `json:"dimension,omitempty" tf:"dimension,omitempty"`
}

type DimensionsParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Dimension []DimensionParameters `json:"dimension" tf:"dimension,omitempty"`
}

type MetricMetadataInitParameters struct {

	// (String) Description
	// Description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block List, Max: 1) Define metadata per metric dimension. (see below for nested schema)
	// Define metadata per metric dimension.
	Dimensions []DimensionsInitParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// (String) Display name
	// Display name
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) The scope of this setting (metric)
	// The scope of this setting (metric)
	MetricID *string `json:"metricId,omitempty" tf:"metric_id,omitempty"`

	// (Block List, Max: 1) Metric properties (see below for nested schema)
	// Metric properties
	MetricProperties []MetricPropertiesInitParameters `json:"metricProperties,omitempty" tf:"metric_properties,omitempty"`

	// (String) Specifies which entity dimension should be used as the primary dimension. The property can only be configured for metrics ingested with the Metrics API.
	// Specifies which entity dimension should be used as the primary dimension. The property can only be configured for metrics ingested with the Metrics API.
	SourceEntityType *string `json:"sourceEntityType,omitempty" tf:"source_entity_type,omitempty"`

	// (Set of String) Tags
	// Tags
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (String) Unit
	// Unit
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// (String) The raw value is stored in bits or bytes. The user interface can display it in these numeral systems:
	// The raw value is stored in bits or bytes. The user interface can display it in these numeral systems:
	//
	// Binary: 1 MiB = 1024 KiB = 1,048,576 bytes
	//
	// Decimal: 1 MB = 1000 kB = 1,000,000 bytes
	//
	// If not set, the decimal system is used.
	UnitDisplayFormat *string `json:"unitDisplayFormat,omitempty" tf:"unit_display_format,omitempty"`
}

type MetricMetadataObservation struct {

	// (String) Description
	// Description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block List, Max: 1) Define metadata per metric dimension. (see below for nested schema)
	// Define metadata per metric dimension.
	Dimensions []DimensionsObservation `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// (String) Display name
	// Display name
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The scope of this setting (metric)
	// The scope of this setting (metric)
	MetricID *string `json:"metricId,omitempty" tf:"metric_id,omitempty"`

	// (Block List, Max: 1) Metric properties (see below for nested schema)
	// Metric properties
	MetricProperties []MetricPropertiesObservation `json:"metricProperties,omitempty" tf:"metric_properties,omitempty"`

	// (String) Specifies which entity dimension should be used as the primary dimension. The property can only be configured for metrics ingested with the Metrics API.
	// Specifies which entity dimension should be used as the primary dimension. The property can only be configured for metrics ingested with the Metrics API.
	SourceEntityType *string `json:"sourceEntityType,omitempty" tf:"source_entity_type,omitempty"`

	// (Set of String) Tags
	// Tags
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (String) Unit
	// Unit
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// (String) The raw value is stored in bits or bytes. The user interface can display it in these numeral systems:
	// The raw value is stored in bits or bytes. The user interface can display it in these numeral systems:
	//
	// Binary: 1 MiB = 1024 KiB = 1,048,576 bytes
	//
	// Decimal: 1 MB = 1000 kB = 1,000,000 bytes
	//
	// If not set, the decimal system is used.
	UnitDisplayFormat *string `json:"unitDisplayFormat,omitempty" tf:"unit_display_format,omitempty"`
}

type MetricMetadataParameters struct {

	// (String) Description
	// Description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block List, Max: 1) Define metadata per metric dimension. (see below for nested schema)
	// Define metadata per metric dimension.
	// +kubebuilder:validation:Optional
	Dimensions []DimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// (String) Display name
	// Display name
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) The scope of this setting (metric)
	// The scope of this setting (metric)
	// +kubebuilder:validation:Optional
	MetricID *string `json:"metricId,omitempty" tf:"metric_id,omitempty"`

	// (Block List, Max: 1) Metric properties (see below for nested schema)
	// Metric properties
	// +kubebuilder:validation:Optional
	MetricProperties []MetricPropertiesParameters `json:"metricProperties,omitempty" tf:"metric_properties,omitempty"`

	// (String) Specifies which entity dimension should be used as the primary dimension. The property can only be configured for metrics ingested with the Metrics API.
	// Specifies which entity dimension should be used as the primary dimension. The property can only be configured for metrics ingested with the Metrics API.
	// +kubebuilder:validation:Optional
	SourceEntityType *string `json:"sourceEntityType,omitempty" tf:"source_entity_type,omitempty"`

	// (Set of String) Tags
	// Tags
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (String) Unit
	// Unit
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// (String) The raw value is stored in bits or bytes. The user interface can display it in these numeral systems:
	// The raw value is stored in bits or bytes. The user interface can display it in these numeral systems:
	//
	// Binary: 1 MiB = 1024 KiB = 1,048,576 bytes
	//
	// Decimal: 1 MB = 1000 kB = 1,000,000 bytes
	//
	// If not set, the decimal system is used.
	// +kubebuilder:validation:Optional
	UnitDisplayFormat *string `json:"unitDisplayFormat,omitempty" tf:"unit_display_format,omitempty"`
}

type MetricPropertiesInitParameters struct {

	// (Boolean) Whether (true or false) the metric is relevant to a problem's impact.
	// Whether (true or false) the metric is relevant to a problem's impact.
	//
	// An impact-relevant metric is highly dependent on other metrics and changes because an underlying root-cause metric has changed.
	ImpactRelevant *bool `json:"impactRelevant,omitempty" tf:"impact_relevant,omitempty"`

	// (Number) The latency of the metric, in minutes.
	// The latency of the metric, in minutes.
	//
	// The latency is the expected reporting delay (for example, caused by constraints of cloud vendors or other third-party data sources) between the observation of a metric data point and its availability in Dynatrace.
	//
	// The allowed value range is from 1 to 60 minutes.
	Latency *float64 `json:"latency,omitempty" tf:"latency,omitempty"`

	// (Number) The maximum allowed value of the metric.
	// The maximum allowed value of the metric.
	MaxValue *float64 `json:"maxValue,omitempty" tf:"max_value,omitempty"`

	// (Number) The minimum allowed value of the metric.
	// The minimum allowed value of the metric.
	MinValue *float64 `json:"minValue,omitempty" tf:"min_value,omitempty"`

	// (Boolean) Whether (true or false) the metric is related to a root cause of a problem.
	// Whether (true or false) the metric is related to a root cause of a problem.
	//
	// A root-cause relevant metric represents a strong indicator for a faulty component.
	RootCauseRelevant *bool `json:"rootCauseRelevant,omitempty" tf:"root_cause_relevant,omitempty"`

	// (String) Possible Values: Error, Score, Unknown
	// Possible Values: `Error`, `Score`, `Unknown`
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

type MetricPropertiesObservation struct {

	// (Boolean) Whether (true or false) the metric is relevant to a problem's impact.
	// Whether (true or false) the metric is relevant to a problem's impact.
	//
	// An impact-relevant metric is highly dependent on other metrics and changes because an underlying root-cause metric has changed.
	ImpactRelevant *bool `json:"impactRelevant,omitempty" tf:"impact_relevant,omitempty"`

	// (Number) The latency of the metric, in minutes.
	// The latency of the metric, in minutes.
	//
	// The latency is the expected reporting delay (for example, caused by constraints of cloud vendors or other third-party data sources) between the observation of a metric data point and its availability in Dynatrace.
	//
	// The allowed value range is from 1 to 60 minutes.
	Latency *float64 `json:"latency,omitempty" tf:"latency,omitempty"`

	// (Number) The maximum allowed value of the metric.
	// The maximum allowed value of the metric.
	MaxValue *float64 `json:"maxValue,omitempty" tf:"max_value,omitempty"`

	// (Number) The minimum allowed value of the metric.
	// The minimum allowed value of the metric.
	MinValue *float64 `json:"minValue,omitempty" tf:"min_value,omitempty"`

	// (Boolean) Whether (true or false) the metric is related to a root cause of a problem.
	// Whether (true or false) the metric is related to a root cause of a problem.
	//
	// A root-cause relevant metric represents a strong indicator for a faulty component.
	RootCauseRelevant *bool `json:"rootCauseRelevant,omitempty" tf:"root_cause_relevant,omitempty"`

	// (String) Possible Values: Error, Score, Unknown
	// Possible Values: `Error`, `Score`, `Unknown`
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

type MetricPropertiesParameters struct {

	// (Boolean) Whether (true or false) the metric is relevant to a problem's impact.
	// Whether (true or false) the metric is relevant to a problem's impact.
	//
	// An impact-relevant metric is highly dependent on other metrics and changes because an underlying root-cause metric has changed.
	// +kubebuilder:validation:Optional
	ImpactRelevant *bool `json:"impactRelevant,omitempty" tf:"impact_relevant,omitempty"`

	// (Number) The latency of the metric, in minutes.
	// The latency of the metric, in minutes.
	//
	// The latency is the expected reporting delay (for example, caused by constraints of cloud vendors or other third-party data sources) between the observation of a metric data point and its availability in Dynatrace.
	//
	// The allowed value range is from 1 to 60 minutes.
	// +kubebuilder:validation:Optional
	Latency *float64 `json:"latency,omitempty" tf:"latency,omitempty"`

	// (Number) The maximum allowed value of the metric.
	// The maximum allowed value of the metric.
	// +kubebuilder:validation:Optional
	MaxValue *float64 `json:"maxValue,omitempty" tf:"max_value,omitempty"`

	// (Number) The minimum allowed value of the metric.
	// The minimum allowed value of the metric.
	// +kubebuilder:validation:Optional
	MinValue *float64 `json:"minValue,omitempty" tf:"min_value,omitempty"`

	// (Boolean) Whether (true or false) the metric is related to a root cause of a problem.
	// Whether (true or false) the metric is related to a root cause of a problem.
	//
	// A root-cause relevant metric represents a strong indicator for a faulty component.
	// +kubebuilder:validation:Optional
	RootCauseRelevant *bool `json:"rootCauseRelevant,omitempty" tf:"root_cause_relevant,omitempty"`

	// (String) Possible Values: Error, Score, Unknown
	// Possible Values: `Error`, `Score`, `Unknown`
	// +kubebuilder:validation:Optional
	ValueType *string `json:"valueType" tf:"value_type,omitempty"`
}

// MetricMetadataSpec defines the desired state of MetricMetadata
type MetricMetadataSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MetricMetadataParameters `json:"forProvider"`
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
	InitProvider MetricMetadataInitParameters `json:"initProvider,omitempty"`
}

// MetricMetadataStatus defines the observed state of MetricMetadata.
type MetricMetadataStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MetricMetadataObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MetricMetadata is the Schema for the MetricMetadatas API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type MetricMetadata struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metricId) || (has(self.initProvider) && has(self.initProvider.metricId))",message="spec.forProvider.metricId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.unit) || (has(self.initProvider) && has(self.initProvider.unit))",message="spec.forProvider.unit is a required parameter"
	Spec   MetricMetadataSpec   `json:"spec"`
	Status MetricMetadataStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MetricMetadataList contains a list of MetricMetadatas
type MetricMetadataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetricMetadata `json:"items"`
}

// Repository type metadata.
var (
	MetricMetadata_Kind             = "MetricMetadata"
	MetricMetadata_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MetricMetadata_Kind}.String()
	MetricMetadata_KindAPIVersion   = MetricMetadata_Kind + "." + CRDGroupVersion.String()
	MetricMetadata_GroupVersionKind = CRDGroupVersion.WithKind(MetricMetadata_Kind)
)

func init() {
	SchemeBuilder.Register(&MetricMetadata{}, &MetricMetadataList{})
}