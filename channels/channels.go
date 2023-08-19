package main

import (
	"fmt"
	// "time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int)
	go fibonacci(50, c)
    for {
        i, going := <- c
        if !going {
            break
        }
        fmt.Println(i)
    }
}
