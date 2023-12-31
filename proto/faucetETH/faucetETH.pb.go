// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: faucetETH/faucetETH.proto

package faucetETH

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type FaucetETHRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletAddress string `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
}

func (x *FaucetETHRequest) Reset() {
	*x = FaucetETHRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faucetETH_faucetETH_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaucetETHRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaucetETHRequest) ProtoMessage() {}

func (x *FaucetETHRequest) ProtoReflect() protoreflect.Message {
	mi := &file_faucetETH_faucetETH_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaucetETHRequest.ProtoReflect.Descriptor instead.
func (*FaucetETHRequest) Descriptor() ([]byte, []int) {
	return file_faucetETH_faucetETH_proto_rawDescGZIP(), []int{0}
}

func (x *FaucetETHRequest) GetWalletAddress() string {
	if x != nil {
		return x.WalletAddress
	}
	return ""
}

type FaucetETHResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionHash string `protobuf:"bytes,1,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	EthBalance      string `protobuf:"bytes,2,opt,name=eth_balance,json=ethBalance,proto3" json:"eth_balance,omitempty"`
	TokenBalance    string `protobuf:"bytes,3,opt,name=token_balance,json=tokenBalance,proto3" json:"token_balance,omitempty"`
}

func (x *FaucetETHResponse) Reset() {
	*x = FaucetETHResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faucetETH_faucetETH_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaucetETHResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaucetETHResponse) ProtoMessage() {}

func (x *FaucetETHResponse) ProtoReflect() protoreflect.Message {
	mi := &file_faucetETH_faucetETH_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaucetETHResponse.ProtoReflect.Descriptor instead.
func (*FaucetETHResponse) Descriptor() ([]byte, []int) {
	return file_faucetETH_faucetETH_proto_rawDescGZIP(), []int{1}
}

func (x *FaucetETHResponse) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *FaucetETHResponse) GetEthBalance() string {
	if x != nil {
		return x.EthBalance
	}
	return ""
}

func (x *FaucetETHResponse) GetTokenBalance() string {
	if x != nil {
		return x.TokenBalance
	}
	return ""
}

var File_faucetETH_faucetETH_proto protoreflect.FileDescriptor

var file_faucetETH_faucetETH_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x45, 0x54, 0x48, 0x2f, 0x66, 0x61, 0x75, 0x63,
	0x65, 0x74, 0x45, 0x54, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x66, 0x61, 0x75,
	0x63, 0x65, 0x74, 0x45, 0x54, 0x48, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x10, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0x45, 0x54,
	0x48, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x84, 0x01, 0x0a, 0x11, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0x45, 0x54, 0x48, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x74, 0x68, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x74, 0x68, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x73, 0x0a, 0x09, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74,
	0x45, 0x54, 0x48, 0x12, 0x66, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x54,
	0x48, 0x12, 0x1b, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x45, 0x54, 0x48, 0x2e, 0x46, 0x61,
	0x75, 0x63, 0x65, 0x74, 0x45, 0x54, 0x48, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x45, 0x54, 0x48, 0x2e, 0x46, 0x61, 0x75, 0x63, 0x65,
	0x74, 0x45, 0x54, 0x48, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74,
	0x45, 0x54, 0x48, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x1c, 0x5a, 0x1a, 0x65,
	0x74, 0x68, 0x2d, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x45, 0x54, 0x48, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_faucetETH_faucetETH_proto_rawDescOnce sync.Once
	file_faucetETH_faucetETH_proto_rawDescData = file_faucetETH_faucetETH_proto_rawDesc
)

func file_faucetETH_faucetETH_proto_rawDescGZIP() []byte {
	file_faucetETH_faucetETH_proto_rawDescOnce.Do(func() {
		file_faucetETH_faucetETH_proto_rawDescData = protoimpl.X.CompressGZIP(file_faucetETH_faucetETH_proto_rawDescData)
	})
	return file_faucetETH_faucetETH_proto_rawDescData
}

var file_faucetETH_faucetETH_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_faucetETH_faucetETH_proto_goTypes = []interface{}{
	(*FaucetETHRequest)(nil),  // 0: faucetETH.FaucetETHRequest
	(*FaucetETHResponse)(nil), // 1: faucetETH.FaucetETHResponse
}
var file_faucetETH_faucetETH_proto_depIdxs = []int32{
	0, // 0: faucetETH.FaucetETH.RequestETH:input_type -> faucetETH.FaucetETHRequest
	1, // 1: faucetETH.FaucetETH.RequestETH:output_type -> faucetETH.FaucetETHResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_faucetETH_faucetETH_proto_init() }
func file_faucetETH_faucetETH_proto_init() {
	if File_faucetETH_faucetETH_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_faucetETH_faucetETH_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaucetETHRequest); i {
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
		file_faucetETH_faucetETH_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaucetETHResponse); i {
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
			RawDescriptor: file_faucetETH_faucetETH_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_faucetETH_faucetETH_proto_goTypes,
		DependencyIndexes: file_faucetETH_faucetETH_proto_depIdxs,
		MessageInfos:      file_faucetETH_faucetETH_proto_msgTypes,
	}.Build()
	File_faucetETH_faucetETH_proto = out.File
	file_faucetETH_faucetETH_proto_rawDesc = nil
	file_faucetETH_faucetETH_proto_goTypes = nil
	file_faucetETH_faucetETH_proto_depIdxs = nil
}
