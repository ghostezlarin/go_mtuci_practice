package main

import (
	"fmt"
	"math"
)

// Интерфейс формы
type Shape interface {
	Area() float64
}

// Прямоугольник
type Rectangle struct {
	Width  float64
	Height float64
}

// Реализация метода Area для прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Круг
type Circle struct {
	Radius float64
}

// Реализация метода Area для круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Функция для вывода площади любой фигуры
func printArea(shape Shape) {
	fmt.Printf("Площадь фигуры: %.2f\n", shape.Area())
}

func main() {
	// Создаем фигуры
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	// Используем интерфейс
	printArea(rect)   // Площадь фигуры: 15.00
	printArea(circle) // Площадь фигуры: 50.27

	// Можно также напрямую вызывать методы
	fmt.Printf("Площадь прямоугольника: %.2f\n", rect.Area())
}
