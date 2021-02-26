package main

import (
	"fmt"
	"os"
	c "./cmd"
)

func main() {
	c.AddCommands()

	if err := c.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}