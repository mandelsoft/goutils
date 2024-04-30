package general

import (
	"fmt"
)

// Must expect a result to be provided without error.
func Must[T any](o T, err error) T {
	if err != nil {
		panic(fmt.Errorf("expected a %T, but got error %w", o, err))
	}
	return o
}
