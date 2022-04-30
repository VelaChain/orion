package keeper

import (
	"context"
	"errors"

	"github.com/VelaChain/orion/x/chios/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//"github.com/cosmos/cosmos-sdk/types/query"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}



// CreatePairPool creates a new pool for the given asset pair and adds
// the pool creator as a liquidity provider for the pool
// TODO
func (k msgServer) CreatePairPool(goCtx context.Context, msg *types.MsgCreatePairPool) (*types.MsgCreatePairPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO make pools meet some min amount

	// get assets from denom and amount
	assetA := types.NewPoolAsset(msg.DenomA, msg.AmountA)
	assetB := types.NewPoolAsset(msg.DenomB, msg.AmountB)
	// get pool name from asset pair
	poolName := types.GetPoolNameFromAssetPair(assetA, assetB)
	// check if pool exists using pool name
	if k.Keeper.HasPool(ctx, poolName) {
		// TODO add to errors
		return &types.MsgCreatePairPoolResponse{}, errors.New("Pool already exists")
	}
	// create pool and send coins
	poolId, shares, err := k.Keeper.CreatePoolPair(ctx, msg)
	if err != nil {
		// TODO add to errors
		return &types.MsgCreatePairPoolResponse{}, err
	}
	// create lp w/ shares created by pool and creator account address
	_, err := k.Keeper.CreateLiqProv(ctx, msg.Creator, shares) 
	if err != nil {
		// TODO add to errors
		return &types.MsgCreatePairPoolResponse{}, err
	}

	// TODO emit event here?

	return &types.MsgCreatePairPoolResponse{PoolId: poolId, Shares: shares}, nil
}

// JoinPairPool adds a new liquidity provider to an existing pool 
// TODO
func (k msgServer) JoinPairPool(goCtx context.Context, msg *types.MsgJoinPairPool) (*types.MsgJoinPairPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// get assets from denom and amount
	assetA := types.NewPoolAsset(msg.DenomA, msg.AmountA)
	assetB := types.NewPoolAsset(msg.DenomB, msg.AmountB)
	// get pool name f/rom asset pair
	poolName := types.GetPoolNameFromAssetPair(assetA, assetB)
	// check if pool exists using pool name
	if k.Keeper.HasPool(ctx, poolName) {
		// TODO add to errors
		return &types.MsgJoinPairPoolResponse{}, errors.New("Pool already exists")
	}
	// check if pool already has liquidity provider using pool name and msg creator
	if k.Keeper.HasLiqProv(ctx, poolName, msg.Creator) {
		return &types.MsgJoinPairPoolResponse{}, errors.New("Liquidity provider already exists")
	}
	// join the pool, send coins and get shares out
	poolId, shares, err := k.Keeper.JoinPoolPair(ctx, msg)
	if err != nil {
		return &types.MsgJoinPairPoolResponse{}, err
	}
	// use shares out to create & join w/ a new liq provider
	_, err := k.Keeper.JoinLiqProv(ctx, msg.Creator, shares) 
	if err != nil {
		return &types.MsgJoinPairPoolResponse{}, err
	}

	// TODO emit event here?

	return &types.MsgJoinPairPoolResponse{PoolId: poolName, Shares: shares}, nil
}

// ExitPairPool removes a liquidity provider entirely from a pool
// TODO
func (k msgServer) ExitPairPool(goCtx context.Context, msg *types.MsgExitPairPool) (*types.MsgExitPairPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if pool already exists
	if !k.Keeper.HasPool(ctx, msg.ShareDenom){
		// TODO add to errors
		return &types.MsgExitPairPoolResponse{}, errors.New("Pool DNE")
	}
	// remove all liquidity from provider, the provider from the list 
	// and get amount of shares removed 
	removedShares, err := k.Keeper.ExitLiqProv(ctx, msg.ShareDenom, msg.Creator)
	if err != nil {
		return &types.MsgExitPairPoolResponse{}, err
	}  
	// use removed shares amount to send coins and get assets out 
	poolName, assets, err := k.Keeper.ExitPoolPair(ctx, msg, removedShares)
	if err != nil {
		return &types.MsgExitPairPoolResponse{}, err
	}


	// TODO emit event here?

	return &types.MsgExitPairPoolResponse{PoolId: poolName, Assets: assets}, nil
}

// SwapPair swaps an asset in for an asset out using balancer's math
// TODO
func (k msgServer) SwapPair(goCtx context.Context, msg *types.MsgSwapPair) (*types.MsgSwapPairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get assets from denom and amount
	assetA := types.NewPoolAsset(msg.DenomA, msg.AmountA)
	assetB := types.NewPoolAsset(msg.DenomB, msg.AmountB)
	// get pool name f/rom asset pair
	poolName := types.GetPoolNameFromAssetPair(assetA, assetB)
	// check pool exists
	if !k.Keeper.HasPool(ctx, poolName) {
		// TODO add to errors
		return &types.MsgSwapPairResponse{}, errors.New("Pool DNE")
	}
	// update balances and send coins
	assetOut, err := k.Keeper.SwapAssetPair(ctx, msg)
	if err != nil {
		return &types.MsgSwapPairResponse{}, err
	}

	// TODO emit event here?

	return &types.MsgSwapPairResponse{Creator: msg.Creator, AssetOut: asset}, nil
} 

// AddLiquidityPair adds liquidity to an existing liqidity provider for a given pair's pool
// TODO
func (k msgServer) AddLiquidityPair(goCtx context.Context, msg *types.MsgAddLiquidityPair) (*types.MsgAddLiquidityPairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get assets from denom and amount
	assetA := types.NewPoolAsset(msg.DenomA, msg.AmountA)
	assetB := types.NewPoolAsset(msg.DenomB, msg.AmountB)
	// get pool name f/rom asset pair
	poolName := types.GetPoolNameFromAssetPair(assetA, assetB)
	// check pool exists
	if !k.Keeper.HasPool(ctx, poolName) {
		// TODO add to errors
		return &types.MsgSwapPairResponse{}, errors.New("Pool DNE")
	}
	// check if liquidity provider exists for pool
	if !k.Keeper.HasLiqProv(ctx, poolName, msg.Creator) {
		// TODO ADD TO errors
		return &Types.MsgSwapPairResponse{}, errors.New("Provider DNE")
	}
	// add to pool and get shares amount
	poolName, shares, err := k.Keeper.AddLiquidity(ctx, msg)
	if err != nil {
		return &types.MsgAddLiquidityPairResponse{}, err
	}
	// use shares amount to add liquidity to the existing provider
	if _, err := k.Keeper.AddLiqToProv(ctx, msg.Creator, shares); err != nil {
		return &types.MsgAddLiquidityPairResponse{}, err
	}
	
	// TODO emit event here?
	
	return &types.MsgAddLiquidityPairResponse{PoolId: poolName, Shares: shares}, nil
}

// RemoveLiquidityPair removes liquidity from a provider but does not remove the provider from the pool
// TODO
func (k msgServer) RemoveLiquidityPair(goCtx context.Context, msg *types.MsgRemoveLiquidityPair) (*types.MsgRemoveLiquidityPairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if pool exists
	if !k.Keeper.HasPool(ctx, msg.SharesDenom){
		// TODO add to errors
		return &types.MsgRemoveLiquidityPairResponse{}, errors.New("Pool DNE")
	}
	// check if provider exists for pool
	if !k.Keeper.HasLiqProv(ctx, msg.SharesDenom, msg.Creator) {
		// TODO add to errors
		return &types.MsgRemoveLiquidityPairResponse{}, errors.New("Provider DNE")
	}
	// remove pool, send coins, and get assets amount out
	assets, err := k.Keeper.RemoveLiquidity(ctx, msg)
	if err != nil {
		return &types.MsgRemoveLiquidityPairResponse{}, err
	}
	// get shares to remove
	shares := types.NewPoolShares(msg.ShareDenom, msg.SharesAmount)
	// remove liquidity from provider
	if err := k.Keeper.RemoveLiqFromProv(ctx, msg.SharesDenom, msg.Creator, shares); err != nil {
		return &types.MsgRemoveLiquidityPairResponse{}, err
	}
		
	// TODO emit event here?

	return &types.MsgRemoveLiquidityPairResponse{Creator: msg.Creator, Assets: assets}, nil
}

