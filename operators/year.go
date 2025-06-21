package main

import "fmt"

func isLeapYear(year int) bool {
	// Год високосный, если:
	// 1. Делится на 4, но не делится на 100 ИЛИ
	// 2. Делится на 400
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func main() {
	var year int

	fmt.Print("Введите год: ")
	fmt.Scanln(&year)

	if isLeapYear(year) {
		fmt.Printf("%d год - високосный\n", year)
	} else {
		fmt.Printf("%d год - не високосный\n", year)
	}

	// Дополнительно: вывод нескольких ближайших високосных годов
	fmt.Println("\nБлижайшие високосные годы:")
	count := 0
	for y := year - 5; y <= year+5; y++ {
		if isLeapYear(y) {
			fmt.Print(y, " ")
			count++
		}
	}
	if count == 0 {
		fmt.Println("Не найдено")
	}
}
