// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: package.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPCClients-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PackageService_GetPackagesWithZones_FullMethodName = "/pb.PackageService/GetPackagesWithZones"
)

// PackageServiceClient is the client API for PackageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PackageServiceClient interface {
	GetPackagesWithZones(ctx context.Context, in *GetPackagesWithZonesRequest, opts ...grpc.CallOption) (*GetPackagesWithZonesResponse, error)
}

type packageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPackageServiceClient(cc grpc.ClientConnInterface) PackageServiceClient {
	return &packageServiceClient{cc}
}

func (c *packageServiceClient) GetPackagesWithZones(ctx context.Context, in *GetPackagesWithZonesRequest, opts ...grpc.CallOption) (*GetPackagesWithZonesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPackagesWithZonesResponse)
	err := c.cc.Invoke(ctx, PackageService_GetPackagesWithZones_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PackageServiceServer is the app API for PackageService service.
// All implementations must embed UnimplementedPackageServiceServer
// for forward compatibility.
type PackageServiceServer interface {
	GetPackagesWithZones(context.Context, *GetPackagesWithZonesRequest) (*GetPackagesWithZonesResponse, error)
	mustEmbedUnimplementedPackageServiceServer()
}

// UnimplementedPackageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPackageServiceServer struct{}

func (UnimplementedPackageServiceServer) GetPackagesWithZones(context.Context, *GetPackagesWithZonesRequest) (*GetPackagesWithZonesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackagesWithZones not implemented")
}
func (UnimplementedPackageServiceServer) mustEmbedUnimplementedPackageServiceServer() {}
func (UnimplementedPackageServiceServer) testEmbeddedByValue()                        {}

// UnsafePackageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PackageServiceServer will
// result in compilation errors.
type UnsafePackageServiceServer interface {
	mustEmbedUnimplementedPackageServiceServer()
}

func RegisterPackageServiceServer(s grpc.ServiceRegistrar, srv PackageServiceServer) {
	// If the following call pancis, it indicates UnimplementedPackageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PackageService_ServiceDesc, srv)
}

func _PackageService_GetPackagesWithZones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPackagesWithZonesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).GetPackagesWithZones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PackageService_GetPackagesWithZones_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).GetPackagesWithZones(ctx, req.(*GetPackagesWithZonesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PackageService_ServiceDesc is the grpc.ServiceDesc for PackageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PackageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PackageService",
	HandlerType: (*PackageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPackagesWithZones",
			Handler:    _PackageService_GetPackagesWithZones_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "package.proto",
}
