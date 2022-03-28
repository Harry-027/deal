package keeper

import (
	"github.com/Harry-027/deal/x/deal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetContractCounter set contractCounter in the store
func (k Keeper) SetContractCounter(ctx sdk.Context, contractCounter types.ContractCounter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractCounterKey))
	b := k.cdc.MustMarshal(&contractCounter)
	store.Set([]byte(contractCounter.DealId), b)
}

// GetContractCounter returns contractCounter
func (k Keeper) GetContractCounter(ctx sdk.Context, dealId string) (val types.ContractCounter, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractCounterKey))

	b := store.Get([]byte(dealId))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllContractCounter(ctx sdk.Context) ([]*types.ContractCounter, error) {
	var contractCounter []*types.ContractCounter
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractCounterKey))
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var counter types.ContractCounter
		ct := iterator.Value()
		if err := k.cdc.Unmarshal(ct, &counter); err != nil {
			return []*types.ContractCounter{}, err
		}
		contractCounter = append(contractCounter, &counter)
	}
	return contractCounter, nil
}

// RemoveContractCounter removes contractCounter from the store
func (k Keeper) RemoveContractCounter(ctx sdk.Context, dealId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractCounterKey))
	store.Delete([]byte(dealId))
}
