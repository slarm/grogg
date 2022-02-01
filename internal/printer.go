package internal

import (
	"bufio"
	"fmt"

	"github.com/pterm/pterm"
)

func PrintResult(m map[string]string) {
	from := pterm.NewRGB(0, 255, 255)
	to := pterm.NewRGB(255, 0, 255)

	i := 0
	for k, v := range m {
		from.Fade(0, float32(len(m)), float32(i), to).Printf("%15s: %s\n", k, v)
		i++
	}
}

func PrintError(s string) {
	pterm.NewRGB(255, 0, 0).Println(s)
}

func PrintLine(s string) {
	pterm.NewRGB(255, 255, 0).Println(s)
}

func PrintSeparator() {
	for i := 0; i < pterm.GetTerminalWidth()/2; i++ {
		fmt.Print("-")
	}
}

func PrintLogo() {
	pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("GROGG")).Render()
}

func PrintPrompt(scanner *bufio.Scanner) bool {
	fmt.Println("Continue? (y/n)")
	for scanner.Scan() {
		if scanner.Text() == "n" {
			break
		} else if scanner.Text() == "y" {
			return true
		}
	}
	return false
}
