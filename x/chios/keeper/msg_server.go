package keeper

import (
	"context"

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

	if k.Has(ctx, types.GetPoolKeyFromPoolName(types.GetPoolNameFromAssets(types.NewPoolAssets(types.NewPoolAsset(msg.DenomA, msg.AmountA), types.NewPoolAsset(msg.DenomB, msg.AmountB))))){
		// TODO add to errors
		return &types.MsgCreatePairPoolResponse{}, errors.New("Pool already exists")
	}

	name, shares, err := k.Keeper.CreatePoolPair(ctx, msg)
	if err != nil {
		return &types.MsgCreatePairPoolResponse{}, err
	}

	// TODO emit event here?

	return &types.MsgCreatePairPoolResponse{PoolId: name, Shares: shares}, nil
}

// TODO
func (k msgServer) JoinPairPool(goCtx context.Context, msg *types.MsgJoinPairPool) (*types.MsgJoinPairPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// check if pool already exists
	if !k.Has(ctx, types.GetPoolKeyFromPoolName(types.GetPoolNameFromAssets(types.NewPoolAssets(types.NewPoolAsset(msg.DenomA, msg.AmountA), types.NewPoolAsset(msg.DenomB, msg.AmountB))))){
		// TODO add to errors
		return &types.MsgJoinPairPoolResponse{}, errors.New("Pool DNE")
	}

	poolName, shares, err := k.Keeper.JoinPoolPair(ctx, msg)
	if err != nil {
		return &types.MsgJoinPairPoolResponse{}, err
	}
	
	// TODO emit event here?

	return &types.MsgJoinPairPoolResponse{PoolId: poolName, Shares: shares}, nil
}

// TODO
func (k msgServer) ExitPairPool(goCtx context.Context, msg *types.MsgExitPairPool) (*types.MsgExitPairPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if pool already exists
	if !k.Has(ctx, types.GetPoolKeyFromPoolName(types.GetPoolNameFromAssets(types.NewPoolAssets(types.NewPoolAsset(msg.DenomA, msg.AmountA), types.NewPoolAsset(msg.DenomB, msg.AmountB))))){
		// TODO add to errors
		return &types.MsgExitPairPoolResponse{}, errors.New("Pool DNE")
	}
	
	poolName, assets, err := k.Keeper.ExitPoolPair(ctx, msg)
	if err != nil {
		return &types.MsgExitPairPoolResponse{}, err
	}

	// TODO emit event here?

	return &types.MsgExitPairPoolResponse{PoolId: poolName, Assets: assets} nil
}

// TODO
func (k msgServer) SwapPair(goCtx context.Context, msg *types.MsgSwapPair) (*types.MsgSwapPairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if pool already exists
	if !k.Has(ctx, types.GetPoolKeyFromPoolName(types.GetPoolNameFromAssets(types.NewPoolAssets(types.NewPoolAsset(msg.DenomA, msg.AmountA), types.NewPoolAsset(msg.DenomB, msg.AmountB))))){
		return &types.MsgSwapPairResponse{}, errors.New("Pool DNE")
	}

	asset, err := k.Keeper.SwapAssetPair(ctx, msg)
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
	if !k.Has(ctx, types.GetPoolKeyFromPoolName(types.GetPoolNameFromAssets(types.NewPoolAssets(types.NewPoolAsset(msg.DenomA, msg.AmountA), types.NewPoolAsset(msg.DenomB, msg.AmountB))))){
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

	if !k.Has(ctx, types.GetPoolKeyFromPoolName(types.GetPoolNameFromAssets(types.NewPoolAssets(types.NewPoolAsset(msg.DenomA, msg.AmountA), types.NewPoolAsset(msg.DenomB, msg.AmountB))))){
		return &types.MsgRemoveLiquidityPairResponse{}, errors.New("Pool DNE")
	}

	assets, err := k.Keeper.RemoveLiquidity(ctx, msg)
	if err != nil {
		return &types.MsgRemoveLiquidityPairResponse{}, err
	}
	// TODO emit event here?

	return &types.MsgRemoveLiquidityPairResponse{Creator: msg.Creator, Assets: assets}, nil
}

