# mtg-html-rules

[![Lint](https://github.com/SethCurry/mtg-html-rules/actions/workflows/lint.yml/badge.svg)](https://github.com/SethCurry/mtg-html-rules/actions/workflows/lint.yml)
[![Test](https://github.com/SethCurry/mtg-html-rules/actions/workflows/go-test.yml/badge.svg)](https://github.com/SethCurry/mtg-html-rules/actions/workflows/go-test.yml)

`mtg-html-rules` is a tool for converting the comprehensive rules for Magic: The Gathering
into a single HTML file with anchors for linking to an individual rule.

There are 2 primary goals for the project:

1. Make the rules more accessible to players by adding features to the rules page and allowing judges to link to specific rules.

2. To allow players and judges with poor internet connections to access the rules.

## Installation

Binaries are located on the [releases page](https://github.com/SethCurry/mtg-html-rules/releases).

## Usage

First, download a copy of the latest comprehensive rules from the [Wizards of the Coast website](https://magic.wizards.com/en/rules).

Then, run the following command:

```bash
mtg-html-rules rules-file.txt
```

It will generate single HTML file named `mtgcr.html` in the current directory.

If you want to output the file to a different path, use the `--output` flag:

```bash
mtg-html-rules rules-file.txt --output=/path/to/output/rules-file.txt
```
