package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"main/internals/types"
	"main/internals/utils"
	"net/http"
)

func BuildStatusColumn(value types.LinkToCheck, grid *fyne.Container, statusWidget *widget.Label, progressBar *widget.ProgressBar) {
	var text *canvas.Text
	response, err := http.Get(value.Url)
	if err != nil || response.StatusCode != 200 {
		text = canvas.NewText("Unavailable", color.RGBA{R: 201, G: 84, B: 60, A: 1})
	} else {
		text = canvas.NewText("Available", color.RGBA{R: 95, G: 237, B: 85, A: 1})
	}
	grid.Remove(statusWidget)
	text.Alignment = fyne.TextAlignTrailing
	grid.Add(text)
	utils.IncrementProgress(progressBar)
}
