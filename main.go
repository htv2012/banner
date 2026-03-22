package banner

import (
	"fmt"
	"io"
)

func Fbanner(w io.Writer, text string) {
	fmt.Fprintln(w, "#")
	fmt.Fprintf(w, "# %s\n", text)
	fmt.Fprintln(w, "#")
}

// Banner displays text to stdout
func Banner(text string) {
	fmt.Println("#")
	fmt.Printf("# %s", text)
	fmt.Println("#")
}
