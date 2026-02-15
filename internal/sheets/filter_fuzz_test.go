package sheets

import (
	"strings"
	"testing"

	"github.com/cheat/cheat/internal/sheet"
)

// FuzzFilter tests the Filter function with various tag combinations
func FuzzFilter(f *testing.F) {
	// Add seed corpus with various tag scenarios
	// Format: "tags to filter by" (comma-separated)
	f.Add("linux")
	f.Add("linux,bash")
	f.Add("linux,bash,ssh")
	f.Add("")
	f.Add(" ")
	f.Add("  linux  ")
	f.Add("linux,")
	f.Add(",linux")
	f.Add(",,")
	f.Add("linux,,bash")
	f.Add("tag-with-dash")
	f.Add("tag_with_underscore")
	f.Add("UPPERCASE")
	f.Add("miXedCase")
	f.Add("🎉emoji")
	f.Add("tag with spaces")
	f.Add("\ttab\ttag")
	f.Add("tag\nwith\nnewline")
	f.Add("very-long-tag-name-that-might-cause-issues-somewhere")
	f.Add(strings.Repeat("a,", 100) + "a")

	f.Fuzz(func(t *testing.T, tagString string) {
		// Split the tag string into individual tags
		var tags []string
		if tagString != "" {
			tags = strings.Split(tagString, ",")
		}

		// Create test data - some sheets with various tags
		cheatpaths := []map[string]sheet.Sheet{
			{
				"sheet1": sheet.Sheet{
					Title: "sheet1",
					Tags:  []string{"linux", "bash"},
				},
				"sheet2": sheet.Sheet{
					Title: "sheet2",
					Tags:  []string{"linux", "ssh", "networking"},
				},
				"sheet3": sheet.Sheet{
					Title: "sheet3",
					Tags:  []string{"UPPERCASE", "miXedCase"},
				},
			},
			{
				"sheet4": sheet.Sheet{
					Title: "sheet4",
					Tags:  []string{"tag with spaces", "🎉emoji"},
				},
				"sheet5": sheet.Sheet{
					Title: "sheet5",
					Tags:  []string{}, // No tags
				},
			},
		}

		// The function should not panic
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Filter panicked with tags %q: %v", tags, r)
				}
			}()

			result := Filter(cheatpaths, tags)

			// Verify invariants
			// 1. Result should have same number of cheatpaths
			if len(result) != len(cheatpaths) {
				t.Errorf("Filter changed number of cheatpaths: got %d, want %d",
					len(result), len(cheatpaths))
			}

			// 2. Each filtered sheet should contain all requested tags
			for _, filteredPath := range result {
				for title, sheet := range filteredPath {
					// Verify this sheet has all the tags we filtered for
					for _, tag := range tags {
						trimmedTag := strings.TrimSpace(tag)
						if trimmedTag == "" {
							continue // Skip empty tags
						}
						if !sheet.Tagged(trimmedTag) {
							t.Errorf("Sheet %q passed filter but doesn't have tag %q",
								title, trimmedTag)
						}
					}
				}
			}

			// 3. Empty tag list should return all sheets
			if len(tags) == 0 || (len(tags) == 1 && tags[0] == "") {
				totalOriginal := 0
				totalFiltered := 0
				for _, path := range cheatpaths {
					totalOriginal += len(path)
				}
				for _, path := range result {
					totalFiltered += len(path)
				}
				if totalFiltered != totalOriginal {
					t.Errorf("Empty filter should return all sheets: got %d, want %d",
						totalFiltered, totalOriginal)
				}
			}
		}()
	})
}

// FuzzFilterEdgeCases tests Filter with extreme inputs
func FuzzFilterEdgeCases(f *testing.F) {
	// Seed with number of tags and tag length
	f.Add(0, 0)
	f.Add(1, 10)
	f.Add(10, 10)
	f.Add(100, 5)
	f.Add(1000, 3)

	f.Fuzz(func(t *testing.T, numTags int, tagLen int) {
		// Limit to reasonable values to avoid memory issues
		if numTags > 1000 || numTags < 0 || tagLen > 100 || tagLen < 0 {
			t.Skip("Skipping unreasonable test case")
		}

		// Generate tags
		tags := make([]string, numTags)
		for i := 0; i < numTags; i++ {
			// Create a tag of specified length
			if tagLen > 0 {
				tags[i] = strings.Repeat("a", tagLen) + string(rune(i%26+'a'))
			}
		}

		// Create a sheet with no tags (should be filtered out)
		cheatpaths := []map[string]sheet.Sheet{
			{
				"test": sheet.Sheet{
					Title: "test",
					Tags:  []string{},
				},
			},
		}

		// Should not panic with many tags
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Filter panicked with %d tags of length %d: %v",
						numTags, tagLen, r)
				}
			}()

			result := Filter(cheatpaths, tags)

			// With non-matching tags, result should be empty
			if numTags > 0 && tagLen > 0 {
				if len(result[0]) != 0 {
					t.Errorf("Expected empty result with non-matching tags, got %d sheets",
						len(result[0]))
				}
			}
		}()
	})
}
