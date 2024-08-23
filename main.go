package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/SethCurry/mtg-html-rules/internal/rulehtml"
	"github.com/SethCurry/mtg-html-rules/pkg/ruleparser"
)

func usage() {
	fmt.Println("mtg-html-rules [--rules <rules file>] [--output <output file>]")
}

func main() {
	outputFilePath := flag.String("output", "./mtgcr.html", "The path to write the HTML rules file to.")
	rulesFilePath := flag.String("rules", "", "The path to the rules file to parse.")
	isHelp := flag.Bool("help", false, "Print the help text and exit.")

	flag.Parse()

	if *isHelp {
		usage()
		os.Exit(0)
	}

	var parsedRules *ruleparser.Rules

	if *rulesFilePath == "" {
		rulesURL, err := ruleparser.GetLatestRulesTxtURL()
		if err != nil {
			fmt.Printf("no rules file provided with --rules flag, and failed to find the URL for the latest rules:\n%s\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Downloading rules from %q\n", rulesURL)

		resp, err := http.Get(rulesURL)
		if err != nil {
			fmt.Printf("failed to download rules from %q:\n%s\n", rulesURL, err.Error())
			os.Exit(1)
		}
		defer resp.Body.Close()

		parsedRules, err = ruleparser.ParseRules(resp.Body)
		if err != nil {
			fmt.Printf("failed to parse rules from %q:\n%s\n", rulesURL, err.Error())
			os.Exit(1)
		}
	} else {
		var err error

		parsedRules, err = ruleparser.ParseFile(*rulesFilePath)
		if err != nil {
			fmt.Printf("failed to parse rules from file %q:\n%s\n", *rulesFilePath, err.Error())
			os.Exit(1)
		}
	}

	outFd, err := os.Create(*outputFilePath)
	if err != nil {
		fmt.Printf("failed to open output file %q:\n%s\n", *outputFilePath, err.Error())
		os.Exit(1)
	}
	defer outFd.Close()

	err = rulehtml.GenerateTemplate(parsedRules, outFd)
	if err != nil {
		fmt.Printf("failed to generate templated HTML from parsed rules:\n%s\n", err.Error())
		os.Exit(1)
	}
}
