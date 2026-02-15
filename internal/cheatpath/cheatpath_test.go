package cheatpath

import (
	"strings"
	"testing"
)

func TestCheatpathValidate(t *testing.T) {
	tests := []struct {
		name      string
		cheatpath Cheatpath
		wantErr   bool
		errMsg    string
	}{
		{
			name: "valid cheatpath",
			cheatpath: Cheatpath{
				Name:     "personal",
				Path:     "/home/user/.config/cheat/personal",
				ReadOnly: false,
				Tags:     []string{"personal"},
			},
			wantErr: false,
		},
		{
			name: "empty name",
			cheatpath: Cheatpath{
				Name:     "",
				Path:     "/home/user/.config/cheat/personal",
				ReadOnly: false,
				Tags:     []string{"personal"},
			},
			wantErr: true,
			errMsg:  "cheatpath name cannot be empty",
		},
		{
			name: "empty path",
			cheatpath: Cheatpath{
				Name:     "personal",
				Path:     "",
				ReadOnly: false,
				Tags:     []string{"personal"},
			},
			wantErr: true,
			errMsg:  "cheatpath path cannot be empty",
		},
		{
			name: "both empty",
			cheatpath: Cheatpath{
				Name:     "",
				Path:     "",
				ReadOnly: true,
				Tags:     nil,
			},
			wantErr: true,
			errMsg:  "cheatpath name cannot be empty",
		},
		{
			name: "minimal valid",
			cheatpath: Cheatpath{
				Name: "x",
				Path: "/",
			},
			wantErr: false,
		},
		{
			name: "with readonly and tags",
			cheatpath: Cheatpath{
				Name:     "community",
				Path:     "/usr/share/cheat",
				ReadOnly: true,
				Tags:     []string{"community", "shared", "readonly"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cheatpath.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("Validate() error = %v, want error containing %q", err, tt.errMsg)
			}
		})
	}
}
