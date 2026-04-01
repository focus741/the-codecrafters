package main

import (
	"os"
	"strings"

	"sentinel-reborn/theInterface"
)

func pipe(s string) string {
	s = theInterface.BinToDec(s)
	s = theInterface.ReplaceHexWithDec(s)
	s = theInterface.Up(s)
	if strings.Contains(s, "(cap,") {
		s = theInterface.CapN(s)
	}
	if strings.Contains(s, "(low,") {
		s = theInterface.LowN(s)
	}
	s = theInterface.UpN(s)
	s = theInterface.Punctuations(s)
	s = theInterface.FixQuotes(s)
	s = theInterface.Vol(s)
	return s
}

func main() {
	in, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.ReplaceAll(string(in), "\r\n", "\n"), "\n")
	for i := range lines {
		lines[i] = pipe(lines[i])
	}
	os.WriteFile("output.txt", []byte(strings.Join(lines, "\n")), 0644)
}
