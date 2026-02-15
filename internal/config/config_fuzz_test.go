package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// FuzzFindLocalCheatpath exercises findLocalCheatpath with randomised
// directory depths and .cheat placements.  For each fuzz input it builds a
// temporary directory hierarchy, places a single .cheat directory at a
// computed level, and asserts that the function always returns it.
func FuzzFindLocalCheatpath(f *testing.F) {
	// Seed corpus: (totalDepth, cheatPlacement)
	f.Add(uint8(1), uint8(0))  // depth 1, .cheat at root
	f.Add(uint8(3), uint8(0))  // depth 3, .cheat at root
	f.Add(uint8(5), uint8(3))  // depth 5, .cheat at level 3
	f.Add(uint8(1), uint8(1))  // depth 1, .cheat at same level as search dir
	f.Add(uint8(10), uint8(5)) // deep hierarchy

	f.Fuzz(func(t *testing.T, totalDepth uint8, cheatPlacement uint8) {
		// Clamp to reasonable values to keep I/O bounded
		depth := int(totalDepth%15) + 1              // 1..15
		cheatAt := int(cheatPlacement) % (depth + 1) // 0..depth (0 = tempDir itself)

		tempDir := t.TempDir()

		// Build chain: tempDir/d0/d1/…/d{depth-1}
		dirs := make([]string, 0, depth+1)
		dirs = append(dirs, tempDir)
		current := tempDir
		for i := 0; i < depth; i++ {
			current = filepath.Join(current, fmt.Sprintf("d%d", i))
			if err := os.Mkdir(current, 0755); err != nil {
				t.Fatalf("mkdir: %v", err)
			}
			dirs = append(dirs, current)
		}

		// Place .cheat at dirs[cheatAt]
		cheatDir := filepath.Join(dirs[cheatAt], ".cheat")
		if err := os.Mkdir(cheatDir, 0755); err != nil {
			t.Fatalf("mkdir .cheat: %v", err)
		}

		// Search from the deepest directory
		result := findLocalCheatpath(current)

		// Invariant 1: must find the .cheat we placed
		if result != cheatDir {
			t.Errorf("depth=%d cheatAt=%d: expected %s, got %s",
				depth, cheatAt, cheatDir, result)
		}

		// Invariant 2: result must end with /.cheat
		if !strings.HasSuffix(result, string(filepath.Separator)+".cheat") {
			t.Errorf("result %q does not end with /.cheat", result)
		}

		// Invariant 3: result must be under tempDir
		if !strings.HasPrefix(result, tempDir) {
			t.Errorf("result %q is not under tempDir %s", result, tempDir)
		}
	})
}

// FuzzFindLocalCheatpathNearestWins verifies that when two .cheat directories
// exist at different levels of the ancestor chain, the nearest one is returned.
func FuzzFindLocalCheatpathNearestWins(f *testing.F) {
	f.Add(uint8(5), uint8(1), uint8(3))
	f.Add(uint8(8), uint8(0), uint8(7))
	f.Add(uint8(3), uint8(0), uint8(2))
	f.Add(uint8(10), uint8(2), uint8(8))

	f.Fuzz(func(t *testing.T, totalDepth, shallowRaw, deepRaw uint8) {
		depth := int(totalDepth%12) + 2 // 2..13 (need room for two placements)
		s := int(shallowRaw) % depth
		d := int(deepRaw) % depth

		// Need two distinct levels
		if s == d {
			d = (d + 1) % depth
		}
		// Ensure s < d (shallow is higher in tree, deep is closer to search dir)
		if s > d {
			s, d = d, s
		}

		tempDir := t.TempDir()

		// Build chain
		dirs := make([]string, 0, depth+1)
		dirs = append(dirs, tempDir)
		current := tempDir
		for i := 0; i < depth; i++ {
			current = filepath.Join(current, fmt.Sprintf("d%d", i))
			if err := os.Mkdir(current, 0755); err != nil {
				t.Fatalf("mkdir: %v", err)
			}
			dirs = append(dirs, current)
		}

		// Place .cheat at both levels
		shallowCheat := filepath.Join(dirs[s], ".cheat")
		deepCheat := filepath.Join(dirs[d], ".cheat")
		if err := os.Mkdir(shallowCheat, 0755); err != nil {
			t.Fatalf("mkdir shallow .cheat: %v", err)
		}
		if err := os.Mkdir(deepCheat, 0755); err != nil {
			t.Fatalf("mkdir deep .cheat: %v", err)
		}

		// Search from the deepest directory — should find the deeper (nearer) .cheat
		result := findLocalCheatpath(current)
		if result != deepCheat {
			t.Errorf("depth=%d shallow=%d deep=%d: expected nearest %s, got %s",
				depth, s, d, deepCheat, result)
		}
	})
}
