package display

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/cheat/cheat/internal/config"
)

// TestWriteToPager tests the writeToPager function
func TestWriteToPager(t *testing.T) {
	// Skip these tests in CI/CD environments where interactive commands might not work
	if os.Getenv("CI") != "" {
		t.Skip("Skipping pager tests in CI environment")
	}

	// Note: We can't easily test os.Exit calls, so we focus on testing writeToPager
	// which contains the core logic

	t.Run("successful pager execution", func(t *testing.T) {
		// Save original stdout
		oldStdout := os.Stdout
		defer func() {
			os.Stdout = oldStdout
		}()

		// Create pipe for capturing output
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Use 'cat' as a simple pager that just outputs input
		conf := config.Config{
			Pager: "cat",
		}

		// This will call os.Exit on error, so we need to be careful
		// We're using 'cat' which should always succeed
		input := "Test output\n"

		// Run in a goroutine to avoid blocking
		done := make(chan bool)
		go func() {
			writeToPager(input, conf)
			done <- true
		}()

		// Wait for completion or timeout
		select {
		case <-done:
			// Success
		}

		// Close write end and read output
		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)

		// Verify output
		if buf.String() != input {
			t.Errorf("expected output %q, got %q", input, buf.String())
		}
	})

	t.Run("pager with arguments", func(t *testing.T) {
		// Save original stdout
		oldStdout := os.Stdout
		defer func() {
			os.Stdout = oldStdout
		}()

		// Create pipe for capturing output
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Use 'cat' with '-A' flag (shows non-printing characters)
		conf := config.Config{
			Pager: "cat -A",
		}

		input := "Test\toutput\n"

		// Run in a goroutine
		done := make(chan bool)
		go func() {
			writeToPager(input, conf)
			done <- true
		}()

		// Wait for completion
		select {
		case <-done:
			// Success
		}

		// Close write end and read output
		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)

		// cat -A shows tabs as ^I and line endings as $
		expected := "Test^Ioutput$\n"
		if buf.String() != expected {
			t.Errorf("expected output %q, got %q", expected, buf.String())
		}
	})
}

// TestWriteToPagerError tests error handling in writeToPager
func TestWriteToPagerError(t *testing.T) {
	if os.Getenv("TEST_PAGER_ERROR_SUBPROCESS") == "1" {
		// This is the subprocess - run the actual test
		conf := config.Config{Pager: "/nonexistent/command"}
		writeToPager("test", conf)
		return
	}

	// Run test in subprocess to handle os.Exit
	cmd := exec.Command(os.Args[0], "-test.run=^TestWriteToPagerError$")
	cmd.Env = append(os.Environ(), "TEST_PAGER_ERROR_SUBPROCESS=1")

	output, err := cmd.CombinedOutput()

	// Should exit with error
	if err == nil {
		t.Error("expected process to exit with error")
	}

	// Should contain error message
	if !strings.Contains(string(output), "failed to write to pager") {
		t.Errorf("expected error message about pager failure, got %q", string(output))
	}
}
