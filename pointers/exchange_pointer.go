package main

import "fmt"

// Функция обмена значений через указатели
func swap(a, b *int) {
	// Временная переменная для хранения значения
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x, y := 10, 20

	fmt.Println("До обмена:")
	fmt.Printf("x = %d, y = %d\n", x, y)

	// Передаем указатели на переменные
	swap(&x, &y)

	fmt.Println("После обмена:")
	fmt.Printf("x = %d, y = %d\n", x, y)
}
