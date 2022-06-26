package keeper

import (
	"github.com/alice/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetWaitingGame set a specific waitingGame in the store from its index
func (k Keeper) SetWaitingGame(ctx sdk.Context, waitingGame types.WaitingGame) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WaitingGameKeyPrefix))
	b := k.cdc.MustMarshal(&waitingGame)
	store.Set(types.WaitingGameKey(
		waitingGame.Index,
	), b)
}

// GetWaitingGame returns a waitingGame from its index
func (k Keeper) GetWaitingGame(
	ctx sdk.Context,
	index string,

) (val types.WaitingGame, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WaitingGameKeyPrefix))

	b := store.Get(types.WaitingGameKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWaitingGame removes a waitingGame from the store
func (k Keeper) RemoveWaitingGame(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WaitingGameKeyPrefix))
	store.Delete(types.WaitingGameKey(
		index,
	))
}

// GetAllWaitingGame returns all waitingGame
func (k Keeper) GetAllWaitingGame(ctx sdk.Context) (list []types.WaitingGame) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WaitingGameKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.WaitingGame
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
