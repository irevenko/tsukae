package draw

import (
	"fmt"
	"math"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func SetupHeader() *widgets.Paragraph {
	header := widgets.NewParagraph()
	header.Border = false
	header.Text = "h & l to switch tabs | q to quit"
	header.SetRect(0, 0, 50, 1)
	header.TextStyle.Bg = ui.ColorBlue

	return header
}

func SetupPieChart(names []string, occurrences []float64, shell string, x1 int, y1 int) *widgets.PieChart {
	pc := widgets.NewPieChart()
	pc.Border = false
	pc.Title = strconv.Itoa(len(names)) + " Most Used " + shell + " Commands"

	if len(names) < 6 {
		pc.SetRect(x1, y1, 50, 25)
	} else if len(names) < 10 {
		pc.SetRect(x1, y1, 60, 30)
	} else {
		pc.SetRect(x1, y1, 65, 35)
	}

	pc.AngleOffset = .15 * math.Pi
	pc.Data = occurrences
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.00f"+" %s", v, names[i])
	}

	return pc
}

func SetupBarChart(names []string, occurrences []float64, shell string, x1 int, y1 int) *widgets.BarChart {
	bc := widgets.NewBarChart()
	bc.Border = false
	bc.Title = strconv.Itoa(len(names)) + " Most Used " + shell + " Commands"
	bc.Data = occurrences
	bc.Labels = names

	if len(names) < 6 {
		bc.SetRect(x1, y1, 170, 13)
	} else if len(names) < 10 {
		bc.SetRect(x1, y1, 170, 18)
	} else {
		bc.SetRect(x1, y1, 170, 23)
	}

	bc.BarWidth = 8
	bc.BarGap = 3

	return bc
}

func SetupList(names []string, occurrences []float64, shell string, x1 int, y1 int) *widgets.List {
	var listData []string
	listColors := []string{"](fg:green)", "](fg:blue)", "](fg:yellow)", "](fg:magenta)", "](fg:cyan)", "](fg:red)"}
	j := 0

	l := widgets.NewList()
	l.Border = false
	l.Title = strconv.Itoa(len(names)) + " Most Used " + shell + " Commands"

	for i := 0; i < len(names); i++ {
		if j > len(listColors)-1 {
			j = 0
		}
		listData = append(listData, "[["+strconv.FormatFloat(occurrences[i], 'f', 0, 64)+"] "+names[i]+listColors[j])
		j++
	}
	l.Rows = listData
	l.SelectedRowStyle = ui.NewStyle(ui.ColorRed)
	l.WrapText = false
	l.SetRect(x1, y1, 30, 30)

	return l
}

func SetupTabPane() *widgets.TabPane {
	tabPane := widgets.NewTabPane("PieChart", "BarChart", "List")
	tabPane.ActiveTabStyle.Fg = ui.ColorGreen
	tabPane.SetRect(0, 1, 30, 4)
	tabPane.Border = true

	return tabPane
}
