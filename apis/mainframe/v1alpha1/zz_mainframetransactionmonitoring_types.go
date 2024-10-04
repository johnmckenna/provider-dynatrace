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

type MainframeTransactionMonitoringInitParameters struct {

	// (Boolean) If enabled, CICS regions belonging to the same CICSPlex will be grouped into a single process group. If disabled, a process group will be created for each CICS region.
	// If enabled, CICS regions belonging to the same CICSPlex will be grouped into a single process group. If disabled, a process group will be created for each CICS region.
	GroupCicsRegions *bool `json:"groupCicsRegions,omitempty" tf:"group_cics_regions,omitempty"`

	// (Boolean) If enabled, IMS regions belonging to the same subsystem will be grouped into a single process group. If disabled, a process group will be created for each IMS region.
	// If enabled, IMS regions belonging to the same subsystem will be grouped into a single process group. If disabled, a process group will be created for each IMS region.
	GroupImsRegions *bool `json:"groupImsRegions,omitempty" tf:"group_ims_regions,omitempty"`

	// (Boolean) If enabled, the CICS Transaction Gateway sensor will trace all EXCI requests including those that are using the TCP/IP or SNA protocol.
	// If enabled, the CICS Transaction Gateway sensor will trace all EXCI requests including those that are using the TCP/IP or SNA protocol.
	MonitorAllCtgProtocols *bool `json:"monitorAllCtgProtocols,omitempty" tf:"monitor_all_ctg_protocols,omitempty"`

	// monitored services. Enable this setting to monitor all incoming web requests. We recommend enabling it only over a short period of time.
	// Dynatrace automatically traces incoming web requests when they are called by already-monitored services. Enable this setting to monitor all incoming web requests. We recommend enabling it only over a short period of time.
	MonitorAllIncomingWebRequests *bool `json:"monitorAllIncomingWebRequests,omitempty" tf:"monitor_all_incoming_web_requests,omitempty"`

	// (Number) We recommend the default limit of 500 nodes. The value 0 means unlimited number of nodes.
	// We recommend the default limit of 500 nodes. The value 0 means unlimited number of nodes.
	NodeLimit *float64 `json:"nodeLimit,omitempty" tf:"node_limit,omitempty"`

	// (Boolean) If enabled, a CICS service will be created for each monitored transaction ID within a process group. If disabled, a CICS service will be created for each monitored CICS region within a process group. We recommend enabling it only when the CICS regions are grouped by their CICSPlex.
	// If enabled, a CICS service will be created for each monitored transaction ID within a process group. If disabled, a CICS service will be created for each monitored CICS region within a process group. We recommend enabling it only when the CICS regions are grouped by their CICSPlex.
	ZosCicsServiceDetectionUsesTransactionID *bool `json:"zosCicsServiceDetectionUsesTransactionId,omitempty" tf:"zos_cics_service_detection_uses_transaction_id,omitempty"`

	// (Boolean) If enabled, an IMS service will be created for each monitored transaction ID within a process group. If disabled, an IMS service will be created for each monitored IMS region within a process group. We recommend enabling it only when the IMS regions are grouped by their subsystem.
	// If enabled, an IMS service will be created for each monitored transaction ID within a process group. If disabled, an IMS service will be created for each monitored IMS region within a process group. We recommend enabling it only when the IMS regions are grouped by their subsystem.
	ZosImsServiceDetectionUsesTransactionID *bool `json:"zosImsServiceDetectionUsesTransactionId,omitempty" tf:"zos_ims_service_detection_uses_transaction_id,omitempty"`
}

