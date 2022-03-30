package keeper

import (
	"context"

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

	err := msg.DealHandlerValidation(goCtx, &contract)
	if err != nil {
		return nil, err
	}

	consumerAddress, err := contract.GetConsumerAddress()
	if err != nil {
		panic("Invalid consumer address")
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, consumerAddress, sdk.NewCoins(contract.GetCoin(contract.Fees)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	contract.Status = types.CANCELLED
	k.Keeper.SetNewContract(ctx, contract)

	return &types.MsgCancelOrderResponse{IdValue: contract.ContractId, ContractStatus: contract.Status}, nil
}
