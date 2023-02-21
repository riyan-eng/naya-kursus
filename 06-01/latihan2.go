package main

import "fmt"

type Tuple struct {
	Char  string
	Count int
}

// type Tungple struct {
// 	Char  string
// 	Count int
// }

func OrderedCount(text string) []Tuple {
	// Implement me! :)

	var tung []Tuple
	// var ting Tuple

	var a = make(map[string]int)
	for _, val := range text {
		a[string(val)]++
	}

	// fmt.Println(b)

	for key, value := range a {
		var ting Tuple
		ting.Char = key
		ting.Count = value
		tung = append(tung, ting)
	}

	// for _, val := range tung {
	// 	fmt.Println(val)
	// }

	// for key, _ := range a {
	// 	fmt.Println(key)
	// 	var ting Tuple
	// 	ting.Char = key
	// 	ting.Count = 0
	// 	// for _, i := range text {
	// 	// 	if key == string(i) {
	// 	// 		ting.Count++
	// 	// 	}
	// 	// }
	// 	tung = append(tung, ting)
	// }

	// for _, val := range text {
	// 	_, is_ada := a[string(val)]

	// 	if is_ada {
	// 		a[string(val)]++
	// 	}
	// }

	// fmt.Println(a)

	// for _, val := range text {
	// 	var ting Tuple
	// 	for _, vaj := range tung {
	// 		if vaj.Char == string(val) {
	// 			ting.Count++
	// 		}
	// 	}

	// 	if ting.Count == 0 {
	// 		tung = append(tung, ting)
	// 	}
	// }
	return tung
}

func main() {
	fmt.Println(OrderedCount("abracadabra"))
	fmt.Println(OrderedCount("Code Wars"))
}
