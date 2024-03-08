package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := -1
	prev1, prev2 := 1, 1
	
	return func() int {
		i++
		if i == 0 {
			return 0	
		} else if i <= 2 {
			return 1	
		}
		prev1, prev2 = prev2, prev1 + prev2
		return prev2			
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}