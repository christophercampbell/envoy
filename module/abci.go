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

	// selfAddr := "" //??? sdk.ValAddress()

	var zeroVal sdk.ValAddress = context.VoteInfos()[0].GetValidator().Address

	var proposerAddr sdk.ValAddress = context.CometInfo().GetProposerAddress()

	fmt.Printf(".....authority: %s\n", am.keeper.GetAuthority())
	fmt.Printf(">>>>>>>>> proposer: %s, other: %s\n", proposerAddr.String(), zeroVal.String())

	return nil
}

func (am AppModule) EndBlock(ctx context.Context) error {
	context := sdk.UnwrapSDKContext(ctx)
	fmt.Printf(".....EXEC_MODE: %d\n", context.ExecMode())
	fmt.Printf("<<<< EndBlock [%d] >>>>\n", context.BlockHeight())
	return nil
}

func (am AppModule) Precommit(ctx context.Context) error {
	context := sdk.UnwrapSDKContext(ctx)
	fmt.Printf(".....EXEC_MODE: %d\n", context.ExecMode())
	fmt.Printf("<<<< Precommit [%d] >>>>\n", context.BlockHeight())
	// Remove expired locks
	currentBlock := sdk.UnwrapSDKContext(ctx).BlockHeight()
	err := am.keeper.Locks.Walk(ctx, nil, func(name string, lock envoy.Lock) (bool, error) {
		expiresAt := int64(lock.AtBlock + lock.NumBlocks)
		if currentBlock >= expiresAt {
			err := am.keeper.Locks.Remove(ctx, name)
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

	if err != nil {
		return err
	}

	return nil
}
