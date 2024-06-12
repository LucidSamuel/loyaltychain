package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "loyaltychain/testutil/keeper"
	"loyaltychain/testutil/nullify"
	"loyaltychain/x/loyalty/keeper"
	"loyaltychain/x/loyalty/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPoint(keeper keeper.Keeper, ctx context.Context, n int) []types.Point {
	items := make([]types.Point, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetPoint(ctx, items[i])
	}
	return items
}

func TestPointGet(t *testing.T) {
	keeper, ctx := keepertest.LoyaltyKeeper(t)
	items := createNPoint(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPoint(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPointRemove(t *testing.T) {
	keeper, ctx := keepertest.LoyaltyKeeper(t)
	items := createNPoint(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePoint(ctx,
			item.Index,
		)
		_, found := keeper.GetPoint(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestPointGetAll(t *testing.T) {
	keeper, ctx := keepertest.LoyaltyKeeper(t)
	items := createNPoint(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPoint(ctx)),
	)
}
