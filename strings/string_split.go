package main

import (
	"fmt"
	"strings"
)

func main() {
	// Исходное предложение
	sentence := "Go создан для эффективной разработки"

	// 1. Разбиение строки на слова
	words := strings.Fields(sentence)

	// 2. Вывод каждого слова на новой строке
	fmt.Println("Разбиение предложения на слова:")
	for i, word := range words {
		fmt.Printf("%d: %s\n", i+1, word)
	}

	// 3. Альтернативный вариант (по другим разделителям)
	complexSentence := "Go,Python;Java C++"
	separators := func(r rune) bool {
		return r == ',' || r == ';' || r == ' '
	}
	words = strings.FieldsFunc(complexSentence, separators)

	fmt.Println("\nРазбиение сложной строки:")
	for _, word := range words {
		fmt.Println(word)
	}
}
