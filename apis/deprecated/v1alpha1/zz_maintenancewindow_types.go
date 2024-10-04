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

type MaintenanceWindowInitParameters struct {

	// (String) A short description of the maintenance purpose
	// A short description of the maintenance purpose
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) The Maintenance Window is enabled or disabled
	// The Maintenance Window is enabled or disabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The name of the maintenance window, displayed in the UI
	// The name of the maintenance window, displayed in the UI
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) The schedule of the maintenance window (see below for nested schema)
	// The schedule of the maintenance window
	Schedule []ScheduleInitParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// (Block List, Max: 1) the tiles this Dashboard consist of (see below for nested schema)
	// the tiles this Dashboard consist of
	Scope []MaintenanceWindowScopeInitParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// (Boolean) Suppress execution of synthetic monitors during the maintenance
	// Suppress execution of synthetic monitors during the maintenance
	SuppressSynthMonExec *bool `json:"suppressSynthMonExec,omitempty" tf:"suppress_synth_mon_exec,omitempty"`

	// (String) The type of suppression of alerting and problem detection during the maintenance
	// The type of suppression of alerting and problem detection during the maintenance
	Suppression *string `json:"suppression,omitempty" tf:"suppression,omitempty"`

	// (String) The type of the maintenance: planned or unplanned
	// The type of the maintenance: planned or unplanned
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type MaintenanceWindowObservation struct {

	// (String) A short description of the maintenance purpose
	// A short description of the maintenance purpose
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) The Maintenance Window is enabled or disabled
	// The Maintenance Window is enabled or disabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the maintenance window, displayed in the UI
	// The name of the maintenance window, displayed in the UI
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) The schedule of the maintenance window (see below for nested schema)
	// The schedule of the maintenance window
	Schedule []ScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// (Block List, Max: 1) the tiles this Dashboard consist of (see below for nested schema)
	// the tiles this Dashboard consist of
	Scope []MaintenanceWindowScopeObservation `json:"scope,omitempty" tf:"scope,omitempty"`

	// (Boolean) Suppress execution of synthetic monitors during the maintenance
	// Suppress execution of synthetic monitors during the maintenance
	SuppressSynthMonExec *bool `json:"suppressSynthMonExec,omitempty" tf:"suppress_synth_mon_exec,omitempty"`

	// (String) The type of suppression of alerting and problem detection during the maintenance
	// The type of suppression of alerting and problem detection during the maintenance
	Suppression *string `json:"suppression,omitempty" tf:"suppression,omitempty"`

	// (String) The type of the maintenance: planned or unplanned
	// The type of the maintenance: planned or unplanned
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type MaintenanceWindowParameters struct {

	// (String) A short description of the maintenance purpose
	// A short description of the maintenance purpose
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) The Maintenance Window is enabled or disabled
	// The Maintenance Window is enabled or disabled
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The name of the maintenance window, displayed in the UI
	// The name of the maintenance window, displayed in the UI
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) The schedule of the maintenance window (see below for nested schema)
	// The schedule of the maintenance window
	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// (Block List, Max: 1) the tiles this Dashboard consist of (see below for nested schema)
	// the tiles this Dashboard consist of
	// +kubebuilder:validation:Optional
	Scope []MaintenanceWindowScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// (Boolean) Suppress execution of synthetic monitors during the maintenance
	// Suppress execution of synthetic monitors during the maintenance
	// +kubebuilder:validation:Optional
	SuppressSynthMonExec *bool `json:"suppressSynthMonExec,omitempty" tf:"suppress_synth_mon_exec,omitempty"`

	// (String) The type of suppression of alerting and problem detection during the maintenance
	// The type of suppression of alerting and problem detection during the maintenance
	// +kubebuilder:validation:Optional
	Suppression *string `json:"suppression,omitempty" tf:"suppression,omitempty"`

	// (String) The type of the maintenance: planned or unplanned
	// The type of the maintenance: planned or unplanned
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type MaintenanceWindowScopeInitParameters struct {

	// (Set of String) A list of Dynatrace entities (for example, hosts or services) to be included in the scope.  Allowed values are Dynatrace entity IDs
	// A list of Dynatrace entities (for example, hosts or services) to be included in the scope.  Allowed values are Dynatrace entity IDs
	// +listType=set
	Entities []*string `json:"entities,omitempty" tf:"entities,omitempty"`

	// (Block List) A list of matching rules for dynamic scope formation.  If several rules are set, the OR logic applies (see below for nested schema)
	// A list of matching rules for dynamic scope formation.  If several rules are set, the OR logic applies
	Matches []MatchesInitParameters `json:"matches,omitempty" tf:"matches,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type MaintenanceWindowScopeObservation struct {

	// (Set of String) A list of Dynatrace entities (for example, hosts or services) to be included in the scope.  Allowed values are Dynatrace entity IDs
	// A list of Dynatrace entities (for example, hosts or services) to be included in the scope.  Allowed values are Dynatrace entity IDs
	// +listType=set
	Entities []*string `json:"entities,omitempty" tf:"entities,omitempty"`

	// (Block List) A list of matching rules for dynamic scope formation.  If several rules are set, the OR logic applies (see below for nested schema)
	// A list of matching rules for dynamic scope formation.  If several rules are set, the OR logic applies
	Matches []MatchesObservation `json:"matches,omitempty" tf:"matches,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type MaintenanceWindowScopeParameters struct {

	// (Set of String) A list of Dynatrace entities (for example, hosts or services) to be included in the scope.  Allowed values are Dynatrace entity IDs
	// A list of Dynatrace entities (for example, hosts or services) to be included in the scope.  Allowed values are Dynatrace entity IDs
	// +kubebuilder:validation:Optional
	// +listType=set
	Entities []*string `json:"entities,omitempty" tf:"entities,omitempty"`

	// (Block List) A list of matching rules for dynamic scope formation.  If several rules are set, the OR logic applies (see below for nested schema)
	// A list of matching rules for dynamic scope formation.  If several rules are set, the OR logic applies
	// +kubebuilder:validation:Optional
	Matches []MatchesParameters `json:"matches,omitempty" tf:"matches,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type MatchesInitParameters struct {

	// (String) The ID of a management zone to which the matched entities must belong
	// The ID of a management zone to which the matched entities must belong
	MzID *string `json:"mzId,omitempty" tf:"mz_id,omitempty"`

	// (String) The logic that applies when several tags are specified: AND/OR.  If not set, the OR logic is used
	// The logic that applies when several tags are specified: AND/OR.  If not set, the OR logic is used
	TagCombination *string `json:"tagCombination,omitempty" tf:"tag_combination,omitempty"`

	// (Block Set) The tag you want to use for matching.  You can use custom tags from the UI, AWS tags, Cloud Foundry tags, OpenShift/Kubernetes, and tags based on environment variables (see below for nested schema)
	// The tag you want to use for matching.  You can use custom tags from the UI, AWS tags, Cloud Foundry tags, OpenShift/Kubernetes, and tags based on environment variables
	Tags []MatchesTagsInitParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// (String) The type of the maintenance: planned or unplanned
	// The type of the Dynatrace entities (for example, hosts or services) you want to pick up by matching
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type MatchesObservation struct {

	// (String) The ID of a management zone to which the matched entities must belong
	// The ID of a management zone to which the matched entities must belong
	MzID *string `json:"mzId,omitempty" tf:"mz_id,omitempty"`

	// (String) The logic that applies when several tags are specified: AND/OR.  If not set, the OR logic is used
	// The logic that applies when several tags are specified: AND/OR.  If not set, the OR logic is used
	TagCombination *string `json:"tagCombination,omitempty" tf:"tag_combination,omitempty"`

	// (Block Set) The tag you want to use for matching.  You can use custom tags from the UI, AWS tags, Cloud Foundry tags, OpenShift/Kubernetes, and tags based on environment variables (see below for nested schema)
	// The tag you want to use for matching.  You can use custom tags from the UI, AWS tags, Cloud Foundry tags, OpenShift/Kubernetes, and tags based on environment variables
	Tags []MatchesTagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	// (String) The type of the maintenance: planned or unplanned
	// The type of the Dynatrace entities (for example, hosts or services) you want to pick up by matching
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type MatchesParameters struct {

	// (String) The ID of a management zone to which the matched entities must belong
	// The ID of a management zone to which the matched entities must belong
	// +kubebuilder:validation:Optional
	MzID *string `json:"mzId,omitempty" tf:"mz_id,omitempty"`

	// (String) The logic that applies when several tags are specified: AND/OR.  If not set, the OR logic is used
	// The logic that applies when several tags are specified: AND/OR.  If not set, the OR logic is used
	// +kubebuilder:validation:Optional
	TagCombination *string `json:"tagCombination,omitempty" tf:"tag_combination,omitempty"`

	// (Block Set) The tag you want to use for matching.  You can use custom tags from the UI, AWS tags, Cloud Foundry tags, OpenShift/Kubernetes, and tags based on environment variables (see below for nested schema)
	// The tag you want to use for matching.  You can use custom tags from the UI, AWS tags, Cloud Foundry tags, OpenShift/Kubernetes, and tags based on environment variables
	// +kubebuilder:validation:Optional
	Tags []MatchesTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// (String) The type of the maintenance: planned or unplanned
	// The type of the Dynatrace entities (for example, hosts or services) you want to pick up by matching
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type MatchesTagsInitParameters struct {

	// (String) The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the CONTEXTLESS value
	// The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the `CONTEXTLESS` value
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String) The key of the tag. Custom tags have the tag value here
	// The key of the tag. Custom tags have the tag value here
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`

	// (String) The value of the tag. Not applicable to custom tags
	// The value of the tag. Not applicable to custom tags
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type MatchesTagsObservation struct {

	// (String) The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the CONTEXTLESS value
	// The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the `CONTEXTLESS` value
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String) The key of the tag. Custom tags have the tag value here
	// The key of the tag. Custom tags have the tag value here
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`

	// (String) The value of the tag. Not applicable to custom tags
	// The value of the tag. Not applicable to custom tags
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type MatchesTagsParameters struct {

	// (String) The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the CONTEXTLESS value
	// The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the `CONTEXTLESS` value
	// +kubebuilder:validation:Optional
	Context *string `json:"context" tf:"context,omitempty"`

	// (String) The key of the tag. Custom tags have the tag value here
	// The key of the tag. Custom tags have the tag value here
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`

	// (String) The value of the tag. Not applicable to custom tags
	// The value of the tag. Not applicable to custom tags
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RecurrenceInitParameters struct {

	// (Number) The day of the month for monthly maintenance.  The value of 31 is treated as the last day of the month for months that don't have a 31st day. The value of 30 is also treated as the last day of the month for February
	// The day of the month for monthly maintenance.  The value of `31` is treated as the last day of the month for months that don't have a 31st day. The value of `30` is also treated as the last day of the month for February
	DayOfMonth *float64 `json:"dayOfMonth,omitempty" tf:"day_of_month,omitempty"`

	// (String) The day of the week for weekly maintenance.  The format is the full name of the day in upper case, for example THURSDAY
	// The day of the week for weekly maintenance.  The format is the full name of the day in upper case, for example `THURSDAY`
	DayOfWeek *string `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// (Number) The duration of the maintenance window in minutes
	// The duration of the maintenance window in minutes
	DurationMinutes *float64 `json:"durationMinutes,omitempty" tf:"duration_minutes,omitempty"`

	// (String) The start time of the maintenance window in HH:mm format
	// The start time of the maintenance window in HH:mm format
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type RecurrenceObservation struct {

	// (Number) The day of the month for monthly maintenance.  The value of 31 is treated as the last day of the month for months that don't have a 31st day. The value of 30 is also treated as the last day of the month for February
	// The day of the month for monthly maintenance.  The value of `31` is treated as the last day of the month for months that don't have a 31st day. The value of `30` is also treated as the last day of the month for February
	DayOfMonth *float64 `json:"dayOfMonth,omitempty" tf:"day_of_month,omitempty"`

	// (String) The day of the week for weekly maintenance.  The format is the full name of the day in upper case, for example THURSDAY
	// The day of the week for weekly maintenance.  The format is the full name of the day in upper case, for example `THURSDAY`
	DayOfWeek *string `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// (Number) The duration of the maintenance window in minutes
	// The duration of the maintenance window in minutes
	DurationMinutes *float64 `json:"durationMinutes,omitempty" tf:"duration_minutes,omitempty"`

	// (String) The start time of the maintenance window in HH:mm format
	// The start time of the maintenance window in HH:mm format
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type RecurrenceParameters struct {

	// (Number) The day of the month for monthly maintenance.  The value of 31 is treated as the last day of the month for months that don't have a 31st day. The value of 30 is also treated as the last day of the month for February
	// The day of the month for monthly maintenance.  The value of `31` is treated as the last day of the month for months that don't have a 31st day. The value of `30` is also treated as the last day of the month for February
	// +kubebuilder:validation:Optional
	DayOfMonth *float64 `json:"dayOfMonth,omitempty" tf:"day_of_month,omitempty"`

	// (String) The day of the week for weekly maintenance.  The format is the full name of the day in upper case, for example THURSDAY
	// The day of the week for weekly maintenance.  The format is the full name of the day in upper case, for example `THURSDAY`
	// +kubebuilder:validation:Optional
	DayOfWeek *string `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// (Number) The duration of the maintenance window in minutes
	// The duration of the maintenance window in minutes
	// +kubebuilder:validation:Optional
	DurationMinutes *float64 `json:"durationMinutes" tf:"duration_minutes,omitempty"`

	// (String) The start time of the maintenance window in HH:mm format
	// The start time of the maintenance window in HH:mm format
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type ScheduleInitParameters struct {

	// mm-dd HH:mm format
	// The end date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// (Block List, Max: 1) The recurrence of the maintenance window (see below for nested schema)
	// The recurrence of the maintenance window
	Recurrence []RecurrenceInitParameters `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// (String) The type of the schedule recurrence. Possible values are DAILY, MONTHLY, ONCE and WEEKLY
	// The type of the schedule recurrence. Possible values are `DAILY`, `MONTHLY`, `ONCE` and `WEEKLY`
	RecurrenceType *string `json:"recurrenceType,omitempty" tf:"recurrence_type,omitempty"`

	// mm-dd HH:mm format
	// The start date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format
	Start *string `json:"start,omitempty" tf:"start,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`

	// (String) The time zone of the start and end time. Default time zone is UTC. You can use either UTC offset UTC+01:00 format or the IANA Time Zone Database format (for example, Europe/Vienna)
	// The time zone of the start and end time. Default time zone is UTC. You can use either UTC offset `UTC+01:00` format or the IANA Time Zone Database format (for example, `Europe/Vienna`)
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ScheduleObservation struct {

	// mm-dd HH:mm format
	// The end date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// (Block List, Max: 1) The recurrence of the maintenance window (see below for nested schema)
	// The recurrence of the maintenance window
	Recurrence []RecurrenceObservation `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// (String) The type of the schedule recurrence. Possible values are DAILY, MONTHLY, ONCE and WEEKLY
	// The type of the schedule recurrence. Possible values are `DAILY`, `MONTHLY`, `ONCE` and `WEEKLY`
	RecurrenceType *string `json:"recurrenceType,omitempty" tf:"recurrence_type,omitempty"`

	// mm-dd HH:mm format
	// The start date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format
	Start *string `json:"start,omitempty" tf:"start,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`

	// (String) The time zone of the start and end time. Default time zone is UTC. You can use either UTC offset UTC+01:00 format or the IANA Time Zone Database format (for example, Europe/Vienna)
	// The time zone of the start and end time. Default time zone is UTC. You can use either UTC offset `UTC+01:00` format or the IANA Time Zone Database format (for example, `Europe/Vienna`)
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ScheduleParameters struct {

	// mm-dd HH:mm format
	// The end date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format
	// +kubebuilder:validation:Optional
	End *string `json:"end" tf:"end,omitempty"`

	// (Block List, Max: 1) The recurrence of the maintenance window (see below for nested schema)
	// The recurrence of the maintenance window
	// +kubebuilder:validation:Optional
	Recurrence []RecurrenceParameters `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// (String) The type of the schedule recurrence. Possible values are DAILY, MONTHLY, ONCE and WEEKLY
	// The type of the schedule recurrence. Possible values are `DAILY`, `MONTHLY`, `ONCE` and `WEEKLY`
	// +kubebuilder:validation:Optional
	RecurrenceType *string `json:"recurrenceType" tf:"recurrence_type,omitempty"`

	// mm-dd HH:mm format
	// The start date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format
	// +kubebuilder:validation:Optional
	Start *string `json:"start" tf:"start,omitempty"`

	// (String) allows for configuring properties that are not explicitly supported by the current version of this provider
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`

	// (String) The time zone of the start and end time. Default time zone is UTC. You can use either UTC offset UTC+01:00 format or the IANA Time Zone Database format (for example, Europe/Vienna)
	// The time zone of the start and end time. Default time zone is UTC. You can use either UTC offset `UTC+01:00` format or the IANA Time Zone Database format (for example, `Europe/Vienna`)
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId" tf:"zone_id,omitempty"`
}

// MaintenanceWindowSpec defines the desired state of MaintenanceWindow
type MaintenanceWindowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MaintenanceWindowParameters `json:"forProvider"`
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
	InitProvider MaintenanceWindowInitParameters `json:"initProvider,omitempty"`
}

// MaintenanceWindowStatus defines the observed state of MaintenanceWindow.
type MaintenanceWindowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MaintenanceWindowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MaintenanceWindow is the Schema for the MaintenanceWindows API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type MaintenanceWindow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.suppression) || (has(self.initProvider) && has(self.initProvider.suppression))",message="spec.forProvider.suppression is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   MaintenanceWindowSpec   `json:"spec"`
	Status MaintenanceWindowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceWindowList contains a list of MaintenanceWindows
type MaintenanceWindowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MaintenanceWindow `json:"items"`
}

// Repository type metadata.
var (
	MaintenanceWindow_Kind             = "MaintenanceWindow"
	MaintenanceWindow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MaintenanceWindow_Kind}.String()
	MaintenanceWindow_KindAPIVersion   = MaintenanceWindow_Kind + "." + CRDGroupVersion.String()
	MaintenanceWindow_GroupVersionKind = CRDGroupVersion.WithKind(MaintenanceWindow_Kind)
)

func init() {
	SchemeBuilder.Register(&MaintenanceWindow{}, &MaintenanceWindowList{})
}
