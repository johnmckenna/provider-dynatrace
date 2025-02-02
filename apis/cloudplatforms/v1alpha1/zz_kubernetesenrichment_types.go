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

type KubernetesEnrichmentInitParameters struct {

	// (Block List, Max: 1) Dynatrace allows to use metadata defined on Kubernetes nodes, namespaces, and pods to set security and cost allocation attributes and dimensions for metrics, events, log, spans, and entities associated with the respective Kubernetes resource.
	// Dynatrace allows to use metadata defined on Kubernetes nodes, namespaces, and pods to set security and cost allocation attributes and dimensions for metrics, events, log, spans, and entities associated with the respective Kubernetes resource.
	//
	// The following annotation keys are considered:
	// * `metadata.dynatrace.com/dt.security_context`
	// * `metadata.dynatrace.com/dt.cost.product`
	// * `metadata.dynatrace.com/dt.cost.costcenter`
	//
	// Pod annotations determine the attributes of data associated with the pod itself, and containers belonging to the pod.
	//
	// Namespace annotations determine the attributes of data associated with the namespace itself, workloads, services, and - if not overwritten on pod level - pods, and containers belonging to the namespace.
	//
	// Node annotations determine the attributes of data associated with only the node.
	//
	// Depending on your specific use case and environment, you have the following enrichment options:
	//
	// **Manual annotation:**
	//
	// Use the aforementioned annotation keys when annotating your namespaces and pods to enrich your Kubernetes data with security and cost allocation attributes.
	//
	// With Dynatrace Operator version 1.3.0, the aforementioned namespace annotations are copied down to pods in the namespace, if they are not yet set on the respective pod.
	//
	// **Rule-based annotation:**
	//
	// If you already have labels or annotations defined on your namespaces, and you want to reuse them for enrichment, you can do so with the help of rules definable here.
	//
	// **Example:**
	//
	// * Namespace label:
	// * `label/example: test-value`
	//
	// * Rule:
	// * `Label`
	// `label/example --> dt.security_context`
	//
	// * Pod annotation:
	// * `metadata.dynatrace.com/dt.security_context: test-value`
	//
	// A maximum of 5 rules can be defined. The first applicable rule will be applied. Preexisting annotations will not be overwritten. For a detailed description of this feature, have a look at our [documentation](https://dt-url.net/pn22sye).
	Rules []RulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// (String) The scope of this setting (KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	// The scope of this setting (KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type KubernetesEnrichmentObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Max: 1) Dynatrace allows to use metadata defined on Kubernetes nodes, namespaces, and pods to set security and cost allocation attributes and dimensions for metrics, events, log, spans, and entities associated with the respective Kubernetes resource.
	// Dynatrace allows to use metadata defined on Kubernetes nodes, namespaces, and pods to set security and cost allocation attributes and dimensions for metrics, events, log, spans, and entities associated with the respective Kubernetes resource.
	//
	// The following annotation keys are considered:
	// * `metadata.dynatrace.com/dt.security_context`
	// * `metadata.dynatrace.com/dt.cost.product`
	// * `metadata.dynatrace.com/dt.cost.costcenter`
	//
	// Pod annotations determine the attributes of data associated with the pod itself, and containers belonging to the pod.
	//
	// Namespace annotations determine the attributes of data associated with the namespace itself, workloads, services, and - if not overwritten on pod level - pods, and containers belonging to the namespace.
	//
	// Node annotations determine the attributes of data associated with only the node.
	//
	// Depending on your specific use case and environment, you have the following enrichment options:
	//
	// **Manual annotation:**
	//
	// Use the aforementioned annotation keys when annotating your namespaces and pods to enrich your Kubernetes data with security and cost allocation attributes.
	//
	// With Dynatrace Operator version 1.3.0, the aforementioned namespace annotations are copied down to pods in the namespace, if they are not yet set on the respective pod.
	//
	// **Rule-based annotation:**
	//
	// If you already have labels or annotations defined on your namespaces, and you want to reuse them for enrichment, you can do so with the help of rules definable here.
	//
	// **Example:**
	//
	// * Namespace label:
	// * `label/example: test-value`
	//
	// * Rule:
	// * `Label`
	// `label/example --> dt.security_context`
	//
	// * Pod annotation:
	// * `metadata.dynatrace.com/dt.security_context: test-value`
	//
	// A maximum of 5 rules can be defined. The first applicable rule will be applied. Preexisting annotations will not be overwritten. For a detailed description of this feature, have a look at our [documentation](https://dt-url.net/pn22sye).
	Rules []RulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`

	// (String) The scope of this setting (KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	// The scope of this setting (KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type KubernetesEnrichmentParameters struct {

	// (Block List, Max: 1) Dynatrace allows to use metadata defined on Kubernetes nodes, namespaces, and pods to set security and cost allocation attributes and dimensions for metrics, events, log, spans, and entities associated with the respective Kubernetes resource.
	// Dynatrace allows to use metadata defined on Kubernetes nodes, namespaces, and pods to set security and cost allocation attributes and dimensions for metrics, events, log, spans, and entities associated with the respective Kubernetes resource.
	//
	// The following annotation keys are considered:
	// * `metadata.dynatrace.com/dt.security_context`
	// * `metadata.dynatrace.com/dt.cost.product`
	// * `metadata.dynatrace.com/dt.cost.costcenter`
	//
	// Pod annotations determine the attributes of data associated with the pod itself, and containers belonging to the pod.
	//
	// Namespace annotations determine the attributes of data associated with the namespace itself, workloads, services, and - if not overwritten on pod level - pods, and containers belonging to the namespace.
	//
	// Node annotations determine the attributes of data associated with only the node.
	//
	// Depending on your specific use case and environment, you have the following enrichment options:
	//
	// **Manual annotation:**
	//
	// Use the aforementioned annotation keys when annotating your namespaces and pods to enrich your Kubernetes data with security and cost allocation attributes.
	//
	// With Dynatrace Operator version 1.3.0, the aforementioned namespace annotations are copied down to pods in the namespace, if they are not yet set on the respective pod.
	//
	// **Rule-based annotation:**
	//
	// If you already have labels or annotations defined on your namespaces, and you want to reuse them for enrichment, you can do so with the help of rules definable here.
	//
	// **Example:**
	//
	// * Namespace label:
	// * `label/example: test-value`
	//
	// * Rule:
	// * `Label`
	// `label/example --> dt.security_context`
	//
	// * Pod annotation:
	// * `metadata.dynatrace.com/dt.security_context: test-value`
	//
	// A maximum of 5 rules can be defined. The first applicable rule will be applied. Preexisting annotations will not be overwritten. For a detailed description of this feature, have a look at our [documentation](https://dt-url.net/pn22sye).
	// +kubebuilder:validation:Optional
	Rules []RulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// (String) The scope of this setting (KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	// The scope of this setting (KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type RuleInitParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The source must follow the syntax of Kubernetes annotation/label keys as defined in the Kubernetes documentation.
	// The source must follow the syntax of Kubernetes annotation/label keys as defined in the [Kubernetes documentation](https://dt-url.net/2c02sbn).
	//
	// `source := (prefix/)?name`
	//
	// `prefix := [a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*`
	//
	// `name := ([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]`
	//
	// Additionally, the name can have at most 63 characters, and the overall length of the source must not exceed 75 characters.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Dt_cost_costcenter, Dt_cost_product, Dt_security_context
	// Possible Values: `Dt_cost_costcenter`, `Dt_cost_product`, `Dt_security_context`
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// (String) Possible Values: ANNOTATION, LABEL
	// Possible Values: `ANNOTATION`, `LABEL`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RuleObservation struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The source must follow the syntax of Kubernetes annotation/label keys as defined in the Kubernetes documentation.
	// The source must follow the syntax of Kubernetes annotation/label keys as defined in the [Kubernetes documentation](https://dt-url.net/2c02sbn).
	//
	// `source := (prefix/)?name`
	//
	// `prefix := [a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*`
	//
	// `name := ([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]`
	//
	// Additionally, the name can have at most 63 characters, and the overall length of the source must not exceed 75 characters.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String) Possible Values: Dt_cost_costcenter, Dt_cost_product, Dt_security_context
	// Possible Values: `Dt_cost_costcenter`, `Dt_cost_product`, `Dt_security_context`
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// (String) Possible Values: ANNOTATION, LABEL
	// Possible Values: `ANNOTATION`, `LABEL`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RuleParameters struct {

	// (Boolean) This setting is enabled (true) or disabled (false)
	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// (String) The source must follow the syntax of Kubernetes annotation/label keys as defined in the Kubernetes documentation.
	// The source must follow the syntax of Kubernetes annotation/label keys as defined in the [Kubernetes documentation](https://dt-url.net/2c02sbn).
	//
	// `source := (prefix/)?name`
	//
	// `prefix := [a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*`
	//
	// `name := ([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]`
	//
	// Additionally, the name can have at most 63 characters, and the overall length of the source must not exceed 75 characters.
	// +kubebuilder:validation:Optional
	Source *string `json:"source" tf:"source,omitempty"`

	// (String) Possible Values: Dt_cost_costcenter, Dt_cost_product, Dt_security_context
	// Possible Values: `Dt_cost_costcenter`, `Dt_cost_product`, `Dt_security_context`
	// +kubebuilder:validation:Optional
	Target *string `json:"target" tf:"target,omitempty"`

	// (String) Possible Values: ANNOTATION, LABEL
	// Possible Values: `ANNOTATION`, `LABEL`
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type RulesInitParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type RulesObservation struct {

	// (Block List, Min: 1) (see below for nested schema)
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type RulesParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule" tf:"rule,omitempty"`
}

// KubernetesEnrichmentSpec defines the desired state of KubernetesEnrichment
type KubernetesEnrichmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KubernetesEnrichmentParameters `json:"forProvider"`
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
	InitProvider KubernetesEnrichmentInitParameters `json:"initProvider,omitempty"`
}

// KubernetesEnrichmentStatus defines the observed state of KubernetesEnrichment.
type KubernetesEnrichmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KubernetesEnrichmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// KubernetesEnrichment is the Schema for the KubernetesEnrichments API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type KubernetesEnrichment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesEnrichmentSpec   `json:"spec"`
	Status            KubernetesEnrichmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubernetesEnrichmentList contains a list of KubernetesEnrichments
type KubernetesEnrichmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesEnrichment `json:"items"`
}

// Repository type metadata.
var (
	KubernetesEnrichment_Kind             = "KubernetesEnrichment"
	KubernetesEnrichment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KubernetesEnrichment_Kind}.String()
	KubernetesEnrichment_KindAPIVersion   = KubernetesEnrichment_Kind + "." + CRDGroupVersion.String()
	KubernetesEnrichment_GroupVersionKind = CRDGroupVersion.WithKind(KubernetesEnrichment_Kind)
)

func init() {
	SchemeBuilder.Register(&KubernetesEnrichment{}, &KubernetesEnrichmentList{})
}
