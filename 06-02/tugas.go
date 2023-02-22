package main

import (
	"fmt"
	"strings"
	"unicode"
)

func solve(str string) string {
	// Your code here and happy coding!
	// var lowerCount int
	var upperCount int
	for _, val := range str {
		if unicode.IsUpper(val) {
			upperCount++
		}
	}

	if upperCount > len(str)/2 {
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)
}

func main() {
	kata := "codE"
	fmt.Println(solve(kata))

	// fmt.Println(lowerCount)
	// fmt.Println(upperCount)
}
