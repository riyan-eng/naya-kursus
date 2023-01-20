package main

import (
	"fmt"
	"strings"
)

func Hero(bullets, dragons int) bool {
	if count := bullets / dragons; count < 2 {
		return false
	} else {
		return true
	}
}

func ReverseWords(str string) string {
	word := strings.Split(str, " ")
	var newWord []string
	for i := len(word) - 1; i >= 0; i-- {
		newWord = append(newWord, word[i])
	}

	return strings.Join(newWord, " ")
}

func main() {
	fmt.Println(Hero(10, 5))
	fmt.Println(Hero(7, 4))
	fmt.Println(Hero(4, 5))
	fmt.Println(Hero(100, 40))
	fmt.Println(Hero(1500, 751))
	fmt.Println(Hero(0, 1))

	fmt.Println()
	fmt.Println(ReverseWords("hello world!"))
	fmt.Println(ReverseWords("yoda doesn't speak like this"))
	fmt.Println(ReverseWords("foobar"))
	fmt.Println(ReverseWords("kata editor"))
	fmt.Println(ReverseWords("row row row your boat"))
}
