package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int, int, int {
	n1, n2 := 0, 1
	return func() int, int, int {
		temp := n1
		n1 = n2
		n2 = temp + n2
		return temp, n1, n2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
