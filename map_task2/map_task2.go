package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go это язык go это просто go это быстро"

	text = strings.ToLower(text)

	words := strings.Split(text, " ")

	wordFrequency := make(map[string]int)

	for _, word := range words {
		wordFrequency[word]++
	}

	fmt.Println("Подсчет частоты слов в тексте:")
	fmt.Printf("Текст: \"%s\"\n\n", text)
	fmt.Println("Результат:")

	for word, count := range wordFrequency {
		fmt.Printf("Слово \"%s\" встречается %d раз(а)\n", word, count)
	}
}
