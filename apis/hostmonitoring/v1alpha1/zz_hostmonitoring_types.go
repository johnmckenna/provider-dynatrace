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

type HostMonitoringInitParameters struct {

	// injection disabled with oneagentctl takes precedence over this setting and cannot be changed from the Dynatrace web UI.
	// An auto-injection disabled with [oneagentctl](https://dt-url.net/oneagentctl) takes precedence over this setting and cannot be changed from the Dynatrace web UI.
	AutoInjection *bool `json:"autoInjection,omitempty" tf:"auto_injection,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// stack monitoring by default, to monitor every aspect of your environment, including all processes, services, and applications detected on your hosts.
	// Dynatrace uses full-stack monitoring by default, to monitor every aspect of your environment, including all processes, services, and applications detected on your hosts.
	//
	// If you turn off full-stack monitoring, Dynatrace will only monitor your infrastructure. You will lose access to application performance, user experience data, code-level visibility and PurePath insights.
	//
	// To learn more, visit [Infrastructure Monitoring mode](https://www.dynatrace.com/support/help/shortlink/infrastructure).
	//
	// Please note that changing the monitoring mode will impact the license consumption of this OneAgent. To learn more, visit [Host units](https://dt-url.net/hi03uns).
	FullStack *bool `json:"fullStack,omitempty" tf:"full_stack,omitempty"`

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`
}

type HostMonitoringObservation struct {

	// injection disabled with oneagentctl takes precedence over this setting and cannot be changed from the Dynatrace web UI.
	// An auto-injection disabled with [oneagentctl](https://dt-url.net/oneagentctl) takes precedence over this setting and cannot be changed from the Dynatrace web UI.
	AutoInjection *bool `json:"autoInjection,omitempty" tf:"auto_injection,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// stack monitoring by default, to monitor every aspect of your environment, including all processes, services, and applications detected on your hosts.
	// Dynatrace uses full-stack monitoring by default, to monitor every aspect of your environment, including all processes, services, and applications detected on your hosts.
	//
	// If you turn off full-stack monitoring, Dynatrace will only monitor your infrastructure. You will lose access to application performance, user experience data, code-level visibility and PurePath insights.
	//
	// To learn more, visit [Infrastructure Monitoring mode](https://www.dynatrace.com/support/help/shortlink/infrastructure).
	//
	// Please note that changing the monitoring mode will impact the license consumption of this OneAgent. To learn more, visit [Host units](https://dt-url.net/hi03uns).
	FullStack *bool `json:"fullStack,omitempty" tf:"full_stack,omitempty"`

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HostMonitoringParameters struct {

	// injection disabled with oneagentctl takes precedence over this setting and cannot be changed from the Dynatrace web UI.
	// An auto-injection disabled with [oneagentctl](https://dt-url.net/oneagentctl) takes precedence over this setting and cannot be changed from the Dynatrace web UI.
	// +kubebuilder:validation:Optional
	AutoInjection *bool `json:"autoInjection,omitempty" tf:"auto_injection,omitempty"`

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// stack monitoring by default, to monitor every aspect of your environment, including all processes, services, and applications detected on your hosts.
	// Dynatrace uses full-stack monitoring by default, to monitor every aspect of your environment, including all processes, services, and applications detected on your hosts.
	//
	// If you turn off full-stack monitoring, Dynatrace will only monitor your infrastructure. You will lose access to application performance, user experience data, code-level visibility and PurePath insights.
	//
	// To learn more, visit [Infrastructure Monitoring mode](https://www.dynatrace.com/support/help/shortlink/infrastructure).
	//
	// Please note that changing the monitoring mode will impact the license consumption of this OneAgent. To learn more, visit [Host units](https://dt-url.net/hi03uns).
	// +kubebuilder:validation:Optional
	FullStack *bool `json:"fullStack,omitempty" tf:"full_stack,omitempty"`

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// +kubebuilder:validation:Optional
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`
}

// HostMonitoringSpec defines the desired state of HostMonitoring
type HostMonitoringSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostMonitoringParameters `json:"forProvider"`
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
	InitProvider HostMonitoringInitParameters `json:"initProvider,omitempty"`
}

// HostMonitoringStatus defines the observed state of HostMonitoring.
type HostMonitoringStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostMonitoringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HostMonitoring is the Schema for the HostMonitorings API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type HostMonitoring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostId) || (has(self.initProvider) && has(self.initProvider.hostId))",message="spec.forProvider.hostId is a required parameter"
	Spec   HostMonitoringSpec   `json:"spec"`
	Status HostMonitoringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostMonitoringList contains a list of HostMonitorings
type HostMonitoringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostMonitoring `json:"items"`
}

// Repository type metadata.
var (
	HostMonitoring_Kind             = "HostMonitoring"
	HostMonitoring_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostMonitoring_Kind}.String()
	HostMonitoring_KindAPIVersion   = HostMonitoring_Kind + "." + CRDGroupVersion.String()
	HostMonitoring_GroupVersionKind = CRDGroupVersion.WithKind(HostMonitoring_Kind)
)

func init() {
	SchemeBuilder.Register(&HostMonitoring{}, &HostMonitoringList{})
}
