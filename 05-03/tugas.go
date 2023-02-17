package main

import "fmt"

func BreakChocolate(n, m int) int {
	// math goes here
	if n < 1 || m < 1 {
		return 0
	}
	return n*m - 1
}

func main() {
	fmt.Println(BreakChocolate(5, 5))
	// BreakChocolate(1, 1)
	fmt.Println(BreakChocolate(1, 1))
}
