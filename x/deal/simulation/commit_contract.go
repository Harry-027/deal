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

func SimulateMsgCommitContract(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount := accs[1]
		dealId := strconv.Itoa(r.Intn(20))
		contractId := strconv.Itoa(r.Intn(5))
		vendorETA := strconv.Itoa(r.Intn(101))
		
		msg := &types.MsgCommitContract{
			Creator: simAccount.Address.String(),
			DealId: dealId,
			ContractId: contractId,
			VendorETA: vendorETA,
		}

		err := SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CommitContract"), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, "approve contract", nil), nil, nil
	}
}
