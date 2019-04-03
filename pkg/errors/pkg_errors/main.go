package main

import (
	"errors"
	"fmt"

	pkgErrors "github.com/pkg/errors"
)

func main() {
	err := fn()
	fmt.Println(err)
	fmt.Println(pkgErrors.Cause(err))

	wrapErr()

	fErr()

	newErr()

	messageErr()

	messagefErr()
}

func fn() error {
	return errors.New("test cause of error of pkgErrors")
}

func wrapErr() {
	cause := pkgErrors.New("whoops")
	err := pkgErrors.Wrap(cause, "oh noes")

	fmt.Printf("%v", err)
}

func fErr() {
	err := pkgErrors.Errorf("whoops: %s", "foo")
	fmt.Printf("%+v", err)
}

func newErr() {
	err := pkgErrors.New("whoops")
	fmt.Println(err)

	fmt.Printf("%+v", err)
}

func messageErr() {
	cause := pkgErrors.New("whoops")
	err := pkgErrors.WithMessage(cause, "oh noes")
	fmt.Println(err)
}

func messagefErr() {
	cause := pkgErrors.New("whoops")
	err := pkgErrors.WithMessagef(cause, "%s", "oh noes")
	fmt.Println(err)
}

func stackErr() {
	cause := pkgErrors.New("whoops")
	err := pkgErrors.WithStack(cause)
	fmt.Println(err)
	fmt.Printf("%+v", err)
}
