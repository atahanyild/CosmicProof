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

    cosmicproof "github.com/atahanyild/ytubc-lambda/x/cosmicproof"
    cosmicproofkeeper "github.com/atahanyild/ytubc-lambda/x/cosmicproof/keeper"
    cosmicprooftypes "github.com/atahanyild/ytubc-lambda/x/cosmicproof/types"
)

type App struct {
    *baseapp.BaseApp
    cdc *codec.Codec

    // Keepers
    AccountKeeper auth.AccountKeeper
    BankKeeper    bank.Keeper
    StakingKeeper staking.Keeper
    SupplyKeeper  supply.Keeper
    cosmicproofKeeper cosmicproofkeeper.Keeper

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
        cosmicproofKeeper: cosmicproofkeeper.NewKeeper(codec, keys[cosmicprooftypes.StoreKey]),
    }

    // Register module routes and query routes
    app.Router().AddRoute(cosmicprooftypes.RouterKey, cosmicproof.NewHandler(app.cosmicproofKeeper))
    app.QueryRouter().AddRoute(cosmicprooftypes.QuerierRoute, cosmicproof.NewQuerier(app.cosmicproofKeeper))

    // Module Manager
    app.mm = module.NewManager(
        auth.NewAppModule(app.AccountKeeper),
        bank.NewAppModule(app.BankKeeper, app.AccountKeeper),
        staking.NewAppModule(app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
        cosmicproof.NewAppModule(app.cosmicproofKeeper),
    )

    // Set the order of begin blockers
    app.mm.SetOrderBeginBlockers(
        cosmicprooftypes.ModuleName,
    )

    // Register modules
    app.mm.RegisterInvariants(&app.CrisisKeeper)
    app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), codec)
    
    // Initialize stores
    app.MountStores(keys[auth.StoreKey], keys[bank.StoreKey], keys[staking.StoreKey], keys[supply.StoreKey], keys[cosmicprooftypes.StoreKey])

    return app
}
