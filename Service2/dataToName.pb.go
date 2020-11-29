// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: dataToName.proto

package dataToName

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type DistributionRequest2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName   string                                     `protobuf:"bytes,1,opt,name=FileName,proto3" json:"FileName,omitempty"`
	TotalParts int32                                      `protobuf:"varint,2,opt,name=TotalParts,proto3" json:"TotalParts,omitempty"`
	Machines   []*DistributionRequest2_MachineInformation `protobuf:"bytes,3,rep,name=Machines,proto3" json:"Machines,omitempty"`
}

func (x *DistributionRequest2) Reset() {
	*x = DistributionRequest2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataToName_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributionRequest2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributionRequest2) ProtoMessage() {}

func (x *DistributionRequest2) ProtoReflect() protoreflect.Message {
	mi := &file_dataToName_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributionRequest2.ProtoReflect.Descriptor instead.
func (*DistributionRequest2) Descriptor() ([]byte, []int) {
	return file_dataToName_proto_rawDescGZIP(), []int{0}
}

func (x *DistributionRequest2) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *DistributionRequest2) GetTotalParts() int32 {
	if x != nil {
		return x.TotalParts
	}
	return 0
}

func (x *DistributionRequest2) GetMachines() []*DistributionRequest2_MachineInformation {
	if x != nil {
		return x.Machines
	}
	return nil
}

type DistributionReply2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName   string                                   `protobuf:"bytes,1,opt,name=FileName,proto3" json:"FileName,omitempty"`
	TotalParts int32                                    `protobuf:"varint,2,opt,name=TotalParts,proto3" json:"TotalParts,omitempty"`
	Machines   []*DistributionReply2_MachineInformation `protobuf:"bytes,3,rep,name=Machines,proto3" json:"Machines,omitempty"`
}

func (x *DistributionReply2) Reset() {
	*x = DistributionReply2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataToName_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributionReply2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributionReply2) ProtoMessage() {}

func (x *DistributionReply2) ProtoReflect() protoreflect.Message {
	mi := &file_dataToName_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributionReply2.ProtoReflect.Descriptor instead.
func (*DistributionReply2) Descriptor() ([]byte, []int) {
	return file_dataToName_proto_rawDescGZIP(), []int{1}
}

func (x *DistributionReply2) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *DistributionReply2) GetTotalParts() int32 {
	if x != nil {
		return x.TotalParts
	}
	return 0
}

func (x *DistributionReply2) GetMachines() []*DistributionReply2_MachineInformation {
	if x != nil {
		return x.Machines
	}
	return nil
}

type DistributionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName   string                                    `protobuf:"bytes,1,opt,name=FileName,proto3" json:"FileName,omitempty"`
	TotalParts int32                                     `protobuf:"varint,2,opt,name=TotalParts,proto3" json:"TotalParts,omitempty"`
	Machines   []*DistributionRequest_MachineInformation `protobuf:"bytes,3,rep,name=Machines,proto3" json:"Machines,omitempty"`
}

func (x *DistributionRequest) Reset() {
	*x = DistributionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataToName_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributionRequest) ProtoMessage() {}

func (x *DistributionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dataToName_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributionRequest.ProtoReflect.Descriptor instead.
func (*DistributionRequest) Descriptor() ([]byte, []int) {
	return file_dataToName_proto_rawDescGZIP(), []int{2}
}

func (x *DistributionRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *DistributionRequest) GetTotalParts() int32 {
	if x != nil {
		return x.TotalParts
	}
	return 0
}

func (x *DistributionRequest) GetMachines() []*DistributionRequest_MachineInformation {
	if x != nil {
		return x.Machines
	}
	return nil
}

type DistributionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName   string                                  `protobuf:"bytes,1,opt,name=FileName,proto3" json:"FileName,omitempty"`
	TotalParts int32                                   `protobuf:"varint,2,opt,name=TotalParts,proto3" json:"TotalParts,omitempty"`
	Machines   []*DistributionReply_MachineInformation `protobuf:"bytes,3,rep,name=Machines,proto3" json:"Machines,omitempty"`
}

