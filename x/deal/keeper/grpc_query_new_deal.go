package keeper

import (
	"context"

	"github.com/Harry-027/deal/x/deal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewDealAll is the query handler to fetch all the new deals
func (k Keeper) NewDealAll(c context.Context, req *types.QueryAllNewDealRequest) (*types.QueryAllNewDealResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var newDeals []types.NewDeal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	newDealStore := prefix.NewStore(store, types.KeyPrefix(types.NewDealKeyPrefix))

	pageRes, err := query.Paginate(newDealStore, req.Pagination, func(key []byte, value []byte) error {
		var newDeal types.NewDeal
		if err := k.cdc.Unmarshal(value, &newDeal); err != nil {
			return err
		}

		newDeals = append(newDeals, newDeal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNewDealResponse{NewDeal: newDeals, Pagination: pageRes}, nil
}

// NewDeal is the query handler to fetch the deal details for a given dealId
func (k Keeper) NewDeal(c context.Context, req *types.QueryGetNewDealRequest) (*types.QueryGetNewDealResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNewDeal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetNewDealResponse{NewDeal: val}, nil
}
