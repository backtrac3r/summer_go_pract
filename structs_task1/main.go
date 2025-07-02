package main

import "fmt"

type Student struct {
	Name       string
	Age        int
	Course     int
	AverageGpa float64
}

func printStudentInfo(s Student) {
	fmt.Println("--- Информация о студенте ---")
	fmt.Printf("Имя: %s\n", s.Name)
	fmt.Printf("Возраст: %d\n", s.Age)
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.2f\n", s.AverageGpa)
	fmt.Println("-----------------------------")
}

func promoteStudent(s *Student) {
	s.Course++
	fmt.Printf("%s переведен(а) на %d курс.\n", s.Name, s.Course)
}

func main() {
	student1 := Student{
		Name:       "Елена",
		Age:        20,
		Course:     3,
		AverageGpa: 4.8,
	}

	printStudentInfo(student1)
	promoteStudent(&student1)
	printStudentInfo(student1)
}
