package main

import (
	"fmt"
	"math"
)

// You can declare a method on non-struct types, too.
// You can only declare a method with a receiver whose type is defined in the same package as the method
// You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
type MyFloat float64

func (x MyFloat) RoundUp() int {
	xBare := float64(x)
	return int(math.Round(xBare))
}

func methods_non_struct() {
	fmt.Println("> Methods for non-struct")
	float_ := MyFloat(1.23)
	fmt.Println(float_.RoundUp())
}

//Methods with pointer receivers can modify the value to which the receiver points
// With a value receiver, the Scale method operates on a copy of the original Vertex value. 
type Rectangle struct {
	width, height int 
}

func (r Rectangle) GetArea() int {
	return r.width * r.height
}

// Try removing the * from the declaration
func (r *Rectangle) Update(width, height int) {
	r.width = width
	r.height = height
}

func methods_pointer() {
	fmt.Println("> pointer receiver")
	r := Rectangle{3,4}
	r.Update(5,6)
	fmt.Println(r.GetArea())
}


func main() {
	methods_non_struct()
	methods_pointer()
	
	Inteface_basic()
	Interface_empty()
	
	Assert_basic()
	Assert_print()

	Reader_basic()
	Reader_Wrap()
}