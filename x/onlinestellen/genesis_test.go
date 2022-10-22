package onlinestellen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "onlinestellen/testutil/keeper"
	"onlinestellen/testutil/nullify"
	"onlinestellen/x/onlinestellen"
	"onlinestellen/x/onlinestellen/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OnlinestellenKeeper(t)
	onlinestellen.InitGenesis(ctx, *k, genesisState)
	got := onlinestellen.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
