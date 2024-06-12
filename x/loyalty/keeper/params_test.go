package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "loyaltychain/testutil/keeper"
	"loyaltychain/x/loyalty/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LoyaltyKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
