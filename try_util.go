package try

import (
	"fmt"

	"github.com/pkg/errors"
)

/*
Converts an arbitrary value to an error. Should be used with `recover()`:

	err := Err(recover())

Caution: `recover()` works ONLY when called DIRECTLY inside a deferred function.
When called ANYWHERE ELSE, even in functions called BY a deferred function,
it's a nop.

When called with `nil`, returns `nil`. When called with a non-nil non-error
value, wraps it into `Val` which implements the `error` interface.
*/
func Err(val interface{}) error {
	if val == nil {
		return nil
	}

	err, _ := val.(error)
	if err != nil {
		return WithStack(err)
	}

	return errors.WithStack(Val{val})
}

/*
Adds a stacktrace via "github.com/pkg/errors", unless the error already has one.
This exists because `errors.WithStack` ALWAYS wraps an error and adds a
stacktrace, even when it would be redundant.

Should be used when it's unknown whether the error has a stacktrace. When the
error is known to NOT have a stacktrace, use `errors.WithStack`, because this
function adds its own frame, and `errors.WithStack` does not.

When called with `nil`, returns `nil`.
*/
func WithStack(err error) error {
	if !HasStack(err) {
		return errors.WithStack(err)
	}
	return err
}

/*
True if this error, or any of the errors it wraps, has a stacktrace provided by
"github.com/pkg/errors".
*/
func HasStack(err error) bool {
	for {
		if err == nil {
			return false
		}

		// Hidden interface implemented by some types in "github.com/pkg/errors".
		_, ok := err.(interface{ StackTrace() errors.StackTrace })
		if ok {
			return true
		}

		cause := errors.Unwrap(err)
		if cause == err {
			return false
		}
		err = cause
	}
}

// Used by `Err()` to wrap non-errors received from `recover()` and convert them
// to errors.
type Val struct{ Value interface{} }

// Implement `error`.
func (self Val) Error() string {
	if self.Value != nil {
		return fmt.Sprint(self.Value)
	}
	return ""
}

// Implement error unwrapping, in case an `error` gets accidentally converted to
// `interface{}` before ending up here.
func (self Val) Unwrap() error {
	err, _ := self.Value.(error)
	return err
}
