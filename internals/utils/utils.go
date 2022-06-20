package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/internals/types"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

var LogsText []string
var LinksToCheck = GetLinksToCheck()
var ProgressBarValue = 0.0

func GetLinksToCheck() []types.LinkToCheck {
	links := make([]types.LinkToCheck, 0)
	sources := GetEnv("URLS_TO_CHECK", "[\n    {\n        \"url\": \"https://aejuice.com\",\n        \"name\": \"Possibility to fetch site\"\n    },\n    {\n        \"url\": \"https://api-free.deepl.com\",\n        \"name\": \"Possibility to translate\"\n    },\n    {\n        \"url\": \"https://aejuice.atlassian.net\",\n        \"name\": \"Possibility to bug reporting #1\"\n    }, \n    {\n        \"url\": \"https://sentry.aejuice.xyz\",\n        \"name\": \"Possibility to bug reporting #2\"\n    },\n    {\n        \"url\": \"https://bitbucket.org\",\n        \"name\": \"Possibility to update\"\n    },\n    {\n        \"url\": \"https://github.com\",\n        \"name\": \"Possibility to download releases\"\n    },\n    {\n        \"url\": \"https://api.pexels.com\",\n        \"name\": \"Possibility to download from stock\"\n    },\n    {\n        \"url\": \"https://nyc3.digitaloceanspaces.com\",\n        \"name\": \"Possibility to download files #1\"\n    },\n    {\n        \"url\": \"https://packmanager.nyc3.digitaloceanspaces.com\",\n        \"name\": \"Possibility to download files #2\"\n    },\n    {\n        \"url\": \"https://aejuice.nyc3.digitaloceanspaces.com\",\n        \"name\": \"Possibility to download files #3\"\n    },\n    {\n        \"url\": \"https://storage.aejuice.xyz\",\n        \"name\": \"Possibility to download files #4\"\n    },\n    {\n        \"url\": \"https://pcdn.aejuice.com\",\n        \"name\": \"Possibility to download files #5\"\n    }\n]")
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

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func GetTracerouteFunction() string {
	if IsWindows() {
		return "tracert"
	}

	return "traceroute"
}

func GetMaxHopsArg() string {
	if IsWindows() {
		return "-h 30"
	}

	return "-m30"
}

func Traceroute(href string) error {
	currentUrl, _ := url.Parse(href)

	if IsWindows() {
		pattern := regexp.MustCompile(`[A-Z]\:\\Windows;`)
		windowsPath := pattern.FindString(os.Getenv("PATH"))
		windowsPath = strings.Replace(windowsPath, ";", "", 1)
		cmdToRun := windowsPath + "\\System32\\TRACERT.EXE"
		args := []string{"-h 30", currentUrl.Host}
		procAttr := new(os.ProcAttr)
		var stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
		var stderr = os.NewFile(uintptr(syscall.Stderr), "/dev/stderr")
		procAttr.Files = []*os.File{os.Stdin, stdout, stderr}
		stdoutScanner := bufio.NewScanner(procAttr.Files[1])
		stderrScanner := bufio.NewScanner(procAttr.Files[2])

		go func() {
			for stderrScanner.Scan() {
				LogsText = append(LogsText, stderrScanner.Text()+"\n")
			}
		}()
		go func() {
			for stdoutScanner.Scan() {
				LogsText = append(LogsText, stdoutScanner.Text()+"\n")
			}
		}()

		if process, err := os.StartProcess(cmdToRun, args, procAttr); err != nil {
			log.Fatal("ERROR Unable to run: \n", cmdToRun, err.Error())
		} else {
			LogsText = append(LogsText, "Running as pid "+strconv.Itoa(process.Pid))
		}
	} else {
		tracerouteFunction := GetTracerouteFunction()
		hopsArg := GetMaxHopsArg()
		cm := exec.Command(tracerouteFunction, hopsArg, currentUrl.Host)
		stdout, _ := cm.StdoutPipe()
		stderr, _ := cm.StderrPipe()
		err := cm.Start()
		if err != nil {
			LogsText = append(LogsText, fmt.Sprintf("%s", err)+"\n")
			log.Fatal(err)
			return err
		}
		stdoutScanner := bufio.NewScanner(stdout)
		stderrScanner := bufio.NewScanner(stderr)
		go func() {
			for stderrScanner.Scan() {
				LogsText = append(LogsText, stderrScanner.Text()+"\n")
			}
		}()
		go func() {
			for stdoutScanner.Scan() {
				LogsText = append(LogsText, stdoutScanner.Text()+"\n")
			}
		}()
		err = cm.Wait()
		if err != nil {
			LogsText = append(LogsText, fmt.Sprintf("%s", err)+"\n")
			log.Fatal(err)
			return err
		}

	}
	return nil
}

func SaveLogs() {
	var bytesArr []byte
	for i := 0; i < len(LogsText); i++ {
		b := []byte(LogsText[i])
		for j := 0; j < len(b); j++ {
			bytesArr = append(bytesArr, b[j])
		}
	}

	err := ioutil.WriteFile("network-diagnostic-tool.log", bytesArr, 0644)
	if err != nil {
		log.Panic(err)
	}
}
