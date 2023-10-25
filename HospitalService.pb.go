// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: HospitalService/HospitalService.proto

package proto

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

type MessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *MessageReq) Reset() {
	*x = MessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HospitalService_HospitalService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageReq) ProtoMessage() {}

func (x *MessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_HospitalService_HospitalService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageReq.ProtoReflect.Descriptor instead.
func (*MessageReq) Descriptor() ([]byte, []int) {
	return file_HospitalService_HospitalService_proto_rawDescGZIP(), []int{0}
}

func (x *MessageReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *MessageReq) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type MessageRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res string `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *MessageRes) Reset() {
	*x = MessageRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HospitalService_HospitalService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRes) ProtoMessage() {}

func (x *MessageRes) ProtoReflect() protoreflect.Message {
	mi := &file_HospitalService_HospitalService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRes.ProtoReflect.Descriptor instead.
func (*MessageRes) Descriptor() ([]byte, []int) {
	return file_HospitalService_HospitalService_proto_rawDescGZIP(), []int{1}
}

func (x *MessageRes) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

type GetClientsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"` // Could be used to send additional request parameters.
}

func (x *GetClientsReq) Reset() {
	*x = GetClientsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HospitalService_HospitalService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsReq) ProtoMessage() {}

func (x *GetClientsReq) ProtoReflect() protoreflect.Message {
	mi := &file_HospitalService_HospitalService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsReq.ProtoReflect.Descriptor instead.
func (*GetClientsReq) Descriptor() ([]byte, []int) {
	return file_HospitalService_HospitalService_proto_rawDescGZIP(), []int{2}
}

func (x *GetClientsReq) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type ClientListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []string `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *ClientListRes) Reset() {
	*x = ClientListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HospitalService_HospitalService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientListRes) ProtoMessage() {}

func (x *ClientListRes) ProtoReflect() protoreflect.Message {
	mi := &file_HospitalService_HospitalService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientListRes.ProtoReflect.Descriptor instead.
func (*ClientListRes) Descriptor() ([]byte, []int) {
	return file_HospitalService_HospitalService_proto_rawDescGZIP(), []int{3}
}

func (x *ClientListRes) GetClients() []string {
	if x != nil {
		return x.Clients
	}
	return nil
}

type ClientIDsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid []string `protobuf:"bytes,1,rep,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *ClientIDsList) Reset() {
	*x = ClientIDsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HospitalService_HospitalService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientIDsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientIDsList) ProtoMessage() {}

func (x *ClientIDsList) ProtoReflect() protoreflect.Message {
	mi := &file_HospitalService_HospitalService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientIDsList.ProtoReflect.Descriptor instead.
func (*ClientIDsList) Descriptor() ([]byte, []int) {
	return file_HospitalService_HospitalService_proto_rawDescGZIP(), []int{4}
}

func (x *ClientIDsList) GetUuid() []string {
	if x != nil {
		return x.Uuid
	}
	return nil
}

type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HospitalService_HospitalService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_HospitalService_HospitalService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_HospitalService_HospitalService_proto_rawDescGZIP(), []int{5}
}

var File_HospitalService_HospitalService_proto protoreflect.FileDescriptor

var file_HospitalService_HospitalService_proto_rawDesc = []byte{
	0x0a, 0x25, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61,
	0x6c, 0x22, 0x3a, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x1e, 0x0a,
	0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x73, 0x22, 0x29, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x23, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x32, 0x98, 0x02, 0x0a, 0x0f, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e,
	0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x68,
	0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x12, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x12, 0x14,
	0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42,
	0x17, 0x5a, 0x15, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HospitalService_HospitalService_proto_rawDescOnce sync.Once
	file_HospitalService_HospitalService_proto_rawDescData = file_HospitalService_HospitalService_proto_rawDesc
)

func file_HospitalService_HospitalService_proto_rawDescGZIP() []byte {
	file_HospitalService_HospitalService_proto_rawDescOnce.Do(func() {
		file_HospitalService_HospitalService_proto_rawDescData = protoimpl.X.CompressGZIP(file_HospitalService_HospitalService_proto_rawDescData)
	})
	return file_HospitalService_HospitalService_proto_rawDescData
}

var file_HospitalService_HospitalService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_HospitalService_HospitalService_proto_goTypes = []interface{}{
	(*MessageReq)(nil),    // 0: hospital.MessageReq
	(*MessageRes)(nil),    // 1: hospital.MessageRes
	(*GetClientsReq)(nil), // 2: hospital.GetClientsReq
	(*ClientListRes)(nil), // 3: hospital.ClientListRes
	(*ClientIDsList)(nil), // 4: hospital.ClientIDsList
	(*EmptyReq)(nil),      // 5: hospital.EmptyReq
}
var file_HospitalService_HospitalService_proto_depIdxs = []int32{
	0, // 0: hospital.HospitalService.SendMessage:input_type -> hospital.MessageReq
	2, // 1: hospital.HospitalService.GetClients:input_type -> hospital.GetClientsReq
	5, // 2: hospital.HospitalService.GetActiveConnections:input_type -> hospital.EmptyReq
	0, // 3: hospital.HospitalService.GetClientIDs:input_type -> hospital.MessageReq
	1, // 4: hospital.HospitalService.SendMessage:output_type -> hospital.MessageRes
	3, // 5: hospital.HospitalService.GetClients:output_type -> hospital.ClientListRes
	3, // 6: hospital.HospitalService.GetActiveConnections:output_type -> hospital.ClientListRes
	3, // 7: hospital.HospitalService.GetClientIDs:output_type -> hospital.ClientListRes
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HospitalService_HospitalService_proto_init() }
func file_HospitalService_HospitalService_proto_init() {
	if File_HospitalService_HospitalService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HospitalService_HospitalService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageReq); i {
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
		file_HospitalService_HospitalService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRes); i {
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
		file_HospitalService_HospitalService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsReq); i {
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
		file_HospitalService_HospitalService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientListRes); i {
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
		file_HospitalService_HospitalService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientIDsList); i {
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
		file_HospitalService_HospitalService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReq); i {
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
			RawDescriptor: file_HospitalService_HospitalService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_HospitalService_HospitalService_proto_goTypes,
		DependencyIndexes: file_HospitalService_HospitalService_proto_depIdxs,
		MessageInfos:      file_HospitalService_HospitalService_proto_msgTypes,
	}.Build()
	File_HospitalService_HospitalService_proto = out.File
	file_HospitalService_HospitalService_proto_rawDesc = nil
	file_HospitalService_HospitalService_proto_goTypes = nil
	file_HospitalService_HospitalService_proto_depIdxs = nil
}
