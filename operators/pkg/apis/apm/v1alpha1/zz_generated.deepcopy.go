// +build !ignore_autogenerated

// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	commonv1alpha1 "github.com/elastic/k8s-operators/operators/pkg/apis/common/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApmServer) DeepCopyInto(out *ApmServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApmServer.
func (in *ApmServer) DeepCopy() *ApmServer {
	if in == nil {
		return nil
	}
	out := new(ApmServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApmServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApmServerList) DeepCopyInto(out *ApmServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApmServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApmServerList.
func (in *ApmServerList) DeepCopy() *ApmServerList {
	if in == nil {
		return nil
	}
	out := new(ApmServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApmServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApmServerSpec) DeepCopyInto(out *ApmServerSpec) {
	*out = *in
	in.Output.DeepCopyInto(&out.Output)
	in.Resources.DeepCopyInto(&out.Resources)
	if in.FeatureFlags != nil {
		in, out := &in.FeatureFlags, &out.FeatureFlags
		*out = make(commonv1alpha1.FeatureFlags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApmServerSpec.
func (in *ApmServerSpec) DeepCopy() *ApmServerSpec {
	if in == nil {
		return nil
	}
	out := new(ApmServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApmServerStatus) DeepCopyInto(out *ApmServerStatus) {
	*out = *in
	out.ReconcilerStatus = in.ReconcilerStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApmServerStatus.
func (in *ApmServerStatus) DeepCopy() *ApmServerStatus {
	if in == nil {
		return nil
	}
	out := new(ApmServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchAuth) DeepCopyInto(out *ElasticsearchAuth) {
	*out = *in
	if in.Inline != nil {
		in, out := &in.Inline, &out.Inline
		*out = new(ElasticsearchInlineAuth)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchAuth.
func (in *ElasticsearchAuth) DeepCopy() *ElasticsearchAuth {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchInlineAuth) DeepCopyInto(out *ElasticsearchInlineAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchInlineAuth.
func (in *ElasticsearchInlineAuth) DeepCopy() *ElasticsearchInlineAuth {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchInlineAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchOutput) DeepCopyInto(out *ElasticsearchOutput) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Auth.DeepCopyInto(&out.Auth)
	in.SSL.DeepCopyInto(&out.SSL)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchOutput.
func (in *ElasticsearchOutput) DeepCopy() *ElasticsearchOutput {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchOutputSSL) DeepCopyInto(out *ElasticsearchOutputSSL) {
	*out = *in
	if in.CertificateAuthoritiesSecret != nil {
		in, out := &in.CertificateAuthoritiesSecret, &out.CertificateAuthoritiesSecret
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchOutputSSL.
func (in *ElasticsearchOutputSSL) DeepCopy() *ElasticsearchOutputSSL {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchOutputSSL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Output) DeepCopyInto(out *Output) {
	*out = *in
	in.Elasticsearch.DeepCopyInto(&out.Elasticsearch)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Output.
func (in *Output) DeepCopy() *Output {
	if in == nil {
		return nil
	}
	out := new(Output)
	in.DeepCopyInto(out)
	return out
}
