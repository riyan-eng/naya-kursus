package main

import "fmt"

func MxDifLg(a1 []string, a2 []string) int {
	// your code
	terbaru := append(a1, a2...)
	fmt.Println(terbaru)
	return 1
}

func main() {
	a1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	a2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	MxDifLg(a1, a2)
}
