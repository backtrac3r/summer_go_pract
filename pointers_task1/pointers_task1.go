package main

import "fmt"

func swapValues(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x := 10
	y := 20

	fmt.Printf("До обмена: x = %d, y = %d\n", x, y)

	swapValues(&x, &y)

	fmt.Printf("После обмена: x = %d, y = %d\n", x, y)
}
