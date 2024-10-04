// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this BusinessEventsBuckets.
func (mg *BusinessEventsBuckets) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BusinessEventsBuckets.
func (mg *BusinessEventsBuckets) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this BusinessEventsBuckets.
func (mg *BusinessEventsBuckets) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this BusinessEventsBuckets.
func (mg *BusinessEventsBuckets) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this BusinessEventsBuckets.
func (mg *BusinessEventsBuckets) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this BusinessEventsBuckets.
func (mg *BusinessEventsBuckets) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BusinessEventsBuckets.
func (mg *BusinessEventsBuckets) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BusinessEventsBuckets.
func (mg *BusinessEventsBuckets) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this BusinessEventsBuckets.
func (mg *BusinessEventsBuckets) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this BusinessEventsBuckets.
func (mg *BusinessEventsBuckets) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this BusinessEventsBuckets.
func (mg *BusinessEventsBuckets) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this BusinessEventsBuckets.
func (mg *BusinessEventsBuckets) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BusinessEventsMetrics.
func (mg *BusinessEventsMetrics) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BusinessEventsMetrics.
func (mg *BusinessEventsMetrics) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this BusinessEventsMetrics.
func (mg *BusinessEventsMetrics) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this BusinessEventsMetrics.
func (mg *BusinessEventsMetrics) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this BusinessEventsMetrics.
func (mg *BusinessEventsMetrics) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this BusinessEventsMetrics.
func (mg *BusinessEventsMetrics) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BusinessEventsMetrics.
func (mg *BusinessEventsMetrics) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BusinessEventsMetrics.
func (mg *BusinessEventsMetrics) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this BusinessEventsMetrics.
func (mg *BusinessEventsMetrics) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this BusinessEventsMetrics.
func (mg *BusinessEventsMetrics) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this BusinessEventsMetrics.
func (mg *BusinessEventsMetrics) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this BusinessEventsMetrics.
func (mg *BusinessEventsMetrics) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BusinessEventsOneagent.
func (mg *BusinessEventsOneagent) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BusinessEventsOneagent.
func (mg *BusinessEventsOneagent) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this BusinessEventsOneagent.
func (mg *BusinessEventsOneagent) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this BusinessEventsOneagent.
func (mg *BusinessEventsOneagent) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this BusinessEventsOneagent.
func (mg *BusinessEventsOneagent) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this BusinessEventsOneagent.
func (mg *BusinessEventsOneagent) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BusinessEventsOneagent.
func (mg *BusinessEventsOneagent) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BusinessEventsOneagent.
func (mg *BusinessEventsOneagent) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this BusinessEventsOneagent.
func (mg *BusinessEventsOneagent) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this BusinessEventsOneagent.
func (mg *BusinessEventsOneagent) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this BusinessEventsOneagent.
func (mg *BusinessEventsOneagent) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this BusinessEventsOneagent.
func (mg *BusinessEventsOneagent) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BusinessEventsOneagentOutgoing.
func (mg *BusinessEventsOneagentOutgoing) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BusinessEventsOneagentOutgoing.
func (mg *BusinessEventsOneagentOutgoing) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this BusinessEventsOneagentOutgoing.
func (mg *BusinessEventsOneagentOutgoing) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this BusinessEventsOneagentOutgoing.
func (mg *BusinessEventsOneagentOutgoing) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this BusinessEventsOneagentOutgoing.
func (mg *BusinessEventsOneagentOutgoing) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this BusinessEventsOneagentOutgoing.
func (mg *BusinessEventsOneagentOutgoing) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BusinessEventsOneagentOutgoing.
func (mg *BusinessEventsOneagentOutgoing) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BusinessEventsOneagentOutgoing.
func (mg *BusinessEventsOneagentOutgoing) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this BusinessEventsOneagentOutgoing.
func (mg *BusinessEventsOneagentOutgoing) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this BusinessEventsOneagentOutgoing.
func (mg *BusinessEventsOneagentOutgoing) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this BusinessEventsOneagentOutgoing.
func (mg *BusinessEventsOneagentOutgoing) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this BusinessEventsOneagentOutgoing.
func (mg *BusinessEventsOneagentOutgoing) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BusinessEventsProcessing.
func (mg *BusinessEventsProcessing) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BusinessEventsProcessing.
func (mg *BusinessEventsProcessing) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this BusinessEventsProcessing.
func (mg *BusinessEventsProcessing) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this BusinessEventsProcessing.
func (mg *BusinessEventsProcessing) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this BusinessEventsProcessing.
func (mg *BusinessEventsProcessing) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this BusinessEventsProcessing.
func (mg *BusinessEventsProcessing) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BusinessEventsProcessing.
func (mg *BusinessEventsProcessing) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BusinessEventsProcessing.
func (mg *BusinessEventsProcessing) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this BusinessEventsProcessing.
func (mg *BusinessEventsProcessing) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this BusinessEventsProcessing.
func (mg *BusinessEventsProcessing) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this BusinessEventsProcessing.
func (mg *BusinessEventsProcessing) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this BusinessEventsProcessing.
func (mg *BusinessEventsProcessing) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BusinessEventsSecurityContext.
func (mg *BusinessEventsSecurityContext) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BusinessEventsSecurityContext.
func (mg *BusinessEventsSecurityContext) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this BusinessEventsSecurityContext.
func (mg *BusinessEventsSecurityContext) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this BusinessEventsSecurityContext.
func (mg *BusinessEventsSecurityContext) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this BusinessEventsSecurityContext.
func (mg *BusinessEventsSecurityContext) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this BusinessEventsSecurityContext.
func (mg *BusinessEventsSecurityContext) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BusinessEventsSecurityContext.
func (mg *BusinessEventsSecurityContext) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BusinessEventsSecurityContext.
func (mg *BusinessEventsSecurityContext) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this BusinessEventsSecurityContext.
func (mg *BusinessEventsSecurityContext) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this BusinessEventsSecurityContext.
func (mg *BusinessEventsSecurityContext) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this BusinessEventsSecurityContext.
func (mg *BusinessEventsSecurityContext) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this BusinessEventsSecurityContext.
func (mg *BusinessEventsSecurityContext) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}