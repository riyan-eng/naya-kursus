package main

import "fmt"

func IsTriangle(a, b, c int) bool {
	if a+b > c && a+c > b && c+b > a {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsTriangle(1, 2, 3))
	fmt.Println(IsTriangle(1, 1, 1))
	fmt.Println(IsTriangle(3, 2, 3))
}
