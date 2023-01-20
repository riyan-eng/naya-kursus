package main

import "fmt"

func main() {
	fmt.Println("--- if, if-else, if-elseif ---")
	number := 10
	if number == 10 {
		fmt.Println("angka 10")
	} else if number == 11 {
		fmt.Println("angka 11")
	} else {
		fmt.Println("angka lain")
	}

	fmt.Println()

	fmt.Println("--- switch case ---")
	switch number {
	case 10:
		fmt.Println("angka 10")
	case 11:
		fmt.Println("angka 11")
	default:
		fmt.Println("angka lain")
	}

	fmt.Println()

	fmt.Println("--- for ---")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
