module banner

import  "fmt"


// Banner displays text to stdout
func Banner(text string) {
	fmt.Println("#")
	fmt.Printf("# %s", text)
	fmt.Println("#")
}

