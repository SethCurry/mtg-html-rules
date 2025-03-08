package ruleparser

import (
	"regexp"
)

// ContentType represents the type of content in a rule, such as text,
// a reference to another rule, or a symbol.
type ContentType string

const (
	// ContentText represents a plain text content element.
	ContentText ContentType = "text"

	// ContentSymbol represents a symbol in the contents,
	// such as a mana symbol or the tap symbol.
	ContentSymbol ContentType = "symbol"

	// ContentRuleReference represents a reference to another rule.
	ContentRuleReference ContentType = "ruleReference"

	// ContentSectionReference represents a reference to a full section.
	ContentSectionReference ContentType = "sectionReference"
)

type ContentElement struct {
	Type  ContentType `json:"type"`
	Value string      `json:"value"`
}

func parseContent(content string) []*ContentElement {
	// https://regex101.com/r/rNco94/7
	refRegex := regexp.MustCompile(`(section \d, “.+”)|(\d{3}(?:\.\d{1,3}[a-z]?)?)(?:–[a-z])?`)
	elements := []*ContentElement{}
	refIndices := refRegex.FindAllStringSubmatchIndex(content, -1)

	lastIndex := 0
	for i := range refIndices {
		// Section reference
		if refIndices[i][2] != -1 {
			elements = append(elements, &ContentElement{
				Type:  ContentText,
				Value: content[lastIndex:refIndices[i][2]],
			})

			elements = append(elements, &ContentElement{
				Type:  ContentSectionReference,
				Value: content[refIndices[i][2]:refIndices[i][3]],
			})

			lastIndex = refIndices[i][3]
			// Rule reference
		} else {
			elements = append(elements, &ContentElement{
				Type:  ContentText,
				Value: content[lastIndex:refIndices[i][0]],
			})

			elements = append(elements, &ContentElement{
				Type:  ContentRuleReference,
				Value: content[refIndices[i][0]:refIndices[i][1]],
			})

			lastIndex = refIndices[i][1]
		}
	}

	if len(content[lastIndex:]) > 0 {
		elements = append(elements, &ContentElement{
			Type:  ContentText,
			Value: content[lastIndex:],
		})
	}
	return elements
}
