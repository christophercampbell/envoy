package module

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keep track of when the node is the leader
type Tracker struct {
	isProposer bool
}

func (rs *Tracker) setIsProposer() {
	rs.isProposer = true
}

func (rs *Tracker) IsProposer() bool {
	return rs.isProposer
}

func (rs *Tracker) Reset() {
	rs.isProposer = false
}

// Manual setting of this is required in the cosmos app
func (rs *Tracker) PrepareProposal(_ sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
	rs.isProposer = true
	return &abci.ResponsePrepareProposal{Txs: req.Txs}, nil
}
