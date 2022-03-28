package deal_test

import (
	"testing"

	keepertest "github.com/Harry-027/deal/testutil/keeper"
	"github.com/Harry-027/deal/testutil/nullify"
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
				DealId: 0,
			},
			{
				DealId: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DealKeeper(t)
	deal.InitGenesis(ctx, *k, genesisState)
	got := deal.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.DealCounter, got.DealCounter)
	require.ElementsMatch(t, genesisState.NewDealList, got.NewDealList)
	// this line is used by starport scaffolding # genesis/test/assert
}
