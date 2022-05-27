package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestGetEnvShouldReturnENVOrFallback(t *testing.T) {
	os.Setenv("TEST_ENV", "Test variable")
	envVar := GetEnv("TEST_ENV", "")
	if envVar == "" {
		t.Fatalf(`ENV should be not empty, because GOOS is defined`)
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

func TestIsWindows(t *testing.T) {
	isWindows := IsWindows()
	if runtime.GOOS == "windows" && !isWindows {
		t.Fatalf(`Should return true on windows arch`)
	}
}

func TestGetTracerouteFunction(t *testing.T) {
	if runtime.GOOS == "windows" {
		if GetTracerouteFunction() != "tracert" {
			t.Fatalf(`Traceroute function should be tracert`)
		}
		return
	}
	if GetTracerouteFunction() != "traceroute" {
		t.Fatalf(`Traceroute function should be traceroute`)
	}
	return
}

func TestTraceroute(t *testing.T) {
	var err error
	go func() {
		err = Traceroute("https://google.com")
	}()

	time.Sleep(3 * time.Second)
	if err != nil {
		t.Fatalf(`Traceroute should work, [` + err.Error() + "]")
	}
	return
}

func TestFileShouBeSaved(t *testing.T) {
	LogsText = []string{"test"}
	SaveLogs()

	fileContent, err := ioutil.ReadFile("network-diagnostic-tool.log")

	defer os.Remove("network-diagnostic-tool.log")
	if err != nil {
		t.Fatalf(`File shoud be created`)
	}
	if string(fileContent) != "test" {
		t.Fatalf(`File shoud contain test value`)
	}
}
