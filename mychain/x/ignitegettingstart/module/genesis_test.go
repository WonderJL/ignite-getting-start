package ignitegettingstart_test

import (
	"testing"

	keepertest "github.com/WonderJL/ignite-getting-start/testutil/keeper"
	"github.com/WonderJL/ignite-getting-start/testutil/nullify"
	ignitegettingstart "github.com/WonderJL/ignite-getting-start/x/ignitegettingstart/module"
	"github.com/WonderJL/ignite-getting-start/x/ignitegettingstart/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IgnitegettingstartKeeper(t)
	ignitegettingstart.InitGenesis(ctx, k, genesisState)
	got := ignitegettingstart.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
