package main

import "fmt"

// Функция принимает переменное количество строк и возвращает самую длинную
func longestString(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}

	longest := strs[0]
	for _, s := range strs[1:] {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	// Пример использования
	fmt.Println(longestString("apple", "banana", "pear"))    // "banana"
	fmt.Println(longestString("cat", "dog", "bird", "fish")) // "bird"
	fmt.Println(longestString("hello", "world"))             // "hello"
	fmt.Println(longestString("single"))                     // "single"
	fmt.Println(longestString())                             // ""
}
