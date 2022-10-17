package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "hellp/testutil/keeper"
	"hellp/x/hellp/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HellpKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
