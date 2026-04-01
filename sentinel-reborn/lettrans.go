package main

import (
	"fmt"
	"strconv"
	"strings"
)

func up(text string) string {
	newtext := []string{}
	txt := strings.Fields(text)
	for i := 0; i < len(txt); i++ {
		if strings.HasPrefix(txt[i], "(up,") {
			p := strings.TrimSuffix(txt[i+1], ")")
			m, err := strconv.Atoi(p)

			if err == nil {
				if m > len(newtext) {
					m = len(newtext)
				}
				for j := 1; j <= m; j++ {
					target := len(newtext) - j
					newtext[target] = strings.ToUpper(newtext[target])
				}
				i++
				continue
			}

		}
		newtext = append(newtext, txt[i])

	}
	return strings.Join(newtext, " ")

}
func main() {
	fmt.Println(up("this is a programing language (up, 4) go"))
}
