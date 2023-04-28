package main

import (
	"github.com/Hayao0819/lico/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(os.Stdin, os.Stdout,  os.Args...); err != nil {
		os.Exit(1)
	}
}
