package loyalty_test

import (
	"testing"

	keepertest "loyaltychain/testutil/keeper"
	"loyaltychain/testutil/nullify"
	loyalty "loyaltychain/x/loyalty/module"
	"loyaltychain/x/loyalty/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PointList: []types.Point{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LoyaltyKeeper(t)
	loyalty.InitGenesis(ctx, k, genesisState)
	got := loyalty.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PointList, got.PointList)
	// this line is used by starport scaffolding # genesis/test/assert
}
