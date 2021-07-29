package try_test

import (
	"fmt"
	"log"
	"os"

	"github.com/mitranim/try"
	"github.com/pkg/errors"
)

func ExampleTo() {
	someFunc := func() (err error) {
		defer try.Rec(&err)
		try.To(errors.New(`failure A`)) // Will panic, error will be returned.
		try.To(errors.New(`failure B`)) // Not executed; would panic.
		try.To(nil)                     // Not executed; would not panic.
		return
	}
	err := someFunc()
	fmt.Println(err)
	// Output:
	// failure A
}

func ExampleRec() {
	someFunc := func() (err error) {
		defer try.Rec(&err)
		try.To(errors.New(`failure A`)) // Will panic, error will be returned.
		try.To(errors.New(`failure B`)) // Not executed; would panic.
		try.To(nil)                     // Not executed; would not panic.
		return
	}
	err := someFunc()
	fmt.Println(err)
	// Output:
	// failure A
}

func ExampleRecOnly() {
	isErrNoFile := func(err error) bool {
		return errors.Is(err, os.ErrNotExist)
	}

	someFunc := func() (err error) {
		defer try.RecOnly(&err, isErrNoFile)
		_ = try.ByteSlice(os.ReadFile(`non-existent-file`))
		fmt.Println(`file exists`)
		return
	}

	err := someFunc()
	fmt.Println(err)
	// Output:
	// open non-existent-file: no such file or directory
}

func ExampleRecChan() {
	someFunc := func(errChan chan error) {
		defer try.RecChan(errChan)
		try.To(errors.New(`failure A`)) // Will panic, error will be sent.
		try.To(errors.New(`failure B`)) // Not executed; would panic.
		try.To(nil)                     // Not executed; would not panic.
	}

	errs := make(chan error, 256)
	someFunc(errs)
}

func ExampleRecWith() {
	bgFun := func() {
		defer try.RecWith(func(err error) {
			log.Printf("failed to X: %+v\n", err)
		})
		panic("fail")
	}
	go bgFun()
}

func ExampleRecWithMessage() {
	someFunc := func() (err error) {
		defer try.RecWithMessage(&err, `failed to X`)
		try.To(errors.New(`failure A`)) // Will panic, error will be wrapped and returned.
		try.To(errors.New(`failure B`)) // Not executed; would panic.
		try.To(nil)                     // Not executed; would not panic.
		return
	}
	err := someFunc()
	fmt.Println(err)
	// Output:
	// failed to X: failure A
}

func ExampleRecWithMessagef() {
	someFunc := func() (err error) {
		defer try.RecWithMessagef(&err, `failed to %v`, `X`)
		try.To(errors.New(`failure A`)) // Will panic, error will be wrapped and returned.
		try.To(errors.New(`failure B`)) // Not executed; would panic.
		try.To(nil)                     // Not executed; would not panic.
		return
	}
	err := someFunc()
	fmt.Println(err)
	// Output:
	// failed to X: failure A
}

func ExampleWithMessage() {
	someFunc := func() (err error) {
		defer try.WithMessage(&err, `failed to X`)
		return errors.New(`failure A`) // Will be wrapped.
		return errors.New(`failure B`) // Not executed; would panic.
		return nil                     // Not executed; would not panic.
	}
	err := someFunc()
	fmt.Println(err)
	// Output:
	// failed to X: failure A
}

func ExampleWithMessagef() {
	someFunc := func() (err error) {
		defer try.WithMessagef(&err, `failed to %v`, `X`)
		return errors.New(`failure A`) // Will panic, error will be wrapped.
		return errors.New(`failure B`) // Not executed; would panic.
		return nil                     // Not executed; would not panic.
	}
	err := someFunc()
	fmt.Println(err)
	// Output:
	// failed to X: failure A
}

func ExampleDetail() {
	someFunc := func() {
		defer try.Detail(`failed to X`)
		try.To(errors.New(`failure A`)) // Will be wrapped.
		try.To(errors.New(`failure B`)) // Not executed; would panic.
		try.To(nil)                     // Not executed; would not panic.
	}
	err := try.Catch(someFunc)
	fmt.Println(err)
	// Output:
	// failed to X: failure A
}

