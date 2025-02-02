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

type HubExtensionConfigInitParameters struct {

	// (String) The name of the Active Gate Group this monitoring configuration will be defined for
	// The name of the Active Gate Group this monitoring configuration will be defined for
	ActiveGateGroup *string `json:"activeGateGroup,omitempty" tf:"active_gate_group,omitempty"`

	// (String) The ID of the host this monitoring configuration will be defined for
	// The ID of the host this monitoring configuration will be defined for
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// (String) The ID of the host group this monitoring configuration will be defined for
	// The ID of the host group this monitoring configuration will be defined for
	HostGroup *string `json:"hostGroup,omitempty" tf:"host_group,omitempty"`

	// (String) The name of the Management Zone this monitoring configuration will be defined for
	// The name of the Management Zone this monitoring configuration will be defined for
	ManagementZone *string `json:"managementZone,omitempty" tf:"management_zone,omitempty"`

	// liberty-cp. You can query for these names using the data source dynatrace_hub_items
	// The fully qualified name of the extension, such as `com.dynatrace.extension.jmx-liberty-cp`. You can query for these names using the data source `dynatrace_hub_items`
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The JSON encoded value for this monitoring configuration
	// The JSON encoded value for this monitoring configuration
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type HubExtensionConfigObservation struct {

	// (String) The name of the Active Gate Group this monitoring configuration will be defined for
	// The name of the Active Gate Group this monitoring configuration will be defined for
	ActiveGateGroup *string `json:"activeGateGroup,omitempty" tf:"active_gate_group,omitempty"`

	// (String) The ID of the host this monitoring configuration will be defined for
	// The ID of the host this monitoring configuration will be defined for
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// (String) The ID of the host group this monitoring configuration will be defined for
	// The ID of the host group this monitoring configuration will be defined for
	HostGroup *string `json:"hostGroup,omitempty" tf:"host_group,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the Management Zone this monitoring configuration will be defined for
	// The name of the Management Zone this monitoring configuration will be defined for
	ManagementZone *string `json:"managementZone,omitempty" tf:"management_zone,omitempty"`

	// liberty-cp. You can query for these names using the data source dynatrace_hub_items
	// The fully qualified name of the extension, such as `com.dynatrace.extension.jmx-liberty-cp`. You can query for these names using the data source `dynatrace_hub_items`
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The JSON encoded value for this monitoring configuration
	// The JSON encoded value for this monitoring configuration
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type HubExtensionConfigParameters struct {

	// (String) The name of the Active Gate Group this monitoring configuration will be defined for
	// The name of the Active Gate Group this monitoring configuration will be defined for
	// +kubebuilder:validation:Optional
	ActiveGateGroup *string `json:"activeGateGroup,omitempty" tf:"active_gate_group,omitempty"`

	// (String) The ID of the host this monitoring configuration will be defined for
	// The ID of the host this monitoring configuration will be defined for
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// (String) The ID of the host group this monitoring configuration will be defined for
	// The ID of the host group this monitoring configuration will be defined for
	// +kubebuilder:validation:Optional
	HostGroup *string `json:"hostGroup,omitempty" tf:"host_group,omitempty"`

	// (String) The name of the Management Zone this monitoring configuration will be defined for
	// The name of the Management Zone this monitoring configuration will be defined for
	// +kubebuilder:validation:Optional
	ManagementZone *string `json:"managementZone,omitempty" tf:"management_zone,omitempty"`

	// liberty-cp. You can query for these names using the data source dynatrace_hub_items
	// The fully qualified name of the extension, such as `com.dynatrace.extension.jmx-liberty-cp`. You can query for these names using the data source `dynatrace_hub_items`
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The JSON encoded value for this monitoring configuration
	// The JSON encoded value for this monitoring configuration
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// HubExtensionConfigSpec defines the desired state of HubExtensionConfig
type HubExtensionConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HubExtensionConfigParameters `json:"forProvider"`
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
	InitProvider HubExtensionConfigInitParameters `json:"initProvider,omitempty"`
}

// HubExtensionConfigStatus defines the observed state of HubExtensionConfig.
type HubExtensionConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HubExtensionConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HubExtensionConfig is the Schema for the HubExtensionConfigs API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type HubExtensionConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	Spec   HubExtensionConfigSpec   `json:"spec"`
	Status HubExtensionConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HubExtensionConfigList contains a list of HubExtensionConfigs
type HubExtensionConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HubExtensionConfig `json:"items"`
}

// Repository type metadata.
var (
	HubExtensionConfig_Kind             = "HubExtensionConfig"
	HubExtensionConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HubExtensionConfig_Kind}.String()
	HubExtensionConfig_KindAPIVersion   = HubExtensionConfig_Kind + "." + CRDGroupVersion.String()
	HubExtensionConfig_GroupVersionKind = CRDGroupVersion.WithKind(HubExtensionConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&HubExtensionConfig{}, &HubExtensionConfigList{})
}
