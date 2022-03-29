package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/deal module sentinel errors
var (
	ErrInvalidOwner      = sdkerrors.Register(ModuleName, 1100, "Owner does not exists")
	ErrInvalidVendor     = sdkerrors.Register(ModuleName, 1101, "Vendor does not exists")
	ErrInvalidConsumer   = sdkerrors.Register(ModuleName, 1102, "Consumer does not exists")
	ErrInvalidCommission = sdkerrors.Register(ModuleName, 1103, "Invalid commission")
	ErrDealNotFound      = sdkerrors.Register(ModuleName, 1104, "Deal not found")
	ErrInvalidETA        = sdkerrors.Register(ModuleName, 1105, "Invalid ETA")
	ErrDescLength        = sdkerrors.Register(ModuleName, 1106, "Desc length too large")
)
