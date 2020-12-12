// Copyright 2020 by codex.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v4.0.0
// source: cello.proto

package rpccello

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	rpcmsgType "hcc/harp/action/grpc/pb/rpcmsgType"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Symbols defined in public import of msgType.proto.

type Empty = rpcmsgType.Empty
type HccError = rpcmsgType.HccError
type Node = rpcmsgType.Node
type NodeDetail = rpcmsgType.NodeDetail
type Server = rpcmsgType.Server
type ServerNode = rpcmsgType.ServerNode
type Quota = rpcmsgType.Quota
type VNC = rpcmsgType.VNC
type Volume = rpcmsgType.Volume
type VolumeAttachment = rpcmsgType.VolumeAttachment
type Pool = rpcmsgType.Pool
type AdaptiveIPSetting = rpcmsgType.AdaptiveIPSetting
type AdaptiveIPAvailableIPList = rpcmsgType.AdaptiveIPAvailableIPList
type AdaptiveIPServer = rpcmsgType.AdaptiveIPServer
type Subnet = rpcmsgType.Subnet
type MetricInfo = rpcmsgType.MetricInfo
type MonitoringData = rpcmsgType.MonitoringData
type NormalAction = rpcmsgType.NormalAction
type HccAction = rpcmsgType.HccAction
type Action = rpcmsgType.Action
type Control = rpcmsgType.Control
type Controls = rpcmsgType.Controls
type ScheduledNodes = rpcmsgType.ScheduledNodes
type ScheduleServer = rpcmsgType.ScheduleServer
type ServerAction = rpcmsgType.ServerAction

type ReqVolumeHandler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume *rpcmsgType.Volume `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *ReqVolumeHandler) Reset() {
	*x = ReqVolumeHandler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqVolumeHandler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqVolumeHandler) ProtoMessage() {}

func (x *ReqVolumeHandler) ProtoReflect() protoreflect.Message {
	mi := &file_cello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqVolumeHandler.ProtoReflect.Descriptor instead.
func (*ReqVolumeHandler) Descriptor() ([]byte, []int) {
	return file_cello_proto_rawDescGZIP(), []int{0}
}

func (x *ReqVolumeHandler) GetVolume() *rpcmsgType.Volume {
	if x != nil {
		return x.Volume
	}
	return nil
}

type ResVolumeHandler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume        *rpcmsgType.Volume     `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
	HccErrorStack []*rpcmsgType.HccError `protobuf:"bytes,2,rep,name=hccErrorStack,proto3" json:"hccErrorStack,omitempty"`
}

func (x *ResVolumeHandler) Reset() {
	*x = ResVolumeHandler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResVolumeHandler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResVolumeHandler) ProtoMessage() {}

func (x *ResVolumeHandler) ProtoReflect() protoreflect.Message {
	mi := &file_cello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResVolumeHandler.ProtoReflect.Descriptor instead.
func (*ResVolumeHandler) Descriptor() ([]byte, []int) {
	return file_cello_proto_rawDescGZIP(), []int{1}
}

func (x *ResVolumeHandler) GetVolume() *rpcmsgType.Volume {
	if x != nil {
		return x.Volume
	}
	return nil
}

func (x *ResVolumeHandler) GetHccErrorStack() []*rpcmsgType.HccError {
	if x != nil {
		return x.HccErrorStack
	}
	return nil
}

type ReqPoolHandler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pool *rpcmsgType.Pool `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
}

func (x *ReqPoolHandler) Reset() {
	*x = ReqPoolHandler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqPoolHandler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPoolHandler) ProtoMessage() {}

func (x *ReqPoolHandler) ProtoReflect() protoreflect.Message {
	mi := &file_cello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPoolHandler.ProtoReflect.Descriptor instead.
func (*ReqPoolHandler) Descriptor() ([]byte, []int) {
	return file_cello_proto_rawDescGZIP(), []int{2}
}

func (x *ReqPoolHandler) GetPool() *rpcmsgType.Pool {
	if x != nil {
		return x.Pool
	}
	return nil
}

type ResPoolHandler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pool          *rpcmsgType.Pool       `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
	HccErrorStack []*rpcmsgType.HccError `protobuf:"bytes,2,rep,name=hccErrorStack,proto3" json:"hccErrorStack,omitempty"`
}

func (x *ResPoolHandler) Reset() {
	*x = ResPoolHandler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResPoolHandler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResPoolHandler) ProtoMessage() {}

