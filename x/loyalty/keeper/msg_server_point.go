package keeper

import (
	"context"

	"loyaltychain/x/loyalty/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
)

func (k msgServer) CreatePoint(goCtx context.Context, msg *types.MsgCreatePoint) (*types.MsgCreatePointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Convert msg.Index to sdk.AccAddress
	indexAddr, err := sdk.AccAddressFromBech32(msg.Index)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid index address")
	}

	// Check if the value already exists
	_, isFound := k.GetPoint(ctx, indexAddr)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var point = types.Point{
		Creator: msg.Creator,
		Index:   msg.Index,
		Owner:   msg.Owner,
		Balance: msg.Balance,
	}

	k.SetPoint(ctx, point)
	return &types.MsgCreatePointResponse{}, nil
}

func (k msgServer) UpdatePoint(goCtx context.Context, msg *types.MsgUpdatePoint) (*types.MsgUpdatePointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Convert msg.Index to sdk.AccAddress
	indexAddr, err := sdk.AccAddressFromBech32(msg.Index)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid index address")
	}

	// Check if the value exists
	valFound, isFound := k.GetPoint(ctx, indexAddr)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var point = types.Point{
		Creator: msg.Creator,
		Index:   msg.Index,
		Owner:   msg.Owner,
		Balance: msg.Balance,
	}

	k.SetPoint(ctx, point)

	return &types.MsgUpdatePointResponse{}, nil
}

func (k msgServer) DeletePoint(goCtx context.Context, msg *types.MsgDeletePoint) (*types.MsgDeletePointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Convert msg.Index to sdk.AccAddress
	indexAddr, err := sdk.AccAddressFromBech32(msg.Index)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid index address")
	}

	// Check if the value exists
	valFound, isFound := k.GetPoint(ctx, indexAddr)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePoint(ctx, indexAddr)

	return &types.MsgDeletePointResponse{}, nil
}
