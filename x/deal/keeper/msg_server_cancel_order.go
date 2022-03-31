package keeper

import (
	"context"
	"github.com/Harry-027/deal/x/deal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// CancelOrder is the tx handler for handling cancel order messages
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

	// Validate is the escrow account has enough balance to process the refund
	moduleAccount := k.auth.GetModuleAddress(types.ModuleName)
	moduleBalance := k.bank.GetBalance(ctx, moduleAccount, types.TOKEN)
	if moduleBalance.IsLT(contract.GetCoin(contract.Fees)) {
		panic("Escrow account insufficient balance")
	}

	// Send coins from contract account to the consumer account
	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, consumerAddress, sdk.NewCoins(contract.GetCoin(contract.Fees)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	// mark contract status as cancelled
	contract.Status = types.CANCELLED
	k.Keeper.SetNewContract(ctx, contract)

	ctx.GasMeter().ConsumeGas(types.PROCESS_GAS, "Cancel Contract")
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.CANCELLED),
			sdk.NewAttribute(types.IDVALUE, contract.ContractId),
		),
	)

	return &types.MsgCancelOrderResponse{IdValue: contract.ContractId, ContractStatus: contract.Status}, nil
}
