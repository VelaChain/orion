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
	// create pool
	poolId, shares, err := k.Keeper.CreatePoolPair(ctx, msg)
	if err != nil {
		// TODO add to errors
		return &types.MsgCreatePairPoolResponse{}, err
	}
	// get creator account address
	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		// TODO add to errors
		return &types.MsgCreatePairPoolResponse{}, err
	}
	// create lp w/ shares and creator account address
	_, err := k.Keeper.CreateLiqProv(ctx, accAddr, shares) 
	if err != nil {
		// TODO add to errors
		return &types.MsgCreatePairPoolResponse{}, err
	}

	// TODO emit event here?

	return &types.MsgCreatePairPoolResponse{PoolId: poolId, Shares: shares}, nil
}

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
		return &types.MsgCreatePairPoolResponse{}, errors.New("Pool already exists")
	}

	poolId, shares, err := k.Keeper.JoinPoolPair(ctx, msg)
	if err != nil {
		return &types.MsgJoinPairPoolResponse{}, err
	}
	
	_, err := k.Keeper.JoinLiqProv(ctx, msg.Creator, shares) 
	if err != nil {
		return &types.MsgJoinPairPoolRespoinse{}, err
	}

	// TODO emit event here?

	return &types.MsgJoinPairPoolResponse{PoolId: poolName, Shares: shares}, nil
}

// TODO
func (k msgServer) ExitPairPool(goCtx context.Context, msg *types.MsgExitPairPool) (*types.MsgExitPairPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if pool already exists
	if !k.Keeper.HasPool(ctx, msg.ShareDenom){
		// TODO add to errors
		return &types.MsgExitPairPoolResponse{}, errors.New("Pool DNE")
	}
	
	removedShares, err := k.Keeper.ExitLiqProv(ctx, msg.ShareDenom, msg.Creator)
	if err != nil {
		return &types.MsgExitPairPoolResponse{}, err
	}  

	poolName, assets, err := k.Keeper.ExitPoolPair(ctx, msg, removedShares)
	if err != nil {
		return &types.MsgExitPairPoolResponse{}, err
	}


	// TODO emit event here?

	return &types.MsgExitPairPoolResponse{PoolId: poolName, Assets: assets}, nil
}

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

	assetOut, err := k.Keeper.SwapAssetPair(ctx, msg)
	if err != nil {
		return &types.MsgSwapPairResponse{}, err
	}

	// TODO emit event here?

	return &types.MsgSwapPairResponse{Creator: msg.Creator, AssetOut: asset}, nil
} 

// TODO
func (k msgServer) AddLiquidityPair(goCtx context.Context, msg *types.MsgAddLiquidityPair) (*types.MsgAddLiquidityPairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if pool exists
	if !ctx.KVStore(k.storeKey).Has(types.GetPoolKeyFromPoolName(types.GetPoolNameFromAssets(types.NewPoolAssets(types.NewPoolAsset(msg.DenomA, msg.AmountA), types.NewPoolAsset(msg.DenomB, msg.AmountB))))){
		return &types.MsgAddLiquidityPairResponse{}, errors.New("Pool DNE")
	}

	poolName, shares, err := k.Keeper.AddLiquidity(ctx, msg)
	if err != nil {
		return &types.MsgAddLiquidityPairResponse{}, err
	}

	// TODO emit event here?
	
	return &types.MsgAddLiquidityPairResponse{PoolId: poolName, Shares: shares}, nil
}

// TODO
func (k msgServer) RemoveLiquidityPair(goCtx context.Context, msg *types.MsgRemoveLiquidityPair) (*types.MsgRemoveLiquidityPairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !ctx.KVStore(k.storeKey).Has(types.GetPoolKeyFromPoolName(msg.SharesDenom)){
		return &types.MsgRemoveLiquidityPairResponse{}, errors.New("Pool DNE")
	}

	assets, err := k.Keeper.RemoveLiquidity(ctx, msg)
	if err != nil {
		return &types.MsgRemoveLiquidityPairResponse{}, err
	}
	// TODO emit event here?

	return &types.MsgRemoveLiquidityPairResponse{Creator: msg.Creator, Assets: assets}, nil
}

