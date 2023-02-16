// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package zbay

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AttendantsClient is the client API for Attendants service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttendantsClient interface {
	Add(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (*Attendant, error)
	Get(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (*Attendant, error)
	Update(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (*Attendant, error)
	List(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (Attendants_ListClient, error)
	Delete(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Login(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (*Attendant, error)
	Certificate(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (*Attendant, error)
}

type attendantsClient struct {
	cc grpc.ClientConnInterface
}

func NewAttendantsClient(cc grpc.ClientConnInterface) AttendantsClient {
	return &attendantsClient{cc}
}

func (c *attendantsClient) Add(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (*Attendant, error) {
	out := new(Attendant)
	err := c.cc.Invoke(ctx, "/zbay.Attendants/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendantsClient) Get(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (*Attendant, error) {
	out := new(Attendant)
	err := c.cc.Invoke(ctx, "/zbay.Attendants/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendantsClient) Update(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (*Attendant, error) {
	out := new(Attendant)
	err := c.cc.Invoke(ctx, "/zbay.Attendants/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendantsClient) List(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (Attendants_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Attendants_ServiceDesc.Streams[0], "/zbay.Attendants/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &attendantsListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Attendants_ListClient interface {
	Recv() (*Attendant, error)
	grpc.ClientStream
}

type attendantsListClient struct {
	grpc.ClientStream
}

func (x *attendantsListClient) Recv() (*Attendant, error) {
	m := new(Attendant)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *attendantsClient) Delete(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zbay.Attendants/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendantsClient) Login(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (*Attendant, error) {
	out := new(Attendant)
	err := c.cc.Invoke(ctx, "/zbay.Attendants/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendantsClient) Certificate(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (*Attendant, error) {
	out := new(Attendant)
	err := c.cc.Invoke(ctx, "/zbay.Attendants/Certificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttendantsServer is the server API for Attendants service.
// All implementations must embed UnimplementedAttendantsServer
// for forward compatibility
type AttendantsServer interface {
	Add(context.Context, *Attendant) (*Attendant, error)
	Get(context.Context, *Attendant) (*Attendant, error)
	Update(context.Context, *Attendant) (*Attendant, error)
	List(*Attendant, Attendants_ListServer) error
	Delete(context.Context, *Attendant) (*emptypb.Empty, error)
	Login(context.Context, *Attendant) (*Attendant, error)
	Certificate(context.Context, *Attendant) (*Attendant, error)
	mustEmbedUnimplementedAttendantsServer()
}

// UnimplementedAttendantsServer must be embedded to have forward compatible implementations.
type UnimplementedAttendantsServer struct {
}

func (UnimplementedAttendantsServer) Add(context.Context, *Attendant) (*Attendant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedAttendantsServer) Get(context.Context, *Attendant) (*Attendant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAttendantsServer) Update(context.Context, *Attendant) (*Attendant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAttendantsServer) List(*Attendant, Attendants_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAttendantsServer) Delete(context.Context, *Attendant) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAttendantsServer) Login(context.Context, *Attendant) (*Attendant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAttendantsServer) Certificate(context.Context, *Attendant) (*Attendant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Certificate not implemented")
}
func (UnimplementedAttendantsServer) mustEmbedUnimplementedAttendantsServer() {}

// UnsafeAttendantsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttendantsServer will
// result in compilation errors.
type UnsafeAttendantsServer interface {
	mustEmbedUnimplementedAttendantsServer()
}

func RegisterAttendantsServer(s grpc.ServiceRegistrar, srv AttendantsServer) {
	s.RegisterService(&Attendants_ServiceDesc, srv)
}

func _Attendants_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Attendant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendantsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Attendants/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendantsServer).Add(ctx, req.(*Attendant))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attendants_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Attendant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendantsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Attendants/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendantsServer).Get(ctx, req.(*Attendant))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attendants_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Attendant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendantsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Attendants/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendantsServer).Update(ctx, req.(*Attendant))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attendants_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Attendant)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AttendantsServer).List(m, &attendantsListServer{stream})
}

type Attendants_ListServer interface {
	Send(*Attendant) error
	grpc.ServerStream
}

type attendantsListServer struct {
	grpc.ServerStream
}

func (x *attendantsListServer) Send(m *Attendant) error {
	return x.ServerStream.SendMsg(m)
}

func _Attendants_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Attendant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendantsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Attendants/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendantsServer).Delete(ctx, req.(*Attendant))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attendants_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Attendant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendantsServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Attendants/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendantsServer).Login(ctx, req.(*Attendant))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attendants_Certificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Attendant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendantsServer).Certificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Attendants/Certificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendantsServer).Certificate(ctx, req.(*Attendant))
	}
	return interceptor(ctx, in, info, handler)
}

// Attendants_ServiceDesc is the grpc.ServiceDesc for Attendants service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Attendants_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zbay.Attendants",
	HandlerType: (*AttendantsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Attendants_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Attendants_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Attendants_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Attendants_Delete_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Attendants_Login_Handler,
		},
		{
			MethodName: "Certificate",
			Handler:    _Attendants_Certificate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Attendants_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "attendant.proto",
}

// AddressesClient is the client API for Addresses service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressesClient interface {
	Add(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error)
	Get(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error)
	Update(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error)
	Delete(ctx context.Context, in *Address, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (Addresses_ListClient, error)
}

type addressesClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressesClient(cc grpc.ClientConnInterface) AddressesClient {
	return &addressesClient{cc}
}

func (c *addressesClient) Add(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/zbay.Addresses/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressesClient) Get(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/zbay.Addresses/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressesClient) Update(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/zbay.Addresses/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressesClient) Delete(ctx context.Context, in *Address, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zbay.Addresses/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressesClient) List(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (Addresses_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Addresses_ServiceDesc.Streams[0], "/zbay.Addresses/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &addressesListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Addresses_ListClient interface {
	Recv() (*Address, error)
	grpc.ClientStream
}

type addressesListClient struct {
	grpc.ClientStream
}

func (x *addressesListClient) Recv() (*Address, error) {
	m := new(Address)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AddressesServer is the server API for Addresses service.
// All implementations must embed UnimplementedAddressesServer
// for forward compatibility
type AddressesServer interface {
	Add(context.Context, *Address) (*Address, error)
	Get(context.Context, *Address) (*Address, error)
	Update(context.Context, *Address) (*Address, error)
	Delete(context.Context, *Address) (*emptypb.Empty, error)
	List(*Attendant, Addresses_ListServer) error
	mustEmbedUnimplementedAddressesServer()
}

// UnimplementedAddressesServer must be embedded to have forward compatible implementations.
type UnimplementedAddressesServer struct {
}

func (UnimplementedAddressesServer) Add(context.Context, *Address) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedAddressesServer) Get(context.Context, *Address) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAddressesServer) Update(context.Context, *Address) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAddressesServer) Delete(context.Context, *Address) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAddressesServer) List(*Attendant, Addresses_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAddressesServer) mustEmbedUnimplementedAddressesServer() {}

// UnsafeAddressesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressesServer will
// result in compilation errors.
type UnsafeAddressesServer interface {
	mustEmbedUnimplementedAddressesServer()
}

func RegisterAddressesServer(s grpc.ServiceRegistrar, srv AddressesServer) {
	s.RegisterService(&Addresses_ServiceDesc, srv)
}

func _Addresses_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressesServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Addresses/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressesServer).Add(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addresses_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Addresses/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressesServer).Get(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addresses_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Addresses/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressesServer).Update(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addresses_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Addresses/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressesServer).Delete(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addresses_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Attendant)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AddressesServer).List(m, &addressesListServer{stream})
}

