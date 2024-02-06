package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/polygon/envoy"
)

var _ envoy.QueryServer = queryServer{}

func NewQueryServerImpl(k Keeper) envoy.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

func (qs queryServer) GetLock(ctx context.Context, req *envoy.QueryGetLockRequest) (*envoy.QueryGetLockResponse, error) {
	lock, err := qs.k.Locks.Get(ctx, req.Name)
	if err == nil {
		return &envoy.QueryGetLockResponse{Lock: &lock}, nil
	}
	if errors.Is(err, collections.ErrNotFound) {
		return &envoy.QueryGetLockResponse{Lock: nil}, nil
	}

	return nil, status.Error(codes.Internal, err.Error())
}
