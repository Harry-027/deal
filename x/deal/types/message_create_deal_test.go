package types

import (
	"testing"

	"github.com/Harry-027/deal/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateDeal_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateDeal
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateDeal{
				Creator: "invalid_address",
				Vendor:  "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateDeal{
				Creator: sample.AccAddress(),
				Vendor:  sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
