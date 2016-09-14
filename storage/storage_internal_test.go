package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSanitizeShort(t *testing.T) {
	testTable := []struct {
		name string
		in   string
		out  string
		err  error
	}{
		{name: "flatten case",
			in: "ABC", out: "abc"},
		{name: "remove special characters",
			in: "a/b-c_d e", out: "abcde"},
		{name: "test empty string after special character removing",
			in: "  ", err: ErrShortEmpty},
	}

	for _, tt := range testTable {
		result, err := sanitizeShort(tt.in)

		require.Equal(t, tt.err, err)
		require.Equal(t, tt.out, result)
	}
}
