// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: utils.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GenValidatorKeys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys uint64 `protobuf:"varint,1,opt,name=keys,proto3" json:"keys,omitempty"`
}

func (x *GenValidatorKeys) Reset() {
	*x = GenValidatorKeys{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utils_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenValidatorKeys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenValidatorKeys) ProtoMessage() {}

func (x *GenValidatorKeys) ProtoReflect() protoreflect.Message {
	mi := &file_utils_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenValidatorKeys.ProtoReflect.Descriptor instead.
func (*GenValidatorKeys) Descriptor() ([]byte, []int) {
	return file_utils_proto_rawDescGZIP(), []int{0}
}

func (x *GenValidatorKeys) GetKeys() uint64 {
	if x != nil {
		return x.Keys
	}
	return 0
}

var File_utils_proto protoreflect.FileDescriptor

var file_utils_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x6e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x32, 0xb9, 0x03, 0x0a, 0x05, 0x55, 0x74, 0x69, 0x6c, 0x73, 0x12, 0x42, 0x0a, 0x0d, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x12, 0x06, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12,
	0x3d, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x12,
	0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x75, 0x74, 0x69, 0x6c,
	0x73, 0x2f, 0x73, 0x74, 0x6f, 0x70, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x12, 0x52,
	0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4b, 0x65,
	0x79, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x4b, 0x65, 0x79, 0x73, 0x1a, 0x09, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x73, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x6b, 0x65, 0x79, 0x3a,
	0x01, 0x2a, 0x12, 0x44, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x61, 0x77, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x08, 0x2e, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22,
	0x14, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x72, 0x61,
	0x77, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x01, 0x2a, 0x12, 0x4d, 0x0a, 0x14, 0x44, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x52, 0x61, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x08, 0x2e, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x03, 0x2e, 0x54, 0x78, 0x22,
	0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f,
	0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x77, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x44, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x52, 0x61, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x08, 0x2e, 0x52, 0x61, 0x77, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x06, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x64, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x72, 0x61, 0x77, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x3a, 0x01, 0x2a, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_utils_proto_rawDescOnce sync.Once
	file_utils_proto_rawDescData = file_utils_proto_rawDesc
)

func file_utils_proto_rawDescGZIP() []byte {
	file_utils_proto_rawDescOnce.Do(func() {
		file_utils_proto_rawDescData = protoimpl.X.CompressGZIP(file_utils_proto_rawDescData)
	})
	return file_utils_proto_rawDescData
}

var file_utils_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_utils_proto_goTypes = []interface{}{
	(*GenValidatorKeys)(nil), // 0: GenValidatorKeys
	(*Empty)(nil),            // 1: Empty
	(*RawData)(nil),          // 2: RawData
	(*Success)(nil),          // 3: Success
	(*KeyPairs)(nil),         // 4: KeyPairs
	(*Tx)(nil),               // 5: Tx
	(*Block)(nil),            // 6: Block
}
var file_utils_proto_depIdxs = []int32{
	1, // 0: Utils.StartProposer:input_type -> Empty
	1, // 1: Utils.StopProposer:input_type -> Empty
	0, // 2: Utils.GenValidatorKey:input_type -> GenValidatorKeys
	2, // 3: Utils.SubmitRawData:input_type -> RawData
	2, // 4: Utils.DecodeRawTransaction:input_type -> RawData
	2, // 5: Utils.DecodeRawBlock:input_type -> RawData
	3, // 6: Utils.StartProposer:output_type -> Success
	3, // 7: Utils.StopProposer:output_type -> Success
	4, // 8: Utils.GenValidatorKey:output_type -> KeyPairs
	3, // 9: Utils.SubmitRawData:output_type -> Success
	5, // 10: Utils.DecodeRawTransaction:output_type -> Tx
	6, // 11: Utils.DecodeRawBlock:output_type -> Block
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_utils_proto_init() }
func file_utils_proto_init() {
	if File_utils_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_utils_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenValidatorKeys); i {
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
			RawDescriptor: file_utils_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_utils_proto_goTypes,
		DependencyIndexes: file_utils_proto_depIdxs,
		MessageInfos:      file_utils_proto_msgTypes,
	}.Build()
	File_utils_proto = out.File
	file_utils_proto_rawDesc = nil
	file_utils_proto_goTypes = nil
	file_utils_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UtilsClient is the client API for Utils service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UtilsClient interface {
	//*
	//Method: StartProposer
	//Input: message Empty
	//Response: message Success
	//Description: Loads the Keystore into the Proposer..
	StartProposer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Success, error)
	//*
	//Method: StopProposer
	//Input: message Password
	//Response: message Success
	//Description: Removes Keystore information and stops the block proposer.
	StopProposer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Success, error)
	//*
	//Method: GenValidatorKey
	//Input: message GenValidatorKeys
	//Response: message KeyPairs
	//Description: Returns private keys generated for validators start.
	GenValidatorKey(ctx context.Context, in *GenValidatorKeys, opts ...grpc.CallOption) (*KeyPairs, error)
	//*
	//Method: SubmitRawData
	//Input: message RawData
	//Response: message Success
	//Description: Broadcast a raw elements of different transactions.
	SubmitRawData(ctx context.Context, in *RawData, opts ...grpc.CallOption) (*Success, error)
	//*
	//Method: DecodeRawTransaction
	//Input: message RawData
	//Response: message Tx
	//Description: Returns a raw transaction on human readable format.
	DecodeRawTransaction(ctx context.Context, in *RawData, opts ...grpc.CallOption) (*Tx, error)
	//*
	//Method: DecodeRawBlock
	//Input: message RawData
	//Response: message Block
	//Description: Returns a raw block on human readable format.
	DecodeRawBlock(ctx context.Context, in *RawData, opts ...grpc.CallOption) (*Block, error)
}

