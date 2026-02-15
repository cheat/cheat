package installer

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrompt(t *testing.T) {
	// Save original stdin/stdout
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	tests := []struct {
		name       string
		prompt     string
		input      string
		defaultVal bool
		want       bool
		wantErr    bool
		wantPrompt string
	}{
		{
			name:       "answer yes",
			prompt:     "Continue?",
			input:      "y\n",
			defaultVal: false,
			want:       true,
			wantPrompt: "Continue?: ",
		},
		{
			name:       "answer yes with uppercase",
			prompt:     "Continue?",
			input:      "Y\n",
			defaultVal: false,
			want:       true,
			wantPrompt: "Continue?: ",
		},
		{
			name:       "answer yes with spaces",
			prompt:     "Continue?",
			input:      "  y  \n",
			defaultVal: false,
			want:       true,
			wantPrompt: "Continue?: ",
		},
		{
			name:       "answer no",
			prompt:     "Continue?",
			input:      "n\n",
			defaultVal: true,
			want:       false,
			wantPrompt: "Continue?: ",
		},
		{
			name:       "answer no with any text",
			prompt:     "Continue?",
			input:      "anything\n",
			defaultVal: true,
			want:       false,
			wantPrompt: "Continue?: ",
		},
		{
			name:       "empty answer uses default true",
			prompt:     "Continue?",
			input:      "\n",
			defaultVal: true,
			want:       true,
			wantPrompt: "Continue?: ",
		},
		{
			name:       "empty answer uses default false",
			prompt:     "Continue?",
			input:      "\n",
			defaultVal: false,
			want:       false,
			wantPrompt: "Continue?: ",
		},
		{
			name:       "whitespace answer uses default",
			prompt:     "Continue?",
			input:      "   \n",
			defaultVal: true,
			want:       true,
			wantPrompt: "Continue?: ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a pipe for stdin
			r, w, _ := os.Pipe()
			os.Stdin = r

			// Create a pipe for stdout to capture the prompt
			rOut, wOut, _ := os.Pipe()
			os.Stdout = wOut

			// Write input to stdin
			go func() {
				defer w.Close()
				io.WriteString(w, tt.input)
			}()

			// Call the function
			got, err := Prompt(tt.prompt, tt.defaultVal)

			// Close stdout write end and read the prompt
			wOut.Close()
			var buf bytes.Buffer
			io.Copy(&buf, rOut)

			// Check error
			if (err != nil) != tt.wantErr {
				t.Errorf("Prompt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check result
			if got != tt.want {
				t.Errorf("Prompt() = %v, want %v", got, tt.want)
			}

			// Check that prompt was displayed correctly
			if buf.String() != tt.wantPrompt {
				t.Errorf("Prompt display = %q, want %q", buf.String(), tt.wantPrompt)
			}
		})
	}
}

func TestPromptError(t *testing.T) {
	// Save original stdin
	oldStdin := os.Stdin
	defer func() {
		os.Stdin = oldStdin
	}()

	// Create a pipe and close it immediately to simulate read error
	r, w, _ := os.Pipe()
	os.Stdin = r
	r.Close()
	w.Close()

	// This should cause a read error
	_, err := Prompt("Test?", false)
	if err == nil {
		t.Error("expected error when reading from closed stdin, got nil")
	}
	if !strings.Contains(err.Error(), "failed to parse input") {
		t.Errorf("expected 'failed to parse input' error, got: %v", err)
	}
}

// TestPromptIntegration provides a simple integration test
func TestPromptIntegration(t *testing.T) {
	// This demonstrates how the prompt would be used in practice
	// It's skipped by default since it requires actual user input
	if os.Getenv("TEST_INTERACTIVE") != "1" {
		t.Skip("Skipping interactive test - set TEST_INTERACTIVE=1 to run")
	}

	fmt.Println("\n=== Interactive Prompt Test ===")
	fmt.Println("You will be prompted to answer a question.")
	fmt.Println("Try different inputs: y, n, Y, N, empty (just press Enter)")

	result, err := Prompt("Would you like to continue? [Y/n]", true)
	if err != nil {
		t.Fatalf("Prompt failed: %v", err)
	}

	fmt.Printf("You answered: %v\n", result)
}
