package keeper

import (
	"context"
	"strconv"
	"strings"

	"cosmossdk.io/errors"

	"github.com/LimeChain/mantrachain/x/token/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nft "github.com/cosmos/cosmos-sdk/x/nft"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) MintNfts(goCtx context.Context, msg *types.MsgMintNfts) (*types.MsgMintNftsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(msg.Receiver) == "" {
		msg.Receiver = msg.Creator
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)

	if err != nil {
		return nil, err
	}

	if len(msg.Nfts.Nfts) == 0 {
		return nil, errors.Wrapf(types.ErrInvalidNftsCount, "nfts length %d invalid, min 1", len(msg.Nfts.Nfts))
	}

	var collectionCreator sdk.AccAddress

	if strings.TrimSpace(msg.CollectionCreator) == "" && !msg.StrictCollection {
		msg.CollectionCreator = msg.Creator
		collectionCreator = creator
	} else {
		collectionCreator, err = sdk.AccAddressFromBech32(msg.CollectionCreator)

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid collection creator")
		}
	}

	collectionController := NewNftCollectionController(ctx, collectionCreator, msg.StrictCollection).
		WithId(msg.CollectionId).
		WithStore(k)

	if !msg.StrictCollection {
		err = collectionController.
			CreateDefaultIfEmptyId().
			Execute()

		if err != nil {
			return nil, err
		}
	}

	// TODO: add operators auth when they are implemented and when the collection is with soul bond nfts
	err = collectionController.
		MustExist().
		IsOpenedOrHasOwner(creator).
		Validate()

	if err != nil {
		return nil, err
	}

	collectionIndex := collectionController.getIndex()
	collectionId := collectionController.getId()

	nftController := NewNftController(ctx, collectionIndex).
		WithMetadata(msg.Nfts.Nfts).
		WithStore(k).
		WithConfiguration(k.GetParams(ctx))

	err = nftController.
		FilterEmptyIds().
		FilterExist().
		Execute()

	if err != nil {
		return nil, err
	}

	err = nftController.
		ValidMetadata().
		Validate()

	if err != nil {
		return nil, err
	}

	nftsMetadata := nftController.getMetadata()

	if len(nftsMetadata) == 0 {
		return nil, errors.Wrap(types.ErrInvalidNftsCount, "existing nfts")
	}

	var newNfts []nft.NFT
	var newNftsMetadata []types.Nft
	var nftsIds []string

	for _, nftMetadata := range nftsMetadata {
		nftIndex := types.GetNftIndex(collectionIndex, nftMetadata.Id)

		newNfts = append(newNfts, nft.NFT{
			ClassId: string(collectionIndex),
			Id:      string(nftIndex),
			Uri:     types.ModuleName,
			UriHash: nftMetadata.Id,
			Data:    nftMetadata.Data,
		})

		newNftsMetadata = append(newNftsMetadata, types.Nft{
			Index:             nftIndex,
			Id:                nftMetadata.Id,
			Images:            nftMetadata.Images,
			Url:               nftMetadata.Url,
			Links:             nftMetadata.Links,
			Title:             nftMetadata.Title,
			Description:       nftMetadata.Description,
			Attributes:        nftMetadata.Attributes,
			CollectionIndex:   collectionIndex,
			CollectionId:      collectionId,
			CollectionCreator: collectionCreator,
			Creator:           creator,
		})

		nftsIds = append(nftsIds, string(nftMetadata.Id))
	}

	nftExecutor := NewNftExecutor(ctx, k.nftKeeper)
	err = nftExecutor.BatchMintNft(newNfts, receiver)
	if err != nil {
		return nil, err
	}

	k.SetNfts(ctx, newNftsMetadata)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.TypeMsgMintNfts),
			sdk.NewAttribute(types.AttributeKeyNftCollectionCreator, collectionCreator.String()),
			sdk.NewAttribute(types.AttributeKeyNftCollectionId, collectionId),
			sdk.NewAttribute(types.AttributeKeyNftsIds, strings.Join(nftsIds[:], ",")),
			sdk.NewAttribute(types.AttributeKeySigner, creator.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, receiver.String()),
		),
	)

	return &types.MsgMintNftsResponse{
		NftsIds:           nftsIds,
		Creator:           creator.String(),
		Receiver:          receiver.String(),
		CollectionCreator: collectionCreator.String(),
		CollectionId:      collectionId,
	}, nil
}

