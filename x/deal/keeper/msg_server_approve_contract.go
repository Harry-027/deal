package keeper

import (
	"context"
	"time"

	"github.com/Harry-027/deal/x/deal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ApproveContract(goCtx context.Context, msg *types.MsgApproveContract) (*types.MsgApproveContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contract, found := k.Keeper.GetNewContract(ctx, msg.DealId, msg.ContractId)
	if !found {
		return nil, types.ErrContractNotFound
	}

	if msg.Creator != contract.Consumer {
		return nil, types.ErrInvalidConsumer
	}

	expiry, err := time.Parse(time.RFC3339, contract.Expiry)
	if err != nil {
		panic("invalid expiry time")
	}

	if ctx.BlockTime().Before(expiry) {
		return nil, types.ErrContractExpired
	}

	if contract.Status != types.COMMITTED {
		return nil, types.ErrNotCommitted
	}

	err = k.bank.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(contract.Consumer), types.ModuleName, sdk.NewCoins(contract.GetCoin()))
	if err != nil {
    	return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	contract.Status = types.APPROVED
	k.Keeper.SetNewContract(ctx, contract)

	return &types.MsgApproveContractResponse{}, nil
}
