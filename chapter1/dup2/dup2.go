// Dup2 prints and counts the text of lines that appears more than once
// in the input , it reads from stdin for a list of files
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int, 0) // here it is explicitly mentioned for 0
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// ignoring potential errors from input.Err
}
