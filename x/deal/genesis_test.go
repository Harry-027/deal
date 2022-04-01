package deal_test

import (
	"testing"

	keepertest "github.com/Harry-027/deal/testutil/keeper"
	"github.com/Harry-027/deal/x/deal"
	"github.com/Harry-027/deal/x/deal/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DealCounter: &types.DealCounter{
			IdValue: 57,
		},
		NewDealList: []types.NewDeal{
			{
				DealId: "1",
			},
			{
				DealId: "2",
			},
		},
		ContractCounter: []*types.ContractCounter{
			{
				IdValue: 74,
				DealId:  "1",
			},
			{
				IdValue: 2,
				DealId:  "2",
			},
		},
		NewContractList: []types.NewContract{
			{
				ContractId: "1",
				DealId:     "1",
			},
			{
				ContractId: "2",
				DealId:     "2",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DealKeeper(t)
	deal.InitGenesis(ctx, *k, genesisState)
	got := deal.ExportGenesis(ctx, *k)

	require.NotNil(t, got)
	require.Equal(t, genesisState.DealCounter, got.DealCounter)
	require.ElementsMatch(t, genesisState.NewDealList, got.NewDealList)
	require.ElementsMatch(t, genesisState.ContractCounter, got.ContractCounter)
	require.ElementsMatch(t, genesisState.NewContractList, got.NewContractList)
	// this line is used by starport scaffolding # genesis/test/assert
}
