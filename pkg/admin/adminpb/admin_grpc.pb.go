// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: admin.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AdminService_AdminLoginRequest_FullMethodName   = "/pb.AdminService/AdminLoginRequest"
	AdminService_AddToAdminWallet_FullMethodName    = "/pb.AdminService/AddToAdminWallet"
	AdminService_AddCategory_FullMethodName         = "/pb.AdminService/AddCategory"
	AdminService_FindCategory_FullMethodName        = "/pb.AdminService/FindCategory"
	AdminService_FindCategories_FullMethodName      = "/pb.AdminService/FindCategories"
	AdminService_EditCategoryAdmin_FullMethodName   = "/pb.AdminService/EditCategoryAdmin"
	AdminService_RemoveCategoryAdmin_FullMethodName = "/pb.AdminService/RemoveCategoryAdmin"
	AdminService_FindProductByID_FullMethodName     = "/pb.AdminService/FindProductByID"
	AdminService_FindAllProducts_FullMethodName     = "/pb.AdminService/FindAllProducts"
	AdminService_RemoveProduct_FullMethodName       = "/pb.AdminService/RemoveProduct"
	AdminService_AdminBlockUser_FullMethodName      = "/pb.AdminService/AdminBlockUser"
	AdminService_FetchOrders_FullMethodName         = "/pb.AdminService/FetchOrders"
	AdminService_FetchOrderByUser_FullMethodName    = "/pb.AdminService/FetchOrderByUser"
	AdminService_FetchOrder_FullMethodName          = "/pb.AdminService/FetchOrder"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	AdminLoginRequest(ctx context.Context, in *AdminLogin, opts ...grpc.CallOption) (*AdminResponse, error)
	AddToAdminWallet(ctx context.Context, in *Amount, opts ...grpc.CallOption) (*AdminResponse, error)
	AddCategory(ctx context.Context, in *AdminCategory, opts ...grpc.CallOption) (*AdminResponse, error)
	FindCategory(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminCategory, error)
	FindCategories(ctx context.Context, in *AdminNoParam, opts ...grpc.CallOption) (*AdminCategoryList, error)
	EditCategoryAdmin(ctx context.Context, in *AdminCategory, opts ...grpc.CallOption) (*AdminCategory, error)
	RemoveCategoryAdmin(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminResponse, error)
	FindProductByID(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminProduct, error)
	FindAllProducts(ctx context.Context, in *AdminNoParam, opts ...grpc.CallOption) (*AdminProductList, error)
	RemoveProduct(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminResponse, error)
	AdminBlockUser(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminResponse, error)
	FetchOrders(ctx context.Context, in *AdminNoParam, opts ...grpc.CallOption) (*AdminOrderList, error)
	FetchOrderByUser(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminOrderList, error)
	FetchOrder(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminOrder, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) AdminLoginRequest(ctx context.Context, in *AdminLogin, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AdminService_AdminLoginRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AddToAdminWallet(ctx context.Context, in *Amount, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AdminService_AddToAdminWallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AddCategory(ctx context.Context, in *AdminCategory, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AdminService_AddCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindCategory(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminCategory, error) {
	out := new(AdminCategory)
	err := c.cc.Invoke(ctx, AdminService_FindCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindCategories(ctx context.Context, in *AdminNoParam, opts ...grpc.CallOption) (*AdminCategoryList, error) {
	out := new(AdminCategoryList)
	err := c.cc.Invoke(ctx, AdminService_FindCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) EditCategoryAdmin(ctx context.Context, in *AdminCategory, opts ...grpc.CallOption) (*AdminCategory, error) {
	out := new(AdminCategory)
	err := c.cc.Invoke(ctx, AdminService_EditCategoryAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) RemoveCategoryAdmin(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AdminService_RemoveCategoryAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindProductByID(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminProduct, error) {
	out := new(AdminProduct)
	err := c.cc.Invoke(ctx, AdminService_FindProductByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindAllProducts(ctx context.Context, in *AdminNoParam, opts ...grpc.CallOption) (*AdminProductList, error) {
	out := new(AdminProductList)
	err := c.cc.Invoke(ctx, AdminService_FindAllProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) RemoveProduct(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AdminService_RemoveProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminBlockUser(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AdminService_AdminBlockUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FetchOrders(ctx context.Context, in *AdminNoParam, opts ...grpc.CallOption) (*AdminOrderList, error) {
	out := new(AdminOrderList)
	err := c.cc.Invoke(ctx, AdminService_FetchOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FetchOrderByUser(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminOrderList, error) {
	out := new(AdminOrderList)
	err := c.cc.Invoke(ctx, AdminService_FetchOrderByUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FetchOrder(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdminOrder, error) {
	out := new(AdminOrder)
	err := c.cc.Invoke(ctx, AdminService_FetchOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	AdminLoginRequest(context.Context, *AdminLogin) (*AdminResponse, error)
	AddToAdminWallet(context.Context, *Amount) (*AdminResponse, error)
	AddCategory(context.Context, *AdminCategory) (*AdminResponse, error)
	FindCategory(context.Context, *AdID) (*AdminCategory, error)
	FindCategories(context.Context, *AdminNoParam) (*AdminCategoryList, error)
	EditCategoryAdmin(context.Context, *AdminCategory) (*AdminCategory, error)
	RemoveCategoryAdmin(context.Context, *AdID) (*AdminResponse, error)
	FindProductByID(context.Context, *AdID) (*AdminProduct, error)
	FindAllProducts(context.Context, *AdminNoParam) (*AdminProductList, error)
	RemoveProduct(context.Context, *AdID) (*AdminResponse, error)
	AdminBlockUser(context.Context, *AdID) (*AdminResponse, error)
	FetchOrders(context.Context, *AdminNoParam) (*AdminOrderList, error)
	FetchOrderByUser(context.Context, *AdID) (*AdminOrderList, error)
	FetchOrder(context.Context, *AdID) (*AdminOrder, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) AdminLoginRequest(context.Context, *AdminLogin) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLoginRequest not implemented")
}
func (UnimplementedAdminServiceServer) AddToAdminWallet(context.Context, *Amount) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToAdminWallet not implemented")
}
func (UnimplementedAdminServiceServer) AddCategory(context.Context, *AdminCategory) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedAdminServiceServer) FindCategory(context.Context, *AdID) (*AdminCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCategory not implemented")
}
func (UnimplementedAdminServiceServer) FindCategories(context.Context, *AdminNoParam) (*AdminCategoryList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCategories not implemented")
}
func (UnimplementedAdminServiceServer) EditCategoryAdmin(context.Context, *AdminCategory) (*AdminCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCategoryAdmin not implemented")
}
func (UnimplementedAdminServiceServer) RemoveCategoryAdmin(context.Context, *AdID) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCategoryAdmin not implemented")
}
func (UnimplementedAdminServiceServer) FindProductByID(context.Context, *AdID) (*AdminProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindProductByID not implemented")
}
func (UnimplementedAdminServiceServer) FindAllProducts(context.Context, *AdminNoParam) (*AdminProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllProducts not implemented")
}
func (UnimplementedAdminServiceServer) RemoveProduct(context.Context, *AdID) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProduct not implemented")
}
func (UnimplementedAdminServiceServer) AdminBlockUser(context.Context, *AdID) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminBlockUser not implemented")
}
func (UnimplementedAdminServiceServer) FetchOrders(context.Context, *AdminNoParam) (*AdminOrderList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchOrders not implemented")
}
func (UnimplementedAdminServiceServer) FetchOrderByUser(context.Context, *AdID) (*AdminOrderList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchOrderByUser not implemented")
}
func (UnimplementedAdminServiceServer) FetchOrder(context.Context, *AdID) (*AdminOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchOrder not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_AdminLoginRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminLoginRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AdminLoginRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminLoginRequest(ctx, req.(*AdminLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AddToAdminWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Amount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AddToAdminWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AddToAdminWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AddToAdminWallet(ctx, req.(*Amount))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AddCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AddCategory(ctx, req.(*AdminCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_FindCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindCategory(ctx, req.(*AdID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_FindCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindCategories(ctx, req.(*AdminNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_EditCategoryAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).EditCategoryAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_EditCategoryAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).EditCategoryAdmin(ctx, req.(*AdminCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_RemoveCategoryAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).RemoveCategoryAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_RemoveCategoryAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).RemoveCategoryAdmin(ctx, req.(*AdID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_FindProductByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindProductByID(ctx, req.(*AdID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_FindAllProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindAllProducts(ctx, req.(*AdminNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_RemoveProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).RemoveProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_RemoveProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).RemoveProduct(ctx, req.(*AdID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminBlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminBlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AdminBlockUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminBlockUser(ctx, req.(*AdID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FetchOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FetchOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_FetchOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FetchOrders(ctx, req.(*AdminNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FetchOrderByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FetchOrderByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_FetchOrderByUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FetchOrderByUser(ctx, req.(*AdID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FetchOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FetchOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_FetchOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FetchOrder(ctx, req.(*AdID))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminLoginRequest",
			Handler:    _AdminService_AdminLoginRequest_Handler,
		},
		{
			MethodName: "AddToAdminWallet",
			Handler:    _AdminService_AddToAdminWallet_Handler,
		},
		{
			MethodName: "AddCategory",
			Handler:    _AdminService_AddCategory_Handler,
		},
		{
			MethodName: "FindCategory",
			Handler:    _AdminService_FindCategory_Handler,
		},
		{
			MethodName: "FindCategories",
			Handler:    _AdminService_FindCategories_Handler,
		},
		{
			MethodName: "EditCategoryAdmin",
			Handler:    _AdminService_EditCategoryAdmin_Handler,
		},
		{
			MethodName: "RemoveCategoryAdmin",
			Handler:    _AdminService_RemoveCategoryAdmin_Handler,
		},
		{
			MethodName: "FindProductByID",
			Handler:    _AdminService_FindProductByID_Handler,
		},
		{
			MethodName: "FindAllProducts",
			Handler:    _AdminService_FindAllProducts_Handler,
		},
		{
			MethodName: "RemoveProduct",
			Handler:    _AdminService_RemoveProduct_Handler,
		},
		{
			MethodName: "AdminBlockUser",
			Handler:    _AdminService_AdminBlockUser_Handler,
		},
		{
			MethodName: "FetchOrders",
			Handler:    _AdminService_FetchOrders_Handler,
		},
		{
			MethodName: "FetchOrderByUser",
			Handler:    _AdminService_FetchOrderByUser_Handler,
		},
		{
			MethodName: "FetchOrder",
			Handler:    _AdminService_FetchOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
