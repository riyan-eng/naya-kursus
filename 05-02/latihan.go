package main

import (
	"fmt"
	"strconv"
	"strings"
)

func MinMax(arr []int) [2]int {
	var min int = arr[0]
	var max int = arr[0]
	for _, val := range arr {
		if min > val {
			min = val
		}
		if max < val {
			max = val
		}
	}
	return [2]int{min, max}
}

func DontGiveMeFive(start int, end int) int {
	var data []int
	for i := start; i <= end; i++ {
		if ah := strconv.Itoa(i); !strings.ContainsAny(ah, "5") {
			data = append(data, i)
		}
	}
	return len(data)
}

func main() {
	// fmt.Println(MinMax([]int{1, 2, 3, 4, 5}))
	// fmt.Println(MinMax([]int{5, 4, 6, 1}))
	// fmt.Println(MinMax([]int{2334454, 5}))

	fmt.Println(DontGiveMeFive(1, 9))
	fmt.Println(DontGiveMeFive(4, 17))
	// DontGiveMeFive(4, 17)
}
