package sheet

import (
	"runtime"
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		// Valid names
		{
			name:    "simple name",
			input:   "docker",
			wantErr: false,
		},
		{
			name:    "name with slash",
			input:   "docker/compose",
			wantErr: false,
		},
		{
			name:    "name with multiple slashes",
			input:   "lang/go/slice",
			wantErr: false,
		},
		{
			name:    "name with dash and underscore",
			input:   "my-cheat_sheet",
			wantErr: false,
		},
		// Invalid names
		{
			name:    "empty name",
			input:   "",
			wantErr: true,
			errMsg:  "empty",
		},
		{
			name:    "parent directory traversal",
			input:   "../etc/passwd",
			wantErr: true,
			errMsg:  "'..'",
		},
		{
			name:    "complex traversal",
			input:   "foo/../../etc/passwd",
			wantErr: true,
			errMsg:  "'..'",
		},
		{
			name:    "absolute path unix",
			input:   "/etc/passwd",
			wantErr: runtime.GOOS != "windows", // /etc/passwd is not absolute on Windows
			errMsg:  "absolute",
		},
		{
			name:    "absolute path windows",
			input:   `C:\evil`,
			wantErr: runtime.GOOS == "windows", // C:\evil is not absolute on Unix
			errMsg:  "absolute",
		},
		{
			name:    "home directory",
			input:   "~/secrets",
			wantErr: true,
			errMsg:  "'~'",
		},
		{
			name:    "just dots",
			input:   "..",
			wantErr: true,
			errMsg:  "'..'",
		},
		{
			name:    "hidden file not allowed",
			input:   ".hidden",
			wantErr: true,
			errMsg:  "cannot start with '.'",
		},
		{
			name:    "current dir is ok",
			input:   "./current",
			wantErr: false,
		},
		{
			name:    "nested hidden file not allowed",
			input:   "config/.gitignore",
			wantErr: true,
			errMsg:  "cannot start with '.'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err != nil && tt.errMsg != "" {
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("Validate(%q) error = %v, want error containing %q", tt.input, err, tt.errMsg)
				}
			}
		})
	}
}
