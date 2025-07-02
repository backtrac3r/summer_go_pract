package main

import "fmt"

func main() {
	fmt.Println("Простые числа до 100:")
	limit := 100

	for number := 2; number <= limit; number++ {
		isPrime := true
		for i := 2; i*i <= number; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", number)
		}
	}
	fmt.Println()
}
