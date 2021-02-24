package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func fetchBashHistory() []string {
	homePath := os.Getenv("HOME")

	rawHistory, err := ioutil.ReadFile(homePath + "/.bash_history")
	if err != nil {
		fmt.Println("Can't open .bash_history in: " + homePath)
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

func main() {
	commandsHist := fetchBashHistory()
	fmt.Println(commandsHist)
}
