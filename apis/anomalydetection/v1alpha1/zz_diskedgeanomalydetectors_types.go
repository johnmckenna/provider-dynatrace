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

type AlertInitParameters struct {

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	SampleCountThresholds []SampleCountThresholdsInitParameters `json:"sampleCountThresholds,omitempty" tf:"sample_count_thresholds,omitempty"`

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	SampleCountThresholdsImmediately []SampleCountThresholdsImmediatelyInitParameters `json:"sampleCountThresholdsImmediately,omitempty" tf:"sample_count_thresholds_immediately,omitempty"`

	// (Number) no documentation available
	// no documentation available
	ThresholdMebibytes *float64 `json:"thresholdMebibytes,omitempty" tf:"threshold_mebibytes,omitempty"`

	// (Number) no documentation available
	// no documentation available
	ThresholdMilliseconds *float64 `json:"thresholdMilliseconds,omitempty" tf:"threshold_milliseconds,omitempty"`

	// (Number) no documentation available
	// no documentation available
	ThresholdNumber *float64 `json:"thresholdNumber,omitempty" tf:"threshold_number,omitempty"`

	// (Number) no documentation available
	// no documentation available
	ThresholdPercent *float64 `json:"thresholdPercent,omitempty" tf:"threshold_percent,omitempty"`

	// (String) Possible Values: AVAILABLE_DISK_SPACE_MEBIBYTES_BELOW, AVAILABLE_DISK_SPACE_PERCENT_BELOW, AVAILABLE_INODES_NUMBER_BELOW, AVAILABLE_INODES_PERCENT_BELOW, READ_ONLY_FILE_SYSTEM, READ_TIME_EXCEEDING, WRITE_TIME_EXCEEDING
	// Possible Values: `AVAILABLE_DISK_SPACE_MEBIBYTES_BELOW`, `AVAILABLE_DISK_SPACE_PERCENT_BELOW`, `AVAILABLE_INODES_NUMBER_BELOW`, `AVAILABLE_INODES_PERCENT_BELOW`, `READ_ONLY_FILE_SYSTEM`, `READ_TIME_EXCEEDING`, `WRITE_TIME_EXCEEDING`
	Trigger *string `json:"trigger,omitempty" tf:"trigger,omitempty"`
}

type AlertObservation struct {

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	SampleCountThresholds []SampleCountThresholdsObservation `json:"sampleCountThresholds,omitempty" tf:"sample_count_thresholds,omitempty"`

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	SampleCountThresholdsImmediately []SampleCountThresholdsImmediatelyObservation `json:"sampleCountThresholdsImmediately,omitempty" tf:"sample_count_thresholds_immediately,omitempty"`

	// (Number) no documentation available
	// no documentation available
	ThresholdMebibytes *float64 `json:"thresholdMebibytes,omitempty" tf:"threshold_mebibytes,omitempty"`

	// (Number) no documentation available
	// no documentation available
	ThresholdMilliseconds *float64 `json:"thresholdMilliseconds,omitempty" tf:"threshold_milliseconds,omitempty"`

	// (Number) no documentation available
	// no documentation available
	ThresholdNumber *float64 `json:"thresholdNumber,omitempty" tf:"threshold_number,omitempty"`

	// (Number) no documentation available
	// no documentation available
	ThresholdPercent *float64 `json:"thresholdPercent,omitempty" tf:"threshold_percent,omitempty"`

	// (String) Possible Values: AVAILABLE_DISK_SPACE_MEBIBYTES_BELOW, AVAILABLE_DISK_SPACE_PERCENT_BELOW, AVAILABLE_INODES_NUMBER_BELOW, AVAILABLE_INODES_PERCENT_BELOW, READ_ONLY_FILE_SYSTEM, READ_TIME_EXCEEDING, WRITE_TIME_EXCEEDING
	// Possible Values: `AVAILABLE_DISK_SPACE_MEBIBYTES_BELOW`, `AVAILABLE_DISK_SPACE_PERCENT_BELOW`, `AVAILABLE_INODES_NUMBER_BELOW`, `AVAILABLE_INODES_PERCENT_BELOW`, `READ_ONLY_FILE_SYSTEM`, `READ_TIME_EXCEEDING`, `WRITE_TIME_EXCEEDING`
	Trigger *string `json:"trigger,omitempty" tf:"trigger,omitempty"`
}