type utilsClient struct {
	cc grpc.ClientConnInterface
}

func NewUtilsClient(cc grpc.ClientConnInterface) UtilsClient {
	return &utilsClient{cc}
}

func (c *utilsClient) StartProposer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/Utils/StartProposer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilsClient) StopProposer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/Utils/StopProposer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilsClient) GenValidatorKey(ctx context.Context, in *GenValidatorKeys, opts ...grpc.CallOption) (*KeyPairs, error) {
	out := new(KeyPairs)
	err := c.cc.Invoke(ctx, "/Utils/GenValidatorKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilsClient) SubmitRawData(ctx context.Context, in *RawData, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/Utils/SubmitRawData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilsClient) DecodeRawTransaction(ctx context.Context, in *RawData, opts ...grpc.CallOption) (*Tx, error) {
	out := new(Tx)
	err := c.cc.Invoke(ctx, "/Utils/DecodeRawTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilsClient) DecodeRawBlock(ctx context.Context, in *RawData, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/Utils/DecodeRawBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UtilsServer is the server API for Utils service.
type UtilsServer interface {
	//*
	//Method: StartProposer
	//Input: message Empty
	//Response: message Success
	//Description: Loads the Keystore into the Proposer..
	StartProposer(context.Context, *Empty) (*Success, error)
	//*
	//Method: StopProposer
	//Input: message Password
	//Response: message Success
	//Description: Removes Keystore information and stops the block proposer.
	StopProposer(context.Context, *Empty) (*Success, error)
	//*
	//Method: GenValidatorKey
	//Input: message GenValidatorKeys
	//Response: message KeyPairs
	//Description: Returns private keys generated for validators start.
	GenValidatorKey(context.Context, *GenValidatorKeys) (*KeyPairs, error)
	//*
	//Method: SubmitRawData
	//Input: message RawData
	//Response: message Success
	//Description: Broadcast a raw elements of different transactions.
	SubmitRawData(context.Context, *RawData) (*Success, error)
	//*
	//Method: DecodeRawTransaction
	//Input: message RawData
	//Response: message Tx
	//Description: Returns a raw transaction on human readable format.
	DecodeRawTransaction(context.Context, *RawData) (*Tx, error)
	//*
	//Method: DecodeRawBlock
	//Input: message RawData
	//Response: message Block
	//Description: Returns a raw block on human readable format.
	DecodeRawBlock(context.Context, *RawData) (*Block, error)
}

// UnimplementedUtilsServer can be embedded to have forward compatible implementations.
type UnimplementedUtilsServer struct {
}

func (*UnimplementedUtilsServer) StartProposer(context.Context, *Empty) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartProposer not implemented")
}
func (*UnimplementedUtilsServer) StopProposer(context.Context, *Empty) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopProposer not implemented")
}
func (*UnimplementedUtilsServer) GenValidatorKey(context.Context, *GenValidatorKeys) (*KeyPairs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenValidatorKey not implemented")
}
func (*UnimplementedUtilsServer) SubmitRawData(context.Context, *RawData) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitRawData not implemented")
}
func (*UnimplementedUtilsServer) DecodeRawTransaction(context.Context, *RawData) (*Tx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeRawTransaction not implemented")
}
func (*UnimplementedUtilsServer) DecodeRawBlock(context.Context, *RawData) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeRawBlock not implemented")
}

func RegisterUtilsServer(s *grpc.Server, srv UtilsServer) {
	s.RegisterService(&_Utils_serviceDesc, srv)
}

func _Utils_StartProposer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilsServer).StartProposer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Utils/StartProposer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilsServer).StartProposer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Utils_StopProposer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilsServer).StopProposer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Utils/StopProposer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilsServer).StopProposer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Utils_GenValidatorKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenValidatorKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilsServer).GenValidatorKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Utils/GenValidatorKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilsServer).GenValidatorKey(ctx, req.(*GenValidatorKeys))
	}
	return interceptor(ctx, in, info, handler)
}

func _Utils_SubmitRawData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilsServer).SubmitRawData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Utils/SubmitRawData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilsServer).SubmitRawData(ctx, req.(*RawData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Utils_DecodeRawTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilsServer).DecodeRawTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Utils/DecodeRawTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilsServer).DecodeRawTransaction(ctx, req.(*RawData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Utils_DecodeRawBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilsServer).DecodeRawBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Utils/DecodeRawBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilsServer).DecodeRawBlock(ctx, req.(*RawData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Utils_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Utils",
	HandlerType: (*UtilsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartProposer",
			Handler:    _Utils_StartProposer_Handler,
		},
		{
			MethodName: "StopProposer",
			Handler:    _Utils_StopProposer_Handler,
		},
		{
			MethodName: "GenValidatorKey",
			Handler:    _Utils_GenValidatorKey_Handler,
		},
		{
			MethodName: "SubmitRawData",
			Handler:    _Utils_SubmitRawData_Handler,
		},
		{
			MethodName: "DecodeRawTransaction",
			Handler:    _Utils_DecodeRawTransaction_Handler,
		},
		{
			MethodName: "DecodeRawBlock",
			Handler:    _Utils_DecodeRawBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "utils.proto",
}
