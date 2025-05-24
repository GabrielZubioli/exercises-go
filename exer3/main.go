package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Digite uma palavra, frase ou número:\n")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")  

	palindromo := strings.ToLower(input)
	invertido := reverse(palindromo)

	fmt.Println("Invertido:", invertido)

	if palindromo == invertido {
		fmt.Println("É um palíndromo!")
	} else {
		fmt.Println("Não é um palíndromo.")
	}
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
