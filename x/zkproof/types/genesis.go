package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type GenesisState struct {
    ZKProofs []ZKProof `json:"zkproofs"`
}

func DefaultGenesis() *GenesisState {
    return &GenesisState{
        ZKProofs: []ZKProof{},
    }
}

func ValidateGenesis(data GenesisState) error {
    for _, record := range data.ZKProofs {
        if record.Proof == "" {
            return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proof cannot be empty")
        }
    }
    return nil
}
