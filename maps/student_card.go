package main

import "fmt"

func main() {
	// Создаем карту для хранения оценок
	grades := make(map[string]int)

	// Добавляем оценки
	addGrade(grades, "Алексей", 5)
	addGrade(grades, "Мария", 4)
	addGrade(grades, "Иван", 3)

	// Поиск оценок
	findGrade(grades, "Мария")
	findGrade(grades, "Петр") // Несуществующий студент

	// Удаление оценки
	removeGrade(grades, "Иван")
	findGrade(grades, "Иван") // Проверка удаления

	// Вывод всех оценок
	fmt.Println("\nВсе оценки:")
	for name, grade := range grades {
		fmt.Printf("%s: %d\n", name, grade)
	}
}

// Функция добавления оценки
func addGrade(grades map[string]int, name string, grade int) {
	grades[name] = grade
	fmt.Printf("Добавлена оценка %d для %s\n", grade, name)
}

// Функция поиска оценки
func findGrade(grades map[string]int, name string) {
	if grade, exists := grades[name]; exists {
		fmt.Printf("Оценка %s: %d\n", name, grade)
	} else {
		fmt.Printf("Студент %s не найден\n", name)
	}
}

// Функция удаления оценки
func removeGrade(grades map[string]int, name string) {
	delete(grades, name)
	fmt.Printf("Оценка для %s удалена\n", name)
}
