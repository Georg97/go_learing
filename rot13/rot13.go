package main

import (
	// "fmt"
	"io"
	"os"
	"strings"
)

func rot13Apply(char byte) byte {
    if char == ' ' {
        return char
    }
	newChar := char + 13
    switch {
    case char < 'z' && char > 'a':
        if newChar > 'z' {
            newChar = 'a' + (newChar - 'z' - 1)
        }
    case char < 'Z' && char > 'A':
        if newChar > 'Z' {
            newChar = 'A' + (newChar - 'Z' - 1)
        }
    default:
        return char
    }
	return newChar
}

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(out []byte) (int, error) {
	n, err := reader.r.Read(out)
	for i := 0; i < n; i++ {
		out[i] = rot13Apply(out[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