type MainframeTransactionMonitoringObservation struct {

	// (Boolean) If enabled, CICS regions belonging to the same CICSPlex will be grouped into a single process group. If disabled, a process group will be created for each CICS region.
	// If enabled, CICS regions belonging to the same CICSPlex will be grouped into a single process group. If disabled, a process group will be created for each CICS region.
	GroupCicsRegions *bool `json:"groupCicsRegions,omitempty" tf:"group_cics_regions,omitempty"`

	// (Boolean) If enabled, IMS regions belonging to the same subsystem will be grouped into a single process group. If disabled, a process group will be created for each IMS region.
	// If enabled, IMS regions belonging to the same subsystem will be grouped into a single process group. If disabled, a process group will be created for each IMS region.
	GroupImsRegions *bool `json:"groupImsRegions,omitempty" tf:"group_ims_regions,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) If enabled, the CICS Transaction Gateway sensor will trace all EXCI requests including those that are using the TCP/IP or SNA protocol.
	// If enabled, the CICS Transaction Gateway sensor will trace all EXCI requests including those that are using the TCP/IP or SNA protocol.
	MonitorAllCtgProtocols *bool `json:"monitorAllCtgProtocols,omitempty" tf:"monitor_all_ctg_protocols,omitempty"`

	// monitored services. Enable this setting to monitor all incoming web requests. We recommend enabling it only over a short period of time.
	// Dynatrace automatically traces incoming web requests when they are called by already-monitored services. Enable this setting to monitor all incoming web requests. We recommend enabling it only over a short period of time.
	MonitorAllIncomingWebRequests *bool `json:"monitorAllIncomingWebRequests,omitempty" tf:"monitor_all_incoming_web_requests,omitempty"`

	// (Number) We recommend the default limit of 500 nodes. The value 0 means unlimited number of nodes.
	// We recommend the default limit of 500 nodes. The value 0 means unlimited number of nodes.
	NodeLimit *float64 `json:"nodeLimit,omitempty" tf:"node_limit,omitempty"`

	// (Boolean) If enabled, a CICS service will be created for each monitored transaction ID within a process group. If disabled, a CICS service will be created for each monitored CICS region within a process group. We recommend enabling it only when the CICS regions are grouped by their CICSPlex.
	// If enabled, a CICS service will be created for each monitored transaction ID within a process group. If disabled, a CICS service will be created for each monitored CICS region within a process group. We recommend enabling it only when the CICS regions are grouped by their CICSPlex.
	ZosCicsServiceDetectionUsesTransactionID *bool `json:"zosCicsServiceDetectionUsesTransactionId,omitempty" tf:"zos_cics_service_detection_uses_transaction_id,omitempty"`

	// (Boolean) If enabled, an IMS service will be created for each monitored transaction ID within a process group. If disabled, an IMS service will be created for each monitored IMS region within a process group. We recommend enabling it only when the IMS regions are grouped by their subsystem.
	// If enabled, an IMS service will be created for each monitored transaction ID within a process group. If disabled, an IMS service will be created for each monitored IMS region within a process group. We recommend enabling it only when the IMS regions are grouped by their subsystem.
	ZosImsServiceDetectionUsesTransactionID *bool `json:"zosImsServiceDetectionUsesTransactionId,omitempty" tf:"zos_ims_service_detection_uses_transaction_id,omitempty"`
}

type MainframeTransactionMonitoringParameters struct {

	// (Boolean) If enabled, CICS regions belonging to the same CICSPlex will be grouped into a single process group. If disabled, a process group will be created for each CICS region.
	// If enabled, CICS regions belonging to the same CICSPlex will be grouped into a single process group. If disabled, a process group will be created for each CICS region.
	// +kubebuilder:validation:Optional
	GroupCicsRegions *bool `json:"groupCicsRegions,omitempty" tf:"group_cics_regions,omitempty"`

	// (Boolean) If enabled, IMS regions belonging to the same subsystem will be grouped into a single process group. If disabled, a process group will be created for each IMS region.
	// If enabled, IMS regions belonging to the same subsystem will be grouped into a single process group. If disabled, a process group will be created for each IMS region.
	// +kubebuilder:validation:Optional
	GroupImsRegions *bool `json:"groupImsRegions,omitempty" tf:"group_ims_regions,omitempty"`

	// (Boolean) If enabled, the CICS Transaction Gateway sensor will trace all EXCI requests including those that are using the TCP/IP or SNA protocol.
	// If enabled, the CICS Transaction Gateway sensor will trace all EXCI requests including those that are using the TCP/IP or SNA protocol.
	// +kubebuilder:validation:Optional
	MonitorAllCtgProtocols *bool `json:"monitorAllCtgProtocols,omitempty" tf:"monitor_all_ctg_protocols,omitempty"`

	// monitored services. Enable this setting to monitor all incoming web requests. We recommend enabling it only over a short period of time.
	// Dynatrace automatically traces incoming web requests when they are called by already-monitored services. Enable this setting to monitor all incoming web requests. We recommend enabling it only over a short period of time.
	// +kubebuilder:validation:Optional
	MonitorAllIncomingWebRequests *bool `json:"monitorAllIncomingWebRequests,omitempty" tf:"monitor_all_incoming_web_requests,omitempty"`

	// (Number) We recommend the default limit of 500 nodes. The value 0 means unlimited number of nodes.
	// We recommend the default limit of 500 nodes. The value 0 means unlimited number of nodes.
	// +kubebuilder:validation:Optional
	NodeLimit *float64 `json:"nodeLimit,omitempty" tf:"node_limit,omitempty"`

	// (Boolean) If enabled, a CICS service will be created for each monitored transaction ID within a process group. If disabled, a CICS service will be created for each monitored CICS region within a process group. We recommend enabling it only when the CICS regions are grouped by their CICSPlex.
	// If enabled, a CICS service will be created for each monitored transaction ID within a process group. If disabled, a CICS service will be created for each monitored CICS region within a process group. We recommend enabling it only when the CICS regions are grouped by their CICSPlex.
	// +kubebuilder:validation:Optional
	ZosCicsServiceDetectionUsesTransactionID *bool `json:"zosCicsServiceDetectionUsesTransactionId,omitempty" tf:"zos_cics_service_detection_uses_transaction_id,omitempty"`

	// (Boolean) If enabled, an IMS service will be created for each monitored transaction ID within a process group. If disabled, an IMS service will be created for each monitored IMS region within a process group. We recommend enabling it only when the IMS regions are grouped by their subsystem.
	// If enabled, an IMS service will be created for each monitored transaction ID within a process group. If disabled, an IMS service will be created for each monitored IMS region within a process group. We recommend enabling it only when the IMS regions are grouped by their subsystem.
	// +kubebuilder:validation:Optional
	ZosImsServiceDetectionUsesTransactionID *bool `json:"zosImsServiceDetectionUsesTransactionId,omitempty" tf:"zos_ims_service_detection_uses_transaction_id,omitempty"`
}

// MainframeTransactionMonitoringSpec defines the desired state of MainframeTransactionMonitoring
type MainframeTransactionMonitoringSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MainframeTransactionMonitoringParameters `json:"forProvider"`
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
	InitProvider MainframeTransactionMonitoringInitParameters `json:"initProvider,omitempty"`
}

// MainframeTransactionMonitoringStatus defines the observed state of MainframeTransactionMonitoring.
type MainframeTransactionMonitoringStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MainframeTransactionMonitoringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MainframeTransactionMonitoring is the Schema for the MainframeTransactionMonitorings API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type MainframeTransactionMonitoring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupCicsRegions) || (has(self.initProvider) && has(self.initProvider.groupCicsRegions))",message="spec.forProvider.groupCicsRegions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupImsRegions) || (has(self.initProvider) && has(self.initProvider.groupImsRegions))",message="spec.forProvider.groupImsRegions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.monitorAllCtgProtocols) || (has(self.initProvider) && has(self.initProvider.monitorAllCtgProtocols))",message="spec.forProvider.monitorAllCtgProtocols is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.monitorAllIncomingWebRequests) || (has(self.initProvider) && has(self.initProvider.monitorAllIncomingWebRequests))",message="spec.forProvider.monitorAllIncomingWebRequests is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nodeLimit) || (has(self.initProvider) && has(self.initProvider.nodeLimit))",message="spec.forProvider.nodeLimit is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zosCicsServiceDetectionUsesTransactionId) || (has(self.initProvider) && has(self.initProvider.zosCicsServiceDetectionUsesTransactionId))",message="spec.forProvider.zosCicsServiceDetectionUsesTransactionId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zosImsServiceDetectionUsesTransactionId) || (has(self.initProvider) && has(self.initProvider.zosImsServiceDetectionUsesTransactionId))",message="spec.forProvider.zosImsServiceDetectionUsesTransactionId is a required parameter"
	Spec   MainframeTransactionMonitoringSpec   `json:"spec"`
	Status MainframeTransactionMonitoringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MainframeTransactionMonitoringList contains a list of MainframeTransactionMonitorings
type MainframeTransactionMonitoringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MainframeTransactionMonitoring `json:"items"`
}

// Repository type metadata.
var (
	MainframeTransactionMonitoring_Kind             = "MainframeTransactionMonitoring"
	MainframeTransactionMonitoring_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MainframeTransactionMonitoring_Kind}.String()
	MainframeTransactionMonitoring_KindAPIVersion   = MainframeTransactionMonitoring_Kind + "." + CRDGroupVersion.String()
	MainframeTransactionMonitoring_GroupVersionKind = CRDGroupVersion.WithKind(MainframeTransactionMonitoring_Kind)
)

func init() {
	SchemeBuilder.Register(&MainframeTransactionMonitoring{}, &MainframeTransactionMonitoringList{})
}
