package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/alice/checkers/testutil/keeper"
	"github.com/alice/checkers/testutil/nullify"
	"github.com/alice/checkers/x/checkers/keeper"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNWaitingGame(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.WaitingGame {
	items := make([]types.WaitingGame, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetWaitingGame(ctx, items[i])
	}
	return items
}

func TestWaitingGameGet(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNWaitingGame(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetWaitingGame(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestWaitingGameRemove(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNWaitingGame(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWaitingGame(ctx,
			item.Index,
		)
		_, found := keeper.GetWaitingGame(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestWaitingGameGetAll(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNWaitingGame(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWaitingGame(ctx)),
	)
}
