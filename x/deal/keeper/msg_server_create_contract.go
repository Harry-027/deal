package keeper

import (
	"context"
	"strconv"

	"time"

	"github.com/Harry-027/deal/x/deal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateContract(goCtx context.Context, msg *types.MsgCreateContract) (*types.MsgCreateContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contractCounter, found := k.Keeper.GetContractCounter(ctx, msg.DealId)
	if !found {
		return nil, types.ErrDealNotFound
	}

	contractId := strconv.FormatUint(contractCounter.IdValue, 10)

	etaInMins, err := strconv.Atoi(msg.OwnerETA)
	if err != nil {
		return nil, types.ErrInvalidETA
	}

	expiryInMins, err := strconv.Atoi(msg.Expiry)
	if err != nil {
		return nil, types.ErrInvalidETA
	}

	ownerETA := ctx.BlockTime().Add(time.Duration(etaInMins) * time.Minute)
	expiry := ctx.BlockTime().Add(time.Duration(expiryInMins) * time.Minute)

	newContract := types.NewContract{
		DealId:     msg.DealId,
		ContractId: contractId,
		Consumer:   msg.Consumer,
		Desc:       msg.Desc,
		OwnerETA:   ownerETA.String(),
		Expiry:     expiry.String(),
	}

	k.Keeper.SetNewContract(ctx, newContract)
	contractCounter.IdValue++
	k.Keeper.SetContractCounter(ctx, contractCounter)

	return &types.MsgCreateContractResponse{IdValue: contractId, ContractStatus: types.INITIATED}, nil
}
