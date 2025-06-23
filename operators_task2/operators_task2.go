package main

import "fmt"

func main() {
	year := 2024

	fmt.Printf("Проверка, является ли %d год високосным.\n", year)

	isLeap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)

	if isLeap {
		fmt.Printf("%d - високосный год.\n", year)
	} else {
		fmt.Printf("%d - не високосный год.\n", year)
	}
}
