package keeper

import (
	"context"
	"strconv"

	"time"

	"github.com/Harry-027/deal/x/deal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateContract is the tx handler to handle create contract messages
func (k msgServer) CreateContract(goCtx context.Context, msg *types.MsgCreateContract) (*types.MsgCreateContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	deal, found := k.Keeper.GetNewDeal(ctx, msg.DealId)
	if !found {
		return nil, types.ErrDealNotFound
	}

	// validate if the tx came from owner
	if msg.Creator != deal.Owner {
		return nil, types.ErrInvalidOwner
	}

	contractCounter, found := k.Keeper.GetContractCounter(ctx, msg.DealId)
	if !found {
		return nil, types.ErrDealNotFound
	}

	contractCounter.IdValue++
	contractId := strconv.FormatUint(contractCounter.IdValue, 10)

	etaInMins, err := strconv.Atoi(msg.OwnerETA)
	if err != nil {
		return nil, types.ErrInvalidTime
	}

	expiryInMins, err := strconv.Atoi(msg.Expiry)
	if err != nil {
		return nil, types.ErrInvalidTime
	}

	expiry := ctx.BlockTime().Add(time.Duration(expiryInMins) * time.Minute)

	// create a new contract under the given dealId
	newContract := types.NewContract{
		DealId:     msg.DealId,
		ContractId: contractId,
		Consumer:   msg.Consumer,
		Desc:       msg.Desc,
		OwnerETA:   uint32(etaInMins),
		Expiry:     expiry.UTC().Format(types.TIME_FORMAT),
		Fees:       msg.Fees,
		StartTime:  ctx.BlockTime().UTC().Format(types.TIME_FORMAT),
		Status:     types.INITIATED,
	}

	k.Keeper.SetNewContract(ctx, newContract)
	k.Keeper.SetContractCounter(ctx, contractCounter)

	ctx.GasMeter().ConsumeGas(types.CREATE_GAS, "Create Contract")
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.INITIATED),
			sdk.NewAttribute(types.IDVALUE, newContract.ContractId),
			sdk.NewAttribute(types.START_TIME, newContract.StartTime),
		),
	)

	return &types.MsgCreateContractResponse{IdValue: contractId, ContractStatus: types.INITIATED}, nil
}