func (k msgServer) BurnNfts(goCtx context.Context, msg *types.MsgBurnNfts) (*types.MsgBurnNftsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, err
	}

	if len(msg.Nfts.NftsIds) == 0 {
		return nil, errors.Wrapf(types.ErrInvalidNftsCount, "nfts length %d invalid, min 1", len(msg.Nfts.NftsIds))
	}

	var collectionCreator sdk.AccAddress

	if strings.TrimSpace(msg.CollectionCreator) == "" && !msg.StrictCollection {
		msg.CollectionCreator = msg.Creator
		collectionCreator = owner
	} else {
		collectionCreator, err = sdk.AccAddressFromBech32(msg.CollectionCreator)

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid collection creator")
		}
	}

	collectionController := NewNftCollectionController(ctx, collectionCreator, msg.StrictCollection).
		WithId(msg.CollectionId).
		WithStore(k)

	err = collectionController.
		MustExist().
		Validate()

	if err != nil {
		return nil, err
	}

	collectionIndex := collectionController.getIndex()
	collectionId := collectionController.getId()

	nftController := NewNftController(ctx, collectionIndex).
		WithIds(msg.Nfts.NftsIds).
		WithStore(k).
		WithConfiguration(k.GetParams(ctx))

	err = nftController.
		ValidMetadataMaxCount().
		Validate()

	if err != nil {
		return nil, err
	}

	// TODO: add operators auth when they are implemented and when the collection is with soul bond nfts
	err = nftController.
		FilterEmptyIds().
		FilterNotExist().
		FilterNotOwnOfClass(owner).
		Execute()

	if err != nil {
		return nil, err
	}

	nftsIds := nftController.getNftsIds()

	if len(nftsIds) == 0 {
		return nil, errors.Wrap(types.ErrInvalidNftsCount, "not existing nfts or not an owner")
	}

	nftsIndexes := nftController.getIndexes()

	nftExecutor := NewNftExecutor(ctx, k.nftKeeper)
	err = nftExecutor.BatchBurnNft(string(collectionIndex), nftsIds)
	if err != nil {
		return nil, err
	}

	k.DeleteApprovedNfts(ctx, collectionIndex, nftsIndexes)
	k.DeleteNfts(ctx, collectionIndex, nftsIndexes)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.TypeMsgBurnNfts),
			sdk.NewAttribute(types.AttributeKeyNftCollectionCreator, collectionCreator.String()),
			sdk.NewAttribute(types.AttributeKeyNftCollectionId, collectionId),
			sdk.NewAttribute(types.AttributeKeyNftsIds, strings.Join(nftsIds[:], ",")),
			sdk.NewAttribute(types.AttributeKeySigner, owner.String()),
		),
	)

	return &types.MsgBurnNftsResponse{
		NftsIds:           nftsIds,
		Burner:            owner.String(),
		CollectionCreator: collectionCreator.String(),
		CollectionId:      collectionId,
	}, nil
}

func (k msgServer) ApproveNfts(goCtx context.Context, msg *types.MsgApproveNfts) (*types.MsgApproveNftsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, err
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)

	if err != nil {
		return nil, err
	}

	if len(msg.Nfts.NftsIds) == 0 {
		return nil, errors.Wrapf(types.ErrInvalidNftsCount, "nfts length %d invalid, min 1", len(msg.Nfts.NftsIds))
	}

	var collectionCreator sdk.AccAddress

	if strings.TrimSpace(msg.CollectionCreator) == "" && !msg.StrictCollection {
		msg.CollectionCreator = msg.Creator
		collectionCreator = owner
	} else {
		collectionCreator, err = sdk.AccAddressFromBech32(msg.CollectionCreator)

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid collection creator")
		}
	}

	collectionController := NewNftCollectionController(ctx, collectionCreator, msg.StrictCollection).
		WithId(msg.CollectionId).
		WithStore(k)

	err = collectionController.
		MustExist().
		Validate()

	if err != nil {
		return nil, err
	}

	collectionIndex := collectionController.getIndex()
	collectionId := collectionController.getId()

	_, areSoulBondedNfts := k.GetSoulBondedNftsCollection(
		ctx,
		collectionController.getId(),
	)

	if areSoulBondedNfts {
		return nil, errors.Wrapf(types.ErrApproveSoulBondedNftsNotSupported, "cannot approve soul bonded nfts")
	}

	nftController := NewNftController(ctx, collectionIndex).
		WithIds(msg.Nfts.NftsIds).
		WithStore(k).
		WithConfiguration(k.GetParams(ctx))

	err = nftController.
		ValidMetadataMaxCount().
		Validate()

	if err != nil {
		return nil, err
	}

	err = nftController.
		FilterEmptyIds().
		FilterNotExist().
		FilterNotOwnOfClass(owner).
		Execute()

	if err != nil {
		return nil, err
	}

	nftsIds := nftController.getNftsIds()

	if len(nftsIds) == 0 {
		return nil, errors.Wrap(types.ErrInvalidNftsCount, "not existing nfts or not an owner")
	}

	nftsIndexes := nftController.getIndexes()

	k.SetApprovedNfts(ctx, collectionIndex, nftsIndexes, owner, receiver, msg.Approved)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.TypeMsgApproveNfts),
			sdk.NewAttribute(types.AttributeKeyNftCollectionCreator, collectionCreator.String()),
			sdk.NewAttribute(types.AttributeKeyNftCollectionId, collectionId),
			sdk.NewAttribute(types.AttributeKeyNftsIds, strings.Join(nftsIds[:], ",")),
			sdk.NewAttribute(types.AttributeKeySigner, owner.String()),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, receiver.String()),
		),
	)

	return &types.MsgApproveNftsResponse{
		NftsIds:           nftsIds,
		Owner:             owner.String(),
		Receiver:          receiver.String(),
		Approved:          msg.Approved,
		CollectionCreator: collectionCreator.String(),
		CollectionId:      collectionId,
	}, nil
}

