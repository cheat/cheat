package main

import "testing"

type PathTestRes struct {
	tag       string
	cheatPath string
}

type PathTest struct {
	path   string
	result PathTestRes
}

var dirtests = []PathTest{
	{"abc", PathTestRes{"abc", ""}},
	{"abc/def", PathTestRes{"abc", "def"}},
	{"a/b/.x", PathTestRes{"a", "b/.x"}},
	{"a/b/c.", PathTestRes{"a", "b/c."}},
	{"a/b/c.x", PathTestRes{"a", "b/c.x"}},
	{"a/b/b/a", PathTestRes{"a", "b/b/a"}},
	{"a/b/c/b", PathTestRes{"a", "b/c/b"}},

	{"b/b", PathTestRes{"b", "b"}},
	{"a/b/b", PathTestRes{"a", "b/b"}},
}

func TestSplitTagAndPath(t *testing.T) {
	for _, test := range dirtests {
		if tag, cheatPath := splitTagAndPath(test.path); tag != test.result.tag || cheatPath != test.result.cheatPath {
			t.Errorf("splitTagAndPath (%q) = %q, %q, want %q, %q", test.path, tag, cheatPath, test.result.tag, test.result.cheatPath)
		}
	}
}
