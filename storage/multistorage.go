package storage

// A storage that will query multiple storages that will only return an ErrShortNotSet if the short code doens't exist in all of them
type MultiStorage []Storage

func (s *MultiStorage) Load(rawShort string) (string, error) {
	short, err := sanitizeShort(rawShort)
	if err != nil {
		return "", err
	}

	for _, store := range *s {
		s, err := store.Load(short)
		if err == ErrShortNotSet {
			continue
		}

		return s, err
	}

	return "", ErrShortNotSet
}