func (k msgServer) ApproveAllNfts(goCtx context.Context, msg *types.MsgApproveAllNfts) (*types.MsgApproveAllNftsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, err
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)

	if err != nil {
		return nil, err
	}

	k.SetApprovedAllNfts(ctx, owner, receiver, msg.Approved)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.TypeMsgApproveAllNfts),
			sdk.NewAttribute(types.AttributeKeyApproved, strconv.FormatBool(msg.Approved)),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, receiver.String()),
		),
	)

	return &types.MsgApproveAllNftsResponse{
		Owner:    owner.String(),
		Receiver: receiver.String(),
		Approved: msg.Approved,
	}, nil
}

func (k msgServer) MintNft(goCtx context.Context, msg *types.MsgMintNft) (*types.MsgMintNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(msg.Receiver) == "" {
		msg.Receiver = msg.Creator
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)

	if err != nil {
		return nil, err
	}

	if msg.Nft == nil {
		return nil, errors.Wrap(types.ErrInvalidNft, "nft cannot be empty")
	}

	var collectionCreator sdk.AccAddress

	if strings.TrimSpace(msg.CollectionCreator) == "" && !msg.StrictCollection {
		msg.CollectionCreator = msg.Creator
		collectionCreator = creator
	} else {
		collectionCreator, err = sdk.AccAddressFromBech32(msg.CollectionCreator)

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid collection creator")
		}
	}

	collectionController := NewNftCollectionController(ctx, collectionCreator, msg.StrictCollection).
		WithId(msg.CollectionId).
		WithStore(k)

	if !msg.StrictCollection {
		err = collectionController.
			CreateDefaultIfEmptyId().
			Execute()

		if err != nil {
			return nil, err
		}
	}

	// TODO: add operators auth when they are implemented and when the collection is with soul bond nfts
	err = collectionController.
		MustExist().
		IsOpenedOrHasOwner(creator).
		Validate()

	if err != nil {
		return nil, err
	}

	collectionIndex := collectionController.getIndex()
	collectionId := collectionController.getId()

	nftController := NewNftController(ctx, collectionIndex).
		WithMetadata([]*types.MsgNftMetadata{msg.Nft}).
		WithStore(k).
		WithConfiguration(k.GetParams(ctx))

	err = nftController.
		FilterEmptyIds().
		FilterExist().
		Execute()

	if err != nil {
		return nil, err
	}

	err = nftController.
		ValidMetadata().
		Validate()

	if err != nil {
		return nil, err
	}

	nftsMetadata := nftController.getMetadata()

	if len(nftsMetadata) == 0 {
		return nil, errors.Wrap(types.ErrInvalidNft, "existing or invalid nft")
	}

	nftMetadata := nftsMetadata[0]
	nftIndex := types.GetNftIndex(collectionIndex, nftMetadata.Id)

	newNft := nft.NFT{
		ClassId: string(collectionIndex),
		Id:      string(nftIndex),
		Uri:     types.ModuleName,
		UriHash: nftMetadata.Id,
		Data:    nftMetadata.Data,
	}

	newNftMetadata := types.Nft{
		Index:             nftIndex,
		Images:            nftMetadata.Images,
		Url:               nftMetadata.Url,
		Links:             nftMetadata.Links,
		Title:             nftMetadata.Title,
		Description:       nftMetadata.Description,
		Attributes:        nftMetadata.Attributes,
		CollectionIndex:   collectionIndex,
		CollectionId:      collectionId,
		CollectionCreator: collectionCreator,
		Creator:           creator,
	}

	nftId := nftMetadata.Id

	nftExecutor := NewNftExecutor(ctx, k.nftKeeper)
	err = nftExecutor.MintNft(newNft, receiver)
	if err != nil {
		return nil, err
	}

	k.SetNft(ctx, newNftMetadata)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.TypeMsgMintNft),
			sdk.NewAttribute(types.AttributeKeyNftCollectionCreator, collectionCreator.String()),
			sdk.NewAttribute(types.AttributeKeyNftCollectionId, collectionId),
			sdk.NewAttribute(types.AttributeKeyNftId, nftId),
			sdk.NewAttribute(types.AttributeKeySigner, creator.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, receiver.String()),
		),
	)

	return &types.MsgMintNftResponse{
		NftId:             nftId,
		Creator:           creator.String(),
		Receiver:          receiver.String(),
		CollectionCreator: collectionCreator.String(),
		CollectionId:      collectionId,
	}, nil
}

