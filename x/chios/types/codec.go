package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	// this line is used by starport scaffolding # 1
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreatePairPool{}, "chios/CreatePairPool", nil)
	cdc.RegisterConcrete(&MsgJoinPairPool{}, "chios/JoinPairPool", nil)
	cdc.RegisterConcrete(&MsgExitPairPool{}, "chios/ExitPairPool", nil)
	cdc.RegisterConcrete(&MsgSwapPair{}, "chios/SwapPair", nil)
	cdc.RegisterConcrete(&MsgAddLiquidityPair{}, "chios/AddLiquidityPair", nil)
	cdc.RegisterConcrete(&MsgRemoveLiquidityPair{}, "chios/RemoveLiquidityPair", nil)


}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePairPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
        &MsgJoinPairPool{},
    )
	registry.RegisterImplementations((*sdk.Msg)(nil),
        &MsgExitPairPool{},
    )
	registry.RegisterImplementations((*sdk.Msg)(nil),
        &MsgSwapPair{},
    )
	registry.RegisterImplementations((*sdk.Msg)(nil),
        &MsgAddLiquidityPair{},
    )
	registry.RegisterImplementations((*sdk.Msg)(nil),
        &MsgRemoveLiquidityPair{},
    )


	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
