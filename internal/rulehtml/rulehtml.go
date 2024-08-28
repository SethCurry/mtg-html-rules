package rulehtml

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/SethCurry/mtg-html-rules/pkg/ruleparser"
)

//go:embed web/vendor/mana/mana.css
var manaCSSTemplate string

//go:embed web/vendor/mana/mana.eot
var manaFontEOT string

//go:embed web/vendor/mana/mana.svg
var manaFontSVG string

//go:embed web/vendor/mana/mana.ttf
var manaFontTTF string

//go:embed web/vendor/mana/mana.woff
var manaFontWOFF string

//go:embed web/vendor/mana/mplantin.eot
var mplantinFontEOT string

//go:embed web/vendor/mana/mplantin.svg
var mplantinFontSVG string

//go:embed web/vendor/mana/mplantin.ttf
var mplantinFontTTF string

//go:embed web/vendor/mana/mplantin.woff
var mplantinFontWOFF string

//go:embed web/menu.js
var menuJS string

//go:embed web/vendor/ufuzzy.min.js
var ufuzzyJS string

//go:embed web/main.css
var mainCSS string

type encodedFont struct {
	EOT  string
	SVG  string
	TTF  string
	WOFF string
}

type fontData struct {
	Mana     encodedFont
	MPlantin encodedFont
}

func newFontData() fontData {
	return fontData{
		Mana: encodedFont{
			EOT:  base64.RawStdEncoding.EncodeToString([]byte(manaFontEOT)),
			SVG:  base64.RawStdEncoding.EncodeToString([]byte(manaFontSVG)),
			TTF:  base64.RawStdEncoding.EncodeToString([]byte(manaFontTTF)),
			WOFF: base64.RawStdEncoding.EncodeToString([]byte(manaFontWOFF)),
		},
		MPlantin: encodedFont{
			EOT:  base64.RawStdEncoding.EncodeToString([]byte(mplantinFontEOT)),
			SVG:  base64.RawStdEncoding.EncodeToString([]byte(mplantinFontSVG)),
			TTF:  base64.RawStdEncoding.EncodeToString([]byte(mplantinFontTTF)),
			WOFF: base64.RawStdEncoding.EncodeToString([]byte(mplantinFontWOFF)),
		},
	}
}

//go:embed rules.tmpl
var rootTemplate string

func manaSymbolToClass(symbol string) (string, error) {
	switch symbol {
	case "t", "T":
		return "ms-tap ms-cost ms-shadow", nil
	case "q", "Q":
		return "ms-untap ms-cost ms-shadow", nil
	case "rn", "rN":
		return "ms-saga", nil
	case "rn1", "rN1":
		return "ms-saga ms-saga-1", nil
	case "rn2", "rN2":
		return "ms-saga ms-saga-2", nil
	case "rn3", "rN3":
		return "ms-saga ms-saga-3", nil
	case "rn4", "rN4":
		return "ms-saga ms-saga-4", nil
	}
	return "ms-cost ms-shadow ms-" + strings.Replace(strings.ToLower(symbol), "/", "", -1), nil
}

func getElementID(elementName string) (string, error) {
	return strings.Replace(elementName, ".", "_", -1), nil
}

type templateData struct {
	Rules    *ruleparser.Rules
	ManaCSS  template.CSS
	MenuJS   template.JS
	UFuzzyJS template.JS
	MainCSS  template.CSS
}

func generateManaCSS() (string, error) {
	parsedTemplate, err := template.New("mana.css").Parse(manaCSSTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse mana.css template: %w", err)
	}

	var manaCSS bytes.Buffer

	err = parsedTemplate.Execute(&manaCSS, newFontData())
	if err != nil {
		return "", fmt.Errorf("failed to execute mana.css template: %w", err)
	}

	return manaCSS.String(), nil
}

func GenerateTemplate(parsedRules *ruleparser.Rules, toWriter io.Writer) error {
	parsedTemplate, err := template.New("rules.tmpl").Funcs(template.FuncMap{
		"ManaClass": manaSymbolToClass,
		"ElementID": getElementID,
	}).Parse(rootTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	manaCSS, err := generateManaCSS()
	if err != nil {
		return fmt.Errorf("failed to generate mana.css: %w", err)
	}

	data := templateData{
		Rules:    parsedRules,
		ManaCSS:  template.CSS(manaCSS),
		MenuJS:   template.JS(menuJS),
		UFuzzyJS: template.JS(ufuzzyJS),
		MainCSS:  template.CSS(mainCSS),
	}

	err = parsedTemplate.Execute(toWriter, &data)
	if err != nil {
		return fmt.Errorf("failed to execute rules template: %w", err)
	}

	return nil
}
