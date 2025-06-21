package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// Исходная строка
	str := "Go - это мощный язык программирования!"

	// 1. Подсчет количества символов (рун)
	charCount := utf8.RuneCountInString(str)
	fmt.Printf("1. Количество символов: %d\n", charCount)

	// 2. Поиск подстроки
	substr := "мощный"
	index := strings.Index(str, substr)
	if index != -1 {
		fmt.Printf("2. Подстрока '%s' найдена на позиции %d\n", substr, index)
	} else {
		fmt.Printf("2. Подстрока '%s' не найдена\n", substr)
	}

	// 3. Изменение регистра
	upper := strings.ToUpper(str)
	lower := strings.ToLower(str)
	fmt.Println("3. В верхнем регистре:", upper)
	fmt.Println("   В нижнем регистре:", lower)

	// 4. Дополнительно: замена подстроки
	newStr := strings.Replace(str, "мощный", "эффективный", 1)
	fmt.Println("4. После замены:", newStr)
}