type Addresses_ListServer interface {
	Send(*Address) error
	grpc.ServerStream
}

type addressesListServer struct {
	grpc.ServerStream
}

func (x *addressesListServer) Send(m *Address) error {
	return x.ServerStream.SendMsg(m)
}

// Addresses_ServiceDesc is the grpc.ServiceDesc for Addresses service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Addresses_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zbay.Addresses",
	HandlerType: (*AddressesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Addresses_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Addresses_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Addresses_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Addresses_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Addresses_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "attendant.proto",
}

// MemosClient is the client API for Memos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemosClient interface {
	Add(ctx context.Context, in *Memo, opts ...grpc.CallOption) (*Memo, error)
	Get(ctx context.Context, in *Memo, opts ...grpc.CallOption) (*Memo, error)
	Update(ctx context.Context, in *Memo, opts ...grpc.CallOption) (*Memo, error)
	Delete(ctx context.Context, in *Memo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (Memos_ListClient, error)
}

type memosClient struct {
	cc grpc.ClientConnInterface
}

func NewMemosClient(cc grpc.ClientConnInterface) MemosClient {
	return &memosClient{cc}
}

func (c *memosClient) Add(ctx context.Context, in *Memo, opts ...grpc.CallOption) (*Memo, error) {
	out := new(Memo)
	err := c.cc.Invoke(ctx, "/zbay.Memos/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memosClient) Get(ctx context.Context, in *Memo, opts ...grpc.CallOption) (*Memo, error) {
	out := new(Memo)
	err := c.cc.Invoke(ctx, "/zbay.Memos/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memosClient) Update(ctx context.Context, in *Memo, opts ...grpc.CallOption) (*Memo, error) {
	out := new(Memo)
	err := c.cc.Invoke(ctx, "/zbay.Memos/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memosClient) Delete(ctx context.Context, in *Memo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zbay.Memos/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memosClient) List(ctx context.Context, in *Attendant, opts ...grpc.CallOption) (Memos_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Memos_ServiceDesc.Streams[0], "/zbay.Memos/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &memosListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Memos_ListClient interface {
	Recv() (*Memo, error)
	grpc.ClientStream
}

type memosListClient struct {
	grpc.ClientStream
}

func (x *memosListClient) Recv() (*Memo, error) {
	m := new(Memo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MemosServer is the server API for Memos service.
// All implementations must embed UnimplementedMemosServer
// for forward compatibility
type MemosServer interface {
	Add(context.Context, *Memo) (*Memo, error)
	Get(context.Context, *Memo) (*Memo, error)
	Update(context.Context, *Memo) (*Memo, error)
	Delete(context.Context, *Memo) (*emptypb.Empty, error)
	List(*Attendant, Memos_ListServer) error
	mustEmbedUnimplementedMemosServer()
}

// UnimplementedMemosServer must be embedded to have forward compatible implementations.
type UnimplementedMemosServer struct {
}

func (UnimplementedMemosServer) Add(context.Context, *Memo) (*Memo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedMemosServer) Get(context.Context, *Memo) (*Memo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMemosServer) Update(context.Context, *Memo) (*Memo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMemosServer) Delete(context.Context, *Memo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMemosServer) List(*Attendant, Memos_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMemosServer) mustEmbedUnimplementedMemosServer() {}

// UnsafeMemosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemosServer will
// result in compilation errors.
type UnsafeMemosServer interface {
	mustEmbedUnimplementedMemosServer()
}

func RegisterMemosServer(s grpc.ServiceRegistrar, srv MemosServer) {
	s.RegisterService(&Memos_ServiceDesc, srv)
}

func _Memos_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Memo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemosServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Memos/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemosServer).Add(ctx, req.(*Memo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Memos_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Memo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemosServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Memos/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemosServer).Get(ctx, req.(*Memo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Memos_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Memo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemosServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Memos/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemosServer).Update(ctx, req.(*Memo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Memos_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Memo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemosServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbay.Memos/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemosServer).Delete(ctx, req.(*Memo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Memos_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Attendant)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MemosServer).List(m, &memosListServer{stream})
}

type Memos_ListServer interface {
	Send(*Memo) error
	grpc.ServerStream
}

type memosListServer struct {
	grpc.ServerStream
}

func (x *memosListServer) Send(m *Memo) error {
	return x.ServerStream.SendMsg(m)
}

// Memos_ServiceDesc is the grpc.ServiceDesc for Memos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Memos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zbay.Memos",
	HandlerType: (*MemosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Memos_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Memos_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Memos_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Memos_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Memos_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "attendant.proto",
}