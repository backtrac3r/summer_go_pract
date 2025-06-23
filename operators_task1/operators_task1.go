package main

import "fmt"

func main() {
	a := 20
	b := 8

	fmt.Printf("Демонстрация арифметических операторов с числами a = %d и b = %d\n", a, b)

	sum := a + b
	fmt.Printf("Сложение (a + b): \t%d\n", sum)

	difference := a - b
	fmt.Printf("Вычитание (a - b): \t%d\n", difference)

	product := a * b
	fmt.Printf("Умножение (a * b): \t%d\n", product)

	quotient := a / b
	fmt.Printf("Деление (a / b): \t%d\n", quotient)

	remainder := a % b
	fmt.Printf("Остаток (a %% b): \t%d\n", remainder)
}
