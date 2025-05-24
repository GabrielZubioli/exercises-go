package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"bufio"
	"os"
)

func normalizeText(input string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return strings.ToLower(re.ReplaceAllString(input, ""))
}

func squareCode(input string) string {
	normalized := normalizeText(input)
	length := len(normalized)

	r := int(math.Floor(math.Sqrt(float64(length))))
	c := int(math.Ceil(math.Sqrt(float64(length))))
	if r*c < length {
		r++
	}
	if c-r > 1 {
		r++
	}

	lines := []string{}
	for i := 0; i < length; i += c {
		end := i + c
		if end > length {
			end = length
		}
		line := normalized[i:end]
		if len(line) < c {
			line += strings.Repeat(" ", c-len(line))
		}
		lines = append(lines, line)
	}

	columns := make([]string, c)
	for i := 0; i < c; i++ {
		col := ""
		for _, line := range lines {
			col += string(line[i])
		}
		columns[i] = col
	}

	return strings.Join(columns, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite uma frase:\n")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	encoded := squareCode(text)
	fmt.Println(encoded)
}
