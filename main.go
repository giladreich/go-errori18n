package main

import (
	"fmt"

	. "github.com/giladreich/go-errori18n/errori18n"
)

func Divide(l int, r int) (int, *Error) {
	if r == 0 {
		return 0, Errorf(ERR_DIVIDE_BY_ZERO, l)
	}
	return l / r, nil
}

func print(err *Error) {
	fmt.Println("\n------")
	fmt.Printf("%#v\n", err.ErrorCode)
	// Implicit calls to defined `*.String()` functions:
	fmt.Printf("%s\n", err.ErrorCode)
	fmt.Printf("%s\n", err)
}

func main() {
	simple := Errorf(ERR_SIMPLE)
	variadic := Errorf(ERR_VARIADIC,
		1.23, 45, "hello", []int{1, 2, 3}, struct{ v1, v2 string }{"hello", "world"})

	print(simple)
	print(variadic)

	SetLanguage(LANG_DE)
	print(simple)
	print(variadic)

	if _, err := Divide(5, 0); err != nil {
		panic(err)
	}
}
