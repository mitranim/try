/*
Shorter error handling in Go, like the rejected "try" proposal
(https://golang.org/design/32437-try-builtin), and fairly similar to Swift's
(but with stacktraces).

Uses a combination of `defer` and panics to make code SIGNIFICANTLY shorter, at
an acceptable runtime cost. Keeps `error` in the function signatures. Also
automatically ensures stacktraces, by tightly coupling
to "github.com/pkg/errors".

See `readme.md` for examples.
*/
package try

/*
Simplifies control flow by panicking on non-nil errors. Should be used in
conjunction with `Rec`:

	func someFunc() (err error) {
		defer try.Rec(&err)
		try.To(errors.New(`A`)) // Will panic and be returned.
		try.To(errors.New(`B`)) // Will panic and be returned.
		try.To(nil)             // Will not panic.
		return
	}

	err := someFunc()

If the error doesn't already have a stacktrace, adds one
via "github.com/pkg/errors". Stacktraces are essential for such exception-like
control flow. Without them, debugging would be incredibly tedious.
*/
func To(err error) {
	if err != nil {
		panic(WithStack(err))
	}
}

/*
A "try" function that takes and returns a value of type `interface{} value. Usage:

	func someFunc() (interface{}, error) {return "val", errors.New(`err`)}
	val := try.Interface(someFunc())
*/
func Interface(val interface{}, err error) interface{} {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `bool`. Usage:

	func someFunc() (bool, error) {return true, errors.New(`err`)}
	val := try.Bool(someFunc())
*/
func Bool(val bool, err error) bool {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `uint8`. Usage:

	func someFunc() (uint8, error) {return 255, errors.New(`err`)}
	val := try.Uint8(someFunc())
*/
func Uint8(val uint8, err error) uint8 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `uint16`. Usage:

	func someFunc() (uint16, error) {return 255, errors.New(`err`)}
	val := try.Uint16(someFunc())
*/
func Uint16(val uint16, err error) uint16 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `uint32`. Usage:

	func someFunc() (uint32, error) {return 255, errors.New(`err`)}
	val := try.Uint32(someFunc())
*/
func Uint32(val uint32, err error) uint32 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `uint64`. Usage:

	func someFunc() (uint64, error) {return 255, errors.New(`err`)}
	val := try.Uint64(someFunc())
*/
func Uint64(val uint64, err error) uint64 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `byte`. Usage:

	func someFunc() (byte, error) {return 255, errors.New(`err`)}
	val := try.Byte(someFunc())
*/
func Byte(val byte, err error) byte {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `int8`. Usage:

	func someFunc() (int8, error) {return 255, errors.New(`err`)}
	val := try.Int8(someFunc())
*/
func Int8(val int8, err error) int8 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `int16`. Usage:

	func someFunc() (int16, error) {return 255, errors.New(`err`)}
	val := try.Int16(someFunc())
*/
func Int16(val int16, err error) int16 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `rune`. Usage:

	func someFunc() (rune, error) {return 255, errors.New(`err`)}
	val := try.Rune(someFunc())
*/
func Rune(val rune, err error) rune {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `int32`. Usage:

	func someFunc() (int32, error) {return 255, errors.New(`err`)}
	val := try.Int32(someFunc())
*/
func Int32(val int32, err error) int32 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `int64`. Usage:

	func someFunc() (int64, error) {return 255, errors.New(`err`)}
	val := try.Int64(someFunc())
*/
func Int64(val int64, err error) int64 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `float32`. Usage:

	func someFunc() (float32, error) {return 255, errors.New(`err`)}
	val := try.Float32(someFunc())
*/
func Float32(val float32, err error) float32 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `float64`. Usage:

	func someFunc() (float64, error) {return 255, errors.New(`err`)}
	val := try.Float64(someFunc())
*/
func Float64(val float64, err error) float64 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `complex64`. Usage:

	func someFunc() (complex64, error) {return 255, errors.New(`err`)}
	val := try.Complex64(someFunc())
*/
func Complex64(val complex64, err error) complex64 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `complex128`. Usage:

	func someFunc() (complex128, error) {return 255, errors.New(`err`)}
	val := try.Complex128(someFunc())
*/
func Complex128(val complex128, err error) complex128 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `string`. Usage:

	func someFunc() (string, error) {return 255, errors.New(`err`)}
	val := try.String(someFunc())
*/
func String(val string, err error) string {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `int`. Usage:

	func someFunc() (int, error) {return 255, errors.New(`err`)}
	val := try.Int(someFunc())
*/
func Int(val int, err error) int {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `uint`. Usage:

	func someFunc() (uint, error) {return 255, errors.New(`err`)}
	val := try.Uint(someFunc())
*/
func Uint(val uint, err error) uint {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `uintptr`. Usage:

	func someFunc() (uintptr, error) {return 255, errors.New(`err`)}
	val := try.Uintptr(someFunc())
*/
func Uintptr(val uintptr, err error) uintptr {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]interface{}`. Usage:

	func someFunc() ([]interface{}, error) {return []interface{}{"val"}, errors.New(`err`)}
	val := try.InterfaceSlice(someFunc())
*/
func InterfaceSlice(val []interface{}, err error) []interface{} {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]bool`. Usage:

	func someFunc() ([]bool, error) {return []bool{true}, errors.New(`err`)}
	val := try.BoolSlice(someFunc())
*/
func BoolSlice(val []bool, err error) []bool {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]uint8`. Usage:

	func someFunc() ([]uint8, error) {return []uint8{255}, errors.New(`err`)}
	val := try.Uint8Slice(someFunc())
*/
func Uint8Slice(val []uint8, err error) []uint8 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]uint16`. Usage:

	func someFunc() ([]uint16, error) {return []uint16{255}, errors.New(`err`)}
	val := try.Uint16Slice(someFunc())
*/
func Uint16Slice(val []uint16, err error) []uint16 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]uint32`. Usage:

	func someFunc() ([]uint32, error) {return []uint32{255}, errors.New(`err`)}
	val := try.Uint32Slice(someFunc())
*/
func Uint32Slice(val []uint32, err error) []uint32 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]uint64`. Usage:

	func someFunc() ([]uint64, error) {return []uint64{255}, errors.New(`err`)}
	val := try.Uint64Slice(someFunc())
*/
func Uint64Slice(val []uint64, err error) []uint64 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]byte`. Usage:

	func someFunc() ([]byte, error) {return []byte{255}, errors.New(`err`)}
	val := try.ByteSlice(someFunc())
*/
func ByteSlice(val []byte, err error) []byte {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]int8`. Usage:

	func someFunc() ([]int8, error) {return []int8{255}, errors.New(`err`)}
	val := try.Int8Slice(someFunc())
*/
func Int8Slice(val []int8, err error) []int8 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]int16`. Usage:

	func someFunc() ([]int16, error) {return []int16{255}, errors.New(`err`)}
	val := try.Int16Slice(someFunc())
*/
func Int16Slice(val []int16, err error) []int16 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]rune`. Usage:

	func someFunc() ([]rune, error) {return []rune{255}, errors.New(`err`)}
	val := try.RuneSlice(someFunc())
*/
func RuneSlice(val []rune, err error) []rune {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]int32`. Usage:

	func someFunc() ([]int32, error) {return []int32{255}, errors.New(`err`)}
	val := try.Int32Slice(someFunc())
*/
func Int32Slice(val []int32, err error) []int32 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]int64`. Usage:

	func someFunc() ([]int64, error) {return []int64{255}, errors.New(`err`)}
	val := try.Int64Slice(someFunc())
*/
func Int64Slice(val []int64, err error) []int64 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]float32`. Usage:

	func someFunc() ([]float32, error) {return []float32{255}, errors.New(`err`)}
	val := try.Float32Slice(someFunc())
*/
func Float32Slice(val []float32, err error) []float32 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]float64`. Usage:

	func someFunc() ([]float64, error) {return []float64{255}, errors.New(`err`)}
	val := try.Float64Slice(someFunc())
*/
func Float64Slice(val []float64, err error) []float64 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]complex64`. Usage:

	func someFunc() ([]complex64, error) {return []complex64{255}, errors.New(`err`)}
	val := try.Complex64Slice(someFunc())
*/
func Complex64Slice(val []complex64, err error) []complex64 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]complex128`. Usage:

	func someFunc() ([]complex128, error) {return []complex128{255}, errors.New(`err`)}
	val := try.Complex128Slice(someFunc())
*/
func Complex128Slice(val []complex128, err error) []complex128 {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]string`. Usage:

	func someFunc() ([]string, error) {return []string{"val"}, errors.New(`err`)}
	val := try.StringSlice(someFunc())
*/
func StringSlice(val []string, err error) []string {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]int`. Usage:

	func someFunc() ([]int, error) {return []int{255}, errors.New(`err`)}
	val := try.IntSlice(someFunc())
*/
func IntSlice(val []int, err error) []int {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]uint`. Usage:

	func someFunc() ([]uint, error) {return []uint{255}, errors.New(`err`)}
	val := try.UintSlice(someFunc())
*/
func UintSlice(val []uint, err error) []uint {
	To(err)
	return val
}

/*
A "try" function that takes and returns a value of type `[]uintptr`. Usage:

	func someFunc() ([]uintptr, error) {return []uintptr{255}, errors.New(`err`)}
	val := try.UintptrSlice(someFunc())
*/
func UintptrSlice(val []uintptr, err error) []uintptr {
	To(err)
	return val
}
