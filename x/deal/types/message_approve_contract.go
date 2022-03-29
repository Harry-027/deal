package types

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveContract = "approve_contract"

var _ sdk.Msg = &MsgApproveContract{}

func NewMsgApproveContract(creator string, dealId string, contractId string) *MsgApproveContract {
	return &MsgApproveContract{
		Creator:    creator,
		DealId:     dealId,
		ContractId: contractId,
	}
}

func (msg *MsgApproveContract) Route() string {
	return RouterKey
}

func (msg *MsgApproveContract) Type() string {
	return TypeMsgApproveContract
}

func (msg *MsgApproveContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgApproveContract) DealHandlerValidation(goCtx context.Context, contract *NewContract) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != contract.Consumer {
		return ErrInvalidConsumer
	}

	expiry, err := time.Parse(time.RFC3339, contract.Expiry)
	if err != nil {
		panic("invalid expiry time")
	}

	if ctx.BlockTime().Before(expiry) {
		return ErrContractExpired
	}

	if contract.Status != COMMITTED {
		return ErrNotCommitted
	}

	return nil
}
