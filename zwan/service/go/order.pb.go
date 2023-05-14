// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: order.proto

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hospital   string                 `protobuf:"bytes,2,opt,name=hospital,proto3" json:"hospital,omitempty"`
	Department string                 `protobuf:"bytes,3,opt,name=department,proto3" json:"department,omitempty"`
	TargetTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=target_time,json=targetTime,proto3" json:"target_time,omitempty"`
	Patient    string                 `protobuf:"bytes,5,opt,name=patient,proto3" json:"patient,omitempty"`
	Companion  string                 `protobuf:"bytes,6,opt,name=companion,proto3" json:"companion,omitempty"`
	Notes      string                 `protobuf:"bytes,7,opt,name=notes,proto3" json:"notes,omitempty"`
	Created    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetHospital() string {
	if x != nil {
		return x.Hospital
	}
	return ""
}

func (x *Order) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *Order) GetTargetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.TargetTime
	}
	return nil
}

func (x *Order) GetPatient() string {
	if x != nil {
		return x.Patient
	}
	return ""
}

func (x *Order) GetCompanion() string {
	if x != nil {
		return x.Companion
	}
	return ""
}

func (x *Order) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *Order) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x7a,
	0x77, 0x61, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x94, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x32, 0xcb, 0x01, 0x0a, 0x06, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0b, 0x2e, 0x7a, 0x77, 0x61,
	0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x0b, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0b, 0x2e,
	0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x0b, 0x2e, 0x7a, 0x77, 0x61,
	0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x1a, 0x0b, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12,
	0x24, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0b, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x1a, 0x0b, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x0b, 0x2e, 0x7a, 0x77, 0x61, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x67, 0x6f, 0x3b, 0x7a, 0x77, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_order_proto_goTypes = []interface{}{
	(*Order)(nil),                 // 0: zwan.Order
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 2: google.protobuf.Empty
}
var file_order_proto_depIdxs = []int32{
	1, // 0: zwan.Order.target_time:type_name -> google.protobuf.Timestamp
	1, // 1: zwan.Order.created:type_name -> google.protobuf.Timestamp
	0, // 2: zwan.Orders.Add:input_type -> zwan.Order
	0, // 3: zwan.Orders.Get:input_type -> zwan.Order
	0, // 4: zwan.Orders.Update:input_type -> zwan.Order
	0, // 5: zwan.Orders.List:input_type -> zwan.Order
	0, // 6: zwan.Orders.Delete:input_type -> zwan.Order
	0, // 7: zwan.Orders.Add:output_type -> zwan.Order
	0, // 8: zwan.Orders.Get:output_type -> zwan.Order
	0, // 9: zwan.Orders.Update:output_type -> zwan.Order
	0, // 10: zwan.Orders.List:output_type -> zwan.Order
	2, // 11: zwan.Orders.Delete:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
