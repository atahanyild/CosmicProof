package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateZKProof = "create_zkproof"

var _ sdk.Msg = &MsgCreateZKProof{}

type MsgCreateZKProof struct {
    Creator      sdk.AccAddress `json:"creator" yaml:"creator"`
    PreviousProof string        `json:"previous_proof" yaml:"previous_proof"`
    BlockHash    string         `json:"block_hash" yaml:"block_hash"`
}

func NewMsgCreateZKProof(creator sdk.AccAddress, previousProof, blockHash string) *MsgCreateZKProof {
    return &MsgCreateZKProof{
        Creator:      creator,
        PreviousProof: previousProof,
        BlockHash:    blockHash,
    }
}

func (msg *MsgCreateZKProof) Route() string {
    return RouterKey
}

func (msg *MsgCreateZKProof) Type() string {
    return TypeMsgCreateZKProof
}

func (msg *MsgCreateZKProof) GetSigners() []sdk.AccAddress {
    return []sdk.AccAddress{msg.Creator}
}

func (msg *MsgCreateZKProof) GetSignBytes() []byte {
    return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgCreateZKProof) ValidateBasic() error {
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
