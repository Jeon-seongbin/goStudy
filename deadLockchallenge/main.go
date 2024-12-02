package main

import "fmt"

func main() {
	c := make(chan int)
	go chan1(c)
	for c1 := range c {
		fmt.Println(c1)
	}

}

func chan1(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
