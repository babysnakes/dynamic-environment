//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerStatus) DeepCopyInto(out *ConsumerStatus) {
	*out = *in
	out.ResourceStatus = in.ResourceStatus
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]StatusError, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerStatus.
func (in *ConsumerStatus) DeepCopy() *ConsumerStatus {
	if in == nil {
		return nil
	}
	out := new(ConsumerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerOverrides) DeepCopyInto(out *ContainerOverrides) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerOverrides.
func (in *ContainerOverrides) DeepCopy() *ContainerOverrides {
	if in == nil {
		return nil
	}
	out := new(ContainerOverrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicEnv) DeepCopyInto(out *DynamicEnv) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicEnv.
func (in *DynamicEnv) DeepCopy() *DynamicEnv {
	if in == nil {
		return nil
	}
	out := new(DynamicEnv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamicEnv) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicEnvList) DeepCopyInto(out *DynamicEnvList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DynamicEnv, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicEnvList.
func (in *DynamicEnvList) DeepCopy() *DynamicEnvList {
	if in == nil {
		return nil
	}
	out := new(DynamicEnvList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamicEnvList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicEnvSpec) DeepCopyInto(out *DynamicEnvSpec) {
	*out = *in
	if in.IstioMatches != nil {
		in, out := &in.IstioMatches, &out.IstioMatches
		*out = make([]IstioMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Subsets != nil {
		in, out := &in.Subsets, &out.Subsets
		*out = make([]Subset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Consumers != nil {
		in, out := &in.Consumers, &out.Consumers
		*out = make([]Subset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicEnvSpec.
func (in *DynamicEnvSpec) DeepCopy() *DynamicEnvSpec {
	if in == nil {
		return nil
	}
	out := new(DynamicEnvSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicEnvStatus) DeepCopyInto(out *DynamicEnvStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubsetsStatus != nil {
		in, out := &in.SubsetsStatus, &out.SubsetsStatus
		*out = make(map[string]SubsetStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ConsumersStatus != nil {
		in, out := &in.ConsumersStatus, &out.ConsumersStatus
		*out = make(map[string]ConsumerStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicEnvStatus.
func (in *DynamicEnvStatus) DeepCopy() *DynamicEnvStatus {
	if in == nil {
		return nil
	}
	out := new(DynamicEnvStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioMatch) DeepCopyInto(out *IstioMatch) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]StringMatch, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SourceLabels != nil {
		in, out := &in.SourceLabels, &out.SourceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioMatch.
func (in *IstioMatch) DeepCopy() *IstioMatch {
	if in == nil {
		return nil
	}
	out := new(IstioMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceStatus) DeepCopyInto(out *ResourceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceStatus.
func (in *ResourceStatus) DeepCopy() *ResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusError) DeepCopyInto(out *StatusError) {
	*out = *in
	in.LastOccurrence.DeepCopyInto(&out.LastOccurrence)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusError.
func (in *StatusError) DeepCopy() *StatusError {
	if in == nil {
		return nil
	}
	out := new(StatusError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringMatch) DeepCopyInto(out *StringMatch) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringMatch.
func (in *StringMatch) DeepCopy() *StringMatch {
	if in == nil {
		return nil
	}
	out := new(StringMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subset) DeepCopyInto(out *Subset) {
	*out = *in
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]ContainerOverrides, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]ContainerOverrides, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subset.
func (in *Subset) DeepCopy() *Subset {
	if in == nil {
		return nil
	}
	out := new(Subset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubsetErrors) DeepCopyInto(out *SubsetErrors) {
	*out = *in
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = make([]StatusError, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DestinationRule != nil {
		in, out := &in.DestinationRule, &out.DestinationRule
		*out = make([]StatusError, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VirtualServices != nil {
		in, out := &in.VirtualServices, &out.VirtualServices
		*out = make([]StatusError, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Subset != nil {
		in, out := &in.Subset, &out.Subset
		*out = make([]StatusError, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubsetErrors.
func (in *SubsetErrors) DeepCopy() *SubsetErrors {
	if in == nil {
		return nil
	}
	out := new(SubsetErrors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubsetMessages) DeepCopyInto(out *SubsetMessages) {
	*out = *in
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DestinationRule != nil {
		in, out := &in.DestinationRule, &out.DestinationRule
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VirtualService != nil {
		in, out := &in.VirtualService, &out.VirtualService
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GlobalErrors != nil {
		in, out := &in.GlobalErrors, &out.GlobalErrors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubsetMessages.
func (in *SubsetMessages) DeepCopy() *SubsetMessages {
	if in == nil {
		return nil
	}
	out := new(SubsetMessages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubsetStatus) DeepCopyInto(out *SubsetStatus) {
	*out = *in
	out.Deployment = in.Deployment
	if in.DestinationRules != nil {
		in, out := &in.DestinationRules, &out.DestinationRules
		*out = make([]ResourceStatus, len(*in))
		copy(*out, *in)
	}
	if in.VirtualServices != nil {
		in, out := &in.VirtualServices, &out.VirtualServices
		*out = make([]ResourceStatus, len(*in))
		copy(*out, *in)
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = new(SubsetErrors)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubsetStatus.
func (in *SubsetStatus) DeepCopy() *SubsetStatus {
	if in == nil {
		return nil
	}
	out := new(SubsetStatus)
	in.DeepCopyInto(out)
	return out
}