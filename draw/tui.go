package draw

import (
	"fmt"
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
	"log"
	"math"
	"strconv"
)

func RenderTui(names []string, occurrences []float64) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	header := widgets.NewParagraph()
	header.Border = false
	header.Text = "h & l to switch tabs | q to quit"
	header.SetRect(0, 0, 50, 1)
	header.TextStyle.Bg = ui.ColorCyan

	pc := widgets.NewPieChart()
	pc.Border = false
	pc.Title = strconv.Itoa(len(names)) + " Most Used Commands"
	pc.SetRect(1, 4, 50, 30)
	pc.AngleOffset = .150 * math.Pi
	pc.Data = occurrences
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.00f" + " %s", v, names[i])
	}

	bc := widgets.NewBarChart()
	bc.Border = false
	bc.Title = strconv.Itoa(len(names)) + " Most Used Commands"
	bc.Data = occurrences
	bc.Labels = names
	bc.SetRect(1, 4, 170, 20)
	bc.BarWidth = 8
	bc.BarGap = 3

	l := widgets.NewList()
	l.Border = false
	l.Title = strconv.Itoa(len(names)) + " Most Used Commands"
	var listData []string
	listColors := []string{"](fg:green)", "](fg:blue)","](fg:yellow)", "](fg:magenta)", "](fg:cyan)", "](fg:red)"}
	j := 0
	for i := 0; i < len(names); i++ {
		if j > len(listColors)-1 {
			j = 0
		}
		listData = append(listData, "[[" + strconv.FormatFloat(occurrences[i], 'f', 0, 64) + "] " + names[i] + listColors[j])
		j++
	}
	l.Rows = listData
	l.SelectedRowStyle = ui.NewStyle(ui.ColorGreen)
	l.WrapText = false
	l.SetRect(1, 4, 25, 30)

	tabpane := widgets.NewTabPane("PieChart", "BarChart", "List")
	tabpane.ActiveTabStyle.Fg = ui.ColorCyan
	tabpane.SetRect(0, 1, 30, 4)
	tabpane.Border = false

	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			ui.Render(pc)
		case 1:
			ui.Render(bc)
		case 2:
			ui.Render(l)
		}
	}

	ui.Render(header, tabpane, pc)
	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabpane.FocusLeft()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		case "l":
			tabpane.FocusRight()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		}
	}
}
