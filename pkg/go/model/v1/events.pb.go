// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: model/v1/events.proto

package model

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

type BlockEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash   string              `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	BlockNumber string              `protobuf:"bytes,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	Network     *BlockEvent_Network `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	Block       *BlockEvent_Block   `protobuf:"bytes,4,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *BlockEvent) Reset() {
	*x = BlockEvent{}
	mi := &file_model_v1_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlockEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockEvent) ProtoMessage() {}

func (x *BlockEvent) ProtoReflect() protoreflect.Message {
	mi := &file_model_v1_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockEvent.ProtoReflect.Descriptor instead.
func (*BlockEvent) Descriptor() ([]byte, []int) {
	return file_model_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *BlockEvent) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *BlockEvent) GetBlockNumber() string {
	if x != nil {
		return x.BlockNumber
	}
	return ""
}

func (x *BlockEvent) GetNetwork() *BlockEvent_Network {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *BlockEvent) GetBlock() *BlockEvent_Block {
	if x != nil {
		return x.Block
	}
	return nil
}

type TransactionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction          *TransactionEvent_Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Network              *TransactionEvent_Network     `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	Addresses            map[string]bool               `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Block                *TransactionEvent_Block       `protobuf:"bytes,4,opt,name=block,proto3" json:"block,omitempty"`
	Logs                 []*TransactionEvent_Log       `protobuf:"bytes,5,rep,name=logs,proto3" json:"logs,omitempty"`
	IsContractDeployment bool                          `protobuf:"varint,6,opt,name=isContractDeployment,proto3" json:"isContractDeployment,omitempty"`
	ContractAddress      string                        `protobuf:"bytes,7,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	TxAddresses          map[string]bool               `protobuf:"bytes,8,rep,name=txAddresses,proto3" json:"txAddresses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *TransactionEvent) Reset() {
	*x = TransactionEvent{}
	mi := &file_model_v1_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionEvent) ProtoMessage() {}

func (x *TransactionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_model_v1_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionEvent.ProtoReflect.Descriptor instead.
func (*TransactionEvent) Descriptor() ([]byte, []int) {
	return file_model_v1_events_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionEvent) GetTransaction() *TransactionEvent_Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TransactionEvent) GetNetwork() *TransactionEvent_Network {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *TransactionEvent) GetAddresses() map[string]bool {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *TransactionEvent) GetBlock() *TransactionEvent_Block {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *TransactionEvent) GetLogs() []*TransactionEvent_Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *TransactionEvent) GetIsContractDeployment() bool {
	if x != nil {
		return x.IsContractDeployment
	}
	return false
}

func (x *TransactionEvent) GetContractAddress() string {
	if x != nil {
		return x.ContractAddress
	}
	return ""
}

func (x *TransactionEvent) GetTxAddresses() map[string]bool {
	if x != nil {
		return x.TxAddresses
	}
	return nil
}

