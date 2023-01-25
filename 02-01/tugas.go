package main

import (
	"fmt"
	"strings"
)

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}

func century(year int) int {
	if year%100 != 0 {
		return (year / 100) + 1
	}
	return year / 100
}

func Accum(s string) string {
	strings.ToLower(s)
	kata := strings.ToLower(s)
	// fmt.Println(kata)
	var katas []string
	for i, val := range kata {
		// fmt.Println(strings.Title(strings.Repeat(string(val), i+1)))
		katas = append(katas, strings.Title(strings.Repeat(string(val), i+1)))
	}
	return strings.Join(katas, "-")
}

func main() {
	// fmt.Println(RemoveChar("eloquent"))
	// fmt.Println(RemoveChar("country"))
	// fmt.Println(RemoveChar("person"))
	// fmt.Println(RemoveChar("place"))

	// fmt.Println(century(1990))
	// fmt.Println(century(1705))
	// fmt.Println(century(1900))
	// fmt.Println(century(1601))
	// fmt.Println(century(2000))
	// fmt.Println(century(89))

	fmt.Println(Accum("ZpglnRxqenU"))
}
