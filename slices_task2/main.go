package main

import "fmt"

func removeAtIndex(slice []int, index int) ([]int, bool) {
	if index < 0 || index >= len(slice) {
		return slice, false
	}
	return append(slice[:index], slice[index+1:]...), true
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Исходный срез:", numbers)

	indexToRemove := 2
	newNumbers, ok := removeAtIndex(numbers, indexToRemove)
	if ok {
		fmt.Printf("Срез после удаления элемента с индексом %d: %v\n", indexToRemove, newNumbers)
	} else {
		fmt.Printf("Не удалось удалить элемент: индекс %d вне диапазона.\n", indexToRemove)
	}

	indexToRemove = 10
	newNumbers, ok = removeAtIndex(numbers, indexToRemove)
	if ok {
		fmt.Printf("Срез после удаления элемента с индексом %d: %v\n", indexToRemove, newNumbers)
	} else {
		fmt.Printf("Не удалось удалить элемент: индекс %d вне диапазона.\n", indexToRemove)
	}
}
