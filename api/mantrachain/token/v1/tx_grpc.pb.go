// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: mantrachain/token/v1/tx.proto

package tokenv1

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
	Msg_UpdateParams_FullMethodName                                   = "/mantrachain.token.v1.Msg/UpdateParams"
	Msg_CreateNftCollection_FullMethodName                            = "/mantrachain.token.v1.Msg/CreateNftCollection"
	Msg_MintNfts_FullMethodName                                       = "/mantrachain.token.v1.Msg/MintNfts"
	Msg_BurnNfts_FullMethodName                                       = "/mantrachain.token.v1.Msg/BurnNfts"
	Msg_TransferNfts_FullMethodName                                   = "/mantrachain.token.v1.Msg/TransferNfts"
	Msg_ApproveNfts_FullMethodName                                    = "/mantrachain.token.v1.Msg/ApproveNfts"
	Msg_ApproveAllNfts_FullMethodName                                 = "/mantrachain.token.v1.Msg/ApproveAllNfts"
	Msg_MintNft_FullMethodName                                        = "/mantrachain.token.v1.Msg/MintNft"
	Msg_BurnNft_FullMethodName                                        = "/mantrachain.token.v1.Msg/BurnNft"
	Msg_TransferNft_FullMethodName                                    = "/mantrachain.token.v1.Msg/TransferNft"
	Msg_ApproveNft_FullMethodName                                     = "/mantrachain.token.v1.Msg/ApproveNft"
	Msg_UpdateGuardSoulBondNftImage_FullMethodName                    = "/mantrachain.token.v1.Msg/UpdateGuardSoulBondNftImage"
	Msg_UpdateRestrictedCollectionNftImage_FullMethodName             = "/mantrachain.token.v1.Msg/UpdateRestrictedCollectionNftImage"
	Msg_UpdateRestrictedCollectionNftImageBatch_FullMethodName        = "/mantrachain.token.v1.Msg/UpdateRestrictedCollectionNftImageBatch"
	Msg_UpdateRestrictedCollectionNftImageGroupedBatch_FullMethodName = "/mantrachain.token.v1.Msg/UpdateRestrictedCollectionNftImageGroupedBatch"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// CreateNftCollection
	CreateNftCollection(ctx context.Context, in *MsgCreateNftCollection, opts ...grpc.CallOption) (*MsgCreateNftCollectionResponse, error)
	// MintNfts
	MintNfts(ctx context.Context, in *MsgMintNfts, opts ...grpc.CallOption) (*MsgMintNftsResponse, error)
	// BurnNfts
	BurnNfts(ctx context.Context, in *MsgBurnNfts, opts ...grpc.CallOption) (*MsgBurnNftsResponse, error)
	// TransferNfts
	TransferNfts(ctx context.Context, in *MsgTransferNfts, opts ...grpc.CallOption) (*MsgTransferNftsResponse, error)
	// ApproveNfts
	ApproveNfts(ctx context.Context, in *MsgApproveNfts, opts ...grpc.CallOption) (*MsgApproveNftsResponse, error)
	// ApproveAllNfts
	ApproveAllNfts(ctx context.Context, in *MsgApproveAllNfts, opts ...grpc.CallOption) (*MsgApproveAllNftsResponse, error)
	// MintNft
	MintNft(ctx context.Context, in *MsgMintNft, opts ...grpc.CallOption) (*MsgMintNftResponse, error)
	// BurnNft
	BurnNft(ctx context.Context, in *MsgBurnNft, opts ...grpc.CallOption) (*MsgBurnNftResponse, error)
	// TransferNft
	TransferNft(ctx context.Context, in *MsgTransferNft, opts ...grpc.CallOption) (*MsgTransferNftResponse, error)
	// ApproveNft
	ApproveNft(ctx context.Context, in *MsgApproveNft, opts ...grpc.CallOption) (*MsgApproveNftResponse, error)
	// UpdateGuardSoulBondNftImage
	UpdateGuardSoulBondNftImage(ctx context.Context, in *MsgUpdateGuardSoulBondNftImage, opts ...grpc.CallOption) (*MsgUpdateGuardSoulBondNftImageResponse, error)
	// UpdateRestrictedCollectionNftImage
	UpdateRestrictedCollectionNftImage(ctx context.Context, in *MsgUpdateRestrictedCollectionNftImage, opts ...grpc.CallOption) (*MsgUpdateRestrictedCollectionNftImageResponse, error)
	// UpdateRestrictedCollectionNftImageBatch
	UpdateRestrictedCollectionNftImageBatch(ctx context.Context, in *MsgUpdateRestrictedCollectionNftImageBatch, opts ...grpc.CallOption) (*MsgUpdateRestrictedCollectionNftImageBatchResponse, error)
	// UpdateRestrictedCollectionNftImageGroupedBatch
	UpdateRestrictedCollectionNftImageGroupedBatch(ctx context.Context, in *MsgUpdateRestrictedCollectionNftImageGroupedBatch, opts ...grpc.CallOption) (*MsgUpdateRestrictedCollectionNftImageGroupedBatchResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateNftCollection(ctx context.Context, in *MsgCreateNftCollection, opts ...grpc.CallOption) (*MsgCreateNftCollectionResponse, error) {
	out := new(MsgCreateNftCollectionResponse)
	err := c.cc.Invoke(ctx, Msg_CreateNftCollection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MintNfts(ctx context.Context, in *MsgMintNfts, opts ...grpc.CallOption) (*MsgMintNftsResponse, error) {
	out := new(MsgMintNftsResponse)
	err := c.cc.Invoke(ctx, Msg_MintNfts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BurnNfts(ctx context.Context, in *MsgBurnNfts, opts ...grpc.CallOption) (*MsgBurnNftsResponse, error) {
	out := new(MsgBurnNftsResponse)
	err := c.cc.Invoke(ctx, Msg_BurnNfts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferNfts(ctx context.Context, in *MsgTransferNfts, opts ...grpc.CallOption) (*MsgTransferNftsResponse, error) {
	out := new(MsgTransferNftsResponse)
	err := c.cc.Invoke(ctx, Msg_TransferNfts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ApproveNfts(ctx context.Context, in *MsgApproveNfts, opts ...grpc.CallOption) (*MsgApproveNftsResponse, error) {
	out := new(MsgApproveNftsResponse)
	err := c.cc.Invoke(ctx, Msg_ApproveNfts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ApproveAllNfts(ctx context.Context, in *MsgApproveAllNfts, opts ...grpc.CallOption) (*MsgApproveAllNftsResponse, error) {
	out := new(MsgApproveAllNftsResponse)
	err := c.cc.Invoke(ctx, Msg_ApproveAllNfts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MintNft(ctx context.Context, in *MsgMintNft, opts ...grpc.CallOption) (*MsgMintNftResponse, error) {
	out := new(MsgMintNftResponse)
	err := c.cc.Invoke(ctx, Msg_MintNft_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BurnNft(ctx context.Context, in *MsgBurnNft, opts ...grpc.CallOption) (*MsgBurnNftResponse, error) {
	out := new(MsgBurnNftResponse)
	err := c.cc.Invoke(ctx, Msg_BurnNft_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferNft(ctx context.Context, in *MsgTransferNft, opts ...grpc.CallOption) (*MsgTransferNftResponse, error) {
	out := new(MsgTransferNftResponse)
	err := c.cc.Invoke(ctx, Msg_TransferNft_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ApproveNft(ctx context.Context, in *MsgApproveNft, opts ...grpc.CallOption) (*MsgApproveNftResponse, error) {
	out := new(MsgApproveNftResponse)
	err := c.cc.Invoke(ctx, Msg_ApproveNft_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGuardSoulBondNftImage(ctx context.Context, in *MsgUpdateGuardSoulBondNftImage, opts ...grpc.CallOption) (*MsgUpdateGuardSoulBondNftImageResponse, error) {
	out := new(MsgUpdateGuardSoulBondNftImageResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateGuardSoulBondNftImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateRestrictedCollectionNftImage(ctx context.Context, in *MsgUpdateRestrictedCollectionNftImage, opts ...grpc.CallOption) (*MsgUpdateRestrictedCollectionNftImageResponse, error) {
	out := new(MsgUpdateRestrictedCollectionNftImageResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateRestrictedCollectionNftImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateRestrictedCollectionNftImageBatch(ctx context.Context, in *MsgUpdateRestrictedCollectionNftImageBatch, opts ...grpc.CallOption) (*MsgUpdateRestrictedCollectionNftImageBatchResponse, error) {
	out := new(MsgUpdateRestrictedCollectionNftImageBatchResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateRestrictedCollectionNftImageBatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateRestrictedCollectionNftImageGroupedBatch(ctx context.Context, in *MsgUpdateRestrictedCollectionNftImageGroupedBatch, opts ...grpc.CallOption) (*MsgUpdateRestrictedCollectionNftImageGroupedBatchResponse, error) {
	out := new(MsgUpdateRestrictedCollectionNftImageGroupedBatchResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateRestrictedCollectionNftImageGroupedBatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// CreateNftCollection
	CreateNftCollection(context.Context, *MsgCreateNftCollection) (*MsgCreateNftCollectionResponse, error)
	// MintNfts
	MintNfts(context.Context, *MsgMintNfts) (*MsgMintNftsResponse, error)
	// BurnNfts
	BurnNfts(context.Context, *MsgBurnNfts) (*MsgBurnNftsResponse, error)
	// TransferNfts
	TransferNfts(context.Context, *MsgTransferNfts) (*MsgTransferNftsResponse, error)
	// ApproveNfts
	ApproveNfts(context.Context, *MsgApproveNfts) (*MsgApproveNftsResponse, error)
	// ApproveAllNfts
	ApproveAllNfts(context.Context, *MsgApproveAllNfts) (*MsgApproveAllNftsResponse, error)
	// MintNft
	MintNft(context.Context, *MsgMintNft) (*MsgMintNftResponse, error)
	// BurnNft
	BurnNft(context.Context, *MsgBurnNft) (*MsgBurnNftResponse, error)
	// TransferNft
	TransferNft(context.Context, *MsgTransferNft) (*MsgTransferNftResponse, error)
	// ApproveNft
	ApproveNft(context.Context, *MsgApproveNft) (*MsgApproveNftResponse, error)
	// UpdateGuardSoulBondNftImage
	UpdateGuardSoulBondNftImage(context.Context, *MsgUpdateGuardSoulBondNftImage) (*MsgUpdateGuardSoulBondNftImageResponse, error)
	// UpdateRestrictedCollectionNftImage
	UpdateRestrictedCollectionNftImage(context.Context, *MsgUpdateRestrictedCollectionNftImage) (*MsgUpdateRestrictedCollectionNftImageResponse, error)
	// UpdateRestrictedCollectionNftImageBatch
	UpdateRestrictedCollectionNftImageBatch(context.Context, *MsgUpdateRestrictedCollectionNftImageBatch) (*MsgUpdateRestrictedCollectionNftImageBatchResponse, error)
	// UpdateRestrictedCollectionNftImageGroupedBatch
	UpdateRestrictedCollectionNftImageGroupedBatch(context.Context, *MsgUpdateRestrictedCollectionNftImageGroupedBatch) (*MsgUpdateRestrictedCollectionNftImageGroupedBatchResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateNftCollection(context.Context, *MsgCreateNftCollection) (*MsgCreateNftCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNftCollection not implemented")
}
func (UnimplementedMsgServer) MintNfts(context.Context, *MsgMintNfts) (*MsgMintNftsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintNfts not implemented")
}
func (UnimplementedMsgServer) BurnNfts(context.Context, *MsgBurnNfts) (*MsgBurnNftsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnNfts not implemented")
}
func (UnimplementedMsgServer) TransferNfts(context.Context, *MsgTransferNfts) (*MsgTransferNftsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferNfts not implemented")
}
func (UnimplementedMsgServer) ApproveNfts(context.Context, *MsgApproveNfts) (*MsgApproveNftsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveNfts not implemented")
}
func (UnimplementedMsgServer) ApproveAllNfts(context.Context, *MsgApproveAllNfts) (*MsgApproveAllNftsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveAllNfts not implemented")
}
func (UnimplementedMsgServer) MintNft(context.Context, *MsgMintNft) (*MsgMintNftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintNft not implemented")
}
func (UnimplementedMsgServer) BurnNft(context.Context, *MsgBurnNft) (*MsgBurnNftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnNft not implemented")
}
func (UnimplementedMsgServer) TransferNft(context.Context, *MsgTransferNft) (*MsgTransferNftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferNft not implemented")
}
func (UnimplementedMsgServer) ApproveNft(context.Context, *MsgApproveNft) (*MsgApproveNftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveNft not implemented")
}
func (UnimplementedMsgServer) UpdateGuardSoulBondNftImage(context.Context, *MsgUpdateGuardSoulBondNftImage) (*MsgUpdateGuardSoulBondNftImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGuardSoulBondNftImage not implemented")
}
func (UnimplementedMsgServer) UpdateRestrictedCollectionNftImage(context.Context, *MsgUpdateRestrictedCollectionNftImage) (*MsgUpdateRestrictedCollectionNftImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRestrictedCollectionNftImage not implemented")
}
func (UnimplementedMsgServer) UpdateRestrictedCollectionNftImageBatch(context.Context, *MsgUpdateRestrictedCollectionNftImageBatch) (*MsgUpdateRestrictedCollectionNftImageBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRestrictedCollectionNftImageBatch not implemented")
}
func (UnimplementedMsgServer) UpdateRestrictedCollectionNftImageGroupedBatch(context.Context, *MsgUpdateRestrictedCollectionNftImageGroupedBatch) (*MsgUpdateRestrictedCollectionNftImageGroupedBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRestrictedCollectionNftImageGroupedBatch not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateNftCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateNftCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateNftCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateNftCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateNftCollection(ctx, req.(*MsgCreateNftCollection))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MintNfts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintNfts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintNfts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MintNfts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintNfts(ctx, req.(*MsgMintNfts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BurnNfts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurnNfts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BurnNfts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_BurnNfts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BurnNfts(ctx, req.(*MsgBurnNfts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferNfts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferNfts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferNfts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransferNfts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferNfts(ctx, req.(*MsgTransferNfts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ApproveNfts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgApproveNfts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ApproveNfts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ApproveNfts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ApproveNfts(ctx, req.(*MsgApproveNfts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ApproveAllNfts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgApproveAllNfts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ApproveAllNfts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ApproveAllNfts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ApproveAllNfts(ctx, req.(*MsgApproveAllNfts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MintNft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintNft)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintNft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MintNft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintNft(ctx, req.(*MsgMintNft))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BurnNft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurnNft)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BurnNft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_BurnNft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BurnNft(ctx, req.(*MsgBurnNft))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferNft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferNft)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferNft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransferNft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferNft(ctx, req.(*MsgTransferNft))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ApproveNft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgApproveNft)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ApproveNft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ApproveNft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ApproveNft(ctx, req.(*MsgApproveNft))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGuardSoulBondNftImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGuardSoulBondNftImage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGuardSoulBondNftImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateGuardSoulBondNftImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGuardSoulBondNftImage(ctx, req.(*MsgUpdateGuardSoulBondNftImage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateRestrictedCollectionNftImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateRestrictedCollectionNftImage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateRestrictedCollectionNftImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateRestrictedCollectionNftImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateRestrictedCollectionNftImage(ctx, req.(*MsgUpdateRestrictedCollectionNftImage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateRestrictedCollectionNftImageBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateRestrictedCollectionNftImageBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateRestrictedCollectionNftImageBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateRestrictedCollectionNftImageBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateRestrictedCollectionNftImageBatch(ctx, req.(*MsgUpdateRestrictedCollectionNftImageBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateRestrictedCollectionNftImageGroupedBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateRestrictedCollectionNftImageGroupedBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateRestrictedCollectionNftImageGroupedBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateRestrictedCollectionNftImageGroupedBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateRestrictedCollectionNftImageGroupedBatch(ctx, req.(*MsgUpdateRestrictedCollectionNftImageGroupedBatch))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mantrachain.token.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateNftCollection",
			Handler:    _Msg_CreateNftCollection_Handler,
		},
		{
			MethodName: "MintNfts",
			Handler:    _Msg_MintNfts_Handler,
		},
		{
			MethodName: "BurnNfts",
			Handler:    _Msg_BurnNfts_Handler,
		},
		{
			MethodName: "TransferNfts",
			Handler:    _Msg_TransferNfts_Handler,
		},
		{
			MethodName: "ApproveNfts",
			Handler:    _Msg_ApproveNfts_Handler,
		},
		{
			MethodName: "ApproveAllNfts",
			Handler:    _Msg_ApproveAllNfts_Handler,
		},
		{
			MethodName: "MintNft",
			Handler:    _Msg_MintNft_Handler,
		},
		{
			MethodName: "BurnNft",
			Handler:    _Msg_BurnNft_Handler,
		},
		{
			MethodName: "TransferNft",
			Handler:    _Msg_TransferNft_Handler,
		},
		{
			MethodName: "ApproveNft",
			Handler:    _Msg_ApproveNft_Handler,
		},
		{
			MethodName: "UpdateGuardSoulBondNftImage",
			Handler:    _Msg_UpdateGuardSoulBondNftImage_Handler,
		},
		{
			MethodName: "UpdateRestrictedCollectionNftImage",
			Handler:    _Msg_UpdateRestrictedCollectionNftImage_Handler,
		},
		{
			MethodName: "UpdateRestrictedCollectionNftImageBatch",
			Handler:    _Msg_UpdateRestrictedCollectionNftImageBatch_Handler,
		},
		{
			MethodName: "UpdateRestrictedCollectionNftImageGroupedBatch",
			Handler:    _Msg_UpdateRestrictedCollectionNftImageGroupedBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mantrachain/token/v1/tx.proto",
}
