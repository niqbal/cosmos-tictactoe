package keeper

import (
	"context"
	"crypto/sha1"
	"strconv"

	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) JoinGame(goCtx context.Context, msg *types.MsgJoinGame) (*types.MsgJoinGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	index := strconv.FormatUint(uint64(msg.IdValue), 10)
	waitingGame, found := k.Keeper.GetWaitingGame(ctx, index)

	if !found {
		panic("waitingGame index failure!")
	}

	_, storedGameFound := k.Keeper.GetStoredGame(ctx, index)

	if storedGameFound {
		panic("This game is already full!")
	}

	newGame := types.New()

	crossPlayer, circlePlayer := assignPlayerTurns(waitingGame, msg)

	storedGame := types.StoredGame{
		IdValue:      msg.IdValue,
		Creator:      waitingGame.Creator,
		Index:        waitingGame.Index,
		Game:         newGame.String(),
		CrossPlayer:  crossPlayer,
		CirclePlayer: circlePlayer,
	}

	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}
	k.Keeper.SetStoredGame(ctx, storedGame)

	return &types.MsgJoinGameResponse{
		Success: true,
	}, nil
}

func assignPlayerTurns(waitingGame types.WaitingGame, msg *types.MsgJoinGame) (string, string) {
	h := sha1.New()
	h.Write([]byte(waitingGame.Creator + msg.Creator))
	hashValue := h.Sum(nil)

	crossPlayer := waitingGame.Creator
	circlePlayer := msg.Creator

	if hashValue[len(hashValue)-1]%2 == 1 {
		crossPlayer = msg.Creator
		circlePlayer = waitingGame.Creator
	}
	return crossPlayer, circlePlayer
}