func ExampleDetailf() {
	someFunc := func() {
		defer try.Detailf(`failed to %v`, `X`)
		try.To(errors.New(`failure A`)) // Will be wrapped.
		try.To(errors.New(`failure B`)) // Not executed; would panic.
		try.To(nil)                     // Not executed; would not panic.
	}
	err := try.Catch(someFunc)
	fmt.Println(err)
	// Output:
	// failed to X: failure A
}

func ExampleTrace() {
	defer try.Trace()
	if false {
		panic("unreachable")
	}
}

func ExampleCatch() {
	someFunc := func() {
		try.To(errors.New(`failure A`)) // Will panic.
		try.To(errors.New(`failure B`)) // Not executed; would panic.
		try.To(nil)                     // Not executed; would not panic.
	}
	err := try.Catch(someFunc)
	fmt.Println(err)
	// Output:
	// failure A
}

func ExampleCatchOnly() {
	isErrNoFile := func(err error) bool {
		return errors.Is(err, os.ErrNotExist)
	}

	maybeRead := func() {
		fmt.Println(try.ByteSlice(os.ReadFile(`non-existent-file`)))
	}

	err := try.CatchOnly(isErrNoFile, maybeRead)
	fmt.Println(err)
	// Output:
	// open non-existent-file: no such file or directory
}

func ExampleCaught() {
	isErrNoFile := func(err error) bool {
		return errors.Is(err, os.ErrNotExist)
	}

	maybeRead := func() {
		fmt.Println(try.ByteSlice(os.ReadFile(`non-existent-file`)))
	}

	fmt.Println(try.Caught(isErrNoFile, maybeRead))
	// Output:
	// true
}

func ExampleIgnoring() {
	isErrNoFile := func(err error) bool {
		return errors.Is(err, os.ErrNotExist)
	}

	maybeRead := func() {
		_ = try.ByteSlice(os.ReadFile(`non-existent-file`))
		fmt.Println(`file exists`)
	}

	try.Ignoring(isErrNoFile, maybeRead)
	// Output:
}

func ExampleIgnore() {
	isErrNoFile := func(err error) bool {
		return errors.Is(err, os.ErrNotExist)
	}

	someFunc := func() {
		defer try.Ignore(isErrNoFile)
		_ = try.ByteSlice(os.ReadFile(`non-existent-file`))
		fmt.Println(`file exists`)
	}

	someFunc()
	// Output:
}

func ExampleTrans() {
	type ErrPub struct{ error }

	toErrPub := func(err error) error {
		if err != nil {
			return ErrPub{err}
		}
		return nil
	}

	someFunc := func() {
		defer try.Trans(toErrPub)
		panic(errors.New(`failure`))
	}

	err := try.Catch(someFunc)
	fmt.Println(errors.As(err, new(ErrPub)))
	// Output:
	// true
}

func ExampleWithTrans() {
	type ErrPub struct{ error }

	toErrPub := func(err error) error {
		if err != nil {
			return ErrPub{err}
		}
		return nil
	}

	someFunc := func() {
		panic(errors.New(`failure`))
	}

	err := try.Catch(func() {
		try.WithTrans(toErrPub, someFunc)
	})
	fmt.Println(errors.As(err, new(ErrPub)))
	// Output:
	// true
}

func ExampleOk() {
	logOk := func() { fmt.Println(`no panic`) }

	funcOk := func() {
		defer try.Ok(logOk)
		fmt.Println(`not panicking`)
	}

	funcFail := func() {
		defer try.Ok(logOk)
		fmt.Println(`panicking`)
		panic(errors.New(`failure`))
	}

	funcOk()
	_ = try.Catch(funcFail)
	// Output:
	// not panicking
	// no panic
	// panicking
}

