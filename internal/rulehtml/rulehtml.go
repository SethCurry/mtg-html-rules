package rulehtml

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/SethCurry/mtg-html-rules/pkg/ruleparser"
)

func getElementID(elementName string) (string, error) {
	return strings.Replace(elementName, ".", "-", -1), nil
}

type templateData struct {
	Introduction *ruleparser.Introduction
	Rules        *ruleparser.Rules
	Glossary     []*ruleparser.GlossaryEntry
	Credits      []string
}

func GenerateTemplate(parsedDoc *ruleparser.ParsedDocument, toWriter io.Writer) error {
	parsedTemplate, err := template.New("rules.tmpl").Funcs(template.FuncMap{
		"ElementID": getElementID,
		"Split":     func(s, sep string) []string { return strings.Split(s, sep) },
		"Index":     func(s, subStr string) int { return strings.Index(s, subStr) },
		"Contains":  func(s, substr string) bool { return strings.Contains(s, substr) },
		"Replace":   func(s, old, new string, n int) string { return strings.Replace(s, old, new, n) },
	}).Parse(rootTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	data := templateData{
		Introduction: parsedDoc.Introduction,
		Rules:        parsedDoc.Rules,
		Glossary:     parsedDoc.Glossary,
		Credits:      parsedDoc.Credits,
	}

	err = parsedTemplate.Execute(toWriter, &data)
	if err != nil {
		return fmt.Errorf("failed to execute rules template: %w", err)
	}

	return nil
}
