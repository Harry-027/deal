package deal

import (
	"math/rand"

	"github.com/Harry-027/deal/testutil/sample"
	dealsimulation "github.com/Harry-027/deal/x/deal/simulation"
	"github.com/Harry-027/deal/x/deal/types"
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
	_ = dealsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateDeal = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDeal int = 100

	opWeightMsgCreateContract = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateContract int = 100

	opWeightMsgCommitContract = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCommitContract int = 100

	opWeightMsgApproveContract = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveContract int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dealGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dealGenesis)
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

	var weightMsgCreateDeal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDeal, &weightMsgCreateDeal, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDeal = defaultWeightMsgCreateDeal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDeal,
		dealsimulation.SimulateMsgCreateDeal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateContract, &weightMsgCreateContract, nil,
		func(_ *rand.Rand) {
			weightMsgCreateContract = defaultWeightMsgCreateContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateContract,
		dealsimulation.SimulateMsgCreateContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCommitContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCommitContract, &weightMsgCommitContract, nil,
		func(_ *rand.Rand) {
			weightMsgCommitContract = defaultWeightMsgCommitContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCommitContract,
		dealsimulation.SimulateMsgCommitContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveContract, &weightMsgApproveContract, nil,
		func(_ *rand.Rand) {
			weightMsgApproveContract = defaultWeightMsgApproveContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveContract,
		dealsimulation.SimulateMsgApproveContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
