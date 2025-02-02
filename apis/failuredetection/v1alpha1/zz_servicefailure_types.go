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

type CustomErrorRuleConditionInitParameters struct {

	// (Boolean) Case sensitive
	// Case sensitive
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// (String) Apply this comparison
	// Apply this comparison
	CompareOperationType *string `json:"compareOperationType,omitempty" tf:"compare_operation_type,omitempty"`

	// (Number) Value
	// Value
	DoubleValue *float64 `json:"doubleValue,omitempty" tf:"double_value,omitempty"`

	// (Number) Value
	// Value
	IntValue *float64 `json:"intValue,omitempty" tf:"int_value,omitempty"`

	// (String) Value
	// Value
	TextValue *string `json:"textValue,omitempty" tf:"text_value,omitempty"`
}

type CustomErrorRuleConditionObservation struct {

	// (Boolean) Case sensitive
	// Case sensitive
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// (String) Apply this comparison
	// Apply this comparison
	CompareOperationType *string `json:"compareOperationType,omitempty" tf:"compare_operation_type,omitempty"`

	// (Number) Value
	// Value
	DoubleValue *float64 `json:"doubleValue,omitempty" tf:"double_value,omitempty"`

	// (Number) Value
	// Value
	IntValue *float64 `json:"intValue,omitempty" tf:"int_value,omitempty"`

	// (String) Value
	// Value
	TextValue *string `json:"textValue,omitempty" tf:"text_value,omitempty"`
}

