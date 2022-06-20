package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/aejuicellc/connection-test-tool/internals/utils"
	"sync"
)

func BuildMainContainer(topLevelCanvas *fyne.Container, progressBar *widget.ProgressBar, app fyne.App) {
	var mutex sync.Mutex
	saveLogsButton := widget.NewButton("Save Logs", utils.SaveLogs)
	status := "hidden"
	showLogsButton := widget.NewButton("Show Logs", func() {
		if status == "hidden" {
			status = "visible"
			logsWindow := CreateLogsWindow(app)
			logsWindow.Show()
			logsWindow.SetOnClosed(func() {
				status = "hidden"
			})
		}
	})

	for _, value := range utils.LinksToCheck {
		grid := container.New(layout.NewGridLayout(2))
		topLevelCanvas.Add(grid)
		nameWidget := CreateLabelWithStringData(value.Name)
		statusWidget := CreateLabelWithStringData("...")
		statusWidget.Alignment = fyne.TextAlignTrailing
		go BuildStatusColumn(value, grid, statusWidget, progressBar, &mutex)

		grid.Add(nameWidget)
		grid.Add(statusWidget)
	}
	topLevelCanvas.Add(saveLogsButton)
	topLevelCanvas.Add(showLogsButton)
}

func CreateLabelWithStringData(value string) *widget.Label {
	str := binding.NewString()
	str.Set(value)
	return widget.NewLabelWithData(str)
}
