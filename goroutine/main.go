package main

import (
	"fmt"
	"runtime"
	"sync"
)

type human interface {
	speak()
}

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("hello")
}

func saySometing(h human) {
	h.speak()
}

func main() {
	// 1
	var wg sync.WaitGroup

	count := 2
	wg.Add(count)

	for i := 0; i < count; i++ {

		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end")

	p1 := person{
		first: "james",
	}

	saySometing(&p1)

	fmt.Println("--")

	gs := 100
	incrementer := 0

	var wg1 sync.WaitGroup
	var mu sync.Mutex
	wg1.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			mu.Unlock()
			wg1.Done()
		}()
	}

	wg1.Wait()
	fmt.Println(incrementer)

	fmt.Println("--")

	var wg2 sync.WaitGroup
	incrementer = 0
	wg2.Add(gs)
	for i := 0; i < 100; i++ {
		go func() {

			wg2.Done()
		}()
	}

	wg2.Wait()
}
