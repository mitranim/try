package try

/*
Recovers from panics, writing the resulting error, if any, to the given pointer.
Should be used together with "try"-style functions.

This function, and all other "rec"-style functions, implicitly add a stacktrace
to a panic-recovered error if it doesn't already have one. This is redundant
with "try" functions, which add stacktraces pointing to a more meaningful
location, but very helpful for debugging unintended panics. When an uncaught
panic crashes the process, Go will print a stacktrace, even if the panic value
was a primitive. But because `Rec` overrides such panics, we lose
this "automatic" stacktrace, and must add our own to ensure we don't make
debugging difficult.
*/
func Rec(err *error) {
	maybeSet(err, Err(recover()))
}

/*
Version of `Rec` that sends the recovered error, if any, to the given channel.
*/
func RecChan(errChan chan<- error) {
	maybeSend(errChan, Err(recover()))
}

/*
Recovery for background goroutines that have nowhere to return their error.
Unlike the other "rec" functions, this doesn't send the error anywhere; instead
it calls the provided function ONLY if the error was non-nil.

Functions that CAN return errors should use the other "rec" functions instead.
*/
func RecWith(fun func(error)) {
	err := Err(recover())
	if err != nil {
		fun(err)
	}
}

/*
Combination of `Rec` and `WithMessage`. Recovers from panics and adds a message.
*/
func RecWithMessage(err *error, msg string) {
	maybeSet(err, Err(recover()))
	WithMessage(err, msg)
}

/*
Combination of `Rec` and `WithMessagef`. Recovers from panics and adds a
message.
*/
func RecWithMessagef(err *error, pattern string, args ...interface{}) {
	maybeSet(err, Err(recover()))
	WithMessagef(err, pattern, args...)
}

// Runs a panicking function, returning the caught error if any.
func Unpanic(fun func()) (err error) {
	defer Rec(&err)
	fun()
	return
}
