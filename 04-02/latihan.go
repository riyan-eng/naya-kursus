package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	var lala int
	if n == 0 {
		lala = 1
	} else {
		lala = n & (n - 1)
	}

	fmt.Println(lala)

	if lala == 0 {
		return true
	}
	return false
}

func main() {
	// fmt.Println(isPowerOfTwo(16))
	// fmt.Println(isPowerOfTwo(4))
	// fmt.Println(isPowerOfTwo(0))
	// fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(3))
}

// 1
// 2
// 4
// 8
// 16
// 32
