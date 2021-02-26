package main

import (
	"fmt"
	"os"
	c "github.com/irevenko/tsukae/cmd"
)

func main() {
	c.AddCommands()

	if err := c.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}