package shell

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	homePath = os.Getenv("HOME")
	bashFile = "/.bash_history"
	zshFile  = "/.zsh_history"
)

func ParseBashHistory() []string {
	rawHistory, err := ioutil.ReadFile(homePath + bashFile)
	if err != nil {
		log.Fatal("Can't find or open " + bashFile + " in: " + homePath)
	}

	historySlice := strings.Split(string(rawHistory), "\n")

	for i, v := range historySlice {
		if strings.HasPrefix(v, "sudo") {
			commandSlice := strings.Split(v, " ")
			historySlice[i] = commandSlice[1] //writing command name after sudo
		} else {
			commandSlice := strings.Split(v, " ")
			historySlice[i] = commandSlice[0] //writing command name only
		}
	}

	return historySlice
}

func ParseZshHistory() []string {
	rawHistory, err := ioutil.ReadFile(homePath + zshFile)
	if err != nil {
		log.Fatal("Can't find or open " + zshFile + " in: " + homePath)
	}

	historySlice := strings.Split(string(rawHistory), "\n")

	for i, v := range historySlice {
		if strings.HasPrefix(v, ":") { //check if there is a timestamp and remove it
			if idx := strings.Index(v, ";"); idx >= 0 {
				historySlice[i] = historySlice[i][idx+1:]
			}
		}
	}

	for i, v := range historySlice {
		if strings.HasPrefix(v, "sudo") {
			commandSlice := strings.Split(v, " ")
			historySlice[i] = commandSlice[1] //writing command name after sudo
		} else {
			commandSlice := strings.Split(v, " ")
			historySlice[i] = commandSlice[0] //writing command name only
		}
	}

	return historySlice
}
