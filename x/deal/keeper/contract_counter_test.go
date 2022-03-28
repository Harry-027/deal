package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/Harry-027/deal/testutil/keeper"
	"github.com/Harry-027/deal/testutil/nullify"
	"github.com/Harry-027/deal/x/deal/keeper"
	"github.com/Harry-027/deal/x/deal/types"
)

const (
	DEALID = "1"
)

func createTestContractCounter(keeper *keeper.Keeper, ctx sdk.Context) types.ContractCounter {
	item := types.ContractCounter{
		IdValue: 1,
		DealId: DEALID,
	}
	keeper.SetContractCounter(ctx, item)
	return item
}

func TestContractCounterGet(t *testing.T) {
	keeper, ctx := keepertest.DealKeeper(t)
	item := createTestContractCounter(keeper, ctx)
	rst, found := keeper.GetContractCounter(ctx, DEALID)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestContractCounterRemove(t *testing.T) {
	keeper, ctx := keepertest.DealKeeper(t)
	createTestContractCounter(keeper, ctx)
	keeper.RemoveContractCounter(ctx, DEALID)
	_, found := keeper.GetContractCounter(ctx, DEALID)
	require.False(t, found)
}
