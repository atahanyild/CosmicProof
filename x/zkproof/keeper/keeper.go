package keeper

import (
    "fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
    "cosmossdk.io/store/prefix"
    "github.com/cosmos/cosmos-sdk/codec"
    paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Keeper struct {
    storeKey   sdk.StoreKey
    cdc        codec.BinaryCodec
    paramSpace paramtypes.Subspace
}

func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, paramSpace paramtypes.Subspace) Keeper {
    return Keeper{
        storeKey:   storeKey,
        cdc:        cdc,
        paramSpace: paramSpace,
    }
}

func (k Keeper) SetZKProof(ctx sdk.Context, blockHeight int64, zkProof []byte) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("zkproof"))
    key := []byte(fmt.Sprintf("%d", blockHeight))
    store.Set(key, zkProof)
}

func (k Keeper) GetZKProof(ctx sdk.Context, blockHeight int64) []byte {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("zkproof"))
    key := []byte(fmt.Sprintf("%d", blockHeight))
    return store.Get(key)
}
