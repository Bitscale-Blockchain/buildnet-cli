package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "bitscale/testutil/keeper"
	"bitscale/x/bitscale/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BitscaleKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
