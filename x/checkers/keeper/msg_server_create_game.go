package keeper

import (
	"context"
	"strconv"

	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	nextGame, found := k.Keeper.GetNextGame(ctx)
	if !found {
		panic("nextGame index failure!")
	}

	newIndex := strconv.FormatUint(nextGame.IdValue, 10)

	storedGame := types.WaitingGame{
		IdValue: nextGame.IdValue,
		Creator: msg.Creator,
		Index:   newIndex,
	}

	k.Keeper.SetWaitingGame(ctx, storedGame)

	nextGame.IdValue++
	k.Keeper.SetNextGame(ctx, nextGame)

	return &types.MsgCreateGameResponse{
		IdValue: newIndex,
	}, nil
}
