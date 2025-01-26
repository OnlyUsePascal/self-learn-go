package main

import "fmt"

func slices_references() {
	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	names := [4]string{
		"Joun",
		"Jeff",
		"Jing",
		"Jang",
	}

	slices1 := names[0:2]
	slices2 := names[1:3]
	slices3 := names[0:3]
	fmt.Printf("%v, %v, %v\n", slices1, slices2, slices3)

	slices3[1] = "XXX"
	fmt.Printf("%v, %v, %v\n", slices1, slices2, slices3)

}

func slices_size() {
	// concept: length & capacity
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

	vals := []int{1,2,3,4,5,6}
	
	// small len, but more capacity
	slices1 := vals[:2]
	fmt.Printf("%v, len: %v, cap: %v\n", slices1, len(slices1), cap(slices1))
	
	// extend length a bit
	slices2 := vals[:len(slices1) + 1 ]
	fmt.Printf("%v, len: %v, cap: %v\n", slices1, len(slices2), cap(slices2))
	
	// u cannot get a length longer than capacity 
	slices2 = vals[:cap(slices2) + 1]
	fmt.Printf("%v, len: %v, cap: %v\n", slices1, len(slices2), cap(slices2))
}


func main() {
	slices_references()
	
	slices_size()
}
