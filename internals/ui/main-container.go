package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"main/internals/utils"
)

func BuildMainContainer(topLevelCanvas *fyne.Container, progressBar *widget.ProgressBar) {
	for _, value := range utils.LinksToCheck {
		grid := container.New(layout.NewGridLayout(2))
		topLevelCanvas.Add(grid)
		nameWidget := createLabelWithStringData(value.Name)
		statusWidget := createLabelWithStringData("...")
		statusWidget.Alignment = fyne.TextAlignTrailing

		go BuildStatusColumn(value, grid, statusWidget, progressBar)

		grid.Add(nameWidget)
		grid.Add(statusWidget)
	}
}

func createLabelWithStringData(value string) *widget.Label {
	str := binding.NewString()
	str.Set(value)

	return widget.NewLabelWithData(str)
}
