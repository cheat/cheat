package sheet

import (
	"reflect"
	"regexp"
	"testing"
)

// TestSearchNoMatch ensures that the expected output is returned when no
// matches are found
func TestSearchNoMatch(t *testing.T) {

	// mock a cheatsheet
	sheet := Sheet{
		Text: "The quick brown fox\njumped over\nthe lazy dog.",
	}

	// compile the search regex
	reg, err := regexp.Compile("(?i)foo")
	if err != nil {
		t.Errorf("failed to compile regex: %v", err)
	}

	// search the sheet
	matches := sheet.Search(reg)

	// assert that no matches were found
	if matches != "" {
		t.Errorf("failure: expected no matches: got: %s", matches)
	}
}

// TestSearchSingleMatch asserts that the expected output is returned
// when a single match is returned
func TestSearchSingleMatch(t *testing.T) {

	// mock a cheatsheet
	sheet := Sheet{
		Text: "The quick brown fox\njumped over\nthe lazy dog.",
	}

	// compile the search regex
	reg, err := regexp.Compile("(?i)fox")
	if err != nil {
		t.Errorf("failed to compile regex: %v", err)
	}

	// search the sheet
	matches := sheet.Search(reg)

	// specify the expected results
	want := "The quick brown fox"

	// assert that the correct matches were returned
	if matches != want {
		t.Errorf(
			"failed to return expected matches: want:\n%s, got:\n%s",
			want,
			matches,
		)
	}
}

// TestSearchMultiMatch asserts that the expected output is returned
// when a multiple matches are returned
func TestSearchMultiMatch(t *testing.T) {

	// mock a cheatsheet
	sheet := Sheet{
		Text: "The quick brown fox\njumped over\nthe lazy dog.",
	}

	// compile the search regex
	reg, err := regexp.Compile("(?i)the")
	if err != nil {
		t.Errorf("failed to compile regex: %v", err)
	}

	// search the sheet
	matches := sheet.Search(reg)

	// specify the expected results
	want := "The quick brown fox\nthe lazy dog."

	// assert that the correct matches were returned
	if !reflect.DeepEqual(matches, want) {
		t.Errorf(
			"failed to return expected matches: want:\n%s, got:\n%s",
			want,
			matches,
		)
	}
}
