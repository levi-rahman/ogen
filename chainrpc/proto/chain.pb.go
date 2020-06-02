// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: chainrpc/proto/chain.proto

package proto

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

type ChainInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash   string `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	BlockHeight uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Validators  uint64 `protobuf:"varint,3,opt,name=validators,proto3" json:"validators,omitempty"`
}

func (x *ChainInfo) Reset() {
	*x = ChainInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_chain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainInfo) ProtoMessage() {}

func (x *ChainInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_chain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainInfo.ProtoReflect.Descriptor instead.
func (*ChainInfo) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_chain_proto_rawDescGZIP(), []int{0}
}

func (x *ChainInfo) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *ChainInfo) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *ChainInfo) GetValidators() uint64 {
	if x != nil {
		return x.Validators
	}
	return 0
}

type BlockSubscribe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludeEpochs bool `protobuf:"varint,1,opt,name=include_epochs,json=includeEpochs,proto3" json:"include_epochs,omitempty"`
}

func (x *BlockSubscribe) Reset() {
	*x = BlockSubscribe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_chain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockSubscribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockSubscribe) ProtoMessage() {}

func (x *BlockSubscribe) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_chain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockSubscribe.ProtoReflect.Descriptor instead.
func (*BlockSubscribe) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_chain_proto_rawDescGZIP(), []int{1}
}

func (x *BlockSubscribe) GetIncludeEpochs() bool {
	if x != nil {
		return x.IncludeEpochs
	}
	return false
}

type BlockWithEpoch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawBlock         string `protobuf:"bytes,1,opt,name=raw_block,json=rawBlock,proto3" json:"raw_block,omitempty"`
	EpochTransitions string `protobuf:"bytes,2,opt,name=epoch_transitions,json=epochTransitions,proto3" json:"epoch_transitions,omitempty"`
}

func (x *BlockWithEpoch) Reset() {
	*x = BlockWithEpoch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_chain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockWithEpoch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockWithEpoch) ProtoMessage() {}

func (x *BlockWithEpoch) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_chain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockWithEpoch.ProtoReflect.Descriptor instead.
func (*BlockWithEpoch) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_chain_proto_rawDescGZIP(), []int{2}
}

func (x *BlockWithEpoch) GetRawBlock() string {
	if x != nil {
		return x.RawBlock
	}
	return ""
}

func (x *BlockWithEpoch) GetEpochTransitions() string {
	if x != nil {
		return x.EpochTransitions
	}
	return ""
}

type AccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Balance string   `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Nonce   string   `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Txs     []string `protobuf:"bytes,4,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *AccountInfo) Reset() {
	*x = AccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_chain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfo) ProtoMessage() {}

func (x *AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_chain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInfo.ProtoReflect.Descriptor instead.
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_chain_proto_rawDescGZIP(), []int{3}
}

func (x *AccountInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *AccountInfo) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *AccountInfo) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *AccountInfo) GetTxs() []string {
	if x != nil {
		return x.Txs
	}
	return nil
}

var File_chainrpc_proto_chain_proto protoreflect.FileDescriptor

var file_chainrpc_proto_chain_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x09, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x37, 0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x45, 0x70, 0x6f, 0x63, 0x68,
	0x73, 0x22, 0x5a, 0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x77, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x61, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x2b, 0x0a, 0x11, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x70, 0x6f,
	0x63, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x69, 0x0a,
	0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x03, 0x74, 0x78, 0x73, 0x32, 0xbd, 0x02, 0x0a, 0x05, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x08,
	0x2e, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x1b, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x06, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x07, 0x2e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x1a, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x08, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x00, 0x12, 0x1b, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68,
	0x1a, 0x08, 0x2e, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x30, 0x01, 0x12, 0x37,
	0x0a, 0x0f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x12, 0x0f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x1a, 0x0f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x08, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x75, 0x73, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chainrpc_proto_chain_proto_rawDescOnce sync.Once
	file_chainrpc_proto_chain_proto_rawDescData = file_chainrpc_proto_chain_proto_rawDesc
)

func file_chainrpc_proto_chain_proto_rawDescGZIP() []byte {
	file_chainrpc_proto_chain_proto_rawDescOnce.Do(func() {
		file_chainrpc_proto_chain_proto_rawDescData = protoimpl.X.CompressGZIP(file_chainrpc_proto_chain_proto_rawDescData)
	})
	return file_chainrpc_proto_chain_proto_rawDescData
}

