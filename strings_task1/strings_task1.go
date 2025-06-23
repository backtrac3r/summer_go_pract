package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	text := "Привет, Мир! Это прекрасный Мир."

	fmt.Printf("Исходная строка: \"%s\"\n\n", text)

	charCount := utf8.RuneCountInString(text)
	fmt.Printf("1. Количество символов (рун): %d\n", charCount)
	fmt.Printf("   Количество байт: %d\n", len(text))

	substring := "Мир"
	contains := strings.Contains(text, substring)
	fmt.Printf("\n2. Строка содержит подстроку \"%s\": %t\n", substring, contains)

	upperText := strings.ToUpper(text)
	lowerText := strings.ToLower(text)
	fmt.Printf("\n3. Строка в верхнем регистре: %s\n", upperText)
	fmt.Printf("   Строка в нижнем регистре: %s\n", lowerText)
}
