package cosmicproof

import (
    "encoding/json"
    "github.com/cosmos/cosmos-sdk/codec"
    "github.com/cosmos/cosmos-sdk/types/module"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/x/params"
    "github.com/atahanyild/ytubc-lambda/x/cosmicproof/keeper"
    "github.com/atahanyild/ytubc-lambda/x/cosmicproof/types"
)

var (
    _ module.AppModule      = AppModule{}
    _ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
    return types.ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.LegacyAmino) {
    types.RegisterCodec(cdc)
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
    return cdc.MustMarshalJSON(types.DefaultGenesis())
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
    var genesisState types.GenesisState
    if err := cdc.UnmarshalJSON(bz, &genesisState); err != nil {
        return err
    }
    return types.ValidateGenesis(genesisState)
}

type AppModule struct {
    AppModuleBasic
    Keeper keeper.Keeper
}

func NewAppModule(k keeper.Keeper) AppModule {
    return AppModule{
        AppModuleBasic: AppModuleBasic{},
        Keeper:         k,
    }
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) Route() sdk.Route {
    return sdk.NewRoute(types.RouterKey, NewHandler(am.Keeper))
}

func (am AppModule) QuerierRoute() string {
    return types.QuerierRoute
}

func (am AppModule) LegacyQuerierHandler(cdc *codec.LegacyAmino) sdk.Querier {
    return nil
}

func (am AppModule) RegisterServices(cfg module.Configurator) {}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
    var genesisState types.GenesisState
    cdc.MustUnmarshalJSON(data, &genesisState)
    return am.Keeper.InitGenesis(ctx, &genesisState)
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
    gs := am.Keeper.ExportGenesis(ctx)
    return cdc.MustMarshalJSON(gs)
}
