// exercises.go Covers Exercises 1.1, 1.2, (TODO 1.3)

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[0:] {
		fmt.Printf("%v: %v\n", i, arg)
	}
}
