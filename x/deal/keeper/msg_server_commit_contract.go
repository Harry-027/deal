package keeper

import (
	"context"
	"strconv"
	"time"

	"github.com/Harry-027/deal/x/deal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CommitContract is the tx handler to handle commit contract messages
func (k msgServer) CommitContract(goCtx context.Context, msg *types.MsgCommitContract) (*types.MsgCommitContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	deal, found := k.Keeper.GetNewDeal(ctx, msg.DealId)
	if !found {
		return nil, types.ErrDealNotFound
	}

	// validate is the transaction is coming from vendor
	if msg.Creator != deal.Vendor {
		return nil, types.ErrInvalidVendor
	}

	contract, found := k.Keeper.GetNewContract(ctx, msg.DealId, msg.ContractId)
	if !found {
		return nil, types.ErrContractNotFound
	}

	expiry, err := time.Parse(types.TIME_FORMAT, contract.Expiry)
	if err != nil {
		panic("invalid expiry time")
	}

	// don't process the expired contracts
	if ctx.BlockTime().After(expiry) {
		return nil, types.ErrContractExpired
	}

	etaInMins, err := strconv.Atoi(msg.VendorETA)
	if err != nil {
		return nil, types.ErrInvalidTime
	}

	// validate the vendor ETA
	if (contract.OwnerETA / 2) < uint32(etaInMins) {
		return nil, types.ErrVendorETA
	}

	// process and emit the custom event
	contract.Status = types.COMMITTED
	contract.VendorETA = uint32(etaInMins)
	k.Keeper.SetNewContract(ctx, contract)

	ctx.GasMeter().ConsumeGas(types.PROCESS_GAS, "Commit Contract")
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.COMMITTED),
			sdk.NewAttribute(types.IDVALUE, contract.ContractId),
			sdk.NewAttribute(types.VENDOR_ETA, strconv.FormatUint(uint64(contract.VendorETA), 10)),
		),
	)

	return &types.MsgCommitContractResponse{IdValue: contract.ContractId, ContractStatus: contract.Status}, nil
}
