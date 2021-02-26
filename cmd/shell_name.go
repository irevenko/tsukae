package cmd

import (
	s "../shell"
)

func GetShellCommandsUsage(shell string, commandsNum int) (n []string, occ []float64) {
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

	return names, occurrences
}