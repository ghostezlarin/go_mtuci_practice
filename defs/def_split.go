package main

import (
	"fmt"
	"sort"
)

// Поиск элемента в срезе
func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Сортировка среза (по возрастанию)
func sortSlice(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Ints(sorted)
	return sorted
}

// Фильтрация среза (оставляем только четные числа)
func filterSlice(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	numbers := []int{5, 2, 8, 3, 1, 9, 4, 6}

	// Поиск элемента
	fmt.Println("Contains 8:", contains(numbers, 8)) // true
	fmt.Println("Contains 7:", contains(numbers, 7)) // false

	// Сортировка
	sorted := sortSlice(numbers)
	fmt.Println("Sorted:", sorted) // [1 2 3 4 5 6 8 9]

	// Фильтрация
	filtered := filterSlice(numbers)
	fmt.Println("Filtered (even):", filtered) // [2 8 4 6]
}
