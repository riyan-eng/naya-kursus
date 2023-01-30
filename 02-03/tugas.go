package main

import "fmt"

func GetSum(a, b int) int {
	// fmt.Println(a, b)
	if a < b {
		count := 0
		for i := a; i <= b; i++ {
			count += i
		}
		return count
	} else if a > b {
		count := 0
		for i := a; i >= b; i-- {
			count += i
		}
		return count
	}
	return a
}

func main() {
	fmt.Println(GetSum(0, 1))   // 0 1
	fmt.Println(GetSum(1, 2))   // 1 2
	fmt.Println(GetSum(5, -1))  // 5 4 3 2 1 0 -1
	fmt.Println(GetSum(17, 17)) // 17 17
	fmt.Println(GetSum(-7, 13)) // 17 17
}
