package go_bitmap

import "fmt"

func toError(e any) error {
	if err, ok := e.(error); ok {
		return err
	}
	return fmt.Errorf("%v", e)
}

func TryCatch(f func(), catch func(e any)) {
	defer func() {
		if r := recover(); r != nil {
			catch(r)
		}
	}()

	f()
}

func Try(f func()) error {
	var err error
	TryCatch(f, func(e any) {
		err = toError(e)
	})
	return err
}

func ThrowError(e any) {
	if e != nil {
		panic(e)
	}
}

func Throwf(format string, e ...any) {
	s := fmt.Sprintf(format, e...)
	ThrowError(s)
}
