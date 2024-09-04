package web

import _ "embed"

//go:generate go run github.com/evanw/esbuild/cmd/esbuild ./javascript/main.js --bundle --minify --outfile=bundle.js

//go:embed bundle.js
var BundleJS string
