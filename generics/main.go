package main

import "fmt"

type Number interface {
	int64 | float64
}

// sum up all ints values in the dict
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// sum up all float values in the dict
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// A generic function to input a map
// and handle both ints and floats via type constraints
func SumGeneric[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Generic function as above with type constraint instead
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	// test map of ints
	ints := make(map[string]int64)
	ints["first"] = 34
	ints["second"] = 12

	// test map of floats
	floats := map[string]float64 {
		"first": 35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: Int: %v and Float: %v\n", 
		SumInts(ints), 
		SumFloats(floats))

	fmt.Printf("Generic     Sums: Int: %v and Float: %v\n",
		SumGeneric[string, int64](ints),
		SumGeneric[string, float64](floats))

	fmt.Printf("Generic (No Type): Int: %v and Float: %v\n",
		SumGeneric(ints),
		SumGeneric(floats))

	fmt.Printf("Generic (Type Interface): Int: %v and Float: %v\n",
		SumGeneric(ints),
		SumGeneric(floats))
}


