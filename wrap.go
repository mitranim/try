package try

import "github.com/pkg/errors"

/*
Deferred version of "github.com/pkg/errors".WithMessage. Wraps a non-nil error
when the current function returns. Unlike `RecWithMessage`, does NOT implicitly
recover. Usage:

	func someFunc() (err error) {
		defer try.WithMessage(&err, `failed to X`)
		return errors.New(`A`) // Will be wrapped.
		return errors.New(`B`) // Will be wrapped.
		return nil             // Will not be wrapped.
	}
*/
func WithMessage(err *error, msg string) {
	if err != nil && *err != nil {
		*err = errors.WithMessage(*err, msg)
	}
}

/*
Deferred version of "github.com/pkg/errors".WithMessagef. Wraps a non-nil error
when the current function returns. Unlike `RecWithMessagef`, does NOT implicitly
recover. Usage:

	func someFunc() (err error) {
		defer try.WithMessagef(&err, `failed to %v`, `X`)
		return errors.New(`A`) // Will be wrapped.
		return errors.New(`B`) // Will be wrapped.
		return nil             // Will not be wrapped.
	}
*/
func WithMessagef(err *error, pattern string, args ...interface{}) {
	if err != nil && *err != nil {
		*err = errors.WithMessagef(*err, pattern, args...)
	}
}
