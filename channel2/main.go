package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan int)
	wg.Add(1)

	go sqare(wg, ch)
	ch <- 9
	// close(ch)
	wg.Wait()

}

func sqare(wg *sync.WaitGroup, ch chan int) {
	// for aa := range ch {
	// fmt.Println(aa)
	// }

	fmt.Println(<-ch)
	wg.Done()
}
