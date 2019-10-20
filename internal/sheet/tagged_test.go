package sheet

import (
	"testing"
)

// TestTagged ensures that tags are properly recognized as being absent or
// present
func TestTagged(t *testing.T) {

	// initialize a cheatsheet
	tags := []string{"foo", "bar", "baz"}
	sheet := Sheet{Tags: tags}

	// assert that set tags are recognized as set
	for _, tag := range tags {
		if sheet.Tagged(tag) == false {
			t.Errorf("failed to recognize tag: %s", tag)
		}
	}

	// assert that unset tags are recognized as unset
	if sheet.Tagged("qux") {
		t.Errorf("failed to recognize absent tag")
	}
}
