package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetEnvShouldReturnENVOrFallback(t *testing.T) {
	envVar := GetEnv("GOPATH", "")
	if envVar == "" {
		t.Fatalf(`ENV should be not empty, because GOPATH is defined`)
	}
	envVar = GetEnv("DOESNT_EXIST", "fallback")
	if envVar != "fallback" {
		t.Fatalf(`ENV should be not empty, fallback is defined`)
	}
}

func TestGetLinksToCheckAlwaysReturnSliceOfLinksToCheck(t *testing.T) {
	linksToCheck := GetLinksToCheck()
	if fmt.Sprintf("%s", reflect.TypeOf(linksToCheck)) != "[]types.LinkToCheck" {
		t.Fatalf(`GetLinksToCheck should only return []types.LinkToCheck`)
	}
}

func TestProgressBarIncrementing(t *testing.T) {
	if ProgressBarValue == 0.0 {
		IncrementProgressValue()
		if ProgressBarValue == 0.0 {
			t.Fatalf(`Progress bar value should be incremented`)
		}
		return
	}

	t.Fatalf(`Progress bar value should be 0 at start`)
}
