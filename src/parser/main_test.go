package parser

import (
	"testing"
)

func TestRuleParser(t *testing.T) {
	example := `902.7. Any abilities of a face-up vanguard card in the command zone function from that zone. The card’s static abilities affect the game, its triggered abilities may trigger, and its activated abilities may be activated.

903. Commander

903.1. In the Commander variant, each deck is led by a legendary creature designated as that deck’s commander. The Commander variant was created and popularized by fans; an independent rules committee maintains additional resources at MTGCommander.net. The Commander variant uses all the normal rules for a Magic game, with the following additions.

903.2. A Commander game may be a two-player game or a multiplayer game. The default multiplayer setup is the Free-for-All variant with the attack multiple players option and without the limited range of influence option. See rule 806, “Free-for-All Variant.”

903.3. Each deck has a legendary creature card designated as its commander. This designation is not a characteristic of the object represented by the card; rather, it is an attribute of the card itself. The card retains this designation even when it changes zones.
Example: A commander that’s been turned face down (due to Ixidron’s effect, for example) is still a commander. A commander that’s copying another card (due to Cytoshape’s effect, for example) is still a commander. A permanent that’s copying a commander (such as a Body Double, for example, copying a commander in a player’s graveyard) is not a commander.`
	rules := parseRules(example)

	if len(rules) != 5 {
		t.Errorf("expected 5 rules, got=%d", len(rules))
	}
	expected_rules := []Rule{
		{
			Code: "902.7.",
			Text: "Any abilities of a face-up vanguard card in the command zone function from that zone. The card’s static abilities affect the game, its triggered abilities may trigger, and its activated abilities may be activated.\n",
		},
		{
			Code: "903.",
			Text: "Commander\n",
		},
		{
			Code: "903.1.",
			Text: "In the Commander variant, each deck is led by a legendary creature designated as that deck’s commander. The Commander variant was created and popularized by fans; an independent rules committee maintains additional resources at MTGCommander.net. The Commander variant uses all the normal rules for a Magic game, with the following additions.\n",
		},
		{
			Code: "903.2.",
			Text: "A Commander game may be a two-player game or a multiplayer game. The default multiplayer setup is the Free-for-All variant with the attack multiple players option and without the limited range of influence option. See rule 806, “Free-for-All Variant.”\n",
		},
		{
			Code: "903.3.",
			Text: "Each deck has a legendary creature card designated as its commander. This designation is not a characteristic of the object represented by the card; rather, it is an attribute of the card itself. The card retains this designation even when it changes zones.\nExample: A commander that’s been turned face down (due to Ixidron’s effect, for example) is still a commander. A commander that’s copying another card (due to Cytoshape’s effect, for example) is still a commander. A permanent that’s copying a commander (such as a Body Double, for example, copying a commander in a player’s graveyard) is not a commander.",
		},
	}

	for i := range rules {
		if rules[i] != expected_rules[i] {
			t.Errorf("expected=%q, got=%q", expected_rules[i], rules[i])
		}
	}
}
