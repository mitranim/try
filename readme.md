## Overview

Shorter error handling in Go, like the rejected [`try` proposal](https://golang.org/design/32437-try-builtin), and fairly similar to Swift's.

Uses a combination of `defer` and panics to make code _significantly_ shorter, at an acceptable runtime cost. **Keeps `error` in the function signatures.** Also automatically ensures stacktraces, by tightly coupling to "github.com/pkg/errors".

See API docs at https://pkg.go.dev/github.com/mitranim/try.

## TOC

* [Why](#why)
* [Limitations](#limitations)
* [Nitpicks](#nitpicks)

## Why

Go wants you to add meaningful context when handling errors. I sympathize with this idea, and do it often. But there's code where annotating every single failure is not practical and/or bloats the code beyond our ability to _read it back_.

```golang
func someFunc() error {
  err := someFunc()
  if err != nil {
    return errors.WithMessage(err, `failed to X`)
  }
  err := someFunc()
  if err != nil {
    return errors.WithMessage(err, `failed to X`)
  }
  err := someFunc()
  if err != nil {
    return errors.WithMessage(err, `failed to X`)
  }
  return nil
}
```

```golang
func someFunc() (err error) {
  defer try.RecWithMessage(&err, `failed to X`)
  try.To(someFunc())
  try.To(someFunc())
  try.To(someFunc())
  return
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

In the current state of Go, functions conforming to this pattern are easier to compose (in a vague sense), lending themselves to shorter code.

## Nitpicks

The term "must" is more conventional in the Go standard library, but this library uses "try" because it's more grammatically flexible: "try string" works, but "must string" would not. The "try" proposal used "try". Swift error handling is very similar and uses "try".

## License

https://unlicense.org

## Misc

I'm receptive to suggestions. If this library _almost_ satisfies you but needs changes, open an issue or chat me up. Contacts: https://mitranim.com/#contacts
