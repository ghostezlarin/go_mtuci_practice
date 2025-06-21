package main

import "fmt"

func main() {
	// Объявление переменных разных типов
	var integerVar int = 42
	var stringVar string = "Привет, Go!"
	var boolVar bool = true
	var floatVar float64 = 3.1415

	// Вывод значений переменных
	fmt.Println("Целое число:", integerVar)
	fmt.Println("Строка:", stringVar)
	fmt.Println("Булево значение:", boolVar)
	fmt.Println("Число с плавающей точкой:", floatVar)

	// Краткое объявление переменных
	shortInt := 100
	shortString := "Краткая запись"
	shortBool := false

	fmt.Println("\nКраткое объявление:")
	fmt.Println(shortInt, shortString, shortBool)
}
