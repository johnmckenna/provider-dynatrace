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

type AdditionalInformationAdditionalInformationInitParameters struct {

	// (String) Name
	// Name
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) no documentation available
	// no documentation available
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// (String) no documentation available
	// no documentation available
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AdditionalInformationAdditionalInformationObservation struct {

	// (String) Name
	// Name
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) no documentation available
	// no documentation available
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// (String) no documentation available
	// no documentation available
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AdditionalInformationAdditionalInformationParameters struct {

	// (String) Name
	// Name
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// (String) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// (String) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type AdditionalInformationInitParameters struct {

	// (Block List, Max: 1) Define key/value pairs that further describe this team — for example, cost center, solution type, or business unit assignments. (see below for nested schema)
	AdditionalInformation []AdditionalInformationAdditionalInformationInitParameters `json:"additionalInformation,omitempty" tf:"additional_information,omitempty"`
}

type AdditionalInformationObservation struct {

	// (Block List, Max: 1) Define key/value pairs that further describe this team — for example, cost center, solution type, or business unit assignments. (see below for nested schema)
	AdditionalInformation []AdditionalInformationAdditionalInformationObservation `json:"additionalInformation,omitempty" tf:"additional_information,omitempty"`
}

type AdditionalInformationParameters struct {

	// (Block List, Max: 1) Define key/value pairs that further describe this team — for example, cost center, solution type, or business unit assignments. (see below for nested schema)
	// +kubebuilder:validation:Optional
	AdditionalInformation []AdditionalInformationAdditionalInformationParameters `json:"additionalInformation" tf:"additional_information,omitempty"`
}

type ContactDetailInitParameters struct {

	// (String) no documentation available
	// no documentation available
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) Possible Values: EMAIL, JIRA, MS_TEAMS, SLACK
	// Possible Values: `EMAIL`, `JIRA`, `MS_TEAMS`, `SLACK`
	IntegrationType *string `json:"integrationType,omitempty" tf:"integration_type,omitempty"`

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	Jira []JiraInitParameters `json:"jira,omitempty" tf:"jira,omitempty"`

	// (String) Team
	// Team
	MsTeams *string `json:"msTeams,omitempty" tf:"ms_teams,omitempty"`

	// (String) Channel
	// Channel
	SlackChannel *string `json:"slackChannel,omitempty" tf:"slack_channel,omitempty"`

	// (String) no documentation available
	// no documentation available
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ContactDetailObservation struct {

	// (String) no documentation available
	// no documentation available
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) Possible Values: EMAIL, JIRA, MS_TEAMS, SLACK
	// Possible Values: `EMAIL`, `JIRA`, `MS_TEAMS`, `SLACK`
	IntegrationType *string `json:"integrationType,omitempty" tf:"integration_type,omitempty"`

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	Jira []JiraObservation `json:"jira,omitempty" tf:"jira,omitempty"`

	// (String) Team
	// Team
	MsTeams *string `json:"msTeams,omitempty" tf:"ms_teams,omitempty"`

	// (String) Channel
	// Channel
	SlackChannel *string `json:"slackChannel,omitempty" tf:"slack_channel,omitempty"`

	// (String) no documentation available
	// no documentation available
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ContactDetailParameters struct {

	// (String) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) Possible Values: EMAIL, JIRA, MS_TEAMS, SLACK
	// Possible Values: `EMAIL`, `JIRA`, `MS_TEAMS`, `SLACK`
	// +kubebuilder:validation:Optional
	IntegrationType *string `json:"integrationType" tf:"integration_type,omitempty"`

	// (Block List, Max: 1) no documentation available (see below for nested schema)
	// no documentation available
	// +kubebuilder:validation:Optional
	Jira []JiraParameters `json:"jira,omitempty" tf:"jira,omitempty"`

	// (String) Team
	// Team
	// +kubebuilder:validation:Optional
	MsTeams *string `json:"msTeams,omitempty" tf:"ms_teams,omitempty"`

	// (String) Channel
	// Channel
	// +kubebuilder:validation:Optional
	SlackChannel *string `json:"slackChannel,omitempty" tf:"slack_channel,omitempty"`

	// (String) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ContactDetailsInitParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	ContactDetail []ContactDetailInitParameters `json:"contactDetail,omitempty" tf:"contact_detail,omitempty"`
}

