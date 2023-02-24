package main

import (
	"fmt"
)

func GroupArray(arr []int, c int) (res [][]int) {
	// Your blaBlabLaaA algorithm
	// var b []int
	// for _, val := range arr {
	// 	// fmt.Println(val)
	// 	a := c - val
	// 	for _, val := range arr {
	// 		if a == val {
	// 			b = append(b, a)
	// 		}
	// 	}
	// }
	// // fmt.Println(b)
	// sort.Ints(b)

	for _, val := range arr {
		// fmt.Println(val)
		for i := 0; i < len(arr); i++ {
			if val+arr[i] == c {
				res = append(res, []int{val, arr[i]})
			}
		}
	}
	return res
}

func main() {
	fmt.Println(GroupArray([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println(GroupArray([]int{3, 4, 5}, 6))
	// fmt.Println(GroupArray([]int{3, 3, 3}, 6))
	fmt.Println(GroupArray([]int{4, 5}, 8))
	fmt.Println(GroupArray([]int{5, 7, 5, 8, 12, 6}, 10))
}
