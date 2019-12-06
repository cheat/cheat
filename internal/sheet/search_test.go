package sheet

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/davecgh/go-spew/spew"
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
	matches := sheet.Search(reg, false)

	// assert that no matches were found
	if len(matches) != 0 {
		t.Errorf("failure: expected no matches: got: %s", spew.Sdump(matches))
	}
}

// TestSearchSingleMatchNoColor asserts that the expected output is returned
// when a single match is returned, and no colorization is applied.
func TestSearchSingleMatchNoColor(t *testing.T) {

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
	matches := sheet.Search(reg, false)

	// specify the expected results
	want := []Match{
		Match{
			Line: 1,
			Text: "The quick brown fox",
		},
	}

	// assert that the correct matches were returned
	if !reflect.DeepEqual(matches, want) {
		t.Errorf(
			"failed to return expected matches: want:\n%s, got:\n%s",
			spew.Sdump(want),
			spew.Sdump(matches),
		)
	}
}

// TestSearchSingleMatchColorized asserts that the expected output is returned
// when a single match is returned, and colorization is applied
func TestSearchSingleMatchColorized(t *testing.T) {

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
	matches := sheet.Search(reg, true)

	// specify the expected results
	want := []Match{
		Match{
			Line: 1,
			Text: "The quick brown \x1b[1;31mfox\x1b[0m",
		},
	}

	// assert that the correct matches were returned
	if !reflect.DeepEqual(matches, want) {
		t.Errorf(
			"failed to return expected matches: want:\n%s, got:\n%s",
			spew.Sdump(want),
			spew.Sdump(matches),
		)
	}
}

// TestSearchMultiMatchNoColor asserts that the expected output is returned
// when a multiple matches are returned, and no colorization is applied
func TestSearchMultiMatchNoColor(t *testing.T) {

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
	matches := sheet.Search(reg, false)

	// specify the expected results
	want := []Match{
		Match{
			Line: 1,
			Text: "The quick brown fox",
		},
		Match{
			Line: 3,
			Text: "the lazy dog.",
		},
	}

	// assert that the correct matches were returned
	if !reflect.DeepEqual(matches, want) {
		t.Errorf(
			"failed to return expected matches: want:\n%s, got:\n%s",
			spew.Sdump(want),
			spew.Sdump(matches),
		)
	}
}

// TestSearchMultiMatchColorized asserts that the expected output is returned
// when a multiple matches are returned, and colorization is applied
func TestSearchMultiMatchColorized(t *testing.T) {

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
	matches := sheet.Search(reg, true)

	// specify the expected results
	want := []Match{
		Match{
			Line: 1,
			Text: "\x1b[1;31mThe\x1b[0m quick brown fox",
		},
		Match{
			Line: 3,
			Text: "\x1b[1;31mthe\x1b[0m lazy dog.",
		},
	}

	// assert that the correct matches were returned
	if !reflect.DeepEqual(matches, want) {
		t.Errorf(
			"failed to return expected matches: want:\n%s, got:\n%s",
			spew.Sdump(want),
			spew.Sdump(matches),
		)
	}
}
