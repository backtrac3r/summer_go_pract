package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Go — это высокоэффективный, лаконичный и современный язык программирования"

	fmt.Println("Исходное предложение:")
	fmt.Println(sentence)
	fmt.Println("\nСлова из предложения:")

	words := strings.Split(sentence, " ")

	for i, word := range words {
		fmt.Printf("%d: %s\n", i+1, word)
	}
}
