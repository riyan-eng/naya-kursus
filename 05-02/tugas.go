package main

import "fmt"

func Number(busStops [][2]int) int {
	// fmt.Println(busStops)
	var sisa int
	for _, val := range busStops {
		sisa += val[0] - val[1]
	}
	return sisa
}

func main() {
	fmt.Println(Number([][2]int{{10, 0}, {3, 5}, {5, 8}}))
}
