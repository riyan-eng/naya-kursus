package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	x := 30
	y := 20
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("swap")
	x, y = y, x
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Println("======")

	// var name string
	// fmt.Println("Enter your name: ")
	// fmt.Scan(&name)
	// fmt.Println("your name is", name)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("my name is: ", scanner.Text())
}
