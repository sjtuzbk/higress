// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1/http_2_rpc.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using Http2Rpc within kubernetes types, where deepcopy-gen is used.
func (in *Http2Rpc) DeepCopyInto(out *Http2Rpc) {
	p := proto.Clone(in).(*Http2Rpc)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Http2Rpc. Required by controller-gen.
func (in *Http2Rpc) DeepCopy() *Http2Rpc {
	if in == nil {
		return nil
	}
	out := new(Http2Rpc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Http2Rpc. Required by controller-gen.
func (in *Http2Rpc) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using DubboService within kubernetes types, where deepcopy-gen is used.
func (in *DubboService) DeepCopyInto(out *DubboService) {
	p := proto.Clone(in).(*DubboService)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DubboService. Required by controller-gen.
func (in *DubboService) DeepCopy() *DubboService {
	if in == nil {
		return nil
	}
	out := new(DubboService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new DubboService. Required by controller-gen.
func (in *DubboService) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Method within kubernetes types, where deepcopy-gen is used.
func (in *Method) DeepCopyInto(out *Method) {
	p := proto.Clone(in).(*Method)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Method. Required by controller-gen.
func (in *Method) DeepCopy() *Method {
	if in == nil {
		return nil
	}
	out := new(Method)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Method. Required by controller-gen.
func (in *Method) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Param within kubernetes types, where deepcopy-gen is used.
func (in *Param) DeepCopyInto(out *Param) {
	p := proto.Clone(in).(*Param)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Param. Required by controller-gen.
func (in *Param) DeepCopy() *Param {
	if in == nil {
		return nil
	}
	out := new(Param)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Param. Required by controller-gen.
func (in *Param) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ParamFromEntireBody within kubernetes types, where deepcopy-gen is used.
func (in *ParamFromEntireBody) DeepCopyInto(out *ParamFromEntireBody) {
	p := proto.Clone(in).(*ParamFromEntireBody)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParamFromEntireBody. Required by controller-gen.
func (in *ParamFromEntireBody) DeepCopy() *ParamFromEntireBody {
	if in == nil {
		return nil
	}
	out := new(ParamFromEntireBody)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ParamFromEntireBody. Required by controller-gen.
func (in *ParamFromEntireBody) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GrpcService within kubernetes types, where deepcopy-gen is used.
func (in *GrpcService) DeepCopyInto(out *GrpcService) {
	p := proto.Clone(in).(*GrpcService)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrpcService. Required by controller-gen.
func (in *GrpcService) DeepCopy() *GrpcService {
	if in == nil {
		return nil
	}
	out := new(GrpcService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GrpcService. Required by controller-gen.
func (in *GrpcService) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