type BlockEvent_Network struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId string `protobuf:"bytes,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
}

func (x *BlockEvent_Network) Reset() {
	*x = BlockEvent_Network{}
	mi := &file_model_v1_events_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlockEvent_Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockEvent_Network) ProtoMessage() {}

func (x *BlockEvent_Network) ProtoReflect() protoreflect.Message {
	mi := &file_model_v1_events_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockEvent_Network.ProtoReflect.Descriptor instead.
func (*BlockEvent_Network) Descriptor() ([]byte, []int) {
	return file_model_v1_events_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BlockEvent_Network) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

type BlockEvent_Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Difficulty       string   `protobuf:"bytes,1,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	ExtraData        string   `protobuf:"bytes,2,opt,name=extraData,proto3" json:"extraData,omitempty"`
	GasLimit         string   `protobuf:"bytes,3,opt,name=gasLimit,proto3" json:"gasLimit,omitempty"`
	GasUsed          string   `protobuf:"bytes,4,opt,name=gasUsed,proto3" json:"gasUsed,omitempty"`
	Hash             string   `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	LogsBloom        string   `protobuf:"bytes,6,opt,name=logsBloom,proto3" json:"logsBloom,omitempty"`
	Miner            string   `protobuf:"bytes,7,opt,name=miner,proto3" json:"miner,omitempty"`
	MixHash          string   `protobuf:"bytes,8,opt,name=mixHash,proto3" json:"mixHash,omitempty"`
	Nonce            string   `protobuf:"bytes,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Number           string   `protobuf:"bytes,10,opt,name=number,proto3" json:"number,omitempty"`
	ParentHash       string   `protobuf:"bytes,11,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	ReceiptsRoot     string   `protobuf:"bytes,12,opt,name=receiptsRoot,proto3" json:"receiptsRoot,omitempty"`
	Sha3Uncles       string   `protobuf:"bytes,13,opt,name=sha3Uncles,proto3" json:"sha3Uncles,omitempty"`
	Size             string   `protobuf:"bytes,14,opt,name=size,proto3" json:"size,omitempty"`
	StateRoot        string   `protobuf:"bytes,15,opt,name=stateRoot,proto3" json:"stateRoot,omitempty"`
	Timestamp        string   `protobuf:"bytes,16,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TotalDifficulty  string   `protobuf:"bytes,17,opt,name=totalDifficulty,proto3" json:"totalDifficulty,omitempty"`
	Transactions     []string `protobuf:"bytes,18,rep,name=transactions,proto3" json:"transactions,omitempty"`
	TransactionsRoot string   `protobuf:"bytes,19,opt,name=transactionsRoot,proto3" json:"transactionsRoot,omitempty"`
	Uncles           []string `protobuf:"bytes,20,rep,name=uncles,proto3" json:"uncles,omitempty"`
	BaseFeePerGas    string   `protobuf:"bytes,21,opt,name=baseFeePerGas,proto3" json:"baseFeePerGas,omitempty"`
}

func (x *BlockEvent_Block) Reset() {
	*x = BlockEvent_Block{}
	mi := &file_model_v1_events_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlockEvent_Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockEvent_Block) ProtoMessage() {}

func (x *BlockEvent_Block) ProtoReflect() protoreflect.Message {
	mi := &file_model_v1_events_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockEvent_Block.ProtoReflect.Descriptor instead.
func (*BlockEvent_Block) Descriptor() ([]byte, []int) {
	return file_model_v1_events_proto_rawDescGZIP(), []int{0, 1}
}

func (x *BlockEvent_Block) GetDifficulty() string {
	if x != nil {
		return x.Difficulty
	}
	return ""
}

func (x *BlockEvent_Block) GetExtraData() string {
	if x != nil {
		return x.ExtraData
	}
	return ""
}

func (x *BlockEvent_Block) GetGasLimit() string {
	if x != nil {
		return x.GasLimit
	}
	return ""
}

func (x *BlockEvent_Block) GetGasUsed() string {
	if x != nil {
		return x.GasUsed
	}
	return ""
}

func (x *BlockEvent_Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *BlockEvent_Block) GetLogsBloom() string {
	if x != nil {
		return x.LogsBloom
	}
	return ""
}

func (x *BlockEvent_Block) GetMiner() string {
	if x != nil {
		return x.Miner
	}
	return ""
}

func (x *BlockEvent_Block) GetMixHash() string {
	if x != nil {
		return x.MixHash
	}
	return ""
}

func (x *BlockEvent_Block) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *BlockEvent_Block) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *BlockEvent_Block) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *BlockEvent_Block) GetReceiptsRoot() string {
	if x != nil {
		return x.ReceiptsRoot
	}
	return ""
}

func (x *BlockEvent_Block) GetSha3Uncles() string {
	if x != nil {
		return x.Sha3Uncles
	}
	return ""
}

func (x *BlockEvent_Block) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *BlockEvent_Block) GetStateRoot() string {
	if x != nil {
		return x.StateRoot
	}
	return ""
}

