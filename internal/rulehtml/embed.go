package rulehtml

import _ "embed"

//go:embed web/vendor/mana/mana.css
var manaCSS string

//go:embed web/embedded-fonts.css
var embeddedFontsCSS string

//go:embed web/main.css
var mainCSS string

//go:embed rules.tmpl
var rootTemplate string
