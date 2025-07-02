package main

import "fmt"

func passByValue(val int) {
	val = 100
	fmt.Printf("Внутри passByValue, val = %d\n", val)
}

func passByPointer(ptr *int) {
	*ptr = 200
	fmt.Printf("Внутри passByPointer, *ptr = %d\n", *ptr)
}

func main() {
	number := 50
	fmt.Printf("Начальное значение number = %d\n\n", number)

	fmt.Println("1. Вызов passByValue(number):")
	passByValue(number)
	fmt.Printf("После passByValue, number = %d (не изменилось)\n\n", number)

	fmt.Println("2. Вызов passByPointer(&number):")
	passByPointer(&number)
	fmt.Printf("После passByPointer, number = %d (изменилось)\n", number)
}
