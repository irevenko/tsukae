package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	blackListPath = "/.config/tsukae/blacklist"
)

func ParseBlackList() []string {
	if _, err := os.Stat(homePath + blackListPath); err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(homePath+"/.config/tsukae", 0755)
			if err != nil {
				log.Fatal("Unable to create tsukae folder in " + homePath + "/.config")
			}

			blackListFile, err := os.OpenFile(homePath+blackListPath, os.O_RDONLY|os.O_CREATE, 0644)
			if err != nil {
				log.Fatal("Unable to create blacklist file in " + homePath + "/.config")
			}
			blackListFile.Close()

			fmt.Println("Created blacklist file in: " + homePath + blackListPath)
			fmt.Println("Put some commands in this file if you want to ignore them (separate each commands by new line)")
		}
	}

	blackList, err := ioutil.ReadFile(homePath + blackListPath)
	if err != nil {
		log.Fatal("Can't read blacklist file in: " + homePath + blackListPath)
	}

	blackListSlice := strings.Split(string(blackList), "\n")

	return blackListSlice
}
