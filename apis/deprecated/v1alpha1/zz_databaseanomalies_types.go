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

type DBConnectFailuresInitParameters struct {

	// (Number) Number of failed database connections during any eval_period minutes period to trigger an alert
	// Number of failed database connections during any **eval_period** minutes period to trigger an alert
	ConnectionFailsCount *float64 `json:"connectionFailsCount,omitempty" tf:"connection_fails_count,omitempty"`

	// (Number) The X minutes time period during which the connection_fails_count is evaluated
	// The *X* minutes time period during which the **connection_fails_count** is evaluated
	EvalPeriod *float64 `json:"evalPeriod,omitempty" tf:"eval_period,omitempty"`
}

type DBConnectFailuresObservation struct {

	// (Number) Number of failed database connections during any eval_period minutes period to trigger an alert
	// Number of failed database connections during any **eval_period** minutes period to trigger an alert
	ConnectionFailsCount *float64 `json:"connectionFailsCount,omitempty" tf:"connection_fails_count,omitempty"`

	// (Number) The X minutes time period during which the connection_fails_count is evaluated
	// The *X* minutes time period during which the **connection_fails_count** is evaluated
	EvalPeriod *float64 `json:"evalPeriod,omitempty" tf:"eval_period,omitempty"`
}

type DBConnectFailuresParameters struct {

	// (Number) Number of failed database connections during any eval_period minutes period to trigger an alert
	// Number of failed database connections during any **eval_period** minutes period to trigger an alert
	// +kubebuilder:validation:Optional
	ConnectionFailsCount *float64 `json:"connectionFailsCount,omitempty" tf:"connection_fails_count,omitempty"`

	// (Number) The X minutes time period during which the connection_fails_count is evaluated
	// The *X* minutes time period during which the **connection_fails_count** is evaluated
	// +kubebuilder:validation:Optional
	EvalPeriod *float64 `json:"evalPeriod,omitempty" tf:"eval_period,omitempty"`
}

