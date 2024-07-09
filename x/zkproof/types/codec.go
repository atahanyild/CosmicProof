package types

import (
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    cdc.RegisterConcrete(MsgCreateZKProof{}, "zkproof/CreateZKProof", nil)
}

func RegisterInterfaces(registry sdk.InterfaceRegistry) {
    registry.RegisterImplementations(
        (*sdk.Msg)(nil),
        &MsgCreateZKProof{},
    )
}
