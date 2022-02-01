package main

import (
	"flag"
	"fmt"
	"grogg/internal"
	"os"
	"strconv"
)

var version string = string("0.4")

var bufferTokenSize *int = flag.Int("line-length", 4096, "Set max line length buffer in bytes")
var customFile *string = flag.String("custom-file", "", "File containing custom Grok patterns")
var silentMode *bool = flag.Bool("silent", false, "Silent mode (summary at the end of execution)")
var failureMode *bool = flag.Bool("failures", false, "Show only failures")
var inputFile *string = flag.String("input", "", "Input file")
var matchString *string = flag.String("pattern", "", "Grok pattern to match")
var multiLinePattern *string = flag.String("multiline", "", "Multiline pattern")
var promptMode *bool = flag.Bool("prompt", false, "Prompt after each line")
var verbose *bool = flag.Bool("verbose", false, "Verbose mode")
var printVersion *bool = flag.Bool("version", false, "Print version")

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Println("Grogg version: " + version)
		os.Exit(0)
	}

	if *inputFile == "" || *matchString == "" {
		internal.PrintLogo()
		flag.PrintDefaults()
		os.Exit(0)
	}

	sc := internal.NewScanConfig()
	sc.BufferTokenSize = *bufferTokenSize
	sc.FailureMode = *failureMode
	sc.InputFile = *inputFile
	sc.MultiLinePattern = *multiLinePattern
	sc.SilentMode = *silentMode
	sc.Verbose = *verbose

	gm := internal.NewGrokMatcher()
	gm.CustomPatterns = *customFile
	gm.MatchString = *matchString
	gm.PromptMode = *promptMode

	err := internal.ScanFile(sc, gm)
	if err != nil {
		panic(err)
	}
	fmt.Println("Matched events: " + strconv.Itoa(gm.Events))
	fmt.Println("Failures: " + strconv.Itoa(gm.Failures))

	os.Exit(0)
}
