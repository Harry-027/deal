package types

import (
	"fmt"
	"strconv"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DealCounter:     nil,
		NewDealList:     []NewDeal{},
		ContractCounter: nil,
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in newDeal
	newDealIndexMap := make(map[string]struct{})

	for _, elem := range gs.NewDealList {
		index := string(NewDealKey(strconv.FormatUint(elem.DealId, 10)))
		if _, ok := newDealIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for newDeal")
		}
		newDealIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
