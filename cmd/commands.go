package cmd

import (
	"log"
	"strconv"

	d "github.com/irevenko/tsukae/draw"
	"github.com/spf13/cobra"
)

var (
	commandsNumber = 7
)

var PieChartFlag bool
var BarChartFlag bool
var ListFlag bool
var ShellHistPath string

var Zsh = &cobra.Command{
	Use:   "zsh",
	Short: "Visualizes the zsh shell commands usage",
	Long:  `tsukae zsh <COMMANDS_NUMBER>`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			num, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal("<COMMANDS_NUMBER> must be an Integer")
			}

			if num < 1 || num > 15 {
				log.Fatal("<COMMANDS_NUMBER> must be between 1-15")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			num, _ := strconv.Atoi(args[0])
			names, occurrences := GetShellCommandsUsage("zsh", num)

			hasFlag := checkForFlags(names, occurrences, num, "zsh")
			if hasFlag == true {
				return
			}

			d.RenderTui(names, occurrences, "zsh")
		} else {
			names, occurrences := GetShellCommandsUsage("zsh", commandsNumber)

			flag := checkForFlags(names, occurrences, commandsNumber, "zsh")
			if flag == true {
				return
			}

			d.RenderTui(names, occurrences, "zsh")
		}
	},
}

var Bash = &cobra.Command{
	Use:   "bash",
	Short: "Visualizes the bash shell commands usage",
	Long:  `tsukae bash <COMMANDS_NUMBER>`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			num, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal("<COMMANDS_NUMBER> must be an Integer")
			}

			if num < 1 || num > 15 {
				log.Fatal("<COMMANDS_NUMBER> must be between 1-15")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			num, _ := strconv.Atoi(args[0])
			names, occurrences := GetShellCommandsUsage("bash", num)

			hasFlag := checkForFlags(names, occurrences, num, "bash")
			if hasFlag == true {
				return
			}

			d.RenderTui(names, occurrences, "bash")
		} else {
			names, occurrences := GetShellCommandsUsage("bash", commandsNumber)

			hasFlag := checkForFlags(names, occurrences, commandsNumber, "bash")
			if hasFlag == true {
				return
			}

			d.RenderTui(names, occurrences, "bash")
		}
	},
}

var Fish = &cobra.Command{
	Use:   "fish",
	Short: "Visualizes the fish shell commands usage",
	Long:  `tsukae fish <COMMANDS_NUMBER>`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			num, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal("<COMMANDS_NUMBER> must be an Integer")
			}

			if num < 1 || num > 15 {
				log.Fatal("<COMMANDS_NUMBER> must be between 1-15")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			num, _ := strconv.Atoi(args[0])
			names, occurrences := GetShellCommandsUsage("fish", num)

			hasFlag := checkForFlags(names, occurrences, num, "fish")
			if hasFlag == true {
				return
			}

			d.RenderTui(names, occurrences, "fish")
		} else {
			names, occurrences := GetShellCommandsUsage("fish", commandsNumber)

			hasFlag := checkForFlags(names, occurrences, commandsNumber, "fish")
			if hasFlag == true {
				return
			}

			d.RenderTui(names, occurrences, "fish")
		}
	},
}

func AddCommands() {
	RootCmd.PersistentFlags().BoolVarP(&PieChartFlag, "piechart", "p", false, "Draw only PieChart")
	RootCmd.PersistentFlags().BoolVarP(&BarChartFlag, "barchart", "b", false, "Draw only BarChart")
	RootCmd.PersistentFlags().BoolVarP(&ListFlag, "list", "l", false, "Draw only List")

	RootCmd.PersistentFlags().StringVarP(&ShellHistPath, "shell-path", "s", "", "Manualy specify the history file")

	RootCmd.AddCommand(Zsh)
	RootCmd.AddCommand(Bash)
	RootCmd.AddCommand(Fish)
}

func checkForFlags(names []string, occurrences []float64, commandsNumber int, shell string) (hasFlag bool) {
	if PieChartFlag == true {
		d.RenderPieChart(names, occurrences, shell, commandsNumber)
		return true
	}
	if BarChartFlag == true {
		d.RenderBarChart(names, occurrences, shell, commandsNumber)
		return true
	}

	if ListFlag == true {
		d.RenderList(names, occurrences, shell, commandsNumber)
		return true
	}

	return false
}
