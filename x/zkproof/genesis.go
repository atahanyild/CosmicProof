package zkproof

import (
    "encoding/json"

    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/atahanyild/ytubc-lambda/x/zkproof/types"
)

func DefaultGenesisState() types.GenesisState {
    return types.GenesisState{}
}

func ValidateGenesis(data types.GenesisState) error {
    // Validate your genesis state here
    return nil
}
