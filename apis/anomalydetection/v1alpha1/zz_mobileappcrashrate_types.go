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

type CrashRateIncreaseCrashRateIncreaseAutoInitParameters struct {

	// (Number) Dynatrace learns the typical crash rate for all app versions and will create an alert if the baseline is violated by more than a specified threshold. Analysis happens based on a sliding window of 10 minutes.
	// Dynatrace learns the typical crash rate for all app versions and will create an alert if the baseline is violated by more than a specified threshold. Analysis happens based on a sliding window of 10 minutes.
	BaselineViolationPercentage *float64 `json:"baselineViolationPercentage,omitempty" tf:"baseline_violation_percentage,omitempty"`

	// (Number) Amount of users
	// Amount of users
	ConcurrentUsers *float64 `json:"concurrentUsers,omitempty" tf:"concurrent_users,omitempty"`

	// (String) Possible Values: Low, Medium, High
	// Possible Values: `Low`, `Medium`, `High`
	Sensitivity *string `json:"sensitivity,omitempty" tf:"sensitivity,omitempty"`
}

type CrashRateIncreaseCrashRateIncreaseAutoObservation struct {

	// (Number) Dynatrace learns the typical crash rate for all app versions and will create an alert if the baseline is violated by more than a specified threshold. Analysis happens based on a sliding window of 10 minutes.
	// Dynatrace learns the typical crash rate for all app versions and will create an alert if the baseline is violated by more than a specified threshold. Analysis happens based on a sliding window of 10 minutes.
	BaselineViolationPercentage *float64 `json:"baselineViolationPercentage,omitempty" tf:"baseline_violation_percentage,omitempty"`

	// (Number) Amount of users
	// Amount of users
	ConcurrentUsers *float64 `json:"concurrentUsers,omitempty" tf:"concurrent_users,omitempty"`

	// (String) Possible Values: Low, Medium, High
	// Possible Values: `Low`, `Medium`, `High`
	Sensitivity *string `json:"sensitivity,omitempty" tf:"sensitivity,omitempty"`
}

type CrashRateIncreaseCrashRateIncreaseAutoParameters struct {

	// (Number) Dynatrace learns the typical crash rate for all app versions and will create an alert if the baseline is violated by more than a specified threshold. Analysis happens based on a sliding window of 10 minutes.
	// Dynatrace learns the typical crash rate for all app versions and will create an alert if the baseline is violated by more than a specified threshold. Analysis happens based on a sliding window of 10 minutes.
	// +kubebuilder:validation:Optional
	BaselineViolationPercentage *float64 `json:"baselineViolationPercentage" tf:"baseline_violation_percentage,omitempty"`

	// (Number) Amount of users
	// Amount of users
	// +kubebuilder:validation:Optional
	ConcurrentUsers *float64 `json:"concurrentUsers" tf:"concurrent_users,omitempty"`

	// (String) Possible Values: Low, Medium, High
	// Possible Values: `Low`, `Medium`, `High`
	// +kubebuilder:validation:Optional
	Sensitivity *string `json:"sensitivity" tf:"sensitivity,omitempty"`
}

type CrashRateIncreaseCrashRateIncreaseFixedInitParameters struct {

	// (Number) Absolute threshold
	// Absolute threshold
	AbsoluteCrashRate *float64 `json:"absoluteCrashRate,omitempty" tf:"absolute_crash_rate,omitempty"`

	// (Number) Amount of users
	// Amount of users
	ConcurrentUsers *float64 `json:"concurrentUsers,omitempty" tf:"concurrent_users,omitempty"`
}

type CrashRateIncreaseCrashRateIncreaseFixedObservation struct {

	// (Number) Absolute threshold
	// Absolute threshold
	AbsoluteCrashRate *float64 `json:"absoluteCrashRate,omitempty" tf:"absolute_crash_rate,omitempty"`

	// (Number) Amount of users
	// Amount of users
	ConcurrentUsers *float64 `json:"concurrentUsers,omitempty" tf:"concurrent_users,omitempty"`
}

type CrashRateIncreaseCrashRateIncreaseFixedParameters struct {

	// (Number) Absolute threshold
	// Absolute threshold
	// +kubebuilder:validation:Optional
	AbsoluteCrashRate *float64 `json:"absoluteCrashRate" tf:"absolute_crash_rate,omitempty"`

	// (Number) Amount of users
	// Amount of users
	// +kubebuilder:validation:Optional
	ConcurrentUsers *float64 `json:"concurrentUsers" tf:"concurrent_users,omitempty"`
}

