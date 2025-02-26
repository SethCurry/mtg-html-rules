package ruleparser

func NewStringScanner(data string) *StringScanner {
	return &StringScanner{
		index: 0,
		data:  data,
	}
}

type StringScanner struct {
	index int
	data  string
}

func (s *StringScanner) Next() (rune, bool) {
	data := []rune(s.data)
	if s.index >= len(data) {
		return 0, false
	}

	c := data[s.index]
	s.index++

	return c, true
}

func (s *StringScanner) ReadUntil(c []rune) string {
	acc := ""
	data := []rune(s.data)
	for s.index < len(data) {
		for _, r := range c {
			if rune(data[s.index]) == r {
				return acc
			}
		}

		acc += string(data[s.index])
		s.index++
	}

	return acc
}
