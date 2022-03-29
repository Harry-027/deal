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
	ErrContractNotFound  = sdkerrors.Register(ModuleName, 1105, "Contract not found")
	ErrContractExpired   = sdkerrors.Register(ModuleName, 1106, "Contract already expired")
	ErrInvalidETA        = sdkerrors.Register(ModuleName, 1107, "Invalid ETA")
	ErrDescLength        = sdkerrors.Register(ModuleName, 1108, "Desc length too large")
)
