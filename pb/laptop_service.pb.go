// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: laptop_service.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ResponseRequestByID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ResponseRequestByID) Reset() {
	*x = ResponseRequestByID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseRequestByID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRequestByID) ProtoMessage() {}

func (x *ResponseRequestByID) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRequestByID.ProtoReflect.Descriptor instead.
func (*ResponseRequestByID) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseRequestByID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ResponsesLaptop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Laptops []*Laptop `protobuf:"bytes,1,rep,name=laptops,proto3" json:"laptops,omitempty"`
}

func (x *ResponsesLaptop) Reset() {
	*x = ResponsesLaptop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsesLaptop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsesLaptop) ProtoMessage() {}

func (x *ResponsesLaptop) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsesLaptop.ProtoReflect.Descriptor instead.
func (*ResponsesLaptop) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResponsesLaptop) GetLaptops() []*Laptop {
	if x != nil {
		return x.Laptops
	}
	return nil
}

type ResponseRequestLaptop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Laptop *Laptop `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
}

func (x *ResponseRequestLaptop) Reset() {
	*x = ResponseRequestLaptop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseRequestLaptop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRequestLaptop) ProtoMessage() {}

func (x *ResponseRequestLaptop) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRequestLaptop.ProtoReflect.Descriptor instead.
func (*ResponseRequestLaptop) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseRequestLaptop) GetLaptop() *Laptop {
	if x != nil {
		return x.Laptop
	}
	return nil
}

type RequestSearchByFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *RequestSearchByFilter) Reset() {
	*x = RequestSearchByFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestSearchByFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestSearchByFilter) ProtoMessage() {}

func (x *RequestSearchByFilter) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestSearchByFilter.ProtoReflect.Descriptor instead.
func (*RequestSearchByFilter) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{3}
}

func (x *RequestSearchByFilter) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type UploadImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*UploadImageRequest_Info
	//	*UploadImageRequest_ChunkData
	Data isUploadImageRequest_Data `protobuf_oneof:"data"`
}

func (x *UploadImageRequest) Reset() {
	*x = UploadImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageRequest) ProtoMessage() {}

func (x *UploadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageRequest.ProtoReflect.Descriptor instead.
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{4}
}

func (m *UploadImageRequest) GetData() isUploadImageRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadImageRequest) GetInfo() *ImageInfo {
	if x, ok := x.GetData().(*UploadImageRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *UploadImageRequest) GetChunkData() []byte {
	if x, ok := x.GetData().(*UploadImageRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isUploadImageRequest_Data interface {
	isUploadImageRequest_Data()
}

type UploadImageRequest_Info struct {
	Info *ImageInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadImageRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*UploadImageRequest_Info) isUploadImageRequest_Data() {}

func (*UploadImageRequest_ChunkData) isUploadImageRequest_Data() {}

type UploadImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageId string `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	Size    uint32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *UploadImageResponse) Reset() {
	*x = UploadImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageResponse) ProtoMessage() {}

func (x *UploadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageResponse.ProtoReflect.Descriptor instead.
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{5}
}

func (x *UploadImageResponse) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *UploadImageResponse) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_laptop_service_proto protoreflect.FileDescriptor

var file_laptop_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x0f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x21,
	0x0a, 0x07, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x07, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70,
	0x73, 0x22, 0x38, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x1f, 0x0a, 0x06, 0x6c, 0x61,
	0x70, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4c, 0x61, 0x70,
	0x74, 0x6f, 0x70, 0x52, 0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x22, 0x38, 0x0a, 0x15, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x5f, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a,
	0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x44, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x32, 0x91, 0x03, 0x0a,
	0x0d, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x16,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x1a, 0x14, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x3d, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x42, 0x79, 0x49, 0x44, 0x12, 0x14, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x38, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x4c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x40, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x42, 0x79, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x30, 0x01, 0x12, 0x40,
	0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x70,
	0x74, 0x6f, 0x70, 0x12, 0x13, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01,
	0x42, 0x05, 0x5a, 0x03, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_laptop_service_proto_rawDescOnce sync.Once
	file_laptop_service_proto_rawDescData = file_laptop_service_proto_rawDesc
)

