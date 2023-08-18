package main

import (
    "strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

    retMap := make(map[string]int) 
    var strings = strings.Fields(s)

    for _, str := range strings {
        retMap[str] = len(str)
    }
	// return map[string]int{"x": 1}
    return retMap
}

func main() {
	wc.Test(WordCount)
}