func (k msgServer) BurnNft(goCtx context.Context, msg *types.MsgBurnNft) (*types.MsgBurnNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(msg.NftId) == "" {
		return nil, errors.Wrap(types.ErrInvalidNft, "nft id cannot be empty")
	}

	var collectionCreator sdk.AccAddress

	if strings.TrimSpace(msg.CollectionCreator) == "" && !msg.StrictCollection {
		msg.CollectionCreator = msg.Creator
		collectionCreator = owner
	} else {
		collectionCreator, err = sdk.AccAddressFromBech32(msg.CollectionCreator)

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid collection creator")
		}
	}

	collectionController := NewNftCollectionController(ctx, collectionCreator, msg.StrictCollection).
		WithId(msg.CollectionId).
		WithStore(k)

	err = collectionController.
		MustExist().
		Validate()

	if err != nil {
		return nil, err
	}

	collectionIndex := collectionController.getIndex()
	collectionId := collectionController.getId()

	nftController := NewNftController(ctx, collectionIndex).
		WithId(msg.NftId).
		WithStore(k).
		WithConfiguration(k.GetParams(ctx))

	// TODO: add operators auth when they are implemented and when the collection is with soul bond nfts
	err = nftController.
		FilterEmptyIds().
		FilterNotExist().
		FilterNotOwnOfClass(owner).
		Execute()

	if err != nil {
		return nil, err
	}

	nftsIds := nftController.getNftsIds()

	if len(nftsIds) == 0 {
		return nil, errors.Wrap(types.ErrInvalidNft, "not existing nft or not an owner")
	}

	nftsIndexes := nftController.getIndexes()

	nftExecutor := NewNftExecutor(ctx, k.nftKeeper)
	err = nftExecutor.BurnNft(string(collectionIndex), nftsIds[0])
	if err != nil {
		return nil, err
	}

	k.DeleteApprovedNft(ctx, collectionIndex, nftsIndexes[0])
	k.DeleteNft(ctx, collectionIndex, nftsIndexes[0])

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.TypeMsgBurnNft),
			sdk.NewAttribute(types.AttributeKeyNftCollectionCreator, collectionCreator.String()),
			sdk.NewAttribute(types.AttributeKeyNftCollectionId, collectionId),
			sdk.NewAttribute(types.AttributeKeyNftId, nftsIds[0]),
			sdk.NewAttribute(types.AttributeKeySigner, owner.String()),
		),
	)

	return &types.MsgBurnNftResponse{
		NftId:             nftsIds[0],
		Burner:            owner.String(),
		CollectionCreator: collectionCreator.String(),
		CollectionId:      collectionId,
	}, nil
}

