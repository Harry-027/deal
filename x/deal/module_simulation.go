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
	defaultWeightMsgCreateDeal int = 90

	opWeightMsgCreateContract = "op_weight_msg_create_chain"
	defaultWeightMsgCreateContract int = 90

	opWeightMsgCommitContract = "op_weight_msg_create_chain"
	defaultWeightMsgCommitContract int = 50

	opWeightMsgApproveContract = "op_weight_msg_create_chain"
	defaultWeightMsgApproveContract int = 40

	opWeightMsgShipOrder = "op_weight_msg_create_chain"
	defaultWeightMsgShipOrder int = 40

	opWeightMsgOrderDelivered = "op_weight_msg_create_chain"
	defaultWeightMsgOrderDelivered int = 40

	opWeightMsgCancelOrder = "op_weight_msg_create_chain"
	defaultWeightMsgCancelOrder int = 20
	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	
	dealGenesis := types.GenesisState{
		DealCounter: &types.DealCounter{
			IdValue: 1,
		},
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

	var weightMsgShipOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgShipOrder, &weightMsgShipOrder, nil,
		func(_ *rand.Rand) {
			weightMsgShipOrder = defaultWeightMsgShipOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgShipOrder,
		dealsimulation.SimulateMsgShipOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgOrderDelivered int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOrderDelivered, &weightMsgOrderDelivered, nil,
		func(_ *rand.Rand) {
			weightMsgOrderDelivered = defaultWeightMsgOrderDelivered
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOrderDelivered,
		dealsimulation.SimulateMsgOrderDelivered(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelOrder, &weightMsgCancelOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCancelOrder = defaultWeightMsgCancelOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelOrder,
		dealsimulation.SimulateMsgCancelOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
