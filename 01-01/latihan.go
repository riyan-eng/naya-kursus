package main

import "fmt"

func main() {
	var counter int // deklasrasi
	counter = 1     // inisialisasi

	var counter2 = 2

	counter3 := 3

	fmt.Println(counter)
	fmt.Println(counter2)
	fmt.Println(counter3)

	var (
		a int = 1
		b int = 2
		c int = 3
		d int = 4
	)

	a, b, c, d = 10, 11, 12, 13
	fmt.Println(a, b, c, d)

	const phi = 3.14

	fmt.Println(phi)
}
