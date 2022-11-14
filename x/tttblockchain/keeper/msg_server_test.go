package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Mesorem/ttt_blockchain/testutil/keeper"
	"github.com/Mesorem/ttt_blockchain/x/tttblockchain/keeper"
	"github.com/Mesorem/ttt_blockchain/x/tttblockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TttblockchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func Test_CreateGame(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	msgResponse, err := msgServer.CreateGame(ctx, &types.MsgCreateGame{
		Creator: alice,
		Cross:   bob,
		Circle:  carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "",
	}, *msgResponse)
}
