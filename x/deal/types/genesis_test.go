package types_test

import (
	"testing"

	"github.com/Harry-027/deal/x/deal/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				DealCounter: &types.DealCounter{
					IdValue: 53,
				},
				NewDealList: []types.NewDeal{
					{
						DealId: "1",
					},
					{
						DealId: "2",
					},
				},
				ContractCounter: []*types.ContractCounter{
					{
						DealId:  "1",
						IdValue: 74,
					},
					{
						DealId:  "2",
						IdValue: 4,
					},
				},
				NewContractList: []types.NewContract{
					{
						DealId:     "1",
						ContractId: "1",
					},
					{
						DealId:     "2",
						ContractId: "2",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated newDeal",
			genState: &types.GenesisState{
				NewDealList: []types.NewDeal{
					{
						DealId: "1",
					},
					{
						DealId: "1",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated newContract",
			genState: &types.GenesisState{
				NewContractList: []types.NewContract{
					{
						ContractId: "1",
					},
					{
						ContractId: "1",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
