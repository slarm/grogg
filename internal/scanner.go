package internal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type ScanConfig struct {
	BufferTokenSize  int
	FailureMode      bool
	InputFile        string
	MultiLinePattern string
	SilentMode       bool
	Verbose          bool
}

func NewScanConfig() *ScanConfig {
	return &ScanConfig{}
}

type multiLine struct {
	lines []string
}

func newMultiLine() *multiLine {
	return &multiLine{}
}

func (m *multiLine) addLine(s string) {
	m.lines = append(m.lines, s)
}

func ScanFile(sc *ScanConfig, gm *GrokMatcher) {
	f, err := os.Open(sc.InputFile)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	ml := newMultiLine()
	mlRegexp := regexp.MustCompile(sc.MultiLinePattern)

	scanner := bufio.NewScanner(f)
	buf := make([]byte, sc.BufferTokenSize)
	scanner.Buffer(buf, sc.BufferTokenSize)

	ch := make(chan string)

	go func() {
		for scanner.Scan() {
			if sc.MultiLinePattern != "" {
				if !mlRegexp.Match([]byte(scanner.Text())) || len(ml.lines) == 0 {
					ml.addLine(scanner.Text())
					continue
				}
				ch <- strings.Join(ml.lines, "\n")
				ml.lines = nil
				ml.addLine(scanner.Text())
				fmt.Println(ml.lines)
			} else {
				ch <- scanner.Text()
			}
		}
		if len(ml.lines) > 0 {
			ch <- strings.Join(ml.lines, "\n")
		}
		close(ch)
	}()

	for l := range ch {
		Parse(l, len(ch), gm, sc)
		if gm.PromptMode {
			in := bufio.NewScanner(os.Stdin)
			con := PrintPrompt(in)
			fmt.Println(con)
			if !con {
				break
			}
		}
	}
}
