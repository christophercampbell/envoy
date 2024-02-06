package envoy

import "cosmossdk.io/collections"

const ModuleName = "envoy"
const MaxLockNameLength = 256
const DefaultBlockTimeout = 8

var (
	ParamsKey = collections.NewPrefix("Params")
	LocksKey  = collections.NewPrefix("Locks/value/")
)