type DatabaseAnomaliesFailureRateInitParameters struct {

	// detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + 1% = 2.5%  Relative: 1.5% + 1.5% * 50% = 2.25% (see below for nested schema)
	// Parameters of failure rate increase auto-detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + **1%** = 2.5%  Relative: 1.5% + 1.5% * **50%** = 2.25%
	Auto []FailureRateAutoInitParameters `json:"auto,omitempty" tf:"auto,omitempty"`

	// (Block List, Max: 1) Fixed thresholds for failure rate increase detection (see below for nested schema)
	// Fixed thresholds for failure rate increase detection
	Thresholds []FailureRateThresholdsInitParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type DatabaseAnomaliesFailureRateObservation struct {

	// detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + 1% = 2.5%  Relative: 1.5% + 1.5% * 50% = 2.25% (see below for nested schema)
	// Parameters of failure rate increase auto-detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + **1%** = 2.5%  Relative: 1.5% + 1.5% * **50%** = 2.25%
	Auto []FailureRateAutoObservation `json:"auto,omitempty" tf:"auto,omitempty"`

	// (Block List, Max: 1) Fixed thresholds for failure rate increase detection (see below for nested schema)
	// Fixed thresholds for failure rate increase detection
	Thresholds []FailureRateThresholdsObservation `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type DatabaseAnomaliesFailureRateParameters struct {

	// detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + 1% = 2.5%  Relative: 1.5% + 1.5% * 50% = 2.25% (see below for nested schema)
	// Parameters of failure rate increase auto-detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + **1%** = 2.5%  Relative: 1.5% + 1.5% * **50%** = 2.25%
	// +kubebuilder:validation:Optional
	Auto []FailureRateAutoParameters `json:"auto,omitempty" tf:"auto,omitempty"`

	// (Block List, Max: 1) Fixed thresholds for failure rate increase detection (see below for nested schema)
	// Fixed thresholds for failure rate increase detection
	// +kubebuilder:validation:Optional
	Thresholds []FailureRateThresholdsParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type DatabaseAnomaliesInitParameters struct {

	// (Block List, Max: 1) Parameters of the failed database connections detection.  The alert is triggered when failed connections number exceeds connectionFailsCount during any timePeriodMinutes minutes period (see below for nested schema)
	// Parameters of the failed database connections detection.  The alert is triggered when failed connections number exceeds **connectionFailsCount** during any **timePeriodMinutes** minutes period
	DBConnectFailures []DBConnectFailuresInitParameters `json:"dbConnectFailures,omitempty" tf:"db_connect_failures,omitempty"`

	// (Block List, Max: 1) Configuration of failure rate increase detection (see below for nested schema)
	// Configuration of failure rate increase detection
	FailureRate []DatabaseAnomaliesFailureRateInitParameters `json:"failureRate,omitempty" tf:"failure_rate,omitempty"`

	// (Block List, Max: 1) Configuration for anomalies regarding load drops and spikes (see below for nested schema)
	// Configuration for anomalies regarding load drops and spikes
	Load []LoadInitParameters `json:"load,omitempty" tf:"load,omitempty"`

	// (Block List, Max: 1) Configuration of response time degradation detection (see below for nested schema)
	// Configuration of response time degradation detection
	ResponseTime []DatabaseAnomaliesResponseTimeInitParameters `json:"responseTime,omitempty" tf:"response_time,omitempty"`
}

type DatabaseAnomaliesObservation struct {

	// (Block List, Max: 1) Parameters of the failed database connections detection.  The alert is triggered when failed connections number exceeds connectionFailsCount during any timePeriodMinutes minutes period (see below for nested schema)
	// Parameters of the failed database connections detection.  The alert is triggered when failed connections number exceeds **connectionFailsCount** during any **timePeriodMinutes** minutes period
	DBConnectFailures []DBConnectFailuresObservation `json:"dbConnectFailures,omitempty" tf:"db_connect_failures,omitempty"`

	// (Block List, Max: 1) Configuration of failure rate increase detection (see below for nested schema)
	// Configuration of failure rate increase detection
	FailureRate []DatabaseAnomaliesFailureRateObservation `json:"failureRate,omitempty" tf:"failure_rate,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Max: 1) Configuration for anomalies regarding load drops and spikes (see below for nested schema)
	// Configuration for anomalies regarding load drops and spikes
	Load []LoadObservation `json:"load,omitempty" tf:"load,omitempty"`

	// (Block List, Max: 1) Configuration of response time degradation detection (see below for nested schema)
	// Configuration of response time degradation detection
	ResponseTime []DatabaseAnomaliesResponseTimeObservation `json:"responseTime,omitempty" tf:"response_time,omitempty"`
}

type DatabaseAnomaliesParameters struct {

	// (Block List, Max: 1) Parameters of the failed database connections detection.  The alert is triggered when failed connections number exceeds connectionFailsCount during any timePeriodMinutes minutes period (see below for nested schema)
	// Parameters of the failed database connections detection.  The alert is triggered when failed connections number exceeds **connectionFailsCount** during any **timePeriodMinutes** minutes period
	// +kubebuilder:validation:Optional
	DBConnectFailures []DBConnectFailuresParameters `json:"dbConnectFailures,omitempty" tf:"db_connect_failures,omitempty"`

	// (Block List, Max: 1) Configuration of failure rate increase detection (see below for nested schema)
	// Configuration of failure rate increase detection
	// +kubebuilder:validation:Optional
	FailureRate []DatabaseAnomaliesFailureRateParameters `json:"failureRate,omitempty" tf:"failure_rate,omitempty"`

	// (Block List, Max: 1) Configuration for anomalies regarding load drops and spikes (see below for nested schema)
	// Configuration for anomalies regarding load drops and spikes
	// +kubebuilder:validation:Optional
	Load []LoadParameters `json:"load,omitempty" tf:"load,omitempty"`

	// (Block List, Max: 1) Configuration of response time degradation detection (see below for nested schema)
	// Configuration of response time degradation detection
	// +kubebuilder:validation:Optional
	ResponseTime []DatabaseAnomaliesResponseTimeParameters `json:"responseTime,omitempty" tf:"response_time,omitempty"`
}

type DatabaseAnomaliesResponseTimeAutoInitParameters struct {

	// (Block List, Max: 1) Configuration for anomalies regarding load drops and spikes (see below for nested schema)
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

type DatabaseAnomaliesResponseTimeAutoObservation struct {

	// (Block List, Max: 1) Configuration for anomalies regarding load drops and spikes (see below for nested schema)
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

type DatabaseAnomaliesResponseTimeAutoParameters struct {

	// (Block List, Max: 1) Configuration for anomalies regarding load drops and spikes (see below for nested schema)
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

type DatabaseAnomaliesResponseTimeInitParameters struct {

	// detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + 1% = 2.5%  Relative: 1.5% + 1.5% * 50% = 2.25% (see below for nested schema)
	// Parameters of the response time degradation auto-detection. Violation of **any** criterion triggers an alert
	Auto []DatabaseAnomaliesResponseTimeAutoInitParameters `json:"auto,omitempty" tf:"auto,omitempty"`

	// (Block List, Max: 1) Fixed thresholds for failure rate increase detection (see below for nested schema)
	// Fixed thresholds for response time degradation detection
	Thresholds []DatabaseAnomaliesResponseTimeThresholdsInitParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type DatabaseAnomaliesResponseTimeObservation struct {

	// detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + 1% = 2.5%  Relative: 1.5% + 1.5% * 50% = 2.25% (see below for nested schema)
	// Parameters of the response time degradation auto-detection. Violation of **any** criterion triggers an alert
	Auto []DatabaseAnomaliesResponseTimeAutoObservation `json:"auto,omitempty" tf:"auto,omitempty"`

	// (Block List, Max: 1) Fixed thresholds for failure rate increase detection (see below for nested schema)
	// Fixed thresholds for response time degradation detection
	Thresholds []DatabaseAnomaliesResponseTimeThresholdsObservation `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type DatabaseAnomaliesResponseTimeParameters struct {

	// detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + 1% = 2.5%  Relative: 1.5% + 1.5% * 50% = 2.25% (see below for nested schema)
	// Parameters of the response time degradation auto-detection. Violation of **any** criterion triggers an alert
	// +kubebuilder:validation:Optional
	Auto []DatabaseAnomaliesResponseTimeAutoParameters `json:"auto,omitempty" tf:"auto,omitempty"`

	// (Block List, Max: 1) Fixed thresholds for failure rate increase detection (see below for nested schema)
	// Fixed thresholds for response time degradation detection
	// +kubebuilder:validation:Optional
	Thresholds []DatabaseAnomaliesResponseTimeThresholdsParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`
}

type DatabaseAnomaliesResponseTimeThresholdsInitParameters struct {

	// (Block List, Max: 1) Configuration for anomalies regarding load drops and spikes (see below for nested schema)
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

type DatabaseAnomaliesResponseTimeThresholdsObservation struct {

	// (Block List, Max: 1) Configuration for anomalies regarding load drops and spikes (see below for nested schema)
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

type DatabaseAnomaliesResponseTimeThresholdsParameters struct {

	// (Block List, Max: 1) Configuration for anomalies regarding load drops and spikes (see below for nested schema)
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

type FailureRateAutoInitParameters struct {

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

type FailureRateAutoObservation struct {

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

type FailureRateAutoParameters struct {

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

type FailureRateThresholdsInitParameters struct {

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

type FailureRateThresholdsObservation struct {

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

type FailureRateThresholdsParameters struct {

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

type LoadDropsInitParameters struct {

	// (Number) Alert if the service stays in abnormal state for at least X minutes
	// Alert if the service stays in abnormal state for at least *X* minutes
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the observed load is more than *X* % of the expected value
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`
}

type LoadDropsObservation struct {

	// (Number) Alert if the service stays in abnormal state for at least X minutes
	// Alert if the service stays in abnormal state for at least *X* minutes
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the observed load is more than *X* % of the expected value
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`
}

type LoadDropsParameters struct {

	// (Number) Alert if the service stays in abnormal state for at least X minutes
	// Alert if the service stays in abnormal state for at least *X* minutes
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// (Number) Alert if the observed load is more than X % of the expected value
	// Alert if the observed load is more than *X* % of the expected value
	// +kubebuilder:validation:Optional
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`
}

type LoadInitParameters struct {

	// (Block List, Max: 1) The configuration of traffic drops detection (see below for nested schema)
	// The configuration of traffic drops detection
	Drops []LoadDropsInitParameters `json:"drops,omitempty" tf:"drops,omitempty"`

	// (Block List, Max: 1) The configuration of traffic spikes detection (see below for nested schema)
	// The configuration of traffic spikes detection
	Spikes []LoadSpikesInitParameters `json:"spikes,omitempty" tf:"spikes,omitempty"`
}

type LoadObservation struct {

	// (Block List, Max: 1) The configuration of traffic drops detection (see below for nested schema)
	// The configuration of traffic drops detection
	Drops []LoadDropsObservation `json:"drops,omitempty" tf:"drops,omitempty"`

	// (Block List, Max: 1) The configuration of traffic spikes detection (see below for nested schema)
	// The configuration of traffic spikes detection
	Spikes []LoadSpikesObservation `json:"spikes,omitempty" tf:"spikes,omitempty"`
}

type LoadParameters struct {

	// (Block List, Max: 1) The configuration of traffic drops detection (see below for nested schema)
	// The configuration of traffic drops detection
	// +kubebuilder:validation:Optional
	Drops []LoadDropsParameters `json:"drops,omitempty" tf:"drops,omitempty"`

	// (Block List, Max: 1) The configuration of traffic spikes detection (see below for nested schema)
	// The configuration of traffic spikes detection
	// +kubebuilder:validation:Optional
	Spikes []LoadSpikesParameters `json:"spikes,omitempty" tf:"spikes,omitempty"`
}

type LoadSpikesInitParameters struct {

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

type LoadSpikesObservation struct {

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

type LoadSpikesParameters struct {

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

// DatabaseAnomaliesSpec defines the desired state of DatabaseAnomalies
type DatabaseAnomaliesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseAnomaliesParameters `json:"forProvider"`
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
	InitProvider DatabaseAnomaliesInitParameters `json:"initProvider,omitempty"`
}

// DatabaseAnomaliesStatus defines the observed state of DatabaseAnomalies.
type DatabaseAnomaliesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseAnomaliesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DatabaseAnomalies is the Schema for the DatabaseAnomaliess API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type DatabaseAnomalies struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAnomaliesSpec   `json:"spec"`
	Status            DatabaseAnomaliesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseAnomaliesList contains a list of DatabaseAnomaliess
type DatabaseAnomaliesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseAnomalies `json:"items"`
}

// Repository type metadata.
var (
	DatabaseAnomalies_Kind             = "DatabaseAnomalies"
	DatabaseAnomalies_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseAnomalies_Kind}.String()
	DatabaseAnomalies_KindAPIVersion   = DatabaseAnomalies_Kind + "." + CRDGroupVersion.String()
	DatabaseAnomalies_GroupVersionKind = CRDGroupVersion.WithKind(DatabaseAnomalies_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabaseAnomalies{}, &DatabaseAnomaliesList{})
}
