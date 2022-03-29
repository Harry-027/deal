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

const (
	DealId = "1" 
)

func createNNewContract(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NewContract {
	items := make([]types.NewContract, n)
	for i := range items {
		items[i].ContractId = strconv.Itoa(i)
		items[i].DealId = DealId
		keeper.SetNewContract(ctx, items[i])
	}
	return items
}

func TestNewContractGet(t *testing.T) {
	keeper, ctx := keepertest.DealKeeper(t)
	items := createNNewContract(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNewContract(ctx,
			DealId,
			item.ContractId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNewContractRemove(t *testing.T) {
	keeper, ctx := keepertest.DealKeeper(t)
	items := createNNewContract(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNewContract(ctx,
			DealId,
			item.ContractId,
		)
		_, found := keeper.GetNewContract(ctx,
			DealId,
			item.ContractId,
		)
		require.False(t, found)
	}
}

func TestNewContractGetAll(t *testing.T) {
	keeper, ctx := keepertest.DealKeeper(t)
	items := createNNewContract(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNewContract(ctx, DealId)),
	)
}
