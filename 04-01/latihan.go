package main

import (
	"fmt"
	"strconv"
	"strings"
)

func plusOne(digits []int) []int {
	// lastIndex := len(digits) - 1
	// last := digits[lastIndex]
	// last++
	// digits[lastIndex] = last
	// if digits[lastIndex] >= 10 {
	// 	digits[lastIndex] = 1
	// 	digits = append(digits, 0)
	// }

	var listString []string
	for _, val := range digits {
		valString := strconv.Itoa(val)
		listString = append(listString, valString)
	}

	stringNum := strings.Join(listString, "")
	fmt.Println(stringNum)
	// mm, _ := strconv.ParseUint(stringNum, 10, 64)
	// fmt.Println(mm)

	number, _ := strconv.Atoi(stringNum)
	fmt.Println(number)
	number++

	var mama []int
	for _, val := range strconv.Itoa(number) {
		nn := string(val)
		num, _ := strconv.Atoi(nn)
		mama = append(mama, num)
	}

	return mama
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9}))
	fmt.Println(plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
}
