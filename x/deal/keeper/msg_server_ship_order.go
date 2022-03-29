package keeper

import (
	"context"
	"time"

	"github.com/Harry-027/deal/x/deal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ShipOrder(goCtx context.Context, msg *types.MsgShipOrder) (*types.MsgShipOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	deal, found := k.Keeper.GetNewDeal(ctx, msg.DealId)
	if !found {
		return nil, types.ErrDealNotFound
	}

	if msg.Creator != deal.Vendor {
		return nil, types.ErrInvalidVendor
	}

	contract, found := k.Keeper.GetNewContract(ctx, msg.DealId, msg.ContractId)
	if !found {
		return nil, types.ErrContractNotFound
	}

	expiry, err := time.Parse(time.RFC3339, contract.Expiry)
	if err != nil {
		panic("invalid expiry time")
	}

	if ctx.BlockTime().Before(expiry) {
		return nil, types.ErrContractExpired
	}

	startTime, err := time.Parse(time.RFC3339, contract.StartTime)
	if err != nil {
		panic("invalid start time")
	}

	shippingExpectedTime := startTime.Add(time.Duration(contract.VendorETA))
	shippingActualTime := ctx.BlockTime()
	if shippingActualTime.After(shippingExpectedTime) {
		shippingTimeDelay := shippingActualTime.Sub(shippingExpectedTime).Minutes()
		contract.ShippingDelay = uint32(shippingTimeDelay)
	}
	contract.Status = types.INDELIVERY
	k.Keeper.SetNewContract(ctx, contract)
	return &types.MsgShipOrderResponse{IdValue: contract.ContractId, ContractStatus: contract.Status}, nil
}