func ExampleFail() {
	logErr := func(err error) { fmt.Println(`caught:`, err) }

	funcOk := func() {
		defer try.Fail(logErr)
		fmt.Println(`not panicking`)
	}

	funcFail := func() {
		defer try.Fail(logErr)
		fmt.Println(`panicking`)
		panic(errors.New(`failure`))
	}

	funcOk()
	_ = try.Catch(funcFail)
	// Output:
	// not panicking
	// panicking
	// caught: failure
}

func ExampleInterface() {
	someFunc := func() (interface{}, error) { return "val", nil }
	fmt.Println(try.Interface(someFunc()))
	// Output:
	// val
}
func ExampleBool() {
	someFunc := func() (bool, error) { return true, nil }
	fmt.Println(try.Bool(someFunc()))
	// Output:
	// true
}

func ExampleUint8() {
	someFunc := func() (uint8, error) { return 255, nil }
	fmt.Println(try.Uint8(someFunc()))
	// Output:
	// 255
}

func ExampleUint16() {
	someFunc := func() (uint16, error) { return 255, nil }
	fmt.Println(try.Uint16(someFunc()))
	// Output:
	// 255
}

func ExampleUint32() {
	someFunc := func() (uint32, error) { return 255, nil }
	fmt.Println(try.Uint32(someFunc()))
	// Output:
	// 255
}

func ExampleUint64() {
	someFunc := func() (uint64, error) { return 255, nil }
	fmt.Println(try.Uint64(someFunc()))
	// Output:
	// 255
}

func ExampleByte() {
	someFunc := func() (byte, error) { return 255, nil }
	fmt.Println(try.Byte(someFunc()))
	// Output:
	// 255
}

func ExampleInt8() {
	someFunc := func() (int8, error) { return -127, nil }
	fmt.Println(try.Int8(someFunc()))
	// Output:
	// -127
}

func ExampleInt16() {
	someFunc := func() (int16, error) { return 255, nil }
	fmt.Println(try.Int16(someFunc()))
	// Output:
	// 255
}

func ExampleRune() {
	someFunc := func() (rune, error) { return 255, nil }
	fmt.Println(try.Rune(someFunc()))
	// Output:
	// 255
}

func ExampleInt32() {
	someFunc := func() (int32, error) { return 255, nil }
	fmt.Println(try.Int32(someFunc()))
	// Output:
	// 255
}

func ExampleInt64() {
	someFunc := func() (int64, error) { return 255, nil }
	fmt.Println(try.Int64(someFunc()))
	// Output:
	// 255
}

func ExampleFloat32() {
	someFunc := func() (float32, error) { return 255, nil }
	fmt.Println(try.Float32(someFunc()))
	// Output:
	// 255
}

func ExampleFloat64() {
	someFunc := func() (float64, error) { return 255, nil }
	fmt.Println(try.Float64(someFunc()))
	// Output:
	// 255
}

func ExampleComplex64() {
	someFunc := func() (complex64, error) { return 255, nil }
	fmt.Println(try.Complex64(someFunc()))
	// Output:
	// (255+0i)
}

func ExampleComplex128() {
	someFunc := func() (complex128, error) { return 255, nil }
	fmt.Println(try.Complex128(someFunc()))
	// Output:
	// (255+0i)
}

func ExampleString() {
	someFunc := func() (string, error) { return "str", nil }
	fmt.Println(try.String(someFunc()))
	// Output:
	// str
}

func ExampleInt() {
	someFunc := func() (int, error) { return 255, nil }
	fmt.Println(try.Int(someFunc()))
	// Output:
	// 255
}

func ExampleUint() {
	someFunc := func() (uint, error) { return 255, nil }
	fmt.Println(try.Uint(someFunc()))
	// Output:
	// 255
}

func ExampleUintptr() {
	someFunc := func() (uintptr, error) { return 255, nil }
	fmt.Println(try.Uintptr(someFunc()))
	// Output:
	// 255
}

func ExampleInterfaceSlice() {
	someFunc := func() ([]interface{}, error) { return []interface{}{"val"}, nil }
	fmt.Println(try.InterfaceSlice(someFunc()))
	// Output:
	// [val]
}

