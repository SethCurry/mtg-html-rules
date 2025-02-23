package ruleparser

import (
	"slices"
	"strings"
)

type Introduction struct {
	EffectiveDate    string
	BodyPreDownload  string
	BodyPostDownload string
}

func ParseIntroduction(intro []string) (*Introduction, error) {
	dateLineIndex := slices.IndexFunc(intro, func(s string) bool {
		return strings.Contains(s, "effective as of")
	})

	splitDateLine := strings.Split(intro[dateLineIndex], "effective as of")

	pre := slices.IndexFunc(intro[dateLineIndex:], func(s string) bool {
		return strings.HasPrefix(s, "This ")
	})

	post := slices.IndexFunc(intro[pre:], func(s string) bool {
		return strings.HasPrefix(s, "Changes ")
	})

	return &Introduction{
		EffectiveDate: strings.TrimRight(strings.TrimSpace(splitDateLine[1]), "."),
		// modify index offsets here since we searched progressively smaller slices to find them
		BodyPreDownload:  intro[pre+dateLineIndex],
		BodyPostDownload: intro[post+pre],
	}, nil
}
