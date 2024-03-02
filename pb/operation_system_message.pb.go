// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: operation_system_message.proto

package pb

import (
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

type Linux_Distro int32

const (
	Linux_UNKNOWN_DISTRO Linux_Distro = 0
	Linux_UBUNTU         Linux_Distro = 1
	Linux_FEDORA         Linux_Distro = 2
	Linux_CENTOS         Linux_Distro = 3
	Linux_DEBIAN         Linux_Distro = 4
	Linux_ARCH           Linux_Distro = 5
	Linux_ALPINE         Linux_Distro = 6
	Linux_Mint           Linux_Distro = 7
)

// Enum value maps for Linux_Distro.
var (
	Linux_Distro_name = map[int32]string{
		0: "UNKNOWN_DISTRO",
		1: "UBUNTU",
		2: "FEDORA",
		3: "CENTOS",
		4: "DEBIAN",
		5: "ARCH",
		6: "ALPINE",
		7: "Mint",
	}
	Linux_Distro_value = map[string]int32{
		"UNKNOWN_DISTRO": 0,
		"UBUNTU":         1,
		"FEDORA":         2,
		"CENTOS":         3,
		"DEBIAN":         4,
		"ARCH":           5,
		"ALPINE":         6,
		"Mint":           7,
	}
)

func (x Linux_Distro) Enum() *Linux_Distro {
	p := new(Linux_Distro)
	*p = x
	return p
}

func (x Linux_Distro) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Linux_Distro) Descriptor() protoreflect.EnumDescriptor {
	return file_operation_system_message_proto_enumTypes[0].Descriptor()
}

func (Linux_Distro) Type() protoreflect.EnumType {
	return &file_operation_system_message_proto_enumTypes[0]
}

func (x Linux_Distro) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Linux_Distro.Descriptor instead.
func (Linux_Distro) EnumDescriptor() ([]byte, []int) {
	return file_operation_system_message_proto_rawDescGZIP(), []int{2, 0}
}

type OS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TypeOs:
	//
	//	*OS_Windows
	//	*OS_Linux
	//	*OS_MacOs
	TypeOs isOS_TypeOs `protobuf_oneof:"type_os"`
}

func (x *OS) Reset() {
	*x = OS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operation_system_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OS) ProtoMessage() {}

func (x *OS) ProtoReflect() protoreflect.Message {
	mi := &file_operation_system_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OS.ProtoReflect.Descriptor instead.
func (*OS) Descriptor() ([]byte, []int) {
	return file_operation_system_message_proto_rawDescGZIP(), []int{0}
}

func (m *OS) GetTypeOs() isOS_TypeOs {
	if m != nil {
		return m.TypeOs
	}
	return nil
}

func (x *OS) GetWindows() *Windows {
	if x, ok := x.GetTypeOs().(*OS_Windows); ok {
		return x.Windows
	}
	return nil
}

func (x *OS) GetLinux() *Linux {
	if x, ok := x.GetTypeOs().(*OS_Linux); ok {
		return x.Linux
	}
	return nil
}

func (x *OS) GetMacOs() *MAC {
	if x, ok := x.GetTypeOs().(*OS_MacOs); ok {
		return x.MacOs
	}
	return nil
}

type isOS_TypeOs interface {
	isOS_TypeOs()
}

type OS_Windows struct {
	Windows *Windows `protobuf:"bytes,1,opt,name=windows,proto3,oneof"`
}

type OS_Linux struct {
	Linux *Linux `protobuf:"bytes,2,opt,name=linux,proto3,oneof"`
}

type OS_MacOs struct {
	MacOs *MAC `protobuf:"bytes,3,opt,name=mac_os,json=macOs,proto3,oneof"`
}

func (*OS_Windows) isOS_TypeOs() {}

func (*OS_Linux) isOS_TypeOs() {}

func (*OS_MacOs) isOS_TypeOs() {}

type Windows struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Edition string `protobuf:"bytes,2,opt,name=edition,proto3" json:"edition,omitempty"`
}

func (x *Windows) Reset() {
	*x = Windows{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operation_system_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Windows) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Windows) ProtoMessage() {}

func (x *Windows) ProtoReflect() protoreflect.Message {
	mi := &file_operation_system_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Windows.ProtoReflect.Descriptor instead.
func (*Windows) Descriptor() ([]byte, []int) {
	return file_operation_system_message_proto_rawDescGZIP(), []int{1}
}

func (x *Windows) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Windows) GetEdition() string {
	if x != nil {
		return x.Edition
	}
	return ""
}