func (x *ResPoolHandler) ProtoReflect() protoreflect.Message {
	mi := &file_cello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResPoolHandler.ProtoReflect.Descriptor instead.
func (*ResPoolHandler) Descriptor() ([]byte, []int) {
	return file_cello_proto_rawDescGZIP(), []int{3}
}

func (x *ResPoolHandler) GetPool() *rpcmsgType.Pool {
	if x != nil {
		return x.Pool
	}
	return nil
}

func (x *ResPoolHandler) GetHccErrorStack() []*rpcmsgType.HccError {
	if x != nil {
		return x.HccErrorStack
	}
	return nil
}

type ReqGetVolume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume *rpcmsgType.Volume `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
	Row    int64              `protobuf:"varint,2,opt,name=row,proto3" json:"row,omitempty"`
	Page   int64              `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ReqGetVolume) Reset() {
	*x = ReqGetVolume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cello_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetVolume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetVolume) ProtoMessage() {}

func (x *ReqGetVolume) ProtoReflect() protoreflect.Message {
	mi := &file_cello_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetVolume.ProtoReflect.Descriptor instead.
func (*ReqGetVolume) Descriptor() ([]byte, []int) {
	return file_cello_proto_rawDescGZIP(), []int{4}
}

func (x *ReqGetVolume) GetVolume() *rpcmsgType.Volume {
	if x != nil {
		return x.Volume
	}
	return nil
}

func (x *ReqGetVolume) GetRow() int64 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *ReqGetVolume) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ResGetVolume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume        []*rpcmsgType.Volume   `protobuf:"bytes,1,rep,name=volume,proto3" json:"volume,omitempty"`
	HccErrorStack []*rpcmsgType.HccError `protobuf:"bytes,2,rep,name=hccErrorStack,proto3" json:"hccErrorStack,omitempty"`
}

func (x *ResGetVolume) Reset() {
	*x = ResGetVolume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cello_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResGetVolume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGetVolume) ProtoMessage() {}

func (x *ResGetVolume) ProtoReflect() protoreflect.Message {
	mi := &file_cello_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGetVolume.ProtoReflect.Descriptor instead.
func (*ResGetVolume) Descriptor() ([]byte, []int) {
	return file_cello_proto_rawDescGZIP(), []int{5}
}

func (x *ResGetVolume) GetVolume() []*rpcmsgType.Volume {
	if x != nil {
		return x.Volume
	}
	return nil
}

func (x *ResGetVolume) GetHccErrorStack() []*rpcmsgType.HccError {
	if x != nil {
		return x.HccErrorStack
	}
	return nil
}

type ReqGetVolumeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume *rpcmsgType.Volume `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
	Row    int64              `protobuf:"varint,2,opt,name=row,proto3" json:"row,omitempty"`
	Page   int64              `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ReqGetVolumeList) Reset() {
	*x = ReqGetVolumeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cello_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetVolumeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetVolumeList) ProtoMessage() {}

func (x *ReqGetVolumeList) ProtoReflect() protoreflect.Message {
	mi := &file_cello_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetVolumeList.ProtoReflect.Descriptor instead.
func (*ReqGetVolumeList) Descriptor() ([]byte, []int) {
	return file_cello_proto_rawDescGZIP(), []int{6}
}

func (x *ReqGetVolumeList) GetVolume() *rpcmsgType.Volume {
	if x != nil {
		return x.Volume
	}
	return nil
}

func (x *ReqGetVolumeList) GetRow() int64 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *ReqGetVolumeList) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ResGetVolumeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume        []*rpcmsgType.Volume   `protobuf:"bytes,1,rep,name=volume,proto3" json:"volume,omitempty"`
	HccErrorStack []*rpcmsgType.HccError `protobuf:"bytes,2,rep,name=hccErrorStack,proto3" json:"hccErrorStack,omitempty"`
}

func (x *ResGetVolumeList) Reset() {
	*x = ResGetVolumeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cello_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResGetVolumeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGetVolumeList) ProtoMessage() {}

func (x *ResGetVolumeList) ProtoReflect() protoreflect.Message {
	mi := &file_cello_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGetVolumeList.ProtoReflect.Descriptor instead.
func (*ResGetVolumeList) Descriptor() ([]byte, []int) {
	return file_cello_proto_rawDescGZIP(), []int{7}
}

func (x *ResGetVolumeList) GetVolume() []*rpcmsgType.Volume {
	if x != nil {
		return x.Volume
	}
	return nil
}

func (x *ResGetVolumeList) GetHccErrorStack() []*rpcmsgType.HccError {
	if x != nil {
		return x.HccErrorStack
	}
	return nil
}

type ReqGetPoolList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pool *rpcmsgType.Pool `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
	Row  int64            `protobuf:"varint,2,opt,name=row,proto3" json:"row,omitempty"`
	Page int64            `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ReqGetPoolList) Reset() {
	*x = ReqGetPoolList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cello_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetPoolList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetPoolList) ProtoMessage() {}

