package keeper

import (
	//"fmt"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	//banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/VelaChain/orion/x/chios/types"
)


// TODO validate everything
func validateCreatePairPoolMsg(ctx sdk.Context, msg types.MsgCreatePairPool) error {
	err := msg.ValidateBasic()
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

// TODO
func validateJoinPairPoolMsg(ctx sdk.Context, msg types.MsgJoinPairPool) error {
	err := msg.ValidateBasic()
	if err != nil {
		return err
	}
	
	return nil
}

// TODO 
func validateExitPairPoolMsg(ctx sdk.Context, msg types.MsgExitPairPool) error {
	err := msg.ValidateBasic()
	if err != nil {
		return err
	}

	return nil
}

// TODO 
func validateSwapPairMsg(ctx sdk.Context, msg types.MsgSwapPair) error {
	err := msg.ValidateBasic()
	if err != nil {
		return err
	}

	return nil
}

// TODO
func validateAddLiquidityMsg(ctx sdk.Context, msg types.MsgAddLiquidityPair) error {
	err := msg.ValidateBasic()
	if err != nil {
		return err
	}

	return nil
}

// TODO
func validateRemoveLiquidityMsg(ctx sdk.Context, msg types.MsgRemoveLiquidityPair) error {
	err := msg.ValidateBasic()
	if err != nil {
		return err
	}

	return nil
}


// when called already know pool DNE 
func (k Keeper) CreatePoolPair(ctx sdk.Context, msg *types.MsgCreatePairPool) (poolID string, sharesOut types.PoolShares, err error) {
	// check msg exists
	if msg == nil {
		// TODO add to errors
		return "", sharesOut, errors.New("create pool msg nil")
	}
	// validate msg data
	if err := validateCreatePairPoolMsg(ctx, *msg); err != nil {
		// TODO add to errors
		return "", sharesOut, err
	}

	// create stuctures to save
	// use msg data to create pool assets
	assetA := types.NewPoolAsset(msg.DenomA, msg.AmountA)
	assetB := types.NewPoolAsset(msg.DenomB, msg.AmountB)
	// wrap pool assets in pool assets
	assets := types.NewPoolAssets(assetA, assetB)
	// get pool name from assets
	poolName := types.GetPoolNameFromAssets(assets)
	// create pool shares
	shares := types.NewPoolShares(poolName, msg.SharesOut)
	// make creator the liquidity provider for created shares
	lp := types.NewLiqProvider(msg.Creator, shares)
	// create new pool w/ lp
	pool, err := types.NewPairPool(poolName, assets, shares, types.NewLiqProviders(lp))
	if err != nil {
		return "", sharesOut, err
	}

	// handle sending coins
	// use msg data to create coins
	coinA := types.CoinFromAsset(assetA)
	coinB := types.CoinFromAsset(assetB)
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
		// TODO add to errors
		return "", sharesOut, err
	}
	
	// w/o account for each pool, do not mint coins, only represent in pool structure
	// update KVStore for pool
	err = k.SetPool(ctx, &pool)
	if err != nil {
		// TODO add to errors
		return "", sharesOut, errors.New("unable to set pool")
	}
	// update KVStore for providers
	err = k.SetLiqProv(ctx, &lp)
	if err != nil {
		// TODO add to errors
		return "", sharesOut, errors.New("unable to set liquidity provider")
	}

	return pool.PoolId, shares, nil
}

