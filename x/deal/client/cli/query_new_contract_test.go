package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Harry-027/deal/testutil/network"
	"github.com/Harry-027/deal/testutil/nullify"
	"github.com/Harry-027/deal/x/deal/client/cli"
	"github.com/Harry-027/deal/x/deal/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithNewContractObjects(t *testing.T, n int) (*network.Network, []types.NewContract) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 1; i < n; i++ {
		newContract := types.NewContract{
			DealId:     "1",
			ContractId: strconv.Itoa(i),
		}
		nullify.Fill(&newContract)
		state.NewContractList = append(state.NewContractList, newContract)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.NewContractList
}

func TestShowNewContract(t *testing.T) {
	net, objs := networkWithNewContractObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc    string
		idIndex string

		args []string
		err  error
		obj  types.NewContract
	}{
		{
			desc:    "found",
			idIndex: objs[0].ContractId,

			args: common,
			obj:  objs[0],
		},
		{
			desc:    "not found",
			idIndex: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				"1",
				tc.idIndex,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowNewContract(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetNewContractResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.NewContract)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.NewContract),
				)
			}
		})
	}
}

func TestListNewContract(t *testing.T) {
	net, objs := networkWithNewContractObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(obj types.NewContract, next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			obj.DealId,
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(objs[i], nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListNewContract(), args)
			require.NoError(t, err)
			var resp types.QueryAllNewContractResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.NewContract), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.NewContract),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(objs[i], next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListNewContract(), args)
			require.NoError(t, err)
			var resp types.QueryAllNewContractResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.NewContract), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.NewContract),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(objs[0], nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListNewContract(), args)
		require.NoError(t, err)
		var resp types.QueryAllNewContractResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.NewContract),
		)
	})
}
