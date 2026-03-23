package banner

import (
	"testing"
)

func TestCreate(t *testing.T) {
	input := "Test Message"
	expected := "\n#\n# Test Message\n#"
	actual := Create(input)

	if actual != expected {
		t.Errorf("Values differ:\nactual:   %q\nexpected: %q", actual, expected)
	}
}
