package main

import "fmt"

type Abser interface{
	Abs() float64
}

type Type1 float64

func (t Type1) Abs() float64{
	return float64(t)
}

type Type2 struct {
	value int
}

func (t *Type2) Abs() float64 {
	return float64(t.value)
}

func Inteface_basic(){
  //There is no explicit declaration of intent, no "implements" keyword.
	fmt.Println("> interface basic")
	var abs Abser

	abs = Type1(1.23)
	fmt.Println(abs.Abs())
	
	// abs = Type2{123} //error
	abs = &Type2{123}
	fmt.Println(abs.Abs())
}

func Interface_empty() {
	// The interface type that specifies zero methods is known as the empty interface:
	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	fmt.Println("> Empty interface")
	var i interface{}
	
	i = "hello"
	describe(i)
	
	i = 1234
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
