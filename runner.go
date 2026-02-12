// file1.go

package main

import (
    "fmt"
    "github.com/Deep-Dey1/go_test"
)

// Exported variable
var ExportedVariable = "Hello, World!"

func main() {
    // Accessing exported identifier in the same file
    fmt.Println(ExportedVariable)

    // Accessing exported identifier from another package
    fmt.Println(test.AnotherExportedVariable)
}
