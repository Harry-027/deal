package keeper

import (
	"github.com/Harry-027/deal/x/deal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNewContract set a specific newContract in the store -  "NewContract/value/{dealId}"
func (k Keeper) SetNewContract(ctx sdk.Context, newContract types.NewContract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(newContract.DealId))
	b := k.cdc.MustMarshal(&newContract)
	store.Set(types.NewContractKey(
		newContract.ContractId,
	), b)
}

// GetNewContract returns a newContract from its index
func (k Keeper) GetNewContract(
	ctx sdk.Context,
	dealId string,
	contractId string,
) (val types.NewContract, found bool) {
	storeKey := types.NewContractKey(dealId)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storeKey)
	b := store.Get(types.NewContractKey(
		contractId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNewContract removes a newContract from the store
func (k Keeper) RemoveNewContract(
	ctx sdk.Context,
	dealId string,
	contractId string,
) { storeKey := types.NewContractKey(dealId)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storeKey)
	store.Delete(types.NewContractKey(
		contractId,
	))
}

// GetAllNewContract returns all newContract
func (k Keeper) GetAllNewContract(ctx sdk.Context, dealId string) (list []types.NewContract) {
	storeKey := types.NewContractKey(dealId)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NewContract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
