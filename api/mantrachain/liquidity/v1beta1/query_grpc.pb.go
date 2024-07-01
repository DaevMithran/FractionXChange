// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: mantrachain/liquidity/v1beta1/query.proto

package liquidityv1beta1

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
	Query_Params_FullMethodName                    = "/mantrachain.liquidity.v1beta1.Query/Params"
	Query_QueryPools_FullMethodName                = "/mantrachain.liquidity.v1beta1.Query/QueryPools"
	Query_QueryPool_FullMethodName                 = "/mantrachain.liquidity.v1beta1.Query/QueryPool"
	Query_QueryPoolByReserveAddress_FullMethodName = "/mantrachain.liquidity.v1beta1.Query/QueryPoolByReserveAddress"
	Query_QueryPoolByPoolCoinDenom_FullMethodName  = "/mantrachain.liquidity.v1beta1.Query/QueryPoolByPoolCoinDenom"
	Query_QueryPairsByDenoms_FullMethodName        = "/mantrachain.liquidity.v1beta1.Query/QueryPairsByDenoms"
	Query_QueryPairs_FullMethodName                = "/mantrachain.liquidity.v1beta1.Query/QueryPairs"
	Query_QueryPair_FullMethodName                 = "/mantrachain.liquidity.v1beta1.Query/QueryPair"
	Query_QueryDepositRequests_FullMethodName      = "/mantrachain.liquidity.v1beta1.Query/QueryDepositRequests"
	Query_QueryDepositRequest_FullMethodName       = "/mantrachain.liquidity.v1beta1.Query/QueryDepositRequest"
	Query_QueryWithdrawRequests_FullMethodName     = "/mantrachain.liquidity.v1beta1.Query/QueryWithdrawRequests"
	Query_QueryWithdrawRequest_FullMethodName      = "/mantrachain.liquidity.v1beta1.Query/QueryWithdrawRequest"
	Query_QueryOrders_FullMethodName               = "/mantrachain.liquidity.v1beta1.Query/QueryOrders"
	Query_QueryOrder_FullMethodName                = "/mantrachain.liquidity.v1beta1.Query/QueryOrder"
	Query_QueryOrdersByOrderer_FullMethodName      = "/mantrachain.liquidity.v1beta1.Query/QueryOrdersByOrderer"
	Query_QueryOrderBooks_FullMethodName           = "/mantrachain.liquidity.v1beta1.Query/QueryOrderBooks"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Pools returns all liquidity pools.
	QueryPools(ctx context.Context, in *QueryPoolsRequest, opts ...grpc.CallOption) (*QueryPoolsResponse, error)
	// Pool returns the specific liquidity pool.
	QueryPool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error)
	// PoolByReserveAddress returns all pools that correspond to the reserve account.
	QueryPoolByReserveAddress(ctx context.Context, in *QueryPoolByReserveAddressRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error)
	// PoolByPoolCoinDenom returns all pools that correspond to the pool coin denom.
	QueryPoolByPoolCoinDenom(ctx context.Context, in *QueryPoolByPoolCoinDenomRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error)
	// Pairs returns all liquidity pairs by denoms.
	QueryPairsByDenoms(ctx context.Context, in *QueryPairsByDenomsRequest, opts ...grpc.CallOption) (*QueryPairsByDenomsResponse, error)
	// Pairs returns all liquidity pairs.
	QueryPairs(ctx context.Context, in *QueryPairsRequest, opts ...grpc.CallOption) (*QueryPairsResponse, error)
	// Pair returns the specific pair.
	QueryPair(ctx context.Context, in *QueryPairRequest, opts ...grpc.CallOption) (*QueryPairResponse, error)
	// DepositRequests returns all deposit requests.
	QueryDepositRequests(ctx context.Context, in *QueryDepositRequestsRequest, opts ...grpc.CallOption) (*QueryDepositRequestsResponse, error)
	// DepositRequest returns the specific deposit request.
	QueryDepositRequest(ctx context.Context, in *QueryDepositRequestRequest, opts ...grpc.CallOption) (*QueryDepositRequestResponse, error)
	// WithdrawRequests returns all withdraw requests.
	QueryWithdrawRequests(ctx context.Context, in *QueryWithdrawRequestsRequest, opts ...grpc.CallOption) (*QueryWithdrawRequestsResponse, error)
	// WithdrawRequest returns the specific withdraw request.
	QueryWithdrawRequest(ctx context.Context, in *QueryWithdrawRequestRequest, opts ...grpc.CallOption) (*QueryWithdrawRequestResponse, error)
	// Orders returns all orders within the pair.
	QueryOrders(ctx context.Context, in *QueryOrdersRequest, opts ...grpc.CallOption) (*QueryOrdersResponse, error)
	// Order returns the specific order.
	QueryOrder(ctx context.Context, in *QueryOrderRequest, opts ...grpc.CallOption) (*QueryOrderResponse, error)
	// OrdersByOrderer returns orders made by an orderer.
	QueryOrdersByOrderer(ctx context.Context, in *QueryOrdersByOrdererRequest, opts ...grpc.CallOption) (*QueryOrdersResponse, error)
	QueryOrderBooks(ctx context.Context, in *QueryOrderBooksRequest, opts ...grpc.CallOption) (*QueryOrderBooksResponse, error)
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

