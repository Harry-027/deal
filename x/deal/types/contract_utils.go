package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// utility func & vars

const (
	INITIATED      string = "INITIATED"
	COMMITTED             = "COMMITTED"
	APPROVED              = "APPROVED"
	INDELIVERY            = "IN-DELIVERY"
	COMPLETED             = "COMPLETED"
	CANCELLED             = "CANCELLED"
	TIME_FORMAT           = "2006-01-02 15:04:05.999999999 +0000 UTC"
	TOKEN                 = "token"
	IDVALUE               = "IdValue"
	OWNER                 = "Owner"
	VENDOR                = "Vendor"
	CONSUMER              = "Consumer"
	START_TIME            = "StartTime"
	VENDOR_ETA            = "VendorETA"
	OWNER_ETA             = "OwnerETA"
	CREATE_GAS     uint64 = 200
	PROCESS_GAS    uint64 = 100
	SETTLEMENT_GAS uint64 = 300
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

func (newContract *NewContract) GetCoin(amount uint64) (fees sdk.Coin) {
	return sdk.NewCoin(TOKEN, sdk.NewInt(int64(amount)))
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
