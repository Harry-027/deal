package types

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelOrder = "cancel_order"

var _ sdk.Msg = &MsgCancelOrder{}

func NewMsgCancelOrder(creator string, dealId string, contractId string) *MsgCancelOrder {
	return &MsgCancelOrder{
		Creator:    creator,
		DealId:     dealId,
		ContractId: contractId,
	}
}

func (msg *MsgCancelOrder) Route() string {
	return RouterKey
}

func (msg *MsgCancelOrder) Type() string {
	return TypeMsgCancelOrder
}

func (msg *MsgCancelOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgCancelOrder) DealHandlerValidation(goCtx context.Context, contract *NewContract) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != contract.Consumer {
		return ErrInvalidConsumer
	}

	if contract.Status != APPROVED {
		return ErrNotApproved
	}

	startTime, err := time.Parse(time.RFC3339, contract.StartTime)
	if err != nil {
		panic("invalid start time")
	}

	deliveryExpectedTime := startTime.Add(time.Duration(contract.OwnerETA))
	timeLimit := uint32(ctx.BlockTime().Sub(deliveryExpectedTime).Minutes())
	if timeLimit < 20 {
		return ErrRefund
	}
	return nil
}