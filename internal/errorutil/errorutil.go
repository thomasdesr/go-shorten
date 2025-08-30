package errorutil

import (
	"fmt"
)

// Wrap annotates `err` with the provided `msg`. It always returns a non-empty
// error that is not nil, even when err is nil or msg is empty.
func Wrap(err error, msg string) error {
	if msg == "" {
		msg = "<wrapped error with no message>"
	}

	if err == nil {
		return fmt.Errorf("%s: <nil error>", msg)
	}
	return fmt.Errorf("%s: %w", msg, err)
}

// Wrapf annotates `err` with the provided format string and arguments.
// It always returns a non-empty error that is not nil, even when err is nil
// or the formatted error message is empty.
func Wrapf(err error, f string, args ...interface{}) error {
	return Wrap(err, fmt.Sprintf(f, args...))
}
