package main

import "fmt"

func main() {
	var faktorial int = 4

	// inisiasi nilai total
	var total int = 1

	for i := 1; i <= faktorial; i++ {
		total = total * i

		fmt.Println("--- faktorialnya ---")
		fmt.Println(i)
		fmt.Println("--- jumlahnya ---")
		fmt.Println(total)
	}
	fmt.Println("--- hasil akhir ---")
	fmt.Println(total)
}
