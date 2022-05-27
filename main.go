package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/getsentry/sentry-go"
	"log"
	"main/internals/ui"
	"main/internals/utils"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: utils.GetEnv("SENTRY_DSN", ""),
	})
	if err != nil {
		log.Panic(err)
	}
	application := app.New()
	currentWindow := application.NewWindow("AEJuice Network Diagnostics Tool")
	currentWindow.Resize(fyne.Size{
		Width: 400,
	})
	progress := widget.NewProgressBar()
	topLevelGrid := container.New(layout.NewGridLayout(1), container.NewVBox(progress))
	currentWindow.SetPadded(true)
	currentWindow.SetContent(topLevelGrid)
	go ui.BuildMainContainer(topLevelGrid, progress, application)
	currentWindow.ShowAndRun()
}
