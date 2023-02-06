package main

import "fmt"

func luas(alas int, tinggi int, result *float32) {
	*result = float32(alas) * float32(tinggi) / 2
	// fmt.Println(*result)
}

func containsDuplicate(nums []int) bool {
	apa := map[int]int{}
	for _, val := range nums {
		apa[val] = val
	}

	if len(apa) == len(nums) {
		return false
	}
	return true
}

func main() {
	// alas := 5
	// tinggi := 7
	// var result float32 = 0

	// luas(alas, tinggi, &result)

	// fmt.Println(result)

	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))

}
