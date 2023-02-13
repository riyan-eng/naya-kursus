package main

import (
	"fmt"
)

func solution(str, ending string) bool {
	// if len(str) == 0 {
	// 	str = " "
	// }
	// if len(ending) == 0 {
	// 	ending = " "
	// }
	lenge := len(str) - len(ending)
	if lenge < 0 {
		return false
	}
	// fmt.Println(lenge)
	// b := str[lenge:]
	// fmt.Println("'", b, "'")

	return str[lenge:] == ending
	// return true
}

func main() {
	fmt.Println(solution("abscsdggff", "bcs"))
	fmt.Println(solution("abscsdggff", "gff"))
	fmt.Println(solution("", ""))
	fmt.Println(solution(" ", ""))
	// fmt.Println(solution("abc", "c"))
	// fmt.Println(solution("banana", "ana"))
	// fmt.Println(solution("a", "z"))
	// fmt.Println(solution("", "t"))
	// fmt.Println(solution("$a = $b + 1", "+1"))
	// fmt.Println(solution("    ", "   "))
	// fmt.Println(solution("1oo", "100"))
}
