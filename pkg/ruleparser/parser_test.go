package ruleparser

import (
	"os"
	"testing"
)

func Test_ParseFile(t *testing.T) {
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

			if len(parsed.Sections) != 9 {
				t.Fatalf("expected 9 sections, got %d", len(parsed.Sections))
			}
		})
	}
}
