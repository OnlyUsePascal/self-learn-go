package main

import (
	"fmt"
	"time"
)

func Assert_basic() {
	// A type assertion provides access to an interface value's underlying concrete value.
	// If i does not hold a T, the statement will trigger a panic.
	// If not, ok will be false and t will be the zero value of type T, and no panic occurs.
	fmt.Println("> assert basic")
	var i interface{} = "Hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// panic
	// f := i.(float32)
	// fmt.Println(f)

	// not panic + zero-value of type
	f, ok := i.(float32)
	fmt.Println(f, ok)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string{
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"send nude !!",
	}
}

func Assert_print() {
	// the idea is The fmt package (and many others) look string interface to print values.
	// even error, where fmt expects to implement Error() function 
	fmt.Println("> Assert print")
	if err := run(); err != nil {
		fmt.Println(err)
	}

}
