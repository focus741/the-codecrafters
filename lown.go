package theinterface

import (
	"fmt"
	"strconv"
	"strings"
)

// I was in the room (low, 2), yes

func main() {
	fmt.Println(lowN("I was in THE THIRD ROOM FLOOR (low, 2), yes"))
}

func lowN(text string) string {
	opening_Marker := strings.Index(text, "(")
	closing_Marker := strings.Index(text, ")")
	marker := text[opening_Marker : closing_Marker+1]

	beforeMarker := text[:opening_Marker]
	words := strings.Fields(beforeMarker)

	num := 0
	for i := 0; i < len(marker)-1; i++ {
		if strings.ContainsAny(string(marker[i]), "123456789") {
			num, _ = strconv.Atoi(string(marker[i]))
		}
	}

	for i := len(words) - num; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}

	words = append(words, text[closing_Marker+1:])

	return strings.Join(words, " ")
}