func ExampleBoolSlice() {
	someFunc := func() ([]bool, error) { return []bool{true}, nil }
	fmt.Println(try.BoolSlice(someFunc()))
	// Output:
	// [true]
}

func ExampleUint8Slice() {
	someFunc := func() ([]uint8, error) { return []uint8{255}, nil }
	fmt.Println(try.Uint8Slice(someFunc()))
	// Output:
	// [255]
}

func ExampleUint16Slice() {
	someFunc := func() ([]uint16, error) { return []uint16{255}, nil }
	fmt.Println(try.Uint16Slice(someFunc()))
	// Output:
	// [255]
}

func ExampleUint32Slice() {
	someFunc := func() ([]uint32, error) { return []uint32{255}, nil }
	fmt.Println(try.Uint32Slice(someFunc()))
	// Output:
	// [255]
}

func ExampleUint64Slice() {
	someFunc := func() ([]uint64, error) { return []uint64{255}, nil }
	fmt.Println(try.Uint64Slice(someFunc()))
	// Output:
	// [255]
}

func ExampleByteSlice() {
	someFunc := func() ([]byte, error) { return []byte{255}, nil }
	fmt.Println(try.ByteSlice(someFunc()))
	// Output:
	// [255]
}

func ExampleInt8Slice() {
	someFunc := func() ([]int8, error) { return []int8{-127}, nil }
	fmt.Println(try.Int8Slice(someFunc()))
	// Output:
	// [-127]
}

func ExampleInt16Slice() {
	someFunc := func() ([]int16, error) { return []int16{255}, nil }
	fmt.Println(try.Int16Slice(someFunc()))
	// Output:
	// [255]
}

func ExampleRuneSlice() {
	someFunc := func() ([]rune, error) { return []rune{255}, nil }
	fmt.Println(try.RuneSlice(someFunc()))
	// Output:
	// [255]
}

func ExampleInt32Slice() {
	someFunc := func() ([]int32, error) { return []int32{255}, nil }
	fmt.Println(try.Int32Slice(someFunc()))
	// Output:
	// [255]
}

func ExampleInt64Slice() {
	someFunc := func() ([]int64, error) { return []int64{255}, nil }
	fmt.Println(try.Int64Slice(someFunc()))
	// Output:
	// [255]
}

func ExampleFloat32Slice() {
	someFunc := func() ([]float32, error) { return []float32{255}, nil }
	fmt.Println(try.Float32Slice(someFunc()))
	// Output:
	// [255]
}

func ExampleFloat64Slice() {
	someFunc := func() ([]float64, error) { return []float64{255}, nil }
	fmt.Println(try.Float64Slice(someFunc()))
	// Output:
	// [255]
}

func ExampleComplex64Slice() {
	someFunc := func() ([]complex64, error) { return []complex64{255}, nil }
	fmt.Println(try.Complex64Slice(someFunc()))
	// Output:
	// [(255+0i)]
}

func ExampleComplex128Slice() {
	someFunc := func() ([]complex128, error) { return []complex128{255}, nil }
	fmt.Println(try.Complex128Slice(someFunc()))
	// Output:
	// [(255+0i)]
}

func ExampleStringSlice() {
	someFunc := func() ([]string, error) { return []string{"val"}, nil }
	fmt.Println(try.StringSlice(someFunc()))
	// Output:
	// [val]
}

func ExampleIntSlice() {
	someFunc := func() ([]int, error) { return []int{255}, nil }
	fmt.Println(try.IntSlice(someFunc()))
	// Output:
	// [255]
}

func ExampleUintSlice() {
	someFunc := func() ([]uint, error) { return []uint{255}, nil }
	fmt.Println(try.UintSlice(someFunc()))
	// Output:
	// [255]
}

func ExampleUintptrSlice() {
	someFunc := func() ([]uintptr, error) { return []uintptr{255}, nil }
	fmt.Println(try.UintptrSlice(someFunc()))
	// Output:
	// [255]
}