type AlertParameters struct {

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	// +kubebuilder:validation:Optional
	SampleCountThresholds []SampleCountThresholdsParameters `json:"sampleCountThresholds,omitempty" tf:"sample_count_thresholds,omitempty"`

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	// +kubebuilder:validation:Optional
	SampleCountThresholdsImmediately []SampleCountThresholdsImmediatelyParameters `json:"sampleCountThresholdsImmediately,omitempty" tf:"sample_count_thresholds_immediately,omitempty"`

	// (Number) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	ThresholdMebibytes *float64 `json:"thresholdMebibytes,omitempty" tf:"threshold_mebibytes,omitempty"`

	// (Number) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	ThresholdMilliseconds *float64 `json:"thresholdMilliseconds,omitempty" tf:"threshold_milliseconds,omitempty"`

	// (Number) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	ThresholdNumber *float64 `json:"thresholdNumber,omitempty" tf:"threshold_number,omitempty"`

	// (Number) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	ThresholdPercent *float64 `json:"thresholdPercent,omitempty" tf:"threshold_percent,omitempty"`

	// (String) Possible Values: AVAILABLE_DISK_SPACE_MEBIBYTES_BELOW, AVAILABLE_DISK_SPACE_PERCENT_BELOW, AVAILABLE_INODES_NUMBER_BELOW, AVAILABLE_INODES_PERCENT_BELOW, READ_ONLY_FILE_SYSTEM, READ_TIME_EXCEEDING, WRITE_TIME_EXCEEDING
	// Possible Values: `AVAILABLE_DISK_SPACE_MEBIBYTES_BELOW`, `AVAILABLE_DISK_SPACE_PERCENT_BELOW`, `AVAILABLE_INODES_NUMBER_BELOW`, `AVAILABLE_INODES_PERCENT_BELOW`, `READ_ONLY_FILE_SYSTEM`, `READ_TIME_EXCEEDING`, `WRITE_TIME_EXCEEDING`
	// +kubebuilder:validation:Optional
	Trigger *string `json:"trigger" tf:"trigger,omitempty"`
}

type AlertsInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Alert []AlertInitParameters `json:"alert,omitempty" tf:"alert,omitempty"`
}

type AlertsObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Alert []AlertObservation `json:"alert,omitempty" tf:"alert,omitempty"`
}

type AlertsParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Alert []AlertParameters `json:"alert" tf:"alert,omitempty"`
}

