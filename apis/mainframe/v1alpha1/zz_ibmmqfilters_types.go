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

type IbmMqFiltersInitParameters struct {

	// (Set of String) CICS: Excluded MQ queues
	// CICS: Excluded MQ queues
	// +listType=set
	CicsMqQueueIDExcludes []*string `json:"cicsMqQueueIdExcludes,omitempty" tf:"cics_mq_queue_id_excludes,omitempty"`

	// (Set of String) CICS: Included MQ queues
	// CICS: Included MQ queues
	// +listType=set
	CicsMqQueueIDIncludes []*string `json:"cicsMqQueueIdIncludes,omitempty" tf:"cics_mq_queue_id_includes,omitempty"`

	// (Set of String) When you add a transaction ID to the exclude list remaining transactions are still monitored.
	// When you add a transaction ID to the exclude list remaining transactions are still monitored.
	// +listType=set
	ImsCrTrnIDExcludes []*string `json:"imsCrTrnIdExcludes,omitempty" tf:"ims_cr_trn_id_excludes,omitempty"`

	// (Set of String) When you add a transaction ID to the include list, all the remaining transactions are ignored.
	// When you add a transaction ID to the include list, all the remaining transactions are ignored.
	// +listType=set
	ImsCrTrnIDIncludes []*string `json:"imsCrTrnIdIncludes,omitempty" tf:"ims_cr_trn_id_includes,omitempty"`

	// (Set of String) IMS: Excluded MQ queues
	// IMS: Excluded MQ queues
	// +listType=set
	ImsMqQueueIDExcludes []*string `json:"imsMqQueueIdExcludes,omitempty" tf:"ims_mq_queue_id_excludes,omitempty"`

	// (Set of String) IMS: Included MQ queues
	// IMS: Included MQ queues
	// +listType=set
	ImsMqQueueIDIncludes []*string `json:"imsMqQueueIdIncludes,omitempty" tf:"ims_mq_queue_id_includes,omitempty"`
}

type IbmMqFiltersObservation struct {

	// (Set of String) CICS: Excluded MQ queues
	// CICS: Excluded MQ queues
	// +listType=set
	CicsMqQueueIDExcludes []*string `json:"cicsMqQueueIdExcludes,omitempty" tf:"cics_mq_queue_id_excludes,omitempty"`

	// (Set of String) CICS: Included MQ queues
	// CICS: Included MQ queues
	// +listType=set
	CicsMqQueueIDIncludes []*string `json:"cicsMqQueueIdIncludes,omitempty" tf:"cics_mq_queue_id_includes,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of String) When you add a transaction ID to the exclude list remaining transactions are still monitored.
	// When you add a transaction ID to the exclude list remaining transactions are still monitored.
	// +listType=set
	ImsCrTrnIDExcludes []*string `json:"imsCrTrnIdExcludes,omitempty" tf:"ims_cr_trn_id_excludes,omitempty"`

	// (Set of String) When you add a transaction ID to the include list, all the remaining transactions are ignored.
	// When you add a transaction ID to the include list, all the remaining transactions are ignored.
	// +listType=set
	ImsCrTrnIDIncludes []*string `json:"imsCrTrnIdIncludes,omitempty" tf:"ims_cr_trn_id_includes,omitempty"`

	// (Set of String) IMS: Excluded MQ queues
	// IMS: Excluded MQ queues
	// +listType=set
	ImsMqQueueIDExcludes []*string `json:"imsMqQueueIdExcludes,omitempty" tf:"ims_mq_queue_id_excludes,omitempty"`

	// (Set of String) IMS: Included MQ queues
	// IMS: Included MQ queues
	// +listType=set
	ImsMqQueueIDIncludes []*string `json:"imsMqQueueIdIncludes,omitempty" tf:"ims_mq_queue_id_includes,omitempty"`
}

type IbmMqFiltersParameters struct {

	// (Set of String) CICS: Excluded MQ queues
	// CICS: Excluded MQ queues
	// +kubebuilder:validation:Optional
	// +listType=set
	CicsMqQueueIDExcludes []*string `json:"cicsMqQueueIdExcludes,omitempty" tf:"cics_mq_queue_id_excludes,omitempty"`

	// (Set of String) CICS: Included MQ queues
	// CICS: Included MQ queues
	// +kubebuilder:validation:Optional
	// +listType=set
	CicsMqQueueIDIncludes []*string `json:"cicsMqQueueIdIncludes,omitempty" tf:"cics_mq_queue_id_includes,omitempty"`

	// (Set of String) When you add a transaction ID to the exclude list remaining transactions are still monitored.
	// When you add a transaction ID to the exclude list remaining transactions are still monitored.
	// +kubebuilder:validation:Optional
	// +listType=set
	ImsCrTrnIDExcludes []*string `json:"imsCrTrnIdExcludes,omitempty" tf:"ims_cr_trn_id_excludes,omitempty"`

	// (Set of String) When you add a transaction ID to the include list, all the remaining transactions are ignored.
	// When you add a transaction ID to the include list, all the remaining transactions are ignored.
	// +kubebuilder:validation:Optional
	// +listType=set
	ImsCrTrnIDIncludes []*string `json:"imsCrTrnIdIncludes,omitempty" tf:"ims_cr_trn_id_includes,omitempty"`

	// (Set of String) IMS: Excluded MQ queues
	// IMS: Excluded MQ queues
	// +kubebuilder:validation:Optional
	// +listType=set
	ImsMqQueueIDExcludes []*string `json:"imsMqQueueIdExcludes,omitempty" tf:"ims_mq_queue_id_excludes,omitempty"`

	// (Set of String) IMS: Included MQ queues
	// IMS: Included MQ queues
	// +kubebuilder:validation:Optional
	// +listType=set
	ImsMqQueueIDIncludes []*string `json:"imsMqQueueIdIncludes,omitempty" tf:"ims_mq_queue_id_includes,omitempty"`
}

// IbmMqFiltersSpec defines the desired state of IbmMqFilters
type IbmMqFiltersSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IbmMqFiltersParameters `json:"forProvider"`
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
	InitProvider IbmMqFiltersInitParameters `json:"initProvider,omitempty"`
}

// IbmMqFiltersStatus defines the observed state of IbmMqFilters.
type IbmMqFiltersStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IbmMqFiltersObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IbmMqFilters is the Schema for the IbmMqFilterss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type IbmMqFilters struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IbmMqFiltersSpec   `json:"spec"`
	Status            IbmMqFiltersStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IbmMqFiltersList contains a list of IbmMqFilterss
type IbmMqFiltersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IbmMqFilters `json:"items"`
}

// Repository type metadata.
var (
	IbmMqFilters_Kind             = "IbmMqFilters"
	IbmMqFilters_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IbmMqFilters_Kind}.String()
	IbmMqFilters_KindAPIVersion   = IbmMqFilters_Kind + "." + CRDGroupVersion.String()
	IbmMqFilters_GroupVersionKind = CRDGroupVersion.WithKind(IbmMqFilters_Kind)
)

func init() {
	SchemeBuilder.Register(&IbmMqFilters{}, &IbmMqFiltersList{})
}