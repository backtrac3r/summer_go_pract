package main

import (
	"fmt"
	"time"
)

func main() {
	currentDate := time.Now()

	fmt.Println("Привет, меня зовут", "Андрей")
	fmt.Println("Текущая дата:", currentDate.Format("2006-01-02"))
}
