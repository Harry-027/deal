package keeper

import (
	v01 "github.com/Harry-027/deal/x/deal/legacy/v01"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrate2to3 - Migrating deal module from version 2 to 3
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v01.MigrateDealStore(ctx, m.keeper.storeKey, m.keeper.cdc) // v01 is package `x/deal/legacy/v01`.
}