type ContactDetailsObservation struct {

	// (Block List, Min: 1) (see below for nested schema)
	ContactDetail []ContactDetailObservation `json:"contactDetail,omitempty" tf:"contact_detail,omitempty"`
}

type ContactDetailsParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	ContactDetail []ContactDetailParameters `json:"contactDetail" tf:"contact_detail,omitempty"`
}

type JiraInitParameters struct {

	// (String) Default Assignee
	// Default Assignee
	DefaultAssignee *string `json:"defaultAssignee,omitempty" tf:"default_assignee,omitempty"`

	// (String) no documentation available
	// no documentation available
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type JiraObservation struct {

	// (String) Default Assignee
	// Default Assignee
	DefaultAssignee *string `json:"defaultAssignee,omitempty" tf:"default_assignee,omitempty"`

	// (String) no documentation available
	// no documentation available
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type JiraParameters struct {

	// (String) Default Assignee
	// Default Assignee
	// +kubebuilder:validation:Optional
	DefaultAssignee *string `json:"defaultAssignee" tf:"default_assignee,omitempty"`

	// (String) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	Project *string `json:"project" tf:"project,omitempty"`
}

type LinkInitParameters struct {

	// (String) Possible Values: DASHBOARD, DOCUMENTATION, HEALTH_APP, REPOSITORY, RUNBOOK, URL, WIKI
	// Possible Values: `DASHBOARD`, `DOCUMENTATION`, `HEALTH_APP`, `REPOSITORY`, `RUNBOOK`, `URL`, `WIKI`
	LinkType *string `json:"linkType,omitempty" tf:"link_type,omitempty"`

	// (String) no documentation available
	// no documentation available
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type LinkObservation struct {

	// (String) Possible Values: DASHBOARD, DOCUMENTATION, HEALTH_APP, REPOSITORY, RUNBOOK, URL, WIKI
	// Possible Values: `DASHBOARD`, `DOCUMENTATION`, `HEALTH_APP`, `REPOSITORY`, `RUNBOOK`, `URL`, `WIKI`
	LinkType *string `json:"linkType,omitempty" tf:"link_type,omitempty"`

	// (String) no documentation available
	// no documentation available
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type LinkParameters struct {

	// (String) Possible Values: DASHBOARD, DOCUMENTATION, HEALTH_APP, REPOSITORY, RUNBOOK, URL, WIKI
	// Possible Values: `DASHBOARD`, `DOCUMENTATION`, `HEALTH_APP`, `REPOSITORY`, `RUNBOOK`, `URL`, `WIKI`
	// +kubebuilder:validation:Optional
	LinkType *string `json:"linkType" tf:"link_type,omitempty"`

	// (String) no documentation available
	// no documentation available
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`
}

type LinksInitParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	Link []LinkInitParameters `json:"link,omitempty" tf:"link,omitempty"`
}

type LinksObservation struct {

	// (Block List, Min: 1) (see below for nested schema)
	Link []LinkObservation `json:"link,omitempty" tf:"link,omitempty"`
}

type LinksParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Link []LinkParameters `json:"link" tf:"link,omitempty"`
}

type OwnershipTeamsInitParameters struct {

	// (Block List, Max: 1) Define key/value pairs that further describe this team — for example, cost center, solution type, or business unit assignments. (see below for nested schema)
	// Define key/value pairs that further describe this team — for example, cost center, solution type, or business unit assignments.
	AdditionalInformation []AdditionalInformationInitParameters `json:"additionalInformation,omitempty" tf:"additional_information,omitempty"`

	// (Block List, Max: 1) Define options for messaging integration or other means of contacting this team. (see below for nested schema)
	// Define options for messaging integration or other means of contacting this team.
	ContactDetails []ContactDetailsInitParameters `json:"contactDetails,omitempty" tf:"contact_details,omitempty"`

	// (String) Description
	// Description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) This field should only be used for the automation purpose when importing team information.
	// This field should only be used for the automation purpose when importing team information.
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// (String) The team identifier is used to reference the team from any entity in Dynatrace. If you are using Kubernetes labels, keep in mind the 63 character limit that they enforce.
	// The team identifier is used to reference the team from any entity in Dynatrace. If you are using Kubernetes labels, keep in mind the 63 character limit that they enforce.
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (Block List, Max: 1) Include links to online resources where information relevant to this team’s responsibilities can be found. (see below for nested schema)
	// Include links to online resources where information relevant to this team’s responsibilities can be found.
	Links []LinksInitParameters `json:"links,omitempty" tf:"links,omitempty"`

	// (String) Team name
	// Team name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) Turn on all responsibility assignments that apply to this team. (see below for nested schema)
	// Turn on all responsibility assignments that apply to this team.
	Responsibilities []ResponsibilitiesInitParameters `json:"responsibilities,omitempty" tf:"responsibilities,omitempty"`

	// (Block List, Max: 1) The supplementary team identifiers can be optionally used in addition to the main team identifier to reference this team from any entity in Dynatrace. Up to 3 supplementary identifiers are supported. (see below for nested schema)
	// The supplementary team identifiers can be optionally used in addition to the main team identifier to reference this team from any entity in Dynatrace. Up to 3 supplementary identifiers are supported.
	SupplementaryIdentifiers []SupplementaryIdentifiersInitParameters `json:"supplementaryIdentifiers,omitempty" tf:"supplementary_identifiers,omitempty"`
}

type OwnershipTeamsObservation struct {

	// (Block List, Max: 1) Define key/value pairs that further describe this team — for example, cost center, solution type, or business unit assignments. (see below for nested schema)
	// Define key/value pairs that further describe this team — for example, cost center, solution type, or business unit assignments.
	AdditionalInformation []AdditionalInformationObservation `json:"additionalInformation,omitempty" tf:"additional_information,omitempty"`

	// (Block List, Max: 1) Define options for messaging integration or other means of contacting this team. (see below for nested schema)
	// Define options for messaging integration or other means of contacting this team.
	ContactDetails []ContactDetailsObservation `json:"contactDetails,omitempty" tf:"contact_details,omitempty"`

	// (String) Description
	// Description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) This field should only be used for the automation purpose when importing team information.
	// This field should only be used for the automation purpose when importing team information.
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The team identifier is used to reference the team from any entity in Dynatrace. If you are using Kubernetes labels, keep in mind the 63 character limit that they enforce.
	// The team identifier is used to reference the team from any entity in Dynatrace. If you are using Kubernetes labels, keep in mind the 63 character limit that they enforce.
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (Block List, Max: 1) Include links to online resources where information relevant to this team’s responsibilities can be found. (see below for nested schema)
	// Include links to online resources where information relevant to this team’s responsibilities can be found.
	Links []LinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	// (String) Team name
	// Team name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) Turn on all responsibility assignments that apply to this team. (see below for nested schema)
	// Turn on all responsibility assignments that apply to this team.
	Responsibilities []ResponsibilitiesObservation `json:"responsibilities,omitempty" tf:"responsibilities,omitempty"`

	// (Block List, Max: 1) The supplementary team identifiers can be optionally used in addition to the main team identifier to reference this team from any entity in Dynatrace. Up to 3 supplementary identifiers are supported. (see below for nested schema)
	// The supplementary team identifiers can be optionally used in addition to the main team identifier to reference this team from any entity in Dynatrace. Up to 3 supplementary identifiers are supported.
	SupplementaryIdentifiers []SupplementaryIdentifiersObservation `json:"supplementaryIdentifiers,omitempty" tf:"supplementary_identifiers,omitempty"`
}

type OwnershipTeamsParameters struct {

	// (Block List, Max: 1) Define key/value pairs that further describe this team — for example, cost center, solution type, or business unit assignments. (see below for nested schema)
	// Define key/value pairs that further describe this team — for example, cost center, solution type, or business unit assignments.
	// +kubebuilder:validation:Optional
	AdditionalInformation []AdditionalInformationParameters `json:"additionalInformation,omitempty" tf:"additional_information,omitempty"`

	// (Block List, Max: 1) Define options for messaging integration or other means of contacting this team. (see below for nested schema)
	// Define options for messaging integration or other means of contacting this team.
	// +kubebuilder:validation:Optional
	ContactDetails []ContactDetailsParameters `json:"contactDetails,omitempty" tf:"contact_details,omitempty"`

	// (String) Description
	// Description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) This field should only be used for the automation purpose when importing team information.
	// This field should only be used for the automation purpose when importing team information.
	// +kubebuilder:validation:Optional
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// (String) The team identifier is used to reference the team from any entity in Dynatrace. If you are using Kubernetes labels, keep in mind the 63 character limit that they enforce.
	// The team identifier is used to reference the team from any entity in Dynatrace. If you are using Kubernetes labels, keep in mind the 63 character limit that they enforce.
	// +kubebuilder:validation:Optional
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (Block List, Max: 1) Include links to online resources where information relevant to this team’s responsibilities can be found. (see below for nested schema)
	// Include links to online resources where information relevant to this team’s responsibilities can be found.
	// +kubebuilder:validation:Optional
	Links []LinksParameters `json:"links,omitempty" tf:"links,omitempty"`

	// (String) Team name
	// Team name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) Turn on all responsibility assignments that apply to this team. (see below for nested schema)
	// Turn on all responsibility assignments that apply to this team.
	// +kubebuilder:validation:Optional
	Responsibilities []ResponsibilitiesParameters `json:"responsibilities,omitempty" tf:"responsibilities,omitempty"`

	// (Block List, Max: 1) The supplementary team identifiers can be optionally used in addition to the main team identifier to reference this team from any entity in Dynatrace. Up to 3 supplementary identifiers are supported. (see below for nested schema)
	// The supplementary team identifiers can be optionally used in addition to the main team identifier to reference this team from any entity in Dynatrace. Up to 3 supplementary identifiers are supported.
	// +kubebuilder:validation:Optional
	SupplementaryIdentifiers []SupplementaryIdentifiersParameters `json:"supplementaryIdentifiers,omitempty" tf:"supplementary_identifiers,omitempty"`
}

type ResponsibilitiesInitParameters struct {

	// (Boolean) Responsible for developing and maintaining high quality software. Development teams are responsible for making code changes to address performance regressions, errors, or security vulnerabilities.
	// Responsible for developing and maintaining high quality software. Development teams are responsible for making code changes to address performance regressions, errors, or security vulnerabilities.
	Development *bool `json:"development,omitempty" tf:"development,omitempty"`

	// (Boolean) Responsible for the administration, management, and support of the IT infrastructure including physical servers, virtualization, and cloud. Teams with infrastructure responsibility are responsible for addressing hardware issues, resource limits, and operating system vulnerabilities.
	// Responsible for the administration, management, and support of the IT infrastructure including physical servers, virtualization, and cloud. Teams with infrastructure responsibility are responsible for addressing hardware issues, resource limits, and operating system vulnerabilities.
	Infrastructure *bool `json:"infrastructure,omitempty" tf:"infrastructure,omitempty"`

	// (Boolean) Responsible for ensuring that applications in development align with business needs and meet the usability requirements of users, stakeholders, customers, and external partners. Teams with line of business responsibility are responsible for understanding the customer experience and how it affects business goals.
	// Responsible for ensuring that applications in development align with business needs and meet the usability requirements of users, stakeholders, customers, and external partners. Teams with line of business responsibility are responsible for understanding the customer experience and how it affects business goals.
	LineOfBusiness *bool `json:"lineOfBusiness,omitempty" tf:"line_of_business,omitempty"`

	// (Boolean) Responsible for deploying and managing software, with a focus on high availability and performance. Teams with operations responsibilities needs to understand the impact, priority, and team responsible for addressing problems detected by Dynatrace.
	// Responsible for deploying and managing software, with a focus on high availability and performance. Teams with operations responsibilities needs to understand the impact, priority, and team responsible for addressing problems detected by Dynatrace.
	Operations *bool `json:"operations,omitempty" tf:"operations,omitempty"`

	// (Boolean) Responsible for the security posture of the organization. Teams with security responsibility must understand the impact, priority, and team responsible for addressing security vulnerabilities.
	// Responsible for the security posture of the organization. Teams with security responsibility must understand the impact, priority, and team responsible for addressing security vulnerabilities.
	Security *bool `json:"security,omitempty" tf:"security,omitempty"`
}

type ResponsibilitiesObservation struct {

	// (Boolean) Responsible for developing and maintaining high quality software. Development teams are responsible for making code changes to address performance regressions, errors, or security vulnerabilities.
	// Responsible for developing and maintaining high quality software. Development teams are responsible for making code changes to address performance regressions, errors, or security vulnerabilities.
	Development *bool `json:"development,omitempty" tf:"development,omitempty"`

	// (Boolean) Responsible for the administration, management, and support of the IT infrastructure including physical servers, virtualization, and cloud. Teams with infrastructure responsibility are responsible for addressing hardware issues, resource limits, and operating system vulnerabilities.
	// Responsible for the administration, management, and support of the IT infrastructure including physical servers, virtualization, and cloud. Teams with infrastructure responsibility are responsible for addressing hardware issues, resource limits, and operating system vulnerabilities.
	Infrastructure *bool `json:"infrastructure,omitempty" tf:"infrastructure,omitempty"`

	// (Boolean) Responsible for ensuring that applications in development align with business needs and meet the usability requirements of users, stakeholders, customers, and external partners. Teams with line of business responsibility are responsible for understanding the customer experience and how it affects business goals.
	// Responsible for ensuring that applications in development align with business needs and meet the usability requirements of users, stakeholders, customers, and external partners. Teams with line of business responsibility are responsible for understanding the customer experience and how it affects business goals.
	LineOfBusiness *bool `json:"lineOfBusiness,omitempty" tf:"line_of_business,omitempty"`

	// (Boolean) Responsible for deploying and managing software, with a focus on high availability and performance. Teams with operations responsibilities needs to understand the impact, priority, and team responsible for addressing problems detected by Dynatrace.
	// Responsible for deploying and managing software, with a focus on high availability and performance. Teams with operations responsibilities needs to understand the impact, priority, and team responsible for addressing problems detected by Dynatrace.
	Operations *bool `json:"operations,omitempty" tf:"operations,omitempty"`

	// (Boolean) Responsible for the security posture of the organization. Teams with security responsibility must understand the impact, priority, and team responsible for addressing security vulnerabilities.
	// Responsible for the security posture of the organization. Teams with security responsibility must understand the impact, priority, and team responsible for addressing security vulnerabilities.
	Security *bool `json:"security,omitempty" tf:"security,omitempty"`
}

type ResponsibilitiesParameters struct {

	// (Boolean) Responsible for developing and maintaining high quality software. Development teams are responsible for making code changes to address performance regressions, errors, or security vulnerabilities.
	// Responsible for developing and maintaining high quality software. Development teams are responsible for making code changes to address performance regressions, errors, or security vulnerabilities.
	// +kubebuilder:validation:Optional
	Development *bool `json:"development" tf:"development,omitempty"`

	// (Boolean) Responsible for the administration, management, and support of the IT infrastructure including physical servers, virtualization, and cloud. Teams with infrastructure responsibility are responsible for addressing hardware issues, resource limits, and operating system vulnerabilities.
	// Responsible for the administration, management, and support of the IT infrastructure including physical servers, virtualization, and cloud. Teams with infrastructure responsibility are responsible for addressing hardware issues, resource limits, and operating system vulnerabilities.
	// +kubebuilder:validation:Optional
	Infrastructure *bool `json:"infrastructure" tf:"infrastructure,omitempty"`

	// (Boolean) Responsible for ensuring that applications in development align with business needs and meet the usability requirements of users, stakeholders, customers, and external partners. Teams with line of business responsibility are responsible for understanding the customer experience and how it affects business goals.
	// Responsible for ensuring that applications in development align with business needs and meet the usability requirements of users, stakeholders, customers, and external partners. Teams with line of business responsibility are responsible for understanding the customer experience and how it affects business goals.
	// +kubebuilder:validation:Optional
	LineOfBusiness *bool `json:"lineOfBusiness" tf:"line_of_business,omitempty"`

	// (Boolean) Responsible for deploying and managing software, with a focus on high availability and performance. Teams with operations responsibilities needs to understand the impact, priority, and team responsible for addressing problems detected by Dynatrace.
	// Responsible for deploying and managing software, with a focus on high availability and performance. Teams with operations responsibilities needs to understand the impact, priority, and team responsible for addressing problems detected by Dynatrace.
	// +kubebuilder:validation:Optional
	Operations *bool `json:"operations" tf:"operations,omitempty"`

	// (Boolean) Responsible for the security posture of the organization. Teams with security responsibility must understand the impact, priority, and team responsible for addressing security vulnerabilities.
	// Responsible for the security posture of the organization. Teams with security responsibility must understand the impact, priority, and team responsible for addressing security vulnerabilities.
	// +kubebuilder:validation:Optional
	Security *bool `json:"security" tf:"security,omitempty"`
}

type SupplementaryIdentifierInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// Supplementary Identifier
	SupplementaryIdentifier *string `json:"supplementaryIdentifier,omitempty" tf:"supplementary_identifier,omitempty"`
}

type SupplementaryIdentifierObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// Supplementary Identifier
	SupplementaryIdentifier *string `json:"supplementaryIdentifier,omitempty" tf:"supplementary_identifier,omitempty"`
}

type SupplementaryIdentifierParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// Supplementary Identifier
	// +kubebuilder:validation:Optional
	SupplementaryIdentifier *string `json:"supplementaryIdentifier" tf:"supplementary_identifier,omitempty"`
}

type SupplementaryIdentifiersInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	SupplementaryIdentifier []SupplementaryIdentifierInitParameters `json:"supplementaryIdentifier,omitempty" tf:"supplementary_identifier,omitempty"`
}

type SupplementaryIdentifiersObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	SupplementaryIdentifier []SupplementaryIdentifierObservation `json:"supplementaryIdentifier,omitempty" tf:"supplementary_identifier,omitempty"`
}

type SupplementaryIdentifiersParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	SupplementaryIdentifier []SupplementaryIdentifierParameters `json:"supplementaryIdentifier" tf:"supplementary_identifier,omitempty"`
}

// OwnershipTeamsSpec defines the desired state of OwnershipTeams
type OwnershipTeamsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OwnershipTeamsParameters `json:"forProvider"`
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
	InitProvider OwnershipTeamsInitParameters `json:"initProvider,omitempty"`
}

// OwnershipTeamsStatus defines the observed state of OwnershipTeams.
type OwnershipTeamsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OwnershipTeamsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OwnershipTeams is the Schema for the OwnershipTeamss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type OwnershipTeams struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identifier) || (has(self.initProvider) && has(self.initProvider.identifier))",message="spec.forProvider.identifier is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.responsibilities) || (has(self.initProvider) && has(self.initProvider.responsibilities))",message="spec.forProvider.responsibilities is a required parameter"
	Spec   OwnershipTeamsSpec   `json:"spec"`
	Status OwnershipTeamsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OwnershipTeamsList contains a list of OwnershipTeamss
type OwnershipTeamsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OwnershipTeams `json:"items"`
}

// Repository type metadata.
var (
	OwnershipTeams_Kind             = "OwnershipTeams"
	OwnershipTeams_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OwnershipTeams_Kind}.String()
	OwnershipTeams_KindAPIVersion   = OwnershipTeams_Kind + "." + CRDGroupVersion.String()
	OwnershipTeams_GroupVersionKind = CRDGroupVersion.WithKind(OwnershipTeams_Kind)
)

func init() {
	SchemeBuilder.Register(&OwnershipTeams{}, &OwnershipTeamsList{})
}