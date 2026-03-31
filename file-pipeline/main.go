package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: The Interface
// ═══════════════════════════════════════════

// ---------- RULE 1: ALL CAPS → Title Case ----------
func allCapsToTitle(line string) string {
	if line == strings.ToUpper(line) && line != "" {
		words := strings.Fields(line)
		for i, w := range words {
			words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
		}
		return strings.Join(words, " ")
	}
	return line
}

// ---------- RULE 2: lowercase → UPPERCASE ----------
func lowerToUpper(line string) string {
	if line == strings.ToLower(line) && line != "" {
		return strings.ToUpper(line)
	}
	return line
}

// ---------- RULE 3: Trim ----------
func trim(line string) string {
	return strings.TrimSpace(line)
}

// ---------- RULE 4: Reverse words if contains REVERSE ----------
func reverseIfNeeded(line string) string {
	if strings.Contains(line, "REVERSE") {
		words := strings.Fields(line)
		for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
			words[i], words[j] = words[j], words[i]
		}
		return strings.Join(words, " ")
	}
	return line
}

// ---------- RULE 5: Flag long lines ----------
func flagLong(line string) string {
	if len(line) > 80 {
		return line + " [TRUNCATED]"
	}
	return line
}

// ---------- MAIN ----------
func main() {

	// ---------- ARGUMENT CHECK ----------
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	if input == output {
		fmt.Println("✗ Input and output cannot be the same file.")
		return
	}

	// ---------- OPEN INPUT ----------
	file, err := os.Open(input)
	if err != nil {
		fmt.Printf("✗ File not found: %s\n", input)
		return
	}
	defer file.Close()

	// ---------- CHECK OUTPUT PATH ----------
	if info, err := os.Stat(output); err == nil && info.IsDir() {
		fmt.Println("✗ Cannot write to output: path is a directory, not a file.")
		return
	}

	scanner := bufio.NewScanner(file)

	var processed []string
	linesRead := 0

	for scanner.Scan() {
		line := scanner.Text()
		linesRead++

		// APPLY RULES IN ORDER
		line = allCapsToTitle(line)
		line = lowerToUpper(line)
		line = trim(line)
		line = reverseIfNeeded(line)
		line = flagLong(line)

		processed = append(processed, line)
	}

	if linesRead == 0 {
		fmt.Println("⚠ Input file is empty. Nothing to process.")
	}

	// ---------- WRITE OUTPUT ----------
	out, err := os.Create(output)
	if err != nil {
		fmt.Println("✗ Failed to write output file.")
		return
	}
	defer out.Close()

	writer := bufio.NewWriter(out)

	// HEADER
	writer.WriteString("SENTINEL INTERFACE REPORT\n")
	writer.WriteString("─────────────────────────\n")

	// NUMBERING (1, 2, 3...)
	for i, line := range processed {
		writer.WriteString(fmt.Sprintf("%d. %s\n", i+1, line))
	}

	// SUMMARY BLOCK
	writer.WriteString("─────────────────────────\n")
	writer.WriteString("SUMMARY\n")
	writer.WriteString(fmt.Sprintf("Lines read : %d\n", linesRead))
	writer.WriteString(fmt.Sprintf("Lines written : %d\n", len(processed)))
	writer.WriteString("Lines removed : 0\n")

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	writer.WriteString(fmt.Sprintf("Processed at : %s\n", timestamp))

	writer.Flush()

	// ---------- TERMINAL OUTPUT ----------
	fmt.Printf("✦ Lines read : %d\n", linesRead)
	fmt.Printf("✦ Lines written : %d\n", len(processed))
	fmt.Println("✦ Lines removed : 0")
	fmt.Println("✦ Rules applied : CAPS→Title, lower→UPPER, Trim, Reverse(REVERSE), Flag long")
	fmt.Printf("✦ Processed at : %s\n", timestamp)
}
