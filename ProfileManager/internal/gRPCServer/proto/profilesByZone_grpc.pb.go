// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: profilesByZone.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProfilesByZoneService_GetProfilesByZone_FullMethodName = "/pb.ProfilesByZoneService/GetProfilesByZone"
)

// ProfilesByZoneServiceClient is the client API for ProfilesByZoneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfilesByZoneServiceClient interface {
	GetProfilesByZone(ctx context.Context, in *GetProfileByZoneRequest, opts ...grpc.CallOption) (*GetProfilesByZoneResponse, error)
}

type profilesByZoneServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilesByZoneServiceClient(cc grpc.ClientConnInterface) ProfilesByZoneServiceClient {
	return &profilesByZoneServiceClient{cc}
}

func (c *profilesByZoneServiceClient) GetProfilesByZone(ctx context.Context, in *GetProfileByZoneRequest, opts ...grpc.CallOption) (*GetProfilesByZoneResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfilesByZoneResponse)
	err := c.cc.Invoke(ctx, ProfilesByZoneService_GetProfilesByZone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilesByZoneServiceServer is the server API for ProfilesByZoneService service.
// All implementations must embed UnimplementedProfilesByZoneServiceServer
// for forward compatibility.
type ProfilesByZoneServiceServer interface {
	GetProfilesByZone(context.Context, *GetProfileByZoneRequest) (*GetProfilesByZoneResponse, error)
	mustEmbedUnimplementedProfilesByZoneServiceServer()
}

// UnimplementedProfilesByZoneServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProfilesByZoneServiceServer struct{}

func (UnimplementedProfilesByZoneServiceServer) GetProfilesByZone(context.Context, *GetProfileByZoneRequest) (*GetProfilesByZoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfilesByZone not implemented")
}
func (UnimplementedProfilesByZoneServiceServer) mustEmbedUnimplementedProfilesByZoneServiceServer() {}
func (UnimplementedProfilesByZoneServiceServer) testEmbeddedByValue()                               {}

// UnsafeProfilesByZoneServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfilesByZoneServiceServer will
// result in compilation errors.
type UnsafeProfilesByZoneServiceServer interface {
	mustEmbedUnimplementedProfilesByZoneServiceServer()
}

func RegisterProfilesByZoneServiceServer(s grpc.ServiceRegistrar, srv ProfilesByZoneServiceServer) {
	// If the following call pancis, it indicates UnimplementedProfilesByZoneServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProfilesByZoneService_ServiceDesc, srv)
}

func _ProfilesByZoneService_GetProfilesByZone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileByZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesByZoneServiceServer).GetProfilesByZone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfilesByZoneService_GetProfilesByZone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesByZoneServiceServer).GetProfilesByZone(ctx, req.(*GetProfileByZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfilesByZoneService_ServiceDesc is the grpc.ServiceDesc for ProfilesByZoneService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfilesByZoneService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProfilesByZoneService",
	HandlerType: (*ProfilesByZoneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfilesByZone",
			Handler:    _ProfilesByZoneService_GetProfilesByZone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profilesByZone.proto",
}