var file_chainrpc_proto_chain_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chainrpc_proto_chain_proto_goTypes = []interface{}{
	(*ChainInfo)(nil),      // 0: ChainInfo
	(*BlockSubscribe)(nil), // 1: BlockSubscribe
	(*BlockWithEpoch)(nil), // 2: BlockWithEpoch
	(*AccountInfo)(nil),    // 3: AccountInfo
	(*Empty)(nil),          // 4: Empty
	(*Hash)(nil),           // 5: Hash
	(*Height)(nil),         // 6: Height
	(*Account)(nil),        // 7: Account
	(*RawData)(nil),        // 8: RawData
	(*Block)(nil),          // 9: Block
}
var file_chainrpc_proto_chain_proto_depIdxs = []int32{
	4, // 0: Chain.GetChainInfo:input_type -> Empty
	5, // 1: Chain.GetRawBlock:input_type -> Hash
	5, // 2: Chain.GetBlock:input_type -> Hash
	6, // 3: Chain.GetBlockHash:input_type -> Height
	7, // 4: Chain.GetAccount:input_type -> Account
	5, // 5: Chain.Sync:input_type -> Hash
	1, // 6: Chain.SubscribeBlocks:input_type -> BlockSubscribe
	7, // 7: Chain.SubscribeAccounts:input_type -> Account
	0, // 8: Chain.GetChainInfo:output_type -> ChainInfo
	8, // 9: Chain.GetRawBlock:output_type -> RawData
	9, // 10: Chain.GetBlock:output_type -> Block
	5, // 11: Chain.GetBlockHash:output_type -> Hash
	3, // 12: Chain.GetAccount:output_type -> AccountInfo
	8, // 13: Chain.Sync:output_type -> RawData
	2, // 14: Chain.SubscribeBlocks:output_type -> BlockWithEpoch
	3, // 15: Chain.SubscribeAccounts:output_type -> AccountInfo
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chainrpc_proto_chain_proto_init() }
func file_chainrpc_proto_chain_proto_init() {
	if File_chainrpc_proto_chain_proto != nil {
		return
	}
	file_chainrpc_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chainrpc_proto_chain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainInfo); i {
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
		file_chainrpc_proto_chain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockSubscribe); i {
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
		file_chainrpc_proto_chain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockWithEpoch); i {
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
		file_chainrpc_proto_chain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInfo); i {
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
			RawDescriptor: file_chainrpc_proto_chain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chainrpc_proto_chain_proto_goTypes,
		DependencyIndexes: file_chainrpc_proto_chain_proto_depIdxs,
		MessageInfos:      file_chainrpc_proto_chain_proto_msgTypes,
	}.Build()
	File_chainrpc_proto_chain_proto = out.File
	file_chainrpc_proto_chain_proto_rawDesc = nil
	file_chainrpc_proto_chain_proto_goTypes = nil
	file_chainrpc_proto_chain_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChainClient is the client API for Chain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChainClient interface {
	GetChainInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChainInfo, error)
	GetRawBlock(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*RawData, error)
	GetBlock(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Block, error)
	GetBlockHash(ctx context.Context, in *Height, opts ...grpc.CallOption) (*Hash, error)
	GetAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountInfo, error)
	Sync(ctx context.Context, in *Hash, opts ...grpc.CallOption) (Chain_SyncClient, error)
	SubscribeBlocks(ctx context.Context, in *BlockSubscribe, opts ...grpc.CallOption) (Chain_SubscribeBlocksClient, error)
	SubscribeAccounts(ctx context.Context, in *Account, opts ...grpc.CallOption) (Chain_SubscribeAccountsClient, error)
}

type chainClient struct {
	cc grpc.ClientConnInterface
}

func NewChainClient(cc grpc.ClientConnInterface) ChainClient {
	return &chainClient{cc}
}

