// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: attendant.proto

package zwan

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Attendant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password    string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Telephone   string                 `protobuf:"bytes,4,opt,name=telephone,proto3" json:"telephone,omitempty"`
	Icon        string                 `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	Signature   string                 `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	Cert        *Certification         `protobuf:"bytes,9,opt,name=cert,proto3" json:"cert,omitempty"`
	Shops       []*Shop                `protobuf:"bytes,10,rep,name=shops,proto3" json:"shops,omitempty"`
	Annotations map[string]string      `protobuf:"bytes,7,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Created     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Attendant) Reset() {
	*x = Attendant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attendant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attendant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attendant) ProtoMessage() {}

func (x *Attendant) ProtoReflect() protoreflect.Message {
	mi := &file_attendant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attendant.ProtoReflect.Descriptor instead.
func (*Attendant) Descriptor() ([]byte, []int) {
	return file_attendant_proto_rawDescGZIP(), []int{0}
}

func (x *Attendant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Attendant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attendant) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Attendant) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *Attendant) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Attendant) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Attendant) GetCert() *Certification {
	if x != nil {
		return x.Cert
	}
	return nil
}

func (x *Attendant) GetShops() []*Shop {
	if x != nil {
		return x.Shops
	}
	return nil
}

func (x *Attendant) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Attendant) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type Certification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName string   `protobuf:"bytes,1,opt,name=fullName,proto3" json:"fullName,omitempty"`
	IdCardNo string   `protobuf:"bytes,2,opt,name=idCardNo,proto3" json:"idCardNo,omitempty"`
	Images   []string `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *Certification) Reset() {
	*x = Certification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attendant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certification) ProtoMessage() {}

func (x *Certification) ProtoReflect() protoreflect.Message {
	mi := &file_attendant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certification.ProtoReflect.Descriptor instead.
func (*Certification) Descriptor() ([]byte, []int) {
	return file_attendant_proto_rawDescGZIP(), []int{1}
}

func (x *Certification) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Certification) GetIdCardNo() string {
	if x != nil {
		return x.IdCardNo
	}
	return ""
}

func (x *Certification) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AttendantId string                 `protobuf:"bytes,2,opt,name=AttendantId,proto3" json:"AttendantId,omitempty"`
	Contact     string                 `protobuf:"bytes,3,opt,name=contact,proto3" json:"contact,omitempty"`
	Telephone   string                 `protobuf:"bytes,4,opt,name=telephone,proto3" json:"telephone,omitempty"`
	Location    string                 `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	Default     bool                   `protobuf:"varint,8,opt,name=default,proto3" json:"default,omitempty"`
	Annotations map[string]string      `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Created     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attendant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_attendant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_attendant_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Address) GetAttendantId() string {
	if x != nil {
		return x.AttendantId
	}
	return ""
}

func (x *Address) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *Address) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *Address) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Address) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

func (x *Address) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Address) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type Shop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Shop) Reset() {
	*x = Shop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attendant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shop) ProtoMessage() {}

func (x *Shop) ProtoReflect() protoreflect.Message {
	mi := &file_attendant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shop.ProtoReflect.Descriptor instead.
func (*Shop) Descriptor() ([]byte, []int) {
	return file_attendant_proto_rawDescGZIP(), []int{3}
}

func (x *Shop) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Shop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Shop) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Memo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AttendantId string                 `protobuf:"bytes,2,opt,name=AttendantId,proto3" json:"AttendantId,omitempty"`
	Title       string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content     string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Location    string                 `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	Annotations map[string]string      `protobuf:"bytes,5,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Created     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	Updated     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *Memo) Reset() {
	*x = Memo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attendant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memo) ProtoMessage() {}

func (x *Memo) ProtoReflect() protoreflect.Message {
	mi := &file_attendant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memo.ProtoReflect.Descriptor instead.
func (*Memo) Descriptor() ([]byte, []int) {
	return file_attendant_proto_rawDescGZIP(), []int{4}
}

func (x *Memo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Memo) GetAttendantId() string {
	if x != nil {
		return x.AttendantId
	}
	return ""
}

func (x *Memo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Memo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Memo) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Memo) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Memo) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Memo) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

