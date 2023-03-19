package bcapi

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrBestchange = errors.New("bcapi")
)

func wrapError(err string) error {
	return fmt.Errorf("%w: %s", ErrBestchange, err)
}

func wrapErrors(errs ...string) error {
	return fmt.Errorf("%w: %s", ErrBestchange, strings.Join(errs, ": "))
}
