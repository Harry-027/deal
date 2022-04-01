package types

import (
	"testing"

	"github.com/Harry-027/deal/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateContract_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateContract
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateContract{
				Creator:  "invalid_address",
				Consumer: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateContract{
				Creator:  sample.AccAddress(),
				Consumer: sample.AccAddress(),
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
