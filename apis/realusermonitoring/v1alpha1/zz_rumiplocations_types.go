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

type RumIPLocationsInitParameters struct {

	// (String) The city name of the location.
	// The city name of the location.
	City *string `json:"city,omitempty" tf:"city,omitempty"`

	// (String) The country code of the location.
	// The country code of the location.
	//
	// Use the alpha-2 code of the [ISO 3166-2 standard](https://dt-url.net/iso3166-2), (for example, `AT` for Austria or `PL` for Poland).
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// (String) Single IP or IP range start address
	// Single IP or IP range start address
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// (String) IP range end
	// IP range end
	IPTo *string `json:"ipTo,omitempty" tf:"ip_to,omitempty"`

	// (Number) Latitude
	// Latitude
	Latitude *float64 `json:"latitude,omitempty" tf:"latitude,omitempty"`

	// (Number) Longitude
	// Longitude
	Longitude *float64 `json:"longitude,omitempty" tf:"longitude,omitempty"`

	// (String) The region code of the location.
	// The region code of the location.
	//
	// For the [USA](https://dt-url.net/iso3166us) or [Canada](https://dt-url.net/iso3166ca) use ISO 3166-2 state codes without `US-` or `CA-` prefix.
	//
	// For the rest of the world use [FIPS 10-4 codes](https://dt-url.net/fipscodes) without country prefix.
	RegionCode *string `json:"regionCode,omitempty" tf:"region_code,omitempty"`
}

type RumIPLocationsObservation struct {

	// (String) The city name of the location.
	// The city name of the location.
	City *string `json:"city,omitempty" tf:"city,omitempty"`

	// (String) The country code of the location.
	// The country code of the location.
	//
	// Use the alpha-2 code of the [ISO 3166-2 standard](https://dt-url.net/iso3166-2), (for example, `AT` for Austria or `PL` for Poland).
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Single IP or IP range start address
	// Single IP or IP range start address
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// (String) IP range end
	// IP range end
	IPTo *string `json:"ipTo,omitempty" tf:"ip_to,omitempty"`

	// (Number) Latitude
	// Latitude
	Latitude *float64 `json:"latitude,omitempty" tf:"latitude,omitempty"`

	// (Number) Longitude
	// Longitude
	Longitude *float64 `json:"longitude,omitempty" tf:"longitude,omitempty"`

	// (String) The region code of the location.
	// The region code of the location.
	//
	// For the [USA](https://dt-url.net/iso3166us) or [Canada](https://dt-url.net/iso3166ca) use ISO 3166-2 state codes without `US-` or `CA-` prefix.
	//
	// For the rest of the world use [FIPS 10-4 codes](https://dt-url.net/fipscodes) without country prefix.
	RegionCode *string `json:"regionCode,omitempty" tf:"region_code,omitempty"`
}

type RumIPLocationsParameters struct {

	// (String) The city name of the location.
	// The city name of the location.
	// +kubebuilder:validation:Optional
	City *string `json:"city,omitempty" tf:"city,omitempty"`

	// (String) The country code of the location.
	// The country code of the location.
	//
	// Use the alpha-2 code of the [ISO 3166-2 standard](https://dt-url.net/iso3166-2), (for example, `AT` for Austria or `PL` for Poland).
	// +kubebuilder:validation:Optional
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// (String) Single IP or IP range start address
	// Single IP or IP range start address
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// (String) IP range end
	// IP range end
	// +kubebuilder:validation:Optional
	IPTo *string `json:"ipTo,omitempty" tf:"ip_to,omitempty"`

	// (Number) Latitude
	// Latitude
	// +kubebuilder:validation:Optional
	Latitude *float64 `json:"latitude,omitempty" tf:"latitude,omitempty"`

	// (Number) Longitude
	// Longitude
	// +kubebuilder:validation:Optional
	Longitude *float64 `json:"longitude,omitempty" tf:"longitude,omitempty"`

	// (String) The region code of the location.
	// The region code of the location.
	//
	// For the [USA](https://dt-url.net/iso3166us) or [Canada](https://dt-url.net/iso3166ca) use ISO 3166-2 state codes without `US-` or `CA-` prefix.
	//
	// For the rest of the world use [FIPS 10-4 codes](https://dt-url.net/fipscodes) without country prefix.
	// +kubebuilder:validation:Optional
	RegionCode *string `json:"regionCode,omitempty" tf:"region_code,omitempty"`
}

// RumIPLocationsSpec defines the desired state of RumIPLocations
type RumIPLocationsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RumIPLocationsParameters `json:"forProvider"`
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
	InitProvider RumIPLocationsInitParameters `json:"initProvider,omitempty"`
}

// RumIPLocationsStatus defines the observed state of RumIPLocations.
type RumIPLocationsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RumIPLocationsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RumIPLocations is the Schema for the RumIPLocationss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type RumIPLocations struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.countryCode) || (has(self.initProvider) && has(self.initProvider.countryCode))",message="spec.forProvider.countryCode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ip) || (has(self.initProvider) && has(self.initProvider.ip))",message="spec.forProvider.ip is a required parameter"
	Spec   RumIPLocationsSpec   `json:"spec"`
	Status RumIPLocationsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RumIPLocationsList contains a list of RumIPLocationss
type RumIPLocationsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RumIPLocations `json:"items"`
}

// Repository type metadata.
var (
	RumIPLocations_Kind             = "RumIPLocations"
	RumIPLocations_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RumIPLocations_Kind}.String()
	RumIPLocations_KindAPIVersion   = RumIPLocations_Kind + "." + CRDGroupVersion.String()
	RumIPLocations_GroupVersionKind = CRDGroupVersion.WithKind(RumIPLocations_Kind)
)

func init() {
	SchemeBuilder.Register(&RumIPLocations{}, &RumIPLocationsList{})
}
