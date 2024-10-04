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

type SessionReplayResourceCaptureInitParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (Boolean) (Field has overlap with dynatrace_web_application) When turned on, all CSS resources from all sessions are captured. For details, see Resource capture.
	// (Field has overlap with `dynatrace_web_application`) When turned on, all CSS resources from all sessions are captured. For details, see [Resource capture](https://dt-url.net/sr-resource-capturing).
	EnableResourceCapturing *bool `json:"enableResourceCapturing,omitempty" tf:"enable_resource_capturing,omitempty"`

	// (Set of String) (Field has overlap with dynatrace_web_application) Add exclusion rules to avoid the capture of resources from certain pages.
	// (Field has overlap with `dynatrace_web_application`) Add exclusion rules to avoid the capture of resources from certain pages.
	// +listType=set
	ResourceCaptureURLExclusionPatternList []*string `json:"resourceCaptureUrlExclusionPatternList,omitempty" tf:"resource_capture_url_exclusion_pattern_list,omitempty"`
}

type SessionReplayResourceCaptureObservation struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (Boolean) (Field has overlap with dynatrace_web_application) When turned on, all CSS resources from all sessions are captured. For details, see Resource capture.
	// (Field has overlap with `dynatrace_web_application`) When turned on, all CSS resources from all sessions are captured. For details, see [Resource capture](https://dt-url.net/sr-resource-capturing).
	EnableResourceCapturing *bool `json:"enableResourceCapturing,omitempty" tf:"enable_resource_capturing,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of String) (Field has overlap with dynatrace_web_application) Add exclusion rules to avoid the capture of resources from certain pages.
	// (Field has overlap with `dynatrace_web_application`) Add exclusion rules to avoid the capture of resources from certain pages.
	// +listType=set
	ResourceCaptureURLExclusionPatternList []*string `json:"resourceCaptureUrlExclusionPatternList,omitempty" tf:"resource_capture_url_exclusion_pattern_list,omitempty"`
}

type SessionReplayResourceCaptureParameters struct {

	// (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (Boolean) (Field has overlap with dynatrace_web_application) When turned on, all CSS resources from all sessions are captured. For details, see Resource capture.
	// (Field has overlap with `dynatrace_web_application`) When turned on, all CSS resources from all sessions are captured. For details, see [Resource capture](https://dt-url.net/sr-resource-capturing).
	// +kubebuilder:validation:Optional
	EnableResourceCapturing *bool `json:"enableResourceCapturing,omitempty" tf:"enable_resource_capturing,omitempty"`

	// (Set of String) (Field has overlap with dynatrace_web_application) Add exclusion rules to avoid the capture of resources from certain pages.
	// (Field has overlap with `dynatrace_web_application`) Add exclusion rules to avoid the capture of resources from certain pages.
	// +kubebuilder:validation:Optional
	// +listType=set
	ResourceCaptureURLExclusionPatternList []*string `json:"resourceCaptureUrlExclusionPatternList,omitempty" tf:"resource_capture_url_exclusion_pattern_list,omitempty"`
}

// SessionReplayResourceCaptureSpec defines the desired state of SessionReplayResourceCapture
type SessionReplayResourceCaptureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SessionReplayResourceCaptureParameters `json:"forProvider"`
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
	InitProvider SessionReplayResourceCaptureInitParameters `json:"initProvider,omitempty"`
}

// SessionReplayResourceCaptureStatus defines the observed state of SessionReplayResourceCapture.
type SessionReplayResourceCaptureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SessionReplayResourceCaptureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SessionReplayResourceCapture is the Schema for the SessionReplayResourceCaptures API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type SessionReplayResourceCapture struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enableResourceCapturing) || (has(self.initProvider) && has(self.initProvider.enableResourceCapturing))",message="spec.forProvider.enableResourceCapturing is a required parameter"
	Spec   SessionReplayResourceCaptureSpec   `json:"spec"`
	Status SessionReplayResourceCaptureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SessionReplayResourceCaptureList contains a list of SessionReplayResourceCaptures
type SessionReplayResourceCaptureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SessionReplayResourceCapture `json:"items"`
}

// Repository type metadata.
var (
	SessionReplayResourceCapture_Kind             = "SessionReplayResourceCapture"
	SessionReplayResourceCapture_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SessionReplayResourceCapture_Kind}.String()
	SessionReplayResourceCapture_KindAPIVersion   = SessionReplayResourceCapture_Kind + "." + CRDGroupVersion.String()
	SessionReplayResourceCapture_GroupVersionKind = CRDGroupVersion.WithKind(SessionReplayResourceCapture_Kind)
)

func init() {
	SchemeBuilder.Register(&SessionReplayResourceCapture{}, &SessionReplayResourceCaptureList{})
}
