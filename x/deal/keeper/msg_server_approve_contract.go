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
	
	err = k.bank.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(contract.Consumer), types.ModuleName, sdk.NewCoins(contract.GetCoin(contract.Fees)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	contract.Status = types.APPROVED
	k.Keeper.SetNewContract(ctx, contract)

	return &types.MsgApproveContractResponse{}, nil
}
