package keeper

import (
    "os/exec"
    "encoding/json"
    "github.com/cosmos/cosmos-sdk/store/prefix"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/codec"
    "github.com/atahanyild/ytubc-lambda/x/cosmicproof/types"
)

//Keeper access to the subset of the state defined by various modules
type Keeper struct {
    storeKey   sdk.StoreKey
    cdc        codec.Codec
}

// Creating new keeper to use
func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey) Keeper {
    return Keeper{
        storeKey:   storeKey,
        cdc:        cdc,
    }
}
 
/* 
   This function executes the js code which is written with "o1.js",it creates zeroknowledge proof by using latest 
   proof and current blockhash as inputs then returns a zeroknowledge 
   proof for this block  
*/
func (k Keeper) Generatecosmicproof(ctx sdk.Context, previousProof string, blockHash string) (string, error) {
    cmd := exec.Command("node", "path_to_generate_proof.js", previousProof, blockHash)
    output, err := cmd.Output()
    if err != nil {
        return "", err
    }
    var proof types.cosmicproof
    err = json.Unmarshal(output, &proof)
    if err != nil {
        return "", err
    }
    return proof.Proof, nil
}


/* 
   This function executes the js code which is written with "o1.js",it verifys zeroknowledge proof by using  
   proof and genesis blockhash as inputs then returns a boolean 
   value for given proof  
*/
func (k Keeper) Verifycosmicproof(ctx sdk.Context, zeroknowledgeProof string, genesisHash string) (bool, error) {
    cmd := exec.Command("node", "path_to_verify_proof.js",zeroknowledgeProof, genesisHash)
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


/* 
   This function saves a proof associated with a specific block height 
   into a prefixed key-value store 
   within the Cosmos SDK's context.
*/
func (k Keeper) Savecosmicproof(ctx sdk.Context, blockHeight int64, proof string) error {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("cosmicproof"))
    key := []byte(string(blockHeight))
    store.Set(key, []byte(proof))
    return nil
}

/* 
   This function gets the proof of specific block height.
*/
func (k Keeper) Getcosmicproof(ctx sdk.Context, blockHeight int64) (string, error) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("cosmicproof"))
    key := []byte(string(blockHeight))
    proof := store.Get(key)
    if proof == nil {
        return "", sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "proof for block height %d not found", blockHeight)
    }
    return string(proof), nil
}
