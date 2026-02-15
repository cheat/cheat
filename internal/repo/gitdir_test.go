package repo

import (
	"os"
	"path/filepath"
	"testing"
)

// setupGitDirTestTree creates a temporary directory structure that exercises
// every case documented in GitDir's comment block. The caller must defer
// os.RemoveAll on the returned root.
//
// Layout:
//
//	root/
//	├── plain/                         # not a repository
//	│   └── sheet
//	├── repo/                          # a repository (.git is a directory)
//	│   ├── .git/
//	│   │   ├── HEAD
//	│   │   ├── objects/
//	│   │   │   └── pack/
//	│   │   └── refs/
//	│   │       └── heads/
//	│   ├── .gitignore
//	│   ├── .gitattributes
//	│   └── sheet
//	├── submodule/                     # a submodule (.git is a file)
//	│   ├── .git                       # file, not directory
//	│   └── sheet
//	├── dotgit-suffix.git/             # directory name ends in .git (#711)
//	│   └── cheat/
//	│       └── sheet
//	├── dotgit-mid.git/                # .git suffix mid-path (#711)
//	│   └── nested/
//	│       └── sheet
//	├── .github/                       # .github directory (not .git)
//	│   └── workflows/
//	│       └── ci.yml
//	└── .hidden/                       # generic hidden directory
//	    └── sheet
func setupGitDirTestTree(t *testing.T) string {
	t.Helper()

	root := t.TempDir()

	dirs := []string{
		// case 1: not a repository
		filepath.Join(root, "plain"),

		// case 2: a repository (.git directory with contents)
		filepath.Join(root, "repo", ".git", "objects", "pack"),
		filepath.Join(root, "repo", ".git", "refs", "heads"),

		// case 4: a submodule (.git is a file)
		filepath.Join(root, "submodule"),

		// case 6: directory name ending in .git (#711)
		filepath.Join(root, "dotgit-suffix.git", "cheat"),
		filepath.Join(root, "dotgit-mid.git", "nested"),

		// .github (should not be confused with .git)
		filepath.Join(root, ".github", "workflows"),

		// generic hidden directory
		filepath.Join(root, ".hidden"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("failed to create dir %s: %v", dir, err)
		}
	}

	files := map[string]string{
		// sheets
		filepath.Join(root, "plain", "sheet"):                      "plain sheet",
		filepath.Join(root, "repo", "sheet"):                       "repo sheet",
		filepath.Join(root, "submodule", "sheet"):                  "submod sheet",
		filepath.Join(root, "dotgit-suffix.git", "cheat", "sheet"): "dotgit sheet",
		filepath.Join(root, "dotgit-mid.git", "nested", "sheet"):   "dotgit nested",
		filepath.Join(root, ".hidden", "sheet"):                    "hidden sheet",

		// git metadata
		filepath.Join(root, "repo", ".git", "HEAD"):           "ref: refs/heads/main\n",
		filepath.Join(root, "repo", ".gitignore"):             "*.tmp\n",
		filepath.Join(root, "repo", ".gitattributes"):         "* text=auto\n",
		filepath.Join(root, "submodule", ".git"):              "gitdir: ../.git/modules/sub\n",
		filepath.Join(root, ".github", "workflows", "ci.yml"): "name: CI\n",
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatalf("failed to write %s: %v", path, err)
		}
	}

	return root
}

