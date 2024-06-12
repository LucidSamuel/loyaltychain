package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
	"loyaltychain/x/loyalty/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PointAll(ctx context.Context, req *types.QueryAllPointRequest) (*types.QueryAllPointResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var points []types.Point

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	pointStore := prefix.NewStore(store, types.KeyPrefix(types.PointKeyPrefix))

	pageRes, err := query.Paginate(pointStore, req.Pagination, func(key []byte, value []byte) error {
		var point types.Point
		if err := k.cdc.Unmarshal(value, &point); err != nil {
			return err
		}

		points = append(points, point)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPointResponse{Point: points, Pagination: pageRes}, nil
}

func (k Keeper) Point(goCtx context.Context, req *types.QueryGetPointRequest) (*types.QueryGetPointResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Type assertion for sdk.Context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Convert req.Index to sdk.AccAddress
	indexAddr, err := sdk.AccAddressFromBech32(req.Index)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid index address")
	}

	val, found := k.GetPoint(ctx, indexAddr)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPointResponse{Point: val}, nil
}
