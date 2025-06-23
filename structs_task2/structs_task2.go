package main

import "fmt"

type Engine struct {
	Volume     float64
	Horsepower int
}

type Car struct {
	Brand  string
	Model  string
	Color  string
	Engine Engine
}

func main() {
	myCar := Car{
		Brand: "Lada",
		Model: "Vesta",
		Color: "Белый",
		Engine: Engine{
			Volume:     1.6,
			Horsepower: 106,
		},
	}

	fmt.Println("--- Информация об автомобиле ---")
	fmt.Printf("Марка: %s\n", myCar.Brand)
	fmt.Printf("Модель: %s\n", myCar.Model)
	fmt.Printf("Цвет: %s\n", myCar.Color)

	fmt.Println("--- Характеристики двигателя ---")
	fmt.Printf("Объем: %.1f л\n", myCar.Engine.Volume)
	fmt.Printf("Мощность: %d л.с.\n", myCar.Engine.Horsepower)
}
