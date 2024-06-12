package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePoint{}

func NewMsgCreatePoint(
	creator string,
	index string,
	owner string,
	balance int32,

) *MsgCreatePoint {
	return &MsgCreatePoint{
		Creator: creator,
		Index:   index,
		Owner:   owner,
		Balance: balance,
	}
}

func (msg *MsgCreatePoint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePoint{}

func NewMsgUpdatePoint(
	creator string,
	index string,
	owner string,
	balance int32,

) *MsgUpdatePoint {
	return &MsgUpdatePoint{
		Creator: creator,
		Index:   index,
		Owner:   owner,
		Balance: balance,
	}
}

func (msg *MsgUpdatePoint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePoint{}

func NewMsgDeletePoint(
	creator string,
	index string,

) *MsgDeletePoint {
	return &MsgDeletePoint{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeletePoint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
