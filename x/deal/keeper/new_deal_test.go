package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/Harry-027/deal/testutil/keeper"
	"github.com/Harry-027/deal/testutil/nullify"
	"github.com/Harry-027/deal/x/deal/keeper"
	"github.com/Harry-027/deal/x/deal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNewDeal(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NewDeal {
	items := make([]types.NewDeal, n)
	for i := range items {
		items[i].DealId = uint64(i)

		keeper.SetNewDeal(ctx, items[i])
	}
	return items
}

func TestNewDealGet(t *testing.T) {
	keeper, ctx := keepertest.DealKeeper(t)
	items := createNNewDeal(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNewDeal(ctx,
			strconv.FormatUint(item.DealId, 10),
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNewDealRemove(t *testing.T) {
	keeper, ctx := keepertest.DealKeeper(t)
	items := createNNewDeal(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNewDeal(ctx,
			strconv.FormatUint(item.DealId, 10),
		)
		_, found := keeper.GetNewDeal(ctx,
			strconv.FormatUint(item.DealId, 10),
		)
		require.False(t, found)
	}
}

func TestNewDealGetAll(t *testing.T) {
	keeper, ctx := keepertest.DealKeeper(t)
	items := createNNewDeal(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNewDeal(ctx)),
	)
}