type Linux struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Distribution  Linux_Distro `protobuf:"varint,1,opt,name=distribution,proto3,enum=Linux_Distro" json:"distribution,omitempty"`
	Version       string       `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	KernelVersion string       `protobuf:"bytes,3,opt,name=kernel_version,json=kernelVersion,proto3" json:"kernel_version,omitempty"`
}

func (x *Linux) Reset() {
	*x = Linux{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operation_system_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Linux) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Linux) ProtoMessage() {}

func (x *Linux) ProtoReflect() protoreflect.Message {
	mi := &file_operation_system_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Linux.ProtoReflect.Descriptor instead.
func (*Linux) Descriptor() ([]byte, []int) {
	return file_operation_system_message_proto_rawDescGZIP(), []int{2}
}

func (x *Linux) GetDistribution() Linux_Distro {
	if x != nil {
		return x.Distribution
	}
	return Linux_UNKNOWN_DISTRO
}

func (x *Linux) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Linux) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

type MAC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version       string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Build         string `protobuf:"bytes,2,opt,name=build,proto3" json:"build,omitempty"`
	KernelVersion string `protobuf:"bytes,3,opt,name=kernel_version,json=kernelVersion,proto3" json:"kernel_version,omitempty"`
}

func (x *MAC) Reset() {
	*x = MAC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operation_system_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MAC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MAC) ProtoMessage() {}

func (x *MAC) ProtoReflect() protoreflect.Message {
	mi := &file_operation_system_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MAC.ProtoReflect.Descriptor instead.
func (*MAC) Descriptor() ([]byte, []int) {
	return file_operation_system_message_proto_rawDescGZIP(), []int{3}
}

func (x *MAC) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *MAC) GetBuild() string {
	if x != nil {
		return x.Build
	}
	return ""
}

func (x *MAC) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

var File_operation_system_message_proto protoreflect.FileDescriptor

var file_operation_system_message_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x74, 0x0a, 0x02, 0x4f, 0x53, 0x12, 0x24, 0x0a, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x73, 0x48, 0x00, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x12, 0x1e, 0x0a, 0x05,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4c, 0x69,
	0x6e, 0x75, 0x78, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x12, 0x1d, 0x0a, 0x06,
	0x6d, 0x61, 0x63, 0x5f, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4d,
	0x41, 0x43, 0x48, 0x00, 0x52, 0x05, 0x6d, 0x61, 0x63, 0x4f, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x6f, 0x73, 0x22, 0x3d, 0x0a, 0x07, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe9, 0x01, 0x0a, 0x05, 0x4c, 0x69, 0x6e, 0x75, 0x78, 0x12,
	0x31, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x4c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x6f, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e,
	0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x6c, 0x0a, 0x06, 0x44, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x12, 0x12, 0x0a,
	0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x44, 0x49, 0x53, 0x54, 0x52, 0x4f, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x42, 0x55, 0x4e, 0x54, 0x55, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x45, 0x44, 0x4f, 0x52, 0x41, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x45, 0x4e,
	0x54, 0x4f, 0x53, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x42, 0x49, 0x41, 0x4e, 0x10,
	0x04, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x52, 0x43, 0x48, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x41,
	0x4c, 0x50, 0x49, 0x4e, 0x45, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x69, 0x6e, 0x74, 0x10,
	0x07, 0x22, 0x5c, 0x0a, 0x03, 0x4d, 0x41, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6b, 0x65, 0x72, 0x6e,
	0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x05, 0x5a, 0x03, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operation_system_message_proto_rawDescOnce sync.Once
	file_operation_system_message_proto_rawDescData = file_operation_system_message_proto_rawDesc
)

func file_operation_system_message_proto_rawDescGZIP() []byte {
	file_operation_system_message_proto_rawDescOnce.Do(func() {
		file_operation_system_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_operation_system_message_proto_rawDescData)
	})
	return file_operation_system_message_proto_rawDescData
}

var file_operation_system_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_operation_system_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_operation_system_message_proto_goTypes = []interface{}{
	(Linux_Distro)(0), // 0: Linux.Distro
	(*OS)(nil),        // 1: OS
	(*Windows)(nil),   // 2: Windows
	(*Linux)(nil),     // 3: Linux
	(*MAC)(nil),       // 4: MAC
}
var file_operation_system_message_proto_depIdxs = []int32{
	2, // 0: OS.windows:type_name -> Windows
	3, // 1: OS.linux:type_name -> Linux
	4, // 2: OS.mac_os:type_name -> MAC
	0, // 3: Linux.distribution:type_name -> Linux.Distro
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_operation_system_message_proto_init() }
func file_operation_system_message_proto_init() {
	if File_operation_system_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_operation_system_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OS); i {
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
		file_operation_system_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Windows); i {
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
		file_operation_system_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Linux); i {
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
		file_operation_system_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MAC); i {
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
	file_operation_system_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OS_Windows)(nil),
		(*OS_Linux)(nil),
		(*OS_MacOs)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_operation_system_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_operation_system_message_proto_goTypes,
		DependencyIndexes: file_operation_system_message_proto_depIdxs,
		EnumInfos:         file_operation_system_message_proto_enumTypes,
		MessageInfos:      file_operation_system_message_proto_msgTypes,
	}.Build()
	File_operation_system_message_proto = out.File
	file_operation_system_message_proto_rawDesc = nil
	file_operation_system_message_proto_goTypes = nil
	file_operation_system_message_proto_depIdxs = nil
}
