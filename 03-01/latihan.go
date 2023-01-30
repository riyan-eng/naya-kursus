package main

import (
	"fmt"
	"sort"
	"strings"
)

func OddCount(n int) int {
	//your code here
	var num int
	if n%2 == 0 {
		num = n / 2
	} else {
		num = (n - 1) / 2
	}
	return num
}

func TwoToOne(s1 string, s2 string) string {
	key := make(map[string]bool)
	listString := strings.Split(s1+s2, "")
	sort.Strings(listString)
	// fmt.Println(listString)
	var eek []string
	for _, val := range listString {
		if _, kk := key[val]; !kk {
			key[val] = true
			// fmt.Println(val)
			eek = append(eek, val)
		}
	}
	// fmt.Println(eek)

	return strings.Join(eek, "")
}

func main() {
	// fmt.Println(OddCount(15))
	// fmt.Println(OddCount(15023))

	a := "xyaabbbccccdefww"
	b := "xxxxyyyyabklmopq"
	fmt.Println(TwoToOne(a, b))
}
