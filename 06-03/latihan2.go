package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	go sendAction(1, ch)

	wg.Add(1)
	go receiveAction(ch, &wg)
	wg.Wait()

}

func sendAction(action int, ch chan<- int) {
	ch <- action
	close(ch)
}

func receiveAction(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	action := <-ch

	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	if action == 0 {
		go actionLike(&wg2)
	} else if action == 1 {
		go actionComment(&wg2)
	} else {
		go error(&wg2)
	}
	wg2.Wait()
}

func actionLike(wg2 *sync.WaitGroup) {
	defer wg2.Done()
	fmt.Println("like")
}

func actionComment(wg2 *sync.WaitGroup) {
	defer wg2.Done()
	fmt.Println("comment")
}

func error(wg2 *sync.WaitGroup) {
	defer wg2.Done()
	fmt.Println("error")
}
