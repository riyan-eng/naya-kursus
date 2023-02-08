package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	var lala int
	if n == 0 {
		lala = 0
	} else {
		lala = n & (n - 1)
	}

	if lala == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(4))
	fmt.Println(isPowerOfTwo(0))
	fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(5))
}

// 1
// 2
// 4
// 8
// 16
// 32
