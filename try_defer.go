package try

import "github.com/pkg/errors"

/*
Must be deferred. Tool for adding a stacktrace to an arbitrary panic. Unlike
the "rec" functions, this does NOT prevent the panic from propagating. It
simply ensures that there's a stacktrace, then re-panics.

Caution: due to idiosyncrasies of `recover()`, this works ONLY when deferred
directly. Anything other than `defer try.Trace()` will NOT work.
*/
func Trace() { To(Err(recover())) }

/*
Must be deferred. Runs the function only if there's no panic. Idempotently adds
a stacktrace.
*/
func Ok(fun func()) {
	To(Err(recover()))
	fun()
}

/*
Must be deferred. Runs the function ONLY if there's an ongoing panic, and then
re-panics. Idempotently adds a stacktrace.
*/
func Fail(fun func(error)) {
	err := Err(recover())
	if err != nil && fun != nil {
		fun(err)
	}
	To(err)
}

/*
Must be deferred. Short for "transmute" or "transform". Catches an ongoing
panic, transforms the error by calling the provided function, and then
re-panics via `To`. Can be used to ignore specific errors, by converting them
to nil, which prevents the second panic. Idempotently adds a stacktrace.
*/
func Trans(fun func(error) error) {
	err := Err(recover())
	if err != nil && fun != nil {
		err = fun(err)
	}
	To(err)
}

/*
Must be deferred. Wraps non-nil panics, prepending the error message and
idempotently adding a stacktrace.
*/
func Detail(msg string) {
	To(errors.WithMessage(Err(recover()), msg))
}

/*
Must be deferred. Wraps non-nil panics, prepending the error message and
idempotently adding a stacktrace.
*/
func Detailf(msg string, args ...interface{}) {
	To(errors.WithMessagef(Err(recover()), msg, args...))
}

/*
Must be deferred. Catches panics; ignores errors that satisfy the provided
test; re-panics on other non-nil errors. Idempotently adds a stacktrace.
*/
func Ignore(test func(error) bool) {
	err := Err(recover())
	if err != nil && test != nil && test(err) {
		return
	}
	To(err)
}

/*
Must be deferred. Recovers from panics, writing the resulting error, if any, to
the given pointer. Should be used together with "try"-style functions.
Idempotently adds a stacktrace.
*/
func Rec(ptr *error) {
	err := Err(recover())
	if err != nil {
		*ptr = err
	}
}

/*
Must be deferred. Filtered version of `Rec`. Recovers from panics that satisfy
the provided test. Re-panics on non-nil errors that don't satisfy the test.
Does NOT check errors that are returned normally, without a panic. Should be
used together with "try"-style functions. Idempotently adds a stacktrace.
*/
func RecOnly(ptr *error, test func(error) bool) {
	err := Err(recover())
	if err != nil {
		*ptr = err
		if test != nil && test(err) {
			return
		}
		panic(err)
	}
}

/*
Must be deferred. Version of `Rec` that sends the recovered error, if any, to
the given channel. Idempotently adds a stacktrace.
*/
func RecChan(errChan chan<- error) {
	err := Err(recover())
	if err != nil {
		select {
		case errChan <- err:
		default:
		}
	}
}

/*
Must be deferred. Recovery for background goroutines that have nowhere to return
their error. Unlike the other "rec" functions, this doesn't send the error
anywhere; instead it calls the provided function ONLY if the error was
non-nil.

Functions that CAN return errors should use the other "rec" functions instead.
*/
func RecWith(fun func(error)) {
	err := Err(recover())
	if err != nil {
		fun(err)
	}
}

/*
Must be deferred. Combination of `Rec` and `WithMessage`. Recovers from panics
and adds a message. Idempotently adds a stacktrace.
*/
func RecWithMessage(ptr *error, msg string) {
	err := Err(recover())
	if err != nil {
		*ptr = errors.WithMessage(err, msg)
	}
}

/*
Must be deferred. Combination of `Rec` and `WithMessagef`. Recovers from panics
and adds a message. Idempotently adds a stacktrace.
*/
func RecWithMessagef(ptr *error, pattern string, args ...interface{}) {
	err := Err(recover())
	if err != nil {
		*ptr = errors.WithMessagef(err, pattern, args...)
	}
}

/*
Must be deferred. Wraps a non-nil error, prepending the message. Unlike
`RecWithMessage`, does NOT implicitly recover or add a stacktrace.
*/
func WithMessage(ptr *error, msg string) {
	if ptr != nil && *ptr != nil {
		*ptr = errors.WithMessage(*ptr, msg)
	}
}

/*
Must be deferred. Wraps a non-nil error, prepending the message. Unlike
`RecWithMessagef`, does NOT implicitly recover or add a stacktrace.
*/
func WithMessagef(ptr *error, pattern string, args ...interface{}) {
	if ptr != nil && *ptr != nil {
		*ptr = errors.WithMessagef(*ptr, pattern, args...)
	}
}
