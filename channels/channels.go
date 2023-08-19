package main

import (
	"fmt"
	"time"
)

func printTick[T any](c <-chan T) {
	for {
		select {
		case <-c:
			fmt.Println("tick")
		}
	}
}

func main() {
	timer := time.Tick(100 * time.Millisecond)
	printTick(timer)
	fmt.Println("Program ran through")
}
