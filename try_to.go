/*
Shorter error handling in Go. Supports two styles:

* Like the rejected "try" proposal (https://golang.org/design/32437-try-builtin).

* Exception-based.

Uses a combination of `defer` and panics to make code SIGNIFICANTLY shorter, at
an acceptable runtime cost. Automatically ensures stacktraces
via "github.com/pkg/errors". You can choose to keep `error` in signatures and
use explicit "try", or drop `error` from signatures and use exceptions.

See `readme.md` and examples.
*/
package try

/*
Simplifies control flow by panicking on non-nil errors. Should be used in
conjunction with `Rec`.

If the error doesn't already have a stacktrace, adds one
via "github.com/pkg/errors". Stacktraces are essential for such exception-like
control flow. Without them, debugging would be incredibly tedious.
*/
func To(err error) {
	if err != nil {
		panic(WithStack(err))
	}
}

// A "try" function that takes and returns a value of type `interface{} value.
func Interface(val interface{}, err error) interface{} {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `bool`.
func Bool(val bool, err error) bool {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `uint8`.
func Uint8(val uint8, err error) uint8 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `uint16`.
func Uint16(val uint16, err error) uint16 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `uint32`.
func Uint32(val uint32, err error) uint32 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `uint64`.
func Uint64(val uint64, err error) uint64 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `byte`.
func Byte(val byte, err error) byte {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `int8`.
func Int8(val int8, err error) int8 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `int16`.
func Int16(val int16, err error) int16 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `rune`.
func Rune(val rune, err error) rune {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `int32`.
func Int32(val int32, err error) int32 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `int64`.
func Int64(val int64, err error) int64 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `float32`.
func Float32(val float32, err error) float32 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `float64`.
func Float64(val float64, err error) float64 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `complex64`.
func Complex64(val complex64, err error) complex64 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `complex128`.
func Complex128(val complex128, err error) complex128 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `string`.
func String(val string, err error) string {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `int`.
func Int(val int, err error) int {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `uint`.
func Uint(val uint, err error) uint {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `uintptr`.
func Uintptr(val uintptr, err error) uintptr {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]interface{}`.
func InterfaceSlice(val []interface{}, err error) []interface{} {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]bool`.
func BoolSlice(val []bool, err error) []bool {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]uint8`.
func Uint8Slice(val []uint8, err error) []uint8 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]uint16`.
func Uint16Slice(val []uint16, err error) []uint16 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]uint32`.
func Uint32Slice(val []uint32, err error) []uint32 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]uint64`.
func Uint64Slice(val []uint64, err error) []uint64 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]byte`.
func ByteSlice(val []byte, err error) []byte {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]int8`.
func Int8Slice(val []int8, err error) []int8 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]int16`.
func Int16Slice(val []int16, err error) []int16 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]rune`.
func RuneSlice(val []rune, err error) []rune {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]int32`.
func Int32Slice(val []int32, err error) []int32 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]int64`.
func Int64Slice(val []int64, err error) []int64 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]float32`.
func Float32Slice(val []float32, err error) []float32 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]float64`.
func Float64Slice(val []float64, err error) []float64 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]complex64`.
func Complex64Slice(val []complex64, err error) []complex64 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]complex128`.
func Complex128Slice(val []complex128, err error) []complex128 {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]string`.
func StringSlice(val []string, err error) []string {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]int`.
func IntSlice(val []int, err error) []int {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]uint`.
func UintSlice(val []uint, err error) []uint {
	To(err)
	return val
}

// A "try" function that takes and returns a value of type `[]uintptr`.
func UintptrSlice(val []uintptr, err error) []uintptr {
	To(err)
	return val
}
