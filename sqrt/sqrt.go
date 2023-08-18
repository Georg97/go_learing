package main

import (
	"fmt"
)

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0.0 {
        return 0.0, ErrNegativeSqrt(x)
    }
    z := 1.0
    prev := 0.0

    for i := 0; i < 10; i++ {
        if i > 0 {
            prev = z
        }
        z -= (z * z - x) / (2 * z)
        if i > 0 && (prev + 0.0001) > z {
            fmt.Println("Condition reached, iterated ", i + 1, " times")
            break
        }
    }
    return z, nil
}

func main() {
    result, err := Sqrt(-2)
    if err != nil {
        fmt.Println("Could not create Root lul")
    }
	fmt.Println(result)
}
