package loyalty

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"loyaltychain/testutil/sample"
	loyaltysimulation "loyaltychain/x/loyalty/simulation"
	"loyaltychain/x/loyalty/types"
)

// avoid unused import issue
var (
	_ = loyaltysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreatePoint = "op_weight_msg_point"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePoint int = 100

	opWeightMsgUpdatePoint = "op_weight_msg_point"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePoint int = 100

	opWeightMsgDeletePoint = "op_weight_msg_point"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePoint int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	loyaltyGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PointList: []types.Point{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&loyaltyGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePoint int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePoint, &weightMsgCreatePoint, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePoint = defaultWeightMsgCreatePoint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePoint,
		loyaltysimulation.SimulateMsgCreatePoint(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePoint int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePoint, &weightMsgUpdatePoint, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePoint = defaultWeightMsgUpdatePoint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePoint,
		loyaltysimulation.SimulateMsgUpdatePoint(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePoint int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePoint, &weightMsgDeletePoint, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePoint = defaultWeightMsgDeletePoint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePoint,
		loyaltysimulation.SimulateMsgDeletePoint(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePoint,
			defaultWeightMsgCreatePoint,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loyaltysimulation.SimulateMsgCreatePoint(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdatePoint,
			defaultWeightMsgUpdatePoint,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loyaltysimulation.SimulateMsgUpdatePoint(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeletePoint,
			defaultWeightMsgDeletePoint,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loyaltysimulation.SimulateMsgDeletePoint(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
