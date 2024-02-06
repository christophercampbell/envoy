package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	"github.com/polygon/envoy"
)

type msgServer struct {
	k Keeper
}

var _ envoy.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) envoy.MsgServer {
	return &msgServer{k: keeper}
}

func (ms msgServer) CreateLock(ctx context.Context, msg *envoy.MsgCreateLock) (*envoy.MsgCreateLockResponse, error) {
	if length := len([]byte(msg.Name)); envoy.MaxLockNameLength < length || length < 1 {
		return nil, envoy.ErrNameTooLong
	}
	if _, err := ms.k.Locks.Get(ctx, msg.Name); err == nil || errors.Is(err, collections.ErrEncoding) {
		return nil, fmt.Errorf("lock already exists with name: %s", msg.Name)
	}
	blockTimeout := msg.NumBlocks
	if blockTimeout == 0 {
		blockTimeout = envoy.DefaultBlockTimeout
	}

	storedLock := envoy.Lock{
		Name:      msg.Name,
		Envoy:     msg.Envoy,   // in future this gets decided by consensus, or maybe just always the leader
		AtBlock:   msg.AtBlock, // this needs to come from chain: "latest"
		NumBlocks: blockTimeout,
	}
	if err := storedLock.Validate(); err != nil {
		return nil, err
	}
	if err := ms.k.Locks.Set(ctx, msg.Name, storedLock); err != nil {
		return nil, err
	}
	return &envoy.MsgCreateLockResponse{}, nil
}
