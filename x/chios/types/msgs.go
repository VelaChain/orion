package types

import (
	//"fmt"
	//"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

const (
	TypeMsgCreatePairPool = "create_pair_pool"
	TypeMsgJoinPairPool = "join_pair_pool"
	TypeMsgExitPairPool = "exit_pair_pool"
	TypeMsgSwapPair = "swap_pair"
	TypeMsgAddLiquidityPair = "add_liquidity_pair"
	TypeMsgRemoveLiquidityPair = "remove_liquidity_pair"
)

var (
	_ sdk.Msg = &MsgCreatePairPool{}
	_ sdk.Msg = &MsgJoinPairPool{}
	_ sdk.Msg = &MsgExitPairPool{}
	_ sdk.Msg = &MsgSwapPair{}
	_ sdk.Msg = &MsgAddLiquidityPair{}
	_ sdk.Msg = &MsgRemoveLiquidityPair{}
)

func NewMsgCreatePairPool(creator string, denomA string, amountA sdk.Int, denomB string, amountB sdk.Int, sharesOut sdk.Int) *MsgCreatePairPool{
	return &MsgCreatePairPool{
		Creator:	creator,
		DenomA:		denomA,
		AmountA:	amountA,
		DenomB:		denomB,
		AmountB:	amountB,
		SharesOut:	sharesOut,
	}
}

func NewMsgJoinPairPool(creator string, denomA string, amountA sdk.Int, denomB string, amountB sdk.Int, sharesOut sdk.Int) *MsgJoinPairPool{
	return &MsgJoinPairPool{
		Creator:	creator,
		DenomA:		denomA,
		AmountA:	amountA,
		DenomB:		denomB,
		AmountB:	amountB,
		SharesOut:	sharesOut,
	}
}

func NewMsgExitPairPool(creator string, sharesDenom string, sharesAmount sdk.Int) *MsgExitPairPool{
	return &MsgExitPairPool{
		Creator:		creator,
		ShareDenom:		sharesDenom,
		ShareAmount:	sharesAmount, 
	}
}

func NewMsgSwapPair(creator string, denomIn string, amountIn sdk.Int, denomOut string, amountOut sdk.Int) *MsgSwapPair {
	return &MsgSwapPair {
		Creator:		creator,
		DenomIn:		denomIn,
		AmountIn:		amountIn,
		DenomOut:		denomOut,
		MinAmountOut:	amountOut,
	}
}

func NewMsgAddLiquidityPair(creator string, denomA string, amountA sdk.Int, denomB string, amountB sdk.Int, sharesOut sdk.Int) *MsgAddLiquidityPair{
	return &MsgAddLiquidityPair{
		Creator:	creator,
		DenomA:		denomA,
		AmountA:	amountA,
		DenomB:		denomB,
		AmountB:	amountB,
		SharesOut:	sharesOut,
	}
}

func NewMsgRemoveLiquidityPair(creator string, sharesDenom string, sharesAmount sdk.Int) *MsgRemoveLiquidityPair {
	return &MsgRemoveLiquidityPair {
		Creator:		creator,
		SharesDenom:	sharesDenom,
		SharesAmount:	sharesAmount,
	}
}

func (msg *MsgCreatePairPool) Route() string {
	return RouterKey
}

func (msg *MsgJoinPairPool) Route() string {
	return RouterKey
}

func (msg *MsgExitPairPool) Route() string {
	return RouterKey
}

func (msg *MsgSwapPair) Route() string {
	return RouterKey
}


func (msg *MsgAddLiquidityPair) Route() string {
	return RouterKey
}

func (msg *MsgRemoveLiquidityPair) Route() string {
	return RouterKey
}

func (msg *MsgCreatePairPool) Type() string {
	return TypeMsgCreatePairPool
}

func (msg *MsgJoinPairPool) Type() string {
	return TypeMsgJoinPairPool
}

func (msg *MsgExitPairPool) Type() string {
	return TypeMsgExitPairPool
}

func (msg *MsgSwapPair) Type() string {
	return TypeMsgSwapPair
}

func (msg *MsgAddLiquidityPair) Type() string {
	return TypeMsgAddLiquidityPair
}

func (msg *MsgRemoveLiquidityPair) Type() string {
	return TypeMsgRemoveLiquidityPair
}

func (msg *MsgCreatePairPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgJoinPairPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgExitPairPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSwapPair) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddLiquidityPair) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveLiquidityPair) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// TODO
func (msg *MsgCreatePairPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}

// TODO
func (msg *MsgJoinPairPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}

// TODO
func (msg *MsgExitPairPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}

// TODO
func (msg *MsgSwapPair) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}

// TODO
func (msg *MsgAddLiquidityPair) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}

// TODO
func (msg *MsgRemoveLiquidityPair) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
