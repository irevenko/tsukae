package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	homePath     = os.Getenv("HOME")
	histFilePath = os.Getenv("HISTFILE")
	bashFile     = "/.bash_history"
	zshFile      = "/.zsh_history"
)

func setHistFile(histFlag string, shell string) []byte {
	if histFlag != "" {
		fileSlice, err := ioutil.ReadFile(histFlag)
		if err != nil {
			log.Fatal("Can't find or open " + histFlag)
		}
		return fileSlice
	}

	if histFilePath != "" {
		rawHistory, err := ioutil.ReadFile(histFilePath)
		if err != nil {
			log.Fatal("Can't find or open " + histFilePath + "\nSpecify absolute hist file path: -m=/your_dir/your_hist_file")
		}
		return rawHistory
	}

	if shell == "bash" {
		rawHistory, err := ioutil.ReadFile(homePath + bashFile)
		if err != nil {
			log.Fatal("Can't find or open " + bashFile + " in: " + homePath + "\nSpecify absolute hist file path: -m=/your_dir/your_hist_file")
		}
		return rawHistory
	}

	if shell == "zsh" {
		rawHistory, err := ioutil.ReadFile(homePath + zshFile)
		if err != nil {
			log.Fatal("Can't find or open " + zshFile + " in: " + homePath + "\nSpecify absolute hist file path: -m=/your_dir/your_hist_file")
		}
		return rawHistory
	}

	return []byte{}
}

func ParseBashHistory() []string {
	rawHistory := setHistFile(ShellHistPath, "bash")

	historySlice := strings.Split(string(rawHistory), "\n")

	for i, v := range historySlice {
		if strings.HasPrefix(v, "sudo") {
			commandSlice := strings.Split(v, " ")
			if len(commandSlice) == 1 { //checking if there is only sudo command
				historySlice[i] = ""
			} else {
				historySlice[i] = commandSlice[1] //writing command name after sudo
			}
		} else {
			commandSlice := strings.Split(v, " ")
			historySlice[i] = commandSlice[0] //writing command name only
		}
	}

	return historySlice
}

func ParseZshHistory() []string {
	rawHistory := setHistFile(ShellHistPath, "zsh")

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
			if len(commandSlice) == 1 { //checking if there is only sudo command
				historySlice[i] = ""
			} else {
				historySlice[i] = commandSlice[1] //writing command name after sudo
			}
		} else {
			commandSlice := strings.Split(v, " ")
			historySlice[i] = commandSlice[0] //writing command name only
		}
	}

	return historySlice
}
