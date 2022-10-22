package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "onlinestellen/testutil/keeper"
	"onlinestellen/x/onlinestellen/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OnlinestellenKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
