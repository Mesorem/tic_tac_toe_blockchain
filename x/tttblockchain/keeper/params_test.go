package keeper_test

import (
	"testing"

	testkeeper "github.com/Mesorem/ttt_blockchain/testutil/keeper"
	"github.com/Mesorem/ttt_blockchain/x/tttblockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TttblockchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
