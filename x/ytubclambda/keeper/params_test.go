package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/atahanyild/ytubc-lambda/testutil/keeper"
	"github.com/atahanyild/ytubc-lambda/x/ytubclambda/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.YtubclambdaKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
