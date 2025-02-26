package ruleparser

import (
	"encoding/json"
	"errors"
	"strings"
)

func NewRules() *Rules {
	return &Rules{
		Sections: make(map[int]*Section),
	}
}

type Rules struct {
	Sections map[int]*Section `json:"sections"`
}

func newParser() *ruleParser {
	return &ruleParser{
		currentSection:    nil,
		currentSubsection: nil,
		currentRule:       nil,
		lastExampler:      nil,
		rules:             NewRules(),
	}
}

type ruleParser struct {
	// currentSection stores a pointer to the last section header that was parsed.
	// This allows subsections and rules to be added to them as that section is being parsed.
	currentSection *Section

	// currentSubsection stores a pointer to the last subsection header that was parsed.
	// This allows rules to be added to it as that subsection is being parsed.
	currentSubsection *Subsection

	// currentRule stores a pointer to the last rule header that was parsed.
	// This allows subrules to be added to it as that rule is being parsed.
	currentRule *Rule

	// lastExampler stores a pointer to the last rule or subrule that was parsed.
	// This allows examples to be added to it as that rule or subrule is being parsed.
	lastExampler Exampler

	rules *Rules
}

func (r *ruleParser) parseSubsection(line string) error {
	num, subsect, err := parseSubsectionLine(line)
	if err != nil {
		return err
	}

	r.currentSubsection = subsect
	r.currentSection.Subsections[num] = subsect

	return nil
}

func (r *ruleParser) parseSection(line string) error {
	num, sect, err := parseSectionLine(line)
	if err != nil {
		return err
	}

	r.currentSection = sect
	r.rules.Sections[num] = sect

	return nil
}

func (r *ruleParser) parseRule(line string) error {
	num, gotRule, err := parseRuleLine(line)
	if err != nil {
		return err
	}

	r.currentRule = gotRule
	r.currentSubsection.Rules[num] = gotRule
	r.lastExampler = gotRule

	return nil
}

func (r *ruleParser) parseSubrule(line string) error {
	l, subrule, err := parseSubruleLine(line)
	if err != nil {
		return err
	}

	r.currentRule.Subrules[l] = subrule
	r.lastExampler = subrule

	return nil
}

func (r *ruleParser) parseExample(line string) {
	example := parseExample(line)

	r.lastExampler.AddExample(example)
}

func (r *ruleParser) handleLine(line string, origLine string) error {
	switch {
	case len(line) == 0:
		return nil
	case isSection(line):
		return r.parseSection(line)
	case isSubsection(line):
		return r.parseSubsection(line)
	case isRule(line):
		return r.parseRule(line)
	case isSubrule(line):
		return r.parseSubrule(line)
	case isExample(line):
		r.parseExample(line)
		return nil
	case strings.HasPrefix(origLine, "    ") || strings.HasPrefix(origLine, "\n"):
		if r.lastExampler != nil {
			r.lastExampler.AddToContents(line)
		} else {
			return errors.New("unexpected content line")
		}
		return nil
	default:
		return errors.New("unknown line type")
	}
}

func ParseRules(cr []string) (*Rules, error) {
	parser := newParser()
	lineNumber := 0
	for i := 0; i < len(cr); i++ {
		origLine := convertEncoding(cr[i])
		line := strings.TrimSpace(origLine)
		origLine = strings.Trim(origLine, "\n")
		err := parser.handleLine(line, origLine)
		if err != nil {
			origLineJSON, jsonErr := json.Marshal(origLine)
			if jsonErr != nil {
				return nil, &ParseError{
					Line:       origLine,
					LineNumber: lineNumber,
					Err:        err,
				}
			}
			return nil, &ParseError{
				Line:       string(origLineJSON),
				LineNumber: lineNumber,
				Err:        err,
			}
		}
	}
	return parser.rules, nil
}
