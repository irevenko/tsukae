package cmd

import (
	"log"
	"strconv"

	s "github.com/irevenko/tsukae/shell"
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

	if len(names) < commandsNumber {
		log.Fatal("Your history is to small! history file must contain at least " + strconv.Itoa(commandsNumber) + " unique commands")
	}

	if len(names) < commandsNum {
		log.Fatal("Your history is to small! Can't retrive " + strconv.Itoa(commandsNum) + " commands")
	}

	names = names[:commandsNum]
	occurrences = occurrences[:commandsNum]

	return names, occurrences
}
