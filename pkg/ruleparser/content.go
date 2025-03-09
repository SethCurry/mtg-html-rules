package ruleparser

import (
	"regexp"
	"strings"
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

	// ContentUrlReference represents a reference to a full section.
	ContentUrlReference ContentType = "urlReference"
)

type ContentElement struct {
	Type  ContentType `json:"type"`
	Value string      `json:"value"`
}

func parseContent(content string) []*ContentElement {
	// https://regex101.com/r/rNco94/9
	refRegex := regexp.MustCompile(`([a-zA-Z0-9/.-]+(?:\.com|\.net)[a-zA-Z0-9/-]*)|(section \d, “.+”)|(\d{3}(?:\.\d{1,3}[a-z]?)?)(?:–[a-z])?`)
	elements := []*ContentElement{}
	refIndices := refRegex.FindAllStringSubmatchIndex(content, -1)

	lastIndex := 0
	for i := range refIndices {
		// Section reference
		if refIndices[i][4] != -1 {
			elements = append(elements, &ContentElement{
				Type:  ContentText,
				Value: content[lastIndex:refIndices[i][4]],
			})

			elements = append(elements, &ContentElement{
				Type:  ContentSectionReference,
				Value: content[refIndices[i][4]:refIndices[i][5]],
			})

			lastIndex = refIndices[i][5]
			// Rule reference
		} else {
			elements = append(elements, &ContentElement{
				Type:  ContentText,
				Value: content[lastIndex:refIndices[i][0]],
			})

			parsedReference := content[refIndices[i][0]:refIndices[i][1]]
			contentType := ContentRuleReference
			if strings.Contains(parsedReference, ".com") || strings.Contains(parsedReference, ".net") {
				contentType = ContentUrlReference
			}

			elements = append(elements, &ContentElement{
				Type:  contentType,
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
