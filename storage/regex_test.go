package storage

import (
	"testing"

	"github.com/pkg/errors"
)

func TestRegexLoad(t *testing.T) {
	r, err := NewRegexFromList(map[string]string{
		`jira/(.+)`:       "https://atlassian.net/browse/$1",
		`(pull|pr)/(\d+)`: "https://github.com/thomaso-mirodin/go-shorten/pull/$2",
	})
	if err != nil {
		t.Fatal(errors.Wrap(err, "failed to create storage.Regex"))
	}

	testTable := []struct {
		in       string
		expected string
		err      error
	}{
		{in: "jira/ABC-1234", expected: "https://atlassian.net/browse/ABC-1234"},
		{in: "pull/1234", expected: "https://github.com/thomaso-mirodin/go-shorten/pull/1234"},
		{in: "pr/1234", expected: "https://github.com/thomaso-mirodin/go-shorten/pull/1234"},
		{in: "pr/", err: ErrShortNotSet},
		{in: "asdflkj", err: ErrShortNotSet},
	}

	for _, tt := range testTable {
		t.Logf("Table: %#v", tt)
		actual, err := r.Load(tt.in)
		if err != tt.err {
			t.Errorf("actual err (%q) != expected err (%q)", err, tt.err)
		}

		if actual != tt.expected {
			t.Errorf("actual result (%q) != expected result (%q)", actual, tt.expected)
		}
	}
}
