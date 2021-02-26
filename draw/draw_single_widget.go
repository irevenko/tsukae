package draw

import (
	ui "github.com/gizak/termui/v3"
	"log"
)

func RenderPieChart(names []string, occurrences []float64, shell string, commandsNum int) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	pieChart := SetupPieChart(names, occurrences, shell, 0,0)

	ui.Render(pieChart)
	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func RenderBarChart(names []string, occurrences []float64, shell string, commandsNum int) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	barChart := SetupBarChart(names, occurrences, shell, 0,0)

	ui.Render(barChart)
	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
func RenderList(names []string, occurrences []float64, shell string, commandsNum int) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	list := SetupList(names, occurrences, shell, 0,0)

	ui.Render(list)
	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}