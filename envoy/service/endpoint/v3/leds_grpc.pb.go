// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: envoy/service/endpoint/v3/leds.proto

package endpointv3

import (
	context "context"
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	LocalityEndpointDiscoveryService_DeltaLocalityEndpoints_FullMethodName = "/envoy.service.endpoint.v3.LocalityEndpointDiscoveryService/DeltaLocalityEndpoints"
)

// LocalityEndpointDiscoveryServiceClient is the client API for LocalityEndpointDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocalityEndpointDiscoveryServiceClient interface {
	// The resource_names_subscribe resource_names_unsubscribe fields in DeltaDiscoveryRequest
	// specify a list of glob collections to subscribe to updates for.
	DeltaLocalityEndpoints(ctx context.Context, opts ...grpc.CallOption) (LocalityEndpointDiscoveryService_DeltaLocalityEndpointsClient, error)
}

type localityEndpointDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalityEndpointDiscoveryServiceClient(cc grpc.ClientConnInterface) LocalityEndpointDiscoveryServiceClient {
	return &localityEndpointDiscoveryServiceClient{cc}
}

func (c *localityEndpointDiscoveryServiceClient) DeltaLocalityEndpoints(ctx context.Context, opts ...grpc.CallOption) (LocalityEndpointDiscoveryService_DeltaLocalityEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &LocalityEndpointDiscoveryService_ServiceDesc.Streams[0], LocalityEndpointDiscoveryService_DeltaLocalityEndpoints_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &localityEndpointDiscoveryServiceDeltaLocalityEndpointsClient{stream}
	return x, nil
}

type LocalityEndpointDiscoveryService_DeltaLocalityEndpointsClient interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type localityEndpointDiscoveryServiceDeltaLocalityEndpointsClient struct {
	grpc.ClientStream
}

func (x *localityEndpointDiscoveryServiceDeltaLocalityEndpointsClient) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *localityEndpointDiscoveryServiceDeltaLocalityEndpointsClient) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LocalityEndpointDiscoveryServiceServer is the server API for LocalityEndpointDiscoveryService service.
// All implementations should embed UnimplementedLocalityEndpointDiscoveryServiceServer
// for forward compatibility
type LocalityEndpointDiscoveryServiceServer interface {
	// The resource_names_subscribe resource_names_unsubscribe fields in DeltaDiscoveryRequest
	// specify a list of glob collections to subscribe to updates for.
	DeltaLocalityEndpoints(LocalityEndpointDiscoveryService_DeltaLocalityEndpointsServer) error
}

// UnimplementedLocalityEndpointDiscoveryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLocalityEndpointDiscoveryServiceServer struct {
}

func (UnimplementedLocalityEndpointDiscoveryServiceServer) DeltaLocalityEndpoints(LocalityEndpointDiscoveryService_DeltaLocalityEndpointsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaLocalityEndpoints not implemented")
}

// UnsafeLocalityEndpointDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocalityEndpointDiscoveryServiceServer will
// result in compilation errors.
type UnsafeLocalityEndpointDiscoveryServiceServer interface {
	mustEmbedUnimplementedLocalityEndpointDiscoveryServiceServer()
}

func RegisterLocalityEndpointDiscoveryServiceServer(s grpc.ServiceRegistrar, srv LocalityEndpointDiscoveryServiceServer) {
	s.RegisterService(&LocalityEndpointDiscoveryService_ServiceDesc, srv)
}

func _LocalityEndpointDiscoveryService_DeltaLocalityEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LocalityEndpointDiscoveryServiceServer).DeltaLocalityEndpoints(&localityEndpointDiscoveryServiceDeltaLocalityEndpointsServer{stream})
}

type LocalityEndpointDiscoveryService_DeltaLocalityEndpointsServer interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type localityEndpointDiscoveryServiceDeltaLocalityEndpointsServer struct {
	grpc.ServerStream
}

func (x *localityEndpointDiscoveryServiceDeltaLocalityEndpointsServer) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *localityEndpointDiscoveryServiceDeltaLocalityEndpointsServer) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LocalityEndpointDiscoveryService_ServiceDesc is the grpc.ServiceDesc for LocalityEndpointDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocalityEndpointDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.endpoint.v3.LocalityEndpointDiscoveryService",
	HandlerType: (*LocalityEndpointDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaLocalityEndpoints",
			Handler:       _LocalityEndpointDiscoveryService_DeltaLocalityEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/endpoint/v3/leds.proto",
}