func (x *ReqGetPoolList) ProtoReflect() protoreflect.Message {
	mi := &file_cello_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetPoolList.ProtoReflect.Descriptor instead.
func (*ReqGetPoolList) Descriptor() ([]byte, []int) {
	return file_cello_proto_rawDescGZIP(), []int{8}
}

func (x *ReqGetPoolList) GetPool() *rpcmsgType.Pool {
	if x != nil {
		return x.Pool
	}
	return nil
}

func (x *ReqGetPoolList) GetRow() int64 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *ReqGetPoolList) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ResGetPoolList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pool          []*rpcmsgType.Pool     `protobuf:"bytes,1,rep,name=pool,proto3" json:"pool,omitempty"`
	HccErrorStack []*rpcmsgType.HccError `protobuf:"bytes,2,rep,name=hccErrorStack,proto3" json:"hccErrorStack,omitempty"`
}

func (x *ResGetPoolList) Reset() {
	*x = ResGetPoolList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cello_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResGetPoolList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGetPoolList) ProtoMessage() {}

func (x *ResGetPoolList) ProtoReflect() protoreflect.Message {
	mi := &file_cello_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGetPoolList.ProtoReflect.Descriptor instead.
func (*ResGetPoolList) Descriptor() ([]byte, []int) {
	return file_cello_proto_rawDescGZIP(), []int{9}
}

func (x *ResGetPoolList) GetPool() []*rpcmsgType.Pool {
	if x != nil {
		return x.Pool
	}
	return nil
}

func (x *ResGetPoolList) GetHccErrorStack() []*rpcmsgType.HccError {
	if x != nil {
		return x.HccErrorStack
	}
	return nil
}

var File_cello_proto protoreflect.FileDescriptor

var file_cello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x52,
	0x70, 0x63, 0x43, 0x65, 0x6c, 0x6c, 0x6f, 0x1a, 0x0d, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x06, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x22, 0x74, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x37, 0x0a, 0x0d, 0x68, 0x63, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x48, 0x63, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x0d, 0x68, 0x63, 0x63, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x22, 0x33, 0x0a, 0x0e, 0x52, 0x65, 0x71,
	0x50, 0x6f, 0x6f, 0x6c, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x04, 0x70,
	0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4d, 0x73, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x22, 0x6c,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x50, 0x6f, 0x6f, 0x6c, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x12, 0x21, 0x0a, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x70,
	0x6f, 0x6f, 0x6c, 0x12, 0x37, 0x0a, 0x0d, 0x68, 0x63, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x48, 0x63, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x0d, 0x68,
	0x63, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x22, 0x5d, 0x0a, 0x0c,
	0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x06,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d,
	0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x70, 0x0a, 0x0c, 0x52,
	0x65, 0x73, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x68, 0x63, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x48, 0x63, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x0d,
	0x68, 0x63, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x22, 0x61, 0x0a,
	0x10, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f,
	0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x22, 0x74, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x37, 0x0a,
	0x0d, 0x68, 0x63, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x48,
	0x63, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x0d, 0x68, 0x63, 0x63, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x22, 0x59, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x6f, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x70, 0x6f, 0x6f, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x22, 0x6c, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x6f, 0x6c,
	0x52, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x12, 0x37, 0x0a, 0x0d, 0x68, 0x63, 0x63, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x48, 0x63, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x0d, 0x68, 0x63, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x32,
	0xe6, 0x02, 0x0a, 0x05, 0x43, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x49, 0x0a, 0x0d, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x52, 0x70, 0x63,
	0x43, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x1a, 0x1a, 0x2e, 0x52, 0x70, 0x63, 0x43, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x50, 0x6f, 0x6f, 0x6c, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x52, 0x70, 0x63, 0x43, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x50, 0x6f, 0x6f, 0x6c, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x1a, 0x18, 0x2e,
	0x52, 0x70, 0x63, 0x43, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x50, 0x6f, 0x6f, 0x6c,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x52, 0x70, 0x63, 0x43, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x1a, 0x16,
	0x2e, 0x52, 0x70, 0x63, 0x43, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x52, 0x70, 0x63, 0x43,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x52, 0x70, 0x63, 0x43, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x18, 0x2e, 0x52, 0x70, 0x63, 0x43, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x52, 0x65,
	0x71, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x52,
	0x70, 0x63, 0x43, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x6f, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x68, 0x63, 0x63, 0x2f,
	0x68, 0x61, 0x72, 0x70, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x50, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cello_proto_rawDescOnce sync.Once
	file_cello_proto_rawDescData = file_cello_proto_rawDesc
)

func file_cello_proto_rawDescGZIP() []byte {
	file_cello_proto_rawDescOnce.Do(func() {
		file_cello_proto_rawDescData = protoimpl.X.CompressGZIP(file_cello_proto_rawDescData)
	})
	return file_cello_proto_rawDescData
}

