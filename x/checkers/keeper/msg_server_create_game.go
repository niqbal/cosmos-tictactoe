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

/*
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	nextGame, found := k.Keeper.GetNextGame(ctx)
	if !found {
		panic("nextGame index failure!")
	}

	newIndex := strconv.FormatUint(nextGame.IdValue, 10)

	newGame := types.New()
	storedGame := types.StoredGame{
		IdValue:      nextGame.IdValue,
		Creator:      msg.Creator,
		Index:        newIndex,
		Game:         newGame.String(),
		CrossPlayer:  msg.Creator,
		CirclePlayer: msg.Creator,
	}
	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}
	k.Keeper.SetStoredGame(ctx, storedGame)

	nextGame.IdValue++
	k.Keeper.SetNextGame(ctx, nextGame)

	return &types.MsgCreateGameResponse{
		IdValue: newIndex,
	}, nil
*/
