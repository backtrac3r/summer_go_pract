package main

import "fmt"

func main() {
	dayNumber := 4

	var dayName string

	fmt.Printf("Определяем день недели по номеру %d\n", dayNumber)

	switch dayNumber {
	case 1:
		dayName = "Понедельник"
	case 2:
		dayName = "Вторник"
	case 3:
		dayName = "Среда"
	case 4:
		dayName = "Четверг"
	case 5:
		dayName = "Пятница"
	case 6:
		dayName = "Суббота"
	case 7:
		dayName = "Воскресенье"
	default:
		dayName = "Ошибка: неверный номер дня"
	}

	fmt.Println("Результат:", dayName)
}