// TODO
func (k Keeper) JoinPoolPair(ctx sdk.Context, msg *types.MsgJoinPairPool) (poolName string, sharesOut types.PoolShares, err error) {
	// check msg exists
	if msg == nil {
		// TODO add to errors
		return "", sharesOut, errors.New("empty msg")
	}
	// validate msg
	if err := validateJoinPairPoolMsg(ctx, *msg); err != nil {
		// TODO add to errors
		return "", sharesOut, err
	}

	// use msg data to create pool assets
	assetA := types.NewPoolAsset(msg.DenomA, msg.AmountA)
	assetB := types.NewPoolAsset(msg.DenomB, msg.AmountB)
	// wrap pool assets in pool assets
	assets := types.NewPoolAssets(assetA, assetB)
	// get pool name from assets
	poolName = types.GetPoolNameFromAssets(assets)
	// get pool
	pool, err := k.GetPool(ctx, poolName)
	if err != nil {
		return "", sharesOut, err
	}
	// check ratio's match
	if !types.ValidJoinRatio(pool.Assets, assets) {
		// TODO add to errors
		return "", sharesOut, errors.New("Invalid join ratio")
	}
	// check shares out >= requested shares
	shareAmountOut := types.GetSharesOut(pool, assets)
	if shareAmountOut.LT(msg.SharesOut) {
		return "", sharesOut, errors.New("Shares out less than requested")
	}
	// create shares 
	poolShares := types.NewPoolShares(poolName, shareAmountOut)

	// update/create liquidty provider

	// check if creator already a liquidity provider
	newProv, err := k.GetLiqProv(ctx, poolName, msg.Creator)
	if err == nil {
		// lp already exists so add to shares amount
		newProv.Liquidity.Amount = newProv.Liquidity.Amount.Add(poolShares.Amount) 
	} else {
		// create lp w/ shares 
		newProv = types.NewLiqProvider(msg.Creator, poolShares) 
	}
	
	// update pool balances (allow assets to be in any order)
	var newAmountA sdk.Int
	var newAmountB sdk.Int

	if pool.Assets.Assets[0].Symbol == msg.DenomA {
		newAmountA = pool.Assets.Assets[0].Amount.Add(msg.AmountA)
		newAmountB = pool.Assets.Assets[1].Amount.Add(msg.AmountB)
	} else {
		newAmountA = pool.Assets.Assets[0].Amount.Add(msg.AmountB)
		newAmountB = pool.Assets.Assets[1].Amount.Add(msg.AmountA)
	}

	// update pool
	pool.Assets.Assets[0].Amount = newAmountA
	pool.Assets.Assets[1].Amount = newAmountB
	pool.Shares.Amount = pool.Shares.Amount.Add(poolShares.Amount)

	// handle sending coins
	// use msg data to create coins
	coinA := types.CoinFromAsset(assetA)
	coinB := types.CoinFromAsset(assetB)
	// get creator address
	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return "", poolShares, err
	}
	// check creator has balance for first coin
	if !k.bankKeeper.HasBalance(ctx, addr, coinA) {
		// TODO add to errors
		return "", poolShares, errors.New("creator has insufficient balance for first asset")
	}
	// check creator has balance for second coin
	if !k.bankKeeper.HasBalance(ctx, addr, coinB) {
		// TODO add to errors
		return "", poolShares, errors.New("creator has insufficient balance for second asset")
	}
	// send coins from creator to pool
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(coinA, coinB))
	if err != nil {
		// TODO add to errors
		return "", poolShares, err
	}

	// coins have been successfully sent
	// "send" shares by setting lp w/ new shares 	
	err = k.SetLiqProv(ctx, &newProv) 
	if err != nil {
		return "", poolShares, err
	}
	// set pool
	err = k.SetPool(ctx, &pool)
	if err != nil {
		return "", poolShares, err
	}

	return "", poolShares, nil
}

// TODO
func (k Keeper) ExitPoolPair(ctx sdk.Context, msg *types.MsgExitPairPool) (poolName string, assets types.PoolAssets, err error) {
	// check msg exists
	if msg == nil {
		// TODO add to errors
		return "", assets, errors.New("empty msg")
	}
	// validate msg
	if err := validateExitPairPoolMsg(ctx, *msg); err != nil {
		// TODO add to errors
		return "", assets, err
	}

	return msg.ShareDenom, assets, nil
}

// TODO
func (k Keeper) SwapAssetPair(ctx sdk.Context, msg *types.MsgSwapPair) (assetOut types.PoolAsset, err error) {
	// check msg exists
	if msg == nil {
		// TODO add to errors
		return assetOut, errors.New("empty msg")
	}
	// validate msg
	if err := validateSwapPairMsg(ctx, *msg); err != nil {
		// TODO add to errors
		return assetOut, err
	}

	return assetOut, nil
}

// TODO
func (k Keeper) AddLiquidity(ctx sdk.Context, msg *types.MsgAddLiquidityPair) (poolName string, shares types.PoolShares, err error) {
	// check msg exists
	if msg == nil {
		// TODO add to errors
		return poolName, shares, errors.New("empty msg")
	}
	// validate msg
	if err := validateAddLiquidityMsg(ctx, *msg); err != nil {
		// TODO add to errors
		return poolName, shares, err
	}
	if !ctx.KVStore(k.storeKey).Has(types.GetPoolKeyFromPoolName(types.GetPoolNameFromAssets(types.NewPoolAssets(types.NewPoolAsset(msg.DenomA, msg.AmountA), types.NewPoolAsset(msg.DenomB, msg.AmountB))))){
		// TODO add to errors
		return poolName, shares, errors.New("Pool DNE") 
	}

	return poolName, shares, nil 
}

// TODO
func (k Keeper) RemoveLiquidity(ctx sdk.Context, msg *types.MsgRemoveLiquidityPair) (assetsOut types.PoolAssets, err error) {
	// check msg exists
	if msg == nil {
		// TODO add to errors
		return assetsOut, errors.New("empty msg")
	}
	if err := validateRemoveLiquidityMsg(ctx, *msg); err != nil {
		// TODO add to errors
		return assetsOut, err
	}

	return assetsOut, nil
}