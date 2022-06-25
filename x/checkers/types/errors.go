package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidCreator      = sdkerrors.Register(ModuleName, 1100, "creator address is invalid: %s")
	ErrGameNotParseable    = sdkerrors.Register(ModuleName, 1101, "game cannot be parsed")
	ErrGameNotFound        = sdkerrors.Register(ModuleName, 1102, "game by id not found: %s")
	ErrCreatorNotPlayer    = sdkerrors.Register(ModuleName, 1103, "message creator is not a player: %s")
	ErrTurnAddressNotValid = sdkerrors.Register(ModuleName, 1104, "message turn is not a player: %s")
	ErrNotPlayerTurn       = sdkerrors.Register(ModuleName, 1105, "player tried to play out of turn: %s")
	ErrWrongMove           = sdkerrors.Register(ModuleName, 1106, "wrong move")
)
