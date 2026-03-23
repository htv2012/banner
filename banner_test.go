package banner

import (
	"testing"
)

func TestCreate(t *testing.T) {
	input := "Test Message"
	expected := "#\n# Test Message\n#"
	actual := Create(input)

	if actual != expected {
		t.Errorf("actual %q differ from expected %q", actual, expected)
	}
}
