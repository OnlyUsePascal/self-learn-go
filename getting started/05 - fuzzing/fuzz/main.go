package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error){
	// b := []byte(s)
	fmt.Printf("input: %q\n", s)
	slices := []rune(s)
	fmt.Printf("runes: %q\n", slices)
	
	if (!utf8.ValidString(s)){
		return s, errors.New("String is not UTF-8 !")
	}

	for i, j := 0, len(slices)-1; i < len(slices)/2; i, j = i+1, j-1 {
		slices[i], slices[j] = slices[j], slices[i]
	}
	return string(slices), nil 
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, _ := Reverse(input)
	doubleRev, _ := Reverse(rev)

	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}
