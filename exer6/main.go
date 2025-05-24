package main

import "fmt"

func main() {
	listA := []int{1, 2, 3, 4, 5}
	listB := []int{2, 3, 4}

	resultado := comparaListas(listA, listB)
	fmt.Println("Resultado:", resultado)
}

func comparaListas(listA, listB []int) string {
	if len(listA) == 0 && len(listB) == 0 {
		return "iguais"
	}

	if listasIguais(listA, listB) {
		return "iguais"
	}

	if contemSublista(listA, listB) {
		return "A é superlista"
	}

	if contemSublista(listB, listA) {
		return "A é sublista"
	}

	return "desiguais"
}

func listasIguais(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func contemSublista(maior, menor []int) bool {
	if len(menor) == 0 {
		return true
	}
	if len(maior) < len(menor) {
		return false
	}

	for i := 0; i <= len(maior)-len(menor); i++ {
		match := true
		for j := range menor {
			if maior[i+j] != menor[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
