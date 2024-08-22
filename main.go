package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SethCurry/mtg-html-rules/internal/rulehtml"
	"github.com/SethCurry/mtg-html-rules/pkg/ruleparser"
)

func usage() {
	fmt.Println("mtg-html-rules <input file> [--output <output file>]")
}

func main() {
	// position argument for input file
	// optional flag for output file

	outputFilePath := flag.String("output", "./mtgcr.html", "The path to write the HTML rules file to.")

	flag.Parse()

	if len(os.Args) < 1 || len(os.Args) > 2 {
		usage()
		os.Exit(1)
	}

	inputFilePath := os.Args[0]

	parsedRules, err := ruleparser.ParseFile(inputFilePath)
	if err != nil {
		fmt.Printf("failed to parse rules from file %q:\n%s\n", inputFilePath, err.Error())
		os.Exit(1)
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
