package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DONE
// You can see how they are constructed below:
var (
	ErrNameDoesNotExist = sdkerrors.Register(ModuleName, 1, "name does not exist")
)
