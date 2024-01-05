package types_test

import (
	"testing"

	"github.com/MANTRA-Finance/aumega/x/rewards/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
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

				SnapshotList: []types.Snapshot{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				SnapshotCount: 2,
				ProviderList: []types.Provider{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated snapshot",
			genState: &types.GenesisState{
				SnapshotList: []types.Snapshot{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid snapshot count",
			genState: &types.GenesisState{
				SnapshotList: []types.Snapshot{
					{
						Id: 1,
					},
				},
				SnapshotCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated provider",
			genState: &types.GenesisState{
				ProviderList: []types.Provider{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
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