func (k msgServer) ApproveNft(goCtx context.Context, msg *types.MsgApproveNft) (*types.MsgApproveNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, err
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)

	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(msg.NftId) == "" {
		return nil, errors.Wrap(types.ErrInvalidNft, "nft id cannot be empty")
	}

	var collectionCreator sdk.AccAddress

	if strings.TrimSpace(msg.CollectionCreator) == "" && !msg.StrictCollection {
		msg.CollectionCreator = msg.Creator
		collectionCreator = owner
	} else {
		collectionCreator, err = sdk.AccAddressFromBech32(msg.CollectionCreator)

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid collection creator")
		}
	}

	collectionController := NewNftCollectionController(ctx, collectionCreator, msg.StrictCollection).
		WithId(msg.CollectionId).
		WithStore(k)

	err = collectionController.
		MustExist().
		Validate()

	if err != nil {
		return nil, err
	}

	collectionIndex := collectionController.getIndex()
	collectionId := collectionController.getId()

	_, areSoulBondedNfts := k.GetSoulBondedNftsCollection(
		ctx,
		collectionController.getId(),
	)

	if areSoulBondedNfts {
		return nil, errors.Wrapf(types.ErrApproveSoulBondedNftNotSupported, "cannot approve soul bonded nft")
	}

	nftController := NewNftController(ctx, collectionIndex).
		WithId(msg.NftId).
		WithStore(k).
		WithConfiguration(k.GetParams(ctx))

	err = nftController.
		FilterEmptyIds().
		FilterNotExist().
		FilterNotOwnOfClass(owner).
		Execute()

	if err != nil {
		return nil, err
	}

	nftsIds := nftController.getNftsIds()

	if len(nftsIds) == 0 {
		return nil, errors.Wrap(types.ErrInvalidNft, "not existing nft or not an owner")
	}

	nftsIndexes := nftController.getIndexes()

	k.SetApprovedNft(ctx, collectionIndex, nftsIndexes[0], owner, receiver, msg.Approved)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.TypeMsgApproveNft),
			sdk.NewAttribute(types.AttributeKeyNftCollectionCreator, collectionCreator.String()),
			sdk.NewAttribute(types.AttributeKeyNftCollectionId, collectionId),
			sdk.NewAttribute(types.AttributeKeyNftId, nftsIds[0]),
			sdk.NewAttribute(types.AttributeKeyApproved, strconv.FormatBool(msg.Approved)),
			sdk.NewAttribute(types.AttributeKeySigner, owner.String()),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, receiver.String()),
		),
	)

	return &types.MsgApproveNftResponse{
		NftId:             nftsIds[0],
		Owner:             owner.String(),
		Receiver:          receiver.String(),
		Approved:          msg.Approved,
		CollectionCreator: collectionCreator.String(),
		CollectionId:      collectionId,
	}, nil
}

func (k msgServer) TransferNft(goCtx context.Context, msg *types.MsgTransferNft) (*types.MsgTransferNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	operator, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, err
	}

	owner, err := sdk.AccAddressFromBech32(msg.Owner)

	if err != nil {
		return nil, err
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)

	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(msg.NftId) == "" {
		return nil, errors.Wrap(types.ErrInvalidNft, "nft id cannot be empty")
	}

	var collectionCreator sdk.AccAddress

	if strings.TrimSpace(msg.CollectionCreator) == "" && !msg.StrictCollection {
		msg.CollectionCreator = msg.Creator
		collectionCreator = operator
	} else {
		collectionCreator, err = sdk.AccAddressFromBech32(msg.CollectionCreator)

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid collection creator")
		}
	}

	collectionController := NewNftCollectionController(ctx, collectionCreator, msg.StrictCollection).
		WithId(msg.CollectionId).
		WithStore(k)

	err = collectionController.
		MustExist().
		Validate()

	if err != nil {
		return nil, err
	}

	collectionIndex := collectionController.getIndex()
	collectionId := collectionController.getId()

	_, areSoulBondedNfts := k.GetSoulBondedNftsCollection(
		ctx,
		collectionController.getId(),
	)

	if areSoulBondedNfts {
		return nil, errors.Wrapf(types.ErrTransferSoulBondedNftNotSupported, "cannot transfer soul bonded nft")
	}

	nftController := NewNftController(ctx, collectionIndex).
		WithId(msg.NftId).
		WithStore(k).
		WithConfiguration(k.GetParams(ctx))

	err = nftController.
		FilterEmptyIds().
		FilterNotExist().
		FilterNotOwnOfClass(owner).
		FilterCannotTransfer(operator).
		Execute()

	if err != nil {
		return nil, err
	}

	nftsIds := nftController.getNftsIds()

	if len(nftsIds) == 0 {
		return nil, errors.Wrap(types.ErrInvalidNft, "not existing nft or no transfer permission")
	}

	nftsIndexes := nftController.getIndexes()

	nftExecutor := NewNftExecutor(ctx, k.nftKeeper)
	err = nftExecutor.TransferNft(string(collectionIndex), string(nftsIndexes[0]), receiver)
	if err != nil {
		return nil, err
	}

	k.DeleteApprovedNft(ctx, collectionIndex, nftsIndexes[0])

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.TypeMsgTransferNft),
			sdk.NewAttribute(types.AttributeKeyNftCollectionCreator, collectionCreator.String()),
			sdk.NewAttribute(types.AttributeKeyNftCollectionId, collectionId),
			sdk.NewAttribute(types.AttributeKeyNftId, nftsIds[0]),
			sdk.NewAttribute(types.AttributeKeySigner, operator.String()),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, receiver.String()),
		),
	)

	return &types.MsgTransferNftResponse{
		NftId:             nftsIds[0],
		Operator:          operator.String(),
		Owner:             owner.String(),
		Receiver:          receiver.String(),
		CollectionCreator: collectionCreator.String(),
		CollectionId:      collectionId,
	}, nil
}

