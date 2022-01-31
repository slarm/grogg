package internal

import (
	"bufio"
	"bytes"
	"testing"
)

func TestSingleMatch(t *testing.T) {
	sc := NewScanConfig()
	sc.BufferTokenSize = 4096
	sc.FailureMode = false
	sc.InputFile = "testdata/single.log"
	sc.MultiLinePattern = ""
	sc.SilentMode = false
	sc.Verbose = true

	gm := NewGrokMatcher()
	gm.MatchString = "%{SYSLOGTIMESTAMP} %{NUMBER} %{LOGLEVEL} %{GREEDYDATA}"
	gm.PromptMode = true

	ScanFile(sc, gm)

	if gm.Failures != 1 && gm.Events != 2 {
		t.Errorf("Expected 0 failures and 2 events, got %d failures and %d events", gm.Failures, gm.Events)
	}
}

func TestMulti(t *testing.T) {
	sc := NewScanConfig()
	sc.BufferTokenSize = 4096
	sc.FailureMode = false
	sc.InputFile = "testdata/multi.log"
	sc.MultiLinePattern = "^\\w{3}\\s\\d{2}\\s(?:\\d{2}:){2}\\d{2}"
	sc.SilentMode = false
	sc.Verbose = true

	gm := NewGrokMatcher()
	gm.MatchString = "%{SYSLOGTIMESTAMP} %{WORD} %{LOGLEVEL} %{MULTILINEDATA}"
	gm.CustomPatterns = "testdata/grok-patterns"

	ScanFile(sc, gm)

	if gm.Failures != 0 && gm.Events != 2 {
		t.Errorf("Expected 0 failures and 2 events, got %d failures and %d events", gm.Failures, gm.Events)
	}
}

func TestPrinter(t *testing.T) {
	PrintLogo()

	buf := bytes.NewBufferString("n")
	PrintPrompt(bufio.NewScanner(buf))

	buf = bytes.NewBufferString("y")
	PrintPrompt(bufio.NewScanner(buf))
}

func TestInputError(t *testing.T) {
	sc := NewScanConfig()
	sc.InputFile = "testdata/fubar.log"

	gm := NewGrokMatcher()

	err := ScanFile(sc, gm)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestCustomPatternError(t *testing.T) {
	gm := NewGrokMatcher()
	gm.CustomPatterns = "testdata/fubar"

	_, err := Match("", gm)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}
