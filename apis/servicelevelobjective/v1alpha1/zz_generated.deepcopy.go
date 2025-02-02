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
func (in *ErrorBudgetBurnRateInitParameters) DeepCopyInto(out *ErrorBudgetBurnRateInitParameters) {
	*out = *in
	if in.BurnRateVisualizationEnabled != nil {
		in, out := &in.BurnRateVisualizationEnabled, &out.BurnRateVisualizationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.FastBurnThreshold != nil {
		in, out := &in.FastBurnThreshold, &out.FastBurnThreshold
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorBudgetBurnRateInitParameters.
func (in *ErrorBudgetBurnRateInitParameters) DeepCopy() *ErrorBudgetBurnRateInitParameters {
	if in == nil {
		return nil
	}
	out := new(ErrorBudgetBurnRateInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorBudgetBurnRateObservation) DeepCopyInto(out *ErrorBudgetBurnRateObservation) {
	*out = *in
	if in.BurnRateVisualizationEnabled != nil {
		in, out := &in.BurnRateVisualizationEnabled, &out.BurnRateVisualizationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.FastBurnThreshold != nil {
		in, out := &in.FastBurnThreshold, &out.FastBurnThreshold
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorBudgetBurnRateObservation.
func (in *ErrorBudgetBurnRateObservation) DeepCopy() *ErrorBudgetBurnRateObservation {
	if in == nil {
		return nil
	}
	out := new(ErrorBudgetBurnRateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorBudgetBurnRateParameters) DeepCopyInto(out *ErrorBudgetBurnRateParameters) {
	*out = *in
	if in.BurnRateVisualizationEnabled != nil {
		in, out := &in.BurnRateVisualizationEnabled, &out.BurnRateVisualizationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.FastBurnThreshold != nil {
		in, out := &in.FastBurnThreshold, &out.FastBurnThreshold
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorBudgetBurnRateParameters.
func (in *ErrorBudgetBurnRateParameters) DeepCopy() *ErrorBudgetBurnRateParameters {
	if in == nil {
		return nil
	}
	out := new(ErrorBudgetBurnRateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLONormalization) DeepCopyInto(out *SLONormalization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLONormalization.
func (in *SLONormalization) DeepCopy() *SLONormalization {
	if in == nil {
		return nil
	}
	out := new(SLONormalization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SLONormalization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLONormalizationInitParameters) DeepCopyInto(out *SLONormalizationInitParameters) {
	*out = *in
	if in.Normalize != nil {
		in, out := &in.Normalize, &out.Normalize
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLONormalizationInitParameters.
func (in *SLONormalizationInitParameters) DeepCopy() *SLONormalizationInitParameters {
	if in == nil {
		return nil
	}
	out := new(SLONormalizationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLONormalizationList) DeepCopyInto(out *SLONormalizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SLONormalization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLONormalizationList.
func (in *SLONormalizationList) DeepCopy() *SLONormalizationList {
	if in == nil {
		return nil
	}
	out := new(SLONormalizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SLONormalizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLONormalizationObservation) DeepCopyInto(out *SLONormalizationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Normalize != nil {
		in, out := &in.Normalize, &out.Normalize
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLONormalizationObservation.
func (in *SLONormalizationObservation) DeepCopy() *SLONormalizationObservation {
	if in == nil {
		return nil
	}
	out := new(SLONormalizationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLONormalizationParameters) DeepCopyInto(out *SLONormalizationParameters) {
	*out = *in
	if in.Normalize != nil {
		in, out := &in.Normalize, &out.Normalize
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLONormalizationParameters.
func (in *SLONormalizationParameters) DeepCopy() *SLONormalizationParameters {
	if in == nil {
		return nil
	}
	out := new(SLONormalizationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLONormalizationSpec) DeepCopyInto(out *SLONormalizationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLONormalizationSpec.
func (in *SLONormalizationSpec) DeepCopy() *SLONormalizationSpec {
	if in == nil {
		return nil
	}
	out := new(SLONormalizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLONormalizationStatus) DeepCopyInto(out *SLONormalizationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLONormalizationStatus.
func (in *SLONormalizationStatus) DeepCopy() *SLONormalizationStatus {
	if in == nil {
		return nil
	}
	out := new(SLONormalizationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLOV2) DeepCopyInto(out *SLOV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLOV2.
func (in *SLOV2) DeepCopy() *SLOV2 {
	if in == nil {
		return nil
	}
	out := new(SLOV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SLOV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLOV2InitParameters) DeepCopyInto(out *SLOV2InitParameters) {
	*out = *in
	if in.CustomDescription != nil {
		in, out := &in.CustomDescription, &out.CustomDescription
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ErrorBudgetBurnRate != nil {
		in, out := &in.ErrorBudgetBurnRate, &out.ErrorBudgetBurnRate
		*out = make([]ErrorBudgetBurnRateInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EvaluationType != nil {
		in, out := &in.EvaluationType, &out.EvaluationType
		*out = new(string)
		**out = **in
	}
	if in.EvaluationWindow != nil {
		in, out := &in.EvaluationWindow, &out.EvaluationWindow
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.LegacyID != nil {
		in, out := &in.LegacyID, &out.LegacyID
		*out = new(string)
		**out = **in
	}
	if in.MetricExpression != nil {
		in, out := &in.MetricExpression, &out.MetricExpression
		*out = new(string)
		**out = **in
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TargetSuccess != nil {
		in, out := &in.TargetSuccess, &out.TargetSuccess
		*out = new(float64)
		**out = **in
	}
	if in.TargetWarning != nil {
		in, out := &in.TargetWarning, &out.TargetWarning
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLOV2InitParameters.
func (in *SLOV2InitParameters) DeepCopy() *SLOV2InitParameters {
	if in == nil {
		return nil
	}
	out := new(SLOV2InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLOV2List) DeepCopyInto(out *SLOV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SLOV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLOV2List.
func (in *SLOV2List) DeepCopy() *SLOV2List {
	if in == nil {
		return nil
	}
	out := new(SLOV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SLOV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLOV2Observation) DeepCopyInto(out *SLOV2Observation) {
	*out = *in
	if in.CustomDescription != nil {
		in, out := &in.CustomDescription, &out.CustomDescription
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ErrorBudgetBurnRate != nil {
		in, out := &in.ErrorBudgetBurnRate, &out.ErrorBudgetBurnRate
		*out = make([]ErrorBudgetBurnRateObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EvaluationType != nil {
		in, out := &in.EvaluationType, &out.EvaluationType
		*out = new(string)
		**out = **in
	}
	if in.EvaluationWindow != nil {
		in, out := &in.EvaluationWindow, &out.EvaluationWindow
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LegacyID != nil {
		in, out := &in.LegacyID, &out.LegacyID
		*out = new(string)
		**out = **in
	}
	if in.MetricExpression != nil {
		in, out := &in.MetricExpression, &out.MetricExpression
		*out = new(string)
		**out = **in
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TargetSuccess != nil {
		in, out := &in.TargetSuccess, &out.TargetSuccess
		*out = new(float64)
		**out = **in
	}
	if in.TargetWarning != nil {
		in, out := &in.TargetWarning, &out.TargetWarning
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLOV2Observation.
func (in *SLOV2Observation) DeepCopy() *SLOV2Observation {
	if in == nil {
		return nil
	}
	out := new(SLOV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLOV2Parameters) DeepCopyInto(out *SLOV2Parameters) {
	*out = *in
	if in.CustomDescription != nil {
		in, out := &in.CustomDescription, &out.CustomDescription
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ErrorBudgetBurnRate != nil {
		in, out := &in.ErrorBudgetBurnRate, &out.ErrorBudgetBurnRate
		*out = make([]ErrorBudgetBurnRateParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EvaluationType != nil {
		in, out := &in.EvaluationType, &out.EvaluationType
		*out = new(string)
		**out = **in
	}
	if in.EvaluationWindow != nil {
		in, out := &in.EvaluationWindow, &out.EvaluationWindow
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.LegacyID != nil {
		in, out := &in.LegacyID, &out.LegacyID
		*out = new(string)
		**out = **in
	}
	if in.MetricExpression != nil {
		in, out := &in.MetricExpression, &out.MetricExpression
		*out = new(string)
		**out = **in
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TargetSuccess != nil {
		in, out := &in.TargetSuccess, &out.TargetSuccess
		*out = new(float64)
		**out = **in
	}
	if in.TargetWarning != nil {
		in, out := &in.TargetWarning, &out.TargetWarning
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLOV2Parameters.
func (in *SLOV2Parameters) DeepCopy() *SLOV2Parameters {
	if in == nil {
		return nil
	}
	out := new(SLOV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLOV2Spec) DeepCopyInto(out *SLOV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLOV2Spec.
func (in *SLOV2Spec) DeepCopy() *SLOV2Spec {
	if in == nil {
		return nil
	}
	out := new(SLOV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLOV2Status) DeepCopyInto(out *SLOV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLOV2Status.
func (in *SLOV2Status) DeepCopy() *SLOV2Status {
	if in == nil {
		return nil
	}
	out := new(SLOV2Status)
	in.DeepCopyInto(out)
	return out
}
