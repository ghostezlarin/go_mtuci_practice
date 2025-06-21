package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := "Go — это язык программирования. Go разработан в Google. Язык Go прост и эффективен."

	// Получаем частоту слов
	wordFrequency := countWords(text)

	// Выводим результаты
	fmt.Println("Частота слов:")
	for word, count := range wordFrequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func countWords(text string) map[string]int {
	// Создаем карту для хранения частоты
	frequency := make(map[string]int)

	// Разбиваем текст на слова
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	// Подсчитываем частоту каждого слова
	for _, word := range words {
		lowerWord := strings.ToLower(word)
		frequency[lowerWord]++
	}

	return frequency
}
