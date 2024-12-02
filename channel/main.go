package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()
	fmt.Println(<-ch)

	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	fmt.Println("--")
	ch2 := make(chan int)
	go inputChan(ch2)
	outputChan(ch2)

	fmt.Println("--")
	ch3 := make(chan int)
	var wg3 sync.WaitGroup

	wg3.Add(2)
	go inputChan1(ch3, &wg3)

	go func(ch3 <-chan int) {
		for v := range ch3 {
			fmt.Println(v)
		}
		wg3.Done()
	}(ch3)
	wg3.Wait()

	fmt.Println("--select-")

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)
	receive(eve, odd, quit)
	// send
	fmt.Println("--select-")

}

func inputChan(c chan<- int) {
	c <- 1
}

func outputChan(c <-chan int) {
	fmt.Println(<-c)
}

func inputChan1(c chan<- int, wg3 *sync.WaitGroup) {

	for i := 0; i < 100; i++ {
		c <- i
	}
	wg3.Done()
	close(c)

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// close(e)
	// close(o)
	q <- 0
	// close(q)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the eve channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel", v)
		case v := <-q:
			fmt.Println("from the quit channel", v)
			return
		}
	}
}
