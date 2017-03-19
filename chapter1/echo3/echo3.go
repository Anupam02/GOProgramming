//echo program from join , so everytime we don't have to create a new string
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println(strings.Join(os.Args[1:]," "))
}
