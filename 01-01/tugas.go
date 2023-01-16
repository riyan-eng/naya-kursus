package main

import "fmt"

func main() {
	x := 30
	y := 20
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("swap")
	x, y = y, x
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
