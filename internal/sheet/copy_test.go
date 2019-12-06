package sheet

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

// TestCopyFlat asserts that Copy correctly copies files at a single level of
// depth
func TestCopyFlat(t *testing.T) {

	// mock a cheatsheet file
	text := "this is the cheatsheet text"
	src, err := ioutil.TempFile("", "foo-src")
	if err != nil {
		t.Errorf("failed to mock cheatsheet: %v", err)
	}
	defer src.Close()
	defer os.Remove(src.Name())

	if _, err := src.WriteString(text); err != nil {
		t.Errorf("failed to write to mock cheatsheet: %v", err)
	}

	// mock a cheatsheet struct
	sheet, err := New("foo", src.Name(), []string{}, false)
	if err != nil {
		t.Errorf("failed to init cheatsheet: %v", err)
	}

	// compute the outfile's path
	outpath := path.Join(os.TempDir(), sheet.Title)
	defer os.Remove(outpath)

	// attempt to copy the cheatsheet
	err = sheet.Copy(outpath)
	if err != nil {
		t.Errorf("failed to copy cheatsheet: %v", err)
	}

	// assert that the destination file contains the correct text
	got, err := ioutil.ReadFile(outpath)
	if err != nil {
		t.Errorf("failed to read destination file: %v", err)
	}
	if string(got) != text {
		t.Errorf(
			"destination file contained wrong text: want: '%s', got: '%s'",
			text,
			got,
		)
	}
}

// TestCopyDeep asserts that Copy correctly copies files at several levels of
// depth
func TestCopyDeep(t *testing.T) {

	// mock a cheatsheet file
	text := "this is the cheatsheet text"
	src, err := ioutil.TempFile("", "foo-src")
	if err != nil {
		t.Errorf("failed to mock cheatsheet: %v", err)
	}
	defer src.Close()
	defer os.Remove(src.Name())

	if _, err := src.WriteString(text); err != nil {
		t.Errorf("failed to write to mock cheatsheet: %v", err)
	}

	// mock a cheatsheet struct
	sheet, err := New("/cheat-tests/alpha/bravo/foo", src.Name(), []string{}, false)
	if err != nil {
		t.Errorf("failed to init cheatsheet: %v", err)
	}

	// compute the outfile's path
	outpath := path.Join(os.TempDir(), sheet.Title)
	defer os.RemoveAll(path.Join(os.TempDir(), "cheat-tests"))

	// attempt to copy the cheatsheet
	err = sheet.Copy(outpath)
	if err != nil {
		t.Errorf("failed to copy cheatsheet: %v", err)
	}

	// assert that the destination file contains the correct text
	got, err := ioutil.ReadFile(outpath)
	if err != nil {
		t.Errorf("failed to read destination file: %v", err)
	}
	if string(got) != text {
		t.Errorf(
			"destination file contained wrong text: want: '%s', got: '%s'",
			text,
			got,
		)
	}
}
