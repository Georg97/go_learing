package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    fmt.Println(dx, dy)

    var pic [][]uint8 = make([][]uint8, dy)
    for i := 0; i < len(pic); i++ {
        pic[i] = make([]uint8, dx)
        for j := 0; j < len(pic[i]); j++ {
            pic[i][j] = uint8((dx+dy)/2)
        }
    }

    return pic
}

func main() {
	pic.Show(Pic)
}
