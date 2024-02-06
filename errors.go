package envoy

import "cosmossdk.io/errors"

var (
	ErrNameTooLong       = errors.Register(ModuleName, 2, "name too long")
	ErrDuplicateLockName = errors.Register(ModuleName, 3, "duplicate lock name")
	ErrInvalidAddress    = errors.Register(ModuleName, 4, "invalid address")
	ErrExpiredLock       = errors.Register(ModuleName, 5, "expired lock")
)
