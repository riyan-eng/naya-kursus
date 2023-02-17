package main

import (
	"fmt"
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

func main() {
	fmt.Println(Gimme([3]int{2, 3, 1}))
	fmt.Println(Gimme([3]int{5, 10, 14}))
	fmt.Println(Gimme([3]int{1, 3, 4}))
}
