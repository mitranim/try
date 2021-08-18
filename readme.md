## Overview

Shorter error handling for Go. Supports two approaches:

* Like the rejected [`try` proposal](https://golang.org/design/32437-try-builtin).
* Exceptions-based.

Features:

* Uses a combination of `defer` and panics to make code _significantly_ shorter, at an acceptable runtime cost.
* Automatically ensures stacktraces via ["github.com/pkg/errors"](https://github.com/pkg/errors).
* You can choose to keep `error` in signatures and use explicit "try".
* You can choose to drop `error` from signatures and use exceptions.

See API docs at https://pkg.go.dev/github.com/mitranim/try.

## TOC

* [Why](#why)
* [Limitations](#limitations)
* [Naming](#naming)
* [Changelog](#changelog)

## Why

Go wants you to add meaningful context when handling errors. I sympathize with this idea, and do it often. But there's code where annotating every single failure is not practical and/or bloats the code beyond our ability to _read it back_.

```golang
func someFuncA() error {
  err := someFuncB()
  if err != nil {
    return errors.WithMessage(err, `failed to X`)
  }
  err = someFuncC()
  if err != nil {
    return errors.WithMessage(err, `failed to X`)
  }
  err = someFuncD()
  if err != nil {
    return errors.WithMessage(err, `failed to X`)
  }
  return nil
}
```

Using the "try" style:

```golang
func someFuncA() (err error) {
  defer try.RecWithMessage(&err, `failed to X`)
  try.To(someFuncB())
  try.To(someFuncC())
  try.To(someFuncD())
  return
}
```

Using the "exceptions" style:

```golang
func someFuncA() {
  defer try.Detail(`failed to X`)
  someFuncB()
  someFuncC()
  someFuncD()
}
```

The code should speak for itself. This won't be usable for _every_ codebase, see [Limitations](#limitations) below, but can be a nice improvement for some.

## Limitations

This has a minor runtime cost. Use this for IO-heavy control code, where the cost is barely measurable. Avoid this in CPU-heavy code or libraries.

This package provides a variety of "try" functions for common cases, but it can't define something generic like the original proposal did. To make your code compatible, prefer to use pointers for "inout" parameters of non-primitive types, and return only `error`:

```golang
func someFunc(input A, out *B) error {
  *out = someOperation(input)
  return someErr
}

var val B
try.To(someFunc(input, &val))
```

...Or use inout parameters and panics:

```golang
func someFunc(input A, out *B) {
  *out = someOperation(input)
}

var val B
someFunc(input, &val)
```

In the current state of Go, functions conforming to this pattern are easier to compose, leading to much shorter code.

## Naming

The term "must" is more conventional in the Go standard library, but this library uses "try" because it's more grammatically flexible: "try string" works, but "must string" would not. The "try" proposal used "try". Swift error handling is very similar and uses "try". (Unlike Swift, we have stacktraces.)

## Changelog

### v0.1.4

Breaking: renamed `Caught` to `CaughtOnly` for consistency, added `Caught`.

### v0.1.3

Added `DetailOnly` and `DetailOnlyf`.

### v0.1.2

Added tools to support the "exceptions" style. For many apps, it's a better fit than either the Go style or the "try" style.

## License

https://unlicense.org

## Misc

I'm receptive to suggestions. If this library _almost_ satisfies you but needs changes, open an issue or chat me up. Contacts: https://mitranim.com/#contacts
