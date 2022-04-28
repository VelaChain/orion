package keeper

import (
	"fmt"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/VelaChain/orion/x/chios/types"
)

func validateCreatePairPoolMsg(ctx sdk.Context, msg types.CreatePairPool) error {
	err := msg.Validate(ctx)
	if err != nil {
		return err
	}
	// TODO check symbol against registry of symbols
	// check amount A
	if !msg.AmountA.IsPositive() {
		// TODO add to errors
		return errors.New("can't create pool: first amount negative or zero")
	}
	// TODO check symbol against registry of symbols
	// check amount B
	if !msg.AmountB.IsPositive() {
		// TODO add to errors
		return errors.New("can't create pool: second amount negative or zero")
	}
	if !msg.SharesOut.IsPositive() {
		// TODO add to errors
		return errors.New("can't create pool: shares out amount negative or zero")
	}

	return nil
}

// when called already know pool DNE 
func (k Keeper) CreatePoolPair(ctx sdk.Context, msg *types.CreatePairPool) (poolID string, sharesOut PoolShares, err error) {
	// validate inputs
	
	// check msg exists
	if msg == nil {
		// TODO add to errors
		return "", sharesOut, errors.New("create pool msg nil")
	}
	// validate msg data
	if err := validateCreatePairPoolMsg(ctx, msg); err != nil {
		return "", sharesOut, err
	}

	// create stuctures to save
	// use msg data to create pool assets
	assetA := NewPoolAsset(msg.DenomA, msg.AmountA)
	assetB := NewPoolAsset(msg.DenomB, msg.AmountB)
	// wrap pool assets in pool assets
	assets := NewPoolAssets(assetA, assetB)
	// get pool name from assets
	poolName := types.GetPoolNameFromAssets(assets)
	// create pool shares
	shares := NewPoolShares(poolName, msg.SharesOut)
	// make creator the liquidity provider for created shares
	lp := NewLiqProvider(creator, shares)
	// create new pool w/ lp
	pool, err := NewPairPool(poolName, assets, shares, types.NewLiqProviders(lp))
	if err != nil {
		return "", sharesOut, err
	}

	// handle sending coins
	// use msg data to create coins
	coinA := type.CoinFromAsset(assetA)
	coinB := type.CoinFromAsset(assetB)
	// get creator address
	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return "", sharesOut, err
	}
	// check creator has balances for first coin
	if !k.bankKeeper.HasBalance(ctx, addr, coinA) {
		// TODO add to errors
		return "", sharesOut, errors.New("creator has insufficient balance for first asset")
	}
	// check creator has balance for second coin
	if !k.bankKeeper.HasBalance(ctx, addr, coinB) {
		// TODO add to errors
		return "", sharesOut, errors.New("creator has insufficient balance for second asset")
	}
	// send coins from creator to pool
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(coinA, coinB))
	if err != nil {
		return "", sharesOut, err
	}
	
	// w/o account for each pool, do not mint coins, only represent in pool structure
	// update KVStore for pool
	err = k.SetPool(ctx, &pool)
	if err != nil {
		// TODO add to errors
		return "", sharesOut errors.New("unable to set pool")
	}
	// update KVStore for providers
	err = k.SetLiqProv(ctx, &lp)
	if err != nil {
		// TODO add to errors
		return "", sharesOut, errros.New("unable to set liquidity provider")
	}

	return pool.PoolId, shares, nil
}

// TODO
func (k Keeper) JoinPoolPair(ctx sdk.Context, *types.MsgJoinPairPool) (shares PoolShares, err error) {
	// 
	return shares, nil
}

// TODO
func (k Keeper) ExitPoolPair(ctx sdk.Context, *types.MsgExitPairPool) (assets types.PoolAssets, err error) {
	// 
	return assets, nil
}

// TODO
func (k Keeper) SwapAssetPair(ctx sdk.Context, *types.MsgSwapPair) (assetOut PoolAsset, err error) {
	//
	return assetOut, nil
}

// TODO
func (k Keeper) AddLiquidity(ctx sdk.Context, *types.MsgAddLiquidityPair) (shares PoolShares, err error) {
	//
	return shares, nil 
}

// TODO
func (k Keeper) RemoveLiquidity(ctx sdk.Context, *types.MsgRemoveLiquidityPair) (assetsOut PoolAssets, err error) {
	//
	return assetsOut, nil
}