func (x *BlockEvent_Block) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *BlockEvent_Block) GetTotalDifficulty() string {
	if x != nil {
		return x.TotalDifficulty
	}
	return ""
}

func (x *BlockEvent_Block) GetTransactions() []string {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *BlockEvent_Block) GetTransactionsRoot() string {
	if x != nil {
		return x.TransactionsRoot
	}
	return ""
}

func (x *BlockEvent_Block) GetUncles() []string {
	if x != nil {
		return x.Uncles
	}
	return nil
}

func (x *BlockEvent_Block) GetBaseFeePerGas() string {
	if x != nil {
		return x.BaseFeePerGas
	}
	return ""
}

type TransactionEvent_Network struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId string `protobuf:"bytes,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
}

func (x *TransactionEvent_Network) Reset() {
	*x = TransactionEvent_Network{}
	mi := &file_model_v1_events_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionEvent_Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionEvent_Network) ProtoMessage() {}

func (x *TransactionEvent_Network) ProtoReflect() protoreflect.Message {
	mi := &file_model_v1_events_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionEvent_Network.ProtoReflect.Descriptor instead.
func (*TransactionEvent_Network) Descriptor() ([]byte, []int) {
	return file_model_v1_events_proto_rawDescGZIP(), []int{1, 0}
}

func (x *TransactionEvent_Network) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

type TransactionEvent_Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash      string `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	BlockNumber    string `protobuf:"bytes,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	BlockTimestamp string `protobuf:"bytes,3,opt,name=blockTimestamp,proto3" json:"blockTimestamp,omitempty"`
	BaseFeePerGas  string `protobuf:"bytes,4,opt,name=baseFeePerGas,proto3" json:"baseFeePerGas,omitempty"`
}

func (x *TransactionEvent_Block) Reset() {
	*x = TransactionEvent_Block{}
	mi := &file_model_v1_events_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionEvent_Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionEvent_Block) ProtoMessage() {}

func (x *TransactionEvent_Block) ProtoReflect() protoreflect.Message {
	mi := &file_model_v1_events_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionEvent_Block.ProtoReflect.Descriptor instead.
func (*TransactionEvent_Block) Descriptor() ([]byte, []int) {
	return file_model_v1_events_proto_rawDescGZIP(), []int{1, 1}
}

func (x *TransactionEvent_Block) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *TransactionEvent_Block) GetBlockNumber() string {
	if x != nil {
		return x.BlockNumber
	}
	return ""
}

func (x *TransactionEvent_Block) GetBlockTimestamp() string {
	if x != nil {
		return x.BlockTimestamp
	}
	return ""
}

func (x *TransactionEvent_Block) GetBaseFeePerGas() string {
	if x != nil {
		return x.BaseFeePerGas
	}
	return ""
}

type TransactionEvent_Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                 string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Nonce                string `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	GasPrice             string `protobuf:"bytes,3,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
	Gas                  string `protobuf:"bytes,4,opt,name=gas,proto3" json:"gas,omitempty"`
	Value                string `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Input                string `protobuf:"bytes,6,opt,name=input,proto3" json:"input,omitempty"`
	V                    string `protobuf:"bytes,7,opt,name=v,proto3" json:"v,omitempty"`
	R                    string `protobuf:"bytes,8,opt,name=r,proto3" json:"r,omitempty"`
	S                    string `protobuf:"bytes,9,opt,name=s,proto3" json:"s,omitempty"`
	To                   string `protobuf:"bytes,10,opt,name=to,proto3" json:"to,omitempty"`
	Hash                 string `protobuf:"bytes,11,opt,name=hash,proto3" json:"hash,omitempty"`
	From                 string `protobuf:"bytes,12,opt,name=from,proto3" json:"from,omitempty"`
	MaxFeePerGas         string `protobuf:"bytes,13,opt,name=maxFeePerGas,proto3" json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas string `protobuf:"bytes,14,opt,name=maxPriorityFeePerGas,proto3" json:"maxPriorityFeePerGas,omitempty"`
}

func (x *TransactionEvent_Transaction) Reset() {
	*x = TransactionEvent_Transaction{}
	mi := &file_model_v1_events_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionEvent_Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionEvent_Transaction) ProtoMessage() {}

func (x *TransactionEvent_Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_model_v1_events_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionEvent_Transaction.ProtoReflect.Descriptor instead.
func (*TransactionEvent_Transaction) Descriptor() ([]byte, []int) {
	return file_model_v1_events_proto_rawDescGZIP(), []int{1, 2}
}

func (x *TransactionEvent_Transaction) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TransactionEvent_Transaction) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *TransactionEvent_Transaction) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

func (x *TransactionEvent_Transaction) GetGas() string {
	if x != nil {
		return x.Gas
	}
	return ""
}

func (x *TransactionEvent_Transaction) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TransactionEvent_Transaction) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *TransactionEvent_Transaction) GetV() string {
	if x != nil {
		return x.V
	}
	return ""
}

func (x *TransactionEvent_Transaction) GetR() string {
	if x != nil {
		return x.R
	}
	return ""
}

func (x *TransactionEvent_Transaction) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *TransactionEvent_Transaction) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TransactionEvent_Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *TransactionEvent_Transaction) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TransactionEvent_Transaction) GetMaxFeePerGas() string {
	if x != nil {
		return x.MaxFeePerGas
	}
	return ""
}

func (x *TransactionEvent_Transaction) GetMaxPriorityFeePerGas() string {
	if x != nil {
		return x.MaxPriorityFeePerGas
	}
	return ""
}

type TransactionEvent_Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address          string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Topics           []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
	Data             string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	BlockNumber      string   `protobuf:"bytes,4,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	TransactionHash  string   `protobuf:"bytes,5,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	TransactionIndex string   `protobuf:"bytes,6,opt,name=transactionIndex,proto3" json:"transactionIndex,omitempty"`
	BlockHash        string   `protobuf:"bytes,7,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	LogIndex         string   `protobuf:"bytes,8,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	Removed          bool     `protobuf:"varint,9,opt,name=removed,proto3" json:"removed,omitempty"`
}

