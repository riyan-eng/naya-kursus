package main

import (
	"fmt"
	"math"
	"sort"
)

func Gimme(array [3]int) int {
	cek := array
	sort.Ints(cek[:])
	la := cek[1]
	na := 0
	for _, val := range array {
		if val == la {
			return na
		}
		na++
	}
	return 0
}

func RoundToNext5(n int) int {
	lala := math.Ceil(float64(n) / 5)
	if lala == -0 {
		lala = 0
	}
	return int(lala) * 5
}

func main() {
	// fmt.Println(Gimme([3]int{2, 3, 1}))
	// fmt.Println(Gimme([3]int{5, 10, 14}))
	// fmt.Println(Gimme([3]int{1, 3, 4}))

	fmt.Println(RoundToNext5(2))
	fmt.Println(RoundToNext5(3))
	fmt.Println(RoundToNext5(6))
	fmt.Println(RoundToNext5(11))
	fmt.Println(RoundToNext5(21))
	fmt.Println(RoundToNext5(-5))
	fmt.Println(RoundToNext5(-2))
}
