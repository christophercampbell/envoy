package module

import (
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"

	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	envoyv1 "github.com/polygon/envoy/api/module/v1"
	"github.com/polygon/envoy/keeper"
)

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

func init() {
	appmodule.Register(
		&envoyv1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Cdc          codec.Codec
	StoreService store.KVStoreService
	AddressCodec address.Codec

	Config *envoyv1.Module
}

type ModuleOutputs struct {
	depinject.Out

	Module  appmodule.AppModule
	Keeper  keeper.Keeper
	Tracker *Tracker
}

func ProvideModule(in ModuleInputs) ModuleOutputs {

	// default to governance as authority if not provided
	authority := authtypes.NewModuleAddress("gov")
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	k := keeper.NewKeeper(in.Cdc, in.AddressCodec, in.StoreService, authority.String())
	t := Tracker{}
	m := NewAppModule(in.Cdc, k, &t)

	return ModuleOutputs{Module: m, Keeper: k, Tracker: &t}
}
