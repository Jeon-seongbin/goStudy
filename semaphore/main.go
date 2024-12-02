package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	semaphore := make(chan int, 10)
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			semaphore <- id
			time.Sleep(1 * time.Second)

			fmt.Printf("gorutine %d released the semapore \n", <-semaphore)
			wg.Done()
		}(i)
		fmt.Println(i)
	}
	wg.Wait()
	fmt.Println("all gorutine done")
}
