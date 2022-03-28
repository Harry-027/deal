package keeper_test

import (
	"testing"

	testkeeper "github.com/Harry-027/deal/testutil/keeper"
	"github.com/Harry-027/deal/x/deal/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DealKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
