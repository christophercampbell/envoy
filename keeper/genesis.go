package keeper

import (
	"context"

	"github.com/polygon/envoy"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *envoy.GenesisState) error {
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}
	for _, lock := range data.Locks {
		if err := k.Locks.Set(ctx, lock.Name, lock); err != nil {
			return err
		}
	}
	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*envoy.GenesisState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	var locks []envoy.Lock
	err = k.Locks.Walk(ctx, nil, func(index string, lock envoy.Lock) (bool, error) {
		locks = append(locks, lock)
		return false, nil
	})
	if err != nil {
		return nil, err
	}
	return &envoy.GenesisState{
		Params: params,
		Locks:  locks,
	}, nil
}
