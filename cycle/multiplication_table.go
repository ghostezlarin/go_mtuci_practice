package main

import "fmt"

func main() {
	// Запрашиваем у пользователя число
	var number int
	fmt.Print("Введите число от 1 до 10: ")
	fmt.Scanln(&number)

	// Проверяем корректность ввода
	if number < 1 || number > 10 {
		fmt.Println("Ошибка: число должно быть от 1 до 10")
		return
	}

	// Выводим таблицу умножения
	fmt.Printf("\nТаблица умножения для %d:\n", number)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d × %d = %d\n", number, i, number*i)
	}
}
