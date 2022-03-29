package keeper

import (
	"github.com/Harry-027/deal/x/deal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNewDeal set a specific newDeal in the store from its index
func (k Keeper) SetNewDeal(ctx sdk.Context, newDeal types.NewDeal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewDealKeyPrefix))
	b := k.cdc.MustMarshal(&newDeal)
	store.Set(types.NewDealKey(
		newDeal.DealId,
	), b)
}

// GetNewDeal returns a newDeal from its index
func (k Keeper) GetNewDeal(
	ctx sdk.Context,
	index string,

) (val types.NewDeal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewDealKeyPrefix))

	b := store.Get(types.NewDealKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNewDeal removes a newDeal from the store
func (k Keeper) RemoveNewDeal(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewDealKeyPrefix))
	store.Delete(types.NewDealKey(
		index,
	))
}

// GetAllNewDeal returns all newDeal
func (k Keeper) GetAllNewDeal(ctx sdk.Context) (list []types.NewDeal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewDealKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NewDeal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
