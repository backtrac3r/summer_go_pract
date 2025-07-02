package main

import "fmt"

// TODO: from lib
const Pi = 3.14159265359
const E = 2.71828182846

func main() {
	var radius = 10.0

	circumference := 2 * Pi * radius

	area := Pi * radius * radius

	fmt.Println("Математические константы:")
	fmt.Println("Число Пи (π):", Pi)
	fmt.Println("Число Эйлера (e):", E)
	fmt.Println("---")
	fmt.Println("Пример использования для круга с радиусом", radius)
	fmt.Println("Длина окружности:", circumference)
	fmt.Println("Площадь круга:", area)
}
