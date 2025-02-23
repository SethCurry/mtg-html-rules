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

	return &Introduction{
		EffectiveDate:    strings.TrimRight(strings.TrimSpace(splitDateLine[1]), "."),
		BodyPreDownload:  intro[3],
		BodyPostDownload: intro[4],
	}, nil
}
