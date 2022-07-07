package keeper

import (
	"regexp"

	"github.com/LimeChain/mantrachain/x/mdb/types"
	"github.com/LimeChain/mantrachain/x/mdb/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type NftCollectionControllerFunc func(controller *NftCollectionController) error

type NftCollectionController struct {
	actions       []NftCollectionControllerFunc
	validators    []NftCollectionControllerFunc
	metadata      *types.MsgCreateNftCollectionMetadata
	nftCollection *types.NftCollection
	store         msgServer
	conf          *types.Params
	ctx           sdk.Context
	creator       sdk.AccAddress
}

func NewNftCollectionController(ctx sdk.Context, creator sdk.AccAddress) *NftCollectionController {
	return &NftCollectionController{
		ctx:     ctx,
		creator: creator,
	}
}

func (c *NftCollectionController) WithMetadata(metadata *types.MsgCreateNftCollectionMetadata) *NftCollectionController {
	c.metadata = metadata
	return c
}

func (c *NftCollectionController) WithStore(store msgServer) *NftCollectionController {
	c.store = store
	return c
}

func (c *NftCollectionController) WithId(id string) *NftCollectionController {
	if c.metadata == nil {
		c.metadata = &types.MsgCreateNftCollectionMetadata{}
	}
	c.metadata.Id = id
	return c
}

func (c *NftCollectionController) WithNftCollection(collection types.NftCollection) *NftCollectionController {
	c.nftCollection = &collection
	return c
}

func (c *NftCollectionController) WithConfiguration(cfg types.Params) *NftCollectionController {
	c.conf = &cfg
	return c
}

func (c *NftCollectionController) CreateDefaultIfNotExists() *NftCollectionController {
	if c.metadata.Id == "" {
		c.actions = append(c.actions, func(controller *NftCollectionController) error {
			return controller.CreateDefault()
		})
	} else {
		c.actions = append(c.actions, func(controller *NftCollectionController) error {
			return c.MustExist().Validate()
		})
	}
	return c
}

func (c *NftCollectionController) CreateDefault() error {
	c.metadata.Id = c.conf.NftCollectionDefaultId
	c.metadata.Opened = true
	c.metadata.Category = string(types.GeneralNftCollectionCat)
	collectionIndex := c.getIndex()

	c.requireNftCollection()

	if c.nftCollection == nil {
		nftExecutor := NewNftExecutor(c.ctx, c.store.nftKeeper)
		err := nftExecutor.SetDefaultClass(collectionIndex)
		if err != nil {
			return err
		}

		newNftCollection := types.NftCollection{
			Index:    collectionIndex,
			Opened:   c.metadata.Opened,
			Category: c.metadata.Category,
			Creator:  c.creator,
			Owner:    c.creator,
		}

		c.store.SetNftCollection(c.ctx, newNftCollection)

		c.nftCollection = &newNftCollection
	}

	return nil
}

func (c *NftCollectionController) Execute() error {
	for _, check := range c.actions {
		if err := check(c); err != nil {
			return err
		}
	}
	return nil
}

func (c *NftCollectionController) Validate() error {
	for _, check := range c.validators {
		if err := check(c); err != nil {
			return err
		}
	}
	return nil
}

func (c *NftCollectionController) MustNotExist() *NftCollectionController {
	c.validators = append(c.validators, func(controller *NftCollectionController) error {
		return controller.mustNotExist()
	})
	return c
}

func (c *NftCollectionController) MustExist() *NftCollectionController {
	c.validators = append(c.validators, func(controller *NftCollectionController) error {
		return controller.mustExist()
	})
	return c
}

func (c *NftCollectionController) MustNotBeDefault() *NftCollectionController {
	c.validators = append(c.validators, func(controller *NftCollectionController) error {
		return controller.mustNotBeDefault()
	})
	return c
}

func (c *NftCollectionController) IsOpenedOrOwner(owner sdk.AccAddress) *NftCollectionController {
	c.validators = append(c.validators, func(controller *NftCollectionController) error {
		return controller.isOpenedOrOwner(owner)
	})
	return c
}

func (c *NftCollectionController) ValidCollectionMetadata() *NftCollectionController {
	c.validators = append(c.validators, func(controller *NftCollectionController) error {
		return controller.validCollectionMetadataId()
	}, func(controller *NftCollectionController) error {
		return controller.validCollectionMetadataName()
	}, func(controller *NftCollectionController) error {
		return controller.validCollectionMetadataCategory()
	}, func(controller *NftCollectionController) error {
		return controller.validCollectionMetadataUrl()
	}, func(controller *NftCollectionController) error {
		return controller.validCollectionMetadataDescription()
	}, func(controller *NftCollectionController) error {
		return controller.validCollectionMetadataSymbol()
	}, func(controller *NftCollectionController) error {
		return controller.validCollectionMetadataOptions()
	}, func(controller *NftCollectionController) error {
		return controller.validCollectionMetadataImages()
	}, func(controller *NftCollectionController) error {
		return controller.validCollectionMetadataLinks()
	}, func(controller *NftCollectionController) error {
		return controller.validCollectionMetadataOpened()
	})
	return c
}

func (c *NftCollectionController) HasOwner(owner sdk.AccAddress) *NftCollectionController {
	c.validators = append(c.validators, func(controller *NftCollectionController) error {
		return controller.hasOwner(owner)
	})
	return c
}

func (c *NftCollectionController) hasOwner(owner sdk.AccAddress) error {
	if err := c.requireNftCollection(); err != nil {
		panic("validation check is not allowed on a non existing nftCollection")
	}
	if owner.Equals(c.nftCollection.Owner) {
		return nil
	}
	return sdkerrors.Wrapf(types.ErrUnauthorized, "unauthorized")
}

func (c *NftCollectionController) mustExist() error {
	return c.requireNftCollection()
}

func (c *NftCollectionController) mustNotBeDefault() error {
	if c.metadata.Id == c.conf.NftCollectionDefaultId {
		return sdkerrors.Wrap(types.ErrInvalidNftCollectionId, c.metadata.Id)
	}
	return nil
}

func (c *NftCollectionController) isOpenedOrOwner(owner sdk.AccAddress) error {
	if err := c.requireNftCollection(); err != nil {
		panic("validation check is not allowed on a non existing nftCollection")
	}

	if c.nftCollection.Opened {
		return nil
	}

	if owner.Equals(c.nftCollection.Owner) {
		return nil
	}
	return sdkerrors.Wrapf(types.ErrUnauthorized, "unauthorized")
}

func (c *NftCollectionController) requireNftCollection() error {
	if c.nftCollection != nil {
		return nil
	}
	creator := sdk.AccAddress(c.creator)
	nftCollection, isFound := c.store.GetNftCollection(c.ctx, creator, c.getIndex())
	if !isFound {
		return sdkerrors.Wrapf(types.ErrNftCollectionDoesNotExist, "not found: %s", c.getIndex())
	}
	c.nftCollection = &nftCollection
	return nil
}

func (c *NftCollectionController) mustNotExist() error {
	err := c.requireNftCollection()
	if err == nil {
		return sdkerrors.Wrapf(types.ErrNftCollectionAlreadyExists, c.metadata.Name)
	}
	return nil
}

func (c *NftCollectionController) validCollectionMetadataCategory() error {
	if c.metadata.Category == "" {
		return nil
	}
	if types.ValidateNftCollectionCategory(types.NftCollectionCategory(c.metadata.Category)) != nil {
		return sdkerrors.Wrapf(types.ErrInvalidNftCollectionCategory, c.metadata.Category)
	}
	return nil
}

func (c *NftCollectionController) validCollectionMetadataSymbol() error {
	if c.metadata.Symbol == "" {
		return nil
	}

	if int32(len(c.metadata.Symbol)) < c.conf.ValidNftCollectionMetadataSymbolMinLength {
		return sdkerrors.Wrapf(types.ErrInvalidNftCollectionSymbol, "%s symbol too short, min %d letters", c.metadata.Symbol, c.conf.ValidNftCollectionMetadataSymbolMinLength)
	}

	if int32(len(c.metadata.Symbol)) > c.conf.ValidNftCollectionMetadataSymbolMaxLength {
		return sdkerrors.Wrapf(types.ErrInvalidNftCollectionSymbol, "%s symbol too long, max %d letters", c.metadata.Symbol, c.conf.ValidNftCollectionMetadataSymbolMaxLength)
	}

	return nil
}

func (c *NftCollectionController) validCollectionMetadataUrl() error {
	if c.metadata.Url == "" {
		return nil
	}

	if !utils.IsUrl(c.metadata.Url) {
		return sdkerrors.Wrapf(types.ErrInvalidNftCollectionUrl, "%s invalid url", c.metadata.Url)
	}

	return nil
}

func (c *NftCollectionController) validCollectionMetadataDescription() error {
	if c.metadata.Description == "" {
		return nil
	}

	if int32(len(c.metadata.Description)) > c.conf.ValidNftCollectionMetadataDescriptionMaxLength {
		return sdkerrors.Wrapf(types.ErrInvalidNftCollectionDescription, "description too long, max %d symbols", c.conf.ValidNftCollectionMetadataDescriptionMaxLength)
	}

	return nil
}

func (c *NftCollectionController) validCollectionMetadataId() error {
	validator := regexp.MustCompile(c.conf.ValidNftCollectionId)

	if !validator.MatchString(c.metadata.Id) {
		return sdkerrors.Wrap(types.ErrInvalidNftCollectionId, c.metadata.Id)
	}

	return nil
}

func (c *NftCollectionController) validCollectionMetadataName() error {
	if len(c.metadata.Name) == 0 {
		return nil
	}

	if int32(len(c.metadata.Name)) > c.conf.ValidNftCollectionMetadataNameMaxLength {
		return sdkerrors.Wrapf(types.ErrInvalidNftCollectionName, "name length %d invalid, max %d", len(c.metadata.Name), c.conf.ValidNftCollectionMetadataNameMaxLength)
	}

	return nil
}

func (c *NftCollectionController) validCollectionMetadataImages() error {
	if len(c.metadata.Images) == 0 {
		return nil
	}
	if int32(len(c.metadata.Images)) > c.conf.ValidNftCollectionMetadataImagesMaxCount {
		return sdkerrors.Wrapf(types.ErrInvalidNftCollectionImagesCount, "images count %d invalid, max %d", len(c.metadata.Images), c.conf.ValidNftCollectionMetadataImagesMaxCount)
	}
	for i, image := range c.metadata.Images {
		if image.Type == "" || int32(len(image.Type)) > c.conf.ValidNftCollectionMetadataImagesTypeMaxLength {
			return sdkerrors.Wrapf(types.ErrInvalidNftCollectionImage, "image type empty or length invalid, index %d , max %d", i, c.conf.ValidNftCollectionMetadataImagesTypeMaxLength)
		}
		if !utils.IsUrl(image.Url) {
			return sdkerrors.Wrapf(types.ErrInvalidNftCollectionImage, "image index %d invalid url", i)
		}
	}
	return nil
}

func (c *NftCollectionController) validCollectionMetadataLinks() error {
	if len(c.metadata.Links) == 0 {
		return nil
	}
	if int32(len(c.metadata.Links)) > c.conf.ValidNftCollectionMetadataLinksMaxCount {
		return sdkerrors.Wrapf(types.ErrInvalidNftCollectionLinksCount, "links count %d invalid, max %d", len(c.metadata.Links), c.conf.ValidNftCollectionMetadataLinksMaxCount)
	}
	for i, link := range c.metadata.Links {
		if link.Type == "" || int32(len(link.Type)) > c.conf.ValidNftCollectionMetadataLinksTypeMaxLength {
			return sdkerrors.Wrapf(types.ErrInvalidNftCollectionLink, "link type empty or length invalid, index %d , max %d", i, c.conf.ValidNftCollectionMetadataLinksTypeMaxLength)
		}
		if !utils.IsUrl(link.Url) {
			return sdkerrors.Wrapf(types.ErrInvalidNftCollectionLink, "link index %d invalid url", i)
		}
	}
	return nil
}

func (c *NftCollectionController) validCollectionMetadataOptions() error {
	if len(c.metadata.Options) == 0 {
		return nil
	}
	if int32(len(c.metadata.Options)) > c.conf.ValidNftCollectionMetadataOptionsMaxCount {
		return sdkerrors.Wrapf(types.ErrInvalidNftCollectionOptionsCount, "options count %d invalid, max %d", len(c.metadata.Options), c.conf.ValidNftCollectionMetadataOptionsMaxCount)
	}
	for i, option := range c.metadata.Options {
		if option.Type == "" || int32(len(option.Type)) > c.conf.ValidNftCollectionMetadataOptionsTypeMaxLength {
			return sdkerrors.Wrapf(types.ErrInvalidNftCollectionOption, "option type empty or length invalid, index %d , max %d", i, c.conf.ValidNftCollectionMetadataOptionsTypeMaxLength)
		}
		if int32(len(option.Value)) > c.conf.ValidNftCollectionMetadataOptionsValueMaxLength || int32(len(option.SubValue)) > c.conf.ValidNftCollectionMetadataOptionsSubValueMaxLength {
			return sdkerrors.Wrapf(types.ErrInvalidNftCollectionOption, "option index %d value/subvalue too long, max %d/%d symbols", i, c.conf.ValidNftCollectionMetadataOptionsValueMaxLength, c.conf.ValidNftCollectionMetadataOptionsSubValueMaxLength)
		}
	}
	return nil
}

func (c *NftCollectionController) validCollectionMetadataOpened() error {
	if c.metadata.Id != c.conf.NftCollectionDefaultId {
		return nil
	}
	if !c.metadata.Opened {
		return sdkerrors.Wrapf(types.ErrInvalidNftCollectionOpened, "collection %d can not be opened", len(c.metadata.Id))
	}
	return nil
}

func (c *NftCollectionController) getId() string {
	if c.metadata.Id == "" {
		return c.conf.NftCollectionDefaultId
	}
	return c.metadata.Id
}

func (c *NftCollectionController) getIndex() []byte {
	id := c.getId()
	return types.GetNftCollectionIndex(c.creator, id)
}