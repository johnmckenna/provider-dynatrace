// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this MonitoredTechnologiesApache.
func (mg *MonitoredTechnologiesApache) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitoredTechnologiesApache.
func (mg *MonitoredTechnologiesApache) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitoredTechnologiesApache.
func (mg *MonitoredTechnologiesApache) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitoredTechnologiesApache.
func (mg *MonitoredTechnologiesApache) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MonitoredTechnologiesApache.
func (mg *MonitoredTechnologiesApache) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitoredTechnologiesApache.
func (mg *MonitoredTechnologiesApache) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitoredTechnologiesApache.
func (mg *MonitoredTechnologiesApache) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitoredTechnologiesApache.
func (mg *MonitoredTechnologiesApache) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitoredTechnologiesApache.
func (mg *MonitoredTechnologiesApache) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitoredTechnologiesApache.
func (mg *MonitoredTechnologiesApache) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MonitoredTechnologiesApache.
func (mg *MonitoredTechnologiesApache) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitoredTechnologiesApache.
func (mg *MonitoredTechnologiesApache) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitoredTechnologiesDotnet.
func (mg *MonitoredTechnologiesDotnet) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitoredTechnologiesDotnet.
func (mg *MonitoredTechnologiesDotnet) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitoredTechnologiesDotnet.
func (mg *MonitoredTechnologiesDotnet) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitoredTechnologiesDotnet.
func (mg *MonitoredTechnologiesDotnet) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MonitoredTechnologiesDotnet.
func (mg *MonitoredTechnologiesDotnet) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitoredTechnologiesDotnet.
func (mg *MonitoredTechnologiesDotnet) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitoredTechnologiesDotnet.
func (mg *MonitoredTechnologiesDotnet) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitoredTechnologiesDotnet.
func (mg *MonitoredTechnologiesDotnet) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitoredTechnologiesDotnet.
func (mg *MonitoredTechnologiesDotnet) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitoredTechnologiesDotnet.
func (mg *MonitoredTechnologiesDotnet) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MonitoredTechnologiesDotnet.
func (mg *MonitoredTechnologiesDotnet) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitoredTechnologiesDotnet.
func (mg *MonitoredTechnologiesDotnet) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitoredTechnologiesGo.
func (mg *MonitoredTechnologiesGo) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitoredTechnologiesGo.
func (mg *MonitoredTechnologiesGo) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitoredTechnologiesGo.
func (mg *MonitoredTechnologiesGo) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitoredTechnologiesGo.
func (mg *MonitoredTechnologiesGo) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MonitoredTechnologiesGo.
func (mg *MonitoredTechnologiesGo) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitoredTechnologiesGo.
func (mg *MonitoredTechnologiesGo) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitoredTechnologiesGo.
func (mg *MonitoredTechnologiesGo) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitoredTechnologiesGo.
func (mg *MonitoredTechnologiesGo) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitoredTechnologiesGo.
func (mg *MonitoredTechnologiesGo) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitoredTechnologiesGo.
func (mg *MonitoredTechnologiesGo) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MonitoredTechnologiesGo.
func (mg *MonitoredTechnologiesGo) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitoredTechnologiesGo.
func (mg *MonitoredTechnologiesGo) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitoredTechnologiesIis.
func (mg *MonitoredTechnologiesIis) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitoredTechnologiesIis.
func (mg *MonitoredTechnologiesIis) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitoredTechnologiesIis.
func (mg *MonitoredTechnologiesIis) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitoredTechnologiesIis.
func (mg *MonitoredTechnologiesIis) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MonitoredTechnologiesIis.
func (mg *MonitoredTechnologiesIis) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitoredTechnologiesIis.
func (mg *MonitoredTechnologiesIis) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitoredTechnologiesIis.
func (mg *MonitoredTechnologiesIis) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitoredTechnologiesIis.
func (mg *MonitoredTechnologiesIis) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitoredTechnologiesIis.
func (mg *MonitoredTechnologiesIis) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitoredTechnologiesIis.
func (mg *MonitoredTechnologiesIis) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MonitoredTechnologiesIis.
func (mg *MonitoredTechnologiesIis) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitoredTechnologiesIis.
func (mg *MonitoredTechnologiesIis) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitoredTechnologiesJava.
func (mg *MonitoredTechnologiesJava) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitoredTechnologiesJava.
func (mg *MonitoredTechnologiesJava) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitoredTechnologiesJava.
func (mg *MonitoredTechnologiesJava) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitoredTechnologiesJava.
func (mg *MonitoredTechnologiesJava) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MonitoredTechnologiesJava.
func (mg *MonitoredTechnologiesJava) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitoredTechnologiesJava.
func (mg *MonitoredTechnologiesJava) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitoredTechnologiesJava.
func (mg *MonitoredTechnologiesJava) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitoredTechnologiesJava.
func (mg *MonitoredTechnologiesJava) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitoredTechnologiesJava.
func (mg *MonitoredTechnologiesJava) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitoredTechnologiesJava.
func (mg *MonitoredTechnologiesJava) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MonitoredTechnologiesJava.
func (mg *MonitoredTechnologiesJava) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitoredTechnologiesJava.
func (mg *MonitoredTechnologiesJava) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitoredTechnologiesNginx.
func (mg *MonitoredTechnologiesNginx) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitoredTechnologiesNginx.
func (mg *MonitoredTechnologiesNginx) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitoredTechnologiesNginx.
func (mg *MonitoredTechnologiesNginx) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitoredTechnologiesNginx.
func (mg *MonitoredTechnologiesNginx) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MonitoredTechnologiesNginx.
func (mg *MonitoredTechnologiesNginx) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitoredTechnologiesNginx.
func (mg *MonitoredTechnologiesNginx) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitoredTechnologiesNginx.
func (mg *MonitoredTechnologiesNginx) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitoredTechnologiesNginx.
func (mg *MonitoredTechnologiesNginx) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitoredTechnologiesNginx.
func (mg *MonitoredTechnologiesNginx) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitoredTechnologiesNginx.
func (mg *MonitoredTechnologiesNginx) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MonitoredTechnologiesNginx.
func (mg *MonitoredTechnologiesNginx) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitoredTechnologiesNginx.
func (mg *MonitoredTechnologiesNginx) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitoredTechnologiesNodeJS.
func (mg *MonitoredTechnologiesNodeJS) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitoredTechnologiesNodeJS.
func (mg *MonitoredTechnologiesNodeJS) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitoredTechnologiesNodeJS.
func (mg *MonitoredTechnologiesNodeJS) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitoredTechnologiesNodeJS.
func (mg *MonitoredTechnologiesNodeJS) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MonitoredTechnologiesNodeJS.
func (mg *MonitoredTechnologiesNodeJS) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitoredTechnologiesNodeJS.
func (mg *MonitoredTechnologiesNodeJS) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitoredTechnologiesNodeJS.
func (mg *MonitoredTechnologiesNodeJS) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitoredTechnologiesNodeJS.
func (mg *MonitoredTechnologiesNodeJS) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitoredTechnologiesNodeJS.
func (mg *MonitoredTechnologiesNodeJS) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitoredTechnologiesNodeJS.
func (mg *MonitoredTechnologiesNodeJS) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MonitoredTechnologiesNodeJS.
func (mg *MonitoredTechnologiesNodeJS) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitoredTechnologiesNodeJS.
func (mg *MonitoredTechnologiesNodeJS) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitoredTechnologiesOpentracing.
func (mg *MonitoredTechnologiesOpentracing) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitoredTechnologiesOpentracing.
func (mg *MonitoredTechnologiesOpentracing) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitoredTechnologiesOpentracing.
func (mg *MonitoredTechnologiesOpentracing) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitoredTechnologiesOpentracing.
func (mg *MonitoredTechnologiesOpentracing) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MonitoredTechnologiesOpentracing.
func (mg *MonitoredTechnologiesOpentracing) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitoredTechnologiesOpentracing.
func (mg *MonitoredTechnologiesOpentracing) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitoredTechnologiesOpentracing.
func (mg *MonitoredTechnologiesOpentracing) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitoredTechnologiesOpentracing.
func (mg *MonitoredTechnologiesOpentracing) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitoredTechnologiesOpentracing.
func (mg *MonitoredTechnologiesOpentracing) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitoredTechnologiesOpentracing.
func (mg *MonitoredTechnologiesOpentracing) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MonitoredTechnologiesOpentracing.
func (mg *MonitoredTechnologiesOpentracing) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitoredTechnologiesOpentracing.
func (mg *MonitoredTechnologiesOpentracing) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitoredTechnologiesPHP.
func (mg *MonitoredTechnologiesPHP) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitoredTechnologiesPHP.
func (mg *MonitoredTechnologiesPHP) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitoredTechnologiesPHP.
func (mg *MonitoredTechnologiesPHP) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitoredTechnologiesPHP.
func (mg *MonitoredTechnologiesPHP) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MonitoredTechnologiesPHP.
func (mg *MonitoredTechnologiesPHP) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitoredTechnologiesPHP.
func (mg *MonitoredTechnologiesPHP) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitoredTechnologiesPHP.
func (mg *MonitoredTechnologiesPHP) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitoredTechnologiesPHP.
func (mg *MonitoredTechnologiesPHP) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitoredTechnologiesPHP.
func (mg *MonitoredTechnologiesPHP) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitoredTechnologiesPHP.
func (mg *MonitoredTechnologiesPHP) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MonitoredTechnologiesPHP.
func (mg *MonitoredTechnologiesPHP) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitoredTechnologiesPHP.
func (mg *MonitoredTechnologiesPHP) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitoredTechnologiesVarnish.
func (mg *MonitoredTechnologiesVarnish) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitoredTechnologiesVarnish.
func (mg *MonitoredTechnologiesVarnish) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitoredTechnologiesVarnish.
func (mg *MonitoredTechnologiesVarnish) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitoredTechnologiesVarnish.
func (mg *MonitoredTechnologiesVarnish) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MonitoredTechnologiesVarnish.
func (mg *MonitoredTechnologiesVarnish) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitoredTechnologiesVarnish.
func (mg *MonitoredTechnologiesVarnish) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitoredTechnologiesVarnish.
func (mg *MonitoredTechnologiesVarnish) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitoredTechnologiesVarnish.
func (mg *MonitoredTechnologiesVarnish) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitoredTechnologiesVarnish.
func (mg *MonitoredTechnologiesVarnish) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitoredTechnologiesVarnish.
func (mg *MonitoredTechnologiesVarnish) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MonitoredTechnologiesVarnish.
func (mg *MonitoredTechnologiesVarnish) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitoredTechnologiesVarnish.
func (mg *MonitoredTechnologiesVarnish) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitoredTechnologiesWsmb.
func (mg *MonitoredTechnologiesWsmb) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitoredTechnologiesWsmb.
func (mg *MonitoredTechnologiesWsmb) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MonitoredTechnologiesWsmb.
func (mg *MonitoredTechnologiesWsmb) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MonitoredTechnologiesWsmb.
func (mg *MonitoredTechnologiesWsmb) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MonitoredTechnologiesWsmb.
func (mg *MonitoredTechnologiesWsmb) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MonitoredTechnologiesWsmb.
func (mg *MonitoredTechnologiesWsmb) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitoredTechnologiesWsmb.
func (mg *MonitoredTechnologiesWsmb) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitoredTechnologiesWsmb.
func (mg *MonitoredTechnologiesWsmb) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MonitoredTechnologiesWsmb.
func (mg *MonitoredTechnologiesWsmb) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MonitoredTechnologiesWsmb.
func (mg *MonitoredTechnologiesWsmb) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MonitoredTechnologiesWsmb.
func (mg *MonitoredTechnologiesWsmb) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MonitoredTechnologiesWsmb.
func (mg *MonitoredTechnologiesWsmb) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
