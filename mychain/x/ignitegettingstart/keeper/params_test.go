package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/WonderJL/ignite-getting-start/testutil/keeper"
	"github.com/WonderJL/ignite-getting-start/x/ignitegettingstart/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.IgnitegettingstartKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
