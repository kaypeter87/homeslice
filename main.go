package main

import (
    "fmt"
    "os"
) 

func main() {
    fmt.Println("I am alive!")
}

func createDotfile(path string) *os.File {
    fmt.Println("Creating dotfile...")
    f, err := os.Create(path)
    if err != nil {
        panic(err)
    }
    return f
}
