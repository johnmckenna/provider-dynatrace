//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomUnits) DeepCopyInto(out *CustomUnits) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomUnits.
func (in *CustomUnits) DeepCopy() *CustomUnits {
	if in == nil {
		return nil
	}
	out := new(CustomUnits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomUnits) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomUnitsInitParameters) DeepCopyInto(out *CustomUnitsInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PluralName != nil {
		in, out := &in.PluralName, &out.PluralName
		*out = new(string)
		**out = **in
	}
	if in.Symbol != nil {
		in, out := &in.Symbol, &out.Symbol
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomUnitsInitParameters.
func (in *CustomUnitsInitParameters) DeepCopy() *CustomUnitsInitParameters {
	if in == nil {
		return nil
	}
	out := new(CustomUnitsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomUnitsList) DeepCopyInto(out *CustomUnitsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomUnits, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomUnitsList.
func (in *CustomUnitsList) DeepCopy() *CustomUnitsList {
	if in == nil {
		return nil
	}
	out := new(CustomUnitsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomUnitsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomUnitsObservation) DeepCopyInto(out *CustomUnitsObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PluralName != nil {
		in, out := &in.PluralName, &out.PluralName
		*out = new(string)
		**out = **in
	}
	if in.Symbol != nil {
		in, out := &in.Symbol, &out.Symbol
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomUnitsObservation.
func (in *CustomUnitsObservation) DeepCopy() *CustomUnitsObservation {
	if in == nil {
		return nil
	}
	out := new(CustomUnitsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomUnitsParameters) DeepCopyInto(out *CustomUnitsParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PluralName != nil {
		in, out := &in.PluralName, &out.PluralName
		*out = new(string)
		**out = **in
	}
	if in.Symbol != nil {
		in, out := &in.Symbol, &out.Symbol
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomUnitsParameters.
func (in *CustomUnitsParameters) DeepCopy() *CustomUnitsParameters {
	if in == nil {
		return nil
	}
	out := new(CustomUnitsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomUnitsSpec) DeepCopyInto(out *CustomUnitsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomUnitsSpec.
func (in *CustomUnitsSpec) DeepCopy() *CustomUnitsSpec {
	if in == nil {
		return nil
	}
	out := new(CustomUnitsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomUnitsStatus) DeepCopyInto(out *CustomUnitsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomUnitsStatus.
func (in *CustomUnitsStatus) DeepCopy() *CustomUnitsStatus {
	if in == nil {
		return nil
	}
	out := new(CustomUnitsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionInitParameters) DeepCopyInto(out *DimensionInitParameters) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionInitParameters.
func (in *DimensionInitParameters) DeepCopy() *DimensionInitParameters {
	if in == nil {
		return nil
	}
	out := new(DimensionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionObservation) DeepCopyInto(out *DimensionObservation) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionObservation.
func (in *DimensionObservation) DeepCopy() *DimensionObservation {
	if in == nil {
		return nil
	}
	out := new(DimensionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionParameters) DeepCopyInto(out *DimensionParameters) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionParameters.
func (in *DimensionParameters) DeepCopy() *DimensionParameters {
	if in == nil {
		return nil
	}
	out := new(DimensionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionsInitParameters) DeepCopyInto(out *DimensionsInitParameters) {
	*out = *in
	if in.Dimension != nil {
		in, out := &in.Dimension, &out.Dimension
		*out = make([]DimensionInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionsInitParameters.
func (in *DimensionsInitParameters) DeepCopy() *DimensionsInitParameters {
	if in == nil {
		return nil
	}
	out := new(DimensionsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionsObservation) DeepCopyInto(out *DimensionsObservation) {
	*out = *in
	if in.Dimension != nil {
		in, out := &in.Dimension, &out.Dimension
		*out = make([]DimensionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionsObservation.
func (in *DimensionsObservation) DeepCopy() *DimensionsObservation {
	if in == nil {
		return nil
	}
	out := new(DimensionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionsParameters) DeepCopyInto(out *DimensionsParameters) {
	*out = *in
	if in.Dimension != nil {
		in, out := &in.Dimension, &out.Dimension
		*out = make([]DimensionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionsParameters.
func (in *DimensionsParameters) DeepCopy() *DimensionsParameters {
	if in == nil {
		return nil
	}
	out := new(DimensionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistogramMetrics) DeepCopyInto(out *HistogramMetrics) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistogramMetrics.
func (in *HistogramMetrics) DeepCopy() *HistogramMetrics {
	if in == nil {
		return nil
	}
	out := new(HistogramMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HistogramMetrics) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistogramMetricsInitParameters) DeepCopyInto(out *HistogramMetricsInitParameters) {
	*out = *in
	if in.EnableHistogramBucketIngest != nil {
		in, out := &in.EnableHistogramBucketIngest, &out.EnableHistogramBucketIngest
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistogramMetricsInitParameters.
func (in *HistogramMetricsInitParameters) DeepCopy() *HistogramMetricsInitParameters {
	if in == nil {
		return nil
	}
	out := new(HistogramMetricsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistogramMetricsList) DeepCopyInto(out *HistogramMetricsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HistogramMetrics, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistogramMetricsList.
func (in *HistogramMetricsList) DeepCopy() *HistogramMetricsList {
	if in == nil {
		return nil
	}
	out := new(HistogramMetricsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HistogramMetricsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistogramMetricsObservation) DeepCopyInto(out *HistogramMetricsObservation) {
	*out = *in
	if in.EnableHistogramBucketIngest != nil {
		in, out := &in.EnableHistogramBucketIngest, &out.EnableHistogramBucketIngest
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistogramMetricsObservation.
func (in *HistogramMetricsObservation) DeepCopy() *HistogramMetricsObservation {
	if in == nil {
		return nil
	}
	out := new(HistogramMetricsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistogramMetricsParameters) DeepCopyInto(out *HistogramMetricsParameters) {
	*out = *in
	if in.EnableHistogramBucketIngest != nil {
		in, out := &in.EnableHistogramBucketIngest, &out.EnableHistogramBucketIngest
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistogramMetricsParameters.
func (in *HistogramMetricsParameters) DeepCopy() *HistogramMetricsParameters {
	if in == nil {
		return nil
	}
	out := new(HistogramMetricsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistogramMetricsSpec) DeepCopyInto(out *HistogramMetricsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistogramMetricsSpec.
func (in *HistogramMetricsSpec) DeepCopy() *HistogramMetricsSpec {
	if in == nil {
		return nil
	}
	out := new(HistogramMetricsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistogramMetricsStatus) DeepCopyInto(out *HistogramMetricsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistogramMetricsStatus.
func (in *HistogramMetricsStatus) DeepCopy() *HistogramMetricsStatus {
	if in == nil {
		return nil
	}
	out := new(HistogramMetricsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricMetadata) DeepCopyInto(out *MetricMetadata) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricMetadata.
func (in *MetricMetadata) DeepCopy() *MetricMetadata {
	if in == nil {
		return nil
	}
	out := new(MetricMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricMetadata) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricMetadataInitParameters) DeepCopyInto(out *MetricMetadataInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]DimensionsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.MetricID != nil {
		in, out := &in.MetricID, &out.MetricID
		*out = new(string)
		**out = **in
	}
	if in.MetricProperties != nil {
		in, out := &in.MetricProperties, &out.MetricProperties
		*out = make([]MetricPropertiesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SourceEntityType != nil {
		in, out := &in.SourceEntityType, &out.SourceEntityType
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
	if in.UnitDisplayFormat != nil {
		in, out := &in.UnitDisplayFormat, &out.UnitDisplayFormat
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricMetadataInitParameters.
func (in *MetricMetadataInitParameters) DeepCopy() *MetricMetadataInitParameters {
	if in == nil {
		return nil
	}
	out := new(MetricMetadataInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricMetadataList) DeepCopyInto(out *MetricMetadataList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetricMetadata, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricMetadataList.
func (in *MetricMetadataList) DeepCopy() *MetricMetadataList {
	if in == nil {
		return nil
	}
	out := new(MetricMetadataList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricMetadataList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricMetadataObservation) DeepCopyInto(out *MetricMetadataObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]DimensionsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MetricID != nil {
		in, out := &in.MetricID, &out.MetricID
		*out = new(string)
		**out = **in
	}
	if in.MetricProperties != nil {
		in, out := &in.MetricProperties, &out.MetricProperties
		*out = make([]MetricPropertiesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SourceEntityType != nil {
		in, out := &in.SourceEntityType, &out.SourceEntityType
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
	if in.UnitDisplayFormat != nil {
		in, out := &in.UnitDisplayFormat, &out.UnitDisplayFormat
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricMetadataObservation.
func (in *MetricMetadataObservation) DeepCopy() *MetricMetadataObservation {
	if in == nil {
		return nil
	}
	out := new(MetricMetadataObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricMetadataParameters) DeepCopyInto(out *MetricMetadataParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]DimensionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.MetricID != nil {
		in, out := &in.MetricID, &out.MetricID
		*out = new(string)
		**out = **in
	}
	if in.MetricProperties != nil {
		in, out := &in.MetricProperties, &out.MetricProperties
		*out = make([]MetricPropertiesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SourceEntityType != nil {
		in, out := &in.SourceEntityType, &out.SourceEntityType
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
	if in.UnitDisplayFormat != nil {
		in, out := &in.UnitDisplayFormat, &out.UnitDisplayFormat
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricMetadataParameters.
func (in *MetricMetadataParameters) DeepCopy() *MetricMetadataParameters {
	if in == nil {
		return nil
	}
	out := new(MetricMetadataParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricMetadataSpec) DeepCopyInto(out *MetricMetadataSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricMetadataSpec.
func (in *MetricMetadataSpec) DeepCopy() *MetricMetadataSpec {
	if in == nil {
		return nil
	}
	out := new(MetricMetadataSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricMetadataStatus) DeepCopyInto(out *MetricMetadataStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricMetadataStatus.
func (in *MetricMetadataStatus) DeepCopy() *MetricMetadataStatus {
	if in == nil {
		return nil
	}
	out := new(MetricMetadataStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricPropertiesInitParameters) DeepCopyInto(out *MetricPropertiesInitParameters) {
	*out = *in
	if in.ImpactRelevant != nil {
		in, out := &in.ImpactRelevant, &out.ImpactRelevant
		*out = new(bool)
		**out = **in
	}
	if in.Latency != nil {
		in, out := &in.Latency, &out.Latency
		*out = new(float64)
		**out = **in
	}
	if in.MaxValue != nil {
		in, out := &in.MaxValue, &out.MaxValue
		*out = new(float64)
		**out = **in
	}
	if in.MinValue != nil {
		in, out := &in.MinValue, &out.MinValue
		*out = new(float64)
		**out = **in
	}
	if in.RootCauseRelevant != nil {
		in, out := &in.RootCauseRelevant, &out.RootCauseRelevant
		*out = new(bool)
		**out = **in
	}
	if in.ValueType != nil {
		in, out := &in.ValueType, &out.ValueType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricPropertiesInitParameters.
func (in *MetricPropertiesInitParameters) DeepCopy() *MetricPropertiesInitParameters {
	if in == nil {
		return nil
	}
	out := new(MetricPropertiesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricPropertiesObservation) DeepCopyInto(out *MetricPropertiesObservation) {
	*out = *in
	if in.ImpactRelevant != nil {
		in, out := &in.ImpactRelevant, &out.ImpactRelevant
		*out = new(bool)
		**out = **in
	}
	if in.Latency != nil {
		in, out := &in.Latency, &out.Latency
		*out = new(float64)
		**out = **in
	}
	if in.MaxValue != nil {
		in, out := &in.MaxValue, &out.MaxValue
		*out = new(float64)
		**out = **in
	}
	if in.MinValue != nil {
		in, out := &in.MinValue, &out.MinValue
		*out = new(float64)
		**out = **in
	}
	if in.RootCauseRelevant != nil {
		in, out := &in.RootCauseRelevant, &out.RootCauseRelevant
		*out = new(bool)
		**out = **in
	}
	if in.ValueType != nil {
		in, out := &in.ValueType, &out.ValueType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricPropertiesObservation.
func (in *MetricPropertiesObservation) DeepCopy() *MetricPropertiesObservation {
	if in == nil {
		return nil
	}
	out := new(MetricPropertiesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricPropertiesParameters) DeepCopyInto(out *MetricPropertiesParameters) {
	*out = *in
	if in.ImpactRelevant != nil {
		in, out := &in.ImpactRelevant, &out.ImpactRelevant
		*out = new(bool)
		**out = **in
	}
	if in.Latency != nil {
		in, out := &in.Latency, &out.Latency
		*out = new(float64)
		**out = **in
	}
	if in.MaxValue != nil {
		in, out := &in.MaxValue, &out.MaxValue
		*out = new(float64)
		**out = **in
	}
	if in.MinValue != nil {
		in, out := &in.MinValue, &out.MinValue
		*out = new(float64)
		**out = **in
	}
	if in.RootCauseRelevant != nil {
		in, out := &in.RootCauseRelevant, &out.RootCauseRelevant
		*out = new(bool)
		**out = **in
	}
	if in.ValueType != nil {
		in, out := &in.ValueType, &out.ValueType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricPropertiesParameters.
func (in *MetricPropertiesParameters) DeepCopy() *MetricPropertiesParameters {
	if in == nil {
		return nil
	}
	out := new(MetricPropertiesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricQuery) DeepCopyInto(out *MetricQuery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricQuery.
func (in *MetricQuery) DeepCopy() *MetricQuery {
	if in == nil {
		return nil
	}
	out := new(MetricQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricQuery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricQueryInitParameters) DeepCopyInto(out *MetricQueryInitParameters) {
	*out = *in
	if in.MetricID != nil {
		in, out := &in.MetricID, &out.MetricID
		*out = new(string)
		**out = **in
	}
	if in.MetricSelector != nil {
		in, out := &in.MetricSelector, &out.MetricSelector
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricQueryInitParameters.
func (in *MetricQueryInitParameters) DeepCopy() *MetricQueryInitParameters {
	if in == nil {
		return nil
	}
	out := new(MetricQueryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricQueryList) DeepCopyInto(out *MetricQueryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetricQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricQueryList.
func (in *MetricQueryList) DeepCopy() *MetricQueryList {
	if in == nil {
		return nil
	}
	out := new(MetricQueryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricQueryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricQueryObservation) DeepCopyInto(out *MetricQueryObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MetricID != nil {
		in, out := &in.MetricID, &out.MetricID
		*out = new(string)
		**out = **in
	}
	if in.MetricSelector != nil {
		in, out := &in.MetricSelector, &out.MetricSelector
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricQueryObservation.
func (in *MetricQueryObservation) DeepCopy() *MetricQueryObservation {
	if in == nil {
		return nil
	}
	out := new(MetricQueryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricQueryParameters) DeepCopyInto(out *MetricQueryParameters) {
	*out = *in
	if in.MetricID != nil {
		in, out := &in.MetricID, &out.MetricID
		*out = new(string)
		**out = **in
	}
	if in.MetricSelector != nil {
		in, out := &in.MetricSelector, &out.MetricSelector
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricQueryParameters.
func (in *MetricQueryParameters) DeepCopy() *MetricQueryParameters {
	if in == nil {
		return nil
	}
	out := new(MetricQueryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricQuerySpec) DeepCopyInto(out *MetricQuerySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricQuerySpec.
func (in *MetricQuerySpec) DeepCopy() *MetricQuerySpec {
	if in == nil {
		return nil
	}
	out := new(MetricQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricQueryStatus) DeepCopyInto(out *MetricQueryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricQueryStatus.
func (in *MetricQueryStatus) DeepCopy() *MetricQueryStatus {
	if in == nil {
		return nil
	}
	out := new(MetricQueryStatus)
	in.DeepCopyInto(out)
	return out
}
