package main

import "fmt"

var grades = make(map[string]int)

func addGrade(name string, grade int) {
	grades[name] = grade
	fmt.Printf("Добавлена/обновлена оценка для студента %s: %d\n", name, grade)
}

func findGrade(name string) (int, bool) {
	grade, found := grades[name]
	return grade, found
}

func deleteGrade(name string) {
	if _, found := grades[name]; found {
		delete(grades, name)
		fmt.Printf("Студент %s удален.\n", name)
	} else {
		fmt.Printf("Студент %s не найден.\n", name)
	}
}

func main() {
	fmt.Println("--- Управление оценками студентов ---")

	addGrade("Иванов", 5)
	addGrade("Петров", 4)
	addGrade("Сидоров", 3)
	fmt.Println("Текущие оценки:", grades)
	fmt.Println("---")

	nameToFind := "Петров"
	if grade, found := findGrade(nameToFind); found {
		fmt.Printf("Оценка студента %s: %d\n", nameToFind, grade)
	} else {
		fmt.Printf("Студент %s не найден.\n", nameToFind)
	}
	fmt.Println("---")

	deleteGrade("Сидоров")
	fmt.Println("Текущие оценки:", grades)
	fmt.Println("---")

	deleteGrade("Кузнецов")
}
