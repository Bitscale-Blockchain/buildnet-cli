package bitscale_test

import (
	"testing"

	keepertest "bitscale/testutil/keeper"
	"bitscale/testutil/nullify"
	bitscale "bitscale/x/bitscale/module"
	"bitscale/x/bitscale/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BitscaleKeeper(t)
	bitscale.InitGenesis(ctx, k, genesisState)
	got := bitscale.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