func (c *queryClient) QueryPools(ctx context.Context, in *QueryPoolsRequest, opts ...grpc.CallOption) (*QueryPoolsResponse, error) {
	out := new(QueryPoolsResponse)
	err := c.cc.Invoke(ctx, Query_QueryPools_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryPool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error) {
	out := new(QueryPoolResponse)
	err := c.cc.Invoke(ctx, Query_QueryPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryPoolByReserveAddress(ctx context.Context, in *QueryPoolByReserveAddressRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error) {
	out := new(QueryPoolResponse)
	err := c.cc.Invoke(ctx, Query_QueryPoolByReserveAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryPoolByPoolCoinDenom(ctx context.Context, in *QueryPoolByPoolCoinDenomRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error) {
	out := new(QueryPoolResponse)
	err := c.cc.Invoke(ctx, Query_QueryPoolByPoolCoinDenom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryPairsByDenoms(ctx context.Context, in *QueryPairsByDenomsRequest, opts ...grpc.CallOption) (*QueryPairsByDenomsResponse, error) {
	out := new(QueryPairsByDenomsResponse)
	err := c.cc.Invoke(ctx, Query_QueryPairsByDenoms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryPairs(ctx context.Context, in *QueryPairsRequest, opts ...grpc.CallOption) (*QueryPairsResponse, error) {
	out := new(QueryPairsResponse)
	err := c.cc.Invoke(ctx, Query_QueryPairs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryPair(ctx context.Context, in *QueryPairRequest, opts ...grpc.CallOption) (*QueryPairResponse, error) {
	out := new(QueryPairResponse)
	err := c.cc.Invoke(ctx, Query_QueryPair_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryDepositRequests(ctx context.Context, in *QueryDepositRequestsRequest, opts ...grpc.CallOption) (*QueryDepositRequestsResponse, error) {
	out := new(QueryDepositRequestsResponse)
	err := c.cc.Invoke(ctx, Query_QueryDepositRequests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryDepositRequest(ctx context.Context, in *QueryDepositRequestRequest, opts ...grpc.CallOption) (*QueryDepositRequestResponse, error) {
	out := new(QueryDepositRequestResponse)
	err := c.cc.Invoke(ctx, Query_QueryDepositRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryWithdrawRequests(ctx context.Context, in *QueryWithdrawRequestsRequest, opts ...grpc.CallOption) (*QueryWithdrawRequestsResponse, error) {
	out := new(QueryWithdrawRequestsResponse)
	err := c.cc.Invoke(ctx, Query_QueryWithdrawRequests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryWithdrawRequest(ctx context.Context, in *QueryWithdrawRequestRequest, opts ...grpc.CallOption) (*QueryWithdrawRequestResponse, error) {
	out := new(QueryWithdrawRequestResponse)
	err := c.cc.Invoke(ctx, Query_QueryWithdrawRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryOrders(ctx context.Context, in *QueryOrdersRequest, opts ...grpc.CallOption) (*QueryOrdersResponse, error) {
	out := new(QueryOrdersResponse)
	err := c.cc.Invoke(ctx, Query_QueryOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryOrder(ctx context.Context, in *QueryOrderRequest, opts ...grpc.CallOption) (*QueryOrderResponse, error) {
	out := new(QueryOrderResponse)
	err := c.cc.Invoke(ctx, Query_QueryOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryOrdersByOrderer(ctx context.Context, in *QueryOrdersByOrdererRequest, opts ...grpc.CallOption) (*QueryOrdersResponse, error) {
	out := new(QueryOrdersResponse)
	err := c.cc.Invoke(ctx, Query_QueryOrdersByOrderer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryOrderBooks(ctx context.Context, in *QueryOrderBooksRequest, opts ...grpc.CallOption) (*QueryOrderBooksResponse, error) {
	out := new(QueryOrderBooksResponse)
	err := c.cc.Invoke(ctx, Query_QueryOrderBooks_FullMethodName, in, out, opts...)
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
	// Pools returns all liquidity pools.
	QueryPools(context.Context, *QueryPoolsRequest) (*QueryPoolsResponse, error)
	// Pool returns the specific liquidity pool.
	QueryPool(context.Context, *QueryPoolRequest) (*QueryPoolResponse, error)
	// PoolByReserveAddress returns all pools that correspond to the reserve account.
	QueryPoolByReserveAddress(context.Context, *QueryPoolByReserveAddressRequest) (*QueryPoolResponse, error)
	// PoolByPoolCoinDenom returns all pools that correspond to the pool coin denom.
	QueryPoolByPoolCoinDenom(context.Context, *QueryPoolByPoolCoinDenomRequest) (*QueryPoolResponse, error)
	// Pairs returns all liquidity pairs by denoms.
	QueryPairsByDenoms(context.Context, *QueryPairsByDenomsRequest) (*QueryPairsByDenomsResponse, error)
	// Pairs returns all liquidity pairs.
	QueryPairs(context.Context, *QueryPairsRequest) (*QueryPairsResponse, error)
	// Pair returns the specific pair.
	QueryPair(context.Context, *QueryPairRequest) (*QueryPairResponse, error)
	// DepositRequests returns all deposit requests.
	QueryDepositRequests(context.Context, *QueryDepositRequestsRequest) (*QueryDepositRequestsResponse, error)
	// DepositRequest returns the specific deposit request.
	QueryDepositRequest(context.Context, *QueryDepositRequestRequest) (*QueryDepositRequestResponse, error)
	// WithdrawRequests returns all withdraw requests.
	QueryWithdrawRequests(context.Context, *QueryWithdrawRequestsRequest) (*QueryWithdrawRequestsResponse, error)
	// WithdrawRequest returns the specific withdraw request.
	QueryWithdrawRequest(context.Context, *QueryWithdrawRequestRequest) (*QueryWithdrawRequestResponse, error)
	// Orders returns all orders within the pair.
	QueryOrders(context.Context, *QueryOrdersRequest) (*QueryOrdersResponse, error)
	// Order returns the specific order.
	QueryOrder(context.Context, *QueryOrderRequest) (*QueryOrderResponse, error)
	// OrdersByOrderer returns orders made by an orderer.
	QueryOrdersByOrderer(context.Context, *QueryOrdersByOrdererRequest) (*QueryOrdersResponse, error)
	QueryOrderBooks(context.Context, *QueryOrderBooksRequest) (*QueryOrderBooksResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) QueryPools(context.Context, *QueryPoolsRequest) (*QueryPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPools not implemented")
}
func (UnimplementedQueryServer) QueryPool(context.Context, *QueryPoolRequest) (*QueryPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPool not implemented")
}
func (UnimplementedQueryServer) QueryPoolByReserveAddress(context.Context, *QueryPoolByReserveAddressRequest) (*QueryPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPoolByReserveAddress not implemented")
}
func (UnimplementedQueryServer) QueryPoolByPoolCoinDenom(context.Context, *QueryPoolByPoolCoinDenomRequest) (*QueryPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPoolByPoolCoinDenom not implemented")
}
func (UnimplementedQueryServer) QueryPairsByDenoms(context.Context, *QueryPairsByDenomsRequest) (*QueryPairsByDenomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPairsByDenoms not implemented")
}
func (UnimplementedQueryServer) QueryPairs(context.Context, *QueryPairsRequest) (*QueryPairsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPairs not implemented")
}
func (UnimplementedQueryServer) QueryPair(context.Context, *QueryPairRequest) (*QueryPairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPair not implemented")
}
func (UnimplementedQueryServer) QueryDepositRequests(context.Context, *QueryDepositRequestsRequest) (*QueryDepositRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDepositRequests not implemented")
}
func (UnimplementedQueryServer) QueryDepositRequest(context.Context, *QueryDepositRequestRequest) (*QueryDepositRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDepositRequest not implemented")
}
func (UnimplementedQueryServer) QueryWithdrawRequests(context.Context, *QueryWithdrawRequestsRequest) (*QueryWithdrawRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryWithdrawRequests not implemented")
}
func (UnimplementedQueryServer) QueryWithdrawRequest(context.Context, *QueryWithdrawRequestRequest) (*QueryWithdrawRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryWithdrawRequest not implemented")
}
func (UnimplementedQueryServer) QueryOrders(context.Context, *QueryOrdersRequest) (*QueryOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrders not implemented")
}
func (UnimplementedQueryServer) QueryOrder(context.Context, *QueryOrderRequest) (*QueryOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrder not implemented")
}
func (UnimplementedQueryServer) QueryOrdersByOrderer(context.Context, *QueryOrdersByOrdererRequest) (*QueryOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrdersByOrderer not implemented")
}
func (UnimplementedQueryServer) QueryOrderBooks(context.Context, *QueryOrderBooksRequest) (*QueryOrderBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrderBooks not implemented")
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

func _Query_QueryPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryPools_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPools(ctx, req.(*QueryPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPool(ctx, req.(*QueryPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryPoolByReserveAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPoolByReserveAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPoolByReserveAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryPoolByReserveAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPoolByReserveAddress(ctx, req.(*QueryPoolByReserveAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryPoolByPoolCoinDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPoolByPoolCoinDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPoolByPoolCoinDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryPoolByPoolCoinDenom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPoolByPoolCoinDenom(ctx, req.(*QueryPoolByPoolCoinDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryPairsByDenoms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPairsByDenomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPairsByDenoms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryPairsByDenoms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPairsByDenoms(ctx, req.(*QueryPairsByDenomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPairsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryPairs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPairs(ctx, req.(*QueryPairsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryPair_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPair(ctx, req.(*QueryPairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryDepositRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDepositRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryDepositRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryDepositRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryDepositRequests(ctx, req.(*QueryDepositRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryDepositRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDepositRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryDepositRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryDepositRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryDepositRequest(ctx, req.(*QueryDepositRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryWithdrawRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWithdrawRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryWithdrawRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryWithdrawRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryWithdrawRequests(ctx, req.(*QueryWithdrawRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryWithdrawRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWithdrawRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryWithdrawRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryWithdrawRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryWithdrawRequest(ctx, req.(*QueryWithdrawRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryOrders(ctx, req.(*QueryOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryOrder(ctx, req.(*QueryOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryOrdersByOrderer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOrdersByOrdererRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryOrdersByOrderer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryOrdersByOrderer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryOrdersByOrderer(ctx, req.(*QueryOrdersByOrdererRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryOrderBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOrderBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryOrderBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryOrderBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryOrderBooks(ctx, req.(*QueryOrderBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mantrachain.liquidity.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "QueryPools",
			Handler:    _Query_QueryPools_Handler,
		},
		{
			MethodName: "QueryPool",
			Handler:    _Query_QueryPool_Handler,
		},
		{
			MethodName: "QueryPoolByReserveAddress",
			Handler:    _Query_QueryPoolByReserveAddress_Handler,
		},
		{
			MethodName: "QueryPoolByPoolCoinDenom",
			Handler:    _Query_QueryPoolByPoolCoinDenom_Handler,
		},
		{
			MethodName: "QueryPairsByDenoms",
			Handler:    _Query_QueryPairsByDenoms_Handler,
		},
		{
			MethodName: "QueryPairs",
			Handler:    _Query_QueryPairs_Handler,
		},
		{
			MethodName: "QueryPair",
			Handler:    _Query_QueryPair_Handler,
		},
		{
			MethodName: "QueryDepositRequests",
			Handler:    _Query_QueryDepositRequests_Handler,
		},
		{
			MethodName: "QueryDepositRequest",
			Handler:    _Query_QueryDepositRequest_Handler,
		},
		{
			MethodName: "QueryWithdrawRequests",
			Handler:    _Query_QueryWithdrawRequests_Handler,
		},
		{
			MethodName: "QueryWithdrawRequest",
			Handler:    _Query_QueryWithdrawRequest_Handler,
		},
		{
			MethodName: "QueryOrders",
			Handler:    _Query_QueryOrders_Handler,
		},
		{
			MethodName: "QueryOrder",
			Handler:    _Query_QueryOrder_Handler,
		},
		{
			MethodName: "QueryOrdersByOrderer",
			Handler:    _Query_QueryOrdersByOrderer_Handler,
		},
		{
			MethodName: "QueryOrderBooks",
			Handler:    _Query_QueryOrderBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mantrachain/liquidity/v1beta1/query.proto",
}
