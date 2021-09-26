// Dup prints the count and text of lines that appear more than once
// from a list of named files, along with the list of filenames in which the dups occurred
// example usage: 'go run main.go test0.txt test1.txt'
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// DupLineFileTracker tracks the names of files that contain the same line along with the count of total duplicates
type DupLineFileTracker struct {
	filenames map[string]string
	count     int
}

func main() {
	counts := make(map[string]*DupLineFileTracker)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			fileTracker, ok := counts[line]
			if !ok {
				fileTracker = &DupLineFileTracker{
					filenames: make(map[string]string),
					count: 0,
				}
				counts[line] = fileTracker
			}
			fileTracker.count++
			fileTracker.filenames[filename] = filename
		}
	}

	for line, fileTracker := range counts {
		if fileTracker.count > 1 {
			fmt.Printf("%d\t%s\n", fileTracker.count, line)
			names := make([]string, 0, len(fileTracker.filenames))
			for k := range fileTracker.filenames {
				names = append(names, k)
			}
			listFiles := "Occurs in the following files: " + strings.Join(names, ", ")
			fmt.Println(listFiles)
		}
	}
}
