package sheet

import (
	"reflect"
	"testing"

	"github.com/cheat/cheat/internal/mock"
)

// TestSheetSuccess asserts that sheets initialize properly
func TestSheetSuccess(t *testing.T) {

	// initialize a sheet
	sheet, err := New(
		"foo",
		mock.Path("sheet/foo"),
		[]string{"alpha", "bravo"},
		false,
	)
	if err != nil {
		t.Errorf("failed to load sheet: %v", err)
	}

	// assert that the sheet loaded correctly
	if sheet.Title != "foo" {
		t.Errorf("failed to init title: want: foo, got: %s", sheet.Title)
	}

	if sheet.Path != mock.Path("sheet/foo") {
		t.Errorf(
			"failed to init path: want: %s, got: %s",
			mock.Path("sheet/foo"),
			sheet.Path,
		)
	}

	wantText := "# To foo the bar:\n  foo bar\n"
	if sheet.Text != wantText {
		t.Errorf("failed to init text: want: %s, got: %s", wantText, sheet.Text)
	}

	// NB: tags should sort alphabetically
	wantTags := []string{"alpha", "bar", "baz", "bravo", "foo"}
	if !reflect.DeepEqual(sheet.Tags, wantTags) {
		t.Errorf("failed to init tags: want: %v, got: %v", wantTags, sheet.Tags)
	}

	if sheet.Syntax != "sh" {
		t.Errorf("failed to init syntax: want: sh, got: %s", sheet.Syntax)
	}

	if sheet.ReadOnly != false {
		t.Errorf("failed to init readonly")
	}
}

// TestSheetFailure asserts that an error is returned if the sheet cannot be
// read
func TestSheetFailure(t *testing.T) {

	// initialize a sheet
	_, err := New(
		"foo",
		mock.Path("/does-not-exist"),
		[]string{"alpha", "bravo"},
		false,
	)
	if err == nil {
		t.Errorf("failed to return an error on unreadable sheet")
	}
}

// TestSheetFrontMatterFailure asserts that an error is returned if the sheet's
// frontmatter cannot be parsed.
func TestSheetFrontMatterFailure(t *testing.T) {

	// initialize a sheet
	_, err := New(
		"foo",
		mock.Path("sheet/bad-fm"),
		[]string{"alpha", "bravo"},
		false,
	)
	if err == nil {
		t.Errorf("failed to return an error on malformed front-matter")
	}
}