type MobileAppCrashRateCrashRateIncreaseInitParameters struct {

	// detected baseline is exceeded by a certain number of users (see below for nested schema)
	// Alert crash rate increases when auto-detected baseline is exceeded by a certain number of users
	CrashRateIncreaseAuto []CrashRateIncreaseCrashRateIncreaseAutoInitParameters `json:"crashRateIncreaseAuto,omitempty" tf:"crash_rate_increase_auto,omitempty"`

	// (Block List, Max: 1) Alert crash rate increases when the defined threshold is exceeded by a certain number of users (see below for nested schema)
	// Alert crash rate increases when the defined threshold is exceeded by a certain number of users
	CrashRateIncreaseFixed []CrashRateIncreaseCrashRateIncreaseFixedInitParameters `json:"crashRateIncreaseFixed,omitempty" tf:"crash_rate_increase_fixed,omitempty"`

	// (String) Possible Values: Auto, Fixed
	// Possible Values: `Auto`, `Fixed`
	DetectionMode *string `json:"detectionMode,omitempty" tf:"detection_mode,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type MobileAppCrashRateCrashRateIncreaseObservation struct {

	// detected baseline is exceeded by a certain number of users (see below for nested schema)
	// Alert crash rate increases when auto-detected baseline is exceeded by a certain number of users
	CrashRateIncreaseAuto []CrashRateIncreaseCrashRateIncreaseAutoObservation `json:"crashRateIncreaseAuto,omitempty" tf:"crash_rate_increase_auto,omitempty"`

	// (Block List, Max: 1) Alert crash rate increases when the defined threshold is exceeded by a certain number of users (see below for nested schema)
	// Alert crash rate increases when the defined threshold is exceeded by a certain number of users
	CrashRateIncreaseFixed []CrashRateIncreaseCrashRateIncreaseFixedObservation `json:"crashRateIncreaseFixed,omitempty" tf:"crash_rate_increase_fixed,omitempty"`

	// (String) Possible Values: Auto, Fixed
	// Possible Values: `Auto`, `Fixed`
	DetectionMode *string `json:"detectionMode,omitempty" tf:"detection_mode,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type MobileAppCrashRateCrashRateIncreaseParameters struct {

	// detected baseline is exceeded by a certain number of users (see below for nested schema)
	// Alert crash rate increases when auto-detected baseline is exceeded by a certain number of users
	// +kubebuilder:validation:Optional
	CrashRateIncreaseAuto []CrashRateIncreaseCrashRateIncreaseAutoParameters `json:"crashRateIncreaseAuto,omitempty" tf:"crash_rate_increase_auto,omitempty"`

	// (Block List, Max: 1) Alert crash rate increases when the defined threshold is exceeded by a certain number of users (see below for nested schema)
	// Alert crash rate increases when the defined threshold is exceeded by a certain number of users
	// +kubebuilder:validation:Optional
	CrashRateIncreaseFixed []CrashRateIncreaseCrashRateIncreaseFixedParameters `json:"crashRateIncreaseFixed,omitempty" tf:"crash_rate_increase_fixed,omitempty"`

	// (String) Possible Values: Auto, Fixed
	// Possible Values: `Auto`, `Fixed`
	// +kubebuilder:validation:Optional
	DetectionMode *string `json:"detectionMode,omitempty" tf:"detection_mode,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type MobileAppCrashRateInitParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (Block List, Min: 1, Max: 1) Crash rate increase (see below for nested schema)
	// Crash rate increase
	CrashRateIncrease []MobileAppCrashRateCrashRateIncreaseInitParameters `json:"crashRateIncrease,omitempty" tf:"crash_rate_increase,omitempty"`
}

type MobileAppCrashRateObservation struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (Block List, Min: 1, Max: 1) Crash rate increase (see below for nested schema)
	// Crash rate increase
	CrashRateIncrease []MobileAppCrashRateCrashRateIncreaseObservation `json:"crashRateIncrease,omitempty" tf:"crash_rate_increase,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MobileAppCrashRateParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (Block List, Min: 1, Max: 1) Crash rate increase (see below for nested schema)
	// Crash rate increase
	// +kubebuilder:validation:Optional
	CrashRateIncrease []MobileAppCrashRateCrashRateIncreaseParameters `json:"crashRateIncrease,omitempty" tf:"crash_rate_increase,omitempty"`
}

// MobileAppCrashRateSpec defines the desired state of MobileAppCrashRate
type MobileAppCrashRateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MobileAppCrashRateParameters `json:"forProvider"`
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
	InitProvider MobileAppCrashRateInitParameters `json:"initProvider,omitempty"`
}

// MobileAppCrashRateStatus defines the observed state of MobileAppCrashRate.
type MobileAppCrashRateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MobileAppCrashRateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MobileAppCrashRate is the Schema for the MobileAppCrashRates API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type MobileAppCrashRate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.crashRateIncrease) || (has(self.initProvider) && has(self.initProvider.crashRateIncrease))",message="spec.forProvider.crashRateIncrease is a required parameter"
	Spec   MobileAppCrashRateSpec   `json:"spec"`
	Status MobileAppCrashRateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MobileAppCrashRateList contains a list of MobileAppCrashRates
type MobileAppCrashRateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MobileAppCrashRate `json:"items"`
}

// Repository type metadata.
var (
	MobileAppCrashRate_Kind             = "MobileAppCrashRate"
	MobileAppCrashRate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MobileAppCrashRate_Kind}.String()
	MobileAppCrashRate_KindAPIVersion   = MobileAppCrashRate_Kind + "." + CRDGroupVersion.String()
	MobileAppCrashRate_GroupVersionKind = CRDGroupVersion.WithKind(MobileAppCrashRate_Kind)
)

func init() {
	SchemeBuilder.Register(&MobileAppCrashRate{}, &MobileAppCrashRateList{})
}
