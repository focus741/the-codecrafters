package main

import (
	"fmt"
	"strings"
)

func Punctuations(text string) string {

	punc := ".,!?:;"
	textSlice := strings.Fields(text)

	for i := 0; i < len(textSlice); i++ {
		for j := 0; j < len(textSlice[i]); j++ {
			var count int = 0
			if strings.Contains(punc, string((textSlice[i])[j])) {
				count += 1
				if count == len(textSlice[i]) {
					textSlice[i-1] = textSlice[i-1] + textSlice[i]

					textSlice[i] = ""

				} else {
					for k := 0; k < len(punc); k++ {
						if strings.HasPrefix(textSlice[i], string(punc[k])) {
							freq := strings.Count(textSlice[i], string(punc[k]))

							textSlice[i-1] = textSlice[i-1] + (textSlice[i])[:freq]

							textSlice[i] = (textSlice[i])[freq:] // I was sitting over there ,and then BAMM !!
						}
					}
				}
			}
		}
	}

	newTextSlice := []string{}
	for _, word := range textSlice {
		if len(word) != 0 {
			newTextSlice = append(newTextSlice, word)
		}
	}

	newText := strings.Join(newTextSlice, " ")

	return newText

}

func main() {
	fmt.Println(Punctuations("I was sitting over there ,and then BAMM !!"))
	fmt.Println(Punctuations("I was thinking ... You were right"))
}
