package main

import (
	"fmt"
	"strings"
)

func FindShort(s string) int {
	// your code
	word := strings.Split(s, " ")
	amount := len(word[0])
	for _, val := range word {
		if len(val) < amount {
			amount = len(val)
		}
	}
	return amount
}

func main() {
	fmt.Println(FindShort("bitcoin take over the world maybe who knows perhaps"))
	fmt.Println(FindShort("turns out random test cases are easier than writing out basic ones"))
	fmt.Println(FindShort("Let's travel abroad shall we"))
}
