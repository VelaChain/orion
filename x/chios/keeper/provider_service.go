package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/VelaChain/orion/x/chios/types"
)

func (k Keeper) CreateLiqProv(ctx sdk.Context, creator string, shares types.PoolShares) (types.LiquidityProvider, error) {
	lp := types.NewLiqProvider(creator, shares)
	err := k.SetLiqProv(ctx, &lp) 
	if err != nil {
		return nil, err
	}

	return lp, nil
}

func (k Keeper) JoinLiqProv(ctx sdk.Context, creator string, shares types.PoolShares) (types.LiquidityProvider, error) {
	var lp LiquidityProvider
	if k.HasLiqProv(ctx, shares.Symbol, creator) {
		lp, err :=  k.AddToLiqProv(ctx, creator, shares)
		if err != nil {
			return nil, err
		}

		return lp, nil
	} 
	
	lp, err := k.CreateLiqProv(ctx, creator, shares)
	if err != nil {
		return nil, err
	} 

	return lp, nil
}

func (k Keeper) AddLiqToProv(ctx sdk.Context, creator string, shares types.PoolShares) (types.LiquidityProvider, error) {
	// Get liq prov
	lp, err := k.GetLiqProv(ctx, shares.Symbol, creator)
	if err != nil {
		return nil, err
	}
	// add liq to prov
	lp.Liquidity.Amount = lp.Liquidity.Amount.Add(shares.Amount)
	// Set liq prov in store
	err := k.SetLiqProv(ctx, &lp)
	if err != nil {
		return nil, err
	}
	return lp, nil
}

func (k Keeper) RemoveLiqFromProv(ctx sdk.Context, poolName string, lpAddr string, shares Types.PoolShares) (types.LiquidityProvider, error) {
	// Get liq prov
	lp, err := k.GetLiqProv(ctx, shares.Symbol, creator)
	if err != nil {
		return nil, err
	}
	// Remove liquidity from provider
	newAmount = lp.Liquidity.Amount.Sub(shares.Amount)
	if newAmount.IsNegative() {
		return nil, errors.New("not enough shares to remove")
	}

	lp.Liquidity.Amount = newAmount
	// Set liq prov in store
	err := k.SetLiqProv(ctx, &lp)
	if err != nil {
		return nil, err 
	}

	return lp, nil 
}

func (k Keeper) RemoveProvider(ctx sdk.Context, poolName string, lpAddr string) error {
	// try to remove liq prov
	removed, err := k.RemoveLiqPro(ctx, poolName, lpAddr)
	if err != nil || !removed {
		return err
	} 
	
	return nil
}

func (k Keeper) ExitLiqProv(ctx sdk.Context, poolName string, lpAddr string) (sdk.Int, error) {
	// Get provider
	lp, err := k.GetLiqProv(ctx, poolName, lpAddr)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	// Get provider's balance
	removeSharesAmount := lp.Liquidity.Amount
	// Clear provider's balance
	if err := k.RemoveLiqFromProv(ctx, poolName, lpAddr, lp.Liquidity); err != nil {
		return sdk.ZeroInt(), err
	}
	// Remove provider from lps 
	if err := k.RemoveProvider(ctx, poolName, lpAddr); err != nil {
		return removeSharesAmount, err
	}

	return removeSharesAmount, nil
}
