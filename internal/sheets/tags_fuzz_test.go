package sheets

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/cheat/cheat/internal/sheet"
)

// FuzzTags tests the Tags function with various tag combinations
func FuzzTags(f *testing.F) {
	// Add seed corpus
	// Format: comma-separated tags that will be distributed across sheets
	f.Add("linux,bash,ssh")
	f.Add("")
	f.Add("single")
	f.Add("duplicate,duplicate,duplicate")
	f.Add("  spaces  ,  around  ,  tags  ")
	f.Add("MiXeD,UPPER,lower")
	f.Add("special-chars,under_score,dot.ted")
	f.Add("emoji🎉,unicode中文,symbols@#$")
	f.Add("\ttab,\nnewline,\rcarriage")
	f.Add(",,,,")                                          // Multiple empty tags
	f.Add(strings.Repeat("tag,", 100))                     // Many tags
	f.Add("a," + strings.Repeat("very-long-tag-name", 10)) // Long tag names

	f.Fuzz(func(t *testing.T, tagString string) {
		// Split tags and distribute them across multiple sheets
		var allTags []string
		if tagString != "" {
			allTags = strings.Split(tagString, ",")
		}

		// Create test cheatpaths with various tag distributions
		cheatpaths := []map[string]sheet.Sheet{}

		// Distribute tags across 3 paths with overlapping tags
		for i := 0; i < 3; i++ {
			path := make(map[string]sheet.Sheet)

			// Each path gets some subset of tags
			for j, tag := range allTags {
				if j%3 == i || j%(i+2) == 0 { // Create some overlap
					sheetName := string(rune('a' + j%26))
					path[sheetName] = sheet.Sheet{
						Title: sheetName,
						Tags:  []string{tag},
					}
				}
			}

			// Add a sheet with multiple tags
			if len(allTags) > 1 {
				path["multi"] = sheet.Sheet{
					Title: "multi",
					Tags:  allTags[:len(allTags)/2+1], // First half of tags
				}
			}

			cheatpaths = append(cheatpaths, path)
		}

		// The function should not panic
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Tags panicked with input %q: %v", tagString, r)
				}
			}()

			result := Tags(cheatpaths)

			// Verify invariants
			// 1. Result should be sorted
			for i := 1; i < len(result); i++ {
				if result[i-1] >= result[i] {
					t.Errorf("Tags not sorted: %q >= %q at positions %d, %d",
						result[i-1], result[i], i-1, i)
				}
			}

			// 2. No duplicates in result
			seen := make(map[string]bool)
			for _, tag := range result {
				if seen[tag] {
					t.Errorf("Duplicate tag in result: %q", tag)
				}
				seen[tag] = true
			}

			// 3. All non-empty tags from input should be in result
			// (This is approximate since we distributed tags in a complex way)
			inputTags := make(map[string]bool)
			for _, tag := range allTags {
				if tag != "" {
					inputTags[tag] = true
				}
			}

			resultTags := make(map[string]bool)
			for _, tag := range result {
				resultTags[tag] = true
			}

			// Result might have fewer tags due to distribution logic,
			// but shouldn't have tags not in the input
			for tag := range resultTags {
				found := false
				for inputTag := range inputTags {
					if tag == inputTag {
						found = true
						break
					}
				}
				if !found && tag != "" {
					t.Errorf("Result contains tag %q not derived from input", tag)
				}
			}

			// 4. Valid UTF-8 (Tags function should filter out invalid UTF-8)
			for _, tag := range result {
				if !utf8.ValidString(tag) {
					t.Errorf("Invalid UTF-8 in tag: %q", tag)
				}
			}
		}()
	})
}
