package keeper

import (
	"context"

	"github.com/VelaChain/orion/x/chios/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
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
	return &types.MsgCreatePairPoolResponse{}, nil
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
func (k msgServer) AddLiquidityPair(goCtx context.Context, msg *types.MsgAddLiquidityPair) (*types.AddLiquidityPairResponse, error) {
	return &types.MsgAddLiquidityPairResponse{}, nil
}

// TODO
func (k msgServer) RemoveLiquidityPair(goCtx context.Context, msg *types.MsgRemoveLiquidityPair) (*types.MsgRemoveLiquidityPairResponse, error) {
	return &types.MsgRemoveLiquidityPairResponse{}, nil
}

