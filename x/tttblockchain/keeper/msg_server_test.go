package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Mesorem/ttt_blockchain/testutil/keeper"
	"github.com/Mesorem/ttt_blockchain/x/tttblockchain/keeper"
	"github.com/Mesorem/ttt_blockchain/x/tttblockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TttblockchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
