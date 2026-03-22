package banner

import (
	"fmt"
)

func Create(text string) string {
	return fmt.Sprintf("#\n# %s\n#\n", text)
}

// Print displays text to stdout
func Print(text string) {
	fmt.Println(Create(text))
}
