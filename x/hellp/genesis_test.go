package hellp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "hellp/testutil/keeper"
	"hellp/testutil/nullify"
	"hellp/x/hellp"
	"hellp/x/hellp/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HellpKeeper(t)
	hellp.InitGenesis(ctx, *k, genesisState)
	got := hellp.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
