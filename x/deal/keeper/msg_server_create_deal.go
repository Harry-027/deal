package keeper

import (
	"context"
	"strconv"

	"github.com/Harry-027/deal/x/deal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateDeal(goCtx context.Context, msg *types.MsgCreateDeal) (*types.MsgCreateDealResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	dealCounter, found := k.Keeper.GetDealCounter(ctx)
	if !found {
		panic("DealCounter not found")
	}

	dealId := strconv.FormatUint(dealCounter.IdValue, 10)
	newDeal := types.NewDeal{
		DealId:     dealId,
		Owner:      msg.Creator,
		Vendor:     msg.Vendor,
		Commission: msg.Commission,
	}

	err := newDeal.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.SetNewDeal(ctx, newDeal)

	dealCounter.IdValue++
	k.Keeper.SetDealCounter(ctx, dealCounter)

	contractCounter := types.ContractCounter{
		DealId:  dealId,
		IdValue: 0,
	}
	k.Keeper.SetContractCounter(ctx, contractCounter)

	return &types.MsgCreateDealResponse{
		IdValue: dealId,
	}, nil
}
