package apis

import (
	"io/ioutil"
	"testing"
)

// TestHowManyPages checks that we return the number of
// pages we expect
func TestHowManyPages(t *testing.T) {
	var pages Pages
	howMany := 1
	lang := "en"
	err := GetWiki(&pages, lang, howMany)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	if len(pages.Items) != howMany {
		t.Fatalf("We should return %d pages, but got %d", howMany, len(pages.Items))
	}
}

// TestSwitchLang verifies that we are getting results
// in the desidered language
func TestSwitchLang(t *testing.T) {
	var pages Pages
	howMany := 2
	lang := "fur"
	err := GetWiki(&pages, lang, howMany)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}	
	if len(pages.Items) != howMany {
		t.Fatalf("We should return %d pages", howMany)
	}

	if pages.Items[0].Pagelanguage != lang {
		t.Fatalf("We should get a page with lang code %s", lang)
	}
}

// TestParsing uses a static input file to test
// parsing of the results
func TestParsing(t *testing.T) {
	testFile := "_testdata/test1.json"
	b, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Error opening the file %s", testFile)
	}
	var pages Pages
	err2 := parseJson(b, &pages)
	if err2 != nil {
		t.Fatalf("Error parsing json: %s", err2)
	}
	// Order is not guaranteed
	position:= 0
	if pages.Items[1].Title != "Mimosciadella" {
		if pages.Items[0].Title != "Mimosciadella" {
			t.Fatalf("Error parsing title, got %s", pages.Items[0].Title)
		} else {
			position= 1
		}
	}
	if pages.Items[position].Length != 9352 {
		t.Fatalf("Error parsing length, got %f", pages.Items[0].Length)
	}

	if pages.Items[position^1].Fullurl != "https://en.wikipedia.org/wiki/Mimosciadella" {
		t.Fatalf("Error parsing URL, got %s", pages.Items[1].Fullurl)
	}
}
