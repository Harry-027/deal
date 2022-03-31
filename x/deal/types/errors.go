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
	ErrInvalidTime        = sdkerrors.Register(ModuleName, 1107, "Invalid Time")
	ErrDescLength        = sdkerrors.Register(ModuleName, 1108, "Desc length too large")
	ErrPaymentFailed     = sdkerrors.Register(ModuleName, 1109, "Payment Failed")
	ErrNotCommitted      = sdkerrors.Register(ModuleName, 1110, "Contract not in commit stage")
	ErrNotShipped        = sdkerrors.Register(ModuleName, 1111, "Order not yet shipped or completed")
	ErrNotApproved       = sdkerrors.Register(ModuleName, 1112, "Contract is not in approved stage")
	ErrRefund            = sdkerrors.Register(ModuleName, 1113, "Refund applicable after 20 mins delay")
	ErrVendorETA         = sdkerrors.Register(ModuleName, 1114, "Vendor ETA should be less than or equal to half of that of Owner ETA")
)
