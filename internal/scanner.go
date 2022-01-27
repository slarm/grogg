package internal

import (
	"bufio"
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

func scanMatcher(ml *multiLine, sc *ScanConfig, scanner *bufio.Scanner, mlRegexp *regexp.Regexp, ch chan string) {
	for scanner.Scan() {
		if sc.MultiLinePattern != "" {
			if !mlRegexp.Match([]byte(scanner.Text())) || len(ml.lines) == 0 {
				ml.addLine(scanner.Text())
				continue
			} else {
				ch <- strings.Join(ml.lines, " ")
				ml.lines = nil
				ml.addLine(scanner.Text())
			}
		} else {
			ch <- scanner.Text()
		}
	}
	if len(ml.lines) > 0 {
		ch <- strings.Join(ml.lines, "\n")
	}
	close(ch)
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

	go scanMatcher(ml, sc, scanner, mlRegexp, ch)

	for l := range ch {
		Parse(l, len(ch), gm, sc)
		if gm.PromptMode {
			con := PrintPrompt()
			if !con {
				break
			}
		}
	}
}
