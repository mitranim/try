package try

/*
Converts a panic to an error, idempotently adding a stacktrace.
*/
func Catch(fun func()) (err error) {
	defer Rec(&err)
	if fun != nil {
		fun()
	}
	return
}

/*
Converts a panic to an error, if the error satisfies the provided test.
Otherwise re-panics. Idempotently adds a stacktrace.
*/
func CatchOnly(test func(error) bool, fun func()) (err error) {
	defer RecOnly(&err, test)
	if fun != nil {
		fun()
	}
	return
}

/*
Shortcut for `CatchOnly() != nil`. Useful when you want to handle and ignore a
specific error, but don't care about its content.
*/
func Caught(test func(error) bool, fun func()) bool {
	return CatchOnly(test, fun) != nil
}

/*
Runs a function, catching and ignoring only the panics that satisfy the provided
test. Idempotently adds a stacktrace to all panics.
*/
func Ignoring(test func(error) bool, fun func()) {
	defer Ignore(test)
	if fun != nil {
		fun()
	}
}

/*
Runs a function, "transmuting" the resulting panics by calling the provided
transformer, which may choose to suppress or wrap specific error types. See
`Trans`.
*/
func WithTrans(trans func(error) error, fun func()) {
	defer Trans(trans)
	if fun != nil {
		fun()
	}
}
