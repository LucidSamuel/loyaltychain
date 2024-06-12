package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "loyaltychain/testutil/keeper"
	"loyaltychain/x/loyalty/keeper"
	"loyaltychain/x/loyalty/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPointMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.LoyaltyKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreatePoint{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreatePoint(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetPoint(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestPointMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdatePoint
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdatePoint{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdatePoint{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdatePoint{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.LoyaltyKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreatePoint{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreatePoint(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdatePoint(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetPoint(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestPointMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeletePoint
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeletePoint{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeletePoint{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeletePoint{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.LoyaltyKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreatePoint(ctx, &types.MsgCreatePoint{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeletePoint(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetPoint(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