func (x *DistributionReply) Reset() {
	*x = DistributionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataToName_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributionReply) ProtoMessage() {}

func (x *DistributionReply) ProtoReflect() protoreflect.Message {
	mi := &file_dataToName_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributionReply.ProtoReflect.Descriptor instead.
func (*DistributionReply) Descriptor() ([]byte, []int) {
	return file_dataToName_proto_rawDescGZIP(), []int{3}
}

func (x *DistributionReply) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *DistributionReply) GetTotalParts() int32 {
	if x != nil {
		return x.TotalParts
	}
	return 0
}

func (x *DistributionReply) GetMachines() []*DistributionReply_MachineInformation {
	if x != nil {
		return x.Machines
	}
	return nil
}

type DistributionRequest2_MachineInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      string  `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Distribution []int32 `protobuf:"varint,2,rep,packed,name=Distribution,proto3" json:"Distribution,omitempty"`
	Status       int32   `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *DistributionRequest2_MachineInformation) Reset() {
	*x = DistributionRequest2_MachineInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataToName_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributionRequest2_MachineInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributionRequest2_MachineInformation) ProtoMessage() {}

func (x *DistributionRequest2_MachineInformation) ProtoReflect() protoreflect.Message {
	mi := &file_dataToName_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributionRequest2_MachineInformation.ProtoReflect.Descriptor instead.
func (*DistributionRequest2_MachineInformation) Descriptor() ([]byte, []int) {
	return file_dataToName_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DistributionRequest2_MachineInformation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DistributionRequest2_MachineInformation) GetDistribution() []int32 {
	if x != nil {
		return x.Distribution
	}
	return nil
}

func (x *DistributionRequest2_MachineInformation) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DistributionReply2_MachineInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      string  `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Distribution []int32 `protobuf:"varint,2,rep,packed,name=Distribution,proto3" json:"Distribution,omitempty"`
	Status       int32   `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *DistributionReply2_MachineInformation) Reset() {
	*x = DistributionReply2_MachineInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataToName_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributionReply2_MachineInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributionReply2_MachineInformation) ProtoMessage() {}

func (x *DistributionReply2_MachineInformation) ProtoReflect() protoreflect.Message {
	mi := &file_dataToName_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributionReply2_MachineInformation.ProtoReflect.Descriptor instead.
func (*DistributionReply2_MachineInformation) Descriptor() ([]byte, []int) {
	return file_dataToName_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DistributionReply2_MachineInformation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DistributionReply2_MachineInformation) GetDistribution() []int32 {
	if x != nil {
		return x.Distribution
	}
	return nil
}

func (x *DistributionReply2_MachineInformation) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DistributionRequest_MachineInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      string  `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Distribution []int32 `protobuf:"varint,2,rep,packed,name=Distribution,proto3" json:"Distribution,omitempty"`
	Status       int32   `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *DistributionRequest_MachineInformation) Reset() {
	*x = DistributionRequest_MachineInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataToName_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributionRequest_MachineInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributionRequest_MachineInformation) ProtoMessage() {}

func (x *DistributionRequest_MachineInformation) ProtoReflect() protoreflect.Message {
	mi := &file_dataToName_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributionRequest_MachineInformation.ProtoReflect.Descriptor instead.
func (*DistributionRequest_MachineInformation) Descriptor() ([]byte, []int) {
	return file_dataToName_proto_rawDescGZIP(), []int{2, 0}
}

func (x *DistributionRequest_MachineInformation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DistributionRequest_MachineInformation) GetDistribution() []int32 {
	if x != nil {
		return x.Distribution
	}
	return nil
}

func (x *DistributionRequest_MachineInformation) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DistributionReply_MachineInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      string  `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Distribution []int32 `protobuf:"varint,2,rep,packed,name=Distribution,proto3" json:"Distribution,omitempty"`
	Status       int32   `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *DistributionReply_MachineInformation) Reset() {
	*x = DistributionReply_MachineInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataToName_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributionReply_MachineInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributionReply_MachineInformation) ProtoMessage() {}

