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

type FailureRatesAutoInitParameters struct {

	// (Number) Absolute increase of failing service calls to trigger an alert, %
	// Absolute increase of failing service calls to trigger an alert, %
	Absolute *float64 `json:"absolute,omitempty" tf:"absolute,omitempty"`

	// (Number) Relative increase of failing service calls to trigger an alert, %
	// Relative increase of failing service calls to trigger an alert, %
	Relative *float64 `json:"relative,omitempty" tf:"relative,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type FailureRatesAutoObservation struct {

	// (Number) Absolute increase of failing service calls to trigger an alert, %
	// Absolute increase of failing service calls to trigger an alert, %
	Absolute *float64 `json:"absolute,omitempty" tf:"absolute,omitempty"`

	// (Number) Relative increase of failing service calls to trigger an alert, %
	// Relative increase of failing service calls to trigger an alert, %
	Relative *float64 `json:"relative,omitempty" tf:"relative,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type FailureRatesAutoParameters struct {

	// (Number) Absolute increase of failing service calls to trigger an alert, %
	// Absolute increase of failing service calls to trigger an alert, %
	// +kubebuilder:validation:Optional
	Absolute *float64 `json:"absolute" tf:"absolute,omitempty"`

	// (Number) Relative increase of failing service calls to trigger an alert, %
	// Relative increase of failing service calls to trigger an alert, %
	// +kubebuilder:validation:Optional
	Relative *float64 `json:"relative" tf:"relative,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type FailureRatesInitParameters struct {

	// detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + 1% = 2.5%  Relative: 1.5% + 1.5% * 50% = 2.25% (see below for nested schema)
	// Parameters of failure rate increase auto-detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + **1%** = 2.5%  Relative: 1.5% + 1.5% * **50%** = 2.25%
	Auto []FailureRatesAutoInitParameters `json:"auto,omitempty" tf:"auto,omitempty"`

	// (Block List, Max: 1) Fixed thresholds for failure rate increase detection (see below for nested schema)
	// Fixed thresholds for failure rate increase detection
	Thresholds []FailureRatesThresholdsInitParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type FailureRatesObservation struct {

	// detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + 1% = 2.5%  Relative: 1.5% + 1.5% * 50% = 2.25% (see below for nested schema)
	// Parameters of failure rate increase auto-detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + **1%** = 2.5%  Relative: 1.5% + 1.5% * **50%** = 2.25%
	Auto []FailureRatesAutoObservation `json:"auto,omitempty" tf:"auto,omitempty"`

	// (Block List, Max: 1) Fixed thresholds for failure rate increase detection (see below for nested schema)
	// Fixed thresholds for failure rate increase detection
	Thresholds []FailureRatesThresholdsObservation `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type FailureRatesParameters struct {

	// detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + 1% = 2.5%  Relative: 1.5% + 1.5% * 50% = 2.25% (see below for nested schema)
	// Parameters of failure rate increase auto-detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + **1%** = 2.5%  Relative: 1.5% + 1.5% * **50%** = 2.25%
	// +kubebuilder:validation:Optional
	Auto []FailureRatesAutoParameters `json:"auto,omitempty" tf:"auto,omitempty"`

	// (Block List, Max: 1) Fixed thresholds for failure rate increase detection (see below for nested schema)
	// Fixed thresholds for failure rate increase detection
	// +kubebuilder:validation:Optional
	Thresholds []FailureRatesThresholdsParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type FailureRatesThresholdsInitParameters struct {

	// (String) Sensitivity of the threshold.  With low sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With high sensitivity, no statistical confidence is used. Each violation triggers alert
	// Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers alert
	Sensitivity *string `json:"sensitivity,omitempty" tf:"sensitivity,omitempty"`

	// minute period to trigger an alert, %
	// Failure rate during any 5-minute period to trigger an alert, %
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type FailureRatesThresholdsObservation struct {

	// (String) Sensitivity of the threshold.  With low sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With high sensitivity, no statistical confidence is used. Each violation triggers alert
	// Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers alert
	Sensitivity *string `json:"sensitivity,omitempty" tf:"sensitivity,omitempty"`

	// minute period to trigger an alert, %
	// Failure rate during any 5-minute period to trigger an alert, %
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type FailureRatesThresholdsParameters struct {

	// (String) Sensitivity of the threshold.  With low sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With high sensitivity, no statistical confidence is used. Each violation triggers alert
	// Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers alert
	// +kubebuilder:validation:Optional
	Sensitivity *string `json:"sensitivity" tf:"sensitivity,omitempty"`

	// minute period to trigger an alert, %
	// Failure rate during any 5-minute period to trigger an alert, %
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type ResponseTimesAutoInitParameters struct {

	// (Block List, Max: 1) The configuration of load spikes detection. Detecting load spikes will be disabled if this block is omitted. (see below for nested schema)
	// Minimal service load to detect response time degradation. Response time degradation of services with smaller load won't trigger alerts. Possible values are `FIFTEEN_REQUESTS_PER_MINUTE`, `FIVE_REQUESTS_PER_MINUTE`, `ONE_REQUEST_PER_MINUTE` and `TEN_REQUESTS_PER_MINUTE`
	Load *string `json:"load,omitempty" tf:"load,omitempty"`

	// (Number) Alert if the response time degrades by more than X milliseconds
	// Alert if the response time degrades by more than *X* milliseconds
	Milliseconds *float64 `json:"milliseconds,omitempty" tf:"milliseconds,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the response time degrades by more than *X* %
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`

	// (Number) Alert if the response time of the slowest 10% degrades by more than X milliseconds
	// Alert if the response time of the slowest 10% degrades by more than *X* milliseconds
	SlowestMilliseconds *float64 `json:"slowestMilliseconds,omitempty" tf:"slowest_milliseconds,omitempty"`

	// (Number) Alert if the response time of the slowest 10% degrades by more than X milliseconds
	// Alert if the response time of the slowest 10% degrades by more than *X* milliseconds
	SlowestPercent *float64 `json:"slowestPercent,omitempty" tf:"slowest_percent,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type ResponseTimesAutoObservation struct {

	// (Block List, Max: 1) The configuration of load spikes detection. Detecting load spikes will be disabled if this block is omitted. (see below for nested schema)
	// Minimal service load to detect response time degradation. Response time degradation of services with smaller load won't trigger alerts. Possible values are `FIFTEEN_REQUESTS_PER_MINUTE`, `FIVE_REQUESTS_PER_MINUTE`, `ONE_REQUEST_PER_MINUTE` and `TEN_REQUESTS_PER_MINUTE`
	Load *string `json:"load,omitempty" tf:"load,omitempty"`

	// (Number) Alert if the response time degrades by more than X milliseconds
	// Alert if the response time degrades by more than *X* milliseconds
	Milliseconds *float64 `json:"milliseconds,omitempty" tf:"milliseconds,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the response time degrades by more than *X* %
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`

	// (Number) Alert if the response time of the slowest 10% degrades by more than X milliseconds
	// Alert if the response time of the slowest 10% degrades by more than *X* milliseconds
	SlowestMilliseconds *float64 `json:"slowestMilliseconds,omitempty" tf:"slowest_milliseconds,omitempty"`

	// (Number) Alert if the response time of the slowest 10% degrades by more than X milliseconds
	// Alert if the response time of the slowest 10% degrades by more than *X* milliseconds
	SlowestPercent *float64 `json:"slowestPercent,omitempty" tf:"slowest_percent,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type ResponseTimesAutoParameters struct {

	// (Block List, Max: 1) The configuration of load spikes detection. Detecting load spikes will be disabled if this block is omitted. (see below for nested schema)
	// Minimal service load to detect response time degradation. Response time degradation of services with smaller load won't trigger alerts. Possible values are `FIFTEEN_REQUESTS_PER_MINUTE`, `FIVE_REQUESTS_PER_MINUTE`, `ONE_REQUEST_PER_MINUTE` and `TEN_REQUESTS_PER_MINUTE`
	// +kubebuilder:validation:Optional
	Load *string `json:"load" tf:"load,omitempty"`

	// (Number) Alert if the response time degrades by more than X milliseconds
	// Alert if the response time degrades by more than *X* milliseconds
	// +kubebuilder:validation:Optional
	Milliseconds *float64 `json:"milliseconds" tf:"milliseconds,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the response time degrades by more than *X* %
	// +kubebuilder:validation:Optional
	Percent *float64 `json:"percent" tf:"percent,omitempty"`

	// (Number) Alert if the response time of the slowest 10% degrades by more than X milliseconds
	// Alert if the response time of the slowest 10% degrades by more than *X* milliseconds
	// +kubebuilder:validation:Optional
	SlowestMilliseconds *float64 `json:"slowestMilliseconds" tf:"slowest_milliseconds,omitempty"`

	// (Number) Alert if the response time of the slowest 10% degrades by more than X milliseconds
	// Alert if the response time of the slowest 10% degrades by more than *X* milliseconds
	// +kubebuilder:validation:Optional
	SlowestPercent *float64 `json:"slowestPercent" tf:"slowest_percent,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type ResponseTimesInitParameters struct {

	// detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + 1% = 2.5%  Relative: 1.5% + 1.5% * 50% = 2.25% (see below for nested schema)
	// Parameters of the response time degradation auto-detection. Violation of **any** criterion triggers an alert
	Auto []ResponseTimesAutoInitParameters `json:"auto,omitempty" tf:"auto,omitempty"`

	// (Block List, Max: 1) Fixed thresholds for failure rate increase detection (see below for nested schema)
	// Fixed thresholds for response time degradation detection
	Thresholds []ResponseTimesThresholdsInitParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type ResponseTimesObservation struct {

	// detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + 1% = 2.5%  Relative: 1.5% + 1.5% * 50% = 2.25% (see below for nested schema)
	// Parameters of the response time degradation auto-detection. Violation of **any** criterion triggers an alert
	Auto []ResponseTimesAutoObservation `json:"auto,omitempty" tf:"auto,omitempty"`

	// (Block List, Max: 1) Fixed thresholds for failure rate increase detection (see below for nested schema)
	// Fixed thresholds for response time degradation detection
	Thresholds []ResponseTimesThresholdsObservation `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type ResponseTimesParameters struct {

	// detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + 1% = 2.5%  Relative: 1.5% + 1.5% * 50% = 2.25% (see below for nested schema)
	// Parameters of the response time degradation auto-detection. Violation of **any** criterion triggers an alert
	// +kubebuilder:validation:Optional
	Auto []ResponseTimesAutoParameters `json:"auto,omitempty" tf:"auto,omitempty"`

	// (Block List, Max: 1) Fixed thresholds for failure rate increase detection (see below for nested schema)
	// Fixed thresholds for response time degradation detection
	// +kubebuilder:validation:Optional
	Thresholds []ResponseTimesThresholdsParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type ResponseTimesThresholdsInitParameters struct {

	// (Block List, Max: 1) The configuration of load spikes detection. Detecting load spikes will be disabled if this block is omitted. (see below for nested schema)
	// Minimal service load to detect response time degradation. Response time degradation of services with smaller load won't trigger alerts. Possible values are `FIFTEEN_REQUESTS_PER_MINUTE`, `FIVE_REQUESTS_PER_MINUTE`, `ONE_REQUEST_PER_MINUTE` and `TEN_REQUESTS_PER_MINUTE`
	Load *string `json:"load,omitempty" tf:"load,omitempty"`

	// (Number) Alert if the response time degrades by more than X milliseconds
	// Response time during any 5-minute period to trigger an alert, in milliseconds
	Milliseconds *float64 `json:"milliseconds,omitempty" tf:"milliseconds,omitempty"`

	// (String) Sensitivity of the threshold.  With low sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With high sensitivity, no statistical confidence is used. Each violation triggers alert
	// Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers an alert
	Sensitivity *string `json:"sensitivity,omitempty" tf:"sensitivity,omitempty"`

	// (Number) Alert if the response time of the slowest 10% degrades by more than X milliseconds
	// Response time of the 10% slowest during any 5-minute period to trigger an alert, in milliseconds
	SlowestMilliseconds *float64 `json:"slowestMilliseconds,omitempty" tf:"slowest_milliseconds,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type ResponseTimesThresholdsObservation struct {

	// (Block List, Max: 1) The configuration of load spikes detection. Detecting load spikes will be disabled if this block is omitted. (see below for nested schema)
	// Minimal service load to detect response time degradation. Response time degradation of services with smaller load won't trigger alerts. Possible values are `FIFTEEN_REQUESTS_PER_MINUTE`, `FIVE_REQUESTS_PER_MINUTE`, `ONE_REQUEST_PER_MINUTE` and `TEN_REQUESTS_PER_MINUTE`
	Load *string `json:"load,omitempty" tf:"load,omitempty"`

	// (Number) Alert if the response time degrades by more than X milliseconds
	// Response time during any 5-minute period to trigger an alert, in milliseconds
	Milliseconds *float64 `json:"milliseconds,omitempty" tf:"milliseconds,omitempty"`

	// (String) Sensitivity of the threshold.  With low sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With high sensitivity, no statistical confidence is used. Each violation triggers alert
	// Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers an alert
	Sensitivity *string `json:"sensitivity,omitempty" tf:"sensitivity,omitempty"`

	// (Number) Alert if the response time of the slowest 10% degrades by more than X milliseconds
	// Response time of the 10% slowest during any 5-minute period to trigger an alert, in milliseconds
	SlowestMilliseconds *float64 `json:"slowestMilliseconds,omitempty" tf:"slowest_milliseconds,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type ResponseTimesThresholdsParameters struct {

	// (Block List, Max: 1) The configuration of load spikes detection. Detecting load spikes will be disabled if this block is omitted. (see below for nested schema)
	// Minimal service load to detect response time degradation. Response time degradation of services with smaller load won't trigger alerts. Possible values are `FIFTEEN_REQUESTS_PER_MINUTE`, `FIVE_REQUESTS_PER_MINUTE`, `ONE_REQUEST_PER_MINUTE` and `TEN_REQUESTS_PER_MINUTE`
	// +kubebuilder:validation:Optional
	Load *string `json:"load" tf:"load,omitempty"`

	// (Number) Alert if the response time degrades by more than X milliseconds
	// Response time during any 5-minute period to trigger an alert, in milliseconds
	// +kubebuilder:validation:Optional
	Milliseconds *float64 `json:"milliseconds" tf:"milliseconds,omitempty"`

	// (String) Sensitivity of the threshold.  With low sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With high sensitivity, no statistical confidence is used. Each violation triggers alert
	// Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers an alert
	// +kubebuilder:validation:Optional
	Sensitivity *string `json:"sensitivity" tf:"sensitivity,omitempty"`

	// (Number) Alert if the response time of the slowest 10% degrades by more than X milliseconds
	// Response time of the 10% slowest during any 5-minute period to trigger an alert, in milliseconds
	// +kubebuilder:validation:Optional
	SlowestMilliseconds *float64 `json:"slowestMilliseconds" tf:"slowest_milliseconds,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type ServiceAnomaliesInitParameters struct {

	// (Block List, Max: 1) Configuration of failure rate increase detection. Detecting failure rate increases will be disabled if this block is omitted. (see below for nested schema)
	// Configuration of failure rate increase detection. Detecting failure rate increases will be disabled if this block is omitted.
	FailureRates []FailureRatesInitParameters `json:"failureRates,omitempty" tf:"failure_rates,omitempty"`

	// (Block List, Max: 1) The configuration of load spikes detection. Detecting load spikes will be disabled if this block is omitted. (see below for nested schema)
	// The configuration of load spikes detection. Detecting load spikes will be disabled if this block is omitted.
	Load []ServiceAnomaliesLoadInitParameters `json:"load,omitempty" tf:"load,omitempty"`

	// (Block List, Max: 1) The configuration of load drops detection. Detecting load drops will be disabled if this block is omitted. (see below for nested schema)
	// The configuration of load drops detection. Detecting load drops will be disabled if this block is omitted.
	LoadDrops []ServiceAnomaliesLoadDropsInitParameters_2 `json:"loadDrops,omitempty" tf:"load_drops,omitempty"`

	// (Block List, Max: 1) Configuration of response time degradation detection. Detecting response time degradation will be disabled if this block is omitted. (see below for nested schema)
	// Configuration of response time degradation detection. Detecting response time degradation will be disabled if this block is omitted.
	ResponseTimes []ResponseTimesInitParameters `json:"responseTimes,omitempty" tf:"response_times,omitempty"`
}

type ServiceAnomaliesLoadDropsInitParameters struct {

	// (Number) Alert if the service stays in abnormal state for at least X minutes
	// Alert if the service stays in abnormal state for at least *X* minutes
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the observed load is more than *X* % of the expected value
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`
}

type ServiceAnomaliesLoadDropsInitParameters_2 struct {

	// (Number) Alert if the service stays in abnormal state for at least X minutes
	// Alert if the service stays in abnormal state for at least *X* minutes
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the observed load is more than *X* % of the expected value
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`
}

type ServiceAnomaliesLoadDropsObservation struct {

	// (Number) Alert if the service stays in abnormal state for at least X minutes
	// Alert if the service stays in abnormal state for at least *X* minutes
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the observed load is more than *X* % of the expected value
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`
}

type ServiceAnomaliesLoadDropsObservation_2 struct {

	// (Number) Alert if the service stays in abnormal state for at least X minutes
	// Alert if the service stays in abnormal state for at least *X* minutes
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the observed load is more than *X* % of the expected value
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`
}

type ServiceAnomaliesLoadDropsParameters struct {

	// (Number) Alert if the service stays in abnormal state for at least X minutes
	// Alert if the service stays in abnormal state for at least *X* minutes
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the observed load is more than *X* % of the expected value
	// +kubebuilder:validation:Optional
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`
}

type ServiceAnomaliesLoadDropsParameters_2 struct {

	// (Number) Alert if the service stays in abnormal state for at least X minutes
	// Alert if the service stays in abnormal state for at least *X* minutes
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the observed load is more than *X* % of the expected value
	// +kubebuilder:validation:Optional
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`
}

type ServiceAnomaliesLoadInitParameters struct {

	// (Block List, Max: 1) The configuration of traffic drops detection (see below for nested schema)
	// The configuration of traffic drops detection
	Drops []ServiceAnomaliesLoadDropsInitParameters `json:"drops,omitempty" tf:"drops,omitempty"`

	// (Block List, Max: 1) The configuration of traffic spikes detection (see below for nested schema)
	// The configuration of traffic spikes detection
	Spikes []ServiceAnomaliesLoadSpikesInitParameters `json:"spikes,omitempty" tf:"spikes,omitempty"`
}

type ServiceAnomaliesLoadObservation struct {

	// (Block List, Max: 1) The configuration of traffic drops detection (see below for nested schema)
	// The configuration of traffic drops detection
	Drops []ServiceAnomaliesLoadDropsObservation `json:"drops,omitempty" tf:"drops,omitempty"`

	// (Block List, Max: 1) The configuration of traffic spikes detection (see below for nested schema)
	// The configuration of traffic spikes detection
	Spikes []ServiceAnomaliesLoadSpikesObservation `json:"spikes,omitempty" tf:"spikes,omitempty"`
}

type ServiceAnomaliesLoadParameters struct {

	// (Block List, Max: 1) The configuration of traffic drops detection (see below for nested schema)
	// The configuration of traffic drops detection
	// +kubebuilder:validation:Optional
	Drops []ServiceAnomaliesLoadDropsParameters `json:"drops,omitempty" tf:"drops,omitempty"`

	// (Block List, Max: 1) The configuration of traffic spikes detection (see below for nested schema)
	// The configuration of traffic spikes detection
	// +kubebuilder:validation:Optional
	Spikes []ServiceAnomaliesLoadSpikesParameters `json:"spikes,omitempty" tf:"spikes,omitempty"`
}

type ServiceAnomaliesLoadSpikesInitParameters struct {

	// (Number) Alert if the service stays in abnormal state for at least X minutes
	// Alert if the service stays in abnormal state for at least *X* minutes
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the observed load is more than *X* % of the expected value
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type ServiceAnomaliesLoadSpikesObservation struct {

	// (Number) Alert if the service stays in abnormal state for at least X minutes
	// Alert if the service stays in abnormal state for at least *X* minutes
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the observed load is more than *X* % of the expected value
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type ServiceAnomaliesLoadSpikesParameters struct {

	// (Number) Alert if the service stays in abnormal state for at least X minutes
	// Alert if the service stays in abnormal state for at least *X* minutes
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the observed load is more than *X* % of the expected value
	// +kubebuilder:validation:Optional
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type ServiceAnomaliesObservation struct {

	// (Block List, Max: 1) Configuration of failure rate increase detection. Detecting failure rate increases will be disabled if this block is omitted. (see below for nested schema)
	// Configuration of failure rate increase detection. Detecting failure rate increases will be disabled if this block is omitted.
	FailureRates []FailureRatesObservation `json:"failureRates,omitempty" tf:"failure_rates,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Max: 1) The configuration of load spikes detection. Detecting load spikes will be disabled if this block is omitted. (see below for nested schema)
	// The configuration of load spikes detection. Detecting load spikes will be disabled if this block is omitted.
	Load []ServiceAnomaliesLoadObservation `json:"load,omitempty" tf:"load,omitempty"`

	// (Block List, Max: 1) The configuration of load drops detection. Detecting load drops will be disabled if this block is omitted. (see below for nested schema)
	// The configuration of load drops detection. Detecting load drops will be disabled if this block is omitted.
	LoadDrops []ServiceAnomaliesLoadDropsObservation_2 `json:"loadDrops,omitempty" tf:"load_drops,omitempty"`

	// (Block List, Max: 1) Configuration of response time degradation detection. Detecting response time degradation will be disabled if this block is omitted. (see below for nested schema)
	// Configuration of response time degradation detection. Detecting response time degradation will be disabled if this block is omitted.
	ResponseTimes []ResponseTimesObservation `json:"responseTimes,omitempty" tf:"response_times,omitempty"`
}

type ServiceAnomaliesParameters struct {

	// (Block List, Max: 1) Configuration of failure rate increase detection. Detecting failure rate increases will be disabled if this block is omitted. (see below for nested schema)
	// Configuration of failure rate increase detection. Detecting failure rate increases will be disabled if this block is omitted.
	// +kubebuilder:validation:Optional
	FailureRates []FailureRatesParameters `json:"failureRates,omitempty" tf:"failure_rates,omitempty"`

	// (Block List, Max: 1) The configuration of load spikes detection. Detecting load spikes will be disabled if this block is omitted. (see below for nested schema)
	// The configuration of load spikes detection. Detecting load spikes will be disabled if this block is omitted.
	// +kubebuilder:validation:Optional
	Load []ServiceAnomaliesLoadParameters `json:"load,omitempty" tf:"load,omitempty"`

	// (Block List, Max: 1) The configuration of load drops detection. Detecting load drops will be disabled if this block is omitted. (see below for nested schema)
	// The configuration of load drops detection. Detecting load drops will be disabled if this block is omitted.
	// +kubebuilder:validation:Optional
	LoadDrops []ServiceAnomaliesLoadDropsParameters_2 `json:"loadDrops,omitempty" tf:"load_drops,omitempty"`

	// (Block List, Max: 1) Configuration of response time degradation detection. Detecting response time degradation will be disabled if this block is omitted. (see below for nested schema)
	// Configuration of response time degradation detection. Detecting response time degradation will be disabled if this block is omitted.
	// +kubebuilder:validation:Optional
	ResponseTimes []ResponseTimesParameters `json:"responseTimes,omitempty" tf:"response_times,omitempty"`
}

// ServiceAnomaliesSpec defines the desired state of ServiceAnomalies
type ServiceAnomaliesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAnomaliesParameters `json:"forProvider"`
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
	InitProvider ServiceAnomaliesInitParameters `json:"initProvider,omitempty"`
}

// ServiceAnomaliesStatus defines the observed state of ServiceAnomalies.
type ServiceAnomaliesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAnomaliesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServiceAnomalies is the Schema for the ServiceAnomaliess API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type ServiceAnomalies struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAnomaliesSpec   `json:"spec"`
	Status            ServiceAnomaliesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAnomaliesList contains a list of ServiceAnomaliess
type ServiceAnomaliesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAnomalies `json:"items"`
}

// Repository type metadata.
var (
	ServiceAnomalies_Kind             = "ServiceAnomalies"
	ServiceAnomalies_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAnomalies_Kind}.String()
	ServiceAnomalies_KindAPIVersion   = ServiceAnomalies_Kind + "." + CRDGroupVersion.String()
	ServiceAnomalies_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAnomalies_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAnomalies{}, &ServiceAnomaliesList{})
}
