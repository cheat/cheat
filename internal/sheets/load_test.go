package sheets

import "testing"

type PathTestRes struct {
	fileName string
	exist    bool
}

type PathTest struct {
	path   string
	result PathTestRes
}

var dirtests = []PathTest{
	{"abc", PathTestRes{"", false}},
	{"abc/def", PathTestRes{"", false}},
	{"a/b/.x", PathTestRes{"", false}},
	{"a/b/c.", PathTestRes{"", false}},
	{"a/b/c.x", PathTestRes{"", false}},
	{"a/b/b/a", PathTestRes{"", false}},
	{"a/b/c/b", PathTestRes{"", false}},

	{"b/b", PathTestRes{"b", true}},
	{"a/b/b", PathTestRes{"a/b", true}},
}

func TestSameNameToParentDir(t *testing.T) {
	for _, test := range dirtests {
		if f, e := sameNameToParentDir(test.path); f != test.result.fileName || e != test.result.exist {
			t.Errorf(" sameNameToParentDir (%q) = %q, %t, want %q, %t", test.path, f, e, test.result.fileName, test.result.exist)
		}
	}
}
