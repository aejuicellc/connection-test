package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"main/internals/utils"
	"time"
)

func CreateLogsWindow(application fyne.App) fyne.Window {
	currentWindow := application.NewWindow("AEJuice Network Diagnostics Tool Logs")
	textBinding := binding.NewString()
	go func() {
		for {
			textBinding.Set(fmt.Sprintf("%s", utils.LogsText))
			time.Sleep(1 * time.Second)
		}
	}()

	textWidget := widget.NewLabelWithData(textBinding)
	scrollContainer := container.NewScroll(textWidget)
	currentWindow.SetContent(scrollContainer)
	currentWindow.Resize(fyne.Size{
		Width:  900,
		Height: 400,
	})
	return currentWindow
}
