package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/WonderJL/ignite-getting-start/testutil/keeper"
	"github.com/WonderJL/ignite-getting-start/x/ignitegettingstart/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.IgnitegettingstartKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
