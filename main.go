package main

import (
    "fmt"
    "os"
    "log"

    "github.com/urfave/cli/v2"
) 

func main() {
    homeslice := &cli.App{
        Name: "homeslice",
        Usage: "A CLI tool that manages your dotfiles with finesse.",
        Action: func(c *cli.Context) error {
            fmt.Println("I am still a baby!")
            return nil
        },
    }

    err := homeslice.Run(os.Args)
        if err != nil {
            log.Fatal(err)
        }
}

// Placeholder functions for later
func createDotfile(path string) *os.File {
    fmt.Println("Creating dotfile...")
    f, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }
    return f
}

func parseDotfile(path string) {
    fmt.Println("Parsing dotfile...")
}

func storeDotfile(path string) {
    fmt.Println("Saving dotfile...")
}
