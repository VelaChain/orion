package types

import (
	
	sdk "github.com/cosmos/cosmos-sdk/types"
)
// make pool assets implement sort
func (pa PoolAssets) Len() int			{ return len(pa.Assets) }
func (pa PoolAssets) Swap(i, j int)		{ pa.Assets[i], pa.Assets[j] = pa.Assets[j], pa.Assets[i] }
func (pa PoolAssets) Less(i, j int) bool	{ return pa.Assets[i].Symbol < pa.Assets[j].Symbol }

// TODO 
func (lp LiquidityProvider) Validate() bool {
	return true
}

// TODO 
func (lps LiquidityProviders) Validate() bool {
	return true
}

// TODO
func (p Pool) Validate() bool {
	return true
}


func NewLiqProviders(lp ...LiquidityProvider) LiquidityProviders {
	var liqProv LiquidityProvider
	for _, prov := range lp {
		liqProv.Providers = append(liqProv.Providers, prov)
	}
	return liqProv
}

// Returns new lp from address and shares
func NewLiqProvider(creator string, shares PoolShares) LiquidityProvider {
	lp := LiquidityProvider{ Creator: creator}
	lp.Liquidity = shares
	return lp
}

// Returns new pool created by sender w/ given poolId, id assets, shares and default fees
func NewPairPool(poolId string, assets PoolAssets, shares PoolShares, lps LiquidityProviders) (pool Pool, err error) {
	p := Pool{PoolId: poolId}
	p.Assets = assets
	p.Shares = shares
	p.Providers = lps
	p.SwapFee, err = sdk.NewDecFromStr(PoolDefaultFee)
	if err != nil {
		// TODO add to errors
		return pool, err
	}
	p.ExitFee, err = sdk.NewDecFromStr(PoolDefaultFee)
	if err != nil {
		// TODO add to errors
		return pool, err
	}	
	return p, nil
}

func NewPoolAsset(symbol string, amount sdk.Int) PoolAsset{
	pa := PoolAsset{ Symbol:symbol }
	pa.Amount = amount
	return pa	
}

func NewPoolAssets(assets ...PoolAsset) PoolAssets {
	var pa PoolAssets
	for _, a := range assets {
		pa.Assets = append(pa.Assets, a)
	}
	return pa 
}

func NewPoolShares(symbol string, amount sdk.Int) PoolShares{
	ps := PoolShares{
		Symbol:	symbol,
		Amount:	amount,
	}
	return ps
}

