package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
    INITIATED string = "INITIATED"
    COMMITTED = "COMMITTED"
    APPROVED = "APPROVED"
    SHIPPED = "SHIPPED"
	INDELIVERY =  "IN-DELIVERY"
	DELIVERED = "DELIVERED"
)

func (newContract *NewContract) ValidateDesc() (err error) {
	if len(newContract.Desc) > 20 {
		return ErrDescLength
	}
	return nil
}

func (newContract *NewContract) GetConsumerAddress() (consumer sdk.AccAddress, err error) {
	consumer, errInvalidConsumer := sdk.AccAddressFromBech32(newContract.Consumer)
	return consumer, sdkerrors.Wrapf(errInvalidConsumer, ErrInvalidConsumer.Error(), newContract.Consumer)
}

func (newContract *NewContract) Validate() (err error) {
	err = newContract.ValidateDesc()
	if err != nil {
		return err
	}

	_, err = newContract.GetConsumerAddress()
	if err != nil {
		return err
	}
	return nil
}