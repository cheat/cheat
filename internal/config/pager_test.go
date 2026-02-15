package config

import (
	"os"
	"runtime"
	"testing"
)

// TestPager tests the Pager function
func TestPager(t *testing.T) {
	// Save original env var
	oldPager := os.Getenv("PAGER")
	defer os.Setenv("PAGER", oldPager)

	t.Run("windows default", func(t *testing.T) {
		if runtime.GOOS != "windows" {
			t.Skip("skipping windows test on non-windows platform")
		}

		os.Setenv("PAGER", "")
		pager := Pager()
		if pager != "more" {
			t.Errorf("expected 'more' on windows, got %s", pager)
		}
	})

	t.Run("PAGER env var", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("skipping non-windows test on windows platform")
		}

		os.Setenv("PAGER", "bat")
		pager := Pager()
		if pager != "bat" {
			t.Errorf("expected PAGER env var value, got %s", pager)
		}
	})

	t.Run("fallback to system pager", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("skipping non-windows test on windows platform")
		}

		os.Setenv("PAGER", "")
		pager := Pager()

		// Should find one of the fallback pagers or return empty string
		validPagers := map[string]bool{
			"":      true, // no pager found
			"pager": true,
			"less":  true,
			"more":  true,
		}

		// Check if it's a path to one of these
		found := false
		for p := range validPagers {
			if p == "" && pager == "" {
				found = true
				break
			}
			if p != "" && (pager == p || len(pager) >= len(p) && pager[len(pager)-len(p):] == p) {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("unexpected pager value: %s", pager)
		}
	})

	t.Run("no pager available", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("skipping non-windows test on windows platform")
		}

		os.Setenv("PAGER", "")

		// Save and modify PATH to ensure no pagers are found
		oldPath := os.Getenv("PATH")
		defer os.Setenv("PATH", oldPath)
		os.Setenv("PATH", "/nonexistent")

		pager := Pager()
		if pager != "" {
			t.Errorf("expected empty string when no pager found, got %s", pager)
		}
	})
}
