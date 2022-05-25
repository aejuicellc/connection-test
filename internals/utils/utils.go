package utils

import (
	"encoding/json"
	"fyne.io/fyne/v2/widget"
	"main/internals/types"
	"os"
)

var LinksToCheck = GetLinksToCheck()
var ProgressBarValue = 0.0

func GetLinksToCheck() []types.LinkToCheck {
	links := make([]types.LinkToCheck, 0)
	sources := GetEnv("URLS_TO_CHECK", "[{\"url\": \"https://aejuice.com\", \"name\": \"Website and API\"},{\"url\": \"https://nyc3.digitaloceanspaces.com/aejuice/update/updater/version.txt?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=SSK2B4AURYVVYMUF75K3%2F20220524%2Fnyc3%2Fs3%2Faws4_request&X-Amz-Date=20220524T214003Z&X-Amz-Expires=604800&X-Amz-SignedHeaders=host&X-Amz-Signature=8a9d408ac988642ffeb3021eb5447f3a423e96ea335aa6ad46ff63a66cb0a83d\", \"name\": \"Possibility to download files\"}]")
	_ = json.Unmarshal([]byte(sources), &links)

	return links
}

func GetEnv(envName string, fallback string) string {
	value := os.Getenv(envName)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func IncrementProgressValue() {
	floatLength := 1.00 / float64(len(LinksToCheck))
	ProgressBarValue = ProgressBarValue + floatLength
}

func IncrementProgress(progressBar *widget.ProgressBar) {
	IncrementProgressValue()
	progressBar.SetValue(ProgressBarValue)
}
