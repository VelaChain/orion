package chios

import (
	"fmt"

	"github.com/VelaChain/orion/x/chios/keeper"
	"github.com/VelaChain/orion/x/chios/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	// this line is used by starport scaffolding # handler/msgServer
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreatePairPool:
			res, err := msgServer.CreatePairPool(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		
		case *types.MsgJoinPairPool:
			res, err := msgServer.JoinPairPool(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		
		case *types.MsgExitPairPool:
			res, err := msgServer.ExitPairPool(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		
		case *types.MsgSwapPair:
			res, err := msgServer.SwapPair(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)			
		
		case *types.MsgAddLiquidityPair:
			res, err := msgServer.AddLiquidityPair(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		
		case *types.MsgRemoveLiquidityPair:
			res, err := msgServer.RemoveLiquidityPair(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		
		// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
