package keeper

import (
	"context"
	"time"

	"github.com/Harry-027/deal/x/deal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ShipOrder is the tx handler for shipOrder messages from Vendor
func (k msgServer) ShipOrder(goCtx context.Context, msg *types.MsgShipOrder) (*types.MsgShipOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	deal, found := k.Keeper.GetNewDeal(ctx, msg.DealId)
	if !found {
		return nil, types.ErrDealNotFound
	}

	// validate if the tx is from vendor
	if msg.Creator != deal.Vendor {
		return nil, types.ErrInvalidVendor
	}

	contract, found := k.Keeper.GetNewContract(ctx, msg.DealId, msg.ContractId)
	if !found {
		return nil, types.ErrContractNotFound
	}

	if contract.Status != types.APPROVED {
		return nil, types.ErrNotApproved
	}

	startTime, err := time.Parse(types.TIME_FORMAT, contract.StartTime)
	if err != nil {
		panic("invalid start time")
	}

	// Calculate shipping delay if any (will be used later to calculate delay penalty)
	shippingExpectedTime := startTime.Add(time.Duration(contract.VendorETA))
	shippingActualTime := ctx.BlockTime()
	if shippingActualTime.After(shippingExpectedTime) {
		shippingTimeDelay := shippingActualTime.Sub(shippingExpectedTime).Minutes()
		contract.ShippingDelay = uint32(shippingTimeDelay)
	}
	// mark the contract status as in delivery
	contract.Status = types.INDELIVERY
	k.Keeper.SetNewContract(ctx, contract)

	ctx.GasMeter().ConsumeGas(types.PROCESS_GAS, "Order shipped")
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.INDELIVERY),
			sdk.NewAttribute(types.IDVALUE, contract.ContractId),
		),
	)

	return &types.MsgShipOrderResponse{IdValue: contract.ContractId, ContractStatus: contract.Status}, nil
}
