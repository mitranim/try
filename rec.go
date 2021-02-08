package try

/*
Recovers from panics, writing the resulting error, if any, to the given pointer.
Should be used together with "try"-style functions.

Usage:

	func someFunc() (err error) {
		defer try.Rec(&err)
		try.To(errors.New(`A`)) // Will panic and be returned.
		try.To(errors.New(`B`)) // Will panic and be returned.
		try.To(nil)             // Will not panic.
		return
	}

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
Usage:

	func someFunc(errChan chan error) {
		defer try.RecChan(errChan)
		try.To(errors.New(`A`)) // Will panic and be sent.
		try.To(errors.New(`B`)) // Will panic and be sent.
		try.To(nil)             // Will not panic, will not be sent.
	}
*/
func RecChan(errChan chan<- error) {
	maybeSend(errChan, Err(recover()))
}

/*
Recovery for background goroutines that have nowhere to return their error.
Unlike the other "rec" functions, this doesn't send the error anywhere; instead
it calls the provided function ONLY if the error was non-nil.

Functions that CAN return errors should use the other "rec" functions instead.

Usage:

	func bgFun() {
		defer try.RecWith(func(err error) {
			log.Printf("failed to X: %+v\n", err)
		})
		panic("fail")
	}

	go bgFun()
*/
func RecWith(fun func(error)) {
	err := Err(recover())
	if err != nil {
		fun(err)
	}
}

/*
Combination of `Rec` and `WithMessage`. Recovers from panics and adds a message.
Usage:

	func someFunc() (err error) {
		defer try.RecWithMessage(&err, `failed to X`)
		try.To(errors.New(`A`)) // Will panic, be wrapped, be returned.
		try.To(errors.New(`B`)) // Will panic, be wrapped, be returned.
		try.To(nil)             // Will not panic, will not be wrapped.
		return
	}
*/
func RecWithMessage(err *error, msg string) {
	maybeSet(err, Err(recover()))
	WithMessage(err, msg)
}

/*
Combination of `Rec` and `WithMessagef`. Recovers from panics and adds a
message. Usage:

	func someFunc() (err error) {
		defer try.RecWithMessagef(&err, `failed to %v`, `X`)
		try.To(errors.New(`A`)) // Will panic, be wrapped, be returned.
		try.To(errors.New(`B`)) // Will panic, be wrapped, be returned.
		try.To(nil)             // Will not panic, will not be wrapped.
		return
	}
*/
func RecWithMessagef(err *error, pattern string, args ...interface{}) {
	maybeSet(err, Err(recover()))
	WithMessagef(err, pattern, args...)
}
