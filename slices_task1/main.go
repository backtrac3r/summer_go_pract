package main

import "fmt"

func main() {
	var fruits []string
	fmt.Printf("Начальный срез: %v, Длина: %d, Емкость: %d\n", fruits, len(fruits), cap(fruits))

	fruits = append(fruits, "Яблоко")
	fmt.Printf("После добавления 'Яблоко': %v, Длина: %d, Емкость: %d\n", fruits, len(fruits), cap(fruits))

	fruits = append(fruits, "Банан", "Апельсин")
	fmt.Printf("После добавления 'Банан', 'Апельсин': %v, Длина: %d, Емкость: %d\n", fruits, len(fruits), cap(fruits))

	fmt.Println("\nВсе элементы среза 'fruits':")
	for i, fruit := range fruits {
		fmt.Printf("%d. %s\n", i+1, fruit)
	}
}
