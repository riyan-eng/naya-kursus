package main

import (
	"fmt"
	"runtime"
	"sort"
	"strings"
)

func OddCount(n int) (num int) {
	//your code here
	if n%2 == 0 {
		num = n / 2
	} else {
		num = (n - 1) / 2
	}
	return
}

func TwoToOne(s1 string, s2 string) string {
	key := make(map[string]bool)
	listString := strings.Split(s1+s2, "")
	fmt.Println(listString)
	sort.Strings(listString)
	fmt.Println(listString)
	var eek []string
	for _, val := range listString {
		// fmt.Println(val)
		// fmt.Println(key)
		if _, kk := key[val]; !kk {
			key[val] = true
			// fmt.Println(key)
			// fmt.Println(li, kk)
			eek = append(eek, val)
		}
		fmt.Println(eek)
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

	go func() {
		fmt.Println("l")
	}()

	go func() {
		fmt.Println("l")
	}()

	fmt.Println(runtime.NumGoroutine())
}
