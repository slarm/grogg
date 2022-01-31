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

func Match(s string, gm *GrokMatcher) (map[string]string, error) {
	g, _ := grok.New()
	if gm.CustomPatterns != "" {
		err := g.AddPatternsFromPath(gm.CustomPatterns)
		if err != nil {
			return nil, err
		}
	}

	values, err := g.Parse(gm.MatchString, s)

	return values, err
}
