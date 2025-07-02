package main

import "fmt"

func main() {
	var a, b float64
	var operator string

	a = 20
	b = 5
	operator = "*"

	var result float64

	fmt.Printf("Вычисляем: %.2f %s %.2f\n", a, operator, b)

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль!")
			return
		}
		result = a / b
	default:
		fmt.Println("Ошибка: неизвестный оператор!")
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}
