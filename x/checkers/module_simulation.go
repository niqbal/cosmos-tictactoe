package checkers

import (
	"math/rand"

	"github.com/alice/checkers/testutil/sample"
	checkerssimulation "github.com/alice/checkers/x/checkers/simulation"
	"github.com/alice/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = checkerssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateGame = "op_weight_msg_create_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGame int = 100

	opWeightMsgJoinGame = "op_weight_msg_join_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgJoinGame int = 100

	opWeightMsgMarkSpace = "op_weight_msg_mark_space"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMarkSpace int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	checkersGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&checkersGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateGame int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateGame, &weightMsgCreateGame, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGame = defaultWeightMsgCreateGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGame,
		checkerssimulation.SimulateMsgCreateGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgJoinGame int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgJoinGame, &weightMsgJoinGame, nil,
		func(_ *rand.Rand) {
			weightMsgJoinGame = defaultWeightMsgJoinGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgJoinGame,
		checkerssimulation.SimulateMsgJoinGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMarkSpace int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMarkSpace, &weightMsgMarkSpace, nil,
		func(_ *rand.Rand) {
			weightMsgMarkSpace = defaultWeightMsgMarkSpace
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMarkSpace,
		checkerssimulation.SimulateMsgMarkSpace(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
