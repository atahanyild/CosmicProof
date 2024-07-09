package app

import (
    "github.com/cosmos/cosmos-sdk/baseapp"
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/module"
    "github.com/cosmos/cosmos-sdk/x/auth"
    "github.com/cosmos/cosmos-sdk/x/bank"
    "github.com/cosmos/cosmos-sdk/x/staking"
    "github.com/cosmos/cosmos-sdk/x/supply"

    zkproof "github.com/atahanyild/ytubc-lambda/x/zkproof"
    zkproofkeeper "github.com/atahanyild/ytubc-lambda/x/zkproof/keeper"
    zkprooftypes "github.com/atahanyild/ytubc-lambda/x/zkproof/types"
)

type App struct {
    *baseapp.BaseApp
    cdc *codec.Codec

    // Keepers
    AccountKeeper auth.AccountKeeper
    BankKeeper    bank.Keeper
    StakingKeeper staking.Keeper
    SupplyKeeper  supply.Keeper
    ZKProofKeeper zkproofkeeper.Keeper

    // Module Manager
    mm *module.Manager
}

func NewApp(...) *App {
    // Initialize your application
    app := &App{
        // Initialize base app and codec
        BaseApp: baseapp.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(codec)),
        cdc:     codec,

        // Initialize keepers
        AccountKeeper: auth.NewAccountKeeper(codec, keys[auth.StoreKey], auth.ProtoBaseAccount),
        BankKeeper:    bank.NewBaseKeeper(codec, keys[bank.StoreKey], app.AccountKeeper),
        StakingKeeper: staking.NewKeeper(codec, keys[staking.StoreKey], app.BankKeeper, app.AccountKeeper, staking.DefaultCodespace),
        SupplyKeeper:  supply.NewKeeper(codec, keys[supply.StoreKey], app.AccountKeeper, app.BankKeeper, app.StakingKeeper, supply.DefaultCodespace),
        ZKProofKeeper: zkproofkeeper.NewKeeper(codec, keys[zkprooftypes.StoreKey]),
    }

    // Register module routes and query routes
    app.Router().AddRoute(zkprooftypes.RouterKey, zkproof.NewHandler(app.ZKProofKeeper))
    app.QueryRouter().AddRoute(zkprooftypes.QuerierRoute, zkproof.NewQuerier(app.ZKProofKeeper))

    // Module Manager
    app.mm = module.NewManager(
        auth.NewAppModule(app.AccountKeeper),
        bank.NewAppModule(app.BankKeeper, app.AccountKeeper),
        staking.NewAppModule(app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
        zkproof.NewAppModule(app.ZKProofKeeper),
    )

    // Set the order of begin blockers
    app.mm.SetOrderBeginBlockers(
        zkprooftypes.ModuleName,
    )

    // Register modules
    app.mm.RegisterInvariants(&app.CrisisKeeper)
    app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), codec)
    
    // Initialize stores
    app.MountStores(keys[auth.StoreKey], keys[bank.StoreKey], keys[staking.StoreKey], keys[supply.StoreKey], keys[zkprooftypes.StoreKey])

    return app
}
