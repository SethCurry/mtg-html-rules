package ruleparser

import (
	"fmt"
	"io"
	"os"
	"slices"
)

type ParsedDocument struct {
	Introduction *Introduction
	Glossary     []*GlossaryEntry
	Rules        *Rules
	Credits      []string
}

func DispatchDocument(doc *ParsedDocument, reader io.Reader) error {
	cr, err := ReadAllLines(reader, false)
	if err != nil {
		panic(err)
	}
	// This all looks a little goofy, but basically we know the intro is at the front of the doc
	// and the glossary/credits are at the back, so we can build & index those first,
	// then whatever's left in the middle must be rules
	firstContentsHit := slices.Index(cr, "Contents")
	introRange := cr[:firstContentsHit]
	doc.Introduction, err = ParseIntroduction(introRange)
	if err != nil {
		panic("Failed to parse the introduction")
	}

	firstCreditsHit := slices.Index(cr, "Credits")

	credits, lastCreditsHit := ParseCredits(cr)
	if lastCreditsHit == -1 {
		panic("Failed to parse the credits")
	}
	doc.Credits = credits

	glossaryEntries, lastGlossaryHit := ParseGlossary(cr[:lastCreditsHit])
	doc.Glossary = glossaryEntries

	doc.Rules, err = ParseRules(cr[firstCreditsHit+1 : lastGlossaryHit])
	if err != nil {
		return err
	}
	return nil
}

func ParseFile(path string) (*ParsedDocument, error) {
	fileDescriptor, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open rules file: %w", err)
	}
	defer fileDescriptor.Close()
	return ParseCrDoc(fileDescriptor)
}

func ParseCrDoc(reader io.Reader) (*ParsedDocument, error) {
	doc := &ParsedDocument{}
	DispatchDocument(doc, reader)
	return doc, nil
}

type ParseError struct {
	Line       string
	LineNumber int
	Err        error
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("error parsing line %d: %s: %v", p.LineNumber, p.Line, p.Err)
}