func (x *TransactionEvent_Log) Reset() {
	*x = TransactionEvent_Log{}
	mi := &file_model_v1_events_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionEvent_Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionEvent_Log) ProtoMessage() {}

func (x *TransactionEvent_Log) ProtoReflect() protoreflect.Message {
	mi := &file_model_v1_events_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionEvent_Log.ProtoReflect.Descriptor instead.
func (*TransactionEvent_Log) Descriptor() ([]byte, []int) {
	return file_model_v1_events_proto_rawDescGZIP(), []int{1, 3}
}

func (x *TransactionEvent_Log) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TransactionEvent_Log) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *TransactionEvent_Log) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *TransactionEvent_Log) GetBlockNumber() string {
	if x != nil {
		return x.BlockNumber
	}
	return ""
}

func (x *TransactionEvent_Log) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *TransactionEvent_Log) GetTransactionIndex() string {
	if x != nil {
		return x.TransactionIndex
	}
	return ""
}

func (x *TransactionEvent_Log) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *TransactionEvent_Log) GetLogIndex() string {
	if x != nil {
		return x.LogIndex
	}
	return ""
}

func (x *TransactionEvent_Log) GetRemoved() bool {
	if x != nil {
		return x.Removed
	}
	return false
}

var File_model_v1_events_proto protoreflect.FileDescriptor

