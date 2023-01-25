package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FakeBin(x string) string {
	var data []string
	for _, val := range x {
		num, _ := strconv.Atoi(string(val))
		if int(num) < 5 {
			num = 0
		} else {
			num = 1
		}
		data = append(data, strconv.Itoa(num))

	}
	return strings.Join(data, "")
}

func DNAtoRNA(dna string) string {
	// your code here
	new := strings.Split(dna, "")
	var newstring []string
	for _, val := range new {
		if val == "T" {
			val = "U"
		}
		newstring = append(newstring, val)
	}
	return strings.Join(newstring, "")
}

func main() {
	// fmt.Println(FakeBin("45385593107843568"))
	// fmt.Println(FakeBin("509321967506747"))

	fmt.Println(DNAtoRNA("GCAT"))
	fmt.Println(DNAtoRNA("TGCAT"))
	fmt.Println(DNAtoRNA("GTCAT"))
	fmt.Println(DNAtoRNA("GCATT"))
}
