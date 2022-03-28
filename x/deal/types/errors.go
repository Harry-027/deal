package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/deal module sentinel errors
var (
	ErrInvalidOwner      = sdkerrors.Register(ModuleName, 1100, "Owner does not exists")
	ErrInvalidVendor     = sdkerrors.Register(ModuleName, 1101, "Vendor does not exists")
	ErrInvalidCommission = sdkerrors.Register(ModuleName, 1102, "Invalid commission")
)
