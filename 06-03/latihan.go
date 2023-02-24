package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	go calculateArea(ch)

	wg.Add(1)
	go getArea(&wg, ch)

	wg.Wait()
}

func calculateArea(ch chan<- int) {
	ch <- 6 * 2
	close(ch)
}

func getArea(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	area := <-ch
	fmt.Println(area)
}