func file_laptop_service_proto_rawDescGZIP() []byte {
	file_laptop_service_proto_rawDescOnce.Do(func() {
		file_laptop_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_laptop_service_proto_rawDescData)
	})
	return file_laptop_service_proto_rawDescData
}

var file_laptop_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_laptop_service_proto_goTypes = []interface{}{
	(*ResponseRequestByID)(nil),   // 0: ResponseRequestByID
	(*ResponsesLaptop)(nil),       // 1: ResponsesLaptop
	(*ResponseRequestLaptop)(nil), // 2: ResponseRequestLaptop
	(*RequestSearchByFilter)(nil), // 3: RequestSearchByFilter
	(*UploadImageRequest)(nil),    // 4: UploadImageRequest
	(*UploadImageResponse)(nil),   // 5: UploadImageResponse
	(*Laptop)(nil),                // 6: Laptop
	(*Filter)(nil),                // 7: Filter
	(*ImageInfo)(nil),             // 8: ImageInfo
	(*empty.Empty)(nil),           // 9: google.protobuf.Empty
}
var file_laptop_service_proto_depIdxs = []int32{
	6,  // 0: ResponsesLaptop.laptops:type_name -> Laptop
	6,  // 1: ResponseRequestLaptop.laptop:type_name -> Laptop
	7,  // 2: RequestSearchByFilter.filter:type_name -> Filter
	8,  // 3: UploadImageRequest.info:type_name -> ImageInfo
	2,  // 4: LaptopService.CreateLaptop:input_type -> ResponseRequestLaptop
	0,  // 5: LaptopService.GetLaptopByID:input_type -> ResponseRequestByID
	9,  // 6: LaptopService.GetAllLaptop:input_type -> google.protobuf.Empty
	0,  // 7: LaptopService.DeleteLaptopByID:input_type -> ResponseRequestByID
	3,  // 8: LaptopService.GetSearchByFilter:input_type -> RequestSearchByFilter
	4,  // 9: LaptopService.UploadImageLaptop:input_type -> UploadImageRequest
	0,  // 10: LaptopService.CreateLaptop:output_type -> ResponseRequestByID
	2,  // 11: LaptopService.GetLaptopByID:output_type -> ResponseRequestLaptop
	1,  // 12: LaptopService.GetAllLaptop:output_type -> ResponsesLaptop
	9,  // 13: LaptopService.DeleteLaptopByID:output_type -> google.protobuf.Empty
	2,  // 14: LaptopService.GetSearchByFilter:output_type -> ResponseRequestLaptop
	5,  // 15: LaptopService.UploadImageLaptop:output_type -> UploadImageResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_laptop_service_proto_init() }
func file_laptop_service_proto_init() {
	if File_laptop_service_proto != nil {
		return
	}
	file_laptop_message_proto_init()
	file_filter_message_proto_init()
	file_image_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_laptop_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseRequestByID); i {
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
		file_laptop_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsesLaptop); i {
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
		file_laptop_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseRequestLaptop); i {
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
		file_laptop_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestSearchByFilter); i {
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
		file_laptop_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageRequest); i {
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
		file_laptop_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageResponse); i {
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
	file_laptop_service_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*UploadImageRequest_Info)(nil),
		(*UploadImageRequest_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_laptop_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_laptop_service_proto_goTypes,
		DependencyIndexes: file_laptop_service_proto_depIdxs,
		MessageInfos:      file_laptop_service_proto_msgTypes,
	}.Build()
	File_laptop_service_proto = out.File
	file_laptop_service_proto_rawDesc = nil
	file_laptop_service_proto_goTypes = nil
	file_laptop_service_proto_depIdxs = nil
}
