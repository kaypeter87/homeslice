package main

import (
    "fmt"
    "os"
    "log"

    "github.com/urfave/cli/v2"
) 

var homeslice = cli.NewApp()

func main() {
    info()
    commands()
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

func info() {
    homeslice.Name = "homeslice"
    homeslice.Usage = "A CLI tool and dotfile manager that organizes your dotfiles."
    homeslice.Authors = []*cli.Author{
        &cli.Author{
            Name: "kaypeter87 (Peter Kay)", 
            Email: "contact@kaypeter.com",
        },
    }
    homeslice.Version = "1.0.0"
}

func commands() {
    homeslice.Commands = []*cli.Command{
        &cli.Command{
            Name:    "test",
            Aliases: []string{"t"},
            Usage:   "Test command",
            Action:  func(c *cli.Context) error {
                t := "This is a test command."
                fmt.Println(t)
                return nil
            },
        },
    }
}
