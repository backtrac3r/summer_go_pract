package main

import "fmt"

func main() {
	var myInt int = -100
	var myUint uint = 100
	var myByte byte = 'a'

	var myFloat32 float32 = 3.14
	var myFloat64 float64 = 3.14159265359

	var myComplex complex128 = complex(2, 3)

	var myString string = "Это строка в Go"

	var myBool bool = true

	fmt.Printf("Тип: %T, Значение: %v\n", myInt, myInt)
	fmt.Printf("Тип: %T, Значение: %v\n", myUint, myUint)
	fmt.Printf("Тип: %T, Значение: %c (как символ), %v (как число)\n", myByte, myByte, myByte)
	fmt.Printf("Тип: %T, Значение: %v\n", myFloat32, myFloat32)
	fmt.Printf("Тип: %T, Значение: %v\n", myFloat64, myFloat64)
	fmt.Printf("Тип: %T, Значение: %v\n", myComplex, myComplex)
	fmt.Printf("Тип: %T, Значение: %q\n", myString, myString)
	fmt.Printf("Тип: %T, Значение: %v\n", myBool, myBool)
}
