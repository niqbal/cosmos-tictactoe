package keeper

import (
	"context"

	"github.com/alice/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WaitingGameAll(c context.Context, req *types.QueryAllWaitingGameRequest) (*types.QueryAllWaitingGameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var waitingGames []types.WaitingGame
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	waitingGameStore := prefix.NewStore(store, types.KeyPrefix(types.WaitingGameKeyPrefix))

	pageRes, err := query.Paginate(waitingGameStore, req.Pagination, func(key []byte, value []byte) error {
		var waitingGame types.WaitingGame
		if err := k.cdc.Unmarshal(value, &waitingGame); err != nil {
			return err
		}

		_, found := k.GetStoredGame(ctx, waitingGame.Index)
		if !found {
			waitingGames = append(waitingGames, waitingGame)
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWaitingGameResponse{WaitingGame: waitingGames, Pagination: pageRes}, nil
}

func (k Keeper) WaitingGame(c context.Context, req *types.QueryGetWaitingGameRequest) (*types.QueryGetWaitingGameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetWaitingGame(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetWaitingGameResponse{WaitingGame: val}, nil
}
