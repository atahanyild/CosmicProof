package keeper_test

import (
    "testing"
    "github.com/stretchr/testify/require"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/store"
    "github.com/cosmos/cosmos-sdk/store/prefix"
    "github.com/cosmos/cosmos-sdk/codec"
    storetypes "github.com/cosmos/cosmos-sdk/store/types"
    paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
    "github.com/cosmos/cosmos-sdk/x/params"
    "github.com/cosmos/cosmos-sdk/x/params/subspace"
    "github.com/cosmos/cosmos-sdk/simapp"
    "github.com/cosmos/cosmos-sdk/codec/types"
    "github.com/cosmos/cosmos-sdk/testutil"
    "github.com/cosmos/cosmos-sdk/x/auth/keeper"
    "github.com/cosmos/cosmos-sdk/x/auth/types"

    "myapp/x/zkproof/keeper"
)

func TestKeeper_SetGetZKProof(t *testing.T) {
    // Initialize the store keys and memory database
    key := sdk.NewKVStoreKey("zkproof")
    tkey := sdk.NewTransientStoreKey("transient_test")
    memDB := store.NewCommitMultiStore(testutil.NewMemoryStore())

    memDB.MountStoreWithDB(key, storetypes.StoreTypeIAVL, nil)
    memDB.MountStoreWithDB(tkey, storetypes.StoreTypeTransient, nil)
    err := memDB.LoadLatestVersion()
    require.NoError(t, err)

    // Create a new codec
    cdc := codec.NewProtoCodec(types.NewInterfaceRegistry())

    // Create a parameter subspace
    paramSpace := paramstypes.NewSubspace(cdc, key, tkey, "zkproofparams")

    // Create a new keeper
    zkKeeper := keeper.NewKeeper(cdc, key, paramSpace)

    // Create a new context
    ctx := sdk.NewContext(memDB, storetypes.CommitID{}, false, testutil.Logger(t))

    // Test data
    blockHeight := int64(100)
    zkProof := []byte("sampleZKProof")

    // Test SetZKProof
    zkKeeper.SetZKProof(ctx, blockHeight, zkProof)

    // Test GetZKProof
    retrievedProof, err := zkKeeper.GetZKProof(ctx, blockHeight)
    require.NoError(t, err)
    require.NotNil(t, retrievedProof)
    require.Equal(t, zkProof, retrievedProof)
}

