package envoy

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//"github.com/polygon/envoy/rules"
)

func (lock Lock) GetEnvoyAddress() (sdk.AccAddress, error) {
	envoy, err := sdk.AccAddressFromBech32(lock.GetEnvoy())
	return envoy, errors.Wrapf(err, ErrInvalidAddress.Error(), lock.Envoy)
}

func (lock Lock) IsExpired(currentBlock uint32) error {
	if lock.AtBlock+lock.NumBlocks > currentBlock {
		return ErrExpiredLock
	}
	return nil
}

func (lock Lock) IsOwner(addr sdk.AccAddress) (bool, error) {
	envoy, err := lock.GetEnvoyAddress()
	if err != nil {
		return false, err
	}
	return envoy.Equals(addr), nil
}

func (lock Lock) Validate() error { // can this take argument??
	_, err := lock.GetEnvoyAddress()
	if err != nil {
		return err
	}
	/*
		err = lock.CheckExpiry(currentBlock)
		if err != nil {
			return err
		}
	*/
	return nil
}
