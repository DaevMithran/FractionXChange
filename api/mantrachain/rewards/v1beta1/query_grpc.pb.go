// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: mantrachain/rewards/v1beta1/query.proto

package rewardsv1beta1

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
	Query_Params_FullMethodName             = "/mantrachain.rewards.v1beta1.Query/Params"
	Query_QuerySnapshot_FullMethodName      = "/mantrachain.rewards.v1beta1.Query/QuerySnapshot"
	Query_QuerySnapshotAll_FullMethodName   = "/mantrachain.rewards.v1beta1.Query/QuerySnapshotAll"
	Query_QueryProvider_FullMethodName      = "/mantrachain.rewards.v1beta1.Query/QueryProvider"
	Query_QueryProviderPairs_FullMethodName = "/mantrachain.rewards.v1beta1.Query/QueryProviderPairs"
	Query_QueryProviderAll_FullMethodName   = "/mantrachain.rewards.v1beta1.Query/QueryProviderAll"
	Query_QueryRewards_FullMethodName       = "/mantrachain.rewards.v1beta1.Query/QueryRewards"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Snapshot items.
	QuerySnapshot(ctx context.Context, in *QueryGetSnapshotRequest, opts ...grpc.CallOption) (*QueryGetSnapshotResponse, error)
	// QuerySnapshotAll
	QuerySnapshotAll(ctx context.Context, in *QueryAllSnapshotRequest, opts ...grpc.CallOption) (*QueryAllSnapshotResponse, error)
	// Queries a list of Provider pairs ids.
	QueryProvider(ctx context.Context, in *QueryGetProviderRequest, opts ...grpc.CallOption) (*QueryGetProviderResponse, error)
	// QueryProviderPairs
	QueryProviderPairs(ctx context.Context, in *QueryGetProviderPairsRequest, opts ...grpc.CallOption) (*QueryGetProviderPairsResponse, error)
	// QueryProviderAll
	QueryProviderAll(ctx context.Context, in *QueryAllProviderRequest, opts ...grpc.CallOption) (*QueryAllProviderResponse, error)
	// Queries a list of Claim items.
	QueryRewards(ctx context.Context, in *QueryGetRewardsRequest, opts ...grpc.CallOption) (*QueryGetRewardsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QuerySnapshot(ctx context.Context, in *QueryGetSnapshotRequest, opts ...grpc.CallOption) (*QueryGetSnapshotResponse, error) {
	out := new(QueryGetSnapshotResponse)
	err := c.cc.Invoke(ctx, Query_QuerySnapshot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QuerySnapshotAll(ctx context.Context, in *QueryAllSnapshotRequest, opts ...grpc.CallOption) (*QueryAllSnapshotResponse, error) {
	out := new(QueryAllSnapshotResponse)
	err := c.cc.Invoke(ctx, Query_QuerySnapshotAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryProvider(ctx context.Context, in *QueryGetProviderRequest, opts ...grpc.CallOption) (*QueryGetProviderResponse, error) {
	out := new(QueryGetProviderResponse)
	err := c.cc.Invoke(ctx, Query_QueryProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryProviderPairs(ctx context.Context, in *QueryGetProviderPairsRequest, opts ...grpc.CallOption) (*QueryGetProviderPairsResponse, error) {
	out := new(QueryGetProviderPairsResponse)
	err := c.cc.Invoke(ctx, Query_QueryProviderPairs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryProviderAll(ctx context.Context, in *QueryAllProviderRequest, opts ...grpc.CallOption) (*QueryAllProviderResponse, error) {
	out := new(QueryAllProviderResponse)
	err := c.cc.Invoke(ctx, Query_QueryProviderAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryRewards(ctx context.Context, in *QueryGetRewardsRequest, opts ...grpc.CallOption) (*QueryGetRewardsResponse, error) {
	out := new(QueryGetRewardsResponse)
	err := c.cc.Invoke(ctx, Query_QueryRewards_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Snapshot items.
	QuerySnapshot(context.Context, *QueryGetSnapshotRequest) (*QueryGetSnapshotResponse, error)
	// QuerySnapshotAll
	QuerySnapshotAll(context.Context, *QueryAllSnapshotRequest) (*QueryAllSnapshotResponse, error)
	// Queries a list of Provider pairs ids.
	QueryProvider(context.Context, *QueryGetProviderRequest) (*QueryGetProviderResponse, error)
	// QueryProviderPairs
	QueryProviderPairs(context.Context, *QueryGetProviderPairsRequest) (*QueryGetProviderPairsResponse, error)
	// QueryProviderAll
	QueryProviderAll(context.Context, *QueryAllProviderRequest) (*QueryAllProviderResponse, error)
	// Queries a list of Claim items.
	QueryRewards(context.Context, *QueryGetRewardsRequest) (*QueryGetRewardsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) QuerySnapshot(context.Context, *QueryGetSnapshotRequest) (*QueryGetSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySnapshot not implemented")
}
func (UnimplementedQueryServer) QuerySnapshotAll(context.Context, *QueryAllSnapshotRequest) (*QueryAllSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySnapshotAll not implemented")
}
func (UnimplementedQueryServer) QueryProvider(context.Context, *QueryGetProviderRequest) (*QueryGetProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProvider not implemented")
}
func (UnimplementedQueryServer) QueryProviderPairs(context.Context, *QueryGetProviderPairsRequest) (*QueryGetProviderPairsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProviderPairs not implemented")
}
func (UnimplementedQueryServer) QueryProviderAll(context.Context, *QueryAllProviderRequest) (*QueryAllProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProviderAll not implemented")
}
func (UnimplementedQueryServer) QueryRewards(context.Context, *QueryGetRewardsRequest) (*QueryGetRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRewards not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QuerySnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QuerySnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QuerySnapshot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QuerySnapshot(ctx, req.(*QueryGetSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QuerySnapshotAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QuerySnapshotAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QuerySnapshotAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QuerySnapshotAll(ctx, req.(*QueryAllSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryProvider(ctx, req.(*QueryGetProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryProviderPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetProviderPairsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryProviderPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryProviderPairs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryProviderPairs(ctx, req.(*QueryGetProviderPairsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryProviderAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryProviderAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryProviderAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryProviderAll(ctx, req.(*QueryAllProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryRewards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryRewards(ctx, req.(*QueryGetRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mantrachain.rewards.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "QuerySnapshot",
			Handler:    _Query_QuerySnapshot_Handler,
		},
		{
			MethodName: "QuerySnapshotAll",
			Handler:    _Query_QuerySnapshotAll_Handler,
		},
		{
			MethodName: "QueryProvider",
			Handler:    _Query_QueryProvider_Handler,
		},
		{
			MethodName: "QueryProviderPairs",
			Handler:    _Query_QueryProviderPairs_Handler,
		},
		{
			MethodName: "QueryProviderAll",
			Handler:    _Query_QueryProviderAll_Handler,
		},
		{
			MethodName: "QueryRewards",
			Handler:    _Query_QueryRewards_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mantrachain/rewards/v1beta1/query.proto",
}
