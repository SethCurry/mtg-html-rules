package rulehtml

import _ "embed"

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

//go:embed web/main.js
var mainJS string

//go:embed web/vendor/ufuzzy.min.js
var ufuzzyJS string

//go:embed web/main.css
var mainCSS string

//go:embed rules.tmpl
var rootTemplate string
