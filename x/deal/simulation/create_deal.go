package simulation

import (
	"math/rand"

	"github.com/Harry-027/deal/x/deal/keeper"
	"github.com/Harry-027/deal/x/deal/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateDeal(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creatorAccount := accs[0]
		vendorAccount := accs[1]
		commission := r.Intn(80) + 5
		msg := &types.MsgCreateDeal{
			Creator: creatorAccount.Address.String(),
			Vendor: vendorAccount.Address.String(),
			Commission: uint64(commission),
		}

		err := SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{creatorAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateDeal"), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, "create deal", nil), nil, nil		
	}
}
