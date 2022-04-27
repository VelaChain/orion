package chios_test

import (
	"testing"

	keepertest "github.com/VelaChain/orion/testutil/keeper"
	"github.com/VelaChain/orion/testutil/nullify"
	"github.com/VelaChain/orion/x/chios"
	"github.com/VelaChain/orion/x/chios/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChiosKeeper(t)
	chios.InitGenesis(ctx, *k, genesisState)
	got := chios.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
