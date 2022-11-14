package tttblockchain_test

import (
	"testing"

	keepertest "github.com/Mesorem/ttt_blockchain/testutil/keeper"
	"github.com/Mesorem/ttt_blockchain/testutil/nullify"
	"github.com/Mesorem/ttt_blockchain/x/tttblockchain"
	"github.com/Mesorem/ttt_blockchain/x/tttblockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: types.SystemInfo{
			NextId: 27,
		},
		StoredGameList: []types.StoredGame{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TttblockchainKeeper(t)
	tttblockchain.InitGenesis(ctx, *k, genesisState)
	got := tttblockchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.StoredGameList, got.StoredGameList)
	// this line is used by starport scaffolding # genesis/test/assert
}
