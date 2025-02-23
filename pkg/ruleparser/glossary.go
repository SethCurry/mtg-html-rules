package ruleparser

type GlossaryEntry struct {
	Term       []*ContentElement   `json:"term"`
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
				Term:       workingSlice[0],
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