var file_cello_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cello_proto_goTypes = []interface{}{
	(*ReqVolumeHandler)(nil),    // 0: RpcCello.ReqVolumeHandler
	(*ResVolumeHandler)(nil),    // 1: RpcCello.ResVolumeHandler
	(*ReqPoolHandler)(nil),      // 2: RpcCello.ReqPoolHandler
	(*ResPoolHandler)(nil),      // 3: RpcCello.ResPoolHandler
	(*ReqGetVolume)(nil),        // 4: RpcCello.ReqGetVolume
	(*ResGetVolume)(nil),        // 5: RpcCello.ResGetVolume
	(*ReqGetVolumeList)(nil),    // 6: RpcCello.ReqGetVolumeList
	(*ResGetVolumeList)(nil),    // 7: RpcCello.ResGetVolumeList
	(*ReqGetPoolList)(nil),      // 8: RpcCello.ReqGetPoolList
	(*ResGetPoolList)(nil),      // 9: RpcCello.ResGetPoolList
	(*rpcmsgType.Volume)(nil),   // 10: MsgType.Volume
	(*rpcmsgType.HccError)(nil), // 11: MsgType.HccError
	(*rpcmsgType.Pool)(nil),     // 12: MsgType.Pool
}
var file_cello_proto_depIdxs = []int32{
	10, // 0: RpcCello.ReqVolumeHandler.volume:type_name -> MsgType.Volume
	10, // 1: RpcCello.ResVolumeHandler.volume:type_name -> MsgType.Volume
	11, // 2: RpcCello.ResVolumeHandler.hccErrorStack:type_name -> MsgType.HccError
	12, // 3: RpcCello.ReqPoolHandler.pool:type_name -> MsgType.Pool
	12, // 4: RpcCello.ResPoolHandler.pool:type_name -> MsgType.Pool
	11, // 5: RpcCello.ResPoolHandler.hccErrorStack:type_name -> MsgType.HccError
	10, // 6: RpcCello.ReqGetVolume.volume:type_name -> MsgType.Volume
	10, // 7: RpcCello.ResGetVolume.volume:type_name -> MsgType.Volume
	11, // 8: RpcCello.ResGetVolume.hccErrorStack:type_name -> MsgType.HccError
	10, // 9: RpcCello.ReqGetVolumeList.volume:type_name -> MsgType.Volume
	10, // 10: RpcCello.ResGetVolumeList.volume:type_name -> MsgType.Volume
	11, // 11: RpcCello.ResGetVolumeList.hccErrorStack:type_name -> MsgType.HccError
	12, // 12: RpcCello.ReqGetPoolList.pool:type_name -> MsgType.Pool
	12, // 13: RpcCello.ResGetPoolList.pool:type_name -> MsgType.Pool
	11, // 14: RpcCello.ResGetPoolList.hccErrorStack:type_name -> MsgType.HccError
	0,  // 15: RpcCello.Cello.VolumeHandler:input_type -> RpcCello.ReqVolumeHandler
	2,  // 16: RpcCello.Cello.PoolHandler:input_type -> RpcCello.ReqPoolHandler
	4,  // 17: RpcCello.Cello.GetVolume:input_type -> RpcCello.ReqGetVolume
	6,  // 18: RpcCello.Cello.GetVolumeList:input_type -> RpcCello.ReqGetVolumeList
	8,  // 19: RpcCello.Cello.GetPoolList:input_type -> RpcCello.ReqGetPoolList
	1,  // 20: RpcCello.Cello.VolumeHandler:output_type -> RpcCello.ResVolumeHandler
	3,  // 21: RpcCello.Cello.PoolHandler:output_type -> RpcCello.ResPoolHandler
	5,  // 22: RpcCello.Cello.GetVolume:output_type -> RpcCello.ResGetVolume
	7,  // 23: RpcCello.Cello.GetVolumeList:output_type -> RpcCello.ResGetVolumeList
	9,  // 24: RpcCello.Cello.GetPoolList:output_type -> RpcCello.ResGetPoolList
	20, // [20:25] is the sub-list for method output_type
	15, // [15:20] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_cello_proto_init() }
func file_cello_proto_init() {
	if File_cello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqVolumeHandler); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResVolumeHandler); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqPoolHandler); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResPoolHandler); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cello_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetVolume); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cello_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResGetVolume); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cello_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetVolumeList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cello_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResGetVolumeList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cello_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetPoolList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cello_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResGetPoolList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cello_proto_goTypes,
		DependencyIndexes: file_cello_proto_depIdxs,
		MessageInfos:      file_cello_proto_msgTypes,
	}.Build()
	File_cello_proto = out.File
	file_cello_proto_rawDesc = nil
	file_cello_proto_goTypes = nil
	file_cello_proto_depIdxs = nil
}
