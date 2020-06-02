// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: chainrpc/proto/common.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_common_proto_rawDescGZIP(), []int{0}
}

type Hash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Hash) Reset() {
	*x = Hash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hash) ProtoMessage() {}

func (x *Hash) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hash.ProtoReflect.Descriptor instead.
func (*Hash) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_common_proto_rawDescGZIP(), []int{1}
}

func (x *Hash) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type Height struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Height) Reset() {
	*x = Height{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Height) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Height) ProtoMessage() {}

func (x *Height) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Height.ProtoReflect.Descriptor instead.
func (*Height) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_common_proto_rawDescGZIP(), []int{2}
}

func (x *Height) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_common_proto_rawDescGZIP(), []int{3}
}

func (x *Success) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Success) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type KeyPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Public  string `protobuf:"bytes,1,opt,name=public,proto3" json:"public,omitempty"`
	Private string `protobuf:"bytes,2,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *KeyPair) Reset() {
	*x = KeyPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyPair) ProtoMessage() {}

func (x *KeyPair) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyPair.ProtoReflect.Descriptor instead.
func (*KeyPair) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_common_proto_rawDescGZIP(), []int{4}
}

func (x *KeyPair) GetPublic() string {
	if x != nil {
		return x.Public
	}
	return ""
}

func (x *KeyPair) GetPrivate() string {
	if x != nil {
		return x.Private
	}
	return ""
}

type RawData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RawData) Reset() {
	*x = RawData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawData) ProtoMessage() {}

func (x *RawData) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawData.ProtoReflect.Descriptor instead.
func (*RawData) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_common_proto_rawDescGZIP(), []int{5}
}

func (x *RawData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type BlockHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version                    int32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Nonce                      int32  `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	TxMerkleRoot               string `protobuf:"bytes,3,opt,name=tx_merkle_root,json=txMerkleRoot,proto3" json:"tx_merkle_root,omitempty"`
	VoteMerkleRoot             string `protobuf:"bytes,4,opt,name=vote_merkle_root,json=voteMerkleRoot,proto3" json:"vote_merkle_root,omitempty"`
	DepositMerkleRoot          string `protobuf:"bytes,5,opt,name=deposit_merkle_root,json=depositMerkleRoot,proto3" json:"deposit_merkle_root,omitempty"`
	ExitMerkleRoot             string `protobuf:"bytes,6,opt,name=exit_merkle_root,json=exitMerkleRoot,proto3" json:"exit_merkle_root,omitempty"`
	VoteSlashingMerkleRoot     string `protobuf:"bytes,7,opt,name=vote_slashing_merkle_root,json=voteSlashingMerkleRoot,proto3" json:"vote_slashing_merkle_root,omitempty"`
	RandaoSlashingMerkleRoot   string `protobuf:"bytes,8,opt,name=randao_slashing_merkle_root,json=randaoSlashingMerkleRoot,proto3" json:"randao_slashing_merkle_root,omitempty"`
	ProposerSlashingMerkleRoot string `protobuf:"bytes,9,opt,name=proposer_slashing_merkle_root,json=proposerSlashingMerkleRoot,proto3" json:"proposer_slashing_merkle_root,omitempty"`
	PrevBlockHash              string `protobuf:"bytes,10,opt,name=prev_block_hash,json=prevBlockHash,proto3" json:"prev_block_hash,omitempty"`
	Timestamp                  int64  `protobuf:"varint,11,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Slot                       uint64 `protobuf:"varint,12,opt,name=slot,proto3" json:"slot,omitempty"`
	StateRoot                  string `protobuf:"bytes,13,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	FeeAddress                 string `protobuf:"bytes,14,opt,name=fee_address,json=feeAddress,proto3" json:"fee_address,omitempty"`
}

func (x *BlockHeader) Reset() {
	*x = BlockHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeader) ProtoMessage() {}

func (x *BlockHeader) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeader.ProtoReflect.Descriptor instead.
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_common_proto_rawDescGZIP(), []int{6}
}

func (x *BlockHeader) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *BlockHeader) GetNonce() int32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *BlockHeader) GetTxMerkleRoot() string {
	if x != nil {
		return x.TxMerkleRoot
	}
	return ""
}

func (x *BlockHeader) GetVoteMerkleRoot() string {
	if x != nil {
		return x.VoteMerkleRoot
	}
	return ""
}

func (x *BlockHeader) GetDepositMerkleRoot() string {
	if x != nil {
		return x.DepositMerkleRoot
	}
	return ""
}

func (x *BlockHeader) GetExitMerkleRoot() string {
	if x != nil {
		return x.ExitMerkleRoot
	}
	return ""
}

func (x *BlockHeader) GetVoteSlashingMerkleRoot() string {
	if x != nil {
		return x.VoteSlashingMerkleRoot
	}
	return ""
}

func (x *BlockHeader) GetRandaoSlashingMerkleRoot() string {
	if x != nil {
		return x.RandaoSlashingMerkleRoot
	}
	return ""
}

func (x *BlockHeader) GetProposerSlashingMerkleRoot() string {
	if x != nil {
		return x.ProposerSlashingMerkleRoot
	}
	return ""
}

