package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDeal{}, "deal/CreateDeal", nil)
	cdc.RegisterConcrete(&MsgCreateContract{}, "deal/CreateContract", nil)
	cdc.RegisterConcrete(&MsgCommitContract{}, "deal/CommitContract", nil)
	cdc.RegisterConcrete(&MsgApproveContract{}, "deal/ApproveContract", nil)
	cdc.RegisterConcrete(&MsgShipOrder{}, "deal/ShipOrder", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDeal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCommitContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgShipOrder{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
