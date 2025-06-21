package main

import (
	"fmt"
	"math"
)

func main() {
	// Объявление констант
	const pi = math.Pi
	const e = math.E
	const goldenRatio = 1.6180339887

	// Демонстрация использования констант
	fmt.Printf("Значение π (пи): %.4f\n", pi)
	fmt.Printf("Значение e (экспонента): %.4f\n", e)
	fmt.Printf("Золотое сечение: %.4f\n", goldenRatio)

	// Примеры вычислений с константами
	radius := 5.0
	circleArea := pi * math.Pow(radius, 2)
	fmt.Printf("\nПлощадь круга с радиусом %.1f: %.2f\n", radius, circleArea)

	x := 2.0
	expResult := math.Pow(e, x)
	fmt.Printf("e^%.1f = %.4f\n", x, expResult)

	// Константы с типизацией
	const precisionPi float64 = 3.141592653589793
	fmt.Printf("\nπ с высокой точностью: %.15f\n", precisionPi)
}
