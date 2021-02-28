package cmd

import (
	"log"
	"strconv"
)

func GetShellCommandsUsage(shell string, commandsNum int) (n []string, occ []float64) {
	var history []string

	if shell == "bash" {
		history = ParseBashHistory()
	}

	if shell == "zsh" {
		history = ParseZshHistory()
	}

	commands := CountCommands(history)
	delete(commands, "")

	names, occurrences := SortCommands(commands)

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
