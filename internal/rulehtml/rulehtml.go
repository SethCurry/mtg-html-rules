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
	Rules 	*ruleparser.Rules
}

func GenerateTemplate(parsedRules *ruleparser.Rules, toWriter io.Writer) error {
	parsedTemplate, err := template.New("rules.tmpl").Funcs(template.FuncMap{
		"ElementID": getElementID,
	}).Parse(rootTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	data := templateData{
		Rules:	parsedRules,
	}

	err = parsedTemplate.Execute(toWriter, &data)
	if err != nil {
		return fmt.Errorf("failed to execute rules template: %w", err)
	}

	return nil
}
