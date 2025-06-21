package main

import "fmt"

// Определение структуры Engine
type Engine struct {
	Type    string
	Power   int     // л.с.
	Volume  float64 // литры
	IsTurbo bool
}

// Метод для вывода информации о двигателе
func (e Engine) PrintSpecs() {
	turbo := ""
	if e.IsTurbo {
		turbo = " (турбо)"
	}
	fmt.Printf("Двигатель: %s, %.1f л, %d л.с.%s\n", e.Type, e.Volume, e.Power, turbo)
}

// Определение структуры Car с вложенной структурой Engine
type Car struct {
	Make    string
	Model   string
	Year    int
	Engine  Engine // Вложенная структура
	Mileage int    // пробег
}

// Функция для создания нового автомобиля
func NewCar(make, model string, year, mileage int, engine Engine) Car {
	return Car{
		Make:    make,
		Model:   model,
		Year:    year,
		Engine:  engine,
		Mileage: mileage,
	}
}

// Метод для вывода информации об автомобиле
func (c Car) PrintInfo() {
	fmt.Printf("\n%s %s %d года\n", c.Make, c.Model, c.Year)
	fmt.Printf("Пробег: %d км\n", c.Mileage)
	c.Engine.PrintSpecs()
}

// Метод для обновления пробега
func (c *Car) UpdateMileage(newMileage int) {
	c.Mileage = newMileage
	fmt.Printf("Пробег %s %s обновлен: %d км\n", c.Make, c.Model, c.Mileage)
}

func main() {
	// Создаем двигатель
	engine := Engine{
		Type:    "V6",
		Power:   300,
		Volume:  3.0,
		IsTurbo: true,
	}

	// Создаем автомобиль
	car := NewCar("Toyota", "Camry", 2022, 15000, engine)

	// Выводим информацию об автомобиле
	car.PrintInfo()

	// Обновляем пробег
	car.UpdateMileage(17500)

	// Создаем и выводим другой автомобиль
	fmt.Println("\nДругой автомобиль:")
	engine2 := Engine{"I4", 150, 2.0, false}
	car2 := NewCar("Honda", "Civic", 2020, 45000, engine2)
	car2.PrintInfo()
}