func (c *chainClient) GetChainInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChainInfo, error) {
	out := new(ChainInfo)
	err := c.cc.Invoke(ctx, "/Chain/GetChainInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetRawBlock(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*RawData, error) {
	out := new(RawData)
	err := c.cc.Invoke(ctx, "/Chain/GetRawBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetBlock(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/Chain/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetBlockHash(ctx context.Context, in *Height, opts ...grpc.CallOption) (*Hash, error) {
	out := new(Hash)
	err := c.cc.Invoke(ctx, "/Chain/GetBlockHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountInfo, error) {
	out := new(AccountInfo)
	err := c.cc.Invoke(ctx, "/Chain/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) Sync(ctx context.Context, in *Hash, opts ...grpc.CallOption) (Chain_SyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chain_serviceDesc.Streams[0], "/Chain/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &chainSyncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chain_SyncClient interface {
	Recv() (*RawData, error)
	grpc.ClientStream
}

type chainSyncClient struct {
	grpc.ClientStream
}

func (x *chainSyncClient) Recv() (*RawData, error) {
	m := new(RawData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chainClient) SubscribeBlocks(ctx context.Context, in *BlockSubscribe, opts ...grpc.CallOption) (Chain_SubscribeBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chain_serviceDesc.Streams[1], "/Chain/SubscribeBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &chainSubscribeBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chain_SubscribeBlocksClient interface {
	Recv() (*BlockWithEpoch, error)
	grpc.ClientStream
}

type chainSubscribeBlocksClient struct {
	grpc.ClientStream
}

func (x *chainSubscribeBlocksClient) Recv() (*BlockWithEpoch, error) {
	m := new(BlockWithEpoch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chainClient) SubscribeAccounts(ctx context.Context, in *Account, opts ...grpc.CallOption) (Chain_SubscribeAccountsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chain_serviceDesc.Streams[2], "/Chain/SubscribeAccounts", opts...)
	if err != nil {
		return nil, err
	}
	x := &chainSubscribeAccountsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chain_SubscribeAccountsClient interface {
	Recv() (*AccountInfo, error)
	grpc.ClientStream
}

type chainSubscribeAccountsClient struct {
	grpc.ClientStream
}

func (x *chainSubscribeAccountsClient) Recv() (*AccountInfo, error) {
	m := new(AccountInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChainServer is the server API for Chain service.
type ChainServer interface {
	GetChainInfo(context.Context, *Empty) (*ChainInfo, error)
	GetRawBlock(context.Context, *Hash) (*RawData, error)
	GetBlock(context.Context, *Hash) (*Block, error)
	GetBlockHash(context.Context, *Height) (*Hash, error)
	GetAccount(context.Context, *Account) (*AccountInfo, error)
	Sync(*Hash, Chain_SyncServer) error
	SubscribeBlocks(*BlockSubscribe, Chain_SubscribeBlocksServer) error
	SubscribeAccounts(*Account, Chain_SubscribeAccountsServer) error
}

// UnimplementedChainServer can be embedded to have forward compatible implementations.
type UnimplementedChainServer struct {
}

func (*UnimplementedChainServer) GetChainInfo(context.Context, *Empty) (*ChainInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChainInfo not implemented")
}
func (*UnimplementedChainServer) GetRawBlock(context.Context, *Hash) (*RawData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRawBlock not implemented")
}
func (*UnimplementedChainServer) GetBlock(context.Context, *Hash) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (*UnimplementedChainServer) GetBlockHash(context.Context, *Height) (*Hash, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHash not implemented")
}
func (*UnimplementedChainServer) GetAccount(context.Context, *Account) (*AccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (*UnimplementedChainServer) Sync(*Hash, Chain_SyncServer) error {
	return status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (*UnimplementedChainServer) SubscribeBlocks(*BlockSubscribe, Chain_SubscribeBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeBlocks not implemented")
}
func (*UnimplementedChainServer) SubscribeAccounts(*Account, Chain_SubscribeAccountsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeAccounts not implemented")
}

func RegisterChainServer(s *grpc.Server, srv ChainServer) {
	s.RegisterService(&_Chain_serviceDesc, srv)
}

func _Chain_GetChainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).GetChainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chain/GetChainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).GetChainInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chain_GetRawBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).GetRawBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chain/GetRawBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).GetRawBlock(ctx, req.(*Hash))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chain_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chain/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).GetBlock(ctx, req.(*Hash))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chain_GetBlockHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Height)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).GetBlockHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chain/GetBlockHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).GetBlockHash(ctx, req.(*Height))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chain_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chain/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).GetAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chain_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Hash)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainServer).Sync(m, &chainSyncServer{stream})
}

type Chain_SyncServer interface {
	Send(*RawData) error
	grpc.ServerStream
}

type chainSyncServer struct {
	grpc.ServerStream
}

func (x *chainSyncServer) Send(m *RawData) error {
	return x.ServerStream.SendMsg(m)
}

func _Chain_SubscribeBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockSubscribe)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainServer).SubscribeBlocks(m, &chainSubscribeBlocksServer{stream})
}

type Chain_SubscribeBlocksServer interface {
	Send(*BlockWithEpoch) error
	grpc.ServerStream
}

type chainSubscribeBlocksServer struct {
	grpc.ServerStream
}

func (x *chainSubscribeBlocksServer) Send(m *BlockWithEpoch) error {
	return x.ServerStream.SendMsg(m)
}

func _Chain_SubscribeAccounts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Account)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainServer).SubscribeAccounts(m, &chainSubscribeAccountsServer{stream})
}

type Chain_SubscribeAccountsServer interface {
	Send(*AccountInfo) error
	grpc.ServerStream
}

type chainSubscribeAccountsServer struct {
	grpc.ServerStream
}

func (x *chainSubscribeAccountsServer) Send(m *AccountInfo) error {
	return x.ServerStream.SendMsg(m)
}

var _Chain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Chain",
	HandlerType: (*ChainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChainInfo",
			Handler:    _Chain_GetChainInfo_Handler,
		},
		{
			MethodName: "GetRawBlock",
			Handler:    _Chain_GetRawBlock_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _Chain_GetBlock_Handler,
		},
		{
			MethodName: "GetBlockHash",
			Handler:    _Chain_GetBlockHash_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Chain_GetAccount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sync",
			Handler:       _Chain_Sync_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeBlocks",
			Handler:       _Chain_SubscribeBlocks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeAccounts",
			Handler:       _Chain_SubscribeAccounts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chainrpc/proto/chain.proto",
}