func TestGitDir(t *testing.T) {
	root := setupGitDirTestTree(t)

	tests := []struct {
		name string
		path string
		want bool
	}{
		// Case 1: not a repository — no .git anywhere in path
		{
			name: "plain directory, no repo",
			path: filepath.Join(root, "plain", "sheet"),
			want: false,
		},

		// Case 2: a repository — paths *inside* .git/ should be detected
		{
			name: "inside .git directory",
			path: filepath.Join(root, "repo", ".git", "HEAD"),
			want: true,
		},
		{
			name: "inside .git/objects",
			path: filepath.Join(root, "repo", ".git", "objects", "pack", "somefile"),
			want: true,
		},
		{
			name: "inside .git/refs",
			path: filepath.Join(root, "repo", ".git", "refs", "heads", "main"),
			want: true,
		},

		// Case 2 (cont.): files *alongside* .git should NOT be detected
		{
			name: "sheet in repo root (beside .git dir)",
			path: filepath.Join(root, "repo", "sheet"),
			want: false,
		},

		// Case 3: .git* files (like .gitignore) should NOT trigger
		{
			name: ".gitignore file",
			path: filepath.Join(root, "repo", ".gitignore"),
			want: false,
		},
		{
			name: ".gitattributes file",
			path: filepath.Join(root, "repo", ".gitattributes"),
			want: false,
		},

		// Case 4: submodule — .git is a file, not a directory
		{
			name: "sheet in submodule (where .git is a file)",
			path: filepath.Join(root, "submodule", "sheet"),
			want: false,
		},

		// Case 6: directory name ends with .git (#711)
		{
			name: "sheet under directory ending in .git",
			path: filepath.Join(root, "dotgit-suffix.git", "cheat", "sheet"),
			want: false,
		},
		{
			name: "sheet under .git-suffixed dir, nested deeper",
			path: filepath.Join(root, "dotgit-mid.git", "nested", "sheet"),
			want: false,
		},

		// .github directory — must not be confused with .git
		{
			name: "file inside .github directory",
			path: filepath.Join(root, ".github", "workflows", "ci.yml"),
			want: false,
		},

		// Hidden directory that is not .git
		{
			name: "file inside generic hidden directory",
			path: filepath.Join(root, ".hidden", "sheet"),
			want: false,
		},

		// Path with no .git at all
		{
			name: "path with no .git component whatsoever",
			path: filepath.Join(root, "nonexistent", "file"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GitDir(tt.path)
			if err != nil {
				t.Fatalf("GitDir(%q) returned unexpected error: %v", tt.path, err)
			}
			if got != tt.want {
				t.Errorf("GitDir(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

// TestGitDirWithNestedGitDir tests a repo inside a .git-suffixed parent
// directory. This is the nastiest combination: a real .git directory that
// appears *after* a .git suffix in the path.
func TestGitDirWithNestedGitDir(t *testing.T) {
	root := t.TempDir()

	// Create: root/cheats.git/repo/.git/HEAD
	//         root/cheats.git/repo/sheet
	gitDir := filepath.Join(root, "cheats.git", "repo", ".git")
	if err := os.MkdirAll(gitDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(gitDir, "HEAD"), []byte("ref: refs/heads/main\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "cheats.git", "repo", "sheet"), []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "sheet beside .git in .git-suffixed parent",
			path: filepath.Join(root, "cheats.git", "repo", "sheet"),
			want: false,
		},
		{
			name: "file inside .git inside .git-suffixed parent",
			path: filepath.Join(root, "cheats.git", "repo", ".git", "HEAD"),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GitDir(tt.path)
			if err != nil {
				t.Fatalf("GitDir(%q) returned unexpected error: %v", tt.path, err)
			}
			if got != tt.want {
				t.Errorf("GitDir(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

// TestGitDirSubmoduleInsideDotGitSuffix tests a submodule (.git file)
// inside a .git-suffixed parent directory.
func TestGitDirSubmoduleInsideDotGitSuffix(t *testing.T) {
	root := t.TempDir()

	// Create: root/personal.git/submod/.git  (file)
	//         root/personal.git/submod/sheet
	subDir := filepath.Join(root, "personal.git", "submod")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatal(err)
	}
	// .git as a file (submodule pointer)
	if err := os.WriteFile(filepath.Join(subDir, ".git"), []byte("gitdir: ../../.git/modules/sub\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(subDir, "sheet"), []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}

	got, err := GitDir(filepath.Join(subDir, "sheet"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got {
		t.Error("GitDir should return false for sheet in submodule under .git-suffixed parent")
	}
}

// TestGitDirIntegrationWalk simulates what sheets.Load does: walking a
// directory tree and checking each path with GitDir. This verifies that
// the function works correctly in the context of filepath.Walk, which is
// how it is actually called.
func TestGitDirIntegrationWalk(t *testing.T) {
	root := setupGitDirTestTree(t)

	// Walk the tree and collect which paths GitDir says to skip
	var skipped []string
	var visited []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		isGit, err := GitDir(path)
		if err != nil {
			return err
		}
		if isGit {
			skipped = append(skipped, path)
		} else {
			visited = append(visited, path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Walk failed: %v", err)
	}

	// Files inside .git/ should be skipped
	expectSkipped := []string{
		filepath.Join(root, "repo", ".git", "HEAD"),
	}
	for _, want := range expectSkipped {
		found := false
		for _, got := range skipped {
			if got == want {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected %q to be skipped, but it was not", want)
		}
	}

	// Sheets should NOT be skipped — including the #711 case
	expectVisited := []string{
		filepath.Join(root, "plain", "sheet"),
		filepath.Join(root, "repo", "sheet"),
		filepath.Join(root, "submodule", "sheet"),
		filepath.Join(root, "dotgit-suffix.git", "cheat", "sheet"),
		filepath.Join(root, "dotgit-mid.git", "nested", "sheet"),
		filepath.Join(root, ".hidden", "sheet"),
	}
	for _, want := range expectVisited {
		found := false
		for _, got := range visited {
			if got == want {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected %q to be visited (not skipped), but it was not found in visited paths", want)
		}
	}
}
