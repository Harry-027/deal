package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Harry-027/deal/testutil/keeper"
	"github.com/Harry-027/deal/x/deal/keeper"
	"github.com/Harry-027/deal/x/deal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DealKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
