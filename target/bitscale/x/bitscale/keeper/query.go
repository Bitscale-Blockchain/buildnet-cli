package keeper

import (
	"bitscale/x/bitscale/types"
)

var _ types.QueryServer = Keeper{}
