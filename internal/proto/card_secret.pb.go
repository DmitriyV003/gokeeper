// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: internal/proto/card_secret.proto

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

type CreateCardSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardholderName string `protobuf:"bytes,1,opt,name=CardholderName,proto3" json:"CardholderName,omitempty"`
	Type           string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	ExpireDate     string `protobuf:"bytes,3,opt,name=ExpireDate,proto3" json:"ExpireDate,omitempty"`
	ValidFrom      string `protobuf:"bytes,4,opt,name=ValidFrom,proto3" json:"ValidFrom,omitempty"`
	AdditionalData string `protobuf:"bytes,5,opt,name=Additional_data,json=AdditionalData,proto3" json:"Additional_data,omitempty"`
	ID             int64  `protobuf:"varint,6,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID         int64  `protobuf:"varint,7,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Number         string `protobuf:"bytes,8,opt,name=Number,proto3" json:"Number,omitempty"`
	SecretCode     string `protobuf:"bytes,9,opt,name=SecretCode,proto3" json:"SecretCode,omitempty"`
}

func (x *CreateCardSecretRequest) Reset() {
	*x = CreateCardSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_card_secret_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCardSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCardSecretRequest) ProtoMessage() {}

func (x *CreateCardSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_card_secret_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCardSecretRequest.ProtoReflect.Descriptor instead.
func (*CreateCardSecretRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_card_secret_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCardSecretRequest) GetCardholderName() string {
	if x != nil {
		return x.CardholderName
	}
	return ""
}

func (x *CreateCardSecretRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateCardSecretRequest) GetExpireDate() string {
	if x != nil {
		return x.ExpireDate
	}
	return ""
}

func (x *CreateCardSecretRequest) GetValidFrom() string {
	if x != nil {
		return x.ValidFrom
	}
	return ""
}

func (x *CreateCardSecretRequest) GetAdditionalData() string {
	if x != nil {
		return x.AdditionalData
	}
	return ""
}

func (x *CreateCardSecretRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CreateCardSecretRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateCardSecretRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *CreateCardSecretRequest) GetSecretCode() string {
	if x != nil {
		return x.SecretCode
	}
	return ""
}

type UpdateCardSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardholderName string `protobuf:"bytes,1,opt,name=CardholderName,proto3" json:"CardholderName,omitempty"`
	Type           string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	ExpireDate     string `protobuf:"bytes,3,opt,name=ExpireDate,proto3" json:"ExpireDate,omitempty"`
	ValidFrom      string `protobuf:"bytes,4,opt,name=ValidFrom,proto3" json:"ValidFrom,omitempty"`
	AdditionalData string `protobuf:"bytes,5,opt,name=Additional_data,json=AdditionalData,proto3" json:"Additional_data,omitempty"`
	ID             int64  `protobuf:"varint,6,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID         int64  `protobuf:"varint,7,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Number         string `protobuf:"bytes,8,opt,name=Number,proto3" json:"Number,omitempty"`
	SecretCode     string `protobuf:"bytes,9,opt,name=SecretCode,proto3" json:"SecretCode,omitempty"`
}

func (x *UpdateCardSecretRequest) Reset() {
	*x = UpdateCardSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_card_secret_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCardSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCardSecretRequest) ProtoMessage() {}

func (x *UpdateCardSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_card_secret_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCardSecretRequest.ProtoReflect.Descriptor instead.
func (*UpdateCardSecretRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_card_secret_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateCardSecretRequest) GetCardholderName() string {
	if x != nil {
		return x.CardholderName
	}
	return ""
}

func (x *UpdateCardSecretRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateCardSecretRequest) GetExpireDate() string {
	if x != nil {
		return x.ExpireDate
	}
	return ""
}

func (x *UpdateCardSecretRequest) GetValidFrom() string {
	if x != nil {
		return x.ValidFrom
	}
	return ""
}

func (x *UpdateCardSecretRequest) GetAdditionalData() string {
	if x != nil {
		return x.AdditionalData
	}
	return ""
}

func (x *UpdateCardSecretRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateCardSecretRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UpdateCardSecretRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *UpdateCardSecretRequest) GetSecretCode() string {
	if x != nil {
		return x.SecretCode
	}
	return ""
}

type SecretCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,6,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *SecretCardResponse) Reset() {
	*x = SecretCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_card_secret_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretCardResponse) ProtoMessage() {}

