package keeper

import (
	"github.com/Harry-027/deal/x/deal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDealCounter set dealCounter in the store
func (k Keeper) SetDealCounter(ctx sdk.Context, dealCounter types.DealCounter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DealCounterKey))
	b := k.cdc.MustMarshal(&dealCounter)
	store.Set([]byte{0}, b)
}

// GetDealCounter returns dealCounter
func (k Keeper) GetDealCounter(ctx sdk.Context) (val types.DealCounter, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DealCounterKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDealCounter removes dealCounter from the store
func (k Keeper) RemoveDealCounter(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DealCounterKey))
	store.Delete([]byte{0})
}
