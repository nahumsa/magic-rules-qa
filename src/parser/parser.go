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

type Keyword struct {
	Title string
	Text  string
}

func parseKeywords(input string) []Keyword {
	re := regexp.MustCompile(`(?<keyword>[^\n]+)\n(?<text>.+)`)

	matches := re.FindAllStringSubmatch(input, -1)

	names := re.SubexpNames()

	var keywords []Keyword

	for _, match := range matches {
		result := make(map[string]string)
		for i, name := range names {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		keyword := Keyword{
			Title: result["keyword"],
			Text:  result["text"],
		}
		keywords = append(keywords, keyword)
	}

	return keywords
}

func parseMainRules(text string) []Rule {
	var rules []Rule

	ruleRegex := regexp.MustCompile(`^(\d+\.)\s(.+)`)

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

func ParseFile(path string) ([]Rule, []Keyword, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	st := strings.Split(string(f), "Glossary")

	rules := parseMainRules(st[1])
	keywords := parseKeywords(st[2])

	return rules, keywords, nil
}
