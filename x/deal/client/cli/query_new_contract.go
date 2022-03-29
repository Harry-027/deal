package cli

import (
	"context"

	"github.com/Harry-027/deal/x/deal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListNewContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-new-contract [dealId]",
		Short: "list all newContract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			dealId := args[0]
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNewContractRequest{
				Pagination: pageReq,
				DealId: dealId,
			}

			res, err := queryClient.NewContractAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowNewContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-new-contract [dealId] [contractId]",
		Short: "shows a newContract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			dealId := args[0]
			contractId := args[1]

			params := &types.QueryGetNewContractRequest{
				DealId: dealId,
				ContractId: contractId,
			}

			res, err := queryClient.NewContract(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
