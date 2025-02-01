package errors

import (
	"errors"
	errx "github.com/pkg/errors"
)

var (
	As = errors.As
	Is = errors.Is

	Cause  = errx.Cause
	Unwrap = errx.Unwrap
	Wrap   = errx.Wrap
)
