package main

import (
	s "./shell"
	d "./draw"
)

func getShellCommandsUsage(shell string, commandsNum int) {
	var history []string

	if shell == "bash" {
		history = s.ParseBashHistory()
	}

	if shell == "zsh" {
		history = s.ParseZshHistory()
	}

	commands := s.CountCommands(history)
	delete(commands, "")
	names, occurrences := s.SortCommands(commands)

	names = names[:commandsNum]
	occurrences = occurrences[:commandsNum]

	d.RenderTui(names, occurrences)

}

func main() {
	getShellCommandsUsage("bash", 11)
}