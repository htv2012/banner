package banner

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestFbanner(t *testing.T) {
	var buf bytes.Buffer
	input := "Test Message"

	Fbanner(&buf, input)

	got := buf.String()
	want := "#\n# Test Message\n#\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestBanner(t *testing.T) {
	// 1. Save the original stdout and restore it when the test finishes
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	// 2. Create a pipe to capture the output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// 3. Run the function
	input := "Hello Go"
	Banner(input)

	// 4. Close the writer and read the captured output
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	got := buf.String()

	// 5. Assertions
	if !strings.Contains(got, input) {
		t.Errorf("Banner() output = %q, want it to contain %q", got, input)
	}

	// Check for the decorative character
	if !strings.Contains(got, "#") {
		t.Errorf("Banner() output missing '#' decoration")
	}
}
