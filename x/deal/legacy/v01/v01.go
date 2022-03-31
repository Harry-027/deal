package v01

import (
	"fmt"
	"github.com/Harry-027/deal/x/deal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateDealStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	return pruneOldContracts(store, cdc)
}

func pruneOldContracts(store sdk.KVStore, cdc codec.BinaryCodec) error {
	dealStore := prefix.NewStore(store, []byte(types.NewDealKeyPrefix))
	dealStoreIter := dealStore.Iterator(nil, nil)
	defer dealStoreIter.Close()
	var deals []types.NewDeal
	for ; dealStoreIter.Valid(); dealStoreIter.Next() {
		var val types.NewDeal
		cdc.MustUnmarshal(dealStoreIter.Value(), &val)
		deals = append(deals, val)
	}

	for _, deal := range deals {
		contractStorePrefixKey := fmt.Sprintf("%s%s/", types.NewContractKeyPrefix, deal.DealId)
		contractStore := prefix.NewStore(store, []byte(contractStorePrefixKey))
		iterator := sdk.KVStorePrefixIterator(contractStore, []byte{})

		defer iterator.Close()

		for ; iterator.Valid(); iterator.Next() {
			var val types.NewContract
			cdc.MustUnmarshal(iterator.Value(), &val)
			if val.Status == "DELIVERED" {
				contractKey := fmt.Sprintf("%s%s/", types.NewContractKeyPrefix, val.ContractId)
				contractStore.Delete([]byte(contractKey))
			}
		}
	}
	return nil
}
