package parser

import (
	"testing"
)

func TestMainRuleParser(t *testing.T) {
	example := `902. Testing
902.7. Any abilities of a face-up vanguard card in the command zone function from that zone. The card’s static abilities affect the game, its triggered abilities may trigger, and its activated abilities may be activated.

903. Commander

903.1. In the Commander variant, each deck is led by a legendary creature designated as that deck’s commander. The Commander variant was created and popularized by fans; an independent rules committee maintains additional resources at MTGCommander.net. The Commander variant uses all the normal rules for a Magic game, with the following additions.

903.2. A Commander game may be a two-player game or a multiplayer game. The default multiplayer setup is the Free-for-All variant with the attack multiple players option and without the limited range of influence option. See rule 806, “Free-for-All Variant.”

903.3. Each deck has a legendary creature card designated as its commander. This designation is not a characteristic of the object represented by the card; rather, it is an attribute of the card itself. The card retains this designation even when it changes zones.
Example: A commander that’s been turned face down (due to Ixidron’s effect, for example) is still a commander. A commander that’s copying another card (due to Cytoshape’s effect, for example) is still a commander. A permanent that’s copying a commander (such as a Body Double, for example, copying a commander in a player’s graveyard) is not a commander.`
	rules := parseMainRules(example)

	if len(rules) != 2 {
		t.Errorf("expected 2 rules, got=%d", len(rules))
	}
	expected_rules := []Rule{
		{
			Code: "902.",
			Text: `Testing
902.7. Any abilities of a face-up vanguard card in the command zone function from that zone. The card’s static abilities affect the game, its triggered abilities may trigger, and its activated abilities may be activated.
`,
		},
		{
			Code: "903.",
			Text: `Commander

903.1. In the Commander variant, each deck is led by a legendary creature designated as that deck’s commander. The Commander variant was created and popularized by fans; an independent rules committee maintains additional resources at MTGCommander.net. The Commander variant uses all the normal rules for a Magic game, with the following additions.

903.2. A Commander game may be a two-player game or a multiplayer game. The default multiplayer setup is the Free-for-All variant with the attack multiple players option and without the limited range of influence option. See rule 806, “Free-for-All Variant.”

903.3. Each deck has a legendary creature card designated as its commander. This designation is not a characteristic of the object represented by the card; rather, it is an attribute of the card itself. The card retains this designation even when it changes zones.
Example: A commander that’s been turned face down (due to Ixidron’s effect, for example) is still a commander. A commander that’s copying another card (due to Cytoshape’s effect, for example) is still a commander. A permanent that’s copying a commander (such as a Body Double, for example, copying a commander in a player’s graveyard) is not a commander.`,
		},
	}

	for i := range expected_rules {
		if rules[i] != expected_rules[i] {
			t.Errorf("expected=%q, got=%q", expected_rules[i], rules[i])
		}
	}
}

func TestSubRuleParser(t *testing.T) {
	example := `903.2. A Commander game may be a two-player game or a multiplayer game. The default multiplayer setup is the Free-for-All variant with the attack multiple players option and without the limited range of influence option. See rule 806, “Free-for-All Variant.”

903.2a Testing

903.2b Testing Again

903.3. each deck has a legendary creature card designated as its commander. this designation is not a characteristic of the object represented by the card; rather, it is an attribute of the card itself. the card retains this designation even when it changes zones.
example: a commander that’s been turned face down (due to ixidron’s effect, for example) is still a commander. a commander that’s copying another card (due to cytoshape’s effect, for example) is still a commander. a permanent that’s copying a commander (such as a body double, for example, copying a commander in a player’s graveyard) is not a commander.`
	rules := parseSubRules(example)

	if len(rules) != 2 {
		t.Errorf("expected 2 rules, got=%d", len(rules))
	}
	expected_rules := []Rule{
		{
			Code: "903.2.",
			Text: `A Commander game may be a two-player game or a multiplayer game. The default multiplayer setup is the Free-for-All variant with the attack multiple players option and without the limited range of influence option. See rule 806, “Free-for-All Variant.”

903.2a Testing

903.2b Testing Again
`,
		},
		{
			Code: "903.3.",
			Text: `each deck has a legendary creature card designated as its commander. this designation is not a characteristic of the object represented by the card; rather, it is an attribute of the card itself. the card retains this designation even when it changes zones.
example: a commander that’s been turned face down (due to ixidron’s effect, for example) is still a commander. a commander that’s copying another card (due to cytoshape’s effect, for example) is still a commander. a permanent that’s copying a commander (such as a body double, for example, copying a commander in a player’s graveyard) is not a commander.`,
		},
	}

	for i := range expected_rules {
		if rules[i] != expected_rules[i] {
			t.Errorf("expected=%q, got=%q", expected_rules[i], rules[i])
		}
	}
}

func TestKeywordParser(t *testing.T) {
	example := `Vigilance
A keyword ability that lets a creature attack without tapping. See rule 702.20, “Vigilance.”

Visit
A keyword ability found on Attraction cards. It provides an effect whenever you roll to visit your attractions and get certain results. See rule 702.159, “Visit.”

Vote
Some cards instruct players to vote from among given options. See rule 701.32, “Vote.”

Walker Token
A Walker token is a 2/2 black Zombie creature token named Walker. For more information on predefined tokens, see rule 111.10.
`
	keywords := parseKeywords(example)

	if len(keywords) != 4 {
		t.Errorf("expected 4 keywords, got=%d", len(keywords))
	}
	expected_kws := []Keyword{
		{
			Title: "Vigilance",
			Text:  "A keyword ability that lets a creature attack without tapping. See rule 702.20, “Vigilance.”",
		},
		{
			Title: "Visit",
			Text:  "A keyword ability found on Attraction cards. It provides an effect whenever you roll to visit your attractions and get certain results. See rule 702.159, “Visit.”",
		},
		{
			Title: "Vote",
			Text:  "Some cards instruct players to vote from among given options. See rule 701.32, “Vote.”",
		},
		{
			Title: "Walker Token",
			Text:  "A Walker token is a 2/2 black Zombie creature token named Walker. For more information on predefined tokens, see rule 111.10.",
		},
	}

	for i := range expected_kws {
		if keywords[i] != expected_kws[i] {
			t.Errorf("expected=%q, got=%q", expected_kws[i], keywords[i])
		}
	}
}
