package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/Harry-027/deal/testutil/keeper"
	"github.com/Harry-027/deal/testutil/nullify"
	"github.com/Harry-027/deal/x/deal/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestNewDealQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DealKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNewDeal(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNewDealRequest
		response *types.QueryGetNewDealResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetNewDealRequest{
				Index: strconv.FormatUint(msgs[0].DealId, 10),
			},
			response: &types.QueryGetNewDealResponse{NewDeal: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetNewDealRequest{
				Index: strconv.FormatUint(msgs[1].DealId, 10),
			},
			response: &types.QueryGetNewDealResponse{NewDeal: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetNewDealRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NewDeal(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestNewDealQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DealKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNewDeal(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllNewDealRequest {
		return &types.QueryAllNewDealRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NewDealAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NewDeal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NewDeal),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NewDealAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NewDeal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NewDeal),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.NewDealAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.NewDeal),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.NewDealAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
