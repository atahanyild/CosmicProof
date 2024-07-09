package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreatecosmicproof = "create_cosmicproof"

var _ sdk.Msg = &MsgCreatecosmicproof{}

type MsgCreatecosmicproof struct {
    Creator      sdk.AccAddress `json:"creator" yaml:"creator"`
    PreviousProof string        `json:"previous_proof" yaml:"previous_proof"`
    BlockHash    string         `json:"block_hash" yaml:"block_hash"`
}

func NewMsgCreatecosmicproof(creator sdk.AccAddress, previousProof, blockHash string) *MsgCreatecosmicproof {
    return &MsgCreatecosmicproof{
        Creator:      creator,
        PreviousProof: previousProof,
        BlockHash:    blockHash,
    }
}

func (msg *MsgCreatecosmicproof) Route() string {
    return RouterKey
}

func (msg *MsgCreatecosmicproof) Type() string {
    return TypeMsgCreatecosmicproof
}

func (msg *MsgCreatecosmicproof) GetSigners() []sdk.AccAddress {
    return []sdk.AccAddress{msg.Creator}
}

func (msg *MsgCreatecosmicproof) GetSignBytes() []byte {
    return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgCreatecosmicproof) ValidateBasic() error {
    if msg.Creator.Empty() {
        return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
    }
    if msg.PreviousProof == "" {
        return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "previous proof can't be empty")
    }
    if msg.BlockHash == "" {
        return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "block hash can't be empty")
    }
    return nil
}
