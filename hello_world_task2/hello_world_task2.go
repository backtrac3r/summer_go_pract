package main

import (
	"fmt"
	"time"
)

func main() {
	var name = "Алексей"

	currentDate := time.Now()

	fmt.Println("Привет, меня зовут", name)
	fmt.Println("Текущая дата:", currentDate.Format("2006-01-02"))
}