type CustomErrorRuleConditionParameters struct {

	// (Boolean) Case sensitive
	// Case sensitive
	// +kubebuilder:validation:Optional
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`

	// (String) Apply this comparison
	// Apply this comparison
	// +kubebuilder:validation:Optional
	CompareOperationType *string `json:"compareOperationType" tf:"compare_operation_type,omitempty"`

	// (Number) Value
	// Value
	// +kubebuilder:validation:Optional
	DoubleValue *float64 `json:"doubleValue,omitempty" tf:"double_value,omitempty"`

	// (Number) Value
	// Value
	// +kubebuilder:validation:Optional
	IntValue *float64 `json:"intValue,omitempty" tf:"int_value,omitempty"`

	// (String) Value
	// Value
	// +kubebuilder:validation:Optional
	TextValue *string `json:"textValue,omitempty" tf:"text_value,omitempty"`
}

type CustomErrorRulesCustomErrorRuleInitParameters struct {

	// (Block List, Min: 1, Max: 1) Request attribute condition (see below for nested schema)
	// Request attribute condition
	Condition []CustomErrorRuleConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// (String) Request attribute
	// Request attribute
	RequestAttribute *string `json:"requestAttribute,omitempty" tf:"request_attribute,omitempty"`
}

type CustomErrorRulesCustomErrorRuleObservation struct {

	// (Block List, Min: 1, Max: 1) Request attribute condition (see below for nested schema)
	// Request attribute condition
	Condition []CustomErrorRuleConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// (String) Request attribute
	// Request attribute
	RequestAttribute *string `json:"requestAttribute,omitempty" tf:"request_attribute,omitempty"`
}

type CustomErrorRulesCustomErrorRuleParameters struct {

	// (Block List, Min: 1, Max: 1) Request attribute condition (see below for nested schema)
	// Request attribute condition
	// +kubebuilder:validation:Optional
	Condition []CustomErrorRuleConditionParameters `json:"condition" tf:"condition,omitempty"`

	// (String) Request attribute
	// Request attribute
	// +kubebuilder:validation:Optional
	RequestAttribute *string `json:"requestAttribute" tf:"request_attribute,omitempty"`
}

type CustomHandledExceptionsCustomHandledExceptionInitParameters struct {

	// (String) The pattern will match if it is contained within the actual class name.
	// The pattern will match if it is contained within the actual class name.
	ClassPattern *string `json:"classPattern,omitempty" tf:"class_pattern,omitempty"`

	// (String) Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	// Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	MessagePattern *string `json:"messagePattern,omitempty" tf:"message_pattern,omitempty"`
}

type CustomHandledExceptionsCustomHandledExceptionObservation struct {

	// (String) The pattern will match if it is contained within the actual class name.
	// The pattern will match if it is contained within the actual class name.
	ClassPattern *string `json:"classPattern,omitempty" tf:"class_pattern,omitempty"`

	// (String) Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	// Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	MessagePattern *string `json:"messagePattern,omitempty" tf:"message_pattern,omitempty"`
}

type CustomHandledExceptionsCustomHandledExceptionParameters struct {

	// (String) The pattern will match if it is contained within the actual class name.
	// The pattern will match if it is contained within the actual class name.
	// +kubebuilder:validation:Optional
	ClassPattern *string `json:"classPattern,omitempty" tf:"class_pattern,omitempty"`

	// (String) Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	// Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	// +kubebuilder:validation:Optional
	MessagePattern *string `json:"messagePattern,omitempty" tf:"message_pattern,omitempty"`
}

type ExceptionRulesCustomErrorRulesInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	CustomErrorRule []CustomErrorRulesCustomErrorRuleInitParameters `json:"customErrorRule,omitempty" tf:"custom_error_rule,omitempty"`
}

type ExceptionRulesCustomErrorRulesObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	CustomErrorRule []CustomErrorRulesCustomErrorRuleObservation `json:"customErrorRule,omitempty" tf:"custom_error_rule,omitempty"`
}

type ExceptionRulesCustomErrorRulesParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	CustomErrorRule []CustomErrorRulesCustomErrorRuleParameters `json:"customErrorRule" tf:"custom_error_rule,omitempty"`
}

type ExceptionRulesCustomHandledExceptionsInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	CustomHandledException []CustomHandledExceptionsCustomHandledExceptionInitParameters `json:"customHandledException,omitempty" tf:"custom_handled_exception,omitempty"`
}

type ExceptionRulesCustomHandledExceptionsObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	CustomHandledException []CustomHandledExceptionsCustomHandledExceptionObservation `json:"customHandledException,omitempty" tf:"custom_handled_exception,omitempty"`
}

type ExceptionRulesCustomHandledExceptionsParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	CustomHandledException []CustomHandledExceptionsCustomHandledExceptionParameters `json:"customHandledException" tf:"custom_handled_exception,omitempty"`
}

type ExceptionRulesIgnoredExceptionsCustomHandledExceptionInitParameters struct {

	// (String) The pattern will match if it is contained within the actual class name.
	// The pattern will match if it is contained within the actual class name.
	ClassPattern *string `json:"classPattern,omitempty" tf:"class_pattern,omitempty"`

	// (String) Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	// Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	MessagePattern *string `json:"messagePattern,omitempty" tf:"message_pattern,omitempty"`
}

type ExceptionRulesIgnoredExceptionsCustomHandledExceptionObservation struct {

	// (String) The pattern will match if it is contained within the actual class name.
	// The pattern will match if it is contained within the actual class name.
	ClassPattern *string `json:"classPattern,omitempty" tf:"class_pattern,omitempty"`

	// (String) Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	// Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	MessagePattern *string `json:"messagePattern,omitempty" tf:"message_pattern,omitempty"`
}

type ExceptionRulesIgnoredExceptionsCustomHandledExceptionParameters struct {

	// (String) The pattern will match if it is contained within the actual class name.
	// The pattern will match if it is contained within the actual class name.
	// +kubebuilder:validation:Optional
	ClassPattern *string `json:"classPattern,omitempty" tf:"class_pattern,omitempty"`

	// (String) Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	// Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	// +kubebuilder:validation:Optional
	MessagePattern *string `json:"messagePattern,omitempty" tf:"message_pattern,omitempty"`
}

type ExceptionRulesIgnoredExceptionsInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	CustomHandledException []ExceptionRulesIgnoredExceptionsCustomHandledExceptionInitParameters `json:"customHandledException,omitempty" tf:"custom_handled_exception,omitempty"`
}

type ExceptionRulesIgnoredExceptionsObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	CustomHandledException []ExceptionRulesIgnoredExceptionsCustomHandledExceptionObservation `json:"customHandledException,omitempty" tf:"custom_handled_exception,omitempty"`
}

type ExceptionRulesIgnoredExceptionsParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	CustomHandledException []ExceptionRulesIgnoredExceptionsCustomHandledExceptionParameters `json:"customHandledException" tf:"custom_handled_exception,omitempty"`
}

type ExceptionRulesSuccessForcingExceptionsCustomHandledExceptionInitParameters struct {

	// (String) The pattern will match if it is contained within the actual class name.
	// The pattern will match if it is contained within the actual class name.
	ClassPattern *string `json:"classPattern,omitempty" tf:"class_pattern,omitempty"`

	// (String) Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	// Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	MessagePattern *string `json:"messagePattern,omitempty" tf:"message_pattern,omitempty"`
}

type ExceptionRulesSuccessForcingExceptionsCustomHandledExceptionObservation struct {

	// (String) The pattern will match if it is contained within the actual class name.
	// The pattern will match if it is contained within the actual class name.
	ClassPattern *string `json:"classPattern,omitempty" tf:"class_pattern,omitempty"`

	// (String) Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	// Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	MessagePattern *string `json:"messagePattern,omitempty" tf:"message_pattern,omitempty"`
}

type ExceptionRulesSuccessForcingExceptionsCustomHandledExceptionParameters struct {

	// (String) The pattern will match if it is contained within the actual class name.
	// The pattern will match if it is contained within the actual class name.
	// +kubebuilder:validation:Optional
	ClassPattern *string `json:"classPattern,omitempty" tf:"class_pattern,omitempty"`

	// (String) Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	// Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
	// +kubebuilder:validation:Optional
	MessagePattern *string `json:"messagePattern,omitempty" tf:"message_pattern,omitempty"`
}

type ExceptionRulesSuccessForcingExceptionsInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	CustomHandledException []ExceptionRulesSuccessForcingExceptionsCustomHandledExceptionInitParameters `json:"customHandledException,omitempty" tf:"custom_handled_exception,omitempty"`
}

type ExceptionRulesSuccessForcingExceptionsObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	CustomHandledException []ExceptionRulesSuccessForcingExceptionsCustomHandledExceptionObservation `json:"customHandledException,omitempty" tf:"custom_handled_exception,omitempty"`
}

type ExceptionRulesSuccessForcingExceptionsParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	CustomHandledException []ExceptionRulesSuccessForcingExceptionsCustomHandledExceptionParameters `json:"customHandledException" tf:"custom_handled_exception,omitempty"`
}

type ServiceFailureExceptionRulesInitParameters struct {

	// (Block List, Max: 1) Some custom error situations are only detectable via a return value or other means. To support such cases, define a request attribute that captures the required data. Then define a custom error rule that determines if the request has failed based on the value of the request attribute. (see below for nested schema)
	// Some custom error situations are only detectable via a return value or other means. To support such cases, [define a request attribute](https://dt-url.net/ys5k0p4y) that captures the required data. Then define a custom error rule that determines if the request has failed based on the value of the request attribute.
	CustomErrorRules []ExceptionRulesCustomErrorRulesInitParameters `json:"customErrorRules,omitempty" tf:"custom_error_rules,omitempty"`

	// handled exceptions that should be treated as service failures. (see below for nested schema)
	// There may be situations where your application code handles exceptions gracefully in a manner that these failures aren't detected by Dynatrace. Use this setting to define specific gracefully-handled exceptions that should be treated as service failures.
	CustomHandledExceptions []ExceptionRulesCustomHandledExceptionsInitParameters `json:"customHandledExceptions,omitempty" tf:"custom_handled_exceptions,omitempty"`

	// (Boolean) Ignore all exceptions
	// Ignore all exceptions
	IgnoreAllExceptions *bool `json:"ignoreAllExceptions,omitempty" tf:"ignore_all_exceptions,omitempty"`

	// (Boolean) Ignore span failure detection
	// Ignore span failure detection
	IgnoreSpanFailureDetection *bool `json:"ignoreSpanFailureDetection,omitempty" tf:"ignore_span_failure_detection,omitempty"`

	// party code indicate a specific response, not an error. Use this setting to instruct Dynatrace to treat such exceptions as non-failed requests.. If an exception matching any of the defined patterns occurs in a request, it will not be considered as a failure. Other exceptions occurring at the same request might still mark the request as failed. (see below for nested schema)
	// Some exceptions that are thrown by legacy or 3rd-party code indicate a specific response, not an error. Use this setting to instruct Dynatrace to treat such exceptions as non-failed requests.. If an exception matching any of the defined patterns occurs in a request, it will not be considered as a failure. Other exceptions occurring at the same request might still mark the request as failed.
	IgnoredExceptions []ExceptionRulesIgnoredExceptionsInitParameters `json:"ignoredExceptions,omitempty" tf:"ignored_exceptions,omitempty"`

	// (Block List, Max: 1) Define exceptions which indicate that a service call should not be considered as failed. E.g. an exception indicating that the client aborted the operation.. If an exception matching any of the defined patterns occurs on the entry node of the service, it will be considered successful. Compared to ignored exceptions, the request will be considered successful even if other exceptions occur in the same request. (see below for nested schema)
	// Define exceptions which indicate that a service call should not be considered as failed. E.g. an exception indicating that the client aborted the operation.. If an exception matching any of the defined patterns occurs on the entry node of the service, it will be considered successful. Compared to ignored exceptions, the request will be considered successful even if other exceptions occur in the same request.
	SuccessForcingExceptions []ExceptionRulesSuccessForcingExceptionsInitParameters `json:"successForcingExceptions,omitempty" tf:"success_forcing_exceptions,omitempty"`
}

type ServiceFailureExceptionRulesObservation struct {

	// (Block List, Max: 1) Some custom error situations are only detectable via a return value or other means. To support such cases, define a request attribute that captures the required data. Then define a custom error rule that determines if the request has failed based on the value of the request attribute. (see below for nested schema)
	// Some custom error situations are only detectable via a return value or other means. To support such cases, [define a request attribute](https://dt-url.net/ys5k0p4y) that captures the required data. Then define a custom error rule that determines if the request has failed based on the value of the request attribute.
	CustomErrorRules []ExceptionRulesCustomErrorRulesObservation `json:"customErrorRules,omitempty" tf:"custom_error_rules,omitempty"`

	// handled exceptions that should be treated as service failures. (see below for nested schema)
	// There may be situations where your application code handles exceptions gracefully in a manner that these failures aren't detected by Dynatrace. Use this setting to define specific gracefully-handled exceptions that should be treated as service failures.
	CustomHandledExceptions []ExceptionRulesCustomHandledExceptionsObservation `json:"customHandledExceptions,omitempty" tf:"custom_handled_exceptions,omitempty"`

	// (Boolean) Ignore all exceptions
	// Ignore all exceptions
	IgnoreAllExceptions *bool `json:"ignoreAllExceptions,omitempty" tf:"ignore_all_exceptions,omitempty"`

	// (Boolean) Ignore span failure detection
	// Ignore span failure detection
	IgnoreSpanFailureDetection *bool `json:"ignoreSpanFailureDetection,omitempty" tf:"ignore_span_failure_detection,omitempty"`

	// party code indicate a specific response, not an error. Use this setting to instruct Dynatrace to treat such exceptions as non-failed requests.. If an exception matching any of the defined patterns occurs in a request, it will not be considered as a failure. Other exceptions occurring at the same request might still mark the request as failed. (see below for nested schema)
	// Some exceptions that are thrown by legacy or 3rd-party code indicate a specific response, not an error. Use this setting to instruct Dynatrace to treat such exceptions as non-failed requests.. If an exception matching any of the defined patterns occurs in a request, it will not be considered as a failure. Other exceptions occurring at the same request might still mark the request as failed.
	IgnoredExceptions []ExceptionRulesIgnoredExceptionsObservation `json:"ignoredExceptions,omitempty" tf:"ignored_exceptions,omitempty"`

	// (Block List, Max: 1) Define exceptions which indicate that a service call should not be considered as failed. E.g. an exception indicating that the client aborted the operation.. If an exception matching any of the defined patterns occurs on the entry node of the service, it will be considered successful. Compared to ignored exceptions, the request will be considered successful even if other exceptions occur in the same request. (see below for nested schema)
	// Define exceptions which indicate that a service call should not be considered as failed. E.g. an exception indicating that the client aborted the operation.. If an exception matching any of the defined patterns occurs on the entry node of the service, it will be considered successful. Compared to ignored exceptions, the request will be considered successful even if other exceptions occur in the same request.
	SuccessForcingExceptions []ExceptionRulesSuccessForcingExceptionsObservation `json:"successForcingExceptions,omitempty" tf:"success_forcing_exceptions,omitempty"`
}

type ServiceFailureExceptionRulesParameters struct {

	// (Block List, Max: 1) Some custom error situations are only detectable via a return value or other means. To support such cases, define a request attribute that captures the required data. Then define a custom error rule that determines if the request has failed based on the value of the request attribute. (see below for nested schema)
	// Some custom error situations are only detectable via a return value or other means. To support such cases, [define a request attribute](https://dt-url.net/ys5k0p4y) that captures the required data. Then define a custom error rule that determines if the request has failed based on the value of the request attribute.
	// +kubebuilder:validation:Optional
	CustomErrorRules []ExceptionRulesCustomErrorRulesParameters `json:"customErrorRules,omitempty" tf:"custom_error_rules,omitempty"`

	// handled exceptions that should be treated as service failures. (see below for nested schema)
	// There may be situations where your application code handles exceptions gracefully in a manner that these failures aren't detected by Dynatrace. Use this setting to define specific gracefully-handled exceptions that should be treated as service failures.
	// +kubebuilder:validation:Optional
	CustomHandledExceptions []ExceptionRulesCustomHandledExceptionsParameters `json:"customHandledExceptions,omitempty" tf:"custom_handled_exceptions,omitempty"`

	// (Boolean) Ignore all exceptions
	// Ignore all exceptions
	// +kubebuilder:validation:Optional
	IgnoreAllExceptions *bool `json:"ignoreAllExceptions" tf:"ignore_all_exceptions,omitempty"`

	// (Boolean) Ignore span failure detection
	// Ignore span failure detection
	// +kubebuilder:validation:Optional
	IgnoreSpanFailureDetection *bool `json:"ignoreSpanFailureDetection" tf:"ignore_span_failure_detection,omitempty"`

	// party code indicate a specific response, not an error. Use this setting to instruct Dynatrace to treat such exceptions as non-failed requests.. If an exception matching any of the defined patterns occurs in a request, it will not be considered as a failure. Other exceptions occurring at the same request might still mark the request as failed. (see below for nested schema)
	// Some exceptions that are thrown by legacy or 3rd-party code indicate a specific response, not an error. Use this setting to instruct Dynatrace to treat such exceptions as non-failed requests.. If an exception matching any of the defined patterns occurs in a request, it will not be considered as a failure. Other exceptions occurring at the same request might still mark the request as failed.
	// +kubebuilder:validation:Optional
	IgnoredExceptions []ExceptionRulesIgnoredExceptionsParameters `json:"ignoredExceptions,omitempty" tf:"ignored_exceptions,omitempty"`

	// (Block List, Max: 1) Define exceptions which indicate that a service call should not be considered as failed. E.g. an exception indicating that the client aborted the operation.. If an exception matching any of the defined patterns occurs on the entry node of the service, it will be considered successful. Compared to ignored exceptions, the request will be considered successful even if other exceptions occur in the same request. (see below for nested schema)
	// Define exceptions which indicate that a service call should not be considered as failed. E.g. an exception indicating that the client aborted the operation.. If an exception matching any of the defined patterns occurs on the entry node of the service, it will be considered successful. Compared to ignored exceptions, the request will be considered successful even if other exceptions occur in the same request.
	// +kubebuilder:validation:Optional
	SuccessForcingExceptions []ExceptionRulesSuccessForcingExceptionsParameters `json:"successForcingExceptions,omitempty" tf:"success_forcing_exceptions,omitempty"`
}

type ServiceFailureInitParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Block List, Max: 1) Customize failure detection for specific exceptions and errors (see below for nested schema)
	// Customize failure detection for specific exceptions and errors
	ExceptionRules []ServiceFailureExceptionRulesInitParameters `json:"exceptionRules,omitempty" tf:"exception_rules,omitempty"`

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`
}

