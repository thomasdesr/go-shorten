package multistorage

// MultiStorageOptions allows you to to configure out the MultiStorage will behave. For example should it Save changes to all underlying packages, or just the first one.
type MultiStorageOption func(*MultiStorage) error

// LoadFirst causes the Multistore it is configuring to return on the first store that doesn't return ErrShortNotSet
func LoadFirst() MultiStorageOption {
	return func(m *MultiStorage) error {
		m.loader = loadFirstFunc
		return nil
	}
}

// LoadCompareAll causes the MultiStorage to try to load the short from all of the underlying stores and then compares them all for equality before returning. If they are not all equal it will return an error
func LoadCompareAllResults() MultiStorageOption {
	return func(m *MultiStorage) error {
		m.loader = loadCompareAllResultsFunc
		return nil
	}
}
