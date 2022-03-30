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
	logger := k.Logger(ctx)

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

	if contract.Status != types.INDELIVERY || contract.Status == types.DELIVERED {
		return nil, types.ErrNotShipped
	}

	startTime, err := time.Parse(types.TIME_FORMAT, contract.StartTime)
	if err != nil {
		panic("invalid start time")
	}

	deliveryExpectedTime := startTime.Add(time.Duration(contract.OwnerETA))
	deliveryActualTime := ctx.BlockTime()

	logger.Debug("deliveryExpectedTime :: ", deliveryExpectedTime)
	logger.Debug("deliveryActualTime :: ", deliveryActualTime)

	if deliveryActualTime.After(deliveryExpectedTime) {
		deliveryTimeDelay := uint32(deliveryActualTime.Sub(deliveryExpectedTime).Minutes())
		logger.Debug("deliveryTimeDelay :: ", deliveryTimeDelay)
		if contract.ShippingDelay != 0 {
			deliveryTimeDelay = deliveryTimeDelay - contract.ShippingDelay
			logger.Debug("deliveryTimeDelay after subtracting shipping delay", deliveryTimeDelay)
		}
		contract.DeliveryDelay = uint32(deliveryTimeDelay)
	}

	timeTaken := uint32(deliveryActualTime.Sub(startTime).Minutes())
	logger.Debug("timeTaken :: ", timeTaken)

	var refundAmount uint64 = 0

	if timeTaken != 0 {
		vendorSlashPercent := uint64(contract.ShippingDelay / timeTaken)
		logger.Debug("vendorSlashPercent :: ", vendorSlashPercent)

		ownerSlashPercent := uint64(contract.DeliveryDelay / timeTaken)
		logger.Debug("ownerSlashPercent :: ", ownerSlashPercent)

		refundAmount = (vendorSlashPercent * contract.Fees) + (ownerSlashPercent * contract.Fees)
		logger.Debug("refundAmount :: ", refundAmount)
	}

	totalPay := contract.Fees - refundAmount
	logger.Debug("TotalPay :: ", totalPay)

	moduleAccount := k.auth.GetModuleAddress(types.ModuleName)
	moduleBalance := k.bank.GetBalance(ctx, moduleAccount, types.TOKEN)
	if moduleBalance.IsLT(contract.GetCoin(totalPay)) {
		panic("Escrow account insufficient balance")
	}

	vendorPay := uint64(0.01 * float64(deal.Commission*totalPay))
	logger.Debug("vendorPay :: ", vendorPay)

	ownerPay := totalPay - vendorPay
	logger.Debug("ownerPay :: ", ownerPay)

	consumerAddress, err := contract.GetConsumerAddress()
	if err != nil {
		panic("Invalid consumer address")
	}

	ownerAddress, err := deal.GetOwnerAddress()
	if err != nil {
		panic("Invalid owner address")
	}

	vendorAddress, err := deal.GetVendorAddress()
	if err != nil {
		panic("Invalid vendor address")
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, consumerAddress, sdk.NewCoins(contract.GetCoin(refundAmount)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddress, sdk.NewCoins(contract.GetCoin(ownerPay)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, vendorAddress, sdk.NewCoins(contract.GetCoin(vendorPay)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	contract.Status = types.DELIVERED
	k.Keeper.SetNewContract(ctx, contract)

	ctx.GasMeter().ConsumeGas(types.SETTLEMENT_GAS, "Order delivered")
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.DELIVERED),
			sdk.NewAttribute(types.IDVALUE, contract.ContractId),
			sdk.NewAttribute(types.CONSUMER, contract.Consumer),
			sdk.NewAttribute(types.OWNER, deal.Owner),
			sdk.NewAttribute(types.VENDOR, deal.Vendor),
		),
	)

	return &types.MsgOrderDeliveredResponse{IdValue: contract.ContractId, ContractStatus: contract.Status}, nil
}
