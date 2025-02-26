package ruleparser

func ParseCredits(cr []string) ([]string, int) {
	creditsChunk := make([]string, 0)
	for i := len(cr) - 1; i >= 0; i-- {
		if cr[i] == "Credits" {
			return creditsChunk, i
		} else {
			creditsChunk = prepend(creditsChunk, cr[i])
		}
	}
	return nil, -1
}
