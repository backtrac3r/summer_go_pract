package main

import "fmt"

type Transport interface {
	Move()
	Stop()
}

type Car struct {
	Model string
}

func (c Car) Move() {
	fmt.Printf("Автомобиль %s едет по дороге.\n", c.Model)
}

func (c Car) Stop() {
	fmt.Printf("Автомобиль %s останавливается.\n", c.Model)
}

type Bicycle struct {
	Type string
}

func (b Bicycle) Move() {
	fmt.Printf("%s велосипед крутит педали.\n", b.Type)
}
func (b Bicycle) Stop() {
	fmt.Printf("%s велосипед тормозит.\n", b.Type)
}

func demonstrateTransport(t Transport) {
	fmt.Printf("--- Демонстрация для [%T] ---\n", t)
	t.Move()
	t.Stop()
	fmt.Println()
}

func main() {
	myCar := Car{Model: "Tesla Model 3"}
	myBike := Bicycle{Type: "Горный"}

	demonstrateTransport(myCar)
	demonstrateTransport(myBike)

	fmt.Println("--- Демонстрация через срез интерфейсов ---")
	vehicles := []Transport{myCar, myBike}
	for _, v := range vehicles {
		v.Move()
	}
}
