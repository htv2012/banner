package banner

import (
	"fmt"
)

// Create constructs a new string which can be printed as a banner
func Create(text string) string {
	return fmt.Sprintf("\n#\n# %s\n#", text)
}

// Print displays text to stdout
func Banner(text string) {
	fmt.Println(Create(text))
}
