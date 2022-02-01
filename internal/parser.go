package internal

import (
	"fmt"
)

func Parse(l string, eof int, gm *GrokMatcher, sc *ScanConfig) error {
	if gm.Events == 0 {
		PrintSeparator()
		fmt.Print("\n")
		if sc.Verbose {
			fmt.Printf("Input file: %s\n", sc.InputFile)
			fmt.Printf("Grok pattern: %s\n", gm.MatchString)
			fmt.Printf("Multiline pattern: %s\n", sc.MultiLinePattern)
			fmt.Printf("Custom patterns file: %s\n", gm.CustomPatterns)
			PrintSeparator()
			fmt.Print("\n")
		}
	}
	gm.IncEvents()
	result, err := Match(l, gm)
	if len(result) == 0 {
		gm.IncFailures()
		if !sc.SilentMode {
			PrintError("Grok failure on: " + l)
			fmt.Print("\n")
		}
	} else {
		if !sc.FailureMode || (!sc.SilentMode && !sc.FailureMode) {
			if sc.Verbose {
				fmt.Print("Matching line:\n     ")
				PrintLine(l)
				fmt.Print("\nResult:\n")
			}
			PrintResult(result)
			PrintSeparator()
			fmt.Print("\n")
		}
	}
	return err
}
