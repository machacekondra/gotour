package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev1, prev2, result := 0, 1, 0

	return func() int {
		prev1 = prev2
		prev2 = result
		result = prev1 + prev2
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

