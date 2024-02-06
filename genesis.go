package envoy

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
func (gs *GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	unique := make(map[string]bool)
	for _, lock := range gs.Locks {
		if length := len([]byte(lock.Name)); MaxLockNameLength < length || length < 1 {
			return ErrNameTooLong
		}
		if _, ok := unique[lock.Name]; ok {
			return ErrDuplicateLockName
		}
		if err := lock.Validate(); err != nil {
			return err
		}
		unique[lock.Name] = true
	}

	return nil
}
