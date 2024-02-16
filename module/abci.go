package module

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/polygon/envoy"
)

// These hooks are all about the FinalizeBlock phase, how do we get at the proposal stage? or maybe vote extensions?

func (am AppModule) BeginBlock(ctx context.Context) error {

	context := sdk.UnwrapSDKContext(ctx)

	fmt.Printf(".....EXEC_MODE: %d\n", context.ExecMode())
	fmt.Printf("<<<< BeginBlock [%d] >>>>\n", context.BlockHeight())
	fmt.Printf("<<<< leader=[%v] >>>>\n", am.tracker.IsProposer())

	// todo: create a job manager that the leader can start new asyc jobs if necessary

	return nil
}

func (am AppModule) EndBlock(ctx context.Context) error {
	context := sdk.UnwrapSDKContext(ctx)
	fmt.Printf(".....EXEC_MODE: %d\n", context.ExecMode())
	fmt.Printf("<<<< EndBlock [%d] >>>>\n", context.BlockHeight())
	fmt.Printf("<<<< leader=[%v] >>>>\n", am.tracker.IsProposer())
	return nil
}

func (am AppModule) Precommit(ctx context.Context) error {
	// manage currently running async jobs

	am.tracker.Reset() // maybe this should be post commit somehow?

	context := sdk.UnwrapSDKContext(ctx)
	fmt.Printf(".....EXEC_MODE: %d\n", context.ExecMode())
	fmt.Printf("<<<< Precommit [%d] >>>>\n", context.BlockHeight())
	fmt.Printf("<<<< leader=[%v] >>>>\n", am.tracker.IsProposer())

	// Remove expired locks
	err := am.expireLocks(context)
	if err != nil {
		return err
	}

	return nil
}

func (am AppModule) expireLocks(context sdk.Context) error {
	currentBlock := context.BlockHeight()
	err := am.keeper.Locks.Walk(context, nil, func(name string, lock envoy.Lock) (bool, error) {
		expiresAt := int64(lock.AtBlock + lock.NumBlocks)
		if currentBlock >= expiresAt {
			err := am.keeper.Locks.Remove(context, name)
			if err != nil {
				return false, err
			}
			expiryEvent := sdk.NewEvent("expired-lock",
				sdk.NewAttribute("name", name),
				sdk.NewAttribute("envoy", lock.Envoy),
				sdk.NewAttribute("expired_at", fmt.Sprintf("%d", expiresAt)),
			)
			context.EventManager().EmitEvent(expiryEvent)
		}
		return false, nil
	})
	return err
}
