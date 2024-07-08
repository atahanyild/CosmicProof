package ytubclambda_test

import (
	"testing"

	keepertest "github.com/atahanyild/ytubc-lambda/testutil/keeper"
	"github.com/atahanyild/ytubc-lambda/testutil/nullify"
	ytubclambda "github.com/atahanyild/ytubc-lambda/x/ytubclambda/module"
	"github.com/atahanyild/ytubc-lambda/x/ytubclambda/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.YtubclambdaKeeper(t)
	ytubclambda.InitGenesis(ctx, k, genesisState)
	got := ytubclambda.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
