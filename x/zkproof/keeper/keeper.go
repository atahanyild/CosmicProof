package keeper

import (
    "os/exec"
    "encoding/json"
    "github.com/cosmos/cosmos-sdk/store/prefix"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/codec"
    "github.com/atahanyild/ytubc-lambda/x/zkproof/types"
)

type Keeper struct {
    storeKey   sdk.StoreKey
    cdc        codec.Codec
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey) Keeper {
    return Keeper{
        storeKey:   storeKey,
        cdc:        cdc,
    }
}

func (k Keeper) GenerateZKProof(ctx sdk.Context, previousProof string, blockHash string) (string, error) {
    cmd := exec.Command("node", "path_to_generate_proof.js", previousProof, blockHash)
    output, err := cmd.Output()
    if err != nil {
        return "", err
    }
    var proof types.ZKProof
    err = json.Unmarshal(output, &proof)
    if err != nil {
        return "", err
    }
    return proof.Proof, nil
}

func (k Keeper) VerifyZKProof(ctx sdk.Context, genesisHash string) (bool, error) {
    cmd := exec.Command("node", "path_to_verify_proof.js", genesisHash)
    output, err := cmd.Output()
    if err != nil {
        return false, err
    }
    var result bool
    err = json.Unmarshal(output, &result)
    if err != nil {
        return false, err
    }
    return result, nil
}

func (k Keeper) SaveZKProof(ctx sdk.Context, blockHeight int64, proof string) error {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("zkproof"))
    key := []byte(string(blockHeight))
    store.Set(key, []byte(proof))
    return nil
}

func (k Keeper) GetZKProof(ctx sdk.Context, blockHeight int64) (string, error) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("zkproof"))
    key := []byte(string(blockHeight))
    proof := store.Get(key)
    if proof == nil {
        return "", sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "proof for block height %d not found", blockHeight)
    }
    return string(proof), nil
}
