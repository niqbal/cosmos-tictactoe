package keeper

import (
	"context"
	"strconv"

	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MarkSpace(goCtx context.Context, msg *types.MsgMarkSpace) (*types.MsgMarkSpaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	index := strconv.FormatUint(uint64(msg.IdValue), 10)
	storedGame, storedGameFound := k.Keeper.GetStoredGame(ctx, index)

	if !storedGameFound {
		panic("Game not found!")
	}

	parsedGame, err := types.Parse(storedGame.Game)
	if err != nil {
		panic("Game is in invalid state")
	}

	playerType := types.EMPTY_MARK
	if msg.Creator == storedGame.CrossPlayer {
		playerType = types.CROSS_MARK
	} else if msg.Creator == storedGame.CirclePlayer {
		playerType = types.CIRCLE_MARK
	} else {
		panic("You are not a player in this game!")
	}

	if parsedGame.GetTurn() != playerType {
		panic("It is not your turn!")
	}

	err = parsedGame.MarkSpace(types.Pos{X: msg.X, Y: msg.Y}, playerType)
	if err != nil {
		panic(err)
	}

	newStoredGame := types.StoredGame{
		IdValue:      msg.IdValue,
		Creator:      storedGame.Creator,
		Index:        storedGame.Index,
		Game:         parsedGame.String(),
		CrossPlayer:  storedGame.CrossPlayer,
		CirclePlayer: storedGame.CirclePlayer,
	}

	err = newStoredGame.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.SetStoredGame(ctx, newStoredGame)

	return &types.MsgMarkSpaceResponse{
		Success: true,
	}, nil
}
