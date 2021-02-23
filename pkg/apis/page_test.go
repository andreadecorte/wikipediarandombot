package apis

import (
	"io/ioutil"
	"strings"
	"testing"
)

// TestParsing uses a static input file to test
// parsing of the results
func TestParsingPage(t *testing.T) {
	testFile := "_testdata/fulltext.json"
	b, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Error opening the file %s", testFile)
	}
	err2, result := parseJsonPage(b)
	if err2 != nil {
		t.Fatalf("Error parsing json: %s", err2)
	}
	index := strings.Index(result, "third")
	if index != 13 {
		t.Fatalf("Word not found: %d", index)
	}
	count := strings.Count(result, "is")
	if count != 345 {
		t.Fatalf("Count error: %d", count)
	}

	words := wordCount(result) 
	if words != 8076 {
		t.Fatalf("Count error: %d", words)
	}
}
