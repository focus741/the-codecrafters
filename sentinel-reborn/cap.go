package main

import (
	"fmt"
	"strings"
)

func parseModifier(t string) string {
	t = strings.TrimSpace(t)
	index := strings.Index(t, "(")
	if strings.Contains(t, "(up)") {
		t = strings.ToUpper(t)
	}
	if strings.Contains(t, "(low)") {
		t = strings.ToLower(t)
	}
	if strings.Contains(t, "(cap)") {
		t = strings.Title(t)
	}
	return t[:index]
}


/*func main() {
	fmt.Println(parseModifier("hello (up)"))
	fmt.Println(parseModifier("Onminyi Andrew Okala (low)"))
	fmt.Println(parseModifier("hello andrew (cap)"))
}
*/
