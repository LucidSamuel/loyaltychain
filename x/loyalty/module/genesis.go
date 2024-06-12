package loyalty

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"loyaltychain/x/loyalty/keeper"
	"loyaltychain/x/loyalty/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the point
	for _, elem := range genState.PointList {
		k.SetPoint(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PointList = k.GetAllPoint(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
