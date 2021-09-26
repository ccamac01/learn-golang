// Echgo prints its command-line arguments (both index and value)
// example usage: 'go run main.go hello world'
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for index, arg := range os.Args[1:] {
		ln := strconv.Itoa(index) + ": " + arg
		fmt.Println(ln)
	}
}
