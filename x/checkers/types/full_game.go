package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (storedGame *StoredGame) GetCreatorAddress() (creator sdk.AccAddress, err error) {
	creator, errCreator := sdk.AccAddressFromBech32(storedGame.Creator)
	return creator, sdkerrors.Wrapf(errCreator, ErrInvalidCreator.Error(), storedGame.Creator)
}

func (storedGame *StoredGame) GetCirclePlayerAddress() (red sdk.AccAddress, err error) {
	red, errRed := sdk.AccAddressFromBech32(storedGame.CirclePlayer)
	return red, sdkerrors.Wrapf(errRed, ErrTurnAddressNotValid.Error(), storedGame.CirclePlayer)
}

func (storedGame *StoredGame) GetCrossPlayerAddress() (red sdk.AccAddress, err error) {
	red, errRed := sdk.AccAddressFromBech32(storedGame.CrossPlayer)
	return red, sdkerrors.Wrapf(errRed, ErrTurnAddressNotValid.Error(), storedGame.CrossPlayer)
}

func (storedGame *StoredGame) ParseGame() (game *Game, err error) {
	game, errGame := Parse(storedGame.Game)
	if errGame != nil {
		return nil, sdkerrors.Wrapf(errGame, ErrGameNotParseable.Error())
	}
	return game, nil
}

func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetCreatorAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.ParseGame()
	if err != nil {
		return err
	}
	_, err = storedGame.GetCirclePlayerAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.GetCrossPlayerAddress()
	return err
}
