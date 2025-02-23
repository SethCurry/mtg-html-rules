package ruleparser

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func Test_ParseRulesSections(t *testing.T) {
	testFiles, err := os.ReadDir("testdata/")
	if err != nil {
		t.Fatalf("failed to list test data files: %v", err)
	}

	for _, v := range testFiles {
		t.Run("parse-"+v.Name(), func(t *testing.T) {
			parsed, err := ParseFile("testdata/" + v.Name())
			if err != nil {
				t.Fatalf("failed to parse file: %v", err)
			}

			if len(parsed.Rules.Sections) != 9 {
				t.Fatalf("expected 9 sections, got %d", len(parsed.Rules.Sections))
			}
		})
	}
}

type introTest struct {
	input string
	want  string
}

func Test_ParseIntroduction(t *testing.T) {
	_, err := os.ReadDir("testdata/")
	if err != nil {
		t.Fatalf("failed to list test data files: %v", err)
	}

	tests := []introTest{
		{input: "20240308.txt", want: "March 8, 2024"},
		{input: "20240410.txt", want: "April 12, 2024"},
	}
	for _, testContext := range tests {
		parsed, err := ParseFile("testdata/" + testContext.input)
		if err != nil {
			t.Fatalf("failed to parse file: %v", err)
		}
		if parsed.Introduction.EffectiveDate != testContext.want {
			t.Fatalf("EffectiveDate did not parse correctly: wanted %s, got %s", testContext.want, parsed.Introduction.EffectiveDate)
		}
		intro, _ := json.Marshal(parsed.Introduction)
		fmt.Printf("%s", intro)
		if len(parsed.Introduction.BodyPreDownload) == 0 || len(parsed.Introduction.BodyPostDownload) == 0 {
			t.Fatalf("Didn't parse the right number of paragraphs from the introduction")
		}
	}
}
