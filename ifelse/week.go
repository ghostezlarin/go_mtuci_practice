package main

import (
	"fmt"
	"os"
)

func main() {
	var dayNumber int

	// Ввод номера дня
	fmt.Print("Введите номер дня недели (1-7): ")
	_, err := fmt.Scanln(&dayNumber)
	if err != nil {
		fmt.Println("Ошибка ввода числа")
		os.Exit(1)
	}

	// Определение дня недели
	var dayName string

	switch dayNumber {
	case 1:
		dayName = "Понедельник"
	case 2:
		dayName = "Вторник"
	case 3:
		dayName = "Среда"
	case 4:
		dayName = "Четверг"
	case 5:
		dayName = "Пятница"
	case 6:
		dayName = "Суббота"
	case 7:
		dayName = "Воскресенье"
	default:
		fmt.Println("Ошибка: номер должен быть от 1 до 7")
		os.Exit(1)
	}

	fmt.Printf("День недели: %s\n", dayName)
}
