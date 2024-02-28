package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDexSwap{}

func NewMsgDexSwap(creator string, assetOne string, assetTwo string, amountIn uint64, amountOut uint64) *MsgDexSwap {
	return &MsgDexSwap{
		Creator:   creator,
		AssetOne:  assetOne,
		AssetTwo:  assetTwo,
		AmountIn:  amountIn,
		AmountOut: amountOut,
	}
}

func (msg *MsgDexSwap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
