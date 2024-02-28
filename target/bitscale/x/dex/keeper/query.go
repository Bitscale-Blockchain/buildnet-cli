package keeper

import (
	"bitscale/x/dex/types"
)

var _ types.QueryServer = Keeper{}
