// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: banktransfer/banktransfer.proto

package banktransfer

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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BankName string `protobuf:"bytes,2,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
	Number   string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banktransfer_banktransfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_banktransfer_banktransfer_proto_msgTypes[0]
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
	return file_banktransfer_banktransfer_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *Account) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DonationId  int32    `protobuf:"varint,2,opt,name=donation_id,json=donationId,proto3" json:"donation_id,omitempty"`
	Amount      float32  `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Reference   string   `protobuf:"bytes,4,opt,name=reference,proto3" json:"reference,omitempty"`
	FromAccount *Account `protobuf:"bytes,5,opt,name=from_account,json=fromAccount,proto3" json:"from_account,omitempty"`
	To_Account  *Account `protobuf:"bytes,6,opt,name=to_Account,json=toAccount,proto3" json:"to_Account,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banktransfer_banktransfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_banktransfer_banktransfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_banktransfer_banktransfer_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Transaction) GetDonationId() int32 {
	if x != nil {
		return x.DonationId
	}
	return 0
}

func (x *Transaction) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *Transaction) GetFromAccount() *Account {
	if x != nil {
		return x.FromAccount
	}
	return nil
}

func (x *Transaction) GetTo_Account() *Account {
	if x != nil {
		return x.To_Account
	}
	return nil
}

type ProcessingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProcessingResponse) Reset() {
	*x = ProcessingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banktransfer_banktransfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessingResponse) ProtoMessage() {}

func (x *ProcessingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banktransfer_banktransfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessingResponse.ProtoReflect.Descriptor instead.
func (*ProcessingResponse) Descriptor() ([]byte, []int) {
	return file_banktransfer_banktransfer_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessingResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_banktransfer_banktransfer_proto protoreflect.FileDescriptor

var file_banktransfer_banktransfer_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x61, 0x6e, 0x6b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x62,
	0x61, 0x6e, 0x6b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x62, 0x61, 0x6e, 0x6b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0xe4, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x62, 0x61, 0x6e, 0x6b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x34, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x09, 0x74, 0x6f,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0xae, 0x01,
	0x0a, 0x0c, 0x42, 0x61, 0x6e, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x44,
	0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12,
	0x19, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x2e, 0x62, 0x61,
	0x6e, 0x6b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x19, 0x2e,
	0x62, 0x61, 0x6e, 0x6b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x43,
	0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x79, 0x61, 0x6b, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x79, 0x61, 0x6b, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_banktransfer_banktransfer_proto_rawDescOnce sync.Once
	file_banktransfer_banktransfer_proto_rawDescData = file_banktransfer_banktransfer_proto_rawDesc
)

func file_banktransfer_banktransfer_proto_rawDescGZIP() []byte {
	file_banktransfer_banktransfer_proto_rawDescOnce.Do(func() {
		file_banktransfer_banktransfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_banktransfer_banktransfer_proto_rawDescData)
	})
	return file_banktransfer_banktransfer_proto_rawDescData
}

var file_banktransfer_banktransfer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_banktransfer_banktransfer_proto_goTypes = []interface{}{
	(*Account)(nil),            // 0: banktransfer.Account
	(*Transaction)(nil),        // 1: banktransfer.Transaction
	(*ProcessingResponse)(nil), // 2: banktransfer.ProcessingResponse
	(*empty.Empty)(nil),        // 3: google.protobuf.Empty
}
var file_banktransfer_banktransfer_proto_depIdxs = []int32{
	0, // 0: banktransfer.Transaction.from_account:type_name -> banktransfer.Account
	0, // 1: banktransfer.Transaction.to_Account:type_name -> banktransfer.Account
	1, // 2: banktransfer.BankTransfer.TransferMoney:input_type -> banktransfer.Transaction
	2, // 3: banktransfer.BankTransfer.ProcessTransactions:input_type -> banktransfer.ProcessingResponse
	3, // 4: banktransfer.BankTransfer.TransferMoney:output_type -> google.protobuf.Empty
	1, // 5: banktransfer.BankTransfer.ProcessTransactions:output_type -> banktransfer.Transaction
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_banktransfer_banktransfer_proto_init() }
func file_banktransfer_banktransfer_proto_init() {
	if File_banktransfer_banktransfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banktransfer_banktransfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_banktransfer_banktransfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_banktransfer_banktransfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessingResponse); i {
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
			RawDescriptor: file_banktransfer_banktransfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_banktransfer_banktransfer_proto_goTypes,
		DependencyIndexes: file_banktransfer_banktransfer_proto_depIdxs,
		MessageInfos:      file_banktransfer_banktransfer_proto_msgTypes,
	}.Build()
	File_banktransfer_banktransfer_proto = out.File
	file_banktransfer_banktransfer_proto_rawDesc = nil
	file_banktransfer_banktransfer_proto_goTypes = nil
	file_banktransfer_banktransfer_proto_depIdxs = nil
}
