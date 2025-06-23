package main

import "fmt"

func main() {
	var userAge int = 35
	var userName string = "Иван"
	var isDeveloper bool = true

	fmt.Println("Имя пользователя:", userName)
	fmt.Println("Возраст:", userAge)
	fmt.Println("Является ли разработчиком:", isDeveloper)

	accountBalance := 10500.75 // float64 будет выведен автоматически
	fmt.Println("Баланс счета:", accountBalance)
}
