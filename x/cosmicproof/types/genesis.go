package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type GenesisState struct {
    cosmicproofs []cosmicproof `json:"cosmicproofs"`
}

func DefaultGenesis() *GenesisState {
    return &GenesisState{
        cosmicproofs: []cosmicproof{},
    }
}

func ValidateGenesis(data GenesisState) error {
    for _, record := range data.cosmicproofs {
        if record.Proof == "" {
            return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proof cannot be empty")
        }
    }
    return nil
}