var file_model_v1_events_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x22, 0xd5, 0x06, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x20,
	0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x36, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x30, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x23, 0x0a, 0x07, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x1a,
	0xf7, 0x04, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66,
	0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x52, 0x6f, 0x6f,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x33, 0x55, 0x6e, 0x63, 0x6c, 0x65, 0x73, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x33, 0x55, 0x6e, 0x63, 0x6c, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63,
	0x75, 0x6c, 0x74, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x6f, 0x6f, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x6e, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x63,
	0x6c, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x46, 0x65, 0x65, 0x50, 0x65,
	0x72, 0x47, 0x61, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x73, 0x65,
	0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x22, 0x9f, 0x0b, 0x0a, 0x10, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x48,
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x47, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12,
	0x36, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x32, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x69,
	0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4d, 0x0a, 0x0b, 0x74, 0x78, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x78, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x74, 0x78, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x1a, 0x23, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x1a, 0x95, 0x01,
	0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x24, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x46, 0x65, 0x65, 0x50,
	0x65, 0x72, 0x47, 0x61, 0x73, 0x1a, 0xcb, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x67,
	0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x76, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x6d,
	0x61, 0x78, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x12,
	0x32, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x46, 0x65,
	0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6d,
	0x61, 0x78, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72,
	0x47, 0x61, 0x73, 0x1a, 0x97, 0x02, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a,
	0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x3c, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x54,
	0x78, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x66, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x14, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_v1_events_proto_rawDescOnce sync.Once
	file_model_v1_events_proto_rawDescData = file_model_v1_events_proto_rawDesc
)

func file_model_v1_events_proto_rawDescGZIP() []byte {
	file_model_v1_events_proto_rawDescOnce.Do(func() {
		file_model_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_v1_events_proto_rawDescData)
	})
	return file_model_v1_events_proto_rawDescData
}

var file_model_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_model_v1_events_proto_goTypes = []any{
	(*BlockEvent)(nil),                   // 0: model.v1.BlockEvent
	(*TransactionEvent)(nil),             // 1: model.v1.TransactionEvent
	(*BlockEvent_Network)(nil),           // 2: model.v1.BlockEvent.Network
	(*BlockEvent_Block)(nil),             // 3: model.v1.BlockEvent.Block
	(*TransactionEvent_Network)(nil),     // 4: model.v1.TransactionEvent.Network
	(*TransactionEvent_Block)(nil),       // 5: model.v1.TransactionEvent.Block
	(*TransactionEvent_Transaction)(nil), // 6: model.v1.TransactionEvent.Transaction
	(*TransactionEvent_Log)(nil),         // 7: model.v1.TransactionEvent.Log
	nil,                                  // 8: model.v1.TransactionEvent.AddressesEntry
	nil,                                  // 9: model.v1.TransactionEvent.TxAddressesEntry
}
var file_model_v1_events_proto_depIdxs = []int32{
	2, // 0: model.v1.BlockEvent.network:type_name -> model.v1.BlockEvent.Network
	3, // 1: model.v1.BlockEvent.block:type_name -> model.v1.BlockEvent.Block
	6, // 2: model.v1.TransactionEvent.transaction:type_name -> model.v1.TransactionEvent.Transaction
	4, // 3: model.v1.TransactionEvent.network:type_name -> model.v1.TransactionEvent.Network
	8, // 4: model.v1.TransactionEvent.addresses:type_name -> model.v1.TransactionEvent.AddressesEntry
	5, // 5: model.v1.TransactionEvent.block:type_name -> model.v1.TransactionEvent.Block
	7, // 6: model.v1.TransactionEvent.logs:type_name -> model.v1.TransactionEvent.Log
	9, // 7: model.v1.TransactionEvent.txAddresses:type_name -> model.v1.TransactionEvent.TxAddressesEntry
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_model_v1_events_proto_init() }
func file_model_v1_events_proto_init() {
	if File_model_v1_events_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_v1_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_v1_events_proto_goTypes,
		DependencyIndexes: file_model_v1_events_proto_depIdxs,
		MessageInfos:      file_model_v1_events_proto_msgTypes,
	}.Build()
	File_model_v1_events_proto = out.File
	file_model_v1_events_proto_rawDesc = nil
	file_model_v1_events_proto_goTypes = nil
	file_model_v1_events_proto_depIdxs = nil
}
