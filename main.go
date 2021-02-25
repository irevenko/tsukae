package main

import (
	"fmt"
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	homePath = os.Getenv("HOME")
	bashFile = "/.bash_history"
	zshFile  = "/.zsh_history"
)

func fetchBashHistory() []string {
	rawHistory, err := ioutil.ReadFile(homePath + bashFile)
	if err != nil {
		log.Fatal("Can't open " + bashFile + " in: " + homePath)
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

func fetchZshHistory() []string {
	rawHistory, err := ioutil.ReadFile(homePath + zshFile)
	if err != nil {
		log.Fatal("Can't open " + zshFile + " in: " + homePath)
	}

	historySlice := strings.Split(string(rawHistory), "\n")

	for i, v := range historySlice {
		if strings.HasPrefix(v, ":") { //check if there is a timestamp and remove it
			commandSlice := strings.Split(v, ";")
			historySlice[i] = commandSlice[1]
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

func countCommands(list []string) map[string]float64 {
	duplicate := map[string]float64{}

	for _, item := range list {
		_, exist := duplicate[item]

		if exist {
			duplicate[item]++
		} else {
			duplicate[item] = 1
		}
	}
	return duplicate
}

func sortCommandsMap(commands map[string]float64) (commandsNames []string, commandsNums []float64) {
	keys := make([]string, 0, len(commands))
	values := make([]float64, 0, len(commands))

	for name := range commands {
		keys = append(keys, name)
	}

	sort.Slice(keys, func(i, j int) bool {
		return commands[keys[i]] > commands[keys[j]]
	})

	for _, name := range keys {
		values = append(values, commands[name])
	}

	return keys, values
}

func renderPieChart(names []string, values []float64) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	commandsNum := strconv.Itoa(len(names))
	pc := widgets.NewPieChart()
	pc.Title = commandsNum + " Most Used Commands"
	pc.SetRect(0, 0, 75, 40)
	pc.AngleOffset = 20 * math.Pi
	pc.Data = values
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%s: " + "%.00f", names[i], v)
	}

	ui.Render(pc)

	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		}
	}
}

func main() {
	history := fetchBashHistory()
	commands := countCommands(history)
	delete(commands, "")

	commandsNames, commandsNums := sortCommandsMap(commands)

	renderPieChart(commandsNames[:10], commandsNums[:10])
}