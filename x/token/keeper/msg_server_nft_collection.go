package keeper

import (
	"context"
	"strings"

	"cosmossdk.io/errors"
	"github.com/LimeChain/mantrachain/x/token/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nft "github.com/cosmos/cosmos-sdk/x/nft"
)

func (k msgServer) CreateNftCollection(goCtx context.Context, msg *types.MsgCreateNftCollection) (*types.MsgCreateNftCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(msg.Collection.Id) == "" {
		return nil, errors.Wrap(types.ErrInvalidNftCollectionId, "collection id should not be empty")
	}

	collectionController := NewNftCollectionController(ctx, creator, false).
		WithMetadata(msg.Collection).
		WithStore(k).
		WithConfiguration(k.GetParams(ctx))

	err = collectionController.
		MustNotExist().
		MustNotBeDefault().
		ValidMetadata().
		Validate()

	if err != nil {
		return nil, err
	}

	collectionIndex := collectionController.getIndex()
	collectionId := collectionController.getId()

	nftExecutor := NewNftExecutor(ctx, k.nftKeeper)
	err = nftExecutor.SetClass(nft.Class{
		Id:          string(collectionIndex),
		Name:        msg.Collection.Name,
		Symbol:      msg.Collection.Symbol,
		Description: msg.Collection.Description,
		Uri:         types.ModuleName,
		UriHash:     collectionId,
		Data:        msg.Collection.Data,
	})

	if err != nil {
		return nil, err
	}

	newNftCollection := types.NftCollection{
		Index:    collectionIndex,
		Id:       collectionId,
		Images:   msg.Collection.Images,
		Url:      msg.Collection.Url,
		Links:    msg.Collection.Links,
		Category: msg.Collection.Category,
		Options:  msg.Collection.Options,
		Opened:   msg.Collection.Opened,
		Creator:  creator,
		Owner:    creator,
	}

	k.SetNftCollection(ctx, newNftCollection)

	if msg.Collection.SoulBondedNfts {
		k.SetSoulBondedNftsCollection(ctx, types.SoulBondedNftsCollection{
			Index: collectionId,
		})
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.TypeMsgCreateNftCollection),
			sdk.NewAttribute(types.AttributeKeyNftCollectionId, collectionId),
			sdk.NewAttribute(types.AttributeKeySigner, creator.String()),
			sdk.NewAttribute(types.AttributeKeyOwner, creator.String()),
		),
	)

	return &types.MsgCreateNftCollectionResponse{
		CollectionId:      collectionId,
		CollectionCreator: creator.String(),
	}, nil
}
