package main

import "fmt"

// Функция удаления элемента по индексу (способ с append)
func removeElement(slice []string, index int) []string {
	// Проверка на корректность индекса
	if index < 0 || index >= len(slice) {
		return slice // возвращаем исходный срез, если индекс невалидный
	}

	// Удаляем элемент, объединяя часть до индекса и после
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	// Исходный срез
	fruits := []string{"Яблоко", "Банан", "Апельсин", "Груша", "Киви"}
	fmt.Println("Исходный срез:", fruits)

	// Удаляем элемент с индексом 2 (Апельсин)
	fruits = removeElement(fruits, 2)
	fmt.Println("После удаления:", fruits)

	// Пробуем удалить несуществующий элемент
	fruits = removeElement(fruits, 10)
	fmt.Println("После попытки удалить несуществующий элемент:", fruits)

	// Удаляем первый элемент
	fruits = removeElement(fruits, 0)
	fmt.Println("После удаления первого элемента:", fruits)

	// Удаляем последний элемент
	fruits = removeElement(fruits, len(fruits)-1)
	fmt.Println("После удаления последнего элемента:", fruits)
}
