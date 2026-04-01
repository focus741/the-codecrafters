package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func hexToDec(hexStr string) (int64, error) {
	dec, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid hex string: %s", hexStr)
	}
	return dec, nil
}

func replaceHexWithDec(input string) string {
	re := regexp.MustCompile(`(?i)\b([0-9a-f]+) \(hex\)`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		parts := strings.Split(match, " ")
		hex := parts[0]
		dec, err := hexToDec(hex)
		if err != nil {
			return match
		}
		return fmt.Sprintf("%d", dec)
	})
}

func main() {
	inputs := []string{
		"1E (hex) files were added",
		"The value is 1E (hex) and 10 (hex)",
		"No hex here",
		"1a2b3c (hex) is a big number",
	}
	for _, s := range inputs {
		fmt.Println(replaceHexWithDec(s))
	}
}
