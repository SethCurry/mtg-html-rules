package ruleparser

import (
	"slices"
	"strings"
)

type Introduction struct {
	EffectiveDate string
	Body          string
}

func ParseIntroduction(intro []string) (*Introduction, error) {
	dateLineIndex := slices.IndexFunc(intro, func(s string) bool {
		return strings.Contains(s, "effective as of")
	})
	splitDateLine := strings.Split(intro[dateLineIndex], "effective as of")

	bodyIndex := slices.IndexFunc(intro[dateLineIndex:], func(s string) bool {
		return strings.HasPrefix(s, "This document")
	})

	return &Introduction{
		EffectiveDate: strings.TrimRight(strings.TrimSpace(splitDateLine[1]), "."),
		Body:          strings.Join(intro[bodyIndex+dateLineIndex:], "\n"),
	}, nil
}
