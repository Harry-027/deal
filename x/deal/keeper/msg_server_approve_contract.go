package keeper

import (
	"context"

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

	err := msg.DealHandlerValidation(goCtx, &contract)
	if err != nil {
		return nil, err
	}

	consumerAddress, err := contract.GetConsumerAddress()
	if err != nil {
		panic("Invalid consumer address")
	}
	err = k.bank.SendCoinsFromAccountToModule(ctx, consumerAddress, types.ModuleName, sdk.NewCoins(contract.GetCoin(contract.Fees)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	contract.Status = types.APPROVED
	k.Keeper.SetNewContract(ctx, contract)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.APPROVED),
			sdk.NewAttribute(types.IDVALUE, contract.ContractId),
		),
	)

	return &types.MsgApproveContractResponse{}, nil
}
