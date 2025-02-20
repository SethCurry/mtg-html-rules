package rulehtml

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/SethCurry/gofall"
)

type CardReference struct {
	Name     string `json:"name"`
	Printing string `json:"printing"`
}

type RuleReference struct {
	Section    int     `json:"section"`
	Subsection int     `json:"subsection"`
	Rule       int     `json:"rule"`
	Subrule    *string `json:"subrule"`
}

func DecomposeRuleReference(reference string) (*RuleReference, error) {
	parts := strings.Split(reference, ".")

	subsection, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	section := subsection / 100

	ruleStr := parts[1]

	ref := &RuleReference{
		Section:    section,
		Subsection: subsection,
	}

	matched, err := regexp.MatchString(`^\d+$`, ruleStr)
	if err != nil {
		return nil, err
	}

	if matched {
		rule, err := strconv.Atoi(ruleStr)
		if err != nil {
			return nil, err
		}

		ref.Rule = rule
	} else {
		ruleNum := ruleStr[:len(ruleStr)-1]
		subrule := string(ruleStr[len(ruleStr)-1])

		rule, err := strconv.Atoi(ruleNum)
		if err != nil {
			return nil, err
		}

		ref.Rule = rule
		ref.Subrule = &subrule
	}

	return ref, nil
}

type CardReferences map[string][]CardReference

func cardReferenceToImageURL(ctx context.Context, name string, client *gofall.Client) (string, error) {
	card, err := client.Card.Named(ctx, name)
	if err != nil {
		return "", err
	}

	return card.ImageURIs.HighestQuality()
}
