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

func createTestDealCounter(keeper *keeper.Keeper, ctx sdk.Context) types.DealCounter {
	item := types.DealCounter{}
	keeper.SetDealCounter(ctx, item)
	return item
}

func TestDealCounterGet(t *testing.T) {
	keeper, ctx := keepertest.DealKeeper(t)
	item := createTestDealCounter(keeper, ctx)
	rst, found := keeper.GetDealCounter(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestDealCounterRemove(t *testing.T) {
	keeper, ctx := keepertest.DealKeeper(t)
	createTestDealCounter(keeper, ctx)
	keeper.RemoveDealCounter(ctx)
	_, found := keeper.GetDealCounter(ctx)
	require.False(t, found)
}
