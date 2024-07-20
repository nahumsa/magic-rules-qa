package parser

import (
	"os"
	"regexp"
	"strings"
)

type Rule struct {
	Code string
	Text string
}

func parseRules(text string) []Rule {
	var rules []Rule

	ruleRegex := regexp.MustCompile(`^(\d+\.\d+[a-z.]*|\d+.)\s(.+)`)

	lines := strings.Split(text, "\n")

	var currentRule Rule

	for _, line := range lines {
		match := ruleRegex.FindStringSubmatch(line)
		if len(match) == 3 {
			ruleNumber := match[1]
			ruleText := match[2]

			if currentRule.Code != "" {
				rules = append(rules, currentRule)
			}

			currentRule = Rule{
				Code: ruleNumber,
				Text: ruleText,
			}
		} else {
			if currentRule.Code != "" {
				currentRule.Text += "\n" + line
			}
		}
	}
	if currentRule.Code != "" {
		rules = append(rules, currentRule)
	}

	return rules
}

func parse_file(path string) ([]Rule, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	st := strings.Split(string(f), "Glossary")

	rules := parseRules(st[1])

	return rules, nil
}
