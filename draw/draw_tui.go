package draw

import (
	ui "github.com/gizak/termui/v3"
	"log"
)

func RenderTui(names []string, occurrences []float64, shell string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	header := SetupHeader()
	tabPane := SetupTabPane()
	list := SetupList(names, occurrences, shell, 1,4)
	pieChart := SetupPieChart(names, occurrences, shell,1, 4)
	barChart := SetupBarChart(names, occurrences, shell,1, 4)

	renderTab := func() {
		switch tabPane.ActiveTabIndex {
		case 0:
			ui.Render(pieChart)
		case 1:
			ui.Render(barChart)
		case 2:
			ui.Render(list)
		}
	}

	ui.Render(header, tabPane, pieChart)
	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabPane.FocusLeft()
			ui.Clear()
			ui.Render(header, tabPane)
			renderTab()
		case "l":
			tabPane.FocusRight()
			ui.Clear()
			ui.Render(header, tabPane)
			renderTab()
		}
	}
}