var File_attendant_proto protoreflect.FileDescriptor

var file_attendant_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x7a, 0x77, 0x61, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x03, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64,
	0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x05,
	0x73, 0x68, 0x6f, 0x70, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x7a, 0x77,
	0x61, 0x6e, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x12, 0x42,
	0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e,
	0x64, 0x61, 0x6e, 0x74, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5f, 0x0a, 0x0d, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x4e,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x4e,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0xe1, 0x02, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x74, 0x74, 0x65,
	0x6e, 0x64, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x7a, 0x77, 0x61,
	0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x3e, 0x0a,
	0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4c, 0x0a,
	0x04, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xef, 0x02, 0x0a, 0x04,
	0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x74, 0x74, 0x65, 0x6e,
	0x64, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x4d,
	0x65, 0x6d, 0x6f, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x3e, 0x0a,
	0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xd3, 0x02,
	0x0a, 0x0a, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x03,
	0x41, 0x64, 0x64, 0x12, 0x0f, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e,
	0x64, 0x61, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65,
	0x6e, 0x64, 0x61, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f,
	0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x1a,
	0x0f, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74,
	0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x7a,
	0x77, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x1a, 0x0f, 0x2e,
	0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x22, 0x00,
	0x12, 0x2c, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e,
	0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x7a, 0x77, 0x61, 0x6e,
	0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x33,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e,
	0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0f, 0x2e, 0x7a,
	0x77, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x1a, 0x0f, 0x2e,
	0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x22, 0x00,
	0x12, 0x31, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12,
	0x0f, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74,
	0x1a, 0x0f, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e,
	0x74, 0x22, 0x00, 0x32, 0xe2, 0x01, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x25, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0d, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x0d, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x0d, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x0d,
	0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12,
	0x28, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x7a, 0x77, 0x61, 0x6e,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x0d, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65,
	0x6e, 0x64, 0x61, 0x6e, 0x74, 0x1a, 0x0d, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x00, 0x30, 0x01, 0x32, 0xc6, 0x01, 0x0a, 0x05, 0x4d, 0x65, 0x6d,
	0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0a, 0x2e, 0x7a, 0x77, 0x61, 0x6e,
	0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x1a, 0x0a, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x4d, 0x65, 0x6d,
	0x6f, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0a, 0x2e, 0x7a, 0x77, 0x61,
	0x6e, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x1a, 0x0a, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x4d, 0x65,
	0x6d, 0x6f, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0a,
	0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x1a, 0x0a, 0x2e, 0x7a, 0x77, 0x61,
	0x6e, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x0a, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x0f, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e,
	0x74, 0x1a, 0x0a, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x11, 0x5a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x6f, 0x3b,
	0x7a, 0x77, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attendant_proto_rawDescOnce sync.Once
	file_attendant_proto_rawDescData = file_attendant_proto_rawDesc
)

func file_attendant_proto_rawDescGZIP() []byte {
	file_attendant_proto_rawDescOnce.Do(func() {
		file_attendant_proto_rawDescData = protoimpl.X.CompressGZIP(file_attendant_proto_rawDescData)
	})
	return file_attendant_proto_rawDescData
}

