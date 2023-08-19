package main

import (
	"fmt"
	"sync"
	// "time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func customlul(val chan int) {
	x := 1
	for {
		val <- x
		x *= 2
	}
}
func main() {
	c := make(chan int)
	val := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

    var wg sync.WaitGroup
    wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("val: %v\n", <-val) // Added a newline for clarity
		}
        wg.Done()
	}()

	fibonacci(c, quit)
    go customlul(val) // Start this goroutine first
	// time.Sleep(50 * time.Millisecond)
    wg.Wait()
}
