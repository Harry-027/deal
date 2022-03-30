package cli

import (
	"strconv"

	"github.com/Harry-027/deal/x/deal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCommitContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit-contract [deal-id] [contract-id] [ETA]",
		Short: "Broadcast message commitContract",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDealId := args[0]
			argContractId := args[1]
			argVendorETA := args[2]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCommitContract(
				clientCtx.GetFromAddress().String(),
				argDealId,
				argContractId,
				argVendorETA,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
