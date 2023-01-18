package main

import (
	"fmt"
	"math"
)

func MakeNegative(x int) int {
	if x > 0 {
		return x * -1
	}
	return x
}

func MakeNegative2(x int) int {
	return int(math.Abs(float64(x))) * -1
}

func PositiveSum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		if num > 0 {
			sum += num
		}
	}
	return sum
}

func main() {
	fmt.Println("--- return negative ---")

	fmt.Println(MakeNegative(1))  // return -1
	fmt.Println(MakeNegative(-5)) // return -5
	fmt.Println(MakeNegative(0))  // return 0

	fmt.Println(MakeNegative2(1))  // return -1
	fmt.Println(MakeNegative2(-5)) // return -5
	fmt.Println(MakeNegative2(0))  // return 0

	fmt.Println()

	fmt.Println("--- sum of positive ---")
	array := []int{1, -4, 7, 12}
	fmt.Println(array)
	fmt.Println(PositiveSum(array))
}
