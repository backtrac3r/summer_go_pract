package main

import (
	"fmt"
	"sort"
)

func findElement(slice []int, element int) int {
	for i, v := range slice {
		if v == element {
			return i
		}
	}
	return -1
}

func sortSlice(slice []int) {
	sort.Ints(slice)
}

func filterEvenNumbers(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	numbers := []int{9, 2, 7, 4, 5, 6, 1, 8, 3}
	fmt.Println("Исходный срез:", numbers)

	elementToFind := 5
	index := findElement(numbers, elementToFind)
	if index != -1 {
		fmt.Printf("Элемент %d найден по индексу %d\n", elementToFind, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", elementToFind)
	}

	evenNumbers := filterEvenNumbers(numbers)
	fmt.Println("Только четные числа:", evenNumbers)

	fmt.Println("Сортировка исходного среза...")
	sortSlice(numbers)
	fmt.Println("Отсортированный срез:", numbers)
}