func (x *BlockHeader) GetPrevBlockHash() string {
	if x != nil {
		return x.PrevBlockHash
	}
	return ""
}

func (x *BlockHeader) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BlockHeader) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *BlockHeader) GetStateRoot() string {
	if x != nil {
		return x.StateRoot
	}
	return ""
}

func (x *BlockHeader) GetFeeAddress() string {
	if x != nil {
		return x.FeeAddress
	}
	return ""
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash            string       `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Header          *BlockHeader `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Txs             []string     `protobuf:"bytes,3,rep,name=txs,proto3" json:"txs,omitempty"`
	Signature       string       `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	RandaoSignature string       `protobuf:"bytes,5,opt,name=randao_signature,json=randaoSignature,proto3" json:"randao_signature,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_common_proto_rawDescGZIP(), []int{7}
}

func (x *Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Block) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetTxs() []string {
	if x != nil {
		return x.Txs
	}
	return nil
}

func (x *Block) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Block) GetRandaoSignature() string {
	if x != nil {
		return x.RandaoSignature
	}
	return ""
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_common_proto_rawDescGZIP(), []int{8}
}

func (x *Account) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// TODO
type Tx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Tx) Reset() {
	*x = Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_proto_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tx) ProtoMessage() {}

func (x *Tx) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_proto_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tx.ProtoReflect.Descriptor instead.
func (*Tx) Descriptor() ([]byte, []int) {
	return file_chainrpc_proto_common_proto_rawDescGZIP(), []int{9}
}

func (x *Tx) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

var File_chainrpc_proto_common_proto protoreflect.FileDescriptor

var file_chainrpc_proto_common_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1a, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x22, 0x20, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x22, 0x39, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x3b, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x1d, 0x0a, 0x07,
	0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xbe, 0x04, 0x0a, 0x0b,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74,
	0x78, 0x5f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x78, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x76, 0x6f, 0x74,
	0x65, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f,
	0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x65,
	0x78, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x69, 0x74, 0x4d, 0x65, 0x72, 0x6b, 0x6c,
	0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x6c,
	0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f,
	0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x76, 0x6f, 0x74, 0x65, 0x53, 0x6c,
	0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74,
	0x12, 0x3d, 0x0a, 0x1b, 0x72, 0x61, 0x6e, 0x64, 0x61, 0x6f, 0x5f, 0x73, 0x6c, 0x61, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x72, 0x61, 0x6e, 0x64, 0x61, 0x6f, 0x53, 0x6c, 0x61,
	0x73, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12,
	0x41, 0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x6c, 0x61, 0x73,
	0x68, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72,
	0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x65,
	0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x65, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x65, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x9c, 0x01, 0x0a,
	0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x78, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x29, 0x0a, 0x10, 0x72, 0x61, 0x6e, 0x64, 0x61, 0x6f, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x61, 0x6e, 0x64,
	0x61, 0x6f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x18, 0x0a, 0x02, 0x54, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x75, 0x73,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chainrpc_proto_common_proto_rawDescOnce sync.Once
	file_chainrpc_proto_common_proto_rawDescData = file_chainrpc_proto_common_proto_rawDesc
)

func file_chainrpc_proto_common_proto_rawDescGZIP() []byte {
	file_chainrpc_proto_common_proto_rawDescOnce.Do(func() {
		file_chainrpc_proto_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_chainrpc_proto_common_proto_rawDescData)
	})
	return file_chainrpc_proto_common_proto_rawDescData
}

var file_chainrpc_proto_common_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_chainrpc_proto_common_proto_goTypes = []interface{}{
	(*Empty)(nil),       // 0: Empty
	(*Hash)(nil),        // 1: Hash
	(*Height)(nil),      // 2: Height
	(*Success)(nil),     // 3: Success
	(*KeyPair)(nil),     // 4: KeyPair
	(*RawData)(nil),     // 5: RawData
	(*BlockHeader)(nil), // 6: BlockHeader
	(*Block)(nil),       // 7: Block
	(*Account)(nil),     // 8: Account
	(*Tx)(nil),          // 9: Tx
}
var file_chainrpc_proto_common_proto_depIdxs = []int32{
	6, // 0: Block.header:type_name -> BlockHeader
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chainrpc_proto_common_proto_init() }
func file_chainrpc_proto_common_proto_init() {
	if File_chainrpc_proto_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chainrpc_proto_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_chainrpc_proto_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hash); i {
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
		file_chainrpc_proto_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Height); i {
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
		file_chainrpc_proto_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
		file_chainrpc_proto_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyPair); i {
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
		file_chainrpc_proto_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawData); i {
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
		file_chainrpc_proto_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeader); i {
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
		file_chainrpc_proto_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_chainrpc_proto_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_chainrpc_proto_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tx); i {
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
			RawDescriptor: file_chainrpc_proto_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chainrpc_proto_common_proto_goTypes,
		DependencyIndexes: file_chainrpc_proto_common_proto_depIdxs,
		MessageInfos:      file_chainrpc_proto_common_proto_msgTypes,
	}.Build()
	File_chainrpc_proto_common_proto = out.File
	file_chainrpc_proto_common_proto_rawDesc = nil
	file_chainrpc_proto_common_proto_goTypes = nil
	file_chainrpc_proto_common_proto_depIdxs = nil
}
