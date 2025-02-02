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

type FrequentIssuesInitParameters struct {

	// (Boolean) Detect frequent issues within applications, enabled (true) or disabled (false)
	// Detect frequent issues within applications, enabled (`true`) or disabled (`false`)
	DetectApps *bool `json:"detectApps,omitempty" tf:"detect_apps,omitempty"`

	// (Boolean) Detect frequent issues within infrastructure, enabled (true) or disabled (false)
	// Detect frequent issues within infrastructure, enabled (`true`) or disabled (`false`)
	DetectInfra *bool `json:"detectInfra,omitempty" tf:"detect_infra,omitempty"`

	// (Boolean) Detect frequent issues within transactions and services, enabled (true) or disabled (false)
	// Detect frequent issues within transactions and services, enabled (`true`) or disabled (`false`)
	DetectTxn *bool `json:"detectTxn,omitempty" tf:"detect_txn,omitempty"`
}

type FrequentIssuesObservation struct {

	// (Boolean) Detect frequent issues within applications, enabled (true) or disabled (false)
	// Detect frequent issues within applications, enabled (`true`) or disabled (`false`)
	DetectApps *bool `json:"detectApps,omitempty" tf:"detect_apps,omitempty"`

	// (Boolean) Detect frequent issues within infrastructure, enabled (true) or disabled (false)
	// Detect frequent issues within infrastructure, enabled (`true`) or disabled (`false`)
	DetectInfra *bool `json:"detectInfra,omitempty" tf:"detect_infra,omitempty"`

	// (Boolean) Detect frequent issues within transactions and services, enabled (true) or disabled (false)
	// Detect frequent issues within transactions and services, enabled (`true`) or disabled (`false`)
	DetectTxn *bool `json:"detectTxn,omitempty" tf:"detect_txn,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FrequentIssuesParameters struct {

	// (Boolean) Detect frequent issues within applications, enabled (true) or disabled (false)
	// Detect frequent issues within applications, enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	DetectApps *bool `json:"detectApps,omitempty" tf:"detect_apps,omitempty"`

	// (Boolean) Detect frequent issues within infrastructure, enabled (true) or disabled (false)
	// Detect frequent issues within infrastructure, enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	DetectInfra *bool `json:"detectInfra,omitempty" tf:"detect_infra,omitempty"`

	// (Boolean) Detect frequent issues within transactions and services, enabled (true) or disabled (false)
	// Detect frequent issues within transactions and services, enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	DetectTxn *bool `json:"detectTxn,omitempty" tf:"detect_txn,omitempty"`
}

// FrequentIssuesSpec defines the desired state of FrequentIssues
type FrequentIssuesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrequentIssuesParameters `json:"forProvider"`
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
	InitProvider FrequentIssuesInitParameters `json:"initProvider,omitempty"`
}

// FrequentIssuesStatus defines the observed state of FrequentIssues.
type FrequentIssuesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrequentIssuesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FrequentIssues is the Schema for the FrequentIssuess API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type FrequentIssues struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.detectApps) || (has(self.initProvider) && has(self.initProvider.detectApps))",message="spec.forProvider.detectApps is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.detectInfra) || (has(self.initProvider) && has(self.initProvider.detectInfra))",message="spec.forProvider.detectInfra is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.detectTxn) || (has(self.initProvider) && has(self.initProvider.detectTxn))",message="spec.forProvider.detectTxn is a required parameter"
	Spec   FrequentIssuesSpec   `json:"spec"`
	Status FrequentIssuesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrequentIssuesList contains a list of FrequentIssuess
type FrequentIssuesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrequentIssues `json:"items"`
}

// Repository type metadata.
var (
	FrequentIssues_Kind             = "FrequentIssues"
	FrequentIssues_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrequentIssues_Kind}.String()
	FrequentIssues_KindAPIVersion   = FrequentIssues_Kind + "." + CRDGroupVersion.String()
	FrequentIssues_GroupVersionKind = CRDGroupVersion.WithKind(FrequentIssues_Kind)
)

func init() {
	SchemeBuilder.Register(&FrequentIssues{}, &FrequentIssuesList{})
}
