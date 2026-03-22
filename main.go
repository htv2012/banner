package banner

import (
	"fmt"
)

// Create constructs a new string which can be printed as a banner
func Create(text string) string {
	return fmt.Sprintf("#\n# %s\n#\n", text)
}

// Print displays text to stdout
func Print(text string) {
	fmt.Println(Create(text))
}
