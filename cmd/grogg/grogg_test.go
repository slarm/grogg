package main

import (
	"grogg/internal"
	"testing"
)

func TestSingleMatch(t *testing.T) {
	sc := internal.NewScanConfig()
	sc.BufferTokenSize = 4096
	sc.FailureMode = false
	sc.InputFile = "testdata/single.log"
	sc.MultiLinePattern = ""
	sc.SilentMode = false
	sc.Verbose = false

	gm := internal.NewGrokMatcher()
	gm.MatchString = "%{SYSLOGTIMESTAMP} %{WORD} %{LOGLEVEL} %{GREEDYDATA}"

	internal.ScanFile(sc, gm)

	if gm.Failures != 0 {
		t.Errorf("Expected 0 failures, got %d", gm.Failures)
	}
}

func TestMulti(t *testing.T) {
	sc := internal.NewScanConfig()
	sc.BufferTokenSize = 4096
	sc.FailureMode = false
	sc.InputFile = "testdata/multi.log"
	sc.MultiLinePattern = "^\\w{3}\\s(\\d{2}:){2}\\d{2}"
	sc.SilentMode = false
	sc.Verbose = false

	gm := internal.NewGrokMatcher()
	gm.MatchString = "%{SYSLOGTIMESTAMP} %{WORD} %{LOGLEVEL} %{MULTILINEDATA}"
	gm.CustomPatterns = "testdata/grok-patterns"

	internal.ScanFile(sc, gm)

	if gm.Failures != 0 {
		t.Errorf("Expected 0 failures, got %d", gm.Failures)
	}
}
