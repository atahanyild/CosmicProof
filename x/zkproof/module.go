package zkproof

import (
    "github.com/cosmos/cosmos-sdk/types/module"
    "github.com/cosmos/cosmos-sdk/codec"
    "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/x/params/types"
    "github.com/atahanyild/ytubc-lambda/x/zkproof/keeper"
    "github.com/atahanyild/ytubc-lambda/x/zkproof/types"
)

type AppModuleBasic struct{}

func (AppModuleBasic) RegisterCodec(cdc *codec.LegacyAmino) {
    // Register your types here
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
    // Default genesis state
    return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config types.TxEncodingConfig, bz json.RawMessage) error {
    // Validate genesis state
    var genesisState types.GenesisState
    if err := cdc.UnmarshalJSON(bz, &genesisState); err != nil {
        return err
    }
    return types.ValidateGenesis(genesisState)
}

type AppModule struct {
    AppModuleBasic
    keeper keeper.Keeper
}

func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
    return AppModule{
        AppModuleBasic: AppModuleBasic{},
        keeper:         keeper,
    }
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) Route() sdk.Route {
    return sdk.NewRoute(types.RouterKey, NewHandler(am.keeper))
}

func (am AppModule) QuerierRoute() string { return types.QuerierRoute }

func (am AppModule) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier {
    return nil
}

func (am AppModule) RegisterServices(cfg module.Configurator) {}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
    // Initialize genesis state
    var genesisState types.GenesisState
    cdc.MustUnmarshalJSON(data, &genesisState)
    return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
    // Export genesis state
    return nil
}

func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
    // Call BeginBlocker here
    BeginBlocker(ctx, req, am.keeper)
}

func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
    // Call EndBlocker here
    return nil
}
