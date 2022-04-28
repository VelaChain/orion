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

	return &types.MsgCreatePairPoolResponse{PoolId: name, Shares: shares}, nil
}

// TODO
func (k msgServer) JoinPairPool(goCtx context.Context, msg *types.MsgJoinPairPool) (*types.MsgJoinPairPoolResponse, error) {
	return &types.MsgJoinPairPoolResponse{}, nil
}

// TODO
func (k msgServer) ExitPairPool(goCtx context.Context, msg *types.MsgExitPairPool) (*types.MsgExitPairPoolResponse, error) {
	return &types.MsgExitPairPoolResponse{}, nil
}

// TODO
func (k msgServer) SwapPair(goCtx context.Context, msg *types.MsgSwapPair) (*types.MsgSwapPairResponse, error) {
	return &types.MsgSwapPairResponse{}, nil
} 

// TODO
func (k msgServer) AddLiquidityPair(goCtx context.Context, msg *types.MsgAddLiquidityPair) (*types.MsgAddLiquidityPairResponse, error) {
	return &types.MsgAddLiquidityPairResponse{}, nil
}

// TODO
func (k msgServer) RemoveLiquidityPair(goCtx context.Context, msg *types.MsgRemoveLiquidityPair) (*types.MsgRemoveLiquidityPairResponse, error) {
	return &types.MsgRemoveLiquidityPairResponse{}, nil
}

