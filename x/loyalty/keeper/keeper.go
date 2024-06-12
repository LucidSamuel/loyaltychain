package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"loyaltychain/x/loyalty/types"
)

// Keeper struct
type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
	logger       log.Logger
	authority    string
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetPoint retrieves the point balance for a given owner address
func (k Keeper) GetPoint(ctx sdk.Context, ownerAddr sdk.AccAddress) (types.Point, bool) {
	store := ctx.KVStore(k.storeService.OpenKVStore(ctx)) // Open the KVStore correctly
	key := []byte(ownerAddr.String())
	bz := store.Get(key)
	if bz == nil {
		return types.Point{}, false
	}

	var point types.Point
	k.cdc.MustUnmarshal(bz, &point)
	return point, true
}

// SetPoint sets the point balance for a given owner address
func (k Keeper) SetPoint(ctx sdk.Context, point types.Point) {
	store := ctx.KVStore(k.storeService.OpenKVStore(ctx)) // Open the KVStore correctly
	key := []byte(point.Owner)
	bz := k.cdc.MustMarshal(&point)
	store.Set(key, bz)
}

// GetReward retrieves the reward details for a given reward item
func (k Keeper) GetReward(ctx sdk.Context, rewardItem string) (types.Reward, bool) {
	store := ctx.KVStore(k.storeService.OpenKVStore(ctx)) // Open the KVStore correctly
	key := []byte(rewardItem)
	bz := store.Get(key)
	if bz == nil {
		return types.Reward{}, false
	}

	var reward types.Reward
	k.cdc.MustUnmarshal(bz, &reward)
	return reward, true
}

// IssuePoints credits points to the owner's balance.
func (k Keeper) IssuePoints(ctx sdk.Context, owner string, amount int) error {
	if amount <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount must be positive")
	}

	ownerAddr, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner address is invalid")
	}

	pointBalance, found := k.GetPoint(ctx, ownerAddr)
	if !found {
		pointBalance = types.Point{
			Owner:   ownerAddr.String(),
			Balance: 0,
		}
	}

	pointBalance.Balance += int32(amount)
	k.SetPoint(ctx, pointBalance)

	return nil
}

// RedeemReward exchanges points for a specified reward item.
func (k Keeper) RedeemReward(ctx sdk.Context, owner string, rewardItem string) error {
	ownerAddr, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner address is invalid")
	}

	reward, found := k.GetReward(ctx, rewardItem)
	if !found {
		return sdkerrors.Wrap(types.ErrRewardNotFound, "reward does not exist")
	}

	pointBalance, found := k.GetPoint(ctx, ownerAddr)
	if !found || pointBalance.Balance < reward.Points {
		return sdkerrors.Wrap(types.ErrInsufficientPoints, "insufficient points to redeem reward")
	}

	pointBalance.Balance -= reward.Points
	k.SetPoint(ctx, pointBalance)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			string(types.EventTypeRedeemReward),
			sdk.NewAttribute(string(types.AttributeKeyRewardItem), rewardItem),
			sdk.NewAttribute(string(types.AttributeKeyOwner), owner),
			sdk.NewAttribute(string(types.AttributeKeyPoints), fmt.Sprintf("%d", reward.Points)),
		),
	)

	return nil
}
