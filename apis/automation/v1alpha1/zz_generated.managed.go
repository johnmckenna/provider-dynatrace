// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this AutomationBusinessCalendar.
func (mg *AutomationBusinessCalendar) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AutomationBusinessCalendar.
func (mg *AutomationBusinessCalendar) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this AutomationBusinessCalendar.
func (mg *AutomationBusinessCalendar) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this AutomationBusinessCalendar.
func (mg *AutomationBusinessCalendar) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this AutomationBusinessCalendar.
func (mg *AutomationBusinessCalendar) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this AutomationBusinessCalendar.
func (mg *AutomationBusinessCalendar) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AutomationBusinessCalendar.
func (mg *AutomationBusinessCalendar) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AutomationBusinessCalendar.
func (mg *AutomationBusinessCalendar) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this AutomationBusinessCalendar.
func (mg *AutomationBusinessCalendar) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this AutomationBusinessCalendar.
func (mg *AutomationBusinessCalendar) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this AutomationBusinessCalendar.
func (mg *AutomationBusinessCalendar) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this AutomationBusinessCalendar.
func (mg *AutomationBusinessCalendar) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AutomationSchedulingRule.
func (mg *AutomationSchedulingRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AutomationSchedulingRule.
func (mg *AutomationSchedulingRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this AutomationSchedulingRule.
func (mg *AutomationSchedulingRule) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this AutomationSchedulingRule.
func (mg *AutomationSchedulingRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this AutomationSchedulingRule.
func (mg *AutomationSchedulingRule) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this AutomationSchedulingRule.
func (mg *AutomationSchedulingRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AutomationSchedulingRule.
func (mg *AutomationSchedulingRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AutomationSchedulingRule.
func (mg *AutomationSchedulingRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this AutomationSchedulingRule.
func (mg *AutomationSchedulingRule) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this AutomationSchedulingRule.
func (mg *AutomationSchedulingRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this AutomationSchedulingRule.
func (mg *AutomationSchedulingRule) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this AutomationSchedulingRule.
func (mg *AutomationSchedulingRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AutomationWorkflow.
func (mg *AutomationWorkflow) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AutomationWorkflow.
func (mg *AutomationWorkflow) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this AutomationWorkflow.
func (mg *AutomationWorkflow) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this AutomationWorkflow.
func (mg *AutomationWorkflow) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this AutomationWorkflow.
func (mg *AutomationWorkflow) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this AutomationWorkflow.
func (mg *AutomationWorkflow) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AutomationWorkflow.
func (mg *AutomationWorkflow) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AutomationWorkflow.
func (mg *AutomationWorkflow) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this AutomationWorkflow.
func (mg *AutomationWorkflow) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this AutomationWorkflow.
func (mg *AutomationWorkflow) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this AutomationWorkflow.
func (mg *AutomationWorkflow) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this AutomationWorkflow.
func (mg *AutomationWorkflow) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AutomationWorkflowAwsConnections.
func (mg *AutomationWorkflowAwsConnections) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AutomationWorkflowAwsConnections.
func (mg *AutomationWorkflowAwsConnections) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this AutomationWorkflowAwsConnections.
func (mg *AutomationWorkflowAwsConnections) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this AutomationWorkflowAwsConnections.
func (mg *AutomationWorkflowAwsConnections) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this AutomationWorkflowAwsConnections.
func (mg *AutomationWorkflowAwsConnections) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this AutomationWorkflowAwsConnections.
func (mg *AutomationWorkflowAwsConnections) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AutomationWorkflowAwsConnections.
func (mg *AutomationWorkflowAwsConnections) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AutomationWorkflowAwsConnections.
func (mg *AutomationWorkflowAwsConnections) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this AutomationWorkflowAwsConnections.
func (mg *AutomationWorkflowAwsConnections) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this AutomationWorkflowAwsConnections.
func (mg *AutomationWorkflowAwsConnections) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this AutomationWorkflowAwsConnections.
func (mg *AutomationWorkflowAwsConnections) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this AutomationWorkflowAwsConnections.
func (mg *AutomationWorkflowAwsConnections) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AutomationWorkflowJira.
func (mg *AutomationWorkflowJira) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AutomationWorkflowJira.
func (mg *AutomationWorkflowJira) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this AutomationWorkflowJira.
func (mg *AutomationWorkflowJira) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this AutomationWorkflowJira.
func (mg *AutomationWorkflowJira) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this AutomationWorkflowJira.
func (mg *AutomationWorkflowJira) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this AutomationWorkflowJira.
func (mg *AutomationWorkflowJira) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AutomationWorkflowJira.
func (mg *AutomationWorkflowJira) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AutomationWorkflowJira.
func (mg *AutomationWorkflowJira) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this AutomationWorkflowJira.
func (mg *AutomationWorkflowJira) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this AutomationWorkflowJira.
func (mg *AutomationWorkflowJira) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this AutomationWorkflowJira.
func (mg *AutomationWorkflowJira) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this AutomationWorkflowJira.
func (mg *AutomationWorkflowJira) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AutomationWorkflowK8SConnections.
func (mg *AutomationWorkflowK8SConnections) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AutomationWorkflowK8SConnections.
func (mg *AutomationWorkflowK8SConnections) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this AutomationWorkflowK8SConnections.
func (mg *AutomationWorkflowK8SConnections) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this AutomationWorkflowK8SConnections.
func (mg *AutomationWorkflowK8SConnections) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this AutomationWorkflowK8SConnections.
func (mg *AutomationWorkflowK8SConnections) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this AutomationWorkflowK8SConnections.
func (mg *AutomationWorkflowK8SConnections) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AutomationWorkflowK8SConnections.
func (mg *AutomationWorkflowK8SConnections) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AutomationWorkflowK8SConnections.
func (mg *AutomationWorkflowK8SConnections) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this AutomationWorkflowK8SConnections.
func (mg *AutomationWorkflowK8SConnections) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this AutomationWorkflowK8SConnections.
func (mg *AutomationWorkflowK8SConnections) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this AutomationWorkflowK8SConnections.
func (mg *AutomationWorkflowK8SConnections) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this AutomationWorkflowK8SConnections.
func (mg *AutomationWorkflowK8SConnections) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AutomationWorkflowSlack.
func (mg *AutomationWorkflowSlack) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AutomationWorkflowSlack.
func (mg *AutomationWorkflowSlack) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this AutomationWorkflowSlack.
func (mg *AutomationWorkflowSlack) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this AutomationWorkflowSlack.
func (mg *AutomationWorkflowSlack) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this AutomationWorkflowSlack.
func (mg *AutomationWorkflowSlack) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this AutomationWorkflowSlack.
func (mg *AutomationWorkflowSlack) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AutomationWorkflowSlack.
func (mg *AutomationWorkflowSlack) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AutomationWorkflowSlack.
func (mg *AutomationWorkflowSlack) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this AutomationWorkflowSlack.
func (mg *AutomationWorkflowSlack) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this AutomationWorkflowSlack.
func (mg *AutomationWorkflowSlack) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this AutomationWorkflowSlack.
func (mg *AutomationWorkflowSlack) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this AutomationWorkflowSlack.
func (mg *AutomationWorkflowSlack) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SiteReliabilityGuardian.
func (mg *SiteReliabilityGuardian) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SiteReliabilityGuardian.
func (mg *SiteReliabilityGuardian) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this SiteReliabilityGuardian.
func (mg *SiteReliabilityGuardian) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this SiteReliabilityGuardian.
func (mg *SiteReliabilityGuardian) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this SiteReliabilityGuardian.
func (mg *SiteReliabilityGuardian) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SiteReliabilityGuardian.
func (mg *SiteReliabilityGuardian) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SiteReliabilityGuardian.
func (mg *SiteReliabilityGuardian) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SiteReliabilityGuardian.
func (mg *SiteReliabilityGuardian) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this SiteReliabilityGuardian.
func (mg *SiteReliabilityGuardian) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this SiteReliabilityGuardian.
func (mg *SiteReliabilityGuardian) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this SiteReliabilityGuardian.
func (mg *SiteReliabilityGuardian) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SiteReliabilityGuardian.
func (mg *SiteReliabilityGuardian) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
