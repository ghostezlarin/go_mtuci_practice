package main

import "fmt"

// Интерфейс транспортного средства
type Transport interface {
	Move(speed int) string
	Stop() string
}

// Автомобиль
type Car struct {
	Model string
}

func (c Car) Move(speed int) string {
	return fmt.Sprintf("%s едет со скоростью %d км/ч", c.Model, speed)
}

func (c Car) Stop() string {
	return fmt.Sprintf("%s остановился", c.Model)
}

// Велосипед
type Bicycle struct {
	Type string
}

func (b Bicycle) Move(speed int) string {
	return fmt.Sprintf("%s велосипед движется со скоростью %d км/ч", b.Type, speed)
}

func (b Bicycle) Stop() string {
	return fmt.Sprintf("%s велосипед остановился", b.Type)
}

// Поезд
type Train struct {
	Number int
}

func (t Train) Move(speed int) string {
	return fmt.Sprintf("Поезд №%d набирает скорость %d км/ч", t.Number, speed)
}

func (t Train) Stop() string {
	return fmt.Sprintf("Поезд №%d прибыл на станцию", t.Number)
}

// Функция для тестирования транспорта
func testTransport(t Transport, speed int) {
	fmt.Println(t.Move(speed))
	fmt.Println(t.Stop())
}

func main() {
	// Создаем транспортные средства
	car := Car{Model: "Toyota Camry"}
	bike := Bicycle{Type: "Горный"}
	train := Train{Number: 123}

	// Тестируем их
	fmt.Println("--- Тест автомобиля ---")
	testTransport(car, 60)

	fmt.Println("\n--- Тест велосипеда ---")
	testTransport(bike, 20)

	fmt.Println("\n--- Тест поезда ---")
	testTransport(train, 100)

	// Можно также использовать интерфейсный тип
	var vehicle Transport = car
	fmt.Println("\n--- Использование интерфейса ---")
	fmt.Println(vehicle.Move(70))
}
