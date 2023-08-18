package main

import (
    "fmt"
    "example/list/listobj"
)

func main() {
    list, _ := listobj.NewList(3)
    list.Add(4)
    list.Add(6)
    list.Add(8)
    fmt.Println(list)
    targetIndex := 6
    fmt.Printf("Index of %v is: %v\n", targetIndex, list.GetIndexOf(targetIndex))
    fmt.Println("Element at index: ", list.Get(3))
}
