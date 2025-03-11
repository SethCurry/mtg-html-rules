package ruleparser

import (
	"regexp"
	"strings"
)

type GlossaryEntry struct {
	Term       string              `json:"term"`
	Definition [][]*ContentElement `json:"examples"`
}

func ParseGlossary(cr []string) ([]*GlossaryEntry, int) {
	workingSlice := make([][]*ContentElement, 0)
	entries := make([]*GlossaryEntry, 0)
	for i := len(cr) - 1; i >= 0; i-- {
		// skip any extra blank lines at the end of the glossary section
		if len(workingSlice) == 0 && len(cr[i]) == 0 {
			continue
		}

		if len(cr[i]) > 0 {
			content := parseContent(cr[i])
			workingSlice = prepend(workingSlice, content)
		} else {
			entries = prepend(entries, &GlossaryEntry{
				Term:       workingSlice[0][0].Value,
				Definition: workingSlice[1:],
			})
			workingSlice = nil
		}

		if cr[i] == "Glossary" {
			return entries, i
		}
	}
	return entries, -1
}

func LinkGlossaryReferences(entries []*GlossaryEntry) {
	// https://regex101.com/r/Y2gKfX/2
	// TODO: I emailed Del to see about getting the one quoted reference un-quoted
	// so maybe we can simplify this in the future
	refRegex := regexp.MustCompile(`See (?:also )?"?((?:[A-Z][^,\n]*(?:, )?)+?)\."?`)

	for _, currentTerm := range entries {
		for j := range currentTerm.Definition {
			updatedDefinitionContent := []*ContentElement{}
			for _, element := range currentTerm.Definition[j] {
				currentDefSlice := element.Value

				if element.Type == ContentText {
					refIndices := refRegex.FindAllStringSubmatchIndex(currentDefSlice, -1)

					if len(refIndices) > 0 {
						updatedDefinitionContent = append(updatedDefinitionContent, &ContentElement{
							Type:  ContentText,
							Value: currentDefSlice[:refIndices[0][2]],
						})
						str := currentDefSlice[refIndices[0][2]:refIndices[0][3]]

						references := strings.Split(str, ",")
						// this definition has a comma in it, so un-split it
						if str == "Active Player, Nonactive Player Order" {
							references[0] = str
							references = references[:1]
						}
						subStringIndices := make([][]int, 0)
						for _, ref := range references {
							ref = strings.TrimSpace(ref)
							index := strings.Index(str, ref)
							subStringIndices = append(subStringIndices, []int{index, index + len(ref)})
						}
						for idx := range subStringIndices {
							updatedDefinitionContent = append(updatedDefinitionContent, &ContentElement{
								Type:  ContentGlossaryReference,
								Value: strings.TrimSpace(references[idx]),
							})
							if len(subStringIndices) > idx+1 {
								updatedDefinitionContent = append(updatedDefinitionContent, &ContentElement{
									Type:  ContentText,
									Value: str[subStringIndices[idx][1]:subStringIndices[idx+1][0]],
								})
							}
						}
						// finally, append everything after the reference(s). This is likely just a full stop.
						updatedDefinitionContent = append(updatedDefinitionContent, &ContentElement{
							Type:  ContentText,
							Value: currentDefSlice[refIndices[0][3]:],
						})
					} else {
						updatedDefinitionContent = append(updatedDefinitionContent, element)
					}
				} else {
					updatedDefinitionContent = append(updatedDefinitionContent, element)
				}
			}
			currentTerm.Definition[j] = updatedDefinitionContent
		}
	}
}
