package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(doWork)

}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func doWork() {
	for i := 0; i < 2000; i++ {
		fmt.Println(i)
	}
}
