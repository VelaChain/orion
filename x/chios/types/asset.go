package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type Assets []Asset

// NewAsset returns a new Asset
func NewAsset(symbol string) Asset {
	return Asset{
		Symbol: symbol,
	}
}

)