var file_attendant_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_attendant_proto_goTypes = []interface{}{
	(*Attendant)(nil),             // 0: zwan.Attendant
	(*Certification)(nil),         // 1: zwan.Certification
	(*Address)(nil),               // 2: zwan.Address
	(*Shop)(nil),                  // 3: zwan.Shop
	(*Memo)(nil),                  // 4: zwan.Memo
	nil,                           // 5: zwan.Attendant.AnnotationsEntry
	nil,                           // 6: zwan.Address.AnnotationsEntry
	nil,                           // 7: zwan.Memo.AnnotationsEntry
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_attendant_proto_depIdxs = []int32{
	1,  // 0: zwan.Attendant.cert:type_name -> zwan.Certification
	3,  // 1: zwan.Attendant.shops:type_name -> zwan.Shop
	5,  // 2: zwan.Attendant.annotations:type_name -> zwan.Attendant.AnnotationsEntry
	8,  // 3: zwan.Attendant.created:type_name -> google.protobuf.Timestamp
	6,  // 4: zwan.Address.annotations:type_name -> zwan.Address.AnnotationsEntry
	8,  // 5: zwan.Address.created:type_name -> google.protobuf.Timestamp
	7,  // 6: zwan.Memo.annotations:type_name -> zwan.Memo.AnnotationsEntry
	8,  // 7: zwan.Memo.created:type_name -> google.protobuf.Timestamp
	8,  // 8: zwan.Memo.updated:type_name -> google.protobuf.Timestamp
	0,  // 9: zwan.Attendants.Add:input_type -> zwan.Attendant
	0,  // 10: zwan.Attendants.Get:input_type -> zwan.Attendant
	0,  // 11: zwan.Attendants.Update:input_type -> zwan.Attendant
	0,  // 12: zwan.Attendants.List:input_type -> zwan.Attendant
	0,  // 13: zwan.Attendants.Delete:input_type -> zwan.Attendant
	0,  // 14: zwan.Attendants.Login:input_type -> zwan.Attendant
	0,  // 15: zwan.Attendants.Certificate:input_type -> zwan.Attendant
	2,  // 16: zwan.Addresses.Add:input_type -> zwan.Address
	2,  // 17: zwan.Addresses.Get:input_type -> zwan.Address
	2,  // 18: zwan.Addresses.Update:input_type -> zwan.Address
	2,  // 19: zwan.Addresses.Delete:input_type -> zwan.Address
	0,  // 20: zwan.Addresses.List:input_type -> zwan.Attendant
	4,  // 21: zwan.Memos.Add:input_type -> zwan.Memo
	4,  // 22: zwan.Memos.Get:input_type -> zwan.Memo
	4,  // 23: zwan.Memos.Update:input_type -> zwan.Memo
	4,  // 24: zwan.Memos.Delete:input_type -> zwan.Memo
	0,  // 25: zwan.Memos.List:input_type -> zwan.Attendant
	0,  // 26: zwan.Attendants.Add:output_type -> zwan.Attendant
	0,  // 27: zwan.Attendants.Get:output_type -> zwan.Attendant
	0,  // 28: zwan.Attendants.Update:output_type -> zwan.Attendant
	0,  // 29: zwan.Attendants.List:output_type -> zwan.Attendant
	9,  // 30: zwan.Attendants.Delete:output_type -> google.protobuf.Empty
	0,  // 31: zwan.Attendants.Login:output_type -> zwan.Attendant
	0,  // 32: zwan.Attendants.Certificate:output_type -> zwan.Attendant
	2,  // 33: zwan.Addresses.Add:output_type -> zwan.Address
	2,  // 34: zwan.Addresses.Get:output_type -> zwan.Address
	2,  // 35: zwan.Addresses.Update:output_type -> zwan.Address
	9,  // 36: zwan.Addresses.Delete:output_type -> google.protobuf.Empty
	2,  // 37: zwan.Addresses.List:output_type -> zwan.Address
	4,  // 38: zwan.Memos.Add:output_type -> zwan.Memo
	4,  // 39: zwan.Memos.Get:output_type -> zwan.Memo
	4,  // 40: zwan.Memos.Update:output_type -> zwan.Memo
	9,  // 41: zwan.Memos.Delete:output_type -> google.protobuf.Empty
	4,  // 42: zwan.Memos.List:output_type -> zwan.Memo
	26, // [26:43] is the sub-list for method output_type
	9,  // [9:26] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_attendant_proto_init() }
func file_attendant_proto_init() {
	if File_attendant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attendant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attendant); i {
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
		file_attendant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certification); i {
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
		file_attendant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_attendant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shop); i {
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
		file_attendant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memo); i {
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
			RawDescriptor: file_attendant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_attendant_proto_goTypes,
		DependencyIndexes: file_attendant_proto_depIdxs,
		MessageInfos:      file_attendant_proto_msgTypes,
	}.Build()
	File_attendant_proto = out.File
	file_attendant_proto_rawDesc = nil
	file_attendant_proto_goTypes = nil
	file_attendant_proto_depIdxs = nil
}
