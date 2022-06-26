package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMarkSpace = "mark_space"

var _ sdk.Msg = &MsgMarkSpace{}

func NewMsgMarkSpace(creator string, idValue uint64, x uint64, y uint64) *MsgMarkSpace {
	return &MsgMarkSpace{
		Creator: creator,
		IdValue: idValue,
		X:       x,
		Y:       y,
	}
}

func (msg *MsgMarkSpace) Route() string {
	return RouterKey
}

func (msg *MsgMarkSpace) Type() string {
	return TypeMsgMarkSpace
}

func (msg *MsgMarkSpace) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMarkSpace) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMarkSpace) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
