package internal

import (
	"github.com/vjeantet/grok"
)

type GrokMatcher struct {
	CustomPatterns string
	Events         int
	Failures       int
	MatchString    string
	PromptMode     bool
	Results        []map[string]string
}

func NewGrokMatcher() *GrokMatcher {
	return &GrokMatcher{}
}

func (g *GrokMatcher) IncFailures() {
	g.Failures++
}

func (g *GrokMatcher) IncEvents() {
	g.Events++
}

func (g *GrokMatcher) AddResult(s map[string]string) {
	g.Results = append(g.Results, s)
}

func Match(s string, gm *GrokMatcher) map[string]string {
	g, _ := grok.New()
	if gm.CustomPatterns != "" {
		err := g.AddPatternsFromPath(gm.CustomPatterns)
		if err != nil {
			panic(err)
		}
	}

	values, err := g.Parse(gm.MatchString, s)
	if err != nil {
		panic(err)
	}
	return values
}
