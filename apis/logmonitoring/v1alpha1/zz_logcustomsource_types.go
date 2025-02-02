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

type ContextContextInitParameters struct {

	// (String) Possible Values: Dt_entity_process_group
	// Possible Values: `Dt_entity_process_group`
	Attribute *string `json:"attribute,omitempty" tf:"attribute,omitempty"`

	// (Set of String)  It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name.
	// no documentation available
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ContextContextObservation struct {

	// (String) Possible Values: Dt_entity_process_group
	// Possible Values: `Dt_entity_process_group`
	Attribute *string `json:"attribute,omitempty" tf:"attribute,omitempty"`

	// (Set of String)  It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name.
	// no documentation available
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ContextContextParameters struct {

	// (String) Possible Values: Dt_entity_process_group
	// Possible Values: `Dt_entity_process_group`
	// +kubebuilder:validation:Optional
	Attribute *string `json:"attribute" tf:"attribute,omitempty"`

	// (Set of String)  It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name.
	// no documentation available
	// +kubebuilder:validation:Optional
	// +listType=set
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ContextInitParameters struct {

	// (Block List, Max: 1) Define Custom Log Source only within context if provided (see below for nested schema)
	Context []ContextContextInitParameters `json:"context,omitempty" tf:"context,omitempty"`
}

type ContextObservation struct {

	// (Block List, Max: 1) Define Custom Log Source only within context if provided (see below for nested schema)
	Context []ContextContextObservation `json:"context,omitempty" tf:"context,omitempty"`
}

type ContextParameters struct {

	// (Block List, Max: 1) Define Custom Log Source only within context if provided (see below for nested schema)
	// +kubebuilder:validation:Optional
	Context []ContextContextParameters `json:"context" tf:"context,omitempty"`
}

type CustomLogSourceInitParameters struct {

	// (Boolean) Accept binary content
	// Accept binary content
	AcceptBinary *bool `json:"acceptBinary,omitempty" tf:"accept_binary,omitempty"`

	// (String) Possible Values: LOG_PATH_PATTERN, WINDOWS_EVENT_LOG
	// Possible Values: `LOG_PATH_PATTERN`, `WINDOWS_EVENT_LOG`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Set of String)  It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name.
	// (Required attribute for cluster v1.291 and under) It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`

	// (Block List, Max: 1)  It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name. (see below for nested schema)
	// (Required attribute for cluster v1.292+) It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name.
	ValuesAndEnrichment []ValuesAndEnrichmentInitParameters `json:"valuesAndEnrichment,omitempty" tf:"values_and_enrichment,omitempty"`
}

type CustomLogSourceObservation struct {

	// (Boolean) Accept binary content
	// Accept binary content
	AcceptBinary *bool `json:"acceptBinary,omitempty" tf:"accept_binary,omitempty"`

	// (String) Possible Values: LOG_PATH_PATTERN, WINDOWS_EVENT_LOG
	// Possible Values: `LOG_PATH_PATTERN`, `WINDOWS_EVENT_LOG`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Set of String)  It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name.
	// (Required attribute for cluster v1.291 and under) It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`

	// (Block List, Max: 1)  It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name. (see below for nested schema)
	// (Required attribute for cluster v1.292+) It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name.
	ValuesAndEnrichment []ValuesAndEnrichmentObservation `json:"valuesAndEnrichment,omitempty" tf:"values_and_enrichment,omitempty"`
}

type CustomLogSourceParameters struct {

	// (Boolean) Accept binary content
	// Accept binary content
	// +kubebuilder:validation:Optional
	AcceptBinary *bool `json:"acceptBinary,omitempty" tf:"accept_binary,omitempty"`

	// (String) Possible Values: LOG_PATH_PATTERN, WINDOWS_EVENT_LOG
	// Possible Values: `LOG_PATH_PATTERN`, `WINDOWS_EVENT_LOG`
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// (Set of String)  It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name.
	// (Required attribute for cluster v1.291 and under) It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name.
	// +kubebuilder:validation:Optional
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`

	// (Block List, Max: 1)  It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name. (see below for nested schema)
	// (Required attribute for cluster v1.292+) It might be either an absolute path to log(s) with optional wildcards or Windows Event Log name.
	// +kubebuilder:validation:Optional
	ValuesAndEnrichment []ValuesAndEnrichmentParameters `json:"valuesAndEnrichment,omitempty" tf:"values_and_enrichment,omitempty"`
}

type CustomLogSourceWithEnrichmentInitParameters struct {

	// (Block List, Max: 1) Optional field that allows to define attributes that will enrich logs. ${N} can be used in attribute value to expand the value matched by wildcards where N denotes the number of the wildcard the expand (see below for nested schema)
	// Optional field that allows to define attributes that will enrich logs. ${N} can be used in attribute value to expand the value matched by wildcards where N denotes the number of the wildcard the expand
	Enrichment []EnrichmentInitParameters `json:"enrichment,omitempty" tf:"enrichment,omitempty"`

	// (String) Values
	// Values
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type CustomLogSourceWithEnrichmentObservation struct {

	// (Block List, Max: 1) Optional field that allows to define attributes that will enrich logs. ${N} can be used in attribute value to expand the value matched by wildcards where N denotes the number of the wildcard the expand (see below for nested schema)
	// Optional field that allows to define attributes that will enrich logs. ${N} can be used in attribute value to expand the value matched by wildcards where N denotes the number of the wildcard the expand
	Enrichment []EnrichmentObservation `json:"enrichment,omitempty" tf:"enrichment,omitempty"`

	// (String) Values
	// Values
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type CustomLogSourceWithEnrichmentParameters struct {

	// (Block List, Max: 1) Optional field that allows to define attributes that will enrich logs. ${N} can be used in attribute value to expand the value matched by wildcards where N denotes the number of the wildcard the expand (see below for nested schema)
	// Optional field that allows to define attributes that will enrich logs. ${N} can be used in attribute value to expand the value matched by wildcards where N denotes the number of the wildcard the expand
	// +kubebuilder:validation:Optional
	Enrichment []EnrichmentParameters `json:"enrichment,omitempty" tf:"enrichment,omitempty"`

	// (String) Values
	// Values
	// +kubebuilder:validation:Optional
	Path *string `json:"path" tf:"path,omitempty"`
}

type EnrichmentEnrichmentInitParameters struct {

	// (String) no documentation available
	// no documentation available
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) Possible Values: LOG_PATH_PATTERN, WINDOWS_EVENT_LOG
	// Possible Values: `Attribute`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) no documentation available
	// no documentation available
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EnrichmentEnrichmentObservation struct {

	// (String) no documentation available
	// no documentation available
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) Possible Values: LOG_PATH_PATTERN, WINDOWS_EVENT_LOG
	// Possible Values: `Attribute`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) no documentation available
	// no documentation available
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EnrichmentEnrichmentParameters struct {

	// (String) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) Possible Values: LOG_PATH_PATTERN, WINDOWS_EVENT_LOG
	// Possible Values: `Attribute`
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// (String) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EnrichmentInitParameters struct {

	// (Block List, Max: 1) Optional field that allows to define attributes that will enrich logs. ${N} can be used in attribute value to expand the value matched by wildcards where N denotes the number of the wildcard the expand (see below for nested schema)
	Enrichment []EnrichmentEnrichmentInitParameters `json:"enrichment,omitempty" tf:"enrichment,omitempty"`
}

type EnrichmentObservation struct {

	// (Block List, Max: 1) Optional field that allows to define attributes that will enrich logs. ${N} can be used in attribute value to expand the value matched by wildcards where N denotes the number of the wildcard the expand (see below for nested schema)
	Enrichment []EnrichmentEnrichmentObservation `json:"enrichment,omitempty" tf:"enrichment,omitempty"`
}

type EnrichmentParameters struct {

	// (Block List, Max: 1) Optional field that allows to define attributes that will enrich logs. ${N} can be used in attribute value to expand the value matched by wildcards where N denotes the number of the wildcard the expand (see below for nested schema)
	// +kubebuilder:validation:Optional
	Enrichment []EnrichmentEnrichmentParameters `json:"enrichment" tf:"enrichment,omitempty"`
}

type LogCustomSourceInitParameters struct {

	// (Block List, Max: 1) Define Custom Log Source only within context if provided (see below for nested schema)
	// Define Custom Log Source only within context if provided
	Context []ContextInitParameters `json:"context,omitempty" tf:"context,omitempty"`

	// (Block List, Min: 1, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	CustomLogSource []CustomLogSourceInitParameters `json:"customLogSource,omitempty" tf:"custom_log_source,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Name
	// Name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type LogCustomSourceObservation struct {

	// (Block List, Max: 1) Define Custom Log Source only within context if provided (see below for nested schema)
	// Define Custom Log Source only within context if provided
	Context []ContextObservation `json:"context,omitempty" tf:"context,omitempty"`

	// (Block List, Min: 1, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	CustomLogSource []CustomLogSourceObservation `json:"customLogSource,omitempty" tf:"custom_log_source,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Name
	// Name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type LogCustomSourceParameters struct {

	// (Block List, Max: 1) Define Custom Log Source only within context if provided (see below for nested schema)
	// Define Custom Log Source only within context if provided
	// +kubebuilder:validation:Optional
	Context []ContextParameters `json:"context,omitempty" tf:"context,omitempty"`

	// (Block List, Min: 1, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	// +kubebuilder:validation:Optional
	CustomLogSource []CustomLogSourceParameters `json:"customLogSource,omitempty" tf:"custom_log_source,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Name
	// Name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type ValuesAndEnrichmentInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	CustomLogSourceWithEnrichment []CustomLogSourceWithEnrichmentInitParameters `json:"customLogSourceWithEnrichment,omitempty" tf:"custom_log_source_with_enrichment,omitempty"`
}

type ValuesAndEnrichmentObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	CustomLogSourceWithEnrichment []CustomLogSourceWithEnrichmentObservation `json:"customLogSourceWithEnrichment,omitempty" tf:"custom_log_source_with_enrichment,omitempty"`
}

type ValuesAndEnrichmentParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	CustomLogSourceWithEnrichment []CustomLogSourceWithEnrichmentParameters `json:"customLogSourceWithEnrichment" tf:"custom_log_source_with_enrichment,omitempty"`
}

// LogCustomSourceSpec defines the desired state of LogCustomSource
type LogCustomSourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogCustomSourceParameters `json:"forProvider"`
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
	InitProvider LogCustomSourceInitParameters `json:"initProvider,omitempty"`
}

// LogCustomSourceStatus defines the observed state of LogCustomSource.
type LogCustomSourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogCustomSourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LogCustomSource is the Schema for the LogCustomSources API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type LogCustomSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.customLogSource) || (has(self.initProvider) && has(self.initProvider.customLogSource))",message="spec.forProvider.customLogSource is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   LogCustomSourceSpec   `json:"spec"`
	Status LogCustomSourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogCustomSourceList contains a list of LogCustomSources
type LogCustomSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogCustomSource `json:"items"`
}

// Repository type metadata.
var (
	LogCustomSource_Kind             = "LogCustomSource"
	LogCustomSource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogCustomSource_Kind}.String()
	LogCustomSource_KindAPIVersion   = LogCustomSource_Kind + "." + CRDGroupVersion.String()
	LogCustomSource_GroupVersionKind = CRDGroupVersion.WithKind(LogCustomSource_Kind)
)

func init() {
	SchemeBuilder.Register(&LogCustomSource{}, &LogCustomSourceList{})
}
