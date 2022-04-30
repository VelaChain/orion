package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//balancer "github.com/osmosis-labs/osmosis/v7/x/gamm/pool-models/balancer"
)
// make pool assets implement sort
func (pa PoolAssets) Len() int			{ return len(pa.Asset) }
func (pa PoolAssets) Swap(i, j int)		{ pa.Asset[i], pa.Asset[j] = pa.Asset[j], pa.Asset[i] }
func (pa PoolAssets) Less(i, j int) bool	{ return pa.Asset[i].Symbol < pa.Asset[j].Symbol }

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
	var liqProvs LiquidityProviders
	for _, prov := range lp {
		liqProvs.Providers = append(liqProvs.Providers, prov)
	}
	return liqProvs
}

// Returns new lp from address and shares
func NewLiqProvider(creator string, shares PoolShares) LiquidityProvider {
	lp := LiquidityProvider{ Creator: creator}
	lp.Liquidity = shares
	return lp
}

// Returns new pool created by sender w/ given poolId, id assets, shares and default fees
func NewPairPool(poolId string, assets PoolAssets, shares PoolShares) (pool Pool, err error) {
	p := Pool{PoolId: poolId}
	p.Assets = assets
	p.Shares = shares
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
		pa.Asset = append(pa.Asset, a)
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

func ValidJoinRatio(poolPA PoolAssets, addPA PoolAssets) bool {
	return (poolPA.Asset[0].Amount.Mul(addPA.Asset[1].Amount)).Equal(poolPA.Asset[1].Amount.Mul(addPA.Asset[0].Amount))
}

func GetSharesOut(p Pool, assetsIn PoolAssets) sdk.Int {
	return assetsIn.Asset[0].Amount.Mul(p.Shares.Amount).Quo(p.Assets.Asset[0].Amount)
}

func GetAssetPairOut(p Pool, shareAmtIn sdk.Int) (sdk.Int, sdk.Int) {
	// get amounts out
	amountOutA :=  shareAmtIn.Quo(p.Shares.Amount).Mul(p.Assets.Asset[0].Amount)
	amountOutB :=  shareAmtIn.Quo(p.Shares.Amount).Mul(p.Assets.Asset[1].Amount)
	return amountOutA, amountOutB
}

// from osmosis' balancer package
func solveConstantFunctionInvariant(
	tokenBalanceFixedBefore,
	tokenBalanceFixedAfter,
	tokenWeightFixed,
	tokenBalanceUnknownBefore,
	tokenWeightUnknown sdk.Dec,
) sdk.Dec {
	// weightRatio = (weightX/weightY)
	weightRatio := tokenWeightFixed.Quo(tokenWeightUnknown)

	// y = balanceXBefore/balanceXAfter
	y := tokenBalanceFixedBefore.Quo(tokenBalanceFixedAfter)

	// amountY = balanceY * (1 - (y ^ weightRatio))
	yToWeightRatio := Pow(y, weightRatio)
	paranthetical := sdk.OneDec().Sub(yToWeightRatio)
	amountY := tokenBalanceUnknownBefore.Mul(paranthetical)
	return amountY
}

func GetAssetPairSwapOut(pa PoolAssets, assetIn PoolAsset, swapFee sdk.Dec) (PoolAsset, error){
	// identify swapping in and out assets from pool
	var poolAssetIn PoolAsset
	var poolAssetOut PoolAsset
	if pa.Asset[0].Symbol == assetIn.Symbol {
		poolAssetIn = pa.Asset[0]
		poolAssetOut = pa.Asset[1]
	} else {
		poolAssetIn = pa.Asset[1]
		poolAssetOut = pa.Asset[0]
	}
	// set up parameters to solve constant invariant func from balancer
	assetInPostFee := assetIn.Amount.ToDec().Mul(sdk.OneDec().Sub(swapFee))
	poolAssetInBalance := poolAssetIn.Amount.ToDec()
	poolPostSwapInBalance := poolAssetInBalance.Add(assetInPostFee)
	// use invariant to get sdk dec amount out
	tokenAmountOut := solveConstantFunctionInvariant(
		poolAssetInBalance,
		poolPostSwapInBalance,
		poolAssetIn.Amount.ToDec(),
		poolAssetOut.Amount.ToDec(),
		poolAssetOut.Amount.ToDec(),
	)

	tokenOutInt := tokenAmountOut.TruncateInt()
	if !tokenOutInt.IsPositive(){
		return nil, errors.New("token out int negative")
	}

	return NewPoolAsset(poolAssetOut.Symbol, tokenOutInt), nil
}