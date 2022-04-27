package types

import (
	
	sdk "github.com/cosmos/cosmos-sdk"
)


// TODO 
func (lp LiquidityProvider) Validate() bool {
	return true
}

// TODO 
func (lps LiquidityProviders) Validate() bool {\
	return true
}

// TODO
func (p Pool) Validate() bool {
	return true
}

func HasLiquidity(p Pool) bool {
	return false
} 

// Returns new lp from address and shares
func NewLiqProvider(creator string, shares PoolShares) LiquidityProvider {
	lp := LiquidityProvider{ Creator: creator}
	lp.Liquidity = shares
	return lp
}

// Returns new pool created by sender w/ given poolId, id assets, shares and default fees
func NewBasicPool(poolId string, count uint64, assets PoolAssets, shares PoolShares, lps LiquidityProviders) Pool {
	p := Pool{PoolId: poollId}
	p.Number = count
	p.PoolAssets = assets
	p.PoolShares = shares
	p.Providers = lps
	p.SwapFee = types.PoolDefaultFee
	p.ExitFee = types.PoolDefaultFee
	return p
}
