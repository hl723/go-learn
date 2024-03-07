package reverse

import "strconv"

// Int reversal the digits in the integer i
func Int(i int) int {
	// convert i to str, reverse the str with String func, convert back to int
	i, _ = strconv.Atoi(String(strconv.Itoa(i)))
	return i
}
