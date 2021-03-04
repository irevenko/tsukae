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
			blackListFile, err := os.Create(homePath + blackListPath)
			if err != nil {
				log.Fatal("Unable to create blacklist file: ", err)
			}
			defer blackListFile.Close()

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
