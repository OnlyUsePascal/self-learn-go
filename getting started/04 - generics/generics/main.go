package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumInts(m map[string]int64) int64 {
	var sum int64 = 0

	for _, val := range m {
		sum += val
	}

	return sum
}

func SumFloats(m map[string]float64) float64 {
	var sum float64 = 0

	for _, val := range m {
		sum += val
	}

	return sum
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var sum V = 0

	for _, val := range m {
		sum += val
	}

	return sum
}

func main() {
	ints := map[string]int64{
		"first": 123,
		"sec":   234,
	}

	floats := map[string]float64{
		"first": 1.23,
		"sec":   2.34,
	}

	// fmt.Println(SumInts(ints), SumFloats(floats))

	// fmt.Println(SumIntsOrFloats[string, int64](ints),
	// 	SumIntsOrFloats[string, float64](floats))

	fmt.Println(SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}
