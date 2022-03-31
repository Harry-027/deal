package simulation

import (
	"math/rand"
	"strconv"

	"github.com/Harry-027/deal/x/deal/keeper"
	"github.com/Harry-027/deal/x/deal/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateContract(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creatorAccount := accs[0]
		consumerAccount := accs[2]
		dealId := strconv.Itoa(r.Intn(30))
		ownerETA := strconv.Itoa(r.Intn(501))
		fees := r.Uint64()
		expiry := strconv.Itoa(r.Int())
		
		msg := &types.MsgCreateContract{
			Creator: creatorAccount.Address.String(),
			DealId: dealId,
			Consumer: consumerAccount.Address.String(),
			Desc: "some random order desc",
			OwnerETA: ownerETA,
			Expiry: expiry,
			Fees: fees,
		}

		err := SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{creatorAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateContract"), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, "create contract", nil), nil, nil
	}
}
