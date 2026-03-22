package banner

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	input := "Test Message"
	expected := "#\n# Test Message\n#\n"
	actual := Create(input)

	if actual != expected {
		t.Errorf("actual %q differ from expected %q", actual, expected)
	}
}

func TestPrint(t *testing.T) {
	// 1. Save the original stdout and restore it when the test finishes
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	// 2. Create a pipe to capture the output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// 3. Run the function
	input := "Hello Go"
	Print(input)

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
