package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DealCounter: &DealCounter{
			IdValue: uint64(1),
		},
		NewDealList:     []NewDeal{},
		ContractCounter: nil,
		NewContractList: []NewContract{},
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
		index := elem.DealId
		if _, ok := newDealIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for newDeal")
		}
		newDealIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in newContract
	newContractIndexMap := make(map[string]struct{})

	for _, elem := range gs.NewContractList {
		if _, ok := newContractIndexMap[elem.ContractId]; ok {
			return fmt.Errorf("duplicated index for newContract")
		}
		newContractIndexMap[elem.ContractId] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