type DiskEdgeAnomalyDetectorsInitParameters struct {

	// (Block List, Max: 1) Alerts (see below for nested schema)
	// Alerts
	Alerts []AlertsInitParameters `json:"alerts,omitempty" tf:"alerts,omitempty"`

	// (Set of String) Disk will be included in this policy if any of the filters match
	// Disk will be included in this policy if **any** of the filters match
	// +listType=set
	DiskNameFilters []*string `json:"diskNameFilters,omitempty" tf:"disk_name_filters,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// value properties to be attached to the triggered event. (see below for nested schema)
	// Set of additional key-value properties to be attached to the triggered event.
	EventProperties []EventPropertiesInitParameters `json:"eventProperties,omitempty" tf:"event_properties,omitempty"`

	// (Block List, Max: 1) The policy will be enabled if all conditions are met (see below for nested schema)
	// The policy will be enabled if **all** conditions are met
	HostMetadataConditions []HostMetadataConditionsInitParameters `json:"hostMetadataConditions,omitempty" tf:"host_metadata_conditions,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (Set of String) Select the operating systems on which policy should be applied
	// Select the operating systems on which policy should be applied
	// +listType=set
	OperatingSystem []*string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// (String) Policy name
	// Policy name
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type DiskEdgeAnomalyDetectorsObservation struct {

	// (Block List, Max: 1) Alerts (see below for nested schema)
	// Alerts
	Alerts []AlertsObservation `json:"alerts,omitempty" tf:"alerts,omitempty"`

	// (Set of String) Disk will be included in this policy if any of the filters match
	// Disk will be included in this policy if **any** of the filters match
	// +listType=set
	DiskNameFilters []*string `json:"diskNameFilters,omitempty" tf:"disk_name_filters,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// value properties to be attached to the triggered event. (see below for nested schema)
	// Set of additional key-value properties to be attached to the triggered event.
	EventProperties []EventPropertiesObservation `json:"eventProperties,omitempty" tf:"event_properties,omitempty"`

	// (Block List, Max: 1) The policy will be enabled if all conditions are met (see below for nested schema)
	// The policy will be enabled if **all** conditions are met
	HostMetadataConditions []HostMetadataConditionsObservation `json:"hostMetadataConditions,omitempty" tf:"host_metadata_conditions,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (Set of String) Select the operating systems on which policy should be applied
	// Select the operating systems on which policy should be applied
	// +listType=set
	OperatingSystem []*string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// (String) Policy name
	// Policy name
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type DiskEdgeAnomalyDetectorsParameters struct {

	// (Block List, Max: 1) Alerts (see below for nested schema)
	// Alerts
	// +kubebuilder:validation:Optional
	Alerts []AlertsParameters `json:"alerts,omitempty" tf:"alerts,omitempty"`

	// (Set of String) Disk will be included in this policy if any of the filters match
	// Disk will be included in this policy if **any** of the filters match
	// +kubebuilder:validation:Optional
	// +listType=set
	DiskNameFilters []*string `json:"diskNameFilters,omitempty" tf:"disk_name_filters,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// value properties to be attached to the triggered event. (see below for nested schema)
	// Set of additional key-value properties to be attached to the triggered event.
	// +kubebuilder:validation:Optional
	EventProperties []EventPropertiesParameters `json:"eventProperties,omitempty" tf:"event_properties,omitempty"`

	// (Block List, Max: 1) The policy will be enabled if all conditions are met (see below for nested schema)
	// The policy will be enabled if **all** conditions are met
	// +kubebuilder:validation:Optional
	HostMetadataConditions []HostMetadataConditionsParameters `json:"hostMetadataConditions,omitempty" tf:"host_metadata_conditions,omitempty"`

	// (String) Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched
	// +kubebuilder:validation:Optional
	InsertAfter *string `json:"insertAfter,omitempty" tf:"insert_after,omitempty"`

	// (Set of String) Select the operating systems on which policy should be applied
	// Select the operating systems on which policy should be applied
	// +kubebuilder:validation:Optional
	// +listType=set
	OperatingSystem []*string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// (String) Policy name
	// Policy name
	// +kubebuilder:validation:Optional
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type EventPropertieInitParameters struct {

	// (String) Type 'dt.' for key hints.
	// Type 'dt.' for key hints.
	MetadataKey *string `json:"metadataKey,omitempty" tf:"metadata_key,omitempty"`

	// (String) no documentation available
	// no documentation available
	MetadataValue *string `json:"metadataValue,omitempty" tf:"metadata_value,omitempty"`
}

type EventPropertieObservation struct {

	// (String) Type 'dt.' for key hints.
	// Type 'dt.' for key hints.
	MetadataKey *string `json:"metadataKey,omitempty" tf:"metadata_key,omitempty"`

	// (String) no documentation available
	// no documentation available
	MetadataValue *string `json:"metadataValue,omitempty" tf:"metadata_value,omitempty"`
}

type EventPropertieParameters struct {

	// (String) Type 'dt.' for key hints.
	// Type 'dt.' for key hints.
	// +kubebuilder:validation:Optional
	MetadataKey *string `json:"metadataKey" tf:"metadata_key,omitempty"`

	// (String) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	MetadataValue *string `json:"metadataValue" tf:"metadata_value,omitempty"`
}

type EventPropertiesInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	EventPropertie []EventPropertieInitParameters `json:"eventPropertie,omitempty" tf:"event_propertie,omitempty"`
}

type EventPropertiesObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	EventPropertie []EventPropertieObservation `json:"eventPropertie,omitempty" tf:"event_propertie,omitempty"`
}

type EventPropertiesParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	EventPropertie []EventPropertieParameters `json:"eventPropertie" tf:"event_propertie,omitempty"`
}

type HostMetadataConditionHostMetadataConditionInitParameters struct {

	// (Boolean) When enabled, the condition requires a metadata key to exist and match the constraints; when disabled, the key is optional but must still match the constrains if it is present.
	// When enabled, the condition requires a metadata key to exist and match the constraints; when disabled, the key is optional but must still match the constrains if it is present.
	KeyMustExist *bool `json:"keyMustExist,omitempty" tf:"key_must_exist,omitempty"`

	// (String) This string has to match a required format.
	// This string has to match a required format.
	//
	// - `$contains(production)` – Matches if `production` appears anywhere in the host metadata value.
	// - `$eq(production)` – Matches if `production` matches the host metadata value exactly.
	// - `$prefix(production)` – Matches if `production` matches the prefix of the host metadata value.
	// - `$suffix(production)` – Matches if `production` matches the suffix of the host metadata value.
	//
	// Available logic operations:
	// - `$not($eq(production))` – Matches if the host metadata value is different from `production`.
	// - `$and($prefix(production),$suffix(main))` – Matches if host metadata value starts with `production` and ends with `main`.
	// - `$or($prefix(production),$suffix(main))` – Matches if host metadata value starts with `production` or ends with `main`.
	//
	// Brackets **(** and **)** that are part of the matched property **must be escaped with a tilde (~)**
	MetadataCondition *string `json:"metadataCondition,omitempty" tf:"metadata_condition,omitempty"`

	// (String) Type 'dt.' for key hints.
	// Key
	MetadataKey *string `json:"metadataKey,omitempty" tf:"metadata_key,omitempty"`
}

type HostMetadataConditionHostMetadataConditionObservation struct {

	// (Boolean) When enabled, the condition requires a metadata key to exist and match the constraints; when disabled, the key is optional but must still match the constrains if it is present.
	// When enabled, the condition requires a metadata key to exist and match the constraints; when disabled, the key is optional but must still match the constrains if it is present.
	KeyMustExist *bool `json:"keyMustExist,omitempty" tf:"key_must_exist,omitempty"`

	// (String) This string has to match a required format.
	// This string has to match a required format.
	//
	// - `$contains(production)` – Matches if `production` appears anywhere in the host metadata value.
	// - `$eq(production)` – Matches if `production` matches the host metadata value exactly.
	// - `$prefix(production)` – Matches if `production` matches the prefix of the host metadata value.
	// - `$suffix(production)` – Matches if `production` matches the suffix of the host metadata value.
	//
	// Available logic operations:
	// - `$not($eq(production))` – Matches if the host metadata value is different from `production`.
	// - `$and($prefix(production),$suffix(main))` – Matches if host metadata value starts with `production` and ends with `main`.
	// - `$or($prefix(production),$suffix(main))` – Matches if host metadata value starts with `production` or ends with `main`.
	//
	// Brackets **(** and **)** that are part of the matched property **must be escaped with a tilde (~)**
	MetadataCondition *string `json:"metadataCondition,omitempty" tf:"metadata_condition,omitempty"`

	// (String) Type 'dt.' for key hints.
	// Key
	MetadataKey *string `json:"metadataKey,omitempty" tf:"metadata_key,omitempty"`
}

type HostMetadataConditionHostMetadataConditionParameters struct {

	// (Boolean) When enabled, the condition requires a metadata key to exist and match the constraints; when disabled, the key is optional but must still match the constrains if it is present.
	// When enabled, the condition requires a metadata key to exist and match the constraints; when disabled, the key is optional but must still match the constrains if it is present.
	// +kubebuilder:validation:Optional
	KeyMustExist *bool `json:"keyMustExist,omitempty" tf:"key_must_exist,omitempty"`

	// (String) This string has to match a required format.
	// This string has to match a required format.
	//
	// - `$contains(production)` – Matches if `production` appears anywhere in the host metadata value.
	// - `$eq(production)` – Matches if `production` matches the host metadata value exactly.
	// - `$prefix(production)` – Matches if `production` matches the prefix of the host metadata value.
	// - `$suffix(production)` – Matches if `production` matches the suffix of the host metadata value.
	//
	// Available logic operations:
	// - `$not($eq(production))` – Matches if the host metadata value is different from `production`.
	// - `$and($prefix(production),$suffix(main))` – Matches if host metadata value starts with `production` and ends with `main`.
	// - `$or($prefix(production),$suffix(main))` – Matches if host metadata value starts with `production` or ends with `main`.
	//
	// Brackets **(** and **)** that are part of the matched property **must be escaped with a tilde (~)**
	// +kubebuilder:validation:Optional
	MetadataCondition *string `json:"metadataCondition" tf:"metadata_condition,omitempty"`

	// (String) Type 'dt.' for key hints.
	// Key
	// +kubebuilder:validation:Optional
	MetadataKey *string `json:"metadataKey" tf:"metadata_key,omitempty"`
}

type HostMetadataConditionInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// no documentation available
	HostMetadataCondition []HostMetadataConditionHostMetadataConditionInitParameters `json:"hostMetadataCondition,omitempty" tf:"host_metadata_condition,omitempty"`
}

type HostMetadataConditionObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// no documentation available
	HostMetadataCondition []HostMetadataConditionHostMetadataConditionObservation `json:"hostMetadataCondition,omitempty" tf:"host_metadata_condition,omitempty"`
}

type HostMetadataConditionParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// no documentation available
	// +kubebuilder:validation:Optional
	HostMetadataCondition []HostMetadataConditionHostMetadataConditionParameters `json:"hostMetadataCondition" tf:"host_metadata_condition,omitempty"`
}

type HostMetadataConditionsInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	HostMetadataCondition []HostMetadataConditionInitParameters `json:"hostMetadataCondition,omitempty" tf:"host_metadata_condition,omitempty"`
}

type HostMetadataConditionsObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	HostMetadataCondition []HostMetadataConditionObservation `json:"hostMetadataCondition,omitempty" tf:"host_metadata_condition,omitempty"`
}

type HostMetadataConditionsParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	HostMetadataCondition []HostMetadataConditionParameters `json:"hostMetadataCondition" tf:"host_metadata_condition,omitempty"`
}

type SampleCountThresholdsImmediatelyInitParameters struct {

	// second samples that form the sliding evaluation window for dealerting.
	// The number of **10-second samples** that form the sliding evaluation window for dealerting.
	DealertingEvaluationWindow *float64 `json:"dealertingEvaluationWindow,omitempty" tf:"dealerting_evaluation_window,omitempty"`

	// second samples within the evaluation window that must be lower than the threshold to close an event
	// The number of **10-second samples** within the evaluation window that must be lower than the threshold to close an event
	DealertingSamples *float64 `json:"dealertingSamples,omitempty" tf:"dealerting_samples,omitempty"`

	// second samples that form the sliding evaluation window to detect violating samples.
	// The number of **10-second samples** that form the sliding evaluation window to detect violating samples.
	ViolatingEvaluationWindow *float64 `json:"violatingEvaluationWindow,omitempty" tf:"violating_evaluation_window,omitempty"`

	// second samples within the evaluation window that must exceed the threshold to trigger an event
	// The number of **10-second samples** within the evaluation window that must exceed the threshold to trigger an event
	ViolatingSamples *float64 `json:"violatingSamples,omitempty" tf:"violating_samples,omitempty"`
}

type SampleCountThresholdsImmediatelyObservation struct {

	// second samples that form the sliding evaluation window for dealerting.
	// The number of **10-second samples** that form the sliding evaluation window for dealerting.
	DealertingEvaluationWindow *float64 `json:"dealertingEvaluationWindow,omitempty" tf:"dealerting_evaluation_window,omitempty"`

	// second samples within the evaluation window that must be lower than the threshold to close an event
	// The number of **10-second samples** within the evaluation window that must be lower than the threshold to close an event
	DealertingSamples *float64 `json:"dealertingSamples,omitempty" tf:"dealerting_samples,omitempty"`

	// second samples that form the sliding evaluation window to detect violating samples.
	// The number of **10-second samples** that form the sliding evaluation window to detect violating samples.
	ViolatingEvaluationWindow *float64 `json:"violatingEvaluationWindow,omitempty" tf:"violating_evaluation_window,omitempty"`

	// second samples within the evaluation window that must exceed the threshold to trigger an event
	// The number of **10-second samples** within the evaluation window that must exceed the threshold to trigger an event
	ViolatingSamples *float64 `json:"violatingSamples,omitempty" tf:"violating_samples,omitempty"`
}

type SampleCountThresholdsImmediatelyParameters struct {

	// second samples that form the sliding evaluation window for dealerting.
	// The number of **10-second samples** that form the sliding evaluation window for dealerting.
	// +kubebuilder:validation:Optional
	DealertingEvaluationWindow *float64 `json:"dealertingEvaluationWindow" tf:"dealerting_evaluation_window,omitempty"`

	// second samples within the evaluation window that must be lower than the threshold to close an event
	// The number of **10-second samples** within the evaluation window that must be lower than the threshold to close an event
	// +kubebuilder:validation:Optional
	DealertingSamples *float64 `json:"dealertingSamples" tf:"dealerting_samples,omitempty"`

	// second samples that form the sliding evaluation window to detect violating samples.
	// The number of **10-second samples** that form the sliding evaluation window to detect violating samples.
	// +kubebuilder:validation:Optional
	ViolatingEvaluationWindow *float64 `json:"violatingEvaluationWindow" tf:"violating_evaluation_window,omitempty"`

	// second samples within the evaluation window that must exceed the threshold to trigger an event
	// The number of **10-second samples** within the evaluation window that must exceed the threshold to trigger an event
	// +kubebuilder:validation:Optional
	ViolatingSamples *float64 `json:"violatingSamples" tf:"violating_samples,omitempty"`
}

type SampleCountThresholdsInitParameters struct {

	// second samples that form the sliding evaluation window for dealerting.
	// The number of **10-second samples** that form the sliding evaluation window for dealerting.
	DealertingEvaluationWindow *float64 `json:"dealertingEvaluationWindow,omitempty" tf:"dealerting_evaluation_window,omitempty"`

	// second samples within the evaluation window that must be lower than the threshold to close an event
	// The number of **10-second samples** within the evaluation window that must be lower than the threshold to close an event
	DealertingSamples *float64 `json:"dealertingSamples,omitempty" tf:"dealerting_samples,omitempty"`

	// second samples that form the sliding evaluation window to detect violating samples.
	// The number of **10-second samples** that form the sliding evaluation window to detect violating samples.
	ViolatingEvaluationWindow *float64 `json:"violatingEvaluationWindow,omitempty" tf:"violating_evaluation_window,omitempty"`

	// second samples within the evaluation window that must exceed the threshold to trigger an event
	// The number of **10-second samples** within the evaluation window that must exceed the threshold to trigger an event
	ViolatingSamples *float64 `json:"violatingSamples,omitempty" tf:"violating_samples,omitempty"`
}

type SampleCountThresholdsObservation struct {

	// second samples that form the sliding evaluation window for dealerting.
	// The number of **10-second samples** that form the sliding evaluation window for dealerting.
	DealertingEvaluationWindow *float64 `json:"dealertingEvaluationWindow,omitempty" tf:"dealerting_evaluation_window,omitempty"`

	// second samples within the evaluation window that must be lower than the threshold to close an event
	// The number of **10-second samples** within the evaluation window that must be lower than the threshold to close an event
	DealertingSamples *float64 `json:"dealertingSamples,omitempty" tf:"dealerting_samples,omitempty"`

	// second samples that form the sliding evaluation window to detect violating samples.
	// The number of **10-second samples** that form the sliding evaluation window to detect violating samples.
	ViolatingEvaluationWindow *float64 `json:"violatingEvaluationWindow,omitempty" tf:"violating_evaluation_window,omitempty"`

	// second samples within the evaluation window that must exceed the threshold to trigger an event
	// The number of **10-second samples** within the evaluation window that must exceed the threshold to trigger an event
	ViolatingSamples *float64 `json:"violatingSamples,omitempty" tf:"violating_samples,omitempty"`
}

type SampleCountThresholdsParameters struct {

	// second samples that form the sliding evaluation window for dealerting.
	// The number of **10-second samples** that form the sliding evaluation window for dealerting.
	// +kubebuilder:validation:Optional
	DealertingEvaluationWindow *float64 `json:"dealertingEvaluationWindow" tf:"dealerting_evaluation_window,omitempty"`

	// second samples within the evaluation window that must be lower than the threshold to close an event
	// The number of **10-second samples** within the evaluation window that must be lower than the threshold to close an event
	// +kubebuilder:validation:Optional
	DealertingSamples *float64 `json:"dealertingSamples" tf:"dealerting_samples,omitempty"`

	// second samples that form the sliding evaluation window to detect violating samples.
	// The number of **10-second samples** that form the sliding evaluation window to detect violating samples.
	// +kubebuilder:validation:Optional
	ViolatingEvaluationWindow *float64 `json:"violatingEvaluationWindow" tf:"violating_evaluation_window,omitempty"`

	// second samples within the evaluation window that must exceed the threshold to trigger an event
	// The number of **10-second samples** within the evaluation window that must exceed the threshold to trigger an event
	// +kubebuilder:validation:Optional
	ViolatingSamples *float64 `json:"violatingSamples" tf:"violating_samples,omitempty"`
}

// DiskEdgeAnomalyDetectorsSpec defines the desired state of DiskEdgeAnomalyDetectors
type DiskEdgeAnomalyDetectorsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskEdgeAnomalyDetectorsParameters `json:"forProvider"`
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
	InitProvider DiskEdgeAnomalyDetectorsInitParameters `json:"initProvider,omitempty"`
}

// DiskEdgeAnomalyDetectorsStatus defines the observed state of DiskEdgeAnomalyDetectors.
type DiskEdgeAnomalyDetectorsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskEdgeAnomalyDetectorsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DiskEdgeAnomalyDetectors is the Schema for the DiskEdgeAnomalyDetectorss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type DiskEdgeAnomalyDetectors struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyName) || (has(self.initProvider) && has(self.initProvider.policyName))",message="spec.forProvider.policyName is a required parameter"
	Spec   DiskEdgeAnomalyDetectorsSpec   `json:"spec"`
	Status DiskEdgeAnomalyDetectorsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskEdgeAnomalyDetectorsList contains a list of DiskEdgeAnomalyDetectorss
type DiskEdgeAnomalyDetectorsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskEdgeAnomalyDetectors `json:"items"`
}

// Repository type metadata.
var (
	DiskEdgeAnomalyDetectors_Kind             = "DiskEdgeAnomalyDetectors"
	DiskEdgeAnomalyDetectors_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiskEdgeAnomalyDetectors_Kind}.String()
	DiskEdgeAnomalyDetectors_KindAPIVersion   = DiskEdgeAnomalyDetectors_Kind + "." + CRDGroupVersion.String()
	DiskEdgeAnomalyDetectors_GroupVersionKind = CRDGroupVersion.WithKind(DiskEdgeAnomalyDetectors_Kind)
)

func init() {
	SchemeBuilder.Register(&DiskEdgeAnomalyDetectors{}, &DiskEdgeAnomalyDetectorsList{})
}
