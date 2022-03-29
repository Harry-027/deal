package keeper

import (
	"context"
	"time"

	"github.com/Harry-027/deal/x/deal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) OrderDelivered(goCtx context.Context, msg *types.MsgOrderDelivered) (*types.MsgOrderDeliveredResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	deal, found := k.Keeper.GetNewDeal(ctx, msg.DealId)
	if !found {
		return nil, types.ErrDealNotFound
	}

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

	if contract.Status != types.INDELIVERY {
		return nil, types.ErrNotShipped
	}

	startTime, err := time.Parse(time.RFC3339, contract.StartTime)
	if err != nil {
		panic("invalid start time")
	}

	deliveryExpectedTime := startTime.Add(time.Duration(contract.OwnerETA))
	deliveryActualTime := ctx.BlockTime()
	if deliveryActualTime.After(deliveryExpectedTime) {
		deliveryTimeDelay := uint32(deliveryActualTime.Sub(deliveryExpectedTime).Minutes())
		if contract.ShippingDelay != 0 {
			deliveryTimeDelay = deliveryTimeDelay - contract.ShippingDelay
		}
		contract.DeliveryDelay = uint32(deliveryTimeDelay)
	}

	timeTaken :=  uint32(deliveryActualTime.Sub(startTime).Minutes())
	vendorSlashPercent := uint64((contract.ShippingDelay / timeTaken) * 100)
	ownerSlashPercent := uint64((contract.DeliveryDelay / timeTaken) * 100)
	refundAmount := (vendorSlashPercent*contract.Fees) + (ownerSlashPercent*contract.Fees)
	totalPay := contract.Fees - refundAmount
	vendorPay := deal.Commission * totalPay 
	ownerPay := totalPay - vendorPay

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(contract.Consumer), sdk.NewCoins(contract.GetCoin(refundAmount)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(deal.Owner), sdk.NewCoins(contract.GetCoin(ownerPay)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(deal.Vendor), sdk.NewCoins(contract.GetCoin(vendorPay)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	contract.Status = types.DELIVERED
	return &types.MsgOrderDeliveredResponse{IdValue: contract.ContractId, ContractStatus: contract.Status}, nil
}