func (x *SecretCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_card_secret_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretCardResponse.ProtoReflect.Descriptor instead.
func (*SecretCardResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_card_secret_proto_rawDescGZIP(), []int{2}
}

func (x *SecretCardResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type CardSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *CardSecretRequest) Reset() {
	*x = CardSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_card_secret_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardSecretRequest) ProtoMessage() {}

func (x *CardSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_card_secret_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardSecretRequest.ProtoReflect.Descriptor instead.
func (*CardSecretRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_card_secret_proto_rawDescGZIP(), []int{3}
}

func (x *CardSecretRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *CardSecretRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DeleteCardSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserID int64 `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *DeleteCardSecretRequest) Reset() {
	*x = DeleteCardSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_card_secret_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCardSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCardSecretRequest) ProtoMessage() {}

func (x *DeleteCardSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_card_secret_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCardSecretRequest.ProtoReflect.Descriptor instead.
func (*DeleteCardSecretRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_card_secret_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCardSecretRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteCardSecretRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

var File_internal_proto_card_secret_proto protoreflect.FileDescriptor

var file_internal_proto_card_secret_proto_rawDesc = []byte{
	0x0a, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x61, 0x72, 0x64, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x43,
	0x61, 0x72, 0x64, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x12,
	0x27, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x9c, 0x02, 0x0a, 0x17, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x61, 0x72, 0x64, 0x68, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x43, 0x61,
	0x72, 0x64, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x27,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x45, 0x0a,
	0x11, 0x43, 0x61, 0x72, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x41, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x32, 0x80, 0x02, 0x0a, 0x11, 0x43, 0x61, 0x72, 0x64,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12,
	0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_card_secret_proto_rawDescOnce sync.Once
	file_internal_proto_card_secret_proto_rawDescData = file_internal_proto_card_secret_proto_rawDesc
)

func file_internal_proto_card_secret_proto_rawDescGZIP() []byte {
	file_internal_proto_card_secret_proto_rawDescOnce.Do(func() {
		file_internal_proto_card_secret_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_card_secret_proto_rawDescData)
	})
	return file_internal_proto_card_secret_proto_rawDescData
}

var file_internal_proto_card_secret_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_proto_card_secret_proto_goTypes = []interface{}{
	(*CreateCardSecretRequest)(nil), // 0: proto.CreateCardSecretRequest
	(*UpdateCardSecretRequest)(nil), // 1: proto.UpdateCardSecretRequest
	(*SecretCardResponse)(nil),      // 2: proto.SecretCardResponse
	(*CardSecretRequest)(nil),       // 3: proto.CardSecretRequest
	(*DeleteCardSecretRequest)(nil), // 4: proto.DeleteCardSecretRequest
}
var file_internal_proto_card_secret_proto_depIdxs = []int32{
	0, // 0: proto.CardSecretService.CreateCardSecret:input_type -> proto.CreateCardSecretRequest
	1, // 1: proto.CardSecretService.UpdateCardSecret:input_type -> proto.UpdateCardSecretRequest
	4, // 2: proto.CardSecretService.DeleteCardSecret:input_type -> proto.DeleteCardSecretRequest
	2, // 3: proto.CardSecretService.CreateCardSecret:output_type -> proto.SecretCardResponse
	2, // 4: proto.CardSecretService.UpdateCardSecret:output_type -> proto.SecretCardResponse
	2, // 5: proto.CardSecretService.DeleteCardSecret:output_type -> proto.SecretCardResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_card_secret_proto_init() }
func file_internal_proto_card_secret_proto_init() {
	if File_internal_proto_card_secret_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_card_secret_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCardSecretRequest); i {
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
		file_internal_proto_card_secret_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCardSecretRequest); i {
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
		file_internal_proto_card_secret_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretCardResponse); i {
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
		file_internal_proto_card_secret_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardSecretRequest); i {
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
		file_internal_proto_card_secret_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCardSecretRequest); i {
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
			RawDescriptor: file_internal_proto_card_secret_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_card_secret_proto_goTypes,
		DependencyIndexes: file_internal_proto_card_secret_proto_depIdxs,
		MessageInfos:      file_internal_proto_card_secret_proto_msgTypes,
	}.Build()
	File_internal_proto_card_secret_proto = out.File
	file_internal_proto_card_secret_proto_rawDesc = nil
	file_internal_proto_card_secret_proto_goTypes = nil
	file_internal_proto_card_secret_proto_depIdxs = nil
}