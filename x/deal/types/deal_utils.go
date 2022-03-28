package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (newDeal *NewDeal) GetOwnerAddress() (owner sdk.AccAddress, err error) {
	owner, errInvalidOwner := sdk.AccAddressFromBech32(newDeal.Owner)
	return owner, sdkerrors.Wrapf(errInvalidOwner, ErrInvalidOwner.Error(), newDeal.Owner)
}

func (newDeal *NewDeal) GetVendorAddress() (vendor sdk.AccAddress, err error) {
	vendor, errInvalidVendor := sdk.AccAddressFromBech32(newDeal.Vendor)
	return vendor, sdkerrors.Wrapf(errInvalidVendor, ErrInvalidVendor.Error(), newDeal.Vendor)
}

func (newDeal *NewDeal) ValidateCommission() (err error) {
	if 1 <= newDeal.Commission && 100 >= newDeal.Commission {
		return nil
	}
	return ErrInvalidCommission
}

func (newDeal *NewDeal) Validate() (err error) {
	_, err = newDeal.GetOwnerAddress()
	if err != nil {
		return err
	}

	_, err = newDeal.GetVendorAddress()
	if err != nil {
		return err
	}

	err = newDeal.ValidateCommission()
	return err
}
