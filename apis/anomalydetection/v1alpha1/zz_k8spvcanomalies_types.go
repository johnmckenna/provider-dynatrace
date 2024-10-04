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

type K8SPvcAnomaliesInitParameters struct {

	// (Block List, Min: 1, Max: 1) Alerts on low disk space in megabytes for a persistent volume claim. (see below for nested schema)
	// Alerts on low disk space in megabytes for a persistent volume claim.
	LowDiskSpaceCritical []LowDiskSpaceCriticalInitParameters `json:"lowDiskSpaceCritical,omitempty" tf:"low_disk_space_critical,omitempty"`

	// (Block List, Min: 1, Max: 1) Alerts on low disk space in % for a persistent volume claim. (see below for nested schema)
	// Alerts on low disk space in % for a persistent volume claim.
	LowDiskSpaceCriticalPercentage []LowDiskSpaceCriticalPercentageInitParameters `json:"lowDiskSpaceCriticalPercentage,omitempty" tf:"low_disk_space_critical_percentage,omitempty"`

	// (String) The scope of this setting (CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	// The scope of this setting (CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type K8SPvcAnomaliesObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Min: 1, Max: 1) Alerts on low disk space in megabytes for a persistent volume claim. (see below for nested schema)
	// Alerts on low disk space in megabytes for a persistent volume claim.
	LowDiskSpaceCritical []LowDiskSpaceCriticalObservation `json:"lowDiskSpaceCritical,omitempty" tf:"low_disk_space_critical,omitempty"`

	// (Block List, Min: 1, Max: 1) Alerts on low disk space in % for a persistent volume claim. (see below for nested schema)
	// Alerts on low disk space in % for a persistent volume claim.
	LowDiskSpaceCriticalPercentage []LowDiskSpaceCriticalPercentageObservation `json:"lowDiskSpaceCriticalPercentage,omitempty" tf:"low_disk_space_critical_percentage,omitempty"`

	// (String) The scope of this setting (CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	// The scope of this setting (CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type K8SPvcAnomaliesParameters struct {

	// (Block List, Min: 1, Max: 1) Alerts on low disk space in megabytes for a persistent volume claim. (see below for nested schema)
	// Alerts on low disk space in megabytes for a persistent volume claim.
	// +kubebuilder:validation:Optional
	LowDiskSpaceCritical []LowDiskSpaceCriticalParameters `json:"lowDiskSpaceCritical,omitempty" tf:"low_disk_space_critical,omitempty"`

	// (Block List, Min: 1, Max: 1) Alerts on low disk space in % for a persistent volume claim. (see below for nested schema)
	// Alerts on low disk space in % for a persistent volume claim.
	// +kubebuilder:validation:Optional
	LowDiskSpaceCriticalPercentage []LowDiskSpaceCriticalPercentageParameters `json:"lowDiskSpaceCriticalPercentage,omitempty" tf:"low_disk_space_critical_percentage,omitempty"`

	// (String) The scope of this setting (CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	// The scope of this setting (CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type LowDiskSpaceCriticalConfigurationInitParameters struct {

	// (Number) within the last
	// within the last
	ObservationPeriodInMinutes *float64 `json:"observationPeriodInMinutes,omitempty" tf:"observation_period_in_minutes,omitempty"`

	// (Number) for at least
	// for at least
	SamplePeriodInMinutes *float64 `json:"samplePeriodInMinutes,omitempty" tf:"sample_period_in_minutes,omitempty"`

	// (Number) the available disk space is below
	// the available disk space is below
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type LowDiskSpaceCriticalConfigurationObservation struct {

	// (Number) within the last
	// within the last
	ObservationPeriodInMinutes *float64 `json:"observationPeriodInMinutes,omitempty" tf:"observation_period_in_minutes,omitempty"`

	// (Number) for at least
	// for at least
	SamplePeriodInMinutes *float64 `json:"samplePeriodInMinutes,omitempty" tf:"sample_period_in_minutes,omitempty"`

	// (Number) the available disk space is below
	// the available disk space is below
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type LowDiskSpaceCriticalConfigurationParameters struct {

	// (Number) within the last
	// within the last
	// +kubebuilder:validation:Optional
	ObservationPeriodInMinutes *float64 `json:"observationPeriodInMinutes" tf:"observation_period_in_minutes,omitempty"`

	// (Number) for at least
	// for at least
	// +kubebuilder:validation:Optional
	SamplePeriodInMinutes *float64 `json:"samplePeriodInMinutes" tf:"sample_period_in_minutes,omitempty"`

	// (Number) the available disk space is below
	// the available disk space is below
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`
}

type LowDiskSpaceCriticalInitParameters struct {

	// (Block List, Max: 1) Alert if (see below for nested schema)
	// Alert if
	Configuration []LowDiskSpaceCriticalConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type LowDiskSpaceCriticalObservation struct {

	// (Block List, Max: 1) Alert if (see below for nested schema)
	// Alert if
	Configuration []LowDiskSpaceCriticalConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type LowDiskSpaceCriticalParameters struct {

	// (Block List, Max: 1) Alert if (see below for nested schema)
	// Alert if
	// +kubebuilder:validation:Optional
	Configuration []LowDiskSpaceCriticalConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type LowDiskSpaceCriticalPercentageConfigurationInitParameters struct {

	// (Number) within the last
	// within the last
	ObservationPeriodInMinutes *float64 `json:"observationPeriodInMinutes,omitempty" tf:"observation_period_in_minutes,omitempty"`

	// (Number) for at least
	// for at least
	SamplePeriodInMinutes *float64 `json:"samplePeriodInMinutes,omitempty" tf:"sample_period_in_minutes,omitempty"`

	// (Number) the available disk space is below
	// the available disk space is below
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type LowDiskSpaceCriticalPercentageConfigurationObservation struct {

	// (Number) within the last
	// within the last
	ObservationPeriodInMinutes *float64 `json:"observationPeriodInMinutes,omitempty" tf:"observation_period_in_minutes,omitempty"`

	// (Number) for at least
	// for at least
	SamplePeriodInMinutes *float64 `json:"samplePeriodInMinutes,omitempty" tf:"sample_period_in_minutes,omitempty"`

	// (Number) the available disk space is below
	// the available disk space is below
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type LowDiskSpaceCriticalPercentageConfigurationParameters struct {

	// (Number) within the last
	// within the last
	// +kubebuilder:validation:Optional
	ObservationPeriodInMinutes *float64 `json:"observationPeriodInMinutes" tf:"observation_period_in_minutes,omitempty"`

	// (Number) for at least
	// for at least
	// +kubebuilder:validation:Optional
	SamplePeriodInMinutes *float64 `json:"samplePeriodInMinutes" tf:"sample_period_in_minutes,omitempty"`

	// (Number) the available disk space is below
	// the available disk space is below
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`
}

type LowDiskSpaceCriticalPercentageInitParameters struct {

	// (Block List, Max: 1) Alert if (see below for nested schema)
	// Alert if
	Configuration []LowDiskSpaceCriticalPercentageConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type LowDiskSpaceCriticalPercentageObservation struct {

	// (Block List, Max: 1) Alert if (see below for nested schema)
	// Alert if
	Configuration []LowDiskSpaceCriticalPercentageConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type LowDiskSpaceCriticalPercentageParameters struct {

	// (Block List, Max: 1) Alert if (see below for nested schema)
	// Alert if
	// +kubebuilder:validation:Optional
	Configuration []LowDiskSpaceCriticalPercentageConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

// K8SPvcAnomaliesSpec defines the desired state of K8SPvcAnomalies
type K8SPvcAnomaliesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     K8SPvcAnomaliesParameters `json:"forProvider"`
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
	InitProvider K8SPvcAnomaliesInitParameters `json:"initProvider,omitempty"`
}

// K8SPvcAnomaliesStatus defines the observed state of K8SPvcAnomalies.
type K8SPvcAnomaliesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        K8SPvcAnomaliesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// K8SPvcAnomalies is the Schema for the K8SPvcAnomaliess API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type K8SPvcAnomalies struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lowDiskSpaceCritical) || (has(self.initProvider) && has(self.initProvider.lowDiskSpaceCritical))",message="spec.forProvider.lowDiskSpaceCritical is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lowDiskSpaceCriticalPercentage) || (has(self.initProvider) && has(self.initProvider.lowDiskSpaceCriticalPercentage))",message="spec.forProvider.lowDiskSpaceCriticalPercentage is a required parameter"
	Spec   K8SPvcAnomaliesSpec   `json:"spec"`
	Status K8SPvcAnomaliesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8SPvcAnomaliesList contains a list of K8SPvcAnomaliess
type K8SPvcAnomaliesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8SPvcAnomalies `json:"items"`
}

// Repository type metadata.
var (
	K8SPvcAnomalies_Kind             = "K8SPvcAnomalies"
	K8SPvcAnomalies_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: K8SPvcAnomalies_Kind}.String()
	K8SPvcAnomalies_KindAPIVersion   = K8SPvcAnomalies_Kind + "." + CRDGroupVersion.String()
	K8SPvcAnomalies_GroupVersionKind = CRDGroupVersion.WithKind(K8SPvcAnomalies_Kind)
)

func init() {
	SchemeBuilder.Register(&K8SPvcAnomalies{}, &K8SPvcAnomaliesList{})
}
