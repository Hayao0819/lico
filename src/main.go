package main

import (
	"os"

	"github.com/Hayao0819/lico/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
