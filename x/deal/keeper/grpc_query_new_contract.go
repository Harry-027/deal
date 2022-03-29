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

func (k Keeper) NewContractAll(c context.Context, req *types.QueryAllNewContractRequest) (*types.QueryAllNewContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var newContracts []types.NewContract
	ctx := sdk.UnwrapSDKContext(c)

	prefixStoreKey := types.NewContractKey(req.DealId)

	store := ctx.KVStore(k.storeKey)

	newContractStore := prefix.NewStore(store, prefixStoreKey)

	pageRes, err := query.Paginate(newContractStore, req.Pagination, func(key []byte, value []byte) error {
		var newContract types.NewContract
		if err := k.cdc.Unmarshal(value, &newContract); err != nil {
			return err
		}

		newContracts = append(newContracts, newContract)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNewContractResponse{NewContract: newContracts, Pagination: pageRes}, nil
}

func (k Keeper) NewContract(c context.Context, req *types.QueryGetNewContractRequest) (*types.QueryGetNewContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNewContract(
		ctx,
		req.DealId,
		req.ContractId,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetNewContractResponse{NewContract: val}, nil
}
