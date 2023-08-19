package main

import (
	"fmt"
	"slices"
	"sync"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		go Walk(t.Left, ch)
	}
	if t.Right != nil {
		go Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	same := false

	ch1 := make(chan int)
	ch2 := make(chan int)

	var arr1 [10]int
	var arr2 [10]int

	go Walk(t1, ch1)
	go Walk(t2, ch2)

    var syncer sync.WaitGroup
    syncer.Add(2)

	fill := func(arr []int, ch chan int) {
		for i := 0; i < 10; i++ {
			arr[i] = <-ch
		}
		// close(ch)
        syncer.Done()
	}
	go fill(arr1[:], ch1)
	go fill(arr2[:], ch2)

    syncer.Wait()
	// for {
	// 	_, open := <-ch1
	// 	_, open2 := <-ch2
	// 	if !open && !open2 {
	// 		break
	// 	}
	// }

    fmt.Printf("Arrs: %v --- %v\n lens: %v | %v\n", arr1, arr2, len(arr1[:]), len(arr2[:]))
	if len(arr1) == len(arr2) {
		for i := 0; i < len(arr1); i++ {
			if !slices.Contains(arr2[:], arr1[i]) {
				return false
			}
		}
		same = true
	}

	return same
}

func main() {
	// c := make(chan int)
	t1 := tree.New(1)
	t2 := tree.New(1)
	// go Walk(t1, c)
	// for val := range c {
	//     fmt.Println(val)
	// }
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-c)
	// }
	// fmt.Println("Arrays are same: ", Same(t1, t2))
	t3 := tree.New(1)
	t4 := tree.New(2)

	// fmt.Println("Arrays1 are same: ", Same(t1, t2))
	// fmt.Println("Arrays2 are same: ", Same(t3, t4))
	fmt.Printf("Works: %v\n", Same(t1, t2) != Same(t3, t4))
}
