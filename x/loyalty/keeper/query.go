package keeper

import (
	"context"

	"loyaltychain/x/loyalty/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) PointsBalance(c context.Context, req *types.QueryPointsBalanceRequest) (*types.QueryPointsBalanceResponse, error) {
	ownerAddr, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner address")
	}

	point, found := k.GetPoint(sdk.UnwrapSDKContext(c), ownerAddr)
	if !found {
		return &types.QueryPointsBalanceResponse{Balance: 0}, nil
	}

	return &types.QueryPointsBalanceResponse{Balance: point.Balance}, nil
}

func (k Keeper) RewardDetails(c context.Context, req *types.QueryRewardDetailsRequest) (*types.QueryRewardDetailsResponse, error) {
	reward, found := k.GetReward(sdk.UnwrapSDKContext(c), req.RewardItem)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrRewardNotFound, "reward not found")
	}

	return &types.QueryRewardDetailsResponse{Points: reward.Points}, nil
}