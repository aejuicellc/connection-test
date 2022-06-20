package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"main/internals/types"
	"main/internals/utils"
	"net/http"
	"sync"
)

func BuildStatusColumn(value types.LinkToCheck, grid *fyne.Container, statusWidget *widget.Label, progressBar *widget.ProgressBar, mutex *sync.Mutex) {
	var text *canvas.Text
	response, err := http.Get(value.Url)
	fmt.Println(response.StatusCode)
	if err != nil {
		text = canvas.NewText("Unavailable", color.RGBA{R: 201, G: 84, B: 60, A: 1})
		utils.LogsText = append(utils.LogsText, fmt.Sprintf("Response status of %s was %s. %s is unavailable. Tracerouting... \n", value.Url, err, value.Url))
		mutex.Lock()
		utils.Traceroute(value.Url)
		mutex.Unlock()
	} else {
		utils.LogsText = append(utils.LogsText, fmt.Sprintf("Response status of %s was %d. %s is available \n", value.Url, response.StatusCode, value.Url))
		text = canvas.NewText("Available", color.RGBA{R: 95, G: 237, B: 85, A: 1})
	}
	grid.Remove(statusWidget)
	text.Alignment = fyne.TextAlignTrailing
	grid.Add(text)
	IncrementProgress(progressBar)
}

func IncrementProgress(progressBar *widget.ProgressBar) {
	utils.IncrementProgressValue()
	progressBar.SetValue(utils.ProgressBarValue)
}
