package keeper

import (
	"github.com/VelaChain/orion/x/amm/types"
)

var _ types.QueryServer = Keeper{}
