package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main(){
	fmt.Print("Digite uma frase\n")
	reader := bufio.NewReader(os.Stdin)
	frase, _ := reader.ReadString('\n')
	frase = strings.ToLower(frase)
	frase = strings.ReplaceAll(frase, " ", "")  

	if isPangrama(frase) {
		fmt.Println("A frase é um pangrama!")
	} else {
		fmt.Println("A frase NÃO é um pangrama.")
	}
}
	func isPangrama(s string) bool {
		letras := make(map[rune]bool)
	
		for _, r := range s {
			if unicode.IsLetter(r) {
				letras[r] = true
			}
		}
	
		for r := 'a'; r <= 'z'; r++ {
			if !letras[r] {
				return false
			}
		}
	
		return true
}