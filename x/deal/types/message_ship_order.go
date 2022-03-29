package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgShipOrder = "ship_order"

var _ sdk.Msg = &MsgShipOrder{}

func NewMsgShipOrder(creator string, dealId string, contractId string) *MsgShipOrder {
	return &MsgShipOrder{
		Creator:        creator,
		DealId:         dealId,
		ContractId:     contractId,
	}
}

func (msg *MsgShipOrder) Route() string {
	return RouterKey
}

func (msg *MsgShipOrder) Type() string {
	return TypeMsgShipOrder
}

func (msg *MsgShipOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgShipOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgShipOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
