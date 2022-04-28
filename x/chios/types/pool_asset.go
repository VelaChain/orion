package types

import(
	//"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// TODO 
func (pa PoolAsset) Validate() bool {
	return true
}

// TODO
func (pa PoolAssets) Validate() bool {
	return true
}

// TODO
func (ps PoolShares) Validate() bool {
	return true
}

func CoinFromAsset(asset PoolAsset) sdk.Coin{
	return sdk.NewCoin(asset.Symbol, asset.Amount)
}