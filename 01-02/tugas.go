package main

import "fmt"

func divisible(n, x, y int) bool {
	if n%x == 0 && n%y == 0 {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println(divisible(3, 1, 3))
	fmt.Println(divisible(12, 2, 6))
	fmt.Println(divisible(100, 5, 3))
	fmt.Println(divisible(12, 7, 5))
}
