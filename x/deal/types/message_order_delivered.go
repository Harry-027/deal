package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOrderDelivered = "order_delivered"

var _ sdk.Msg = &MsgOrderDelivered{}

func NewMsgOrderDelivered(creator string, dealId string, contractId string) *MsgOrderDelivered {
	return &MsgOrderDelivered{
		Creator:    creator,
		DealId:     dealId,
		ContractId: contractId,
	}
}

func (msg *MsgOrderDelivered) Route() string {
	return RouterKey
}

func (msg *MsgOrderDelivered) Type() string {
	return TypeMsgOrderDelivered
}

func (msg *MsgOrderDelivered) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOrderDelivered) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOrderDelivered) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
