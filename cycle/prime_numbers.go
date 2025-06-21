package main

import "fmt"

func main() {
	fmt.Println("Простые числа до 100:")

	for num := 2; num <= 100; num++ {
		isPrime := true

		// Проверяем, делится ли число на что-то кроме 1 и самого себя
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Print(num, " ")
		}
	}
	fmt.Println() // Переход на новую строку после вывода
}
