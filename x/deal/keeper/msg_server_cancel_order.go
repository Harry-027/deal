package keeper

import (
	"context"
	"time"

	"github.com/Harry-027/deal/x/deal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CancelOrder(goCtx context.Context, msg *types.MsgCancelOrder) (*types.MsgCancelOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contract, found := k.Keeper.GetNewContract(ctx, msg.DealId, msg.ContractId)
	if !found {
		return nil, types.ErrContractNotFound
	}

	if msg.Creator != contract.Consumer {
		return nil, types.ErrInvalidConsumer
	}

	if contract.Status != types.APPROVED {
		return nil, types.ErrNotApproved
	}

	startTime, err := time.Parse(time.RFC3339, contract.StartTime)
	if err != nil {
		panic("invalid start time")
	}

	deliveryExpectedTime := startTime.Add(time.Duration(contract.OwnerETA))
	timeLimit := uint32(ctx.BlockTime().Sub(deliveryExpectedTime).Minutes())
	if timeLimit < 20 {
		return nil, types.ErrRefund
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(contract.Consumer), sdk.NewCoins(contract.GetCoin(contract.Fees)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	contract.Status = types.CANCELLED
	k.Keeper.SetNewContract(ctx, contract)

	return &types.MsgCancelOrderResponse{IdValue: contract.ContractId, ContractStatus: contract.Status}, nil
}
