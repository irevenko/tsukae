package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "tsukae",
	Short: "Visualize your shell commands usage",
	Long:  `Complete documentation is available at https://github.com/irevenko/tsukae`,
}