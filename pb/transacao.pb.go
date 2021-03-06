// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/transacao.proto

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

type QixRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transacao    string `protobuf:"bytes,1,opt,name=transacao,proto3" json:"transacao,omitempty"`
	Complexidade int32  `protobuf:"varint,2,opt,name=complexidade,proto3" json:"complexidade,omitempty"`
}

func (x *QixRequest) Reset() {
	*x = QixRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transacao_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QixRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QixRequest) ProtoMessage() {}

func (x *QixRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transacao_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QixRequest.ProtoReflect.Descriptor instead.
func (*QixRequest) Descriptor() ([]byte, []int) {
	return file_proto_transacao_proto_rawDescGZIP(), []int{0}
}

func (x *QixRequest) GetTransacao() string {
	if x != nil {
		return x.Transacao
	}
	return ""
}

func (x *QixRequest) GetComplexidade() int32 {
	if x != nil {
		return x.Complexidade
	}
	return 0
}

type QixReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servidor string `protobuf:"bytes,1,opt,name=servidor,proto3" json:"servidor,omitempty"`
	Sucesso  bool   `protobuf:"varint,2,opt,name=sucesso,proto3" json:"sucesso,omitempty"`
}

func (x *QixReply) Reset() {
	*x = QixReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transacao_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QixReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QixReply) ProtoMessage() {}

func (x *QixReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transacao_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QixReply.ProtoReflect.Descriptor instead.
func (*QixReply) Descriptor() ([]byte, []int) {
	return file_proto_transacao_proto_rawDescGZIP(), []int{1}
}

func (x *QixReply) GetServidor() string {
	if x != nil {
		return x.Servidor
	}
	return ""
}

func (x *QixReply) GetSucesso() bool {
	if x != nil {
		return x.Sucesso
	}
	return false
}

type LoadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servidor string `protobuf:"bytes,1,opt,name=servidor,proto3" json:"servidor,omitempty"`
	Load     int64  `protobuf:"varint,2,opt,name=load,proto3" json:"load,omitempty"`
}

func (x *LoadInfo) Reset() {
	*x = LoadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transacao_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadInfo) ProtoMessage() {}

func (x *LoadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transacao_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadInfo.ProtoReflect.Descriptor instead.
func (*LoadInfo) Descriptor() ([]byte, []int) {
	return file_proto_transacao_proto_rawDescGZIP(), []int{2}
}

func (x *LoadInfo) GetServidor() string {
	if x != nil {
		return x.Servidor
	}
	return ""
}

func (x *LoadInfo) GetLoad() int64 {
	if x != nil {
		return x.Load
	}
	return 0
}

type LoadRecebido struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sucesso bool `protobuf:"varint,1,opt,name=sucesso,proto3" json:"sucesso,omitempty"`
}

func (x *LoadRecebido) Reset() {
	*x = LoadRecebido{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transacao_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadRecebido) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadRecebido) ProtoMessage() {}

func (x *LoadRecebido) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transacao_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadRecebido.ProtoReflect.Descriptor instead.
func (*LoadRecebido) Descriptor() ([]byte, []int) {
	return file_proto_transacao_proto_rawDescGZIP(), []int{3}
}

func (x *LoadRecebido) GetSucesso() bool {
	if x != nil {
		return x.Sucesso
	}
	return false
}

var File_proto_transacao_proto protoreflect.FileDescriptor

var file_proto_transacao_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x61,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x71, 0x69, 0x78, 0x32, 0x22, 0x4e, 0x0a,
	0x0a, 0x51, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x61, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x61, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x69, 0x64, 0x61, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x69, 0x64, 0x61, 0x64, 0x65, 0x22, 0x40, 0x0a,
	0x08, 0x51, 0x69, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x64, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x65, 0x73, 0x73, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x22,
	0x3a, 0x0a, 0x08, 0x4c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x4c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x63, 0x65, 0x62, 0x69, 0x64, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x32, 0x35, 0x0a, 0x03, 0x51, 0x69, 0x78, 0x12, 0x2e, 0x0a, 0x08,
	0x45, 0x6e, 0x76, 0x69, 0x61, 0x51, 0x69, 0x78, 0x12, 0x10, 0x2e, 0x71, 0x69, 0x78, 0x32, 0x2e,
	0x51, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x71, 0x69, 0x78,
	0x32, 0x2e, 0x51, 0x69, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x39, 0x0a, 0x04,
	0x4c, 0x6f, 0x61, 0x64, 0x12, 0x31, 0x0a, 0x09, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x4c, 0x6f, 0x61,
	0x64, 0x12, 0x0e, 0x2e, 0x71, 0x69, 0x78, 0x32, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x12, 0x2e, 0x71, 0x69, 0x78, 0x32, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x63,
	0x65, 0x62, 0x69, 0x64, 0x6f, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_transacao_proto_rawDescOnce sync.Once
	file_proto_transacao_proto_rawDescData = file_proto_transacao_proto_rawDesc
)

func file_proto_transacao_proto_rawDescGZIP() []byte {
	file_proto_transacao_proto_rawDescOnce.Do(func() {
		file_proto_transacao_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_transacao_proto_rawDescData)
	})
	return file_proto_transacao_proto_rawDescData
}

var file_proto_transacao_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_transacao_proto_goTypes = []interface{}{
	(*QixRequest)(nil),   // 0: qix2.QixRequest
	(*QixReply)(nil),     // 1: qix2.QixReply
	(*LoadInfo)(nil),     // 2: qix2.LoadInfo
	(*LoadRecebido)(nil), // 3: qix2.LoadRecebido
}
var file_proto_transacao_proto_depIdxs = []int32{
	0, // 0: qix2.Qix.EnviaQix:input_type -> qix2.QixRequest
	2, // 1: qix2.Load.EnviaLoad:input_type -> qix2.LoadInfo
	1, // 2: qix2.Qix.EnviaQix:output_type -> qix2.QixReply
	3, // 3: qix2.Load.EnviaLoad:output_type -> qix2.LoadRecebido
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_transacao_proto_init() }
func file_proto_transacao_proto_init() {
	if File_proto_transacao_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_transacao_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QixRequest); i {
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
		file_proto_transacao_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QixReply); i {
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
		file_proto_transacao_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadInfo); i {
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
		file_proto_transacao_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadRecebido); i {
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
			RawDescriptor: file_proto_transacao_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_transacao_proto_goTypes,
		DependencyIndexes: file_proto_transacao_proto_depIdxs,
		MessageInfos:      file_proto_transacao_proto_msgTypes,
	}.Build()
	File_proto_transacao_proto = out.File
	file_proto_transacao_proto_rawDesc = nil
	file_proto_transacao_proto_goTypes = nil
	file_proto_transacao_proto_depIdxs = nil
}
