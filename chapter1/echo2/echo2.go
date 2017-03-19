// echo2 prints its command line arguments
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg //everytime a new string s is made(immutable) and garbage collected
        sep = " "
    }
    fmt.Println(s)
}