func (k msgServer) TransferNfts(goCtx context.Context, msg *types.MsgTransferNfts) (*types.MsgTransferNftsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	operator, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, err
	}

	owner, err := sdk.AccAddressFromBech32(msg.Owner)

	if err != nil {
		return nil, err
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)

	if err != nil {
		return nil, err
	}

	if len(msg.Nfts.NftsIds) == 0 {
		return nil, errors.Wrapf(types.ErrInvalidNftsCount, "nfts length %d invalid, min 1", len(msg.Nfts.NftsIds))
	}

	var collectionCreator sdk.AccAddress

	if strings.TrimSpace(msg.CollectionCreator) == "" && !msg.StrictCollection {
		msg.CollectionCreator = msg.Creator
		collectionCreator = operator
	} else {
		collectionCreator, err = sdk.AccAddressFromBech32(msg.CollectionCreator)

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid collection creator")
		}
	}

	collectionController := NewNftCollectionController(ctx, collectionCreator, msg.StrictCollection).
		WithId(msg.CollectionId).
		WithStore(k)

	err = collectionController.
		MustExist().
		Validate()

	if err != nil {
		return nil, err
	}

	collectionIndex := collectionController.getIndex()
	collectionId := collectionController.getId()

	_, areSoulBondedNfts := k.GetSoulBondedNftsCollection(
		ctx,
		collectionController.getId(),
	)

	if areSoulBondedNfts {
		return nil, errors.Wrapf(types.ErrTransferSoulBondedNftsNotSupported, "cannot transfer soul bonded nfts")
	}

	nftController := NewNftController(ctx, collectionIndex).
		WithIds(msg.Nfts.NftsIds).
		WithStore(k).
		WithConfiguration(k.GetParams(ctx))

	err = nftController.
		ValidMetadataMaxCount().
		Validate()

	if err != nil {
		return nil, err
	}

	err = nftController.
		FilterEmptyIds().
		FilterNotExist().
		FilterNotOwnOfClass(owner).
		FilterCannotTransfer(operator).
		Execute()

	if err != nil {
		return nil, err
	}

	nftsIds := nftController.getNftsIds()

	if len(nftsIds) == 0 {
		return nil, errors.Wrap(types.ErrInvalidNftsCount, "not existing nfts or no transfer permission")
	}

	nftsIndexes := nftController.getIndexes()

	nftExecutor := NewNftExecutor(ctx, k.nftKeeper)
	err = nftExecutor.BatchTransferNft(string(collectionIndex), nftsIds, receiver)
	if err != nil {
		return nil, err
	}

	k.DeleteApprovedNfts(ctx, collectionIndex, nftsIndexes)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.TypeMsgTransferNfts),
			sdk.NewAttribute(types.AttributeKeyNftCollectionCreator, collectionCreator.String()),
			sdk.NewAttribute(types.AttributeKeyNftCollectionId, collectionId),
			sdk.NewAttribute(types.AttributeKeyNftsIds, strings.Join(nftsIds[:], ",")),
			sdk.NewAttribute(types.AttributeKeySigner, operator.String()),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, receiver.String()),
		),
	)

	return &types.MsgTransferNftsResponse{
		NftsIds:           nftsIds,
		Operator:          operator.String(),
		Owner:             owner.String(),
		Receiver:          receiver.String(),
		CollectionCreator: collectionCreator.String(),
		CollectionId:      collectionId,
	}, nil
}
