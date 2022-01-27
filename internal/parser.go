package internal

import (
	"fmt"
)

func Parse(l string, eof int, gm *GrokMatcher, sc *ScanConfig) {
	gm.IncEvents()
	result := Match(l, gm)
	if len(result) == 0 {
		gm.IncFailures()
		if !sc.SilentMode {
			PrintLine("Grok failure on: " + l)
			fmt.Print("\n")
		}
	} else {
		if !sc.FailureMode || (!sc.SilentMode && !sc.FailureMode) {
			PrintResult(result)
			if sc.Verbose {
				PrintLine(l)
				PrintSeparator()
			}
			fmt.Print("\n")
		}
	}
}