func (x *DistributionReply_MachineInformation) ProtoReflect() protoreflect.Message {
	mi := &file_dataToName_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributionReply_MachineInformation.ProtoReflect.Descriptor instead.
func (*DistributionReply_MachineInformation) Descriptor() ([]byte, []int) {
	return file_dataToName_proto_rawDescGZIP(), []int{3, 0}
}

func (x *DistributionReply_MachineInformation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DistributionReply_MachineInformation) GetDistribution() []int32 {
	if x != nil {
		return x.Distribution
	}
	return nil
}

func (x *DistributionReply_MachineInformation) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_dataToName_proto protoreflect.FileDescriptor

var file_dataToName_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x8f,
	0x02, 0x0a, 0x14, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61,
	0x72, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x08, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x4e, 0x61,
	0x6d, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x73, 0x1a, 0x6a, 0x0a, 0x12, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x8b, 0x02, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61,
	0x72, 0x74, 0x73, 0x12, 0x4d, 0x0a, 0x08, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x4e, 0x61,
	0x6d, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x32, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x73, 0x1a, 0x6a, 0x0a, 0x12, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8d,
	0x02, 0x0a, 0x13, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x72,
	0x74, 0x73, 0x12, 0x4e, 0x0a, 0x08, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x4e, 0x61, 0x6d,
	0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x73, 0x1a, 0x6a, 0x0a, 0x12, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x89,
	0x02, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x73,
	0x12, 0x4c, 0x0a, 0x08, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x1a, 0x6a,
	0x0a, 0x12, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xc9, 0x01, 0x0a, 0x11, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5c, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x1f, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x54, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x56,
	0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x32, 0x1a, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x4e, 0x61, 0x6d,
	0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dataToName_proto_rawDescOnce sync.Once
	file_dataToName_proto_rawDescData = file_dataToName_proto_rawDesc
)

func file_dataToName_proto_rawDescGZIP() []byte {
	file_dataToName_proto_rawDescOnce.Do(func() {
		file_dataToName_proto_rawDescData = protoimpl.X.CompressGZIP(file_dataToName_proto_rawDescData)
	})
	return file_dataToName_proto_rawDescData
}

var file_dataToName_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dataToName_proto_goTypes = []interface{}{
	(*DistributionRequest2)(nil),                    // 0: dataToName.DistributionRequest2
	(*DistributionReply2)(nil),                      // 1: dataToName.DistributionReply2
	(*DistributionRequest)(nil),                     // 2: dataToName.DistributionRequest
	(*DistributionReply)(nil),                       // 3: dataToName.DistributionReply
	(*DistributionRequest2_MachineInformation)(nil), // 4: dataToName.DistributionRequest2.MachineInformation
	(*DistributionReply2_MachineInformation)(nil),   // 5: dataToName.DistributionReply2.MachineInformation
	(*DistributionRequest_MachineInformation)(nil),  // 6: dataToName.DistributionRequest.MachineInformation
	(*DistributionReply_MachineInformation)(nil),    // 7: dataToName.DistributionReply.MachineInformation
}
var file_dataToName_proto_depIdxs = []int32{
	4, // 0: dataToName.DistributionRequest2.Machines:type_name -> dataToName.DistributionRequest2.MachineInformation
	5, // 1: dataToName.DistributionReply2.Machines:type_name -> dataToName.DistributionReply2.MachineInformation
	6, // 2: dataToName.DistributionRequest.Machines:type_name -> dataToName.DistributionRequest.MachineInformation
	7, // 3: dataToName.DistributionReply.Machines:type_name -> dataToName.DistributionReply.MachineInformation
	2, // 4: dataToName.DataToNameService.SendDistributionProposal:input_type -> dataToName.DistributionRequest
	0, // 5: dataToName.DataToNameService.SendDistribution:input_type -> dataToName.DistributionRequest2
	3, // 6: dataToName.DataToNameService.SendDistributionProposal:output_type -> dataToName.DistributionReply
	1, // 7: dataToName.DataToNameService.SendDistribution:output_type -> dataToName.DistributionReply2
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_dataToName_proto_init() }
func file_dataToName_proto_init() {
	if File_dataToName_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dataToName_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributionRequest2); i {
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
		file_dataToName_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributionReply2); i {
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
		file_dataToName_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributionRequest); i {
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
		file_dataToName_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributionReply); i {
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
		file_dataToName_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributionRequest2_MachineInformation); i {
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
		file_dataToName_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributionReply2_MachineInformation); i {
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
		file_dataToName_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributionRequest_MachineInformation); i {
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
		file_dataToName_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributionReply_MachineInformation); i {
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
			RawDescriptor: file_dataToName_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dataToName_proto_goTypes,
		DependencyIndexes: file_dataToName_proto_depIdxs,
		MessageInfos:      file_dataToName_proto_msgTypes,
	}.Build()
	File_dataToName_proto = out.File
	file_dataToName_proto_rawDesc = nil
	file_dataToName_proto_goTypes = nil
	file_dataToName_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataToNameServiceClient is the client API for DataToNameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataToNameServiceClient interface {
	SendDistributionProposal(ctx context.Context, in *DistributionRequest, opts ...grpc.CallOption) (*DistributionReply, error)
	SendDistribution(ctx context.Context, in *DistributionRequest2, opts ...grpc.CallOption) (*DistributionReply2, error)
}

type dataToNameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataToNameServiceClient(cc grpc.ClientConnInterface) DataToNameServiceClient {
	return &dataToNameServiceClient{cc}
}

func (c *dataToNameServiceClient) SendDistributionProposal(ctx context.Context, in *DistributionRequest, opts ...grpc.CallOption) (*DistributionReply, error) {
	out := new(DistributionReply)
	err := c.cc.Invoke(ctx, "/dataToName.DataToNameService/SendDistributionProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataToNameServiceClient) SendDistribution(ctx context.Context, in *DistributionRequest2, opts ...grpc.CallOption) (*DistributionReply2, error) {
	out := new(DistributionReply2)
	err := c.cc.Invoke(ctx, "/dataToName.DataToNameService/SendDistribution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataToNameServiceServer is the server API for DataToNameService service.
type DataToNameServiceServer interface {
	SendDistributionProposal(context.Context, *DistributionRequest) (*DistributionReply, error)
	SendDistribution(context.Context, *DistributionRequest2) (*DistributionReply2, error)
}

// UnimplementedDataToNameServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDataToNameServiceServer struct {
}

func (*UnimplementedDataToNameServiceServer) SendDistributionProposal(context.Context, *DistributionRequest) (*DistributionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDistributionProposal not implemented")
}
func (*UnimplementedDataToNameServiceServer) SendDistribution(context.Context, *DistributionRequest2) (*DistributionReply2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDistribution not implemented")
}

func RegisterDataToNameServiceServer(s *grpc.Server, srv DataToNameServiceServer) {
	s.RegisterService(&_DataToNameService_serviceDesc, srv)
}

func _DataToNameService_SendDistributionProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DistributionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataToNameServiceServer).SendDistributionProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataToName.DataToNameService/SendDistributionProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataToNameServiceServer).SendDistributionProposal(ctx, req.(*DistributionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataToNameService_SendDistribution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DistributionRequest2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataToNameServiceServer).SendDistribution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataToName.DataToNameService/SendDistribution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataToNameServiceServer).SendDistribution(ctx, req.(*DistributionRequest2))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataToNameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataToName.DataToNameService",
	HandlerType: (*DataToNameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendDistributionProposal",
			Handler:    _DataToNameService_SendDistributionProposal_Handler,
		},
		{
			MethodName: "SendDistribution",
			Handler:    _DataToNameService_SendDistribution_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataToName.proto",
}
