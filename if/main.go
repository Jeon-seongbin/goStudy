package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	if z := 2 * rand.Intn(40); z >= 40 {
		fmt.Println("10")
	} else {
		fmt.Println("20")
	}

	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	go func() {
		time.Sleep(d1 * time.Microsecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Microsecond)
		ch2 <- 42
	}()

	select {
	case v1 := <-ch1:
		fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v2)
	}
}
