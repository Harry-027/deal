package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCommitContract = "commit_contract"

var _ sdk.Msg = &MsgCommitContract{}

func NewMsgCommitContract(creator, dealId, contractId, vendorETA string) *MsgCommitContract {
	return &MsgCommitContract{
		Creator:    creator,
		DealId:     dealId,
		ContractId: contractId,
		VendorETA: vendorETA,
	}
}

func (msg *MsgCommitContract) Route() string {
	return RouterKey
}

func (msg *MsgCommitContract) Type() string {
	return TypeMsgCommitContract
}

func (msg *MsgCommitContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCommitContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCommitContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