type ServiceFailureObservation struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Block List, Max: 1) Customize failure detection for specific exceptions and errors (see below for nested schema)
	// Customize failure detection for specific exceptions and errors
	ExceptionRules []ServiceFailureExceptionRulesObservation `json:"exceptionRules,omitempty" tf:"exception_rules,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`
}

type ServiceFailureParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Block List, Max: 1) Customize failure detection for specific exceptions and errors (see below for nested schema)
	// Customize failure detection for specific exceptions and errors
	// +kubebuilder:validation:Optional
	ExceptionRules []ServiceFailureExceptionRulesParameters `json:"exceptionRules,omitempty" tf:"exception_rules,omitempty"`

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`
}

// ServiceFailureSpec defines the desired state of ServiceFailure
type ServiceFailureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceFailureParameters `json:"forProvider"`
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
	InitProvider ServiceFailureInitParameters `json:"initProvider,omitempty"`
}

// ServiceFailureStatus defines the observed state of ServiceFailure.
type ServiceFailureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceFailureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServiceFailure is the Schema for the ServiceFailures API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type ServiceFailure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceId) || (has(self.initProvider) && has(self.initProvider.serviceId))",message="spec.forProvider.serviceId is a required parameter"
	Spec   ServiceFailureSpec   `json:"spec"`
	Status ServiceFailureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceFailureList contains a list of ServiceFailures
type ServiceFailureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceFailure `json:"items"`
}

// Repository type metadata.
var (
	ServiceFailure_Kind             = "ServiceFailure"
	ServiceFailure_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceFailure_Kind}.String()
	ServiceFailure_KindAPIVersion   = ServiceFailure_Kind + "." + CRDGroupVersion.String()
	ServiceFailure_GroupVersionKind = CRDGroupVersion.WithKind(ServiceFailure_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceFailure{}, &ServiceFailureList{})
}
