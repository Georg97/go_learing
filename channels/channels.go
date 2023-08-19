package main

import "fmt"

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

	go customlul(val)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("val: %v", <-val)
		}
	}()
	fibonacci(c, quit)
}

