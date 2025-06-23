package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name       string
	BirthYear  int
	Course     int
	AverageGpa float64
}

func (s Student) CalculateAge() int {
	return time.Now().Year() - s.BirthYear
}

func (s Student) GetStatus() string {
	if s.AverageGpa >= 4.5 {
		return "Отличник"
	} else if s.AverageGpa >= 3.5 {
		return "Хорошист"
	} else if s.AverageGpa >= 2.5 {
		return "Троечник"
	}
	return "Неуспевающий"
}

func main() {
	student1 := Student{
		Name:       "Анна",
		BirthYear:  2003,
		Course:     3,
		AverageGpa: 4.8,
	}

	fmt.Printf("Студент: %s\n", student1.Name)
	fmt.Printf("Возраст: %d\n", student1.CalculateAge())
	fmt.Printf("Статус: %s\n", student1.GetStatus())
}
