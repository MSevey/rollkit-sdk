package staking

import (
	"context"
	"encoding/json"

	"cosmossdk.io/core/appmodule"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/exported"
	"github.com/decentrio/rollkit-sdk/x/staking/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
	_ module.HasServices         = AppModule{}
	_ module.HasInvariants       = AppModule{}
	_ module.HasABCIEndBlock     = AppModule{}
	_ module.HasGenesis          = AppModule{}

	_ appmodule.AppModule       = AppModule{}
	_ appmodule.HasBeginBlocker = AppModule{}
)

type AppModuleBasic struct {
	staking.AppModuleBasic
}

// AppModule embeds the Cosmos SDK's x/staking AppModule where we only override specific methods.
type AppModule struct {
	staking.AppModule
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper, ak types.AccountKeeper, bk types.BankKeeper, ls exported.Subspace) AppModule {
	return AppModule{
		AppModule: staking.NewAppModule(cdc, &keeper.Keeper, ak, bk, ls),
		keeper:    keeper,
	}
}

func (am AppModule) EndBlock(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	_, err := am.keeper.EndBlocker(ctx)
	if err != nil {
		return nil, err
	}

	return []abci.ValidatorUpdate{}, nil
}

// BeginBlock returns the begin blocker for the staking module.
func (am AppModule) BeginBlock(ctx context.Context) error {
	return am.keeper.BeginBlocker(ctx)
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	_ = am.keeper.InitGenesis(ctx, &genesisState)
}
