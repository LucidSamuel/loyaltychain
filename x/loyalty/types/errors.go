package types

import (
	sdkerrors "cosmossdk.io/errors"
)

// Error definitions
var (
	ErrRewardNotFound     = sdkerrors.Register(ModuleName, 1100, "reward not found")
	ErrInsufficientPoints = sdkerrors.Register(ModuleName, 1101, "insufficient points")
	ErrInvalidSigner      = sdkerrors.Register(ModuleName, 1102, "invalid signer")
	ErrPointNotFound      = sdkerrors.Register(ModuleName, 1103, "point not found")
	ErrInvalidRequest     = sdkerrors.Register(ModuleName, 1104, "invalid request")
	ErrInvalidAddress     = sdkerrors.Register(ModuleName, 1105, "invalid address")
	ErrKeyNotFound        = sdkerrors.Register(ModuleName, 1106, "key not found")
	ErrUnauthorized       = sdkerrors.Register(ModuleName, 1107, "unauthorized")
)

