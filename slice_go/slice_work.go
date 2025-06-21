package main

import "fmt"

func main() {
	// 1. Создаем пустой срез строк
	var slice []string

	// 2. Динамически добавляем элементы
	slice = append(slice, "Первый")
	slice = append(slice, "Второй")
	slice = append(slice, "Третий")

	// 3. Добавляем сразу несколько элементов
	slice = append(slice, "Четвертый", "Пятый")

	// 4. Выводим все элементы
	fmt.Println("Срез строк:")
	for i, value := range slice {
		fmt.Printf("Индекс: %d, Значение: %s\n", i, value)
	}

	// 5. Выводим длину и вместимость
	fmt.Printf("\nДлина: %d, Вместимость: %d\n", len(slice), cap(slice))
}
