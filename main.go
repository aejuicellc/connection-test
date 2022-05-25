package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"main/internals/ui"
)

func main() {
	application := app.New()
	currentWindow := application.NewWindow("AEJuice Network Diagnostics Tool")
	progress := widget.NewProgressBar()
	topLevelGrid := container.New(layout.NewGridLayout(1), container.NewVBox(progress))
	currentWindow.SetPadded(true)
	currentWindow.SetContent(topLevelGrid)

	go ui.BuildMainContainer(topLevelGrid, progress)

	currentWindow.ShowAndRun()
}
