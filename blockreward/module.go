package blockreward

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/yourchain/evmos/x/blockreward/keeper"
)

type AppModule struct {
	keeper keeper.Keeper
}

func NewAppModule(k keeper.Keeper) AppModule { return AppModule{keeper: k} }

func (am AppModule) Name() string { return "blockreward" }

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) Route() sdk.Route { return sdk.Route{} }
func (am AppModule) QuerierRoute() string { return "" }
func (am AppModule) LegacyQuerierHandler(*module.LegacyQuerierCdc) sdk.Querier { return nil }

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {}
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage     { return nil }
func (am AppModule) RegisterServices(cfg module.Configurator)                               {}

func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	_ = am.keeper.MintAndDistribute(ctx)
}